// Code generated by protoc-gen-go.
// source: pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	pps/pps.proto

It has these top-level messages:
	Transform
	Job
	JobInfo
	JobInfos
	Pipeline
	PipelineInfo
	PipelineInfos
	CreateJobRequest
	InspectJobRequest
	ListJobRequest
	CreatePipelineRequest
	InspectPipelineRequest
	ListPipelineRequest
	DeletePipelineRequest
	StartJobRequest
	StartJobResponse
	FinishJobRequest
*/
package pps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/google-protobuf"
import google_protobuf1 "go.pedge.io/google-protobuf"
import pfs "github.com/pachyderm/pachyderm/src/pfs"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JobState int32

const (
	JobState_JOB_STATE_RUNNING JobState = 0
	JobState_JOB_STATE_FAILURE JobState = 1
	JobState_JOB_STATE_SUCCESS JobState = 2
)

var JobState_name = map[int32]string{
	0: "JOB_STATE_RUNNING",
	1: "JOB_STATE_FAILURE",
	2: "JOB_STATE_SUCCESS",
}
var JobState_value = map[string]int32{
	"JOB_STATE_RUNNING": 0,
	"JOB_STATE_FAILURE": 1,
	"JOB_STATE_SUCCESS": 2,
}

func (x JobState) String() string {
	return proto.EnumName(JobState_name, int32(x))
}

type Transform struct {
	Image string   `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Cmd   []string `protobuf:"bytes,2,rep,name=cmd" json:"cmd,omitempty"`
	Stdin string   `protobuf:"bytes,3,opt,name=stdin" json:"stdin,omitempty"`
}

func (m *Transform) Reset()         { *m = Transform{} }
func (m *Transform) String() string { return proto.CompactTextString(m) }
func (*Transform) ProtoMessage()    {}

type Job struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}

// TODO: add created at?
type JobInfo struct {
	Job          *Job                        `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	Transform    *Transform                  `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Pipeline     *Pipeline                   `protobuf:"bytes,3,opt,name=pipeline" json:"pipeline,omitempty"`
	Shards       uint64                      `protobuf:"varint,4,opt,name=shards" json:"shards,omitempty"`
	InputCommit  []*pfs.Commit               `protobuf:"bytes,5,rep,name=input_commit" json:"input_commit,omitempty"`
	ParentJob    *Job                        `protobuf:"bytes,6,opt,name=parent_job" json:"parent_job,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp `protobuf:"bytes,7,opt,name=created_at" json:"created_at,omitempty"`
	OutputCommit *pfs.Commit                 `protobuf:"bytes,8,opt,name=output_commit" json:"output_commit,omitempty"`
	State        JobState                    `protobuf:"varint,9,opt,name=state,enum=pachyderm.pps.JobState" json:"state,omitempty"`
}

func (m *JobInfo) Reset()         { *m = JobInfo{} }
func (m *JobInfo) String() string { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()    {}

func (m *JobInfo) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *JobInfo) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *JobInfo) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *JobInfo) GetInputCommit() []*pfs.Commit {
	if m != nil {
		return m.InputCommit
	}
	return nil
}

func (m *JobInfo) GetParentJob() *Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

func (m *JobInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *JobInfo) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

type JobInfos struct {
	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info" json:"job_info,omitempty"`
}

func (m *JobInfos) Reset()         { *m = JobInfos{} }
func (m *JobInfos) String() string { return proto.CompactTextString(m) }
func (*JobInfos) ProtoMessage()    {}

func (m *JobInfos) GetJobInfo() []*JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

type Pipeline struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}

// TODO: add created at?
type PipelineInfo struct {
	Pipeline   *Pipeline   `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	Transform  *Transform  `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Shards     uint64      `protobuf:"varint,3,opt,name=shards" json:"shards,omitempty"`
	InputRepo  []*pfs.Repo `protobuf:"bytes,4,rep,name=input_repo" json:"input_repo,omitempty"`
	OutputRepo *pfs.Repo   `protobuf:"bytes,5,opt,name=output_repo" json:"output_repo,omitempty"`
}

