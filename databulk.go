// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/steel-gareth/firefly-go/internal/apiquery"
	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
)

// The &quot;data&quot;-endpoints manage generic Firefly III and user-specific
// data.
//
// DataBulkService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataBulkService] method instead.
type DataBulkService struct {
	options []option.RequestOption
}

// NewDataBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataBulkService(opts ...option.RequestOption) (r DataBulkService) {
	r = DataBulkService{}
	r.options = opts
	return
}

// Allows you to update transactions in bulk.
func (r *DataBulkService) UpdateTransactions(ctx context.Context, params DataBulkUpdateTransactionsParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/data/bulk/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type DataBulkUpdateTransactionsParams struct {
	// The JSON query.
	Query    string            `query:"query" api:"required" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [DataBulkUpdateTransactionsParams]'s query parameters as
// `url.Values`.
func (r DataBulkUpdateTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
