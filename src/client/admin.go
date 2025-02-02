package client

import (
	"io"

	"github.com/gogo/protobuf/types"
	"github.com/pachyderm/pachyderm/src/client/admin"
	"github.com/pachyderm/pachyderm/src/client/pkg/errors"
	"github.com/pachyderm/pachyderm/src/client/pkg/grpcutil"
	"github.com/pachyderm/pachyderm/src/client/pkg/pbutil"
	"github.com/pachyderm/pachyderm/src/client/pps"
)

// InspectCluster retrieves cluster state
func (c APIClient) InspectCluster() (*admin.ClusterInfo, error) {
	clusterInfo, err := c.AdminAPIClient.InspectCluster(c.Ctx(), &types.Empty{})
	if err != nil {
		return nil, grpcutil.ScrubGRPC(err)
	}
	return clusterInfo, nil
}

// Extract all cluster state, call f with each operation.
func (c APIClient) Extract(objects, enterprise, auth bool, f func(op *admin.Op) error) error {
	extractClient, err := c.AdminAPIClient.Extract(c.Ctx(), &admin.ExtractRequest{
		NoObjects:    !objects,
		NoEnterprise: !enterprise,
		NoAuth:       !auth,
	})
	if err != nil {
		return grpcutil.ScrubGRPC(err)
	}
	for {
		op, err := extractClient.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return grpcutil.ScrubGRPC(err)
		}
		if err := f(op); err != nil {
			return err
		}
	}
	return nil
}

// ExtractAll cluster state as a slice of operations.
func (c APIClient) ExtractAll(objects, enterprise, auth bool) ([]*admin.Op, error) {
	var result []*admin.Op
	if err := c.Extract(objects, enterprise, auth, func(op *admin.Op) error {
		result = append(result, op)
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}

// ExtractWriter extracts all cluster state and marshals it to w.
func (c APIClient) ExtractWriter(objects, enterprise, auth bool, w io.Writer) error {
	writer := pbutil.NewWriter(w)
	return c.Extract(objects, enterprise, auth, func(op *admin.Op) error {
		_, err := writer.Write(op)
		return err
	})
}

// ExtractURL extracts all cluster state and marshalls it to object storage.
func (c APIClient) ExtractURL(url string, enterprise, auth bool) error {
	extractClient, err := c.AdminAPIClient.Extract(c.Ctx(), &admin.ExtractRequest{
		URL:          url,
		NoEnterprise: !enterprise,
		NoAuth:       !auth,
	})
	if err != nil {
		return grpcutil.ScrubGRPC(err)
	}
	resp, err := extractClient.Recv()
	if err == nil {
		return errors.Errorf("unexpected response from extract: %v", resp)
	}
	if !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}

// ExtractPipeline extracts a single pipeline.
func (c APIClient) ExtractPipeline(pipelineName string) (*pps.CreatePipelineRequest, error) {
	op, err := c.AdminAPIClient.ExtractPipeline(c.Ctx(), &admin.ExtractPipelineRequest{Pipeline: NewPipeline(pipelineName)})
	if err != nil {
		return nil, grpcutil.ScrubGRPC(err)
	}
	if op.Op2_0 == nil || op.Op2_0.Pipeline == nil {
		return nil, errors.Errorf("malformed response is missing pipeline")
	}
	return op.Op2_0.Pipeline, nil
}

// Restore cluster state from an extract series of operations.
func (c APIClient) Restore(ops []*admin.Op) (retErr error) {
	restoreClient, err := c.AdminAPIClient.Restore(c.Ctx())
	if err != nil {
		return grpcutil.ScrubGRPC(err)
	}
	defer func() {
		if _, err := restoreClient.CloseAndRecv(); err != nil && retErr == nil {
			retErr = grpcutil.ScrubGRPC(err)
		}
	}()
	for _, op := range ops {
		if err := restoreClient.Send(&admin.RestoreRequest{Op: op}); err != nil {
			return grpcutil.ScrubGRPC(err)
		}
	}
	return nil
}

// RestoreReader restores cluster state from a reader containing marshaled ops.
// Such as those written by ExtractWriter.
func (c APIClient) RestoreReader(r io.Reader) (retErr error) {
	restoreClient, err := c.AdminAPIClient.Restore(c.Ctx())
	if err != nil {
		return grpcutil.ScrubGRPC(err)
	}
	defer func() {
		if _, err := restoreClient.CloseAndRecv(); err != nil && retErr == nil {
			retErr = grpcutil.ScrubGRPC(err)
		}
	}()
	reader := pbutil.NewReader(r)
	op := &admin.Op{}
	for {
		if err := reader.Read(op); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if err := restoreClient.Send(&admin.RestoreRequest{Op: op}); err != nil {
			return grpcutil.ScrubGRPC(err)
		}
	}
	return nil
}

// RestoreFrom restores state from another cluster which can be access through otherC.
func (c APIClient) RestoreFrom(objects, enterprise, auth bool, otherC *APIClient) (retErr error) {
	restoreClient, err := c.AdminAPIClient.Restore(c.Ctx())
	if err != nil {
		return grpcutil.ScrubGRPC(err)
	}
	defer func() {
		if _, err := restoreClient.CloseAndRecv(); err != nil && retErr == nil {
			retErr = grpcutil.ScrubGRPC(err)
		}
	}()
	return otherC.Extract(objects, enterprise, auth, func(op *admin.Op) error {
		return restoreClient.Send(&admin.RestoreRequest{Op: op})
	})
}

// RestoreURL restures cluster state from object storage.
func (c APIClient) RestoreURL(url string) (retErr error) {
	restoreClient, err := c.AdminAPIClient.Restore(c.Ctx())
	if err != nil {
		return grpcutil.ScrubGRPC(err)
	}
	defer func() {
		if _, err := restoreClient.CloseAndRecv(); err != nil && retErr == nil {
			retErr = grpcutil.ScrubGRPC(err)
		}
	}()
	return grpcutil.ScrubGRPC(restoreClient.Send(&admin.RestoreRequest{URL: url}))
}