func (m *PipelineInfo) Reset()         { *m = PipelineInfo{} }
func (m *PipelineInfo) String() string { return proto.CompactTextString(m) }
func (*PipelineInfo) ProtoMessage()    {}

func (m *PipelineInfo) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *PipelineInfo) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *PipelineInfo) GetInputRepo() []*pfs.Repo {
	if m != nil {
		return m.InputRepo
	}
	return nil
}

func (m *PipelineInfo) GetOutputRepo() *pfs.Repo {
	if m != nil {
		return m.OutputRepo
	}
	return nil
}

type PipelineInfos struct {
	PipelineInfo []*PipelineInfo `protobuf:"bytes,1,rep,name=pipeline_info" json:"pipeline_info,omitempty"`
}

func (m *PipelineInfos) Reset()         { *m = PipelineInfos{} }
func (m *PipelineInfos) String() string { return proto.CompactTextString(m) }
func (*PipelineInfos) ProtoMessage()    {}

func (m *PipelineInfos) GetPipelineInfo() []*PipelineInfo {
	if m != nil {
		return m.PipelineInfo
	}
	return nil
}

type CreateJobRequest struct {
	Transform   *Transform    `protobuf:"bytes,1,opt,name=transform" json:"transform,omitempty"`
	Pipeline    *Pipeline     `protobuf:"bytes,2,opt,name=pipeline" json:"pipeline,omitempty"`
	Shards      uint64        `protobuf:"varint,3,opt,name=shards" json:"shards,omitempty"`
	InputCommit []*pfs.Commit `protobuf:"bytes,4,rep,name=input_commit" json:"input_commit,omitempty"`
	ParentJob   *Job          `protobuf:"bytes,5,opt,name=parent_job" json:"parent_job,omitempty"`
}

func (m *CreateJobRequest) Reset()         { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()    {}

func (m *CreateJobRequest) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *CreateJobRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *CreateJobRequest) GetInputCommit() []*pfs.Commit {
	if m != nil {
		return m.InputCommit
	}
	return nil
}

func (m *CreateJobRequest) GetParentJob() *Job {
	if m != nil {
		return m.ParentJob
	}
	return nil
}

type InspectJobRequest struct {
	Job         *Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	BlockOutput bool `protobuf:"varint,2,opt,name=block_output" json:"block_output,omitempty"`
	BlockState  bool `protobuf:"varint,3,opt,name=block_state" json:"block_state,omitempty"`
}

func (m *InspectJobRequest) Reset()         { *m = InspectJobRequest{} }
func (m *InspectJobRequest) String() string { return proto.CompactTextString(m) }
func (*InspectJobRequest) ProtoMessage()    {}

