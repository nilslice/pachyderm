package server

import (
	"github.com/pachyderm/pachyderm/src/client/admin"
	auth1_11 "github.com/pachyderm/pachyderm/src/client/admin/v1_11/auth"
	pfs1_11 "github.com/pachyderm/pachyderm/src/client/admin/v1_11/pfs"
	pps1_11 "github.com/pachyderm/pachyderm/src/client/admin/v1_11/pps"
	auth1_12 "github.com/pachyderm/pachyderm/src/client/admin/v1_12/auth"
	enterprise1_12 "github.com/pachyderm/pachyderm/src/client/admin/v1_12/enterprise"
	pfs1_12 "github.com/pachyderm/pachyderm/src/client/admin/v1_12/pfs"
	pps1_12 "github.com/pachyderm/pachyderm/src/client/admin/v1_12/pps"
	"github.com/pachyderm/pachyderm/src/client/pkg/errors"
)

func convert1_11Repo(r *pfs1_11.Repo) *pfs1_12.Repo {
	if r == nil {
		return nil
	}
	return &pfs1_12.Repo{
		Name: r.Name,
	}
}

func convert1_11Commit(c *pfs1_11.Commit) *pfs1_12.Commit {
	if c == nil {
		return nil
	}
	return &pfs1_12.Commit{
		Repo: convert1_11Repo(c.Repo),
		ID:   c.ID,
	}
}

func convert1_11Provenance(provenance *pfs1_11.CommitProvenance) *pfs1_12.CommitProvenance {
	if provenance == nil {
		return nil
	}
	return &pfs1_12.CommitProvenance{
		Commit: convert1_11Commit(provenance.Commit),
		Branch: convert1_11Branch(provenance.Branch),
	}
}

func convert1_11Provenances(provenances []*pfs1_11.CommitProvenance) []*pfs1_12.CommitProvenance {
	if provenances == nil {
		return nil
	}
	result := make([]*pfs1_12.CommitProvenance, 0, len(provenances))
	for _, p := range provenances {
		result = append(result, convert1_11Provenance(p))
	}
	return result
}

func convert1_11Job(j *pps1_11.CreateJobRequest) *pps1_12.CreateJobRequest {
	if j == nil {
		return nil
	}
	return &pps1_12.CreateJobRequest{
		Pipeline:      convert1_11Pipeline(j.Pipeline),
		OutputCommit:  convert1_11Commit(j.OutputCommit),
		Restart:       j.Restart,
		DataProcessed: j.DataProcessed,
		DataSkipped:   j.DataSkipped,
		DataTotal:     j.DataTotal,
		DataFailed:    j.DataFailed,
		DataRecovered: j.DataRecovered,
		Stats:         convert1_11Stats(j.Stats),
		StatsCommit:   convert1_11Commit(j.StatsCommit),
		State:         pps1_12.JobState(j.State),
		Reason:        j.Reason,
		Started:       j.Started,
		Finished:      j.Finished,
	}
}

func convert1_11Stats(s *pps1_11.ProcessStats) *pps1_12.ProcessStats {
	if s == nil {
		return nil
	}
	return &pps1_12.ProcessStats{
		DownloadTime:  s.DownloadTime,
		ProcessTime:   s.ProcessTime,
		UploadTime:    s.UploadTime,
		DownloadBytes: s.DownloadBytes,
		UploadBytes:   s.UploadBytes,
	}
}

func convert1_11CreateObject(o *pfs1_11.CreateObjectRequest) *pfs1_12.CreateObjectRequest {
	if o == nil {
		return nil
	}
	return &pfs1_12.CreateObjectRequest{
		Object:   convert1_11Object(o.Object),
		BlockRef: convert1_11BlockRef(o.BlockRef),
	}
}

func convert1_11Object(o *pfs1_11.Object) *pfs1_12.Object {
	if o == nil {
		return nil
	}
	return &pfs1_12.Object{
		Hash: o.Hash,
	}
}

