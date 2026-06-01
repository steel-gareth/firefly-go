// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package emceesprodtesting5

import (
	"context"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the emcees-prod-testing-5 API. You should not instantiate this
// client directly, and instead use the [NewClient] method instead.
type Client struct {
	options []option.RequestOption
	// Auto-complete endpoints show basic information about Firefly III models, like
	// the name and maybe some amounts. They all support a search query and can be used
	// to autocomplete data in forms. Autocomplete return values always have a
	// &quot;name&quot;-field.
	Autocomplete AutocompleteService
	Chart        ChartService
	// The &quot;data&quot;-endpoints manage generic Firefly III and user-specific
	// data.
	Data    DataService
	Insight InsightService
	// Endpoints that deliver all of the user&#039;s asset, expense and other accounts
	// (and the metadata) together with related transactions, piggy banks and other
	// objects. Also delivers endpoints for CRUD operations for accounts.
	Accounts AccountService
	// Endpoints to manage the attachments of the authenticated user, including up- and
	// downloading of the files.
	Attachments AttachmentService
	// Endpoints to manage the total available amount that the user has made available
	// to themselves. Used in periodic budgeting.
	AvailableBudgets AvailableBudgetService
	// Endpoints to manage a user&#039;s bills and all related objects.
	Bills BillService
	// Endpoints to manage a user&#039;s budgets and get info on the related objects,
	// like limits.
	Budgets BudgetService
	// Endpoints to manage a user&#039;s categories and get information on transactions
	// and other related objects.
	Categories CategoryService
	// All currency exchange rates.
	ExchangeRates ExchangeRateService
	// Endpoints to manage links between transactions, and manage the type of links
	// available.
	LinkTypes LinkTypeService
	// Endpoints to manage links between transactions, and manage the type of links
	// available.
	TransactionLinks TransactionLinkService
	// Endpoints to control and manage all of the user&#039;s object groups. Can only
	// be created in conjunction with another object (for example a piggy bank) and
	// will auto-delete when no other items are linked to it.
	ObjectGroups ObjectGroupService
	// Endpoints to control and manage all of the user&#039;s piggy banks and related
	// objects and information.
	PiggyBanks PiggyBankService
	// Use these endpoints to manage the user&#039;s recurring transactions, trigger
	// the creation of transactions and manage the settings.
	Recurrences RecurrenceService
	// Manage all of the user&#039;s groups of rules and trigger the execution of
	// entire groups.
	RuleGroups RuleGroupService
	// These endpoints can be used to manage all of the user&#039;s rules. Also
	// includes triggers to execute or test rules and individual triggers.
	Rules RuleService
	// This endpoint manages all of the user&#039;s tags.
	Tags TagService
	// Endpoints to manage the currencies in Firefly III. Depending on the user&#039;s
	// role you can also disable and enable them, or add new ones.
	Currencies CurrencyService
	// The most-used endpoints in Firefly III, these endpoints are used to manage the
	// user&#039;s transactions.
	TransactionJournals TransactionJournalService
	// The most-used endpoints in Firefly III, these endpoints are used to manage the
	// user&#039;s transactions.
	Transactions TransactionService
	// User groups are the objects around which &quot;financial administrations&quot;
	// are built.
	UserGroups UserGroupService
	// Endpoints that allow you to search through the user&#039;s financial data.
	// Different from the autocomplete endpoints, the search accepts more advanced
	// arguments.
	Search SearchService
	// These endpoints deliver summaries, like sums, lists of numbers and other
	// processed information. Mainly used for the main dashboard and pretty specific
	// for Firefly III itself.
	Summary SummaryService
	// These endpoints deliver general system information, version- and meta
	// information.
	About AboutService
	// These endpoints deliver general system information, version- and meta
	// information.
	Batch BatchService
	// These endpoints allow you to manage and update the Firefly III configuration.
	// You need to have the &quot;owner&quot; role to update configuration.
	Configuration ConfigurationService
	// These endpoints deliver general system information, version- and meta
	// information.
	Cron CronService
	// Use these endpoints to manage the users registered within Firefly III. You need
	// to have the &quot;owner&quot; role to access these endpoints.
	Users UserService
	// These endpoints can be used to manage the user&#039;s preferences, including
	// some hidden ones.
	Preferences PreferenceService
	// These endpoints can be used to manage the user&#039;s webhooks and triggers them
	// if necessary.
	Webhooks WebhookService
}

// DefaultClientOptions read from the environment
// (EMCEES_PROD_TESTING_5_BEARER_TOKEN, EMCEES_PROD_TESTING_5_BASE_URL). This
// should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("EMCEES_PROD_TESTING_5_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("EMCEES_PROD_TESTING_5_BEARER_TOKEN"); ok {
		defaults = append(defaults, option.WithBearerToken(o))
	}
	if o, ok := os.LookupEnv("EMCEES_PROD_TESTING_5_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (EMCEES_PROD_TESTING_5_BEARER_TOKEN,
// EMCEES_PROD_TESTING_5_BASE_URL). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{options: opts}

	r.Autocomplete = NewAutocompleteService(opts...)
	r.Chart = NewChartService(opts...)
	r.Data = NewDataService(opts...)
	r.Insight = NewInsightService(opts...)
	r.Accounts = NewAccountService(opts...)
	r.Attachments = NewAttachmentService(opts...)
	r.AvailableBudgets = NewAvailableBudgetService(opts...)
	r.Bills = NewBillService(opts...)
	r.Budgets = NewBudgetService(opts...)
	r.Categories = NewCategoryService(opts...)
	r.ExchangeRates = NewExchangeRateService(opts...)
	r.LinkTypes = NewLinkTypeService(opts...)
	r.TransactionLinks = NewTransactionLinkService(opts...)
	r.ObjectGroups = NewObjectGroupService(opts...)
	r.PiggyBanks = NewPiggyBankService(opts...)
	r.Recurrences = NewRecurrenceService(opts...)
	r.RuleGroups = NewRuleGroupService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Tags = NewTagService(opts...)
	r.Currencies = NewCurrencyService(opts...)
	r.TransactionJournals = NewTransactionJournalService(opts...)
	r.Transactions = NewTransactionService(opts...)
	r.UserGroups = NewUserGroupService(opts...)
	r.Search = NewSearchService(opts...)
	r.Summary = NewSummaryService(opts...)
	r.About = NewAboutService(opts...)
	r.Batch = NewBatchService(opts...)
	r.Configuration = NewConfigurationService(opts...)
	r.Cron = NewCronService(opts...)
	r.Users = NewUserService(opts...)
	r.Preferences = NewPreferenceService(opts...)
	r.Webhooks = NewWebhookService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
