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

// Endpoints to manage a user&#039;s budgets and get info on the related objects,
// like limits.
//
// BudgetLimitService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBudgetLimitService] method instead.
type BudgetLimitService struct {
	options []option.RequestOption
}

// NewBudgetLimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBudgetLimitService(opts ...option.RequestOption) (r BudgetLimitService) {
	r = BudgetLimitService{}
	r.options = opts
	return
}

// Store a new budget limit under this budget.
func (r *BudgetLimitService) New(ctx context.Context, id string, params BudgetLimitNewParams, opts ...option.RequestOption) (res *BudgetLimitSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/budgets/%s/limits", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get single budget limit.
func (r *BudgetLimitService) Get(ctx context.Context, limitID int64, params BudgetLimitGetParams, opts ...option.RequestOption) (res *BudgetLimitSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/budgets/%s/limits/%v", url.PathEscape(params.ID), limitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update existing budget limit.
func (r *BudgetLimitService) Update(ctx context.Context, limitID string, params BudgetLimitUpdateParams, opts ...option.RequestOption) (res *BudgetLimitSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if limitID == "" {
		err = errors.New("missing required limitId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/budgets/%s/limits/%s", url.PathEscape(params.ID), url.PathEscape(limitID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Delete a budget limit.
func (r *BudgetLimitService) Delete(ctx context.Context, limitID string, params BudgetLimitDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if limitID == "" {
		err = errors.New("missing required limitId parameter")
		return err
	}
	path := fmt.Sprintf("v1/budgets/%s/limits/%s", url.PathEscape(params.ID), url.PathEscape(limitID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get all budget limits for this budget and the money spent, and money left. You
// can limit the list by submitting a date range as well. The "spent" array for
// each budget limit is NOT influenced by the start and end date of your query, but
// by the start and end date of the budget limit itself.
func (r *BudgetLimitService) List0(ctx context.Context, id string, params BudgetLimitList0Params, opts ...option.RequestOption) (res *BudgetLimitArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/budgets/%s/limits", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get all budget limits for for this date range.
func (r *BudgetLimitService) List1(ctx context.Context, params BudgetLimitList1Params, opts ...option.RequestOption) (res *BudgetLimitArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/budget-limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all the transactions within one budget limit. The start and end date are
// dictated by the budget limit.
func (r *BudgetLimitService) ListTransactions(ctx context.Context, limitID string, params BudgetLimitListTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if limitID == "" {
		err = errors.New("missing required limitId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/budgets/%s/limits/%s/transactions", url.PathEscape(params.ID), url.PathEscape(limitID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type BudgetLimitArray struct {
	Data []BudgetLimitRead `json:"data" api:"required"`
	Meta Meta              `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BudgetLimitArray) RawJSON() string { return r.JSON.raw }
func (r *BudgetLimitArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetLimitRead struct {
	ID         string                    `json:"id" api:"required"`
	Attributes BudgetLimitReadAttributes `json:"attributes" api:"required"`
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
func (r BudgetLimitRead) RawJSON() string { return r.JSON.raw }
func (r *BudgetLimitRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetLimitReadAttributes struct {
	Amount string `json:"amount"`
	// The budget ID of the associated budget.
	BudgetID  string    `json:"budget_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The currency code of the currency associated with this object.
	CurrencyCode          string `json:"currency_code"`
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	// The currency ID of the currency associated with this object.
	CurrencyID string `json:"currency_id"`
	// The currency name of the currency associated with this object.
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	// End date of the budget limit.
	End time.Time `json:"end" format:"date-time"`
	// Some notes for this specific budget limit.
	Notes string `json:"notes" api:"nullable"`
	// Indicates whether the object has a currency setting. If false, the object uses
	// the administration's primary currency.
	ObjectHasCurrencySetting bool `json:"object_has_currency_setting"`
	// The amount of this budget limit in the user's primary currency, if the original
	// amount is in a different currency.
	PcAmount string `json:"pc_amount" api:"nullable"`
	// Amount(s) spent in the primary currency in the database for this budget limit.
	PcSpent []ArrayEntryWithCurrencyAndSum `json:"pc_spent"`
	// Period of the budget limit. Only used when auto-generated by auto-budget.
	Period string `json:"period" api:"nullable"`
	// The currency code of the administration's primary currency.
	PrimaryCurrencyCode string `json:"primary_currency_code"`
	// The currency decimal places of the administration's primary currency.
	PrimaryCurrencyDecimalPlaces int64 `json:"primary_currency_decimal_places"`
	// The currency ID of the administration's primary currency.
	PrimaryCurrencyID string `json:"primary_currency_id"`
	// The currency name of the administration's primary currency.
	PrimaryCurrencyName string `json:"primary_currency_name"`
	// The currency symbol of the administration's primary currency.
	PrimaryCurrencySymbol string `json:"primary_currency_symbol"`
	// Amount(s) spent in the currencies in the database for this budget limit.
	Spent []ArrayEntryWithCurrencyAndSum `json:"spent"`
	// Start date of the budget limit.
	Start     time.Time `json:"start" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount                       respjson.Field
		BudgetID                     respjson.Field
		CreatedAt                    respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		End                          respjson.Field
		Notes                        respjson.Field
		ObjectHasCurrencySetting     respjson.Field
		PcAmount                     respjson.Field
		PcSpent                      respjson.Field
		Period                       respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		Spent                        respjson.Field
		Start                        respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BudgetLimitReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *BudgetLimitReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetLimitSingle struct {
	Data BudgetLimitRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BudgetLimitSingle) RawJSON() string { return r.JSON.raw }
func (r *BudgetLimitSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetLimitNewParams struct {
	Amount string `json:"amount" api:"required"`
	// End date of the budget limit.
	End time.Time `json:"end" api:"required" format:"date"`
	// Start date of the budget limit.
	Start time.Time `json:"start" api:"required" format:"date"`
	// Some notes for this specific budget limit.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's primary
	// currency.
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's primary
	// currency.
	CurrencyID param.Opt[string] `json:"currency_id,omitzero"`
	// Whether or not to fire the webhooks that are related to this event.
	FireWebhooks param.Opt[bool]   `json:"fire_webhooks,omitzero"`
	XTraceID     param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r BudgetLimitNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BudgetLimitNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BudgetLimitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetLimitGetParams struct {
	ID       string            `path:"id" api:"required" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type BudgetLimitUpdateParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Some notes for this specific budget limit.
	Notes  param.Opt[string] `json:"notes,omitzero"`
	Amount param.Opt[string] `json:"amount,omitzero"`
	// The currency code of the currency associated with this object.
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// The currency ID of the currency associated with this object.
	CurrencyID param.Opt[string] `json:"currency_id,omitzero"`
	// The currency name of the currency associated with this object.
	CurrencyName param.Opt[string] `json:"currency_name,omitzero"`
	// End date of the budget limit.
	End param.Opt[time.Time] `json:"end,omitzero" format:"date-time"`
	// Whether or not to fire the webhooks that are related to this event.
	FireWebhooks param.Opt[bool] `json:"fire_webhooks,omitzero"`
	// Start date of the budget limit.
	Start    param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r BudgetLimitUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BudgetLimitUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BudgetLimitUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetLimitDeleteParams struct {
	ID       string            `path:"id" api:"required" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type BudgetLimitList0Params struct {
	// A date formatted YYYY-MM-DD.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BudgetLimitList0Params]'s query parameters as `url.Values`.
func (r BudgetLimitList0Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BudgetLimitList1Params struct {
	// A date formatted YYYY-MM-DD.
	End time.Time `query:"end" api:"required" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    time.Time         `query:"start" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BudgetLimitList1Params]'s query parameters as `url.Values`.
func (r BudgetLimitList1Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BudgetLimitListTransactionsParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "all", "withdrawal", "withdrawals", "expense", "deposit", "deposits",
	// "income", "transfer", "transfers", "opening_balance", "reconciliation",
	// "special", "specials", "default".
	Type TransactionTypeFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BudgetLimitListTransactionsParams]'s query parameters as
// `url.Values`.
func (r BudgetLimitListTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
