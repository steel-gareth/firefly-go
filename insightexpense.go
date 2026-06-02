// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/steel-gareth/firefly-go/internal/apijson"
	"github.com/steel-gareth/firefly-go/internal/apiquery"
	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
	"github.com/steel-gareth/firefly-go/packages/respjson"
)

// The &quot;insight&quot; endpoints try to deliver sums, balances and insightful
// information in the broadest sense of the word.
//
// InsightExpenseService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightExpenseService] method instead.
type InsightExpenseService struct {
	options []option.RequestOption
}

// NewInsightExpenseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInsightExpenseService(opts ...option.RequestOption) (r InsightExpenseService) {
	r = InsightExpenseService{}
	r.options = opts
	return
}

// This endpoint gives a sum of the total expenses made by the user.
func (r *InsightExpenseService) GetTotal(ctx context.Context, params InsightExpenseGetTotalParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/total"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, grouped by asset
// account.
func (r *InsightExpenseService) ListByAssetAccount(ctx context.Context, params InsightExpenseListByAssetAccountParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/asset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, grouped by (any)
// bill.
func (r *InsightExpenseService) ListByBill(ctx context.Context, params InsightExpenseListByBillParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/bill"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, grouped by (any)
// budget.
func (r *InsightExpenseService) ListByBudget(ctx context.Context, params InsightExpenseListByBudgetParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/budget"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, grouped by (any)
// category.
func (r *InsightExpenseService) ListByCategory(ctx context.Context, params InsightExpenseListByCategoryParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, grouped by
// expense account.
func (r *InsightExpenseService) ListByExpenseAccount(ctx context.Context, params InsightExpenseListByExpenseAccountParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/expense"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, grouped by (any)
// tag.
func (r *InsightExpenseService) ListByTag(ctx context.Context, params InsightExpenseListByTagParams, opts ...option.RequestOption) (res *[]InsightGroupEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/tag"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, including only
// expenses with no bill.
func (r *InsightExpenseService) ListWithoutBill(ctx context.Context, params InsightExpenseListWithoutBillParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/no-bill"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, including only
// expenses with no budget.
func (r *InsightExpenseService) ListWithoutBudget(ctx context.Context, params InsightExpenseListWithoutBudgetParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/no-budget"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, including only
// expenses with no category.
func (r *InsightExpenseService) ListWithoutCategory(ctx context.Context, params InsightExpenseListWithoutCategoryParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/no-category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint gives a summary of the expenses made by the user, including only
// expenses with no tag.
func (r *InsightExpenseService) ListWithoutTag(ctx context.Context, params InsightExpenseListWithoutTagParams, opts ...option.RequestOption) (res *[]InsightTotalEntry, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/insight/expense/no-tag"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type InsightGroupEntry struct {
	// This ID is a reference to the original object.
	ID string `json:"id"`
	// The currency code of the expenses listed for this account.
	CurrencyCode string `json:"currency_code"`
	// The currency ID of the expenses listed for this account.
	CurrencyID string `json:"currency_id"`
	// The amount spent or earned between start date and end date, a number defined as
	// a string, for this object and all asset accounts.
	Difference string `json:"difference"`
	// The amount spent or earned between start date and end date, a number as a float,
	// for this object and all asset accounts. May have rounding errors.
	DifferenceFloat float64 `json:"difference_float"`
	// This is the name of the object.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CurrencyCode    respjson.Field
		CurrencyID      respjson.Field
		Difference      respjson.Field
		DifferenceFloat respjson.Field
		Name            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGroupEntry) RawJSON() string { return r.JSON.raw }
func (r *InsightGroupEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightTotalEntry struct {
	// The currency code of the expenses listed for this expense account.
	CurrencyCode string `json:"currency_code"`
	// The currency ID of the expenses listed for this expense account.
	CurrencyID string `json:"currency_id"`
	// The amount spent between start date and end date, defined as a string, for this
	// expense account and all asset accounts.
	Difference string `json:"difference"`
	// The amount spent between start date and end date, defined as a string, for this
	// expense account and all asset accounts. This number is a float (double) and may
	// have rounding errors.
	DifferenceFloat float64 `json:"difference_float"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyCode    respjson.Field
		CurrencyID      respjson.Field
		Difference      respjson.Field
		DifferenceFloat respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightTotalEntry) RawJSON() string { return r.JSON.raw }
func (r *InsightTotalEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightExpenseGetTotalParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseGetTotalParams]'s query parameters as
// `url.Values`.
func (r InsightExpenseGetTotalParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListByAssetAccountParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListByAssetAccountParams]'s query parameters
// as `url.Values`.
func (r InsightExpenseListByAssetAccountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListByBillParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	// The bills to be included in the results.
	Bills []int64 `query:"bills,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListByBillParams]'s query parameters as
// `url.Values`.
func (r InsightExpenseListByBillParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListByBudgetParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	// The budgets to be included in the results.
	Budgets []int64 `query:"budgets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListByBudgetParams]'s query parameters as
// `url.Values`.
func (r InsightExpenseListByBudgetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListByCategoryParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	// The categories to be included in the results.
	Categories []int64 `query:"categories,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListByCategoryParams]'s query parameters as
// `url.Values`.
func (r InsightExpenseListByCategoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListByExpenseAccountParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you add the accounts ID's of
	// expense accounts, only those accounts are included in the results. If you
	// include ID's of asset accounts or liabilities, only withdrawals from those asset
	// accounts / liabilities will be included. You can combine both asset / liability
	// and expense account ID's. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListByExpenseAccountParams]'s query
// parameters as `url.Values`.
func (r InsightExpenseListByExpenseAccountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListByTagParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	// The tags to be included in the results.
	Tags []int64 `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListByTagParams]'s query parameters as
// `url.Values`.
func (r InsightExpenseListByTagParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListWithoutBillParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListWithoutBillParams]'s query parameters as
// `url.Values`.
func (r InsightExpenseListWithoutBillParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListWithoutBudgetParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListWithoutBudgetParams]'s query parameters
// as `url.Values`.
func (r InsightExpenseListWithoutBudgetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListWithoutCategoryParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListWithoutCategoryParams]'s query parameters
// as `url.Values`.
func (r InsightExpenseListWithoutCategoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightExpenseListWithoutTagParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// The accounts to be included in the results. If you include ID's of asset
	// accounts or liabilities, only withdrawals from those asset accounts /
	// liabilities will be included. Other account ID's will be ignored.
	Accounts []int64 `query:"accounts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightExpenseListWithoutTagParams]'s query parameters as
// `url.Values`.
func (r InsightExpenseListWithoutTagParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
