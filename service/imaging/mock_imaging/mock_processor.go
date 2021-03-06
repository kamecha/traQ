// Code generated by MockGen. DO NOT EDIT.
// Source: processor.go

// Package mock_imaging is a generated GoMock package.
package mock_imaging

import (
	bytes "bytes"
	gomock "github.com/golang/mock/gomock"
	image "image"
	io "io"
	reflect "reflect"
)

// MockProcessor is a mock of Processor interface
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// Thumbnail mocks base method
func (m *MockProcessor) Thumbnail(src io.ReadSeeker) (image.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Thumbnail", src)
	ret0, _ := ret[0].(image.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Thumbnail indicates an expected call of Thumbnail
func (mr *MockProcessorMockRecorder) Thumbnail(src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Thumbnail", reflect.TypeOf((*MockProcessor)(nil).Thumbnail), src)
}

// Fit mocks base method
func (m *MockProcessor) Fit(src io.ReadSeeker, width, height int) (image.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fit", src, width, height)
	ret0, _ := ret[0].(image.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fit indicates an expected call of Fit
func (mr *MockProcessorMockRecorder) Fit(src, width, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fit", reflect.TypeOf((*MockProcessor)(nil).Fit), src, width, height)
}

// FitAnimationGIF mocks base method
func (m *MockProcessor) FitAnimationGIF(src io.Reader, width, height int) (*bytes.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FitAnimationGIF", src, width, height)
	ret0, _ := ret[0].(*bytes.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FitAnimationGIF indicates an expected call of FitAnimationGIF
func (mr *MockProcessorMockRecorder) FitAnimationGIF(src, width, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FitAnimationGIF", reflect.TypeOf((*MockProcessor)(nil).FitAnimationGIF), src, width, height)
}
