// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw-plugin-yarpc
// @generated

package matchingservicetest

import (
	"context"
	"go.uber.org/yarpc"
	"github.com/golang/mock/gomock"
	"github.com/uber/cadence/.gen/go/matching"
	"github.com/uber/cadence/.gen/go/shared"
	"github.com/uber/cadence/.gen/go/matching/matchingserviceclient"
)

// MockClient implements a gomock-compatible mock client for service
// MatchingService.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

var _ matchingserviceclient.Interface = (*MockClient)(nil)

type _MockClientRecorder struct {
	mock *MockClient
}

// Build a new mock client for service MatchingService.
//
// 	mockCtrl := gomock.NewController(t)
// 	client := matchingservicetest.NewMockClient(mockCtrl)
//
// Use EXPECT() to set expectations on the mock.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

// EXPECT returns an object that allows you to define an expectation on the
// MatchingService mock client.
func (m *MockClient) EXPECT() *_MockClientRecorder {
	return m.recorder
}

// AddActivityTask responds to a AddActivityTask call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().AddActivityTask(gomock.Any(), ...).Return(...)
// 	... := client.AddActivityTask(...)
func (m *MockClient) AddActivityTask(
	ctx context.Context,
	_AddRequest *matching.AddActivityTaskRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _AddRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "AddActivityTask", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) AddActivityTask(
	ctx interface{},
	_AddRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _AddRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "AddActivityTask", args...)
}

// AddDecisionTask responds to a AddDecisionTask call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().AddDecisionTask(gomock.Any(), ...).Return(...)
// 	... := client.AddDecisionTask(...)
func (m *MockClient) AddDecisionTask(
	ctx context.Context,
	_AddRequest *matching.AddDecisionTaskRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _AddRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "AddDecisionTask", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) AddDecisionTask(
	ctx interface{},
	_AddRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _AddRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "AddDecisionTask", args...)
}

// PollForActivityTask responds to a PollForActivityTask call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().PollForActivityTask(gomock.Any(), ...).Return(...)
// 	... := client.PollForActivityTask(...)
func (m *MockClient) PollForActivityTask(
	ctx context.Context,
	_PollRequest *matching.PollForActivityTaskRequest,
	opts ...yarpc.CallOption,
) (success *shared.PollForActivityTaskResponse, err error) {

	args := []interface{}{ctx, _PollRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "PollForActivityTask", args...)
	success, _ = ret[i].(*shared.PollForActivityTaskResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) PollForActivityTask(
	ctx interface{},
	_PollRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _PollRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "PollForActivityTask", args...)
}

// PollForDecisionTask responds to a PollForDecisionTask call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().PollForDecisionTask(gomock.Any(), ...).Return(...)
// 	... := client.PollForDecisionTask(...)
func (m *MockClient) PollForDecisionTask(
	ctx context.Context,
	_PollRequest *matching.PollForDecisionTaskRequest,
	opts ...yarpc.CallOption,
) (success *matching.PollForDecisionTaskResponse, err error) {

	args := []interface{}{ctx, _PollRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "PollForDecisionTask", args...)
	success, _ = ret[i].(*matching.PollForDecisionTaskResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) PollForDecisionTask(
	ctx interface{},
	_PollRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _PollRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "PollForDecisionTask", args...)
}
