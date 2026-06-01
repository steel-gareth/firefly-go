// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apijson"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apiquery"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/respjson"
)

// The &quot;insight&quot; endpoints try to deliver sums, balances and insightful
// information in the broadest sense of the word.
//
// InsightTransferService contains methods and other services that help with
// interacting with the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightTransferService] method instead.
type InsightTransferService struct {
	options []option.RequestOption
}

// NewInsightTransferService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInsightTransferService(opts ...option.RequestOption) (r InsightTransferService) {
	r = InsightTransferService{}
	r.options = opts
	return
}

// This endpoint gives a sum of the total amount transfers made by the user.
func (r *InsightTransferService) GetTotal(ctx context.Context, params InsightTransferGetTotalParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/transfer/total"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the transfers made by the user, grouped by
// asset account or lability.
func (r *InsightTransferService) ListByAssetAccount(ctx context.Context, params InsightTransferListByAssetAccountParams, opts ...option.RequestOption) (res *[]InsightTransferListByAssetAccountResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/transfer/asset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the transfers made by the user, grouped by
// (any) category.
func (r *InsightTransferService) ListByCategory(ctx context.Context, params InsightTransferListByCategoryParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/transfer/category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the transfers created by the user, grouped by
// (any) tag.
func (r *InsightTransferService) ListByTag(ctx context.Context, params InsightTransferListByTagParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/transfer/tag"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the transfers made by the user, including only
// transfers with no category.
func (r *InsightTransferService) ListWithoutCategory(ctx context.Context, params InsightTransferListWithoutCategoryParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/transfer/no-category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the transfers made by the user, including only
// transfers with no tag.
func (r *InsightTransferService) ListWithoutTag(ctx context.Context, params InsightTransferListWithoutTagParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/transfer/no-tag"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type InsightTransferListByAssetAccountResponse struct {
	// This ID is a reference to the original object.
	ID string `json:"id"`
	// The currency code of the expenses listed for this account.
	CurrencyCode string `json:"currency_code"`
	// The currency ID of the expenses listed for this account.
	CurrencyID string `json:"currency_id"`
	// The total amount transferred between start date and end date, a number defined
	// as a string, for this asset account.
	Difference string `json:"difference"`
	// The total amount transferred between start date and end date, a number as a
	// float, for this asset account. May have rounding errors.
	DifferenceFloat float64 `json:"difference_float"`
	// The total amount transferred TO this account between start date and end date, a
	// number defined as a string, for this asset account.
	In string `json:"in"`
	// The total amount transferred FROM this account between start date and end date,
	// a number as a float, for this asset account. May have rounding errors.
	InFloat float64 `json:"in_float"`
	// This is the name of the object.
	Name string `json:"name"`
	// The total amount transferred FROM this account between start date and end date,
	// a number defined as a string, for this asset account.
	Out string `json:"out"`
	// The total amount transferred TO this account between start date and end date, a
	// number as a float, for this asset account. May have rounding errors.
	OutFloat float64 `json:"out_float"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CurrencyCode    respjson.Field
		CurrencyID      respjson.Field
		Difference      respjson.Field
		DifferenceFloat respjson.Field
		In              respjson.Field
		InFloat         respjson.Field
		Name            respjson.Field
		Out             respjson.Field
		OutFloat        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightTransferListByAssetAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *InsightTransferListByAssetAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightTransferGetTotalParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only transfers between those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightTransferGetTotalParams]'s query parameters as
// `url.Values`.
func (r InsightTransferGetTotalParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightTransferListByAssetAccountParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only transfers between those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightTransferListByAssetAccountParams]'s query parameters
// as `url.Values`.
func (r InsightTransferListByAssetAccountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightTransferListByCategoryParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only transfers between those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	// The categories to be included in the results.
	Categories []int64 `query:"categories,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightTransferListByCategoryParams]'s query parameters as
// `url.Values`.
func (r InsightTransferListByCategoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightTransferListByTagParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only transfers between those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	// The tags to be included in the results.
	Tags []int64 `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightTransferListByTagParams]'s query parameters as
// `url.Values`.
func (r InsightTransferListByTagParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightTransferListWithoutCategoryParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only transfers between those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightTransferListWithoutCategoryParams]'s query
// parameters as `url.Values`.
func (r InsightTransferListWithoutCategoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightTransferListWithoutTagParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only transfers from those asset accounts / liabilities
	// will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightTransferListWithoutTagParams]'s query parameters as
// `url.Values`.
func (r InsightTransferListWithoutTagParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
