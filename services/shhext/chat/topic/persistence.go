package topic

import (
	"database/sql"
	"strings"
)

type PersistenceService interface {
	AddTopic(identity []byte, secret []byte, installationID string) error
	GetTopic(identity []byte, installationIDs []string) (*TopicResponse, error)
}

type TopicResponse struct {
	secret          []byte
	installationIDs map[string]bool
}

type SQLLitePersistence struct {
	db *sql.DB
}

func NewSQLLitePersistence(db *sql.DB) *SQLLitePersistence {
	return &SQLLitePersistence{db: db}
}

func (s *SQLLitePersistence) AddTopic(identity []byte, secret []byte, installationID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	insertTopicStmt, err := tx.Prepare("INSERT INTO topics(identity, secret) VALUES (?, ?)")
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer insertTopicStmt.Close()

	_, err = insertTopicStmt.Exec(identity, secret)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	insertInstallationIDStmt, err := tx.Prepare("INSERT INTO topic_installation_ids(id, identity_id) VALUES (?, ?)")
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer insertInstallationIDStmt.Close()

	_, err = insertInstallationIDStmt.Exec(installationID, identity)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *SQLLitePersistence) GetTopic(identity []byte, installationIDs []string) (*TopicResponse, error) {
	response := &TopicResponse{
		installationIDs: make(map[string]bool),
	}
	args := make([]interface{}, len(installationIDs)+1)
	args[0] = identity
	for i, installationID := range installationIDs {
		args[i+1] = installationID
	}

	/* #nosec */
	query := `SELECT secret, id
	          FROM topics t
		  JOIN
		  topic_installation_ids tid
		  ON t.identity = tid.identity_id
		  WHERE
		  t.identity = ?
		  AND
		  tid.id IN (?` + strings.Repeat(",?", len(installationIDs)-1) + `)`

	rows, err := s.db.Query(query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		var installationID string
		var secret []byte
		err = rows.Scan(&secret, &installationID)
		if err != nil {
			return nil, err
		}

		response.secret = secret
		response.installationIDs[installationID] = true
	}

	return response, nil
}