func convert1_11BlockRef(b *pfs1_11.BlockRef) *pfs1_12.BlockRef {
	if b == nil {
		return nil
	}
	return &pfs1_12.BlockRef{
		Block: &pfs1_12.Block{
			Hash: b.Block.Hash,
		},
		Range: &pfs1_12.ByteRange{
			Lower: b.Range.Lower,
			Upper: b.Range.Upper,
		},
	}
}

func convert1_11Objects(objects []*pfs1_11.Object) []*pfs1_12.Object {
	if objects == nil {
		return nil
	}
	result := make([]*pfs1_12.Object, 0, len(objects))
	for _, o := range objects {
		result = append(result, convert1_11Object(o))
	}
	return result
}

func convert1_11Tag(tag *pfs1_11.Tag) *pfs1_12.Tag {
	if tag == nil {
		return nil
	}
	return &pfs1_12.Tag{
		Name: tag.Name,
	}
}

func convert1_11Tags(tags []*pfs1_11.Tag) []*pfs1_12.Tag {
	if tags == nil {
		return nil
	}
	result := make([]*pfs1_12.Tag, 0, len(tags))
	for _, t := range tags {
		result = append(result, convert1_11Tag(t))
	}
	return result
}

func convert1_11Branch(b *pfs1_11.Branch) *pfs1_12.Branch {
	if b == nil {
		return nil
	}
	return &pfs1_12.Branch{
		Repo: convert1_11Repo(b.Repo),
		Name: b.Name,
	}
}

func convert1_11Branches(branches []*pfs1_11.Branch) []*pfs1_12.Branch {
	if branches == nil {
		return nil
	}
	result := make([]*pfs1_12.Branch, 0, len(branches))
	for _, b := range branches {
		result = append(result, convert1_11Branch(b))
	}
	return result
}

func convert1_11Pipeline(p *pps1_11.Pipeline) *pps1_12.Pipeline {
	if p == nil {
		return nil
	}
	return &pps1_12.Pipeline{
		Name: p.Name,
	}
}

func convert1_11SecretMount(s *pps1_11.SecretMount) *pps1_12.SecretMount {
	if s == nil {
		return nil
	}
	return &pps1_12.SecretMount{
		Name:      s.Name,
		Key:       s.Key,
		MountPath: s.MountPath,
		EnvVar:    s.EnvVar,
	}
}

func convert1_11SecretMounts(secrets []*pps1_11.SecretMount) []*pps1_12.SecretMount {
	if secrets == nil {
		return nil
	}
	result := make([]*pps1_12.SecretMount, 0, len(secrets))
	for _, s := range secrets {
		result = append(result, convert1_11SecretMount(s))
	}
	return result
}

func convert1_11Transform(t *pps1_11.Transform) *pps1_12.Transform {
	if t == nil {
		return nil
	}
	return &pps1_12.Transform{
		Image:            t.Image,
		Cmd:              t.Cmd,
		ErrCmd:           t.ErrCmd,
		Env:              t.Env,
		Secrets:          convert1_11SecretMounts(t.Secrets),
		ImagePullSecrets: t.ImagePullSecrets,
		Stdin:            t.Stdin,
		ErrStdin:         t.ErrStdin,
		AcceptReturnCode: t.AcceptReturnCode,
		Debug:            t.Debug,
		User:             t.User,
		WorkingDir:       t.WorkingDir,
		Dockerfile:       t.Dockerfile,
	}
}

func convert1_11ParallelismSpec(s *pps1_11.ParallelismSpec) *pps1_12.ParallelismSpec {
	if s == nil {
		return nil
	}
	return &pps1_12.ParallelismSpec{
		Constant:    s.Constant,
		Coefficient: s.Coefficient,
	}
}

func convert1_11HashtreeSpec(h *pps1_11.HashtreeSpec) *pps1_12.HashtreeSpec {
	if h == nil {
		return nil
	}
	return &pps1_12.HashtreeSpec{
		Constant: h.Constant,
	}
}

func convert1_11Egress(e *pps1_11.Egress) *pps1_12.Egress {
	if e == nil {
		return nil
	}
	return &pps1_12.Egress{
		URL: e.URL,
	}
}

