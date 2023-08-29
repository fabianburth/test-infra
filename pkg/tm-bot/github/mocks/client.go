// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/test-infra/pkg/tm-bot/github (interfaces: Client)

// Package mock_github is a generated GoMock package.
package mock_github

import (
	context "context"
	json "encoding/json"
	reflect "reflect"

	semver "github.com/Masterminds/semver/v3"
	github "github.com/gardener/test-infra/pkg/tm-bot/github"
	ghval "github.com/gardener/test-infra/pkg/tm-bot/github/ghval"
	gomock "github.com/golang/mock/gomock"
	github0 "github.com/google/go-github/v54/github"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockClient) Client() *github0.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*github0.Client)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockClient)(nil).Client))
}

// Comment mocks base method.
func (m *MockClient) Comment(arg0 context.Context, arg1 *github.GenericRequestEvent, arg2 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Comment", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Comment indicates an expected call of Comment.
func (mr *MockClientMockRecorder) Comment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Comment", reflect.TypeOf((*MockClient)(nil).Comment), arg0, arg1, arg2)
}

// GetConfig mocks base method.
func (m *MockClient) GetConfig(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockClientMockRecorder) GetConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockClient)(nil).GetConfig), arg0, arg1)
}

// GetContent mocks base method.
func (m *MockClient) GetContent(arg0 context.Context, arg1 *github.GenericRequestEvent, arg2 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContent indicates an expected call of GetContent.
func (mr *MockClientMockRecorder) GetContent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockClient)(nil).GetContent), arg0, arg1, arg2)
}

// GetHead mocks base method.
func (m *MockClient) GetHead(arg0 context.Context, arg1 *github.GenericRequestEvent) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHead", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHead indicates an expected call of GetHead.
func (mr *MockClientMockRecorder) GetHead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHead", reflect.TypeOf((*MockClient)(nil).GetHead), arg0, arg1)
}

// GetIssue mocks base method.
func (m *MockClient) GetIssue(arg0 *github.GenericRequestEvent) (*github0.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssue", arg0)
	ret0, _ := ret[0].(*github0.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssue indicates an expected call of GetIssue.
func (mr *MockClientMockRecorder) GetIssue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssue", reflect.TypeOf((*MockClient)(nil).GetIssue), arg0)
}

// GetPullRequest mocks base method.
func (m *MockClient) GetPullRequest(arg0 context.Context, arg1 *github.GenericRequestEvent) (*github0.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPullRequest", arg0, arg1)
	ret0, _ := ret[0].(*github0.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPullRequest indicates an expected call of GetPullRequest.
func (mr *MockClientMockRecorder) GetPullRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequest", reflect.TypeOf((*MockClient)(nil).GetPullRequest), arg0, arg1)
}

// GetRawConfig mocks base method.
func (m *MockClient) GetRawConfig(arg0 string) (json.RawMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawConfig", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawConfig indicates an expected call of GetRawConfig.
func (mr *MockClientMockRecorder) GetRawConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawConfig", reflect.TypeOf((*MockClient)(nil).GetRawConfig), arg0)
}

// GetVersions mocks base method.
func (m *MockClient) GetVersions(arg0, arg1 string) ([]*semver.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersions", arg0, arg1)
	ret0, _ := ret[0].([]*semver.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersions indicates an expected call of GetVersions.
func (mr *MockClientMockRecorder) GetVersions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersions", reflect.TypeOf((*MockClient)(nil).GetVersions), arg0, arg1)
}

// IsAuthorized mocks base method.
func (m *MockClient) IsAuthorized(arg0 github.AuthorizationType, arg1 *github.GenericRequestEvent) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthorized", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAuthorized indicates an expected call of IsAuthorized.
func (mr *MockClientMockRecorder) IsAuthorized(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorized", reflect.TypeOf((*MockClient)(nil).IsAuthorized), arg0, arg1)
}

// ResolveConfigValue mocks base method.
func (m *MockClient) ResolveConfigValue(arg0 context.Context, arg1 *github.GenericRequestEvent, arg2 *ghval.GitHubValue) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveConfigValue", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveConfigValue indicates an expected call of ResolveConfigValue.
func (mr *MockClientMockRecorder) ResolveConfigValue(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveConfigValue", reflect.TypeOf((*MockClient)(nil).ResolveConfigValue), arg0, arg1, arg2)
}

// UpdateComment mocks base method.
func (m *MockClient) UpdateComment(arg0 *github.GenericRequestEvent, arg1 int64, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockClientMockRecorder) UpdateComment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockClient)(nil).UpdateComment), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockClient) UpdateStatus(arg0 context.Context, arg1 *github.GenericRequestEvent, arg2 github.State, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockClientMockRecorder) UpdateStatus(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockClient)(nil).UpdateStatus), arg0, arg1, arg2, arg3, arg4)
}
