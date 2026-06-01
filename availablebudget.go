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

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apijson"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apiquery"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/respjson"
)

// Endpoints to manage the total available amount that the user has made available
// to themselves. Used in periodic budgeting.
//
// AvailableBudgetService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAvailableBudgetService] method instead.
type AvailableBudgetService struct {
	options []option.RequestOption
}

// NewAvailableBudgetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAvailableBudgetService(opts ...option.RequestOption) (r AvailableBudgetService) {
	r = AvailableBudgetService{}
	r.options = opts
	return
}

// Get a single available budget, by ID.
func (r *AvailableBudgetService) Get(ctx context.Context, id string, query AvailableBudgetGetParams, opts ...option.RequestOption) (res *AvailableBudgetGetResponse, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/available-budgets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Firefly III calculates the total amount of money budgeted in so-called
// "available budgets". This endpoint returns all of these amounts and the periods
// for which they are calculated.
func (r *AvailableBudgetService) List(ctx context.Context, params AvailableBudgetListParams, opts ...option.RequestOption) (res *AvailableBudgetArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/available-budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ArrayEntryWithCurrencyAndSum struct {
	CurrencyCode string `json:"currency_code"`
	// Number of decimals supported by the currency
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	CurrencyID            string `json:"currency_id"`
	CurrencySymbol        string `json:"currency_symbol"`
	// The amount earned, spent or transferred.
	Sum string `json:"sum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyCode          respjson.Field
		CurrencyDecimalPlaces respjson.Field
		CurrencyID            respjson.Field
		CurrencySymbol        respjson.Field
		Sum                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArrayEntryWithCurrencyAndSum) RawJSON() string { return r.JSON.raw }
func (r *ArrayEntryWithCurrencyAndSum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableBudgetArray struct {
	Data []AvailableBudgetRead `json:"data" api:"required"`
	Meta Meta                  `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailableBudgetArray) RawJSON() string { return r.JSON.raw }
func (r *AvailableBudgetArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableBudgetRead struct {
	ID         string                        `json:"id" api:"required"`
	Attributes AvailableBudgetReadAttributes `json:"attributes" api:"required"`
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
func (r AvailableBudgetRead) RawJSON() string { return r.JSON.raw }
func (r *AvailableBudgetRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableBudgetReadAttributes struct {
	// The amount of this available budget in the currency of this available budget.
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The currency code of the currency associated with this object.
	CurrencyCode          string `json:"currency_code"`
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	// The currency ID of the currency associated with this object.
	CurrencyID string `json:"currency_id"`
	// The currency name of the currency associated with this object.
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	// End date of the available budget.
	End time.Time `json:"end" format:"date-time"`
	// Indicates whether the object has a currency setting. If false, the object uses
	// the administration's primary currency.
	ObjectHasCurrencySetting bool `json:"object_has_currency_setting"`
	// The amount of this available budget in the primary currency (pc) of this
	// administration.
	PcAmount string `json:"pc_amount"`
	// The amount spent in budgets in the primary currency (pc) of this administration.
	PcSpentInBudgets []ArrayEntryWithCurrencyAndSum `json:"pc_spent_in_budgets"`
	// The amount spent outside of budgets in the primary currency (pc) of this
	// administration.
	PcSpentOutsideBudgets []ArrayEntryWithCurrencyAndSum `json:"pc_spent_outside_budgets"`
	// The currency code of the administration's primary currency.
	PrimaryCurrencyCode string `json:"primary_currency_code"`
	// The currency decimal places of the administration's primary currency.
	PrimaryCurrencyDecimalPlaces int64 `json:"primary_currency_decimal_places"`
	// The currency ID of the administration's primary currency.
	PrimaryCurrencyID string `json:"primary_currency_id"`
	// The currency name of the administration's primary currency.
	PrimaryCurrencyName string `json:"primary_currency_name"`
	// The currency symbol of the administration's primary currency.
	PrimaryCurrencySymbol string                         `json:"primary_currency_symbol"`
	SpentInBudgets        []ArrayEntryWithCurrencyAndSum `json:"spent_in_budgets"`
	SpentOutsideBudgets   []ArrayEntryWithCurrencyAndSum `json:"spent_outside_budgets"`
	// Start date of the available budget.
	Start     time.Time `json:"start" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount                       respjson.Field
		CreatedAt                    respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		End                          respjson.Field
		ObjectHasCurrencySetting     respjson.Field
		PcAmount                     respjson.Field
		PcSpentInBudgets             respjson.Field
		PcSpentOutsideBudgets        respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		SpentInBudgets               respjson.Field
		SpentOutsideBudgets          respjson.Field
		Start                        respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailableBudgetReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *AvailableBudgetReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableBudgetGetResponse struct {
	Data AvailableBudgetRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailableBudgetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AvailableBudgetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailableBudgetGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type AvailableBudgetListParams struct {
	// A date formatted YYYY-MM-DD.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AvailableBudgetListParams]'s query parameters as
// `url.Values`.
func (r AvailableBudgetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
