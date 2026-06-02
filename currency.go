// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"errors"
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

// Endpoints to manage the currencies in Firefly III. Depending on the user&#039;s
// role you can also disable and enable them, or add new ones.
//
// CurrencyService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrencyService] method instead.
type CurrencyService struct {
	options []option.RequestOption
	// Endpoints to manage the currencies in Firefly III. Depending on the user&#039;s
	// role you can also disable and enable them, or add new ones.
	Primary CurrencyPrimaryService
}

// NewCurrencyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCurrencyService(opts ...option.RequestOption) (r CurrencyService) {
	r = CurrencyService{}
	r.options = opts
	r.Primary = NewCurrencyPrimaryService(opts...)
	return
}

// Creates a new currency. The data required can be submitted as a JSON body or as
// a list of parameters.
func (r *CurrencyService) New(ctx context.Context, params CurrencyNewParams, opts ...option.RequestOption) (res *CurrencySingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/currencies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single currency.
func (r *CurrencyService) Get(ctx context.Context, code string, query CurrencyGetParams, opts ...option.RequestOption) (res *CurrencySingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update existing currency.
func (r *CurrencyService) Update(ctx context.Context, code string, params CurrencyUpdateParams, opts ...option.RequestOption) (res *CurrencySingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all currencies.
func (r *CurrencyService) List(ctx context.Context, params CurrencyListParams, opts ...option.RequestOption) (res *CurrencyListResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/currencies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a currency.
func (r *CurrencyService) Delete(ctx context.Context, code string, body CurrencyDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return err
	}
	path := fmt.Sprintf("v1/currencies/%s", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Disable a currency.
func (r *CurrencyService) Disable(ctx context.Context, code string, body CurrencyDisableParams, opts ...option.RequestOption) (res *CurrencySingle, err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/disable", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Enable a single currency.
func (r *CurrencyService) Enable(ctx context.Context, code string, body CurrencyEnableParams, opts ...option.RequestOption) (res *CurrencySingle, err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/enable", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// List all accounts with this currency.
func (r *CurrencyService) ListAccounts(ctx context.Context, code string, params CurrencyListAccountsParams, opts ...option.RequestOption) (res *AccountArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/accounts", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all available budgets with this currency.
func (r *CurrencyService) ListAvailableBudgets(ctx context.Context, code string, params CurrencyListAvailableBudgetsParams, opts ...option.RequestOption) (res *AvailableBudgetArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/available-budgets", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all bills with this currency.
func (r *CurrencyService) ListBills(ctx context.Context, code string, params CurrencyListBillsParams, opts ...option.RequestOption) (res *BillArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/bills", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all budget limits with this currency
func (r *CurrencyService) ListBudgetLimits(ctx context.Context, code string, params CurrencyListBudgetLimitsParams, opts ...option.RequestOption) (res *BudgetLimitArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/budget-limits", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all recurring transactions with this currency.
func (r *CurrencyService) ListRecurrences(ctx context.Context, code string, params CurrencyListRecurrencesParams, opts ...option.RequestOption) (res *RecurrenceArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/recurrences", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all rules with this currency.
func (r *CurrencyService) ListRules(ctx context.Context, code string, params CurrencyListRulesParams, opts ...option.RequestOption) (res *RuleArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/rules", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all transactions with this currency.
func (r *CurrencyService) ListTransactions(ctx context.Context, code string, params CurrencyListTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/currencies/%s/transactions", url.PathEscape(code))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type CurrencyRead struct {
	ID         string                 `json:"id" api:"required"`
	Attributes CurrencyReadAttributes `json:"attributes" api:"required"`
	// Immutable value
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrencyRead) RawJSON() string { return r.JSON.raw }
func (r *CurrencyRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyReadAttributes struct {
	Code      string    `json:"code" api:"required"`
	Name      string    `json:"name" api:"required"`
	Symbol    string    `json:"symbol" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Supports 0-16 decimals.
	DecimalPlaces int64 `json:"decimal_places"`
	// Defaults to true
	Enabled bool `json:"enabled"`
	// Is the primary currency?
	Primary   bool      `json:"primary"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code          respjson.Field
		Name          respjson.Field
		Symbol        respjson.Field
		CreatedAt     respjson.Field
		DecimalPlaces respjson.Field
		Enabled       respjson.Field
		Primary       respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrencyReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *CurrencyReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencySingle struct {
	Data CurrencyRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrencySingle) RawJSON() string { return r.JSON.raw }
func (r *CurrencySingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyListResponse struct {
	Data  []CurrencyRead `json:"data" api:"required"`
	Links PageLink       `json:"links" api:"required"`
	Meta  Meta           `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Links       respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrencyListResponse) RawJSON() string { return r.JSON.raw }
func (r *CurrencyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyNewParams struct {
	Code   string `json:"code" api:"required"`
	Name   string `json:"name" api:"required"`
	Symbol string `json:"symbol" api:"required"`
	// Supports 0-16 decimals.
	DecimalPlaces param.Opt[int64] `json:"decimal_places,omitzero"`
	// Defaults to true
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Make this currency the primary currency for the current administration. You can
	// set this value to FALSE, in which case nothing will change to the primary
	// currency. If you set it to TRUE, the current primary currency will no longer be
	// the primary currency.
	Primary  param.Opt[bool]   `json:"primary,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r CurrencyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CurrencyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CurrencyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type CurrencyUpdateParams struct {
	// The currency code
	Code param.Opt[string] `json:"code,omitzero"`
	// How many decimals to use when displaying this currency. Between 0 and 16.
	DecimalPlaces param.Opt[int64] `json:"decimal_places,omitzero"`
	// If the currency is enabled
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The currency name
	Name param.Opt[string] `json:"name,omitzero"`
	// The currency symbol
	Symbol   param.Opt[string] `json:"symbol,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// If the currency must be the primary for the user. You can only submit TRUE.
	// Submitting FALSE will not drop this currency as the primary currency, because
	// then the system would be without one.
	//
	// Any of true.
	Primary bool `json:"primary,omitzero"`
	paramObj
}

func (r CurrencyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CurrencyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CurrencyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyListParams]'s query parameters as `url.Values`.
func (r CurrencyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CurrencyDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type CurrencyDisableParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type CurrencyEnableParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type CurrencyListAccountsParams struct {
	// A date formatted YYYY-MM-DD. When added to the request, Firefly III will show
	// the account's balance on that day.
	Date param.Opt[time.Time] `query:"date,omitzero" format:"date" json:"-"`
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

// URLQuery serializes [CurrencyListAccountsParams]'s query parameters as
// `url.Values`.
func (r CurrencyListAccountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CurrencyListAvailableBudgetsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyListAvailableBudgetsParams]'s query parameters as
// `url.Values`.
func (r CurrencyListAvailableBudgetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CurrencyListBillsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyListBillsParams]'s query parameters as
// `url.Values`.
func (r CurrencyListBillsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CurrencyListBudgetLimitsParams struct {
	// End date for the budget limit list.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Start date for the budget limit list.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyListBudgetLimitsParams]'s query parameters as
// `url.Values`.
func (r CurrencyListBudgetLimitsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CurrencyListRecurrencesParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyListRecurrencesParams]'s query parameters as
// `url.Values`.
func (r CurrencyListRecurrencesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CurrencyListRulesParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyListRulesParams]'s query parameters as
// `url.Values`.
func (r CurrencyListRulesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CurrencyListTransactionsParams struct {
	// A date formatted YYYY-MM-DD, to limit the list of transactions.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD, to limit the list of transactions.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "all", "withdrawal", "withdrawals", "expense", "deposit", "deposits",
	// "income", "transfer", "transfers", "opening_balance", "reconciliation",
	// "special", "specials", "default".
	Type TransactionTypeFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CurrencyListTransactionsParams]'s query parameters as
// `url.Values`.
func (r CurrencyListTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