func convert1_11GPUSpec(g *pps1_11.GPUSpec) *pps1_12.GPUSpec {
	if g == nil {
		return nil
	}
	return &pps1_12.GPUSpec{
		Type:   g.Type,
		Number: g.Number,
	}
}

func convert1_11ResourceSpec(r *pps1_11.ResourceSpec) *pps1_12.ResourceSpec {
	if r == nil {
		return nil
	}
	return &pps1_12.ResourceSpec{
		Cpu:    r.Cpu,
		Memory: r.Memory,
		Gpu:    convert1_11GPUSpec(r.Gpu),
		Disk:   r.Disk,
	}
}

func convert1_11PFSInput(p *pps1_11.PFSInput) *pps1_12.PFSInput {
	if p == nil {
		return nil
	}
	return &pps1_12.PFSInput{
		Name:       p.Name,
		Repo:       p.Repo,
		Branch:     p.Branch,
		Commit:     p.Commit,
		Glob:       p.Glob,
		JoinOn:     p.JoinOn,
		Lazy:       p.Lazy,
		EmptyFiles: p.EmptyFiles,
	}
}

func convert1_11CronInput(i *pps1_11.CronInput) *pps1_12.CronInput {
	if i == nil {
		return nil
	}
	return &pps1_12.CronInput{
		Name:      i.Name,
		Repo:      i.Repo,
		Commit:    i.Commit,
		Spec:      i.Spec,
		Overwrite: i.Overwrite,
		Start:     i.Start,
	}
}

func convert1_11GitInput(i *pps1_11.GitInput) *pps1_12.GitInput {
	if i == nil {
		return nil
	}
	return &pps1_12.GitInput{
		Name:   i.Name,
		URL:    i.URL,
		Branch: i.Branch,
		Commit: i.Commit,
	}
}

func convert1_11Input(i *pps1_11.Input) *pps1_12.Input {
	if i == nil {
		return nil
	}
	return &pps1_12.Input{
		Pfs:   convert1_11PFSInput(i.Pfs),
		Cross: convert1_11Inputs(i.Cross),
		Union: convert1_11Inputs(i.Union),
		Cron:  convert1_11CronInput(i.Cron),
		Git:   convert1_11GitInput(i.Git),
	}
}

func convert1_11Inputs(inputs []*pps1_11.Input) []*pps1_12.Input {
	if inputs == nil {
		return nil
	}
	result := make([]*pps1_12.Input, 0, len(inputs))
	for _, i := range inputs {
		result = append(result, convert1_11Input(i))
	}
	return result
}

func convert1_11Service(s *pps1_11.Service) *pps1_12.Service {
	if s == nil {
		return nil
	}
	return &pps1_12.Service{
		InternalPort: s.InternalPort,
		ExternalPort: s.ExternalPort,
		IP:           s.IP,
		Type:         s.Type,
	}
}

func convert1_11Spout(s *pps1_11.Spout) *pps1_12.Spout {
	if s == nil {
		return nil
	}
	return &pps1_12.Spout{
		Overwrite: s.Overwrite,
		Service:   convert1_11Service(s.Service),
		Marker:    s.Marker,
	}
}

func convert1_11Metadata(s *pps1_11.Metadata) *pps1_12.Metadata {
	if s == nil {
		return nil
	}
	if s.Annotations == nil {
		return nil
	}
	return &pps1_12.Metadata{
		Annotations: s.Annotations,
	}
}

func convert1_11ChunkSpec(c *pps1_11.ChunkSpec) *pps1_12.ChunkSpec {
	if c == nil {
		return nil
	}
	return &pps1_12.ChunkSpec{
		Number:    c.Number,
		SizeBytes: c.SizeBytes,
	}
}

func convert1_11SchedulingSpec(s *pps1_11.SchedulingSpec) *pps1_12.SchedulingSpec {
	if s == nil {
		return nil
	}
	return &pps1_12.SchedulingSpec{
		NodeSelector:      s.NodeSelector,
		PriorityClassName: s.PriorityClassName,
	}
}

