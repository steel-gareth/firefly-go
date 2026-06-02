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

// Endpoints that deliver all of the user&#039;s asset, expense and other accounts
// (and the metadata) together with related transactions, piggy banks and other
// objects. Also delivers endpoints for CRUD operations for accounts.
//
// AccountService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.options = opts
	return
}

// Creates a new account. The data required can be submitted as a JSON body or as a
// list of parameters (in key=value pairs, like a webform).
func (r *AccountService) New(ctx context.Context, params AccountNewParams, opts ...option.RequestOption) (res *AccountSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single account by its ID.
func (r *AccountService) Get(ctx context.Context, id string, params AccountGetParams, opts ...option.RequestOption) (res *AccountSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Used to update a single account. All fields that are not submitted will be
// cleared (set to NULL). The model will tell you which fields are mandatory.
func (r *AccountService) Update(ctx context.Context, id string, params AccountUpdateParams, opts ...option.RequestOption) (res *AccountSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// This endpoint returns a list of all the accounts owned by the authenticated
// user.
func (r *AccountService) List(ctx context.Context, params AccountListParams, opts ...option.RequestOption) (res *AccountArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Will permanently delete an account. Any associated transactions and piggy banks
// are ALSO deleted. Cannot be recovered from.
func (r *AccountService) Delete(ctx context.Context, id string, body AccountDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Lists all attachments.
func (r *AccountService) ListAttachments(ctx context.Context, id string, params AccountListAttachmentsParams, opts ...option.RequestOption) (res *AttachmentArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/attachments", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint returns a list of all the piggy banks connected to the account.
func (r *AccountService) ListPiggyBanks(ctx context.Context, id string, params AccountListPiggyBanksParams, opts ...option.RequestOption) (res *PiggyBankArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/piggy-banks", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// This endpoint returns a list of all the transactions connected to the account.
func (r *AccountService) ListTransactions(ctx context.Context, id string, params AccountListTransactionsParams, opts ...option.RequestOption) (res *TransactionArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s/transactions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type AccountArray struct {
	Data []AccountRead `json:"data" api:"required"`
	Meta Meta          `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountArray) RawJSON() string { return r.JSON.raw }
func (r *AccountArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRead struct {
	ID         string                `json:"id" api:"required"`
	Attributes AccountReadAttributes `json:"attributes" api:"required"`
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
func (r AccountRead) RawJSON() string { return r.JSON.raw }
func (r *AccountRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountReadAttributes struct {
	Name string `json:"name" api:"required"`
	// Can only be one one these account types. import, initial-balance and
	// reconciliation cannot be set manually.
	//
	// Any of "asset", "expense", "import", "revenue", "cash", "liability",
	// "liabilities", "initial-balance", "reconciliation".
	Type          ShortAccountTypeProperty `json:"type" api:"required"`
	AccountNumber string                   `json:"account_number" api:"nullable"`
	// Is only mandatory when the type is asset.
	//
	// Any of "defaultAsset", "sharedAsset", "savingAsset", "ccAsset",
	// "cashWalletAsset".
	AccountRole AccountRoleProperty `json:"account_role" api:"nullable"`
	Active      bool                `json:"active"`
	// If you submit a start AND end date, this will be the difference between those
	// two moments.
	BalanceDifference string    `json:"balance_difference"`
	Bic               string    `json:"bic" api:"nullable"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// Mandatory when the account_role is ccAsset. Can only be monthlyFull or null.
	//
	// Any of "monthlyFull".
	CreditCardType CreditCardTypeProperty `json:"credit_card_type" api:"nullable"`
	// The currency code of the currency associated with this object.
	CurrencyCode          string `json:"currency_code"`
	CurrencyDecimalPlaces int64  `json:"currency_decimal_places"`
	// The currency ID of the currency associated with this object.
	CurrencyID string `json:"currency_id"`
	// The currency name of the currency associated with this object.
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	// The current balance of the account in the account's currency. If the account has
	// no currency, this is the balance in the administration's primary currency.
	// Either way, the `currency_*` fields reflect the currency used.
	CurrentBalance string `json:"current_balance"`
	// The timestamp for this date is always 23:59:59, to indicate it's the balance at
	// the very END of that particular day.
	CurrentBalanceDate time.Time `json:"current_balance_date" format:"date-time"`
	// In liability accounts (loans, debts and mortgages), this is the amount of debt
	// in the account's currency (see the `currency_*` fields). In asset accounts, this
	// is NULL.
	DebtAmount      string `json:"debt_amount" api:"nullable"`
	Iban            string `json:"iban" api:"nullable"`
	IncludeNetWorth bool   `json:"include_net_worth"`
	// Mandatory when type is liability. Interest percentage.
	Interest string `json:"interest" api:"nullable"`
	// Mandatory when type is liability. Period over which the interest is calculated.
	//
	// Any of "daily", "weekly", "monthly", "quarterly", "half-year", "yearly".
	InterestPeriod InterestPeriodProperty `json:"interest_period" api:"nullable"`
	// Last activity of the account.
	LastActivity time.Time `json:"last_activity" api:"nullable" format:"date-time"`
	// Latitude of the accounts's location, if applicable. Can be used to draw a map.
	Latitude float64 `json:"latitude" api:"nullable"`
	// 'credit' indicates somebody owes you the liability. 'debit' Indicates you owe
	// this debt yourself. Works only for liabilities.
	//
	// Any of "credit", "debit".
	LiabilityDirection LiabilityDirectionProperty `json:"liability_direction" api:"nullable"`
	// Mandatory when type is liability. Specifies the exact type.
	//
	// Any of "loan", "debt", "mortgage".
	LiabilityType LiabilityTypeProperty `json:"liability_type" api:"nullable"`
	// Latitude of the accounts's location, if applicable. Can be used to draw a map.
	Longitude float64 `json:"longitude" api:"nullable"`
	// Mandatory when the account_role is ccAsset. Moment at which CC payment
	// installments are asked for by the bank.
	MonthlyPaymentDate time.Time `json:"monthly_payment_date" api:"nullable" format:"date-time"`
	Notes              string    `json:"notes" api:"nullable"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupID string `json:"object_group_id" api:"nullable"`
	// The order of the group. At least 1, for the highest sorting.
	ObjectGroupOrder int64 `json:"object_group_order" api:"nullable"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle string `json:"object_group_title" api:"nullable"`
	// Indicates whether the account has a currency setting. If false, the account uses
	// the administration's primary currency. Asset accounts and liability accounts
	// always have a currency setting, while expense and revenue accounts do not.
	ObjectHasCurrencySetting bool `json:"object_has_currency_setting"`
	// Represents the opening balance, the initial amount this account holds in the
	// currency of the account or the administration's primary currency if the account
	// has no currency. Either way, the `currency_*` fields reflect the currency used.
	OpeningBalance string `json:"opening_balance"`
	// Represents the date of the opening balance.
	OpeningBalanceDate time.Time `json:"opening_balance_date" api:"nullable" format:"date-time"`
	// Order of the account. Is NULL if account is not asset or liability.
	Order int64 `json:"order" api:"nullable"`
	// If you submit a start AND end date, this will be the difference in the currency
	// of the account or the administration's primary currency between those two
	// moments.
	PcBalanceDifference string `json:"pc_balance_difference" api:"nullable"`
	// The current balance of the account in the administration's primary currency. The
	// `primary_currency_*` fields reflect the currency used. This field is NULL if the
	// user does have 'convert to primary' set to true in their settings.
	PcCurrentBalance string `json:"pc_current_balance" api:"nullable"`
	// In liability accounts (loans, debts and mortgages), this is the amount of debt
	// in the administration's primary currency (see the `currency_*` fields. In asset
	// accounts, this is NULL.
	PcDebtAmount string `json:"pc_debt_amount" api:"nullable"`
	// The opening balance of the account in the administration's primary currency
	// (pc). The `primary_currency_*` fields reflect the currency used. This field is
	// NULL if the user does have 'convert to primary' set to true in their settings.
	PcOpeningBalance string `json:"pc_opening_balance"`
	// The virtual balance of the account in the administration's primary currency
	// (pc). The `primary_currency_*` fields reflect the currency used. This field is
	// NULL if the user does have 'convert to primary' set to true in their settings.
	PcVirtualBalance string `json:"pc_virtual_balance"`
	// The currency code of the administration's primary currency.
	PrimaryCurrencyCode string `json:"primary_currency_code"`
	// The currency decimal places of the administration's primary currency.
	PrimaryCurrencyDecimalPlaces int64 `json:"primary_currency_decimal_places"`
	// The currency ID of the administration's primary currency.
	PrimaryCurrencyID string `json:"primary_currency_id"`
	// The currency name of the administration's primary currency.
	PrimaryCurrencyName string `json:"primary_currency_name"`
	// The currency symbol of the administration's primary currency.
	PrimaryCurrencySymbol string    `json:"primary_currency_symbol"`
	UpdatedAt             time.Time `json:"updated_at" format:"date-time"`
	// The virtual balance of the account in the account's currency or the
	// administration's primary currency if the account has no currency.
	VirtualBalance string `json:"virtual_balance"`
	// Zoom level for the map, if drawn. This to set the box right. Unfortunately this
	// is a proprietary value because each map provider has different zoom levels.
	ZoomLevel int64 `json:"zoom_level" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                         respjson.Field
		Type                         respjson.Field
		AccountNumber                respjson.Field
		AccountRole                  respjson.Field
		Active                       respjson.Field
		BalanceDifference            respjson.Field
		Bic                          respjson.Field
		CreatedAt                    respjson.Field
		CreditCardType               respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		CurrentBalance               respjson.Field
		CurrentBalanceDate           respjson.Field
		DebtAmount                   respjson.Field
		Iban                         respjson.Field
		IncludeNetWorth              respjson.Field
		Interest                     respjson.Field
		InterestPeriod               respjson.Field
		LastActivity                 respjson.Field
		Latitude                     respjson.Field
		LiabilityDirection           respjson.Field
		LiabilityType                respjson.Field
		Longitude                    respjson.Field
		MonthlyPaymentDate           respjson.Field
		Notes                        respjson.Field
		ObjectGroupID                respjson.Field
		ObjectGroupOrder             respjson.Field
		ObjectGroupTitle             respjson.Field
		ObjectHasCurrencySetting     respjson.Field
		OpeningBalance               respjson.Field
		OpeningBalanceDate           respjson.Field
		Order                        respjson.Field
		PcBalanceDifference          respjson.Field
		PcCurrentBalance             respjson.Field
		PcDebtAmount                 respjson.Field
		PcOpeningBalance             respjson.Field
		PcVirtualBalance             respjson.Field
		PrimaryCurrencyCode          respjson.Field
		PrimaryCurrencyDecimalPlaces respjson.Field
		PrimaryCurrencyID            respjson.Field
		PrimaryCurrencyName          respjson.Field
		PrimaryCurrencySymbol        respjson.Field
		UpdatedAt                    respjson.Field
		VirtualBalance               respjson.Field
		ZoomLevel                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *AccountReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Is only mandatory when the type is asset.
type AccountRoleProperty string

const (
	AccountRolePropertyDefaultAsset    AccountRoleProperty = "defaultAsset"
	AccountRolePropertySharedAsset     AccountRoleProperty = "sharedAsset"
	AccountRolePropertySavingAsset     AccountRoleProperty = "savingAsset"
	AccountRolePropertyCcAsset         AccountRoleProperty = "ccAsset"
	AccountRolePropertyCashWalletAsset AccountRoleProperty = "cashWalletAsset"
)

type AccountSingle struct {
	Data AccountRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountSingle) RawJSON() string { return r.JSON.raw }
func (r *AccountSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTypeFilter string

const (
	AccountTypeFilterAll                   AccountTypeFilter = "all"
	AccountTypeFilterAsset                 AccountTypeFilter = "asset"
	AccountTypeFilterCash                  AccountTypeFilter = "cash"
	AccountTypeFilterExpense               AccountTypeFilter = "expense"
	AccountTypeFilterRevenue               AccountTypeFilter = "revenue"
	AccountTypeFilterSpecial               AccountTypeFilter = "special"
	AccountTypeFilterHidden                AccountTypeFilter = "hidden"
	AccountTypeFilterLiability             AccountTypeFilter = "liability"
	AccountTypeFilterLiabilities           AccountTypeFilter = "liabilities"
	AccountTypeFilterDefaultAccount        AccountTypeFilter = "Default account"
	AccountTypeFilterCashAccount           AccountTypeFilter = "Cash account"
	AccountTypeFilterAssetAccount          AccountTypeFilter = "Asset account"
	AccountTypeFilterExpenseAccount        AccountTypeFilter = "Expense account"
	AccountTypeFilterRevenueAccount        AccountTypeFilter = "Revenue account"
	AccountTypeFilterInitialBalanceAccount AccountTypeFilter = "Initial balance account"
	AccountTypeFilterBeneficiaryAccount    AccountTypeFilter = "Beneficiary account"
	AccountTypeFilterImportAccount         AccountTypeFilter = "Import account"
	AccountTypeFilterReconciliationAccount AccountTypeFilter = "Reconciliation account"
	AccountTypeFilterLoan                  AccountTypeFilter = "Loan"
	AccountTypeFilterDebt                  AccountTypeFilter = "Debt"
	AccountTypeFilterMortgage              AccountTypeFilter = "Mortgage"
)

type AttachmentArray struct {
	Data []AttachmentRead `json:"data" api:"required"`
	Meta Meta             `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentArray) RawJSON() string { return r.JSON.raw }
func (r *AttachmentArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mandatory when the account_role is ccAsset. Can only be monthlyFull or null.
type CreditCardTypeProperty string

const (
	CreditCardTypePropertyMonthlyFull CreditCardTypeProperty = "monthlyFull"
)

// Mandatory when type is liability. Period over which the interest is calculated.
type InterestPeriodProperty string

const (
	InterestPeriodPropertyDaily     InterestPeriodProperty = "daily"
	InterestPeriodPropertyWeekly    InterestPeriodProperty = "weekly"
	InterestPeriodPropertyMonthly   InterestPeriodProperty = "monthly"
	InterestPeriodPropertyQuarterly InterestPeriodProperty = "quarterly"
	InterestPeriodPropertyHalfYear  InterestPeriodProperty = "half-year"
	InterestPeriodPropertyYearly    InterestPeriodProperty = "yearly"
)

// 'credit' indicates somebody owes you the liability. 'debit' Indicates you owe
// this debt yourself. Works only for liabilities.
type LiabilityDirectionProperty string

const (
	LiabilityDirectionPropertyCredit LiabilityDirectionProperty = "credit"
	LiabilityDirectionPropertyDebit  LiabilityDirectionProperty = "debit"
)

// Mandatory when type is liability. Specifies the exact type.
type LiabilityTypeProperty string

const (
	LiabilityTypePropertyLoan     LiabilityTypeProperty = "loan"
	LiabilityTypePropertyDebt     LiabilityTypeProperty = "debt"
	LiabilityTypePropertyMortgage LiabilityTypeProperty = "mortgage"
)

type Meta struct {
	Pagination MetaPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Meta) RawJSON() string { return r.JSON.raw }
func (r *Meta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetaPagination struct {
	Count       int64 `json:"count"`
	CurrentPage int64 `json:"current_page"`
	PerPage     int64 `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int64 `json:"total_pages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		CurrentPage respjson.Field
		PerPage     respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetaPagination) RawJSON() string { return r.JSON.raw }
func (r *MetaPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageLink struct {
	First string `json:"first" format:"uri"`
	Last  string `json:"last" format:"uri"`
	Next  string `json:"next" api:"nullable" format:"uri"`
	Prev  string `json:"prev" api:"nullable" format:"uri"`
	Self  string `json:"self" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		First       respjson.Field
		Last        respjson.Field
		Next        respjson.Field
		Prev        respjson.Field
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageLink) RawJSON() string { return r.JSON.raw }
func (r *PageLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PiggyBankArray struct {
	Data  []PiggyBankRead `json:"data" api:"required"`
	Links PageLink        `json:"links" api:"required"`
	Meta  Meta            `json:"meta" api:"required"`
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
func (r PiggyBankArray) RawJSON() string { return r.JSON.raw }
func (r *PiggyBankArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Can only be one one these account types. import, initial-balance and
// reconciliation cannot be set manually.
type ShortAccountTypeProperty string

const (
	ShortAccountTypePropertyAsset          ShortAccountTypeProperty = "asset"
	ShortAccountTypePropertyExpense        ShortAccountTypeProperty = "expense"
	ShortAccountTypePropertyImport         ShortAccountTypeProperty = "import"
	ShortAccountTypePropertyRevenue        ShortAccountTypeProperty = "revenue"
	ShortAccountTypePropertyCash           ShortAccountTypeProperty = "cash"
	ShortAccountTypePropertyLiability      ShortAccountTypeProperty = "liability"
	ShortAccountTypePropertyLiabilities    ShortAccountTypeProperty = "liabilities"
	ShortAccountTypePropertyInitialBalance ShortAccountTypeProperty = "initial-balance"
	ShortAccountTypePropertyReconciliation ShortAccountTypeProperty = "reconciliation"
)

type TransactionArray struct {
	Data  []TransactionRead `json:"data" api:"required"`
	Links PageLink          `json:"links" api:"required"`
	Meta  Meta              `json:"meta" api:"required"`
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
func (r TransactionArray) RawJSON() string { return r.JSON.raw }
func (r *TransactionArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionTypeFilter string

const (
	TransactionTypeFilterAll            TransactionTypeFilter = "all"
	TransactionTypeFilterWithdrawal     TransactionTypeFilter = "withdrawal"
	TransactionTypeFilterWithdrawals    TransactionTypeFilter = "withdrawals"
	TransactionTypeFilterExpense        TransactionTypeFilter = "expense"
	TransactionTypeFilterDeposit        TransactionTypeFilter = "deposit"
	TransactionTypeFilterDeposits       TransactionTypeFilter = "deposits"
	TransactionTypeFilterIncome         TransactionTypeFilter = "income"
	TransactionTypeFilterTransfer       TransactionTypeFilter = "transfer"
	TransactionTypeFilterTransfers      TransactionTypeFilter = "transfers"
	TransactionTypeFilterOpeningBalance TransactionTypeFilter = "opening_balance"
	TransactionTypeFilterReconciliation TransactionTypeFilter = "reconciliation"
	TransactionTypeFilterSpecial        TransactionTypeFilter = "special"
	TransactionTypeFilterSpecials       TransactionTypeFilter = "specials"
	TransactionTypeFilterDefault        TransactionTypeFilter = "default"
)

type AccountNewParams struct {
	Name string `json:"name" api:"required"`
	// Can only be one one these account types. import, initial-balance and
	// reconciliation cannot be set manually.
	//
	// Any of "asset", "expense", "import", "revenue", "cash", "liability",
	// "liabilities", "initial-balance", "reconciliation".
	Type          ShortAccountTypeProperty `json:"type,omitzero" api:"required"`
	AccountNumber param.Opt[string]        `json:"account_number,omitzero"`
	Bic           param.Opt[string]        `json:"bic,omitzero"`
	Iban          param.Opt[string]        `json:"iban,omitzero"`
	// Mandatory when type is liability. Interest percentage.
	Interest param.Opt[string] `json:"interest,omitzero"`
	// Latitude of the accounts's location, if applicable. Can be used to draw a map.
	Latitude param.Opt[float64] `json:"latitude,omitzero"`
	// Latitude of the accounts's location, if applicable. Can be used to draw a map.
	Longitude param.Opt[float64] `json:"longitude,omitzero"`
	// Mandatory when the account_role is ccAsset. Moment at which CC payment
	// installments are asked for by the bank.
	MonthlyPaymentDate param.Opt[time.Time] `json:"monthly_payment_date,omitzero" format:"date-time"`
	Notes              param.Opt[string]    `json:"notes,omitzero"`
	// Represents the date of the opening balance.
	OpeningBalanceDate param.Opt[time.Time] `json:"opening_balance_date,omitzero" format:"date-time"`
	// Zoom level for the map, if drawn. This to set the box right. Unfortunately this
	// is a proprietary value because each map provider has different zoom levels.
	ZoomLevel param.Opt[int64] `json:"zoom_level,omitzero"`
	// If omitted, defaults to true.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's financial
	// administration's currency.
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's financial
	// administration's currency.
	CurrencyID param.Opt[string] `json:"currency_id,omitzero"`
	// If omitted, defaults to true.
	IncludeNetWorth param.Opt[bool] `json:"include_net_worth,omitzero"`
	// Represents the opening balance, the initial amount this account holds.
	OpeningBalance param.Opt[string] `json:"opening_balance,omitzero"`
	// Order of the account
	Order          param.Opt[int64]  `json:"order,omitzero"`
	VirtualBalance param.Opt[string] `json:"virtual_balance,omitzero"`
	XTraceID       param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Is only mandatory when the type is asset.
	//
	// Any of "defaultAsset", "sharedAsset", "savingAsset", "ccAsset",
	// "cashWalletAsset".
	AccountRole AccountRoleProperty `json:"account_role,omitzero"`
	// Mandatory when the account_role is ccAsset. Can only be monthlyFull or null.
	//
	// Any of "monthlyFull".
	CreditCardType CreditCardTypeProperty `json:"credit_card_type,omitzero"`
	// Mandatory when type is liability. Period over which the interest is calculated.
	//
	// Any of "daily", "weekly", "monthly", "quarterly", "half-year", "yearly".
	InterestPeriod InterestPeriodProperty `json:"interest_period,omitzero"`
	// 'credit' indicates somebody owes you the liability. 'debit' Indicates you owe
	// this debt yourself. Works only for liabilities.
	//
	// Any of "credit", "debit".
	LiabilityDirection LiabilityDirectionProperty `json:"liability_direction,omitzero"`
	// Mandatory when type is liability. Specifies the exact type.
	//
	// Any of "loan", "debt", "mortgage".
	LiabilityType LiabilityTypeProperty `json:"liability_type,omitzero"`
	paramObj
}

func (r AccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetParams struct {
	// A date formatted YYYY-MM-DD. When added to the request, Firefly III will show
	// the account's balance on that day.
	Date param.Opt[time.Time] `query:"date,omitzero" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD. Must be after "start". Can not be the same as
	// "start". May be omitted.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD. May be omitted.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AccountGetParams]'s query parameters as `url.Values`.
func (r AccountGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountUpdateParams struct {
	Name          string            `json:"name" api:"required"`
	Type          any               `json:"type,omitzero" api:"required"`
	AccountNumber param.Opt[string] `json:"account_number,omitzero"`
	Bic           param.Opt[string] `json:"bic,omitzero"`
	Iban          param.Opt[string] `json:"iban,omitzero"`
	// Mandatory when type is liability. Interest percentage.
	Interest param.Opt[string] `json:"interest,omitzero"`
	// Latitude of the account's location, if applicable. Can be used to draw a map. If
	// omitted, the existing location will be kept. If submitted as NULL, the current
	// location will be removed.
	Latitude param.Opt[float64] `json:"latitude,omitzero"`
	// Latitude of the account's location, if applicable. Can be used to draw a map. If
	// omitted, the existing location will be kept. If submitted as NULL, the current
	// location will be removed.
	Longitude param.Opt[float64] `json:"longitude,omitzero"`
	// Mandatory when the account_role is ccAsset. Moment at which CC payment
	// installments are asked for by the bank.
	MonthlyPaymentDate param.Opt[time.Time] `json:"monthly_payment_date,omitzero" format:"date-time"`
	Notes              param.Opt[string]    `json:"notes,omitzero"`
	OpeningBalanceDate param.Opt[time.Time] `json:"opening_balance_date,omitzero" format:"date-time"`
	// Zoom level for the map, if drawn. This to set the box right. Unfortunately this
	// is a proprietary value because each map provider has different zoom levels. If
	// omitted, the existing location will be kept. If submitted as NULL, the current
	// location will be removed.
	ZoomLevel param.Opt[int64] `json:"zoom_level,omitzero"`
	// If omitted, defaults to true.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's financial
	// administration's currency.
	CurrencyCode param.Opt[string] `json:"currency_code,omitzero"`
	// Use either currency_id or currency_code. Defaults to the user's financial
	// administration's currency.
	CurrencyID param.Opt[string] `json:"currency_id,omitzero"`
	// If omitted, defaults to true.
	IncludeNetWorth param.Opt[bool]   `json:"include_net_worth,omitzero"`
	OpeningBalance  param.Opt[string] `json:"opening_balance,omitzero"`
	// Order of the account
	Order          param.Opt[int64]  `json:"order,omitzero"`
	VirtualBalance param.Opt[string] `json:"virtual_balance,omitzero"`
	XTraceID       param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Is only mandatory when the type is asset.
	//
	// Any of "defaultAsset", "sharedAsset", "savingAsset", "ccAsset",
	// "cashWalletAsset".
	AccountRole AccountRoleProperty `json:"account_role,omitzero"`
	// Mandatory when the account_role is ccAsset. Can only be monthlyFull or null.
	//
	// Any of "monthlyFull".
	CreditCardType CreditCardTypeProperty `json:"credit_card_type,omitzero"`
	// Mandatory when type is liability. Period over which the interest is calculated.
	//
	// Any of "daily", "weekly", "monthly", "quarterly", "half-year", "yearly".
	InterestPeriod InterestPeriodProperty `json:"interest_period,omitzero"`
	// Mandatory when type is liability. Specifies the exact type.
	//
	// Any of "loan", "debt", "mortgage".
	LiabilityType LiabilityTypeProperty `json:"liability_type,omitzero"`
	paramObj
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListParams struct {
	// A date formatted YYYY-MM-DD. When added to the request, Firefly III will show
	// the account's balance on that day.
	Date param.Opt[time.Time] `query:"date,omitzero" format:"date" json:"-"`
	// A date formatted YYYY-MM-DD. Must be after "start". Can not be the same as
	// "start". May be omitted.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// A date formatted YYYY-MM-DD. May be omitted.
	Start    param.Opt[time.Time] `query:"start,omitzero" format:"date" json:"-"`
	XTraceID param.Opt[string]    `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Any of "all", "asset", "cash", "expense", "revenue", "special", "hidden",
	// "liability", "liabilities", "Default account", "Cash account", "Asset account",
	// "Expense account", "Revenue account", "Initial balance account", "Beneficiary
	// account", "Import account", "Reconciliation account", "Loan", "Debt",
	// "Mortgage".
	Type AccountTypeFilter `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type AccountListAttachmentsParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListAttachmentsParams]'s query parameters as
// `url.Values`.
func (r AccountListAttachmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountListPiggyBanksParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListPiggyBanksParams]'s query parameters as
// `url.Values`.
func (r AccountListPiggyBanksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountListTransactionsParams struct {
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

// URLQuery serializes [AccountListTransactionsParams]'s query parameters as
// `url.Values`.
func (r AccountListTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
