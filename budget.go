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

// Endpoints to manage a user&#039;s budgets and get info on the related objects,
// like limits.
//
// BudgetService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBudgetService] method instead.
type BudgetService struct {
	options []option.RequestOption
	// Endpoints to manage a user&#039;s budgets and get info on the related objects,
	// like limits.
	Limits BudgetLimitService
}

// NewBudgetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBudgetService(opts ...option.RequestOption) (r BudgetService) {
	r = BudgetService{}
	r.options = opts
	r.Limits = NewBudgetLimitService(opts...)
	return
}

// Creates a new budget. The data required can be submitted as a JSON body or as a
// list of parameters.
func (r *BudgetService) New(ctx context.Context, params BudgetNewParams, opts ...option.RequestOption) (res *BudgetSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single budget. If the start date and end date are submitted as well, the
// "spent" array will be updated accordingly.
func (r *BudgetService) Get(ctx context.Context, id string, params BudgetGetParams, opts ...option.RequestOption) (res *BudgetSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/budgets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Update existing budget. This endpoint cannot be used to set budget amount
// limits.
func (r *BudgetService) Update(ctx context.Context, id string, params BudgetUpdateParams, opts ...option.RequestOption) (res *BudgetSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/budgets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all the budgets the user has made. If the start date and end date are
// submitted as well, the "spent" array will be updated accordingly.
func (r *BudgetService) List(ctx context.Context, params BudgetListParams, opts ...option.RequestOption) (res *BudgetListResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a budget. Transactions will not be deleted.
func (r *BudgetService) Delete(ctx context.Context, id string, body BudgetDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/budgets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Lists all attachments.
func (r *BudgetService) ListAttachments(ctx context.Context, id string, params BudgetListAttachmentsParams, opts ...option.RequestOption) (res *AttachmentArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/budgets/%s/attachments", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get all transactions linked to a budget, possibly limited by start and end
func (r *BudgetService) ListTransactions(ctx context.Context, id string, params BudgetListTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/budgets/%s/transactions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get all transactions NOT linked to a budget, possibly limited by start and end
func (r *BudgetService) ListTransactionsWithoutBudget(ctx context.Context, params BudgetListTransactionsWithoutBudgetParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/budgets/transactions-without-budget"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Period for the auto budget
type AutoBudgetPeriod string

const (
	AutoBudgetPeriodDaily     AutoBudgetPeriod = "daily"
	AutoBudgetPeriodWeekly    AutoBudgetPeriod = "weekly"
	AutoBudgetPeriodMonthly   AutoBudgetPeriod = "monthly"
	AutoBudgetPeriodQuarterly AutoBudgetPeriod = "quarterly"
	AutoBudgetPeriodHalfYear  AutoBudgetPeriod = "half-year"
	AutoBudgetPeriodYearly    AutoBudgetPeriod = "yearly"
)

// The type of auto-budget that Firefly III must create.
type AutoBudgetType string

const (
	AutoBudgetTypeReset    AutoBudgetType = "reset"
	AutoBudgetTypeRollover AutoBudgetType = "rollover"
	AutoBudgetTypeNone     AutoBudgetType = "none"
)

type BudgetRead struct {
	ID         string               `json:"id" api:"required"`
	Attributes BudgetReadAttributes `json:"attributes" api:"required"`
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
func (r BudgetRead) RawJSON() string { return r.JSON.raw }
func (r *BudgetRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetReadAttributes struct {
	Name   string `json:"name" api:"required"`
	Active bool   `json:"active"`
	// The amount for the auto-budget, if set.
	AutoBudgetAmount string `json:"auto_budget_amount" api:"nullable"`
	// Period for the auto budget
	//
	// Any of "daily", "weekly", "monthly", "quarterly", "half-year", "yearly".
	AutoBudgetPeriod AutoBudgetPeriod `json:"auto_budget_period" api:"nullable"`
	// The type of auto-budget that Firefly III must create.
	//
	// Any of "reset", "rollover", "none".
	AutoBudgetType AutoBudgetType `json:"auto_budget_type" api:"nullable"`
	CreatedAt      time.Time      `json:"created_at" format:"date-time"`
	// The currency code of the currency associated with this object.
	CurrencyCode          string `json:"currency_code"`
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	// The currency ID of the currency associated with this object.
	CurrencyID string `json:"currency_id"`
	// The currency name of the currency associated with this object.
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	Notes          string `json:"notes" api:"nullable"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupID string `json:"object_group_id" api:"nullable"`
	// The order of the group. At least 1, for the highest sorting.
	ObjectGroupOrder int64 `json:"object_group_order" api:"nullable"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle string `json:"object_group_title" api:"nullable"`
	// Indicates whether the object has a currency setting. If false, the object uses
	// the administration's primary currency.
	ObjectHasCurrencySetting bool  `json:"object_has_currency_setting"`
	Order                    int64 `json:"order"`
	// The amount for the auto-budget, if set in the primary currency of the
	// administration.
	PcAutoBudgetAmount string `json:"pc_auto_budget_amount" api:"nullable"`
	// Information on how much was spent in this budget. Is only filled in when the
	// start and end date are submitted. It is converted to the primary currency of the
	// administration.
	PcSpent []ArrayEntryWithCurrencyAndSum `json:"pc_spent"`
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
	// Information on how much was spent in this budget. Is only filled in when the
	// start and end date are submitted.
	Spent     []ArrayEntryWithCurrencyAndSum `json:"spent"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                         respjson.Field
		Active                       respjson.Field
		AutoBudgetAmount             respjson.Field
		AutoBudgetPeriod             respjson.Field
		AutoBudgetType               respjson.Field
		CreatedAt                    respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		Notes                        respjson.Field
		ObjectGroupID                respjson.Field
		ObjectGroupOrder             respjson.Field
		ObjectGroupTitle             respjson.Field
		ObjectHasCurrencySetting     respjson.Field
		Order                        respjson.Field
		PcAutoBudgetAmount           respjson.Field
		PcSpent                      respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		Spent                        respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BudgetReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *BudgetReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetSingle struct {
	Data BudgetRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BudgetSingle) RawJSON() string { return r.JSON.raw }
func (r *BudgetSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetListResponse struct {
	Data []BudgetRead `json:"data" api:"required"`
	Meta Meta         `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BudgetListResponse) RawJSON() string { return r.JSON.raw }
func (r *BudgetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetNewParams struct {
	Name             string            `json:"name" api:"required"`
	AutoBudgetAmount param.Opt[string] `json:"auto_budget_amount,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's financial
	// administration's currency.
	AutoBudgetCurrencyCode param.Opt[string] `json:"auto_budget_currency_code,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's financial
	// administration's currency.
	AutoBudgetCurrencyID param.Opt[string] `json:"auto_budget_currency_id,omitzero"`
	Notes                param.Opt[string] `json:"notes,omitzero"`
	Active               param.Opt[bool]   `json:"active,omitzero"`
	// Whether or not to fire the webhooks that are related to this event.
	FireWebhooks param.Opt[bool]   `json:"fire_webhooks,omitzero"`
	XTraceID     param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Period for the auto budget
	//
	// Any of "daily", "weekly", "monthly", "quarterly", "half-year", "yearly".
	AutoBudgetPeriod AutoBudgetPeriod `json:"auto_budget_period,omitzero"`
	// The type of auto-budget that Firefly III must create.
	//
	// Any of "reset", "rollover", "none".
	AutoBudgetType AutoBudgetType `json:"auto_budget_type,omitzero"`
	paramObj
}

func (r BudgetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BudgetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BudgetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetGetParams struct {
	// A date formatted YYYY-MM-DD, to get info on how much the user has spent.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD, to get info on how much the user has spent.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BudgetGetParams]'s query parameters as `url.Values`.
func (r BudgetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BudgetUpdateParams struct {
	Name             string            `json:"name" api:"required"`
	AutoBudgetAmount param.Opt[string] `json:"auto_budget_amount,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's financial
	// administration's currency.
	AutoBudgetCurrencyCode param.Opt[string] `json:"auto_budget_currency_code,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's financial
	// administration's currency.
	AutoBudgetCurrencyID param.Opt[string] `json:"auto_budget_currency_id,omitzero"`
	Notes                param.Opt[string] `json:"notes,omitzero"`
	Active               param.Opt[bool]   `json:"active,omitzero"`
	// Whether or not to fire the webhooks that are related to this event.
	FireWebhooks param.Opt[bool]   `json:"fire_webhooks,omitzero"`
	Order        param.Opt[int64]  `json:"order,omitzero"`
	XTraceID     param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Period for the auto budget
	//
	// Any of "daily", "weekly", "monthly", "quarterly", "half-year", "yearly".
	AutoBudgetPeriod AutoBudgetPeriod `json:"auto_budget_period,omitzero"`
	// The type of auto-budget that Firefly III must create.
	//
	// Any of "reset", "rollover", "none".
	AutoBudgetType AutoBudgetType `json:"auto_budget_type,omitzero"`
	paramObj
}

func (r BudgetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BudgetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BudgetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetListParams struct {
	// A date formatted YYYY-MM-DD, to get info on how much the user has spent. You
	// must submit both start and end.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD, to get info on how much the user has spent. You
	// must submit both start and end.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BudgetListParams]'s query parameters as `url.Values`.
func (r BudgetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BudgetDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type BudgetListAttachmentsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BudgetListAttachmentsParams]'s query parameters as
// `url.Values`.
func (r BudgetListAttachmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BudgetListTransactionsParams struct {
	// A date formatted YYYY-MM-DD.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "all", "withdrawal", "withdrawals", "expense", "deposit", "deposits",
	// "income", "transfer", "transfers", "opening_balance", "reconciliation",
	// "special", "specials", "default".
	Type TransactionTypeFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BudgetListTransactionsParams]'s query parameters as
// `url.Values`.
func (r BudgetListTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BudgetListTransactionsWithoutBudgetParams struct {
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

// URLQuery serializes [BudgetListTransactionsWithoutBudgetParams]'s query
// parameters as `url.Values`.
func (r BudgetListTransactionsWithoutBudgetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
