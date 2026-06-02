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

// Endpoints that allow you to search through the user&#039;s financial data.
// Different from the autocomplete endpoints, the search accepts more advanced
// arguments.
//
// SearchService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.options = opts
	return
}

// Search for accounts
func (r *SearchService) Accounts(ctx context.Context, params SearchAccountsParams, opts ...option.RequestOption) (res *AccountArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/search/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Searches through the users transactions.
func (r *SearchService) Transactions(ctx context.Context, params SearchTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/search/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type SearchAccountsParams struct {
	// The account field(s) you want to search in.
	//
	// Any of "all", "iban", "name", "number", "id".
	Field SearchAccountsParamsField `query:"field,omitzero" api:"required" json:"-"`
	// The query you wish to search for.
	Query string `query:"query" api:"required" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "all", "asset", "cash", "expense", "revenue", "special", "hidden",
	// "liability", "liabilities", "Default account", "Cash account", "Asset account",
	// "Expense account", "Revenue account", "Initial balance account", "Beneficiary
	// account", "Import account", "Reconciliation account", "Loan", "Debt",
	// "Mortgage".
	Type AccountTypeFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchAccountsParams]'s query parameters as `url.Values`.
func (r SearchAccountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The account field(s) you want to search in.
type SearchAccountsParamsField string

const (
	SearchAccountsParamsFieldAll    SearchAccountsParamsField = "all"
	SearchAccountsParamsFieldIban   SearchAccountsParamsField = "iban"
	SearchAccountsParamsFieldName   SearchAccountsParamsField = "name"
	SearchAccountsParamsFieldNumber SearchAccountsParamsField = "number"
	SearchAccountsParamsFieldID     SearchAccountsParamsField = "id"
)

type SearchTransactionsParams struct {
	// The query you wish to search for.
	Query string `query:"query" api:"required" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [SearchTransactionsParams]'s query parameters as
// `url.Values`.
func (r SearchTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
