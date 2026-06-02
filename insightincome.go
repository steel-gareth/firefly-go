// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/steel-gareth/firefly-go/internal/apiquery"
	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
)

// The &quot;insight&quot; endpoints try to deliver sums, balances and insightful
// information in the broadest sense of the word.
//
// InsightIncomeService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightIncomeService] method instead.
type InsightIncomeService struct {
	options []option.RequestOption
}

// NewInsightIncomeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInsightIncomeService(opts ...option.RequestOption) (r InsightIncomeService) {
	r = InsightIncomeService{}
	r.options = opts
	return
}

// This endpoint gives a sum of the total income received by the user.
func (r *InsightIncomeService) GetTotal(ctx context.Context, params InsightIncomeGetTotalParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/income/total"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the income received by the user, grouped by
// asset account.
func (r *InsightIncomeService) ListByAssetAccount(ctx context.Context, params InsightIncomeListByAssetAccountParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/income/asset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the income received by the user, grouped by
// (any) category.
func (r *InsightIncomeService) ListByCategory(ctx context.Context, params InsightIncomeListByCategoryParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/income/category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the income received by the user, grouped by
// revenue account.
func (r *InsightIncomeService) ListByRevenueAccount(ctx context.Context, params InsightIncomeListByRevenueAccountParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/income/revenue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the income received by the user, grouped by
// (any) tag.
func (r *InsightIncomeService) ListByTag(ctx context.Context, params InsightIncomeListByTagParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/income/tag"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the income received by the user, including only
// income with no category.
func (r *InsightIncomeService) ListWithoutCategory(ctx context.Context, params InsightIncomeListWithoutCategoryParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/income/no-category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the income received by the user, including only
// income with no tag.
func (r *InsightIncomeService) ListWithoutTag(ctx context.Context, params InsightIncomeListWithoutTagParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/income/no-tag"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type InsightIncomeGetTotalParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only deposits to those asset accounts / liabilities
	// will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightIncomeGetTotalParams]'s query parameters as
// `url.Values`.
func (r InsightIncomeGetTotalParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightIncomeListByAssetAccountParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only deposits to those asset accounts / liabilities
	// will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightIncomeListByAssetAccountParams]'s query parameters
// as `url.Values`.
func (r InsightIncomeListByAssetAccountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightIncomeListByCategoryParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only deposits to those asset accounts / liabilities
	// will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	// The categories to be included in the results.
	Categories []int64 `query:"categories,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightIncomeListByCategoryParams]'s query parameters as
// `url.Values`.
func (r InsightIncomeListByCategoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightIncomeListByRevenueAccountParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you add the accounts ID's of
	// revenue accounts, only those accounts are included in the results. If you
	// include ID's of asset accounts or liabilities, only deposits to those asset
	// accounts / liabilities will be included. You can combine both asset / liability
	// and deposit account ID's. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightIncomeListByRevenueAccountParams]'s query parameters
// as `url.Values`.
func (r InsightIncomeListByRevenueAccountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightIncomeListByTagParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only deposits to those asset accounts / liabilities
	// will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	// The tags to be included in the results.
	Tags []int64 `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightIncomeListByTagParams]'s query parameters as
// `url.Values`.
func (r InsightIncomeListByTagParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightIncomeListWithoutCategoryParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only deposits to those asset accounts / liabilities
	// will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightIncomeListWithoutCategoryParams]'s query parameters
// as `url.Values`.
func (r InsightIncomeListWithoutCategoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightIncomeListWithoutTagParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only deposits to those asset accounts / liabilities
	// will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightIncomeListWithoutTagParams]'s query parameters as
// `url.Values`.
func (r InsightIncomeListWithoutTagParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