func (m *InspectJobRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type ListJobRequest struct {
	Pipeline    *Pipeline     `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	InputCommit []*pfs.Commit `protobuf:"bytes,2,rep,name=input_commit" json:"input_commit,omitempty"`
}

func (m *ListJobRequest) Reset()         { *m = ListJobRequest{} }
func (m *ListJobRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobRequest) ProtoMessage()    {}

func (m *ListJobRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *ListJobRequest) GetInputCommit() []*pfs.Commit {
	if m != nil {
		return m.InputCommit
	}
	return nil
}

type CreatePipelineRequest struct {
	Pipeline  *Pipeline   `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	Transform *Transform  `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Shards    uint64      `protobuf:"varint,3,opt,name=shards" json:"shards,omitempty"`
	InputRepo []*pfs.Repo `protobuf:"bytes,4,rep,name=input_repo" json:"input_repo,omitempty"`
}

func (m *CreatePipelineRequest) Reset()         { *m = CreatePipelineRequest{} }
func (m *CreatePipelineRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePipelineRequest) ProtoMessage()    {}

func (m *CreatePipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *CreatePipelineRequest) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *CreatePipelineRequest) GetInputRepo() []*pfs.Repo {
	if m != nil {
		return m.InputRepo
	}
	return nil
}

type InspectPipelineRequest struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *InspectPipelineRequest) Reset()         { *m = InspectPipelineRequest{} }
func (m *InspectPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*InspectPipelineRequest) ProtoMessage()    {}

func (m *InspectPipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type ListPipelineRequest struct {
}

func (m *ListPipelineRequest) Reset()         { *m = ListPipelineRequest{} }
func (m *ListPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*ListPipelineRequest) ProtoMessage()    {}

type DeletePipelineRequest struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *DeletePipelineRequest) Reset()         { *m = DeletePipelineRequest{} }
func (m *DeletePipelineRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePipelineRequest) ProtoMessage()    {}

func (m *DeletePipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type StartJobRequest struct {
	Job *Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
}

func (m *StartJobRequest) Reset()         { *m = StartJobRequest{} }
func (m *StartJobRequest) String() string { return proto.CompactTextString(m) }
func (*StartJobRequest) ProtoMessage()    {}

func (m *StartJobRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type StartJobResponse struct {
	Transform    *Transform    `protobuf:"bytes,1,opt,name=transform" json:"transform,omitempty"`
	InputCommit  []*pfs.Commit `protobuf:"bytes,2,rep,name=input_commit" json:"input_commit,omitempty"`
	OutputCommit *pfs.Commit   `protobuf:"bytes,3,opt,name=output_commit" json:"output_commit,omitempty"`
	Shard        *pfs.Shard    `protobuf:"bytes,4,opt,name=shard" json:"shard,omitempty"`
}

func (m *StartJobResponse) Reset()         { *m = StartJobResponse{} }
func (m *StartJobResponse) String() string { return proto.CompactTextString(m) }
func (*StartJobResponse) ProtoMessage()    {}

func (m *StartJobResponse) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *StartJobResponse) GetInputCommit() []*pfs.Commit {
	if m != nil {
		return m.InputCommit
	}
	return nil
}

func (m *StartJobResponse) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

func (m *StartJobResponse) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

type FinishJobRequest struct {
	Job     *Job       `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	Shard   *pfs.Shard `protobuf:"bytes,2,opt,name=shard" json:"shard,omitempty"`
	Success bool       `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
}

func (m *FinishJobRequest) Reset()         { *m = FinishJobRequest{} }
func (m *FinishJobRequest) String() string { return proto.CompactTextString(m) }
func (*FinishJobRequest) ProtoMessage()    {}

func (m *FinishJobRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *FinishJobRequest) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

func init() {
	proto.RegisterType((*Transform)(nil), "pachyderm.pps.Transform")
	proto.RegisterType((*Job)(nil), "pachyderm.pps.Job")
	proto.RegisterType((*JobInfo)(nil), "pachyderm.pps.JobInfo")
	proto.RegisterType((*JobInfos)(nil), "pachyderm.pps.JobInfos")
	proto.RegisterType((*Pipeline)(nil), "pachyderm.pps.Pipeline")
	proto.RegisterType((*PipelineInfo)(nil), "pachyderm.pps.PipelineInfo")
	proto.RegisterType((*PipelineInfos)(nil), "pachyderm.pps.PipelineInfos")
	proto.RegisterType((*CreateJobRequest)(nil), "pachyderm.pps.CreateJobRequest")
	proto.RegisterType((*InspectJobRequest)(nil), "pachyderm.pps.InspectJobRequest")
	proto.RegisterType((*ListJobRequest)(nil), "pachyderm.pps.ListJobRequest")
	proto.RegisterType((*CreatePipelineRequest)(nil), "pachyderm.pps.CreatePipelineRequest")
	proto.RegisterType((*InspectPipelineRequest)(nil), "pachyderm.pps.InspectPipelineRequest")
	proto.RegisterType((*ListPipelineRequest)(nil), "pachyderm.pps.ListPipelineRequest")
	proto.RegisterType((*DeletePipelineRequest)(nil), "pachyderm.pps.DeletePipelineRequest")
	proto.RegisterType((*StartJobRequest)(nil), "pachyderm.pps.StartJobRequest")
	proto.RegisterType((*StartJobResponse)(nil), "pachyderm.pps.StartJobResponse")
	proto.RegisterType((*FinishJobRequest)(nil), "pachyderm.pps.FinishJobRequest")
	proto.RegisterEnum("pachyderm.pps.JobState", JobState_name, JobState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for JobAPI service

type JobAPIClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error)
	InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
}

type jobAPIClient struct {
	cc *grpc.ClientConn
}

func NewJobAPIClient(cc *grpc.ClientConn) JobAPIClient {
	return &jobAPIClient{cc}
}

func (c *jobAPIClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/pachyderm.pps.JobAPI/CreateJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobAPIClient) InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.JobAPI/InspectJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobAPIClient) ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.JobAPI/ListJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JobAPI service

type JobAPIServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*Job, error)
	InspectJob(context.Context, *InspectJobRequest) (*JobInfo, error)
	ListJob(context.Context, *ListJobRequest) (*JobInfos, error)
}

func RegisterJobAPIServer(s *grpc.Server, srv JobAPIServer) {
	s.RegisterService(&_JobAPI_serviceDesc, srv)
}

func _JobAPI_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobAPIServer).CreateJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _JobAPI_InspectJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InspectJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobAPIServer).InspectJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _JobAPI_ListJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobAPIServer).ListJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _JobAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.JobAPI",
	HandlerType: (*JobAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _JobAPI_CreateJob_Handler,
		},
		{
			MethodName: "InspectJob",
			Handler:    _JobAPI_InspectJob_Handler,
		},
		{
			MethodName: "ListJob",
			Handler:    _JobAPI_ListJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for PipelineAPI service

type PipelineAPIClient interface {
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error)
	ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type pipelineAPIClient struct {
	cc *grpc.ClientConn
}

func NewPipelineAPIClient(cc *grpc.ClientConn) PipelineAPIClient {
	return &pipelineAPIClient{cc}
}

func (c *pipelineAPIClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/CreatePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineAPIClient) InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/InspectPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineAPIClient) ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/ListPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineAPIClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/DeletePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PipelineAPI service

type PipelineAPIServer interface {
	CreatePipeline(context.Context, *CreatePipelineRequest) (*google_protobuf.Empty, error)
	InspectPipeline(context.Context, *InspectPipelineRequest) (*PipelineInfo, error)
	ListPipeline(context.Context, *ListPipelineRequest) (*PipelineInfos, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*google_protobuf.Empty, error)
}

func RegisterPipelineAPIServer(s *grpc.Server, srv PipelineAPIServer) {
	s.RegisterService(&_PipelineAPI_serviceDesc, srv)
}

func _PipelineAPI_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).CreatePipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PipelineAPI_InspectPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InspectPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).InspectPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PipelineAPI_ListPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).ListPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PipelineAPI_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).DeletePipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PipelineAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.PipelineAPI",
	HandlerType: (*PipelineAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _PipelineAPI_CreatePipeline_Handler,
		},
		{
			MethodName: "InspectPipeline",
			Handler:    _PipelineAPI_InspectPipeline_Handler,
		},
		{
			MethodName: "ListPipeline",
			Handler:    _PipelineAPI_ListPipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _PipelineAPI_DeletePipeline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for InternalJobAPI service

type InternalJobAPIClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error)
	FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type internalJobAPIClient struct {
	cc *grpc.ClientConn
}

func NewInternalJobAPIClient(cc *grpc.ClientConn) InternalJobAPIClient {
	return &internalJobAPIClient{cc}
}

func (c *internalJobAPIClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error) {
	out := new(StartJobResponse)
	err := grpc.Invoke(ctx, "/pachyderm.pps.InternalJobAPI/StartJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalJobAPIClient) FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.InternalJobAPI/FinishJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InternalJobAPI service

type InternalJobAPIServer interface {
	StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error)
	FinishJob(context.Context, *FinishJobRequest) (*google_protobuf.Empty, error)
}

func RegisterInternalJobAPIServer(s *grpc.Server, srv InternalJobAPIServer) {
	s.RegisterService(&_InternalJobAPI_serviceDesc, srv)
}

func _InternalJobAPI_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(InternalJobAPIServer).StartJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _InternalJobAPI_FinishJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(FinishJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(InternalJobAPIServer).FinishJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _InternalJobAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.InternalJobAPI",
	HandlerType: (*InternalJobAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _InternalJobAPI_StartJob_Handler,
		},
		{
			MethodName: "FinishJob",
			Handler:    _InternalJobAPI_FinishJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
