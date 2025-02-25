// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/marmotedu/iam/internal/authzserver/store (interfaces: StoreClient)

// Package store is a generated GoMock package.
package store

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/marmotedu/api/proto/apiserver/v1"
	ladon "github.com/ory/ladon"
)

// MockStoreClient is a mock of StoreClient interface.
type MockStoreClient struct {
	ctrl     *gomock.Controller
	recorder *MockStoreClientMockRecorder
}

// MockStoreClientMockRecorder is the mock recorder for MockStoreClient.
type MockStoreClientMockRecorder struct {
	mock *MockStoreClient
}

// NewMockStoreClient creates a new mock instance.
func NewMockStoreClient(ctrl *gomock.Controller) *MockStoreClient {
	mock := &MockStoreClient{ctrl: ctrl}
	mock.recorder = &MockStoreClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreClient) EXPECT() *MockStoreClientMockRecorder {
	return m.recorder
}

// GetPolicies mocks base method.
func (m *MockStoreClient) GetPolicies() (map[string][]*ladon.DefaultPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicies")
	ret0, _ := ret[0].(map[string][]*ladon.DefaultPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicies indicates an expected call of GetPolicies.
func (mr *MockStoreClientMockRecorder) GetPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicies", reflect.TypeOf((*MockStoreClient)(nil).GetPolicies))
}

// GetSecrets mocks base method.
func (m *MockStoreClient) GetSecrets() (map[string]*v1.SecretInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecrets")
	ret0, _ := ret[0].(map[string]*v1.SecretInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecrets indicates an expected call of GetSecrets.
func (mr *MockStoreClientMockRecorder) GetSecrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecrets", reflect.TypeOf((*MockStoreClient)(nil).GetSecrets))
}
