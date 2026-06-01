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

// Endpoints to control and manage all of the user&#039;s piggy banks and related
// objects and information.
//
// PiggyBankService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPiggyBankService] method instead.
type PiggyBankService struct {
	options []option.RequestOption
}

// NewPiggyBankService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPiggyBankService(opts ...option.RequestOption) (r PiggyBankService) {
	r = PiggyBankService{}
	r.options = opts
	return
}

// Creates a new piggy bank. The data required can be submitted as a JSON body or
// as a list of parameters.
func (r *PiggyBankService) New(ctx context.Context, params PiggyBankNewParams, opts ...option.RequestOption) (res *PiggyBankSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/piggy-banks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single piggy bank.
func (r *PiggyBankService) Get(ctx context.Context, id string, query PiggyBankGetParams, opts ...option.RequestOption) (res *PiggyBankSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/piggy-banks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update existing piggy bank.
func (r *PiggyBankService) Update(ctx context.Context, id string, params PiggyBankUpdateParams, opts ...option.RequestOption) (res *PiggyBankSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/piggy-banks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all piggy banks.
func (r *PiggyBankService) List(ctx context.Context, params PiggyBankListParams, opts ...option.RequestOption) (res *PiggyBankArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/piggy-banks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a piggy bank.
func (r *PiggyBankService) Delete(ctx context.Context, id string, body PiggyBankDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/piggy-banks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Lists all attachments.
func (r *PiggyBankService) ListAttachments(ctx context.Context, id string, params PiggyBankListAttachmentsParams, opts ...option.RequestOption) (res *AttachmentArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/piggy-banks/%s/attachments", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all events linked to a piggy bank (adding and removing money).
func (r *PiggyBankService) ListEvents(ctx context.Context, id string, params PiggyBankListEventsParams, opts ...option.RequestOption) (res *PiggyBankEventArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/piggy-banks/%s/events", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type PiggyBankEventArray struct {
	Data  []PiggyBankEventArrayData `json:"data" api:"required"`
	Links PageLink                  `json:"links" api:"required"`
	Meta  Meta                      `json:"meta" api:"required"`
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
func (r PiggyBankEventArray) RawJSON() string { return r.JSON.raw }
func (r *PiggyBankEventArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankEventArrayData struct {
	ID         string                            `json:"id" api:"required"`
	Attributes PiggyBankEventArrayDataAttributes `json:"attributes" api:"required"`
	Links      ObjectLink                        `json:"links" api:"required"`
	// Immutable value
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		Links       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PiggyBankEventArrayData) RawJSON() string { return r.JSON.raw }
func (r *PiggyBankEventArrayData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankEventArrayDataAttributes struct {
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
	// Indicates whether the object has a currency setting. If false, the object uses
	// the administration's primary currency.
	ObjectHasCurrencySetting bool   `json:"object_has_currency_setting"`
	PcAmount                 string `json:"pc_amount"`
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
	// The transaction group associated with the event.
	TransactionGroupID string `json:"transaction_group_id" api:"nullable"`
	// The journal associated with the event.
	TransactionJournalID string    `json:"transaction_journal_id" api:"nullable"`
	UpdatedAt            time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount                       respjson.Field
		CreatedAt                    respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		ObjectHasCurrencySetting     respjson.Field
		PcAmount                     respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		TransactionGroupID           respjson.Field
		TransactionJournalID         respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PiggyBankEventArrayDataAttributes) RawJSON() string { return r.JSON.raw }
func (r *PiggyBankEventArrayDataAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankRead struct {
	ID         string                  `json:"id" api:"required"`
	Attributes PiggyBankReadAttributes `json:"attributes" api:"required"`
	Links      ObjectLink              `json:"links" api:"required"`
	// Immutable value
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		Links       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PiggyBankRead) RawJSON() string { return r.JSON.raw }
func (r *PiggyBankRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankReadAttributes struct {
	AccountID    any                              `json:"account_id" api:"required"`
	Name         string                           `json:"name" api:"required"`
	TargetAmount string                           `json:"target_amount" api:"required"`
	Accounts     []PiggyBankReadAttributesAccount `json:"accounts"`
	Active       bool                             `json:"active"`
	CreatedAt    time.Time                        `json:"created_at" format:"date-time"`
	// The currency code of the currency associated with this object.
	CurrencyCode          string `json:"currency_code"`
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	// The currency ID of the currency associated with this object.
	CurrencyID string `json:"currency_id"`
	// The currency name of the currency associated with this object.
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	CurrentAmount  string `json:"current_amount"`
	LeftToSave     string `json:"left_to_save" api:"nullable"`
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
	// The current amount in the primary currency of the administration.
	PcCurrentAmount string `json:"pc_current_amount"`
	PcLeftToSave    string `json:"pc_left_to_save" api:"nullable"`
	PcSavePerMonth  string `json:"pc_save_per_month" api:"nullable"`
	// The target amount in the primary currency of the administration.
	PcTargetAmount string `json:"pc_target_amount" api:"nullable"`
	// The percentage of the target amount that has been saved, if a target amount is
	// set.
	Percentage int64 `json:"percentage" api:"nullable"`
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
	SavePerMonth          string `json:"save_per_month" api:"nullable"`
	// The date you started with this piggy bank.
	StartDate time.Time `json:"start_date" format:"date-time"`
	// The date you intend to finish saving money.
	TargetDate time.Time `json:"target_date" api:"nullable" format:"date-time"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID                    respjson.Field
		Name                         respjson.Field
		TargetAmount                 respjson.Field
		Accounts                     respjson.Field
		Active                       respjson.Field
		CreatedAt                    respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		CurrentAmount                respjson.Field
		LeftToSave                   respjson.Field
		Notes                        respjson.Field
		ObjectGroupID                respjson.Field
		ObjectGroupOrder             respjson.Field
		ObjectGroupTitle             respjson.Field
		ObjectHasCurrencySetting     respjson.Field
		Order                        respjson.Field
		PcCurrentAmount              respjson.Field
		PcLeftToSave                 respjson.Field
		PcSavePerMonth               respjson.Field
		PcTargetAmount               respjson.Field
		Percentage                   respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		SavePerMonth                 respjson.Field
		StartDate                    respjson.Field
		TargetDate                   respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PiggyBankReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *PiggyBankReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankReadAttributesAccount struct {
	// The ID of the account.
	AccountID     string `json:"account_id"`
	CurrentAmount string `json:"current_amount"`
	Name          string `json:"name"`
	// If convertToPrimary is on, this will show the amount in the primary currency.
	PcCurrentAmount string `json:"pc_current_amount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID       respjson.Field
		CurrentAmount   respjson.Field
		Name            respjson.Field
		PcCurrentAmount respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PiggyBankReadAttributesAccount) RawJSON() string { return r.JSON.raw }
func (r *PiggyBankReadAttributesAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankSingle struct {
	Data PiggyBankRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PiggyBankSingle) RawJSON() string { return r.JSON.raw }
func (r *PiggyBankSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankNewParams struct {
	TargetAmount param.Opt[string] `json:"target_amount,omitzero" api:"required"`
	AccountID    any               `json:"account_id,omitzero" api:"required"`
	Name         string            `json:"name" api:"required"`
	// The date you started with this piggy bank.
	StartDate time.Time         `json:"start_date" api:"required" format:"date"`
	Notes     param.Opt[string] `json:"notes,omitzero"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupID param.Opt[string] `json:"object_group_id,omitzero"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle param.Opt[string] `json:"object_group_title,omitzero"`
	// The date you intend to finish saving money.
	TargetDate    param.Opt[time.Time]        `json:"target_date,omitzero" format:"date"`
	CurrentAmount param.Opt[string]           `json:"current_amount,omitzero"`
	Order         param.Opt[int64]            `json:"order,omitzero"`
	XTraceID      param.Opt[string]           `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	Accounts      []PiggyBankNewParamsAccount `json:"accounts,omitzero"`
	paramObj
}

func (r PiggyBankNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PiggyBankNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PiggyBankNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type PiggyBankNewParamsAccount struct {
	// The ID of the account.
	ID param.Opt[string] `json:"id,omitzero" api:"required"`
	// The name of the account.
	Name param.Opt[string] `json:"name,omitzero"`
	// The amount saved currently.
	CurrentAmount param.Opt[string] `json:"current_amount,omitzero"`
	paramObj
}

func (r PiggyBankNewParamsAccount) MarshalJSON() (data []byte, err error) {
	type shadow PiggyBankNewParamsAccount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PiggyBankNewParamsAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PiggyBankUpdateParams struct {
	Notes param.Opt[string] `json:"notes,omitzero"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupID param.Opt[string] `json:"object_group_id,omitzero"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle param.Opt[string] `json:"object_group_title,omitzero"`
	TargetAmount     param.Opt[string] `json:"target_amount,omitzero"`
	// The date you intend to finish saving money.
	TargetDate param.Opt[time.Time] `json:"target_date,omitzero" format:"date"`
	Name       param.Opt[string]    `json:"name,omitzero"`
	Order      param.Opt[int64]     `json:"order,omitzero"`
	// The date you started with this piggy bank.
	StartDate param.Opt[time.Time]           `json:"start_date,omitzero" format:"date"`
	XTraceID  param.Opt[string]              `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	Accounts  []PiggyBankUpdateParamsAccount `json:"accounts,omitzero"`
	paramObj
}

func (r PiggyBankUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PiggyBankUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PiggyBankUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type PiggyBankUpdateParamsAccount struct {
	ID any `json:"id,omitzero" api:"required"`
	// The ID of the account.
	AccountID param.Opt[string] `json:"account_id,omitzero"`
	// The amount saved currently.
	CurrentAmount param.Opt[string] `json:"current_amount,omitzero"`
	// The name of the account.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r PiggyBankUpdateParamsAccount) MarshalJSON() (data []byte, err error) {
	type shadow PiggyBankUpdateParamsAccount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PiggyBankUpdateParamsAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PiggyBankListParams]'s query parameters as `url.Values`.
func (r PiggyBankListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PiggyBankDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PiggyBankListAttachmentsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PiggyBankListAttachmentsParams]'s query parameters as
// `url.Values`.
func (r PiggyBankListAttachmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PiggyBankListEventsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PiggyBankListEventsParams]'s query parameters as
// `url.Values`.
func (r PiggyBankListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
