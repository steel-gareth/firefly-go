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

// Use these endpoints to manage the user&#039;s recurring transactions, trigger
// the creation of transactions and manage the settings.
//
// RecurrenceService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecurrenceService] method instead.
type RecurrenceService struct {
	options []option.RequestOption
}

// NewRecurrenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecurrenceService(opts ...option.RequestOption) (r RecurrenceService) {
	r = RecurrenceService{}
	r.options = opts
	return
}

// Creates a new recurring transaction. The data required can be submitted as a
// JSON body or as a list of parameters.
func (r *RecurrenceService) New(ctx context.Context, params RecurrenceNewParams, opts ...option.RequestOption) (res *RecurrenceSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/recurrences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single recurring transaction.
func (r *RecurrenceService) Get(ctx context.Context, id string, query RecurrenceGetParams, opts ...option.RequestOption) (res *RecurrenceSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/recurrences/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update existing recurring transaction.
func (r *RecurrenceService) Update(ctx context.Context, id string, params RecurrenceUpdateParams, opts ...option.RequestOption) (res *RecurrenceSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/recurrences/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all recurring transactions.
func (r *RecurrenceService) List(ctx context.Context, params RecurrenceListParams, opts ...option.RequestOption) (res *RecurrenceArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/recurrences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a recurring transaction. Transactions created by the recurring
// transaction will not be deleted.
func (r *RecurrenceService) Delete(ctx context.Context, id string, body RecurrenceDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/recurrences/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// List all transactions created by a recurring transaction, optionally limited to
// the date ranges specified.
func (r *RecurrenceService) ListTransactions(ctx context.Context, id string, params RecurrenceListTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/recurrences/%s/transactions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Trigger the creation of a transaction for a specific recurring transaction. All
// recurrences have a set of future occurrences. For those moments, you can trigger
// the creation of the transaction. That means the transaction will be created NOW,
// instead of on the indicated date. The transaction will be dated to _today_.
//
// So, if you recurring transaction that occurs every Monday, you can trigger the
// creation of a transaction for Monday in two weeks, today. On that Monday two
// weeks from now, no transaction will be created. Instead, the transaction is
// created right now, and dated _today_.
func (r *RecurrenceService) TriggerTransaction(ctx context.Context, id string, params RecurrenceTriggerTransactionParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/recurrences/%s/trigger", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type AccountTypeProperty string

const (
	AccountTypePropertyDefaultAccount        AccountTypeProperty = "Default account"
	AccountTypePropertyCashAccount           AccountTypeProperty = "Cash account"
	AccountTypePropertyAssetAccount          AccountTypeProperty = "Asset account"
	AccountTypePropertyExpenseAccount        AccountTypeProperty = "Expense account"
	AccountTypePropertyRevenueAccount        AccountTypeProperty = "Revenue account"
	AccountTypePropertyInitialBalanceAccount AccountTypeProperty = "Initial balance account"
	AccountTypePropertyBeneficiaryAccount    AccountTypeProperty = "Beneficiary account"
	AccountTypePropertyImportAccount         AccountTypeProperty = "Import account"
	AccountTypePropertyReconciliationAccount AccountTypeProperty = "Reconciliation account"
	AccountTypePropertyLoan                  AccountTypeProperty = "Loan"
	AccountTypePropertyDebt                  AccountTypeProperty = "Debt"
	AccountTypePropertyMortgage              AccountTypeProperty = "Mortgage"
)

type RecurrenceArray struct {
	Data  []RecurrenceRead `json:"data" api:"required"`
	Links PageLink         `json:"links" api:"required"`
	Meta  Meta             `json:"meta" api:"required"`
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
func (r RecurrenceArray) RawJSON() string { return r.JSON.raw }
func (r *RecurrenceArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurrenceRead struct {
	ID         string                   `json:"id" api:"required"`
	Attributes RecurrenceReadAttributes `json:"attributes" api:"required"`
	Links      ObjectLink               `json:"links" api:"required"`
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
func (r RecurrenceRead) RawJSON() string { return r.JSON.raw }
func (r *RecurrenceRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurrenceReadAttributes struct {
	// If the recurrence is even active.
	Active bool `json:"active"`
	// Whether or not to fire the rules after the creation of a transaction.
	ApplyRules bool      `json:"apply_rules"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// Not to be confused with the description of the actual transaction(s) being
	// created.
	Description string `json:"description"`
	// First time the recurring transaction will fire. Must be after today.
	FirstDate time.Time `json:"first_date" format:"date"`
	// Last time the recurring transaction has fired.
	LatestDate time.Time `json:"latest_date" api:"nullable" format:"date"`
	Notes      string    `json:"notes" api:"nullable"`
	// Max number of created transactions. Use either this field or repeat_until.
	NrOfRepetitions int64 `json:"nr_of_repetitions" api:"nullable"`
	// Date until the recurring transaction can fire. Use either this field or
	// repetitions.
	RepeatUntil  time.Time                             `json:"repeat_until" api:"nullable" format:"date"`
	Repetitions  []RecurrenceReadAttributesRepetition  `json:"repetitions"`
	Title        string                                `json:"title"`
	Transactions []RecurrenceReadAttributesTransaction `json:"transactions"`
	// Any of "withdrawal", "transfer", "deposit".
	Type      RecurrenceTransactionType `json:"type"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active          respjson.Field
		ApplyRules      respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		FirstDate       respjson.Field
		LatestDate      respjson.Field
		Notes           respjson.Field
		NrOfRepetitions respjson.Field
		RepeatUntil     respjson.Field
		Repetitions     respjson.Field
		Title           respjson.Field
		Transactions    respjson.Field
		Type            respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecurrenceReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *RecurrenceReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurrenceReadAttributesRepetition struct {
	// Information that defined the type of repetition.
	//
	//   - For 'daily', this is empty.
	//   - For 'weekly', it is day of the week between 1 and 7 (Monday - Sunday).
	//   - For 'ndom', it is '1,2' or '4,5' or something else, where the first number is
	//     the week in the month, and the second number is the day in the week (between 1
	//     and 7). '2,3' means: the 2nd Wednesday of the month
	//   - For 'monthly' it is the day of the month (1 - 31)
	//   - For yearly, it is a full date, ie '2026-04-01'. The year you use does not
	//     matter.
	Moment string `json:"moment" api:"required"`
	// The type of the repetition. ndom means: the n-th weekday of the month, where you
	// can also specify which day of the week.
	//
	// Any of "daily", "weekly", "ndom", "monthly", "yearly".
	Type      RecurrenceRepetitionType `json:"type" api:"required"`
	ID        string                   `json:"id"`
	CreatedAt time.Time                `json:"created_at" format:"date-time"`
	// Auto-generated repetition description.
	Description string `json:"description"`
	// Array of future dates when the repetition will apply to. Auto generated.
	Occurrences []time.Time `json:"occurrences" format:"date-time"`
	// How many occurrences to skip. 0 means skip nothing. 1 means every other.
	Skip      int64     `json:"skip"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// How to respond when the recurring transaction falls in the weekend. Possible
	// values:
	//
	// 1. Do nothing, just create it
	// 2. Create no transaction.
	// 3. Skip to the previous Friday.
	// 4. Skip to the next Monday.
	Weekend int64 `json:"weekend"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Moment      respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Occurrences respjson.Field
		Skip        respjson.Field
		UpdatedAt   respjson.Field
		Weekend     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecurrenceReadAttributesRepetition) RawJSON() string { return r.JSON.raw }
func (r *RecurrenceReadAttributesRepetition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurrenceReadAttributesTransaction struct {
	// Amount of the transaction.
	Amount      string `json:"amount" api:"required"`
	Description string `json:"description" api:"required"`
	ID          string `json:"id"`
	// The budget ID for this transaction.
	BudgetID string `json:"budget_id"`
	// The name of the budget to be used. If the budget name is unknown, the ID will be
	// used or the value will be ignored.
	BudgetName string `json:"budget_name" api:"nullable"`
	// Category ID for this transaction.
	CategoryID string `json:"category_id"`
	// Category name for this transaction.
	CategoryName string `json:"category_name"`
	// The currency code of the currency associated with this object.
	CurrencyCode          string `json:"currency_code"`
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	// The currency ID of the currency associated with this object.
	CurrencyID string `json:"currency_id"`
	// The currency name of the currency associated with this object.
	CurrencyName    string `json:"currency_name"`
	CurrencySymbol  string `json:"currency_symbol"`
	DestinationIban string `json:"destination_iban" api:"nullable"`
	// ID of the destination account. Submit either this or destination_name.
	DestinationID string `json:"destination_id"`
	// Name of the destination account. Submit either this or destination_id.
	DestinationName string `json:"destination_name"`
	// Any of "Default account", "Cash account", "Asset account", "Expense account",
	// "Revenue account", "Initial balance account", "Beneficiary account", "Import
	// account", "Reconciliation account", "Loan", "Debt", "Mortgage".
	DestinationType AccountTypeProperty `json:"destination_type"`
	// Foreign amount of the transaction.
	ForeignAmount       string `json:"foreign_amount" api:"nullable"`
	ForeignCurrencyCode string `json:"foreign_currency_code" api:"nullable"`
	// Number of decimals in the currency
	ForeignCurrencyDecimalPlaces int64  `json:"foreign_currency_decimal_places" api:"nullable"`
	ForeignCurrencyID            string `json:"foreign_currency_id" api:"nullable"`
	ForeignCurrencyName          string `json:"foreign_currency_name" api:"nullable"`
	ForeignCurrencySymbol        string `json:"foreign_currency_symbol" api:"nullable"`
	// Indicates whether the object has a currency setting. If false, the object uses
	// the administration's primary currency.
	ObjectHasCurrencySetting bool `json:"object_has_currency_setting"`
	// Amount of the transaction in primary currency.
	PcAmount string `json:"pc_amount"`
	// Foreign amount of the transaction.
	PcForeignAmount string `json:"pc_foreign_amount" api:"nullable"`
	PiggyBankID     string `json:"piggy_bank_id" api:"nullable"`
	PiggyBankName   string `json:"piggy_bank_name" api:"nullable"`
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
	SourceIban            string `json:"source_iban" api:"nullable"`
	// ID of the source account. Submit either this or source_name.
	SourceID string `json:"source_id"`
	// Name of the source account. Submit either this or source_id.
	SourceName string `json:"source_name"`
	// Any of "Default account", "Cash account", "Asset account", "Expense account",
	// "Revenue account", "Initial balance account", "Beneficiary account", "Import
	// account", "Reconciliation account", "Loan", "Debt", "Mortgage".
	SourceType       AccountTypeProperty `json:"source_type"`
	SubscriptionID   string              `json:"subscription_id" api:"nullable"`
	SubscriptionName string              `json:"subscription_name" api:"nullable"`
	// Array of tags.
	Tags []string `json:"tags" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount                       respjson.Field
		Description                  respjson.Field
		ID                           respjson.Field
		BudgetID                     respjson.Field
		BudgetName                   respjson.Field
		CategoryID                   respjson.Field
		CategoryName                 respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		DestinationIban              respjson.Field
		DestinationID                respjson.Field
		DestinationName              respjson.Field
		DestinationType              respjson.Field
		ForeignAmount                respjson.Field
		ForeignCurrencyCode          respjson.Field
		ForeignCurrencyDecimalPlaces respjson.Field
		ForeignCurrencyID            respjson.Field
		ForeignCurrencyName          respjson.Field
		ForeignCurrencySymbol        respjson.Field
		ObjectHasCurrencySetting     respjson.Field
		PcAmount                     respjson.Field
		PcForeignAmount              respjson.Field
		PiggyBankID                  respjson.Field
		PiggyBankName                respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		SourceIban                   respjson.Field
		SourceID                     respjson.Field
		SourceName                   respjson.Field
		SourceType                   respjson.Field
		SubscriptionID               respjson.Field
		SubscriptionName             respjson.Field
		Tags                         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecurrenceReadAttributesTransaction) RawJSON() string { return r.JSON.raw }
func (r *RecurrenceReadAttributesTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the repetition. ndom means: the n-th weekday of the month, where you
// can also specify which day of the week.
type RecurrenceRepetitionType string

const (
	RecurrenceRepetitionTypeDaily   RecurrenceRepetitionType = "daily"
	RecurrenceRepetitionTypeWeekly  RecurrenceRepetitionType = "weekly"
	RecurrenceRepetitionTypeNdom    RecurrenceRepetitionType = "ndom"
	RecurrenceRepetitionTypeMonthly RecurrenceRepetitionType = "monthly"
	RecurrenceRepetitionTypeYearly  RecurrenceRepetitionType = "yearly"
)

type RecurrenceSingle struct {
	Data RecurrenceRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecurrenceSingle) RawJSON() string { return r.JSON.raw }
func (r *RecurrenceSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurrenceTransactionType string

const (
	RecurrenceTransactionTypeWithdrawal RecurrenceTransactionType = "withdrawal"
	RecurrenceTransactionTypeTransfer   RecurrenceTransactionType = "transfer"
	RecurrenceTransactionTypeDeposit    RecurrenceTransactionType = "deposit"
)

type RecurrenceNewParams struct {
	// Date until the recurring transaction can fire. Use either this field or
	// repetitions.
	RepeatUntil param.Opt[time.Time] `json:"repeat_until,omitzero" api:"required" format:"date"`
	// First time the recurring transaction will fire. Must be after today.
	FirstDate    time.Time                        `json:"first_date" api:"required" format:"date"`
	Repetitions  []RecurrenceNewParamsRepetition  `json:"repetitions,omitzero" api:"required"`
	Title        string                           `json:"title" api:"required"`
	Transactions []RecurrenceNewParamsTransaction `json:"transactions,omitzero" api:"required"`
	// Any of "withdrawal", "transfer", "deposit".
	Type  RecurrenceTransactionType `json:"type,omitzero" api:"required"`
	Notes param.Opt[string]         `json:"notes,omitzero"`
	// Max number of created transactions. Use either this field or repeat_until.
	NrOfRepetitions param.Opt[int64] `json:"nr_of_repetitions,omitzero"`
	// If the recurrence is even active.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Whether or not to fire the rules after the creation of a transaction.
	ApplyRules param.Opt[bool] `json:"apply_rules,omitzero"`
	// Not to be confused with the description of the actual transaction(s) being
	// created.
	Description param.Opt[string] `json:"description,omitzero"`
	XTraceID    param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r RecurrenceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RecurrenceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecurrenceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Moment, Type are required.
type RecurrenceNewParamsRepetition struct {
	// Information that defined the type of repetition.
	//
	//   - For 'daily', this is empty.
	//   - For 'weekly', it is day of the week between 1 and 7 (Monday - Sunday).
	//   - For 'ndom', it is '1,2' or '4,5' or something else, where the first number is
	//     the week in the month, and the second number is the day in the week (between 1
	//     and 7). '2,3' means: the 2nd Wednesday of the month
	//   - For 'monthly' it is the day of the month (1 - 31)
	//   - For yearly, it is a full date, ie '2026-04-01'. The year you use does not
	//     matter.
	Moment string `json:"moment" api:"required"`
	// The type of the repetition. ndom means: the n-th weekday of the month, where you
	// can also specify which day of the week.
	//
	// Any of "daily", "weekly", "ndom", "monthly", "yearly".
	Type RecurrenceRepetitionType `json:"type,omitzero" api:"required"`
	// How many occurrences to skip. 0 means skip nothing. 1 means every other.
	Skip param.Opt[int64] `json:"skip,omitzero"`
	// How to respond when the recurring transaction falls in the weekend. Possible
	// values:
	//
	// 1. Do nothing, just create it
	// 2. Create no transaction.
	// 3. Skip to the previous Friday.
	// 4. Skip to the next Monday.
	Weekend param.Opt[int64] `json:"weekend,omitzero"`
	paramObj
}

func (r RecurrenceNewParamsRepetition) MarshalJSON() (data []byte, err error) {
	type shadow RecurrenceNewParamsRepetition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecurrenceNewParamsRepetition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Description, DestinationID, SourceID are required.
type RecurrenceNewParamsTransaction struct {
	// Amount of the transaction.
	Amount      string `json:"amount" api:"required"`
	Description string `json:"description" api:"required"`
	// ID of the destination account.
	DestinationID string `json:"destination_id" api:"required"`
	// ID of the source account.
	SourceID string `json:"source_id" api:"required"`
	// Optional.
	BillID param.Opt[string] `json:"bill_id,omitzero"`
	// Foreign amount of the transaction.
	ForeignAmount param.Opt[string] `json:"foreign_amount,omitzero"`
	// Submit either a foreign_currency_id or a foreign_currency_code, or neither.
	ForeignCurrencyCode param.Opt[string] `json:"foreign_currency_code,omitzero"`
	// Submit either a foreign_currency_id or a foreign_currency_code, or neither.
	ForeignCurrencyID param.Opt[string] `json:"foreign_currency_id,omitzero"`
	// Optional.
	PiggyBankID param.Opt[string] `json:"piggy_bank_id,omitzero"`
	// The budget ID for this transaction.
	BudgetID param.Opt[string] `json:"budget_id,omitzero"`
	// Category ID for this transaction.
	CategoryID param.Opt[string] `json:"category_id,omitzero"`
	// Submit either a currency_id or a currency_code.
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// Submit either a currency_id or a currency_code.
	CurrencyID param.Opt[string] `json:"currency_id,omitzero"`
	// Array of tags.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r RecurrenceNewParamsTransaction) MarshalJSON() (data []byte, err error) {
	type shadow RecurrenceNewParamsTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecurrenceNewParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurrenceGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type RecurrenceUpdateParams struct {
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Max number of created transactions. Use either this field or repeat_until.
	NrOfRepetitions param.Opt[int64] `json:"nr_of_repetitions,omitzero"`
	// Date until when the recurring transaction can fire. After that date, it's
	// basically inactive. Use either this field or repetitions.
	RepeatUntil param.Opt[time.Time] `json:"repeat_until,omitzero" format:"date"`
	// If the recurrence is even active.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Whether or not to fire the rules after the creation of a transaction.
	ApplyRules param.Opt[bool] `json:"apply_rules,omitzero"`
	// Not to be confused with the description of the actual transaction(s) being
	// created.
	Description param.Opt[string] `json:"description,omitzero"`
	// First time the recurring transaction will fire.
	FirstDate    param.Opt[time.Time]                `json:"first_date,omitzero" format:"date"`
	Title        param.Opt[string]                   `json:"title,omitzero"`
	XTraceID     param.Opt[string]                   `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	Repetitions  []RecurrenceUpdateParamsRepetition  `json:"repetitions,omitzero"`
	Transactions []RecurrenceUpdateParamsTransaction `json:"transactions,omitzero"`
	paramObj
}

func (r RecurrenceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RecurrenceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecurrenceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurrenceUpdateParamsRepetition struct {
	// Information that defined the type of repetition.
	//
	//   - For 'daily', this is empty.
	//   - For 'weekly', it is day of the week between 1 and 7 (Monday - Sunday).
	//   - For 'ndom', it is '1,2' or '4,5' or something else, where the first number is
	//     the week in the month, and the second number is the day in the week (between 1
	//     and 7). '2,3' means: the 2nd Wednesday of the month
	//   - For 'monthly' it is the day of the month (1 - 31)
	//   - For yearly, it is a full date, ie '2026-04-01'. The year you use does not
	//     matter.
	Moment param.Opt[string] `json:"moment,omitzero"`
	// How many occurrences to skip. 0 means skip nothing. 1 means every other.
	Skip param.Opt[int64] `json:"skip,omitzero"`
	// How to respond when the recurring transaction falls in the weekend. Possible
	// values:
	//
	// 1. Do nothing, just create it
	// 2. Create no transaction.
	// 3. Skip to the previous Friday.
	// 4. Skip to the next Monday.
	Weekend param.Opt[int64] `json:"weekend,omitzero"`
	// The type of the repetition. ndom means: the n-th weekday of the month, where you
	// can also specify which day of the week.
	//
	// Any of "daily", "weekly", "ndom", "monthly", "yearly".
	Type RecurrenceRepetitionType `json:"type,omitzero"`
	paramObj
}

func (r RecurrenceUpdateParamsRepetition) MarshalJSON() (data []byte, err error) {
	type shadow RecurrenceUpdateParamsRepetition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecurrenceUpdateParamsRepetition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type RecurrenceUpdateParamsTransaction struct {
	ID string `json:"id" api:"required"`
	// Optional.
	BillID param.Opt[string] `json:"bill_id,omitzero"`
	// Foreign amount of the transaction.
	ForeignAmount param.Opt[string] `json:"foreign_amount,omitzero"`
	// Submit either a foreign_currency_id or a foreign_currency_code, or neither.
	ForeignCurrencyID param.Opt[string] `json:"foreign_currency_id,omitzero"`
	PiggyBankID       param.Opt[string] `json:"piggy_bank_id,omitzero"`
	// Amount of the transaction.
	Amount param.Opt[string] `json:"amount,omitzero"`
	// The budget ID for this transaction.
	BudgetID param.Opt[string] `json:"budget_id,omitzero"`
	// Category ID for this transaction.
	CategoryID param.Opt[string] `json:"category_id,omitzero"`
	// Submit either a currency_id or a currency_code.
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// Submit either a currency_id or a currency_code.
	CurrencyID  param.Opt[string] `json:"currency_id,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// ID of the destination account. Submit either this or destination_name.
	DestinationID param.Opt[string] `json:"destination_id,omitzero"`
	// ID of the source account. Submit either this or source_name.
	SourceID param.Opt[string] `json:"source_id,omitzero"`
	// Array of tags.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r RecurrenceUpdateParamsTransaction) MarshalJSON() (data []byte, err error) {
	type shadow RecurrenceUpdateParamsTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecurrenceUpdateParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecurrenceListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [RecurrenceListParams]'s query parameters as `url.Values`.
func (r RecurrenceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RecurrenceDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type RecurrenceListTransactionsParams struct {
	// A date formatted YYYY-MM-DD. Both the start and end date must be present.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD. Both the start and end date must be present.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "all", "withdrawal", "withdrawals", "expense", "deposit", "deposits",
	// "income", "transfer", "transfers", "opening_balance", "reconciliation",
	// "special", "specials", "default".
	Type TransactionTypeFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecurrenceListTransactionsParams]'s query parameters as
// `url.Values`.
func (r RecurrenceListTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RecurrenceTriggerTransactionParams struct {
	// A date formatted YYYY-MM-DD. This is the date for which you want the recurrence
	// to fire. You can take the date from the list of occurrences in the recurring
	// transaction.
	Date     time.Time         `query:"date" api:"required" format:"date" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [RecurrenceTriggerTransactionParams]'s query parameters as
// `url.Values`.
func (r RecurrenceTriggerTransactionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
