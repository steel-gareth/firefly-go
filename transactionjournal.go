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

// The most-used endpoints in Firefly III, these endpoints are used to manage the
// user&#039;s transactions.
//
// TransactionJournalService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionJournalService] method instead.
type TransactionJournalService struct {
	options []option.RequestOption
}

// NewTransactionJournalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionJournalService(opts ...option.RequestOption) (r TransactionJournalService) {
	r = TransactionJournalService{}
	r.options = opts
	return
}

// Get a single transaction by underlying journal (split).
func (r *TransactionJournalService) Get(ctx context.Context, id string, query TransactionJournalGetParams, opts ...option.RequestOption) (res *TransactionSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction-journals/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete an individual journal (split) from a transaction.
func (r *TransactionJournalService) Delete(ctx context.Context, id string, body TransactionJournalDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/transaction-journals/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Lists all the transaction links for an individual journal (a split). Don't use
// the group ID, you need the actual underlying journal (the split).
func (r *TransactionJournalService) ListLinks(ctx context.Context, id string, params TransactionJournalListLinksParams, opts ...option.RequestOption) (res *TransactionLinkArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transaction-journals/%s/links", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type TransactionRead struct {
	ID         string                    `json:"id" api:"required"`
	Attributes TransactionReadAttributes `json:"attributes" api:"required"`
	Links      ObjectLink                `json:"links" api:"required"`
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
func (r TransactionRead) RawJSON() string { return r.JSON.raw }
func (r *TransactionRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionReadAttributes struct {
	Transactions []TransactionReadAttributesTransaction `json:"transactions" api:"required"`
	CreatedAt    time.Time                              `json:"created_at" format:"date-time"`
	// Title of the transaction if it has been split in more than one piece. Empty
	// otherwise.
	GroupTitle string    `json:"group_title" api:"nullable"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	// User ID
	User string `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transactions respjson.Field
		CreatedAt    respjson.Field
		GroupTitle   respjson.Field
		UpdatedAt    respjson.Field
		User         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *TransactionReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionReadAttributesTransaction struct {
	// Amount of the transaction.
	Amount string `json:"amount" api:"required"`
	// Date of the transaction
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Description of the transaction.
	Description string `json:"description" api:"required"`
	// ID of the destination account. For a deposit or a transfer, this must always be
	// an asset account. For withdrawals this must be an expense account.
	DestinationID string `json:"destination_id" api:"required"`
	// ID of the source account. For a withdrawal or a transfer, this must always be an
	// asset account. For deposits, this must be a revenue account.
	SourceID string `json:"source_id" api:"required"`
	// Any of "withdrawal", "deposit", "transfer", "reconciliation", "opening balance".
	Type TransactionTypeProperty `json:"type" api:"required"`
	// The associated subscription ID for this transaction. `bill` refers to the OLD
	// name for subscriptions and this field will be removed.
	BillID string `json:"bill_id" api:"nullable"`
	// The associated subscription name for this transaction. `bill` refers to the OLD
	// name for subscriptions and this field will be removed.
	BillName string    `json:"bill_name" api:"nullable"`
	BookDate time.Time `json:"book_date" api:"nullable" format:"date-time"`
	// The budget ID for this transaction.
	BudgetID string `json:"budget_id" api:"nullable"`
	// The name of the budget used.
	BudgetName string `json:"budget_name" api:"nullable"`
	// The category ID for this transaction.
	CategoryID string `json:"category_id" api:"nullable"`
	// The name of the category to be used. If the category is unknown, it will be
	// created. If the ID and the name point to different categories, the ID overrules
	// the name.
	CategoryName string `json:"category_name" api:"nullable"`
	// Currency code for the currency of this transaction.
	CurrencyCode string `json:"currency_code"`
	// Number of decimals used in this currency.
	CurrencyDecimalPlaces int64 `json:"currency_decimal_places"`
	// Currency ID for the currency of this transaction.
	CurrencyID string `json:"currency_id"`
	// Currency name for the currency of this transaction.
	CurrencyName string `json:"currency_name"`
	// Currency symbol for the currency of this transaction.
	CurrencySymbol string `json:"currency_symbol"`
	// The balance of the destination account. This is the balance in the account's
	// currency which may be different from this transaction, and is not provided in
	// this model.
	DestinationBalanceAfter string `json:"destination_balance_after" api:"nullable"`
	DestinationIban         string `json:"destination_iban" api:"nullable"`
	// Name of the destination account. You can submit the name instead of the ID. For
	// everything except transfers, the account will be auto-generated if unknown, so
	// submitting a name is enough.
	DestinationName string `json:"destination_name" api:"nullable"`
	// Any of "Default account", "Cash account", "Asset account", "Expense account",
	// "Revenue account", "Initial balance account", "Beneficiary account", "Import
	// account", "Reconciliation account", "Loan", "Debt", "Mortgage".
	DestinationType AccountTypeProperty `json:"destination_type"`
	DueDate         time.Time           `json:"due_date" api:"nullable" format:"date-time"`
	// Reference to external ID in other systems.
	ExternalID string `json:"external_id" api:"nullable"`
	// External, custom URL for this transaction.
	ExternalURL string `json:"external_url" api:"nullable"`
	// The amount in the set foreign currency. May be NULL if the transaction does not
	// have a foreign amount.
	ForeignAmount string `json:"foreign_amount" api:"nullable"`
	// Currency code of the foreign currency. Default is NULL.
	ForeignCurrencyCode string `json:"foreign_currency_code" api:"nullable"`
	// Number of decimals in the foreign currency.
	ForeignCurrencyDecimalPlaces int64 `json:"foreign_currency_decimal_places" api:"nullable"`
	// Currency ID of the foreign currency, if this transaction has a foreign amount.
	ForeignCurrencyID     string `json:"foreign_currency_id" api:"nullable"`
	ForeignCurrencySymbol string `json:"foreign_currency_symbol" api:"nullable"`
	// If the transaction has attachments.
	HasAttachments bool `json:"has_attachments"`
	// Hash value of original import transaction (for duplicate detection).
	ImportHashV2 string    `json:"import_hash_v2" api:"nullable"`
	InterestDate time.Time `json:"interest_date" api:"nullable" format:"date-time"`
	// Reference to internal reference of other systems.
	InternalReference string    `json:"internal_reference" api:"nullable"`
	InvoiceDate       time.Time `json:"invoice_date" api:"nullable" format:"date-time"`
	// Latitude of the transaction's location, if applicable. Can be used to draw a
	// map.
	Latitude float64 `json:"latitude" api:"nullable"`
	// Latitude of the transaction's location, if applicable. Can be used to draw a
	// map.
	Longitude float64 `json:"longitude" api:"nullable"`
	Notes     string  `json:"notes" api:"nullable"`
	// Indicates whether the transaction has a currency setting. For transactions this
	// is always true.
	ObjectHasCurrencySetting bool `json:"object_has_currency_setting"`
	// Order of this entry in the list of transactions.
	Order int64 `json:"order" api:"nullable"`
	// System generated identifier for original creator of transaction.
	OriginalSource string    `json:"original_source" api:"nullable"`
	PaymentDate    time.Time `json:"payment_date" api:"nullable" format:"date-time"`
	// Amount of the transaction in the primary currency of this administration. The
	// `primary_currency_*` fields reflect the currency used. This field is NULL if the
	// user does have 'convert to primary' set to true in their settings.
	PcAmount string `json:"pc_amount"`
	// The balance of the destination account in the primary currency of this
	// administration. The `primary_currency_*` fields reflect the currency used. This
	// field is NULL if the user does have 'convert to primary' set to true in their
	// settings.
	PcDestinationBalanceAfter string `json:"pc_destination_balance_after" api:"nullable"`
	// Foreign amount of the transaction in the primary currency of this
	// administration. The `primary_currency_*` fields reflect the currency used. This
	// field is NULL if the user does have 'convert to primary' set to true in their
	// settings.
	PcForeignAmount string `json:"pc_foreign_amount"`
	// The balance of the source account in the primary currency of this
	// administration. The `primary_currency_*` fields reflect the currency used. This
	// field is NULL if the user does have 'convert to primary' set to true in their
	// settings.
	PcSourceBalanceAfter string `json:"pc_source_balance_after" api:"nullable"`
	// Returns the primary currency code of the administration. This currency is used
	// as the currency for all `pc_*` amount and balance fields of this account.
	PrimaryCurrencyCode string `json:"primary_currency_code" api:"nullable"`
	// See the other `primary_*` fields.
	PrimaryCurrencyDecimalPlaces int64 `json:"primary_currency_decimal_places" api:"nullable"`
	// Returns the primary currency ID of the administration. This currency is used as
	// the currency for all `pc_*` amount and balance fields of this account.
	PrimaryCurrencyID string `json:"primary_currency_id" api:"nullable"`
	// See the other `primary_*` fields.
	PrimaryCurrencySymbol string    `json:"primary_currency_symbol" api:"nullable"`
	ProcessDate           time.Time `json:"process_date" api:"nullable" format:"date-time"`
	// If the transaction has been reconciled already. When you set this, the amount
	// can no longer be edited by the user.
	Reconciled bool `json:"reconciled"`
	// The # of the current transaction created under this recurrence.
	RecurrenceCount int64 `json:"recurrence_count" api:"nullable"`
	// Reference to recurrence that made the transaction.
	RecurrenceID string `json:"recurrence_id" api:"nullable"`
	// Total number of transactions expected to be created by this recurrence
	// repetition. Will be 0 if infinite.
	RecurrenceTotal int64 `json:"recurrence_total" api:"nullable"`
	// SEPA Batch ID
	SepaBatchID string `json:"sepa_batch_id" api:"nullable"`
	// SEPA Clearing Code
	SepaCc string `json:"sepa_cc" api:"nullable"`
	// SEPA Creditor Identifier
	SepaCi string `json:"sepa_ci" api:"nullable"`
	// SEPA Country
	SepaCountry string `json:"sepa_country" api:"nullable"`
	// SEPA end-to-end Identifier
	SepaCtID string `json:"sepa_ct_id" api:"nullable"`
	// SEPA Opposing Account Identifier
	SepaCtOp string `json:"sepa_ct_op" api:"nullable"`
	// SEPA mandate identifier
	SepaDB string `json:"sepa_db" api:"nullable"`
	// SEPA External Purpose indicator
	SepaEp string `json:"sepa_ep" api:"nullable"`
	// The balance of the source account. This is the balance in the account's currency
	// which may be different from this transaction, and is not provided in this model.
	SourceBalanceAfter string `json:"source_balance_after" api:"nullable"`
	SourceIban         string `json:"source_iban" api:"nullable"`
	// Name of the source account. For a withdrawal or a transfer, this must always be
	// an asset account. For deposits, this must be a revenue account. Can be used
	// instead of the source_id. If the transaction is a deposit, the source_name can
	// be filled in freely: the account will be created based on the name.
	SourceName string `json:"source_name" api:"nullable"`
	// Any of "Default account", "Cash account", "Asset account", "Expense account",
	// "Revenue account", "Initial balance account", "Beneficiary account", "Import
	// account", "Reconciliation account", "Loan", "Debt", "Mortgage".
	SourceType AccountTypeProperty `json:"source_type"`
	// The associated subscription ID for this transaction.
	SubscriptionID string `json:"subscription_id" api:"nullable"`
	// The associated subscription name for this transaction.
	SubscriptionName string `json:"subscription_name" api:"nullable"`
	// Array of tags.
	Tags []string `json:"tags" api:"nullable"`
	// ID of the underlying transaction journal. Each transaction consists of a
	// transaction group (see the top ID) and one or more journals making up the splits
	// of the transaction.
	TransactionJournalID string `json:"transaction_journal_id"`
	// User ID
	User string `json:"user"`
	// Zoom level for the map, if drawn. This to set the box right. Unfortunately this
	// is a proprietary value because each map provider has different zoom levels.
	ZoomLevel int64 `json:"zoom_level" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount                       respjson.Field
		Date                         respjson.Field
		Description                  respjson.Field
		DestinationID                respjson.Field
		SourceID                     respjson.Field
		Type                         respjson.Field
		BillID                       respjson.Field
		BillName                     respjson.Field
		BookDate                     respjson.Field
		BudgetID                     respjson.Field
		BudgetName                   respjson.Field
		CategoryID                   respjson.Field
		CategoryName                 respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		DestinationBalanceAfter      respjson.Field
		DestinationIban              respjson.Field
		DestinationName              respjson.Field
		DestinationType              respjson.Field
		DueDate                      respjson.Field
		ExternalID                   respjson.Field
		ExternalURL                  respjson.Field
		ForeignAmount                respjson.Field
		ForeignCurrencyCode          respjson.Field
		ForeignCurrencyDecimalPlaces respjson.Field
		ForeignCurrencyID            respjson.Field
		ForeignCurrencySymbol        respjson.Field
		HasAttachments               respjson.Field
		ImportHashV2                 respjson.Field
		InterestDate                 respjson.Field
		InternalReference            respjson.Field
		InvoiceDate                  respjson.Field
		Latitude                     respjson.Field
		Longitude                    respjson.Field
		Notes                        respjson.Field
		ObjectHasCurrencySetting     respjson.Field
		Order                        respjson.Field
		OriginalSource               respjson.Field
		PaymentDate                  respjson.Field
		PcAmount                     respjson.Field
		PcDestinationBalanceAfter    respjson.Field
		PcForeignAmount              respjson.Field
		PcSourceBalanceAfter         respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		ProcessDate                  respjson.Field
		Reconciled                   respjson.Field
		RecurrenceCount              respjson.Field
		RecurrenceID                 respjson.Field
		RecurrenceTotal              respjson.Field
		SepaBatchID                  respjson.Field
		SepaCc                       respjson.Field
		SepaCi                       respjson.Field
		SepaCountry                  respjson.Field
		SepaCtID                     respjson.Field
		SepaCtOp                     respjson.Field
		SepaDB                       respjson.Field
		SepaEp                       respjson.Field
		SourceBalanceAfter           respjson.Field
		SourceIban                   respjson.Field
		SourceName                   respjson.Field
		SourceType                   respjson.Field
		SubscriptionID               respjson.Field
		SubscriptionName             respjson.Field
		Tags                         respjson.Field
		TransactionJournalID         respjson.Field
		User                         respjson.Field
		ZoomLevel                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionReadAttributesTransaction) RawJSON() string { return r.JSON.raw }
func (r *TransactionReadAttributesTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionSingle struct {
	Data TransactionRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionSingle) RawJSON() string { return r.JSON.raw }
func (r *TransactionSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionJournalGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type TransactionJournalDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type TransactionJournalListLinksParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionJournalListLinksParams]'s query parameters as
// `url.Values`.
func (r TransactionJournalListLinksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
