// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apijson"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/apiquery"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/respjson"
)

// Auto-complete endpoints show basic information about Firefly III models, like
// the name and maybe some amounts. They all support a search query and can be used
// to autocomplete data in forms. Autocomplete return values always have a
// &quot;name&quot;-field.
//
// AutocompleteService contains methods and other services that help with
// interacting with the emcees-prod-testing-5 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutocompleteService] method instead.
type AutocompleteService struct {
	options []option.RequestOption
}

// NewAutocompleteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutocompleteService(opts ...option.RequestOption) (r AutocompleteService) {
	r = AutocompleteService{}
	r.options = opts
	return
}

// Returns all accounts of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListAccounts(ctx context.Context, params AutocompleteListAccountsParams, opts ...option.RequestOption) (res *[]AutocompleteListAccountsResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all bills of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListBills(ctx context.Context, params AutocompleteListBillsParams, opts ...option.RequestOption) (res *[]AutocompleteBill, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/bills"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all budgets of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListBudgets(ctx context.Context, params AutocompleteListBudgetsParams, opts ...option.RequestOption) (res *[]AutocompleteListBudgetsResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all categories of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListCategories(ctx context.Context, params AutocompleteListCategoriesParams, opts ...option.RequestOption) (res *[]AutocompleteListCategoriesResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all currencies of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListCurrencies(ctx context.Context, params AutocompleteListCurrenciesParams, opts ...option.RequestOption) (res *[]AutocompleteListCurrenciesResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/currencies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all currencies of the user returned in a basic auto-complete array. This
// endpoint is DEPRECATED and I suggest you DO NOT use it.
func (r *AutocompleteService) ListCurrenciesWithCode(ctx context.Context, params AutocompleteListCurrenciesWithCodeParams, opts ...option.RequestOption) (res *[]AutocompleteListCurrenciesWithCodeResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/currencies-with-code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all object groups of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListObjectGroups(ctx context.Context, params AutocompleteListObjectGroupsParams, opts ...option.RequestOption) (res *[]AutocompleteListObjectGroupsResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/object-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all piggy banks of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListPiggyBanks(ctx context.Context, params AutocompleteListPiggyBanksParams, opts ...option.RequestOption) (res *[]AutocompleteListPiggyBanksResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/piggy-banks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all piggy banks of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListPiggyBanksWithBalance(ctx context.Context, params AutocompleteListPiggyBanksWithBalanceParams, opts ...option.RequestOption) (res *[]AutocompleteListPiggyBanksWithBalanceResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/piggy-banks-with-balance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all recurring transactions of the user returned in a basic auto-complete
// array.
func (r *AutocompleteService) ListRecurringTransactions(ctx context.Context, params AutocompleteListRecurringTransactionsParams, opts ...option.RequestOption) (res *[]AutocompleteListRecurringTransactionsResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all rule groups of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListRuleGroups(ctx context.Context, params AutocompleteListRuleGroupsParams, opts ...option.RequestOption) (res *[]AutocompleteListRuleGroupsResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/rule-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all rules of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListRules(ctx context.Context, params AutocompleteListRulesParams, opts ...option.RequestOption) (res *[]AutocompleteListRulesResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all subscriptions of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListSubscriptions(ctx context.Context, params AutocompleteListSubscriptionsParams, opts ...option.RequestOption) (res *[]AutocompleteBill, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all tags of the user returned in a basic auto-complete array.
func (r *AutocompleteService) ListTags(ctx context.Context, params AutocompleteListTagsParams, opts ...option.RequestOption) (res *[]AutocompleteListTagsResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all transaction types returned in a basic auto-complete array. English
// only.
func (r *AutocompleteService) ListTransactionTypes(ctx context.Context, params AutocompleteListTransactionTypesParams, opts ...option.RequestOption) (res *[]AutocompleteListTransactionTypesResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/transaction-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all transaction descriptions of the user returned in a basic
// auto-complete array.
func (r *AutocompleteService) ListTransactions(ctx context.Context, params AutocompleteListTransactionsParams, opts ...option.RequestOption) (res *[]AutocompleteListTransactionsResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns all transactions, complemented with their ID, of the user returned in a
// basic auto-complete array. This endpoint is DEPRECATED and I suggest you DO NOT
// use it.
func (r *AutocompleteService) ListTransactionsWithID(ctx context.Context, params AutocompleteListTransactionsWithIDParams, opts ...option.RequestOption) (res *[]AutocompleteListTransactionsWithIDResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/autocomplete/transactions-with-id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type AutocompleteBill struct {
	ID string `json:"id" api:"required"`
	// Name of the bill found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Is the bill active or not?
	Active bool `json:"active"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Active      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteBill) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteBill) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListAccountsResponse struct {
	ID string `json:"id" api:"required"`
	// Currency code for the currency used by this account. If the user prefers amounts
	// converted to their primary currency, this primary currency is used instead.
	CurrencyCode string `json:"currency_code" api:"required"`
	// Number of decimal places for the currency used by this account. If the user
	// prefers amounts converted to their primary currency, this primary currency is
	// used instead.
	CurrencyDecimalPlaces int64 `json:"currency_decimal_places" api:"required"`
	// ID for the currency used by this account. If the user prefers amounts converted
	// to their primary currency, this primary currency is used instead.
	CurrencyID string `json:"currency_id" api:"required"`
	// Currency name for the currency used by this account. If the user prefers amounts
	// converted to their primary currency, this primary currency is used instead.
	CurrencyName string `json:"currency_name" api:"required"`
	// Currency symbol for the currency used by this account. If the user prefers
	// amounts converted to their primary currency, this primary currency is used
	// instead.
	CurrencySymbol string `json:"currency_symbol" api:"required"`
	// Name of the account found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Asset accounts and liabilities have a second field with the given date's account
	// balance in the account currency or primary currency.
	NameWithBalance string `json:"name_with_balance" api:"required"`
	// Account type of the account found by the auto-complete search.
	Type string `json:"type" api:"required"`
	// Code for the currency used by this account. Even if "convertToPrimary" is on,
	// the account currency code is displayed here.
	AccountCurrencyCode string `json:"account_currency_code"`
	// Number of decimal places for the currency used by this account. Even if
	// "convertToPrimary" is on, the account currency code is displayed here.
	AccountCurrencyDecimalPlaces int64 `json:"account_currency_decimal_places"`
	// ID for the currency used by this account. Even if "convertToPrimary" is on, the
	// account currency ID is displayed here.
	AccountCurrencyID string `json:"account_currency_id"`
	// Name for the currency used by this account. Even if "convertToPrimary" is on,
	// the account currency name is displayed here.
	AccountCurrencyName string `json:"account_currency_name"`
	// Code for the currency used by this account. Even if "convertToPrimary" is on,
	// the account currency code is displayed here.
	AccountCurrencySymbol string `json:"account_currency_symbol"`
	// Is the bill active or not?
	Active bool `json:"active"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		CurrencyCode                 respjson.Field
		CurrencyDecimalPlaces        respjson.Field
		CurrencyID                   respjson.Field
		CurrencyName                 respjson.Field
		CurrencySymbol               respjson.Field
		Name                         respjson.Field
		NameWithBalance              respjson.Field
		Type                         respjson.Field
		AccountCurrencyCode          respjson.Field
		AccountCurrencyDecimalPlaces respjson.Field
		AccountCurrencyID            respjson.Field
		AccountCurrencyName          respjson.Field
		AccountCurrencySymbol        respjson.Field
		Active                       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListAccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListAccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListBudgetsResponse struct {
	ID string `json:"id" api:"required"`
	// Name of the budget found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Is the budget active or not?
	Active bool `json:"active"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Active      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListBudgetsResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListBudgetsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListCategoriesResponse struct {
	ID string `json:"id" api:"required"`
	// Name of the category found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListCategoriesResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListCategoriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListCurrenciesResponse struct {
	ID string `json:"id" api:"required"`
	// Currency code.
	Code          string `json:"code" api:"required"`
	DecimalPlaces int64  `json:"decimal_places" api:"required"`
	// Currency name.
	Name   string `json:"name" api:"required"`
	Symbol string `json:"symbol" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Code          respjson.Field
		DecimalPlaces respjson.Field
		Name          respjson.Field
		Symbol        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListCurrenciesResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListCurrenciesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListCurrenciesWithCodeResponse struct {
	ID string `json:"id" api:"required"`
	// Currency code.
	Code          string `json:"code" api:"required"`
	DecimalPlaces int64  `json:"decimal_places" api:"required"`
	// Currency name with the code between brackets.
	Name   string `json:"name" api:"required"`
	Symbol string `json:"symbol" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Code          respjson.Field
		DecimalPlaces respjson.Field
		Name          respjson.Field
		Symbol        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListCurrenciesWithCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListCurrenciesWithCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListObjectGroupsResponse struct {
	ID string `json:"id" api:"required"`
	// Title of the object group found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Title of the object group found by an auto-complete search.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListObjectGroupsResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListObjectGroupsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListPiggyBanksResponse struct {
	ID string `json:"id" api:"required"`
	// Name of the piggy bank found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Currency code for this piggy bank. This will always be the currency of the piggy
	// bank, never the user's primary currency.
	CurrencyCode string `json:"currency_code"`
	// Number of decimal places for the currency used by this piggy bank. This will
	// always be the currency of the piggy bank, never the user's primary currency.
	CurrencyDecimalPlaces int64 `json:"currency_decimal_places"`
	// Currency ID for this piggy bank. This will always be the currency of the piggy
	// bank, never the user's primary currency.
	CurrencyID string `json:"currency_id"`
	// Currency name for the currency used by this piggy bank. This will always be the
	// currency of the piggy bank, never the user's primary currency.
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupID string `json:"object_group_id" api:"nullable"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle string `json:"object_group_title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Name                  respjson.Field
		CurrencyCode          respjson.Field
		CurrencyDecimalPlaces respjson.Field
		CurrencyID            respjson.Field
		CurrencyName          respjson.Field
		CurrencySymbol        respjson.Field
		ObjectGroupID         respjson.Field
		ObjectGroupTitle      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListPiggyBanksResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListPiggyBanksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListPiggyBanksWithBalanceResponse struct {
	ID string `json:"id" api:"required"`
	// Name of the piggy bank found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Currency code for the currency used by this piggy bank. This will always be the
	// piggy bank's currency, never the primary currency.
	CurrencyCode string `json:"currency_code"`
	// Currency decimal places for the currency used by this piggy bank. This will
	// always be the piggy bank's currency, never the primary currency.
	CurrencyDecimalPlaces int64 `json:"currency_decimal_places"`
	// Currency ID for the currency used by this piggy bank. This will always be the
	// piggy bank's currency, never the primary currency.
	CurrencyID string `json:"currency_id"`
	// Currency symbol for the currency used by this piggy bank. This will always be
	// the piggy bank's currency, never the primary currency.
	CurrencySymbol string `json:"currency_symbol"`
	// Name of the piggy bank found by an auto-complete search, including the currently
	// saved amount and the target amount.
	NameWithBalance string `json:"name_with_balance"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupID string `json:"object_group_id" api:"nullable"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle string `json:"object_group_title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Name                  respjson.Field
		CurrencyCode          respjson.Field
		CurrencyDecimalPlaces respjson.Field
		CurrencyID            respjson.Field
		CurrencySymbol        respjson.Field
		NameWithBalance       respjson.Field
		ObjectGroupID         respjson.Field
		ObjectGroupTitle      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListPiggyBanksWithBalanceResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListPiggyBanksWithBalanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListRecurringTransactionsResponse struct {
	ID string `json:"id" api:"required"`
	// Name of the recurrence found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Is the recurring transaction active or not?
	Active bool `json:"active"`
	// Description of the recurrence found by auto-complete.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Active      respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListRecurringTransactionsResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListRecurringTransactionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListRuleGroupsResponse struct {
	ID string `json:"id" api:"required"`
	// Name of the rule group found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Is the bill active or not?
	Active bool `json:"active"`
	// Description of the rule group found by auto-complete.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Active      respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListRuleGroupsResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListRuleGroupsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListRulesResponse struct {
	ID string `json:"id" api:"required"`
	// Name of the rule found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Is the bill active or not?
	Active bool `json:"active"`
	// Description of the rule found by auto-complete.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Active      respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListRulesResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListRulesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListTagsResponse struct {
	ID string `json:"id" api:"required"`
	// Name of the tag found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Name of the tag found by an auto-complete search.
	Tag string `json:"tag" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListTagsResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListTagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListTransactionTypesResponse struct {
	ID string `json:"id" api:"required"`
	// Type of the object found by an auto-complete search.
	Name string `json:"name" api:"required"`
	// Name of the object found by an auto-complete search.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListTransactionTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListTransactionTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListTransactionsResponse struct {
	// The ID of a transaction journal (basically a single split).
	ID string `json:"id" api:"required"`
	// Transaction description
	Description string `json:"description" api:"required"`
	// Transaction description
	Name string `json:"name" api:"required"`
	// The ID of the underlying transaction group.
	TransactionGroupID string `json:"transaction_group_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Description        respjson.Field
		Name               respjson.Field
		TransactionGroupID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListTransactionsResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListTransactionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListTransactionsWithIDResponse struct {
	// The ID of a transaction journal (basically a single split).
	ID string `json:"id" api:"required"`
	// Transaction description with ID in the name.
	Description string `json:"description" api:"required"`
	// Transaction description with ID in the name.
	Name string `json:"name" api:"required"`
	// The ID of the underlying transaction group.
	TransactionGroupID string `json:"transaction_group_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Description        respjson.Field
		Name               respjson.Field
		TransactionGroupID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteListTransactionsWithIDResponse) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteListTransactionsWithIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteListAccountsParams struct {
	// If the account is an asset account or a liability, the autocomplete will also
	// return the balance of the account on this date.
	Date param.Opt[string] `query:"date,omitzero" json:"-"`
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	// Optional filter on the account type(s) used in the autocomplete.
	Types []AccountTypeFilter `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListAccountsParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListAccountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListBillsParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListBillsParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListBillsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListBudgetsParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListBudgetsParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListBudgetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListCategoriesParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListCategoriesParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListCategoriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListCurrenciesParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListCurrenciesParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListCurrenciesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListCurrenciesWithCodeParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListCurrenciesWithCodeParams]'s query
// parameters as `url.Values`.
func (r AutocompleteListCurrenciesWithCodeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListObjectGroupsParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListObjectGroupsParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListObjectGroupsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListPiggyBanksParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListPiggyBanksParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListPiggyBanksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListPiggyBanksWithBalanceParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListPiggyBanksWithBalanceParams]'s query
// parameters as `url.Values`.
func (r AutocompleteListPiggyBanksWithBalanceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListRecurringTransactionsParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListRecurringTransactionsParams]'s query
// parameters as `url.Values`.
func (r AutocompleteListRecurringTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListRuleGroupsParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListRuleGroupsParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListRuleGroupsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListRulesParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListRulesParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListRulesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListSubscriptionsParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListSubscriptionsParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListSubscriptionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListTagsParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListTagsParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListTagsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListTransactionTypesParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListTransactionTypesParams]'s query parameters
// as `url.Values`.
func (r AutocompleteListTransactionTypesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListTransactionsParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListTransactionsParams]'s query parameters as
// `url.Values`.
func (r AutocompleteListTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AutocompleteListTransactionsWithIDParams struct {
	// The number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The autocomplete search query.
	Query    param.Opt[string] `query:"query,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteListTransactionsWithIDParams]'s query
// parameters as `url.Values`.
func (r AutocompleteListTransactionsWithIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
