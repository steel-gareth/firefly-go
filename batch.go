// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
)

// These endpoints deliver general system information, version- and meta
// information.
//
// BatchService contains methods and other services that help with interacting with
// the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchService] method instead.
type BatchService struct {
	options []option.RequestOption
}

// NewBatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBatchService(opts ...option.RequestOption) (r BatchService) {
	r = BatchService{}
	r.options = opts
	return
}

// summary: Finish a batch of unprocessed transactions.
func (r *BatchService) Finish(ctx context.Context, body BatchFinishParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/batch/finish"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type BatchFinishParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
