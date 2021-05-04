// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloud/ec2_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	aws "github.com/aws/aws-sdk-go/aws"
	ec2metadata "github.com/aws/aws-sdk-go/aws/ec2metadata"
	request "github.com/aws/aws-sdk-go/aws/request"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
)

// MockEC2 is a mock of EC2 interface.
type MockEC2 struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MockRecorder
}

// MockEC2MockRecorder is the mock recorder for MockEC2.
type MockEC2MockRecorder struct {
	mock *MockEC2
}

// NewMockEC2 creates a new mock instance.
func NewMockEC2(ctrl *gomock.Controller) *MockEC2 {
	mock := &MockEC2{ctrl: ctrl}
	mock.recorder = &MockEC2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2) EXPECT() *MockEC2MockRecorder {
	return m.recorder
}

// AttachVolumeWithContext mocks base method.
func (m *MockEC2) AttachVolumeWithContext(ctx aws.Context, input *ec2.AttachVolumeInput, opts ...request.Option) (*ec2.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AttachVolumeWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachVolumeWithContext indicates an expected call of AttachVolumeWithContext.
func (mr *MockEC2MockRecorder) AttachVolumeWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolumeWithContext", reflect.TypeOf((*MockEC2)(nil).AttachVolumeWithContext), varargs...)
}

// CreateSnapshotWithContext mocks base method.
func (m *MockEC2) CreateSnapshotWithContext(ctx aws.Context, input *ec2.CreateSnapshotInput, opts ...request.Option) (*ec2.Snapshot, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSnapshotWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshotWithContext indicates an expected call of CreateSnapshotWithContext.
func (mr *MockEC2MockRecorder) CreateSnapshotWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshotWithContext", reflect.TypeOf((*MockEC2)(nil).CreateSnapshotWithContext), varargs...)
}

// CreateVolumeWithContext mocks base method.
func (m *MockEC2) CreateVolumeWithContext(ctx aws.Context, input *ec2.CreateVolumeInput, opts ...request.Option) (*ec2.Volume, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateVolumeWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolumeWithContext indicates an expected call of CreateVolumeWithContext.
func (mr *MockEC2MockRecorder) CreateVolumeWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolumeWithContext", reflect.TypeOf((*MockEC2)(nil).CreateVolumeWithContext), varargs...)
}

// DeleteSnapshotWithContext mocks base method.
func (m *MockEC2) DeleteSnapshotWithContext(ctx aws.Context, input *ec2.DeleteSnapshotInput, opts ...request.Option) (*ec2.DeleteSnapshotOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSnapshotWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DeleteSnapshotOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnapshotWithContext indicates an expected call of DeleteSnapshotWithContext.
func (mr *MockEC2MockRecorder) DeleteSnapshotWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshotWithContext", reflect.TypeOf((*MockEC2)(nil).DeleteSnapshotWithContext), varargs...)
}

// DeleteVolumeWithContext mocks base method.
func (m *MockEC2) DeleteVolumeWithContext(ctx aws.Context, input *ec2.DeleteVolumeInput, opts ...request.Option) (*ec2.DeleteVolumeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteVolumeWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DeleteVolumeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVolumeWithContext indicates an expected call of DeleteVolumeWithContext.
func (mr *MockEC2MockRecorder) DeleteVolumeWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolumeWithContext", reflect.TypeOf((*MockEC2)(nil).DeleteVolumeWithContext), varargs...)
}

// DescribeAvailabilityZonesWithContext mocks base method.
func (m *MockEC2) DescribeAvailabilityZonesWithContext(ctx aws.Context, input *ec2.DescribeAvailabilityZonesInput, opts ...request.Option) (*ec2.DescribeAvailabilityZonesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAvailabilityZonesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeAvailabilityZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAvailabilityZonesWithContext indicates an expected call of DescribeAvailabilityZonesWithContext.
func (mr *MockEC2MockRecorder) DescribeAvailabilityZonesWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAvailabilityZonesWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeAvailabilityZonesWithContext), varargs...)
}

// DescribeInstancesWithContext mocks base method.
func (m *MockEC2) DescribeInstancesWithContext(ctx aws.Context, input *ec2.DescribeInstancesInput, opts ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstancesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstancesWithContext indicates an expected call of DescribeInstancesWithContext.
func (mr *MockEC2MockRecorder) DescribeInstancesWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstancesWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeInstancesWithContext), varargs...)
}

