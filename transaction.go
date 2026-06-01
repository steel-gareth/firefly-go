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
)

// The most-used endpoints in Firefly III, these endpoints are used to manage the
// user&#039;s transactions.
//
// TransactionService contains methods and other services that help with
// interacting with the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	options []option.RequestOption
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r TransactionService) {
	r = TransactionService{}
	r.options = opts
	return
}

// Creates a new transaction. The data required can be submitted as a JSON body or
// as a list of parameters.
func (r *TransactionService) New(ctx context.Context, params TransactionNewParams, opts ...option.RequestOption) (res *TransactionSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single transaction.
func (r *TransactionService) Get(ctx context.Context, id string, query TransactionGetParams, opts ...option.RequestOption) (res *TransactionSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transactions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing transaction.
func (r *TransactionService) Update(ctx context.Context, id string, params TransactionUpdateParams, opts ...option.RequestOption) (res *TransactionSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transactions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all the user's transactions.
func (r *TransactionService) List(ctx context.Context, params TransactionListParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a transaction.
func (r *TransactionService) Delete(ctx context.Context, id string, body TransactionDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/transactions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Lists all attachments.
func (r *TransactionService) ListAttachments(ctx context.Context, id string, params TransactionListAttachmentsParams, opts ...option.RequestOption) (res *AttachmentArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transactions/%s/attachments", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Lists all piggy bank events.
func (r *TransactionService) ListPiggyBankEvents(ctx context.Context, id string, params TransactionListPiggyBankEventsParams, opts ...option.RequestOption) (res *PiggyBankEventArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transactions/%s/piggy-bank-events", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type TransactionTypeProperty string

const (
	TransactionTypePropertyWithdrawal     TransactionTypeProperty = "withdrawal"
	TransactionTypePropertyDeposit        TransactionTypeProperty = "deposit"
	TransactionTypePropertyTransfer       TransactionTypeProperty = "transfer"
	TransactionTypePropertyReconciliation TransactionTypeProperty = "reconciliation"
	TransactionTypePropertyOpeningBalance TransactionTypeProperty = "opening balance"
)

type TransactionNewParams struct {
	Transactions []TransactionNewParamsTransaction `json:"transactions,omitzero" api:"required"`
	// Title of the transaction if it has been split in more than one piece. Empty
	// otherwise.
	GroupTitle param.Opt[string] `json:"group_title,omitzero"`
	// Whether or not to apply rules when submitting transaction.
	ApplyRules param.Opt[bool] `json:"apply_rules,omitzero"`
	// Break if the submitted transaction exists already.
	ErrorIfDuplicateHash param.Opt[bool] `json:"error_if_duplicate_hash,omitzero"`
	// Whether or not to fire the webhooks that are related to this event.
	FireWebhooks param.Opt[bool]   `json:"fire_webhooks,omitzero"`
	XTraceID     param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r TransactionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Date, Description, Type are required.
type TransactionNewParamsTransaction struct {
	// Amount of the transaction.
	Amount string `json:"amount" api:"required"`
	// Date of the transaction
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Description of the transaction.
	Description string `json:"description" api:"required"`
	// Any of "withdrawal", "deposit", "transfer", "reconciliation", "opening balance".
	Type TransactionTypeProperty `json:"type,omitzero" api:"required"`
	// Optional. Use either this or the bill_name
	BillID param.Opt[string] `json:"bill_id,omitzero"`
	// Optional. Use either this or the bill_id
	BillName param.Opt[string]    `json:"bill_name,omitzero"`
	BookDate param.Opt[time.Time] `json:"book_date,omitzero" format:"date-time"`
	// The budget ID for this transaction.
	BudgetID param.Opt[string] `json:"budget_id,omitzero"`
	// The name of the budget to be used. If the budget name is unknown, the ID will be
	// used or the value will be ignored.
	BudgetName param.Opt[string] `json:"budget_name,omitzero"`
	// The category ID for this transaction.
	CategoryID param.Opt[string] `json:"category_id,omitzero"`
	// The name of the category to be used. If the category is unknown, it will be
	// created. If the ID and the name point to different categories, the ID overrules
	// the name.
	CategoryName param.Opt[string] `json:"category_name,omitzero"`
	// Currency code. Default is the source account's currency, or the user's financial
	// administration's primary currency. The value you submit may be overruled by the
	// source or destination account.
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// Currency ID. Default is the source account's currency, or the user's financial
	// administration's currency. The value you submit may be overruled by the source
	// or destination account.
	CurrencyID param.Opt[string] `json:"currency_id,omitzero"`
	// ID of the destination account. For a deposit or a transfer, this must always be
	// an asset account. For withdrawals this must be an expense account.
	DestinationID param.Opt[string] `json:"destination_id,omitzero"`
	// Name of the destination account. You can submit the name instead of the ID. For
	// everything except transfers, the account will be auto-generated if unknown, so
	// submitting a name is enough.
	DestinationName param.Opt[string]    `json:"destination_name,omitzero"`
	DueDate         param.Opt[time.Time] `json:"due_date,omitzero" format:"date-time"`
	// Reference to external ID in other systems.
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	// External, custom URL for this transaction.
	ExternalURL param.Opt[string] `json:"external_url,omitzero"`
	// The amount in a foreign currency.
	ForeignAmount param.Opt[string] `json:"foreign_amount,omitzero"`
	// Currency code of the foreign currency. Default is NULL. Can be used instead of
	// the foreign_currency_id, but this or the ID is required when submitting a
	// foreign amount.
	ForeignCurrencyCode param.Opt[string] `json:"foreign_currency_code,omitzero"`
	// Currency ID of the foreign currency. Default is null. Is required when you
	// submit a foreign amount.
	ForeignCurrencyID param.Opt[string]    `json:"foreign_currency_id,omitzero"`
	InterestDate      param.Opt[time.Time] `json:"interest_date,omitzero" format:"date-time"`
	// Reference to internal reference of other systems.
	InternalReference param.Opt[string]    `json:"internal_reference,omitzero"`
	InvoiceDate       param.Opt[time.Time] `json:"invoice_date,omitzero" format:"date-time"`
	Notes             param.Opt[string]    `json:"notes,omitzero"`
	// Order of this entry in the list of transactions.
	Order       param.Opt[int64]     `json:"order,omitzero"`
	PaymentDate param.Opt[time.Time] `json:"payment_date,omitzero" format:"date-time"`
	// Optional. Use either this or the piggy_bank_name
	PiggyBankID param.Opt[int64] `json:"piggy_bank_id,omitzero"`
	// Optional. Use either this or the piggy_bank_id
	PiggyBankName param.Opt[string]    `json:"piggy_bank_name,omitzero"`
	ProcessDate   param.Opt[time.Time] `json:"process_date,omitzero" format:"date-time"`
	// SEPA Batch ID
	SepaBatchID param.Opt[string] `json:"sepa_batch_id,omitzero"`
	// SEPA Clearing Code
	SepaCc param.Opt[string] `json:"sepa_cc,omitzero"`
	// SEPA Creditor Identifier
	SepaCi param.Opt[string] `json:"sepa_ci,omitzero"`
	// SEPA Country
	SepaCountry param.Opt[string] `json:"sepa_country,omitzero"`
	// SEPA end-to-end Identifier
	SepaCtID param.Opt[string] `json:"sepa_ct_id,omitzero"`
	// SEPA Opposing Account Identifier
	SepaCtOp param.Opt[string] `json:"sepa_ct_op,omitzero"`
	// SEPA mandate identifier
	SepaDB param.Opt[string] `json:"sepa_db,omitzero"`
	// SEPA External Purpose indicator
	SepaEp param.Opt[string] `json:"sepa_ep,omitzero"`
	// ID of the source account. For a withdrawal or a transfer, this must always be an
	// asset account. For deposits, this must be a revenue account.
	SourceID param.Opt[string] `json:"source_id,omitzero"`
	// Name of the source account. For a withdrawal or a transfer, this must always be
	// an asset account. For deposits, this must be a revenue account. Can be used
	// instead of the source_id. If the transaction is a deposit, the source_name can
	// be filled in freely: the account will be created based on the name.
	SourceName param.Opt[string] `json:"source_name,omitzero"`
	// If the transaction has been reconciled already. When you set this, the amount
	// can no longer be edited by the user.
	Reconciled param.Opt[bool] `json:"reconciled,omitzero"`
	// Array of tags.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r TransactionNewParamsTransaction) MarshalJSON() (data []byte, err error) {
	type shadow TransactionNewParamsTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionNewParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type TransactionUpdateParams struct {
	// Title of the transaction if it has been split in more than one piece. Empty
	// otherwise.
	GroupTitle param.Opt[string] `json:"group_title,omitzero"`
	// Whether or not to apply rules when submitting transaction.
	ApplyRules param.Opt[bool] `json:"apply_rules,omitzero"`
	// Whether or not to fire the webhooks that are related to this event.
	FireWebhooks param.Opt[bool]                      `json:"fire_webhooks,omitzero"`
	XTraceID     param.Opt[string]                    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	Transactions []TransactionUpdateParamsTransaction `json:"transactions,omitzero"`
	paramObj
}

func (r TransactionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionUpdateParamsTransaction struct {
	// Optional. Use either this or the bill_name
	BillID param.Opt[string] `json:"bill_id,omitzero"`
	// Optional. Use either this or the bill_id
	BillName param.Opt[string]    `json:"bill_name,omitzero"`
	BookDate param.Opt[time.Time] `json:"book_date,omitzero" format:"date-time"`
	// The budget ID for this transaction.
	BudgetID param.Opt[string] `json:"budget_id,omitzero"`
	// The category ID for this transaction.
	CategoryID param.Opt[string] `json:"category_id,omitzero"`
	// The name of the category to be used. If the category is unknown, it will be
	// created. If the ID and the name point to different categories, the ID overrules
	// the name.
	CategoryName param.Opt[string] `json:"category_name,omitzero"`
	// Currency code. Default is the source account's currency, or the user's financial
	// administration's primary currency. Can be used instead of currency_id.
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// Currency ID. Default is the source account's currency, or the user's financial
	// administration's primary currency. Can be used instead of currency_code.
	CurrencyID      param.Opt[string] `json:"currency_id,omitzero"`
	DestinationIban param.Opt[string] `json:"destination_iban,omitzero"`
	// ID of the destination account. For a deposit or a transfer, this must always be
	// an asset account. For withdrawals this must be an expense account.
	DestinationID param.Opt[string] `json:"destination_id,omitzero"`
	// Name of the destination account. You can submit the name instead of the ID. For
	// everything except transfers, the account will be auto-generated if unknown, so
	// submitting a name is enough.
	DestinationName param.Opt[string]    `json:"destination_name,omitzero"`
	DueDate         param.Opt[time.Time] `json:"due_date,omitzero" format:"date-time"`
	// Reference to external ID in other systems.
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	// External, custom URL for this transaction.
	ExternalURL param.Opt[string] `json:"external_url,omitzero"`
	// The amount in a foreign currency.
	ForeignAmount param.Opt[string] `json:"foreign_amount,omitzero"`
	// Currency code of the foreign currency. Default is NULL. Can be used instead of
	// the foreign_currency_id, but this or the ID is required when submitting a
	// foreign amount.
	ForeignCurrencyCode param.Opt[string] `json:"foreign_currency_code,omitzero"`
	// Currency ID of the foreign currency. Default is null. Is required when you
	// submit a foreign amount.
	ForeignCurrencyID param.Opt[string]    `json:"foreign_currency_id,omitzero"`
	InterestDate      param.Opt[time.Time] `json:"interest_date,omitzero" format:"date-time"`
	// Reference to internal reference of other systems.
	InternalReference param.Opt[string]    `json:"internal_reference,omitzero"`
	InvoiceDate       param.Opt[time.Time] `json:"invoice_date,omitzero" format:"date-time"`
	Notes             param.Opt[string]    `json:"notes,omitzero"`
	// Order of this entry in the list of transactions.
	Order       param.Opt[int64]     `json:"order,omitzero"`
	PaymentDate param.Opt[time.Time] `json:"payment_date,omitzero" format:"date-time"`
	ProcessDate param.Opt[time.Time] `json:"process_date,omitzero" format:"date-time"`
	// SEPA Batch ID
	SepaBatchID param.Opt[string] `json:"sepa_batch_id,omitzero"`
	// SEPA Clearing Code
	SepaCc param.Opt[string] `json:"sepa_cc,omitzero"`
	// SEPA Creditor Identifier
	SepaCi param.Opt[string] `json:"sepa_ci,omitzero"`
	// SEPA Country
	SepaCountry param.Opt[string] `json:"sepa_country,omitzero"`
	// SEPA end-to-end Identifier
	SepaCtID param.Opt[string] `json:"sepa_ct_id,omitzero"`
	// SEPA Opposing Account Identifier
	SepaCtOp param.Opt[string] `json:"sepa_ct_op,omitzero"`
	// SEPA mandate identifier
	SepaDB param.Opt[string] `json:"sepa_db,omitzero"`
	// SEPA External Purpose indicator
	SepaEp     param.Opt[string] `json:"sepa_ep,omitzero"`
	SourceIban param.Opt[string] `json:"source_iban,omitzero"`
	// ID of the source account. For a withdrawal or a transfer, this must always be an
	// asset account. For deposits, this must be a revenue account.
	SourceID param.Opt[string] `json:"source_id,omitzero"`
	// Name of the source account. For a withdrawal or a transfer, this must always be
	// an asset account. For deposits, this must be a revenue account. Can be used
	// instead of the source_id. If the transaction is a deposit, the source_name can
	// be filled in freely: the account will be created based on the name.
	SourceName param.Opt[string] `json:"source_name,omitzero"`
	// Amount of the transaction.
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Date of the transaction
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Description of the transaction.
	Description param.Opt[string] `json:"description,omitzero"`
	// If the transaction has been reconciled already. When you set this, the amount
	// can no longer be edited by the user.
	Reconciled param.Opt[bool] `json:"reconciled,omitzero"`
	// Transaction journal ID of current transaction (split).
	TransactionJournalID param.Opt[string] `json:"transaction_journal_id,omitzero"`
	// Array of tags.
	Tags []string `json:"tags,omitzero"`
	// Any of "withdrawal", "deposit", "transfer", "reconciliation", "opening balance".
	Type TransactionTypeProperty `json:"type,omitzero"`
	paramObj
}

func (r TransactionUpdateParamsTransaction) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParamsTransaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParamsTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListParams struct {
	// A date formatted YYYY-MM-DD. This is the end date of the selected range
	// (inclusive).
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD. This is the start date of the selected range
	// (inclusive).
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "all", "withdrawal", "withdrawals", "expense", "deposit", "deposits",
	// "income", "transfer", "transfers", "opening_balance", "reconciliation",
	// "special", "specials", "default".
	Type TransactionTypeFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransactionDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type TransactionListAttachmentsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListAttachmentsParams]'s query parameters as
// `url.Values`.
func (r TransactionListAttachmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransactionListPiggyBankEventsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListPiggyBankEventsParams]'s query parameters as
// `url.Values`.
func (r TransactionListPiggyBankEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