func convert1_11Acl(acl *auth1_11.SetACLRequest) *auth1_12.SetACLRequest {
	req := &auth1_12.SetACLRequest{
		Repo:    acl.Repo,
		Entries: make([]*auth1_12.ACLEntry, len(acl.Entries)),
	}

	for i, entry := range acl.Entries {
		req.Entries[i] = &auth1_12.ACLEntry{
			Username: entry.Username,
			Scope:    auth1_12.Scope(entry.Scope),
		}
	}
	return req
}

func convert1_11ClusterRoleBinding(bindings *auth1_11.ModifyClusterRoleBindingRequest) *auth1_12.ModifyClusterRoleBindingRequest {
	req := &auth1_12.ModifyClusterRoleBindingRequest{
		Principal: bindings.Principal,
		Roles: &auth1_12.ClusterRoles{
			Roles: make([]auth1_12.ClusterRole, len(bindings.Roles.Roles)),
		},
	}

	for i, role := range bindings.Roles.Roles {
		req.Roles.Roles[i] = auth1_12.ClusterRole(role)
	}
	return req
}

func convert1_11AuthConfig(config *auth1_11.SetConfigurationRequest) *auth1_12.SetConfigurationRequest {
	req := &auth1_12.SetConfigurationRequest{
		Configuration: &auth1_12.AuthConfig{
			LiveConfigVersion: config.Configuration.LiveConfigVersion,
		},
	}
	if config.Configuration.SAMLServiceOptions != nil {
		req.Configuration.SAMLServiceOptions = &auth1_12.AuthConfig_SAMLServiceOptions{
			ACSURL:          config.Configuration.SAMLServiceOptions.ACSURL,
			MetadataURL:     config.Configuration.SAMLServiceOptions.MetadataURL,
			DashURL:         config.Configuration.SAMLServiceOptions.DashURL,
			SessionDuration: config.Configuration.SAMLServiceOptions.SessionDuration,
			DebugLogging:    config.Configuration.SAMLServiceOptions.DebugLogging,
		}
	}
	return req
}

