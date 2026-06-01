// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5

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

// Endpoints to manage a user&#039;s bills and all related objects.
//
// BillService contains methods and other services that help with interacting with
// the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillService] method instead.
type BillService struct {
	options []option.RequestOption
}

// NewBillService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBillService(opts ...option.RequestOption) (r BillService) {
	r = BillService{}
	r.options = opts
	return
}

// Creates a new bill. The data required can be submitted as a JSON body or as a
// list of parameters.
func (r *BillService) New(ctx context.Context, params BillNewParams, opts ...option.RequestOption) (res *BillSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/bills"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single bill.
func (r *BillService) Get(ctx context.Context, id string, params BillGetParams, opts ...option.RequestOption) (res *BillSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bills/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Update existing bill.
func (r *BillService) Update(ctx context.Context, id string, params BillUpdateParams, opts ...option.RequestOption) (res *BillSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bills/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// This endpoint will list all the user's bills.
func (r *BillService) List(ctx context.Context, params BillListParams, opts ...option.RequestOption) (res *BillArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/bills"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a bill. This will not delete any associated rules. Will not remove
// associated transactions. WILL remove all associated attachments.
func (r *BillService) Delete(ctx context.Context, id string, body BillDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/bills/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// This endpoint will list all attachments linked to the bill.
func (r *BillService) ListAttachments(ctx context.Context, id string, params BillListAttachmentsParams, opts ...option.RequestOption) (res *AttachmentArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bills/%s/attachments", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint will list all rules that have an action to set the bill to this
// bill.
func (r *BillService) ListRules(ctx context.Context, id string, query BillListRulesParams, opts ...option.RequestOption) (res *RuleArray, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bills/%s/rules", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint will list all transactions linked to this bill.
func (r *BillService) ListTransactions(ctx context.Context, id string, params BillListTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bills/%s/transactions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type BillArray struct {
	Data []BillRead `json:"data" api:"required"`
	Meta Meta       `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillArray) RawJSON() string { return r.JSON.raw }
func (r *BillArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillRead struct {
	ID         string             `json:"id" api:"required"`
	Attributes BillReadAttributes `json:"attributes" api:"required"`
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
func (r BillRead) RawJSON() string { return r.JSON.raw }
func (r *BillRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillReadAttributes struct {
	// If the subscription is active.
	Active bool `json:"active"`
	// The average amount that is expected for this subscription in the subscription's
	// currency.
	AmountAvg string `json:"amount_avg"`
	// The maximum amount that is expected for this subscription in the subscription's
	// currency.
	AmountMax string `json:"amount_max"`
	// The minimum amount that is expected for this subscription in the subscription's
	// currency.
	AmountMin string    `json:"amount_min"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The currency code of the currency associated with this object.
	CurrencyCode          string `json:"currency_code"`
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	// The currency ID of the currency associated with this object.
	CurrencyID string `json:"currency_id"`
	// The currency name of the currency associated with this object.
	CurrencyName   string    `json:"currency_name"`
	CurrencySymbol string    `json:"currency_symbol"`
	Date           time.Time `json:"date" format:"date-time"`
	// The date after which this subscription is no longer valid or applicable
	EndDate time.Time `json:"end_date" api:"nullable" format:"date-time"`
	// The date before which the subscription must be renewed (or cancelled)
	ExtensionDate time.Time `json:"extension_date" api:"nullable" format:"date-time"`
	// The name of the subscription.
	Name string `json:"name"`
	// When the subscription is expected to be due.
	NextExpectedMatch time.Time `json:"next_expected_match" api:"nullable" format:"date-time"`
	// Formatted (locally) when the subscription is due.
	NextExpectedMatchDiff string `json:"next_expected_match_diff" api:"nullable"`
	Notes                 string `json:"notes" api:"nullable"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupID string `json:"object_group_id" api:"nullable"`
	// The order of the group. At least 1, for the highest sorting.
	ObjectGroupOrder int64 `json:"object_group_order" api:"nullable"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle string `json:"object_group_title" api:"nullable"`
	// Indicates whether the object has a currency setting. If false, the object uses
	// the administration's primary currency.
	ObjectHasCurrencySetting bool `json:"object_has_currency_setting"`
	// Order of the subscription.
	Order int64 `json:"order"`
	// Array of past transactions when the subscription was paid.
	PaidDates []BillReadAttributesPaidDate `json:"paid_dates"`
	// Array of future dates when the bill is expected to be paid. Autogenerated.
	PayDates []time.Time `json:"pay_dates" format:"date-time"`
	// The average amount that is expected for this subscription in the
	// administration's primary currency.
	PcAmountAvg string `json:"pc_amount_avg"`
	// The maximum amount that is expected for this subscription in the
	// administration's primary currency.
	PcAmountMax string `json:"pc_amount_max"`
	// The minimum amount that is expected for this subscription in the
	// administration's primary currency.
	PcAmountMin string `json:"pc_amount_min"`
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
	// How often the bill must be paid.
	//
	// Any of "weekly", "monthly", "quarterly", "half-year", "yearly".
	RepeatFreq BillRepeatFrequency `json:"repeat_freq"`
	// How often the subscription will be skipped. 1 means a bi-monthly subscription.
	Skip      int64     `json:"skip"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active                       respjson.Field
		AmountAvg                    respjson.Field
		AmountMax                    respjson.Field
		AmountMin                    respjson.Field
		CreatedAt                    respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		Date                         respjson.Field
		EndDate                      respjson.Field
		ExtensionDate                respjson.Field
		Name                         respjson.Field
		NextExpectedMatch            respjson.Field
		NextExpectedMatchDiff        respjson.Field
		Notes                        respjson.Field
		ObjectGroupID                respjson.Field
		ObjectGroupOrder             respjson.Field
		ObjectGroupTitle             respjson.Field
		ObjectHasCurrencySetting     respjson.Field
		Order                        respjson.Field
		PaidDates                    respjson.Field
		PayDates                     respjson.Field
		PcAmountAvg                  respjson.Field
		PcAmountMax                  respjson.Field
		PcAmountMin                  respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		RepeatFreq                   respjson.Field
		Skip                         respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *BillReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillReadAttributesPaidDate struct {
	// The amount that was paid for this subscription in the subscription's currency.
	Amount string `json:"amount"`
	// The currency code of the currency associated with this object.
	CurrencyCode          string `json:"currency_code"`
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	// The currency ID of the currency associated with this object.
	CurrencyID string `json:"currency_id"`
	// The currency name of the currency associated with this object.
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	// Date the bill was paid.
	Date time.Time `json:"date" format:"date-time"`
	// The foreign amount that was paid for this subscription in the subscription's
	// currency.
	ForeignAmount string `json:"foreign_amount"`
	// The amount that was paid for this subscription in the administration's primary
	// currency.
	PcAmount string `json:"pc_amount"`
	// The foreign amount that was paid for this subscription in the administration's
	// primary currency.
	PcForeignAmount string `json:"pc_foreign_amount"`
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
	// ID of this subscription.
	SubscriptionID string `json:"subscription_id"`
	// Transaction group ID of the transaction linked to this subscription.
	TransactionGroupID string `json:"transaction_group_id"`
	// Transaction journal ID of the transaction linked to this subscription.
	TransactionJournalID string `json:"transaction_journal_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount                       respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		Date                         respjson.Field
		ForeignAmount                respjson.Field
		PcAmount                     respjson.Field
		PcForeignAmount              respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		SubscriptionID               respjson.Field
		TransactionGroupID           respjson.Field
		TransactionJournalID         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillReadAttributesPaidDate) RawJSON() string { return r.JSON.raw }
func (r *BillReadAttributesPaidDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How often the bill must be paid.
type BillRepeatFrequency string

const (
	BillRepeatFrequencyWeekly    BillRepeatFrequency = "weekly"
	BillRepeatFrequencyMonthly   BillRepeatFrequency = "monthly"
	BillRepeatFrequencyQuarterly BillRepeatFrequency = "quarterly"
	BillRepeatFrequencyHalfYear  BillRepeatFrequency = "half-year"
	BillRepeatFrequencyYearly    BillRepeatFrequency = "yearly"
)

type BillSingle struct {
	Data BillRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillSingle) RawJSON() string { return r.JSON.raw }
func (r *BillSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RuleArray struct {
	Data  []RuleRead `json:"data" api:"required"`
	Links PageLink   `json:"links" api:"required"`
	Meta  Meta       `json:"meta" api:"required"`
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
func (r RuleArray) RawJSON() string { return r.JSON.raw }
func (r *RuleArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillNewParams struct {
	AmountMax string    `json:"amount_max" api:"required"`
	AmountMin string    `json:"amount_min" api:"required"`
	Date      time.Time `json:"date" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// How often the bill must be paid.
	//
	// Any of "weekly", "monthly", "quarterly", "half-year", "yearly".
	RepeatFreq BillRepeatFrequency `json:"repeat_freq,omitzero" api:"required"`
	Notes      param.Opt[string]   `json:"notes,omitzero"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupID param.Opt[string] `json:"object_group_id,omitzero"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle param.Opt[string] `json:"object_group_title,omitzero"`
	// If the bill is active.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Use either currency_id or currency_code
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// Use either currency_id or currency_code
	CurrencyID param.Opt[string] `json:"currency_id,omitzero"`
	// The date after which this bill is no longer valid or applicable
	EndDate param.Opt[time.Time] `json:"end_date,omitzero" format:"date-time"`
	// The date before which the bill must be renewed (or cancelled)
	ExtensionDate param.Opt[time.Time] `json:"extension_date,omitzero" format:"date-time"`
	// How often the bill must be skipped. 1 means a bi-monthly bill.
	Skip     param.Opt[int64]  `json:"skip,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r BillNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BillNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BillNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillGetParams struct {
	// A date formatted YYYY-MM-DD. If it is added to the request, Firefly III will
	// calculate the appropriate payment and paid dates.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD. If it is are added to the request, Firefly III will
	// calculate the appropriate payment and paid dates.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BillGetParams]'s query parameters as `url.Values`.
func (r BillGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillUpdateParams struct {
	Name  string            `json:"name" api:"required"`
	Notes param.Opt[string] `json:"notes,omitzero"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupID param.Opt[string] `json:"object_group_id,omitzero"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle param.Opt[string] `json:"object_group_title,omitzero"`
	// If the bill is active.
	Active    param.Opt[bool]   `json:"active,omitzero"`
	AmountMax param.Opt[string] `json:"amount_max,omitzero"`
	AmountMin param.Opt[string] `json:"amount_min,omitzero"`
	// Use either currency_id or currency_code
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// Use either currency_id or currency_code
	CurrencyID param.Opt[string]    `json:"currency_id,omitzero"`
	Date       param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// The date after which this bill is no longer valid or applicable
	EndDate param.Opt[time.Time] `json:"end_date,omitzero" format:"date-time"`
	// The date before which the bill must be renewed (or cancelled)
	ExtensionDate param.Opt[time.Time] `json:"extension_date,omitzero" format:"date-time"`
	// How often the bill must be skipped. 1 means a bi-monthly bill.
	Skip     param.Opt[int64]  `json:"skip,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// How often the bill must be paid.
	//
	// Any of "weekly", "monthly", "quarterly", "half-year", "yearly".
	RepeatFreq BillRepeatFrequency `json:"repeat_freq,omitzero"`
	paramObj
}

func (r BillUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BillUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BillUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillListParams struct {
	// A date formatted YYYY-MM-DD. If it is added to the request, Firefly III will
	// calculate the appropriate payment and paid dates.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD. If it is are added to the request, Firefly III will
	// calculate the appropriate payment and paid dates.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BillListParams]'s query parameters as `url.Values`.
func (r BillListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type BillListAttachmentsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BillListAttachmentsParams]'s query parameters as
// `url.Values`.
func (r BillListAttachmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillListRulesParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type BillListTransactionsParams struct {
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

// URLQuery serializes [BillListTransactionsParams]'s query parameters as
// `url.Values`.
func (r BillListTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