// DescribeSnapshotsWithContext mocks base method.
func (m *MockEC2) DescribeSnapshotsWithContext(ctx aws.Context, input *ec2.DescribeSnapshotsInput, opts ...request.Option) (*ec2.DescribeSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSnapshotsWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSnapshotsWithContext indicates an expected call of DescribeSnapshotsWithContext.
func (mr *MockEC2MockRecorder) DescribeSnapshotsWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshotsWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeSnapshotsWithContext), varargs...)
}

// DescribeVolumesModificationsWithContext mocks base method.
func (m *MockEC2) DescribeVolumesModificationsWithContext(ctx aws.Context, input *ec2.DescribeVolumesModificationsInput, opts ...request.Option) (*ec2.DescribeVolumesModificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVolumesModificationsWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVolumesModificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVolumesModificationsWithContext indicates an expected call of DescribeVolumesModificationsWithContext.
func (mr *MockEC2MockRecorder) DescribeVolumesModificationsWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVolumesModificationsWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeVolumesModificationsWithContext), varargs...)
}

// DescribeVolumesWithContext mocks base method.
func (m *MockEC2) DescribeVolumesWithContext(ctx aws.Context, input *ec2.DescribeVolumesInput, opts ...request.Option) (*ec2.DescribeVolumesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVolumesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVolumesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVolumesWithContext indicates an expected call of DescribeVolumesWithContext.
func (mr *MockEC2MockRecorder) DescribeVolumesWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVolumesWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeVolumesWithContext), varargs...)
}

// DetachVolumeWithContext mocks base method.
func (m *MockEC2) DetachVolumeWithContext(ctx aws.Context, input *ec2.DetachVolumeInput, opts ...request.Option) (*ec2.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DetachVolumeWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachVolumeWithContext indicates an expected call of DetachVolumeWithContext.
func (mr *MockEC2MockRecorder) DetachVolumeWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachVolumeWithContext", reflect.TypeOf((*MockEC2)(nil).DetachVolumeWithContext), varargs...)
}

// ModifyVolumeWithContext mocks base method.
func (m *MockEC2) ModifyVolumeWithContext(ctx aws.Context, input *ec2.ModifyVolumeInput, opts ...request.Option) (*ec2.ModifyVolumeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModifyVolumeWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.ModifyVolumeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyVolumeWithContext indicates an expected call of ModifyVolumeWithContext.
func (mr *MockEC2MockRecorder) ModifyVolumeWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyVolumeWithContext", reflect.TypeOf((*MockEC2)(nil).ModifyVolumeWithContext), varargs...)
}

// MockEC2Metadata is a mock of EC2Metadata interface.
type MockEC2Metadata struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MetadataMockRecorder
}

// MockEC2MetadataMockRecorder is the mock recorder for MockEC2Metadata.
type MockEC2MetadataMockRecorder struct {
	mock *MockEC2Metadata
}

// NewMockEC2Metadata creates a new mock instance.
func NewMockEC2Metadata(ctrl *gomock.Controller) *MockEC2Metadata {
	mock := &MockEC2Metadata{ctrl: ctrl}
	mock.recorder = &MockEC2MetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2Metadata) EXPECT() *MockEC2MetadataMockRecorder {
	return m.recorder
}

// Available mocks base method.
func (m *MockEC2Metadata) Available() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Available")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Available indicates an expected call of Available.
func (mr *MockEC2MetadataMockRecorder) Available() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Available", reflect.TypeOf((*MockEC2Metadata)(nil).Available))
}

// GetInstanceIdentityDocument mocks base method.
func (m *MockEC2Metadata) GetInstanceIdentityDocument() (ec2metadata.EC2InstanceIdentityDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceIdentityDocument")
	ret0, _ := ret[0].(ec2metadata.EC2InstanceIdentityDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceIdentityDocument indicates an expected call of GetInstanceIdentityDocument.
func (mr *MockEC2MetadataMockRecorder) GetInstanceIdentityDocument() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceIdentityDocument", reflect.TypeOf((*MockEC2Metadata)(nil).GetInstanceIdentityDocument))
}

// GetMetadata mocks base method.
func (m *MockEC2Metadata) GetMetadata(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockEC2MetadataMockRecorder) GetMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockEC2Metadata)(nil).GetMetadata), arg0)
}