func convert1_11Op(op *admin.Op1_11) (*admin.Op1_12, error) {
	switch {
	case op.CreateObject != nil:
		return &admin.Op1_12{
			CreateObject: convert1_11CreateObject(op.CreateObject),
		}, nil
	case op.Job != nil:
		return &admin.Op1_12{
			Job: convert1_11Job(op.Job),
		}, nil
	case op.Tag != nil:
		if !objHashRE.MatchString(op.Tag.Object.Hash) {
			return nil, errors.Errorf("invalid object hash in op: %q", op)
		}
		return &admin.Op1_12{
			Tag: &pfs1_12.TagObjectRequest{
				Object: convert1_11Object(op.Tag.Object),
				Tags:   convert1_11Tags(op.Tag.Tags),
			},
		}, nil
	case op.Repo != nil:
		return &admin.Op1_12{
			Repo: &pfs1_12.CreateRepoRequest{
				Repo:        convert1_11Repo(op.Repo.Repo),
				Description: op.Repo.Description,
			},
		}, nil
	case op.Commit != nil:
		return &admin.Op1_12{
			Commit: &pfs1_12.BuildCommitRequest{
				Parent:     convert1_11Commit(op.Commit.Parent),
				Branch:     op.Commit.Branch,
				Provenance: convert1_11Provenances(op.Commit.Provenance),
				Tree:       convert1_11Object(op.Commit.Tree),
				Trees:      convert1_11Objects(op.Commit.Trees),
				Datums:     convert1_11Object(op.Commit.Datums),
				ID:         op.Commit.ID,
				SizeBytes:  op.Commit.SizeBytes,
			},
		}, nil
	case op.Branch != nil:
		newOp := &admin.Op1_12{
			Branch: &pfs1_12.CreateBranchRequest{
				Head:       convert1_11Commit(op.Branch.Head),
				Branch:     convert1_11Branch(op.Branch.Branch),
				Provenance: convert1_11Branches(op.Branch.Provenance),
			},
		}
		if newOp.Branch.Branch == nil {
			newOp.Branch.Branch = &pfs1_12.Branch{
				Repo: convert1_11Repo(op.Branch.Head.Repo),
				Name: op.Branch.SBranch,
			}
		}
		return newOp, nil
	case op.Pipeline != nil:
		return &admin.Op1_12{
			Pipeline: &pps1_12.CreatePipelineRequest{
				Pipeline:         convert1_11Pipeline(op.Pipeline.Pipeline),
				Transform:        convert1_11Transform(op.Pipeline.Transform),
				ParallelismSpec:  convert1_11ParallelismSpec(op.Pipeline.ParallelismSpec),
				HashtreeSpec:     convert1_11HashtreeSpec(op.Pipeline.HashtreeSpec),
				Egress:           convert1_11Egress(op.Pipeline.Egress),
				Update:           op.Pipeline.Update,
				OutputBranch:     op.Pipeline.OutputBranch,
				ResourceRequests: convert1_11ResourceSpec(op.Pipeline.ResourceRequests),
				ResourceLimits:   convert1_11ResourceSpec(op.Pipeline.ResourceLimits),
				Input:            convert1_11Input(op.Pipeline.Input),
				Description:      op.Pipeline.Description,
				CacheSize:        op.Pipeline.CacheSize,
				EnableStats:      op.Pipeline.EnableStats,
				Reprocess:        op.Pipeline.Reprocess,
				MaxQueueSize:     op.Pipeline.MaxQueueSize,
				Service:          convert1_11Service(op.Pipeline.Service),
				Spout:            convert1_11Spout(op.Pipeline.Spout),
				ChunkSpec:        convert1_11ChunkSpec(op.Pipeline.ChunkSpec),
				DatumTimeout:     op.Pipeline.DatumTimeout,
				JobTimeout:       op.Pipeline.JobTimeout,
				Salt:             op.Pipeline.Salt,
				Standby:          op.Pipeline.Standby,
				DatumTries:       op.Pipeline.DatumTries,
				SchedulingSpec:   convert1_11SchedulingSpec(op.Pipeline.SchedulingSpec),
				PodSpec:          op.Pipeline.PodSpec,
				PodPatch:         op.Pipeline.PodPatch,
				SpecCommit:       convert1_11Commit(op.Pipeline.SpecCommit),
				Metadata:         convert1_11Metadata(op.Pipeline.Metadata),
			},
		}, nil
	case op.SetAcl != nil:
		return &admin.Op1_12{
			SetAcl: convert1_11Acl(op.SetAcl),
		}, nil
	case op.SetClusterRoleBinding != nil:
		return &admin.Op1_12{
			SetClusterRoleBinding: convert1_11ClusterRoleBinding(op.SetClusterRoleBinding),
		}, nil
	case op.SetAuthConfig != nil:
		return &admin.Op1_12{
			SetAuthConfig: convert1_11AuthConfig(op.SetAuthConfig),
		}, nil
	case op.ActivateAuth != nil:
		return &admin.Op1_12{
			ActivateAuth: &auth1_12.ActivateRequest{},
		}, nil
	case op.RestoreAuthToken != nil:
		return &admin.Op1_12{
			RestoreAuthToken: &auth1_12.RestoreAuthTokenRequest{
				Token: &auth1_12.HashedAuthToken{
					HashedToken: op.RestoreAuthToken.Token.HashedToken,
					Expiration:  op.RestoreAuthToken.Token.Expiration,
					TokenInfo: &auth1_12.TokenInfo{
						Subject: op.RestoreAuthToken.Token.TokenInfo.Subject,
						Source:  auth1_12.TokenInfo_TokenSource(op.RestoreAuthToken.Token.TokenInfo.Source),
					},
				},
			},
		}, nil
	case op.ActivateEnterprise != nil:
		return &admin.Op1_12{
			ActivateEnterprise: &enterprise1_12.ActivateRequest{
				ActivationCode: op.ActivateEnterprise.ActivationCode,
			},
		}, nil
	case op.CheckAuthToken != nil:
		return &admin.Op1_12{
			CheckAuthToken: &admin.CheckAuthToken{},
		}, nil
	default:
		return nil, errors.Errorf("unrecognized 1.9 op type:\n%+v", op)
	}
	return nil, errors.Errorf("internal error: convert1_11Op() didn't return a 1.9 op for:\n%+v", op)
}
