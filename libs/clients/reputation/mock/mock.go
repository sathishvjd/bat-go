// Code generated by MockGen. DO NOT EDIT.
// Source: ./libs/clients/reputation/client.go

// Package mock_reputation is a generated GoMock package.
package mock_reputation

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	decimal "github.com/shopspring/decimal"
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

// CreateReputationSummary mocks base method.
func (m *MockClient) CreateReputationSummary(ctx context.Context, paymentID, geoCountry string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReputationSummary", ctx, paymentID, geoCountry)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReputationSummary indicates an expected call of CreateReputationSummary.
func (mr *MockClientMockRecorder) CreateReputationSummary(ctx, paymentID, geoCountry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReputationSummary", reflect.TypeOf((*MockClient)(nil).CreateReputationSummary), ctx, paymentID, geoCountry)
}

// IsDrainReputable mocks base method.
func (m *MockClient) IsDrainReputable(ctx context.Context, id, promotionID uuid.UUID, withdrawAmount decimal.Decimal) (bool, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDrainReputable", ctx, id, promotionID, withdrawAmount)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsDrainReputable indicates an expected call of IsDrainReputable.
func (mr *MockClientMockRecorder) IsDrainReputable(ctx, id, promotionID, withdrawAmount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDrainReputable", reflect.TypeOf((*MockClient)(nil).IsDrainReputable), ctx, id, promotionID, withdrawAmount)
}

// IsLinkingReputable mocks base method.
func (m *MockClient) IsLinkingReputable(ctx context.Context, id uuid.UUID, country string) (bool, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLinkingReputable", ctx, id, country)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsLinkingReputable indicates an expected call of IsLinkingReputable.
func (mr *MockClientMockRecorder) IsLinkingReputable(ctx, id, country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLinkingReputable", reflect.TypeOf((*MockClient)(nil).IsLinkingReputable), ctx, id, country)
}

// IsWalletAdsReputable mocks base method.
func (m *MockClient) IsWalletAdsReputable(ctx context.Context, id uuid.UUID, platform string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWalletAdsReputable", ctx, id, platform)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsWalletAdsReputable indicates an expected call of IsWalletAdsReputable.
func (mr *MockClientMockRecorder) IsWalletAdsReputable(ctx, id, platform interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWalletAdsReputable", reflect.TypeOf((*MockClient)(nil).IsWalletAdsReputable), ctx, id, platform)
}

// IsWalletOnPlatform mocks base method.
func (m *MockClient) IsWalletOnPlatform(ctx context.Context, id uuid.UUID, platform string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWalletOnPlatform", ctx, id, platform)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsWalletOnPlatform indicates an expected call of IsWalletOnPlatform.
func (mr *MockClientMockRecorder) IsWalletOnPlatform(ctx, id, platform interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWalletOnPlatform", reflect.TypeOf((*MockClient)(nil).IsWalletOnPlatform), ctx, id, platform)
}

// IsWalletReputable mocks base method.
func (m *MockClient) IsWalletReputable(ctx context.Context, id uuid.UUID, platform string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWalletReputable", ctx, id, platform)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsWalletReputable indicates an expected call of IsWalletReputable.
func (mr *MockClientMockRecorder) IsWalletReputable(ctx, id, platform interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWalletReputable", reflect.TypeOf((*MockClient)(nil).IsWalletReputable), ctx, id, platform)
}
