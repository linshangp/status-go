// Code generated by MockGen. DO NOT EDIT.
// Source: services/status/service.go

// Package status is a generated GoMock package.
package status

import (
	ecdsa "crypto/ecdsa"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	account "github.com/status-im/status-go/account"
	"github.com/status-im/status-go/account/generator"
	types "github.com/status-im/status-go/eth-node/types"
)

// MockWhisperService is a mock of WhisperService interface
type MockWhisperService struct {
	ctrl     *gomock.Controller
	recorder *MockWhisperServiceMockRecorder
}

// MockWhisperServiceMockRecorder is the mock recorder for MockWhisperService
type MockWhisperServiceMockRecorder struct {
	mock *MockWhisperService
}

// NewMockWhisperService creates a new mock instance
func NewMockWhisperService(ctrl *gomock.Controller) *MockWhisperService {
	mock := &MockWhisperService{ctrl: ctrl}
	mock.recorder = &MockWhisperServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWhisperService) EXPECT() *MockWhisperServiceMockRecorder {
	return m.recorder
}

// AddKeyPair mocks base method
func (m *MockWhisperService) AddKeyPair(key *ecdsa.PrivateKey) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKeyPair", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddKeyPair indicates an expected call of AddKeyPair
func (mr *MockWhisperServiceMockRecorder) AddKeyPair(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKeyPair", reflect.TypeOf((*MockWhisperService)(nil).AddKeyPair), key)
}

// MockAccountManager is a mock of AccountManager interface
type MockAccountManager struct {
	ctrl     *gomock.Controller
	recorder *MockAccountManagerMockRecorder
}

// MockAccountManagerMockRecorder is the mock recorder for MockAccountManager
type MockAccountManagerMockRecorder struct {
	mock *MockAccountManager
}

// NewMockAccountManager creates a new mock instance
func NewMockAccountManager(ctrl *gomock.Controller) *MockAccountManager {
	mock := &MockAccountManager{ctrl: ctrl}
	mock.recorder = &MockAccountManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountManager) EXPECT() *MockAccountManagerMockRecorder {
	return m.recorder
}

// AddressToDecryptedAccount mocks base method
func (m *MockAccountManager) AddressToDecryptedAccount(arg0, arg1 string) (types.Account, *types.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressToDecryptedAccount", arg0, arg1)
	ret0, _ := ret[0].(types.Account)
	ret1, _ := ret[1].(*types.Key)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddressToDecryptedAccount indicates an expected call of AddressToDecryptedAccount
func (mr *MockAccountManagerMockRecorder) AddressToDecryptedAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressToDecryptedAccount", reflect.TypeOf((*MockAccountManager)(nil).AddressToDecryptedAccount), arg0, arg1)
}

// SelectAccount mocks base method
func (m *MockAccountManager) SelectAccount(arg0 account.LoginParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectAccount indicates an expected call of SelectAccount
func (mr *MockAccountManagerMockRecorder) SelectAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAccount", reflect.TypeOf((*MockAccountManager)(nil).SelectAccount), arg0)
}

// CreateAccount mocks base method
func (m *MockAccountManager) CreateAccount(password string) (generator.GeneratedAccountInfo, account.Info, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", password)
	ret0, _ := ret[0].(generator.GeneratedAccountInfo)
	ret1, _ := ret[1].(account.Info)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockAccountManagerMockRecorder) CreateAccount(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountManager)(nil).CreateAccount), password)
}
