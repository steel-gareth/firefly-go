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
	shimjson "github.com/stainless-sdks/emcees-prod-testing-5-go/internal/encoding/json"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/respjson"
)

// All currency exchange rates.
//
// ExchangeRateService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExchangeRateService] method instead.
type ExchangeRateService struct {
	options []option.RequestOption
}

// NewExchangeRateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExchangeRateService(opts ...option.RequestOption) (r ExchangeRateService) {
	r = ExchangeRateService{}
	r.options = opts
	return
}

// Stores a new exchange rate. The data required can be submitted as a JSON body or
// as a list of parameters.
func (r *ExchangeRateService) New(ctx context.Context, params ExchangeRateNewParams, opts ...option.RequestOption) (res *CurrencyExchangeRateSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/exchange-rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List a single specific exchange rate by its ID.
func (r *ExchangeRateService) Get(ctx context.Context, id string, params ExchangeRateGetParams, opts ...option.RequestOption) (res *CurrencyExchangeRateSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/exchange-rates/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Used to update a single currency exchange rate by its ID. Including the from/to
// currency is optional.
func (r *ExchangeRateService) Update(ctx context.Context, id string, params ExchangeRateUpdateParams, opts ...option.RequestOption) (res *CurrencyExchangeRateSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/exchange-rates/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List exchange rates that Firefly III knows.
func (r *ExchangeRateService) List(ctx context.Context, params ExchangeRateListParams, opts ...option.RequestOption) (res *CurrencyExchangeRateArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/exchange-rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a specific currency exchange rate by its internal ID.
func (r *ExchangeRateService) Delete(ctx context.Context, id string, body ExchangeRateDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/exchange-rates/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Stores a new set of exchange rates for this pair. The date is variable, and the
// data required can be submitted as a JSON body.
func (r *ExchangeRateService) NewByCurrencies(ctx context.Context, to string, params ExchangeRateNewByCurrenciesParams, opts ...option.RequestOption) (res *CurrencyExchangeRateArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.From == "" {
		err = errors.New("missing required from parameter")
		return nil, err
	}
	if to == "" {
		err = errors.New("missing required to parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/exchange-rates/by-currencies/%s/%s", url.PathEscape(params.From), url.PathEscape(to))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Stores a new set of exchange rates. The date is fixed (in the URL parameter) and
// the data required can be submitted as a JSON body.
func (r *ExchangeRateService) NewByDate(ctx context.Context, date string, params ExchangeRateNewByDateParams, opts ...option.RequestOption) (res *CurrencyExchangeRateArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if date == "" {
		err = errors.New("missing required date parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/exchange-rates/by-date/%s", url.PathEscape(date))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Deletes ALL currency exchange rates from 'from' to 'to'. It's important to know
// that the reverse exchange rates (from 'to' to 'from') will not be deleted and
// Firefly III will still be able to infer the correct exchange rate from the
// reverse one.
func (r *ExchangeRateService) DeleteAllByCurrencies(ctx context.Context, to string, params ExchangeRateDeleteAllByCurrenciesParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.From == "" {
		err = errors.New("missing required from parameter")
		return err
	}
	if to == "" {
		err = errors.New("missing required to parameter")
		return err
	}
	path := fmt.Sprintf("v1/exchange-rates/%s/%s", url.PathEscape(params.From), url.PathEscape(to))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delete the currency exchange rate from 'from' to 'to' on the specified date.
// It's important to know that the reverse exchange rate (from 'to' to 'from') will
// not be deleted and Firefly III will still be able to infer the correct exchange
// rate from the reverse one.
func (r *ExchangeRateService) DeleteByDate(ctx context.Context, date string, params ExchangeRateDeleteByDateParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.From == "" {
		err = errors.New("missing required from parameter")
		return err
	}
	if params.To == "" {
		err = errors.New("missing required to parameter")
		return err
	}
	if date == "" {
		err = errors.New("missing required date parameter")
		return err
	}
	path := fmt.Sprintf("v1/exchange-rates/%s/%s/%s", url.PathEscape(params.From), url.PathEscape(params.To), url.PathEscape(date))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// List all exchange rates from/to the mentioned currencies.
func (r *ExchangeRateService) ListByCurrencies(ctx context.Context, to string, params ExchangeRateListByCurrenciesParams, opts ...option.RequestOption) (res *CurrencyExchangeRateArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.From == "" {
		err = errors.New("missing required from parameter")
		return nil, err
	}
	if to == "" {
		err = errors.New("missing required to parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/exchange-rates/%s/%s", url.PathEscape(params.From), url.PathEscape(to))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List the exchange rate for the from and to-currency on the requested date.
func (r *ExchangeRateService) GetByDate(ctx context.Context, date string, params ExchangeRateGetByDateParams, opts ...option.RequestOption) (res *CurrencyExchangeRateArray, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.From == "" {
		err = errors.New("missing required from parameter")
		return nil, err
	}
	if params.To == "" {
		err = errors.New("missing required to parameter")
		return nil, err
	}
	if date == "" {
		err = errors.New("missing required date parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/exchange-rates/%s/%s/%s", url.PathEscape(params.From), url.PathEscape(params.To), url.PathEscape(date))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Used to update a single currency exchange rate by its currency codes and date
func (r *ExchangeRateService) UpdateByDate(ctx context.Context, date string, params ExchangeRateUpdateByDateParams, opts ...option.RequestOption) (res *CurrencyExchangeRateSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.From == "" {
		err = errors.New("missing required from parameter")
		return nil, err
	}
	if params.To == "" {
		err = errors.New("missing required to parameter")
		return nil, err
	}
	if date == "" {
		err = errors.New("missing required date parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/exchange-rates/%s/%s/%s", url.PathEscape(params.From), url.PathEscape(params.To), url.PathEscape(date))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type CurrencyExchangeRateArray struct {
	Data  []CurrencyExchangeRateRead `json:"data" api:"required"`
	Links PageLink                   `json:"links" api:"required"`
	Meta  Meta                       `json:"meta" api:"required"`
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
func (r CurrencyExchangeRateArray) RawJSON() string { return r.JSON.raw }
func (r *CurrencyExchangeRateArray) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyExchangeRateRead struct {
	ID         string                             `json:"id"`
	Attributes CurrencyExchangeRateReadAttributes `json:"attributes"`
	Links      ObjectLink                         `json:"links"`
	// Immutable value
	Type string `json:"type"`
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
func (r CurrencyExchangeRateRead) RawJSON() string { return r.JSON.raw }
func (r *CurrencyExchangeRateRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyExchangeRateReadAttributes struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Date and time of the exchange rate.
	Date time.Time `json:"date" format:"date-time"`
	// Base currency code for this exchange rate entry.
	FromCurrencyCode string `json:"from_currency_code"`
	// Base currency decimal places for this exchange rate entry.
	FromCurrencyDecimalPlaces int64 `json:"from_currency_decimal_places"`
	// Base currency ID for this exchange rate entry.
	FromCurrencyID string `json:"from_currency_id"`
	// Base currency name for this exchange rate entry.
	FromCurrencyName string `json:"from_currency_name"`
	// Base currency symbol for this exchange rate entry.
	FromCurrencySymbol string `json:"from_currency_symbol"`
	// The actual exchange rate. How many 'to' currency will you get for 1 'from'
	// currency?
	Rate string `json:"rate"`
	// Destination currency code for this exchange rate entry.
	ToCurrencyCode string `json:"to_currency_code"`
	// Destination currency decimal places for this exchange rate entry.
	ToCurrencyDecimalPlaces int64 `json:"to_currency_decimal_places"`
	// Destination currency ID for this exchange rate entry.
	ToCurrencyID string `json:"to_currency_id"`
	// Destination currency name for this exchange rate entry.
	ToCurrencyName string `json:"to_currency_name"`
	// Destination currency symbol for this exchange rate entry.
	ToCurrencySymbol string    `json:"to_currency_symbol"`
	UpdatedAt        time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt                 respjson.Field
		Date                      respjson.Field
		FromCurrencyCode          respjson.Field
		FromCurrencyDecimalPlaces respjson.Field
		FromCurrencyID            respjson.Field
		FromCurrencyName          respjson.Field
		FromCurrencySymbol        respjson.Field
		Rate                      respjson.Field
		ToCurrencyCode            respjson.Field
		ToCurrencyDecimalPlaces   respjson.Field
		ToCurrencyID              respjson.Field
		ToCurrencyName            respjson.Field
		ToCurrencySymbol          respjson.Field
		UpdatedAt                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrencyExchangeRateReadAttributes) RawJSON() string { return r.JSON.raw }
func (r *CurrencyExchangeRateReadAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CurrencyExchangeRateSingle struct {
	Data CurrencyExchangeRateRead `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CurrencyExchangeRateSingle) RawJSON() string { return r.JSON.raw }
func (r *CurrencyExchangeRateSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExchangeRateNewParams struct {
	// The date to which the exchange rate is applicable.
	Date time.Time `json:"date" api:"required" format:"date"`
	// The base currency code.
	From  string `json:"from" api:"required"`
	Rates any    `json:"rates,omitzero" api:"required"`
	// The destination currency code.
	To string `json:"to" api:"required"`
	// The exchange rate from the base currency to the destination currency.
	Rate     param.Opt[string] `json:"rate,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExchangeRateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExchangeRateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExchangeRateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExchangeRateGetParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ExchangeRateGetParams]'s query parameters as `url.Values`.
func (r ExchangeRateGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExchangeRateUpdateParams struct {
	// The date to which the exchange rate is applicable.
	Date time.Time `json:"date" api:"required" format:"date"`
	// The exchange rate from the base currency to the destination currency.
	Rate string `json:"rate" api:"required"`
	// The base currency code.
	From param.Opt[string] `json:"from,omitzero"`
	// The destination currency code.
	To       param.Opt[string] `json:"to,omitzero"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExchangeRateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExchangeRateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExchangeRateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExchangeRateListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ExchangeRateListParams]'s query parameters as `url.Values`.
func (r ExchangeRateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExchangeRateDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ExchangeRateNewByCurrenciesParams struct {
	From string `path:"from" api:"required" json:"-"`
	// The actual entries for this data set. They 'key' value is the date in
	// YYYY-MM-DD. The value is the exchange rate.
	Body     map[string]string
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExchangeRateNewByCurrenciesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ExchangeRateNewByCurrenciesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExchangeRateNewByDateParams struct {
	Date any `json:"date,omitzero" api:"required"`
	// The 'from'-currency
	From string `json:"from" api:"required"`
	// The actual entries for this data set. They 'key' value is 'to' currency. The
	// value is the exchange rate.
	Rates    map[string]string `json:"rates,omitzero" api:"required"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExchangeRateNewByDateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExchangeRateNewByDateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExchangeRateNewByDateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExchangeRateDeleteAllByCurrenciesParams struct {
	From     string            `path:"from" api:"required" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ExchangeRateDeleteByDateParams struct {
	From     string            `path:"from" api:"required" json:"-"`
	To       string            `path:"to" api:"required" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ExchangeRateListByCurrenciesParams struct {
	From string `path:"from" api:"required" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ExchangeRateListByCurrenciesParams]'s query parameters as
// `url.Values`.
func (r ExchangeRateListByCurrenciesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExchangeRateGetByDateParams struct {
	From string `path:"from" api:"required" json:"-"`
	To   string `path:"to" api:"required" json:"-"`
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ExchangeRateGetByDateParams]'s query parameters as
// `url.Values`.
func (r ExchangeRateGetByDateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExchangeRateUpdateByDateParams struct {
	From string `path:"from" api:"required" json:"-"`
	To   string `path:"to" api:"required" json:"-"`
	// The exchange rate from the base currency to the destination currency.
	Rate     string            `json:"rate" api:"required"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExchangeRateUpdateByDateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExchangeRateUpdateByDateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExchangeRateUpdateByDateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
