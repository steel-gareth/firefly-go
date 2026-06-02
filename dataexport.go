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

// The &quot;data&quot;-endpoints manage generic Firefly III and user-specific
// data.
//
// DataExportService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataExportService] method instead.
type DataExportService struct {
	options []option.RequestOption
}

// NewDataExportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataExportService(opts ...option.RequestOption) (r DataExportService) {
	r = DataExportService{}
	r.options = opts
	return
}

// This endpoint allows you to export your accounts from Firefly III into a file.
// Currently supports CSV exports only.
func (r *DataExportService) ExportAccounts(ctx context.Context, params DataExportExportAccountsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "v1/data/export/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint allows you to export your bills from Firefly III into a file.
// Currently supports CSV exports only.
func (r *DataExportService) ExportBills(ctx context.Context, params DataExportExportBillsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "v1/data/export/bills"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint allows you to export your budgets and associated budget data from
// Firefly III into a file. Currently supports CSV exports only.
func (r *DataExportService) ExportBudgets(ctx context.Context, params DataExportExportBudgetsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "v1/data/export/budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint allows you to export your categories from Firefly III into a file.
// Currently supports CSV exports only.
func (r *DataExportService) ExportCategories(ctx context.Context, params DataExportExportCategoriesParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "v1/data/export/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint allows you to export your piggy banks from Firefly III into a
// file. Currently supports CSV exports only.
func (r *DataExportService) ExportPiggyBanks(ctx context.Context, params DataExportExportPiggyBanksParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "v1/data/export/piggy-banks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint allows you to export your recurring transactions from Firefly III
// into a file. Currently supports CSV exports only.
func (r *DataExportService) ExportRecurring(ctx context.Context, params DataExportExportRecurringParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "v1/data/export/recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint allows you to export your rules and rule groups from Firefly III
// into a file. Currently supports CSV exports only.
func (r *DataExportService) ExportRules(ctx context.Context, params DataExportExportRulesParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "v1/data/export/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint allows you to export your tags from Firefly III into a file.
// Currently supports CSV exports only.
func (r *DataExportService) ExportTags(ctx context.Context, params DataExportExportTagsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "v1/data/export/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint allows you to export transactions from Firefly III into a file.
// Currently supports CSV exports only.
func (r *DataExportService) ExportTransactions(ctx context.Context, params DataExportExportTransactionsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "v1/data/export/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ExportFileFilter string

const (
	ExportFileFilterCsv ExportFileFilter = "csv"
)

type DataExportExportAccountsParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "csv".
	Type ExportFileFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataExportExportAccountsParams]'s query parameters as
// `url.Values`.
func (r DataExportExportAccountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportExportBillsParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "csv".
	Type ExportFileFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataExportExportBillsParams]'s query parameters as
// `url.Values`.
func (r DataExportExportBillsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportExportBudgetsParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "csv".
	Type ExportFileFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataExportExportBudgetsParams]'s query parameters as
// `url.Values`.
func (r DataExportExportBudgetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportExportCategoriesParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "csv".
	Type ExportFileFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataExportExportCategoriesParams]'s query parameters as
// `url.Values`.
func (r DataExportExportCategoriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportExportPiggyBanksParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "csv".
	Type ExportFileFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataExportExportPiggyBanksParams]'s query parameters as
// `url.Values`.
func (r DataExportExportPiggyBanksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportExportRecurringParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "csv".
	Type ExportFileFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataExportExportRecurringParams]'s query parameters as
// `url.Values`.
func (r DataExportExportRecurringParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportExportRulesParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "csv".
	Type ExportFileFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataExportExportRulesParams]'s query parameters as
// `url.Values`.
func (r DataExportExportRulesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportExportTagsParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "csv".
	Type ExportFileFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataExportExportTagsParams]'s query parameters as
// `url.Values`.
func (r DataExportExportTagsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportExportTransactionsParams struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start time.Time `query:"start" api:"required" format:"date" json:"-"`
	// Limit the export of transactions to these accounts only. Only asset accounts
	// will be accepted. Other types will be silently dropped.
	Accounts param.Opt[string] `query:"accounts,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "csv".
	Type ExportFileFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataExportExportTransactionsParams]'s query parameters as
// `url.Values`.
func (r DataExportExportTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
