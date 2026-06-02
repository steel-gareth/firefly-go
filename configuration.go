// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firefly

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/steel-gareth/firefly-go/internal/apijson"
	"github.com/steel-gareth/firefly-go/internal/requestconfig"
	"github.com/steel-gareth/firefly-go/option"
	"github.com/steel-gareth/firefly-go/packages/param"
	"github.com/steel-gareth/firefly-go/packages/respjson"
)

// These endpoints allow you to manage and update the Firefly III configuration.
// You need to have the &quot;owner&quot; role to update configuration.
//
// ConfigurationService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigurationService] method instead.
type ConfigurationService struct {
	options []option.RequestOption
}

// NewConfigurationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigurationService(opts ...option.RequestOption) (r ConfigurationService) {
	r = ConfigurationService{}
	r.options = opts
	return
}

// Returns all editable and not-editable configuration values for this Firefly III
// installation
func (r *ConfigurationService) Get(ctx context.Context, query ConfigurationGetParams, opts ...option.RequestOption) (res *[]Configuration, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/configuration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns one configuration variable for this Firefly III installation
func (r *ConfigurationService) GetValue(ctx context.Context, name ConfigValueFilter, query ConfigurationGetValueParams, opts ...option.RequestOption) (res *ConfigurationSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/configuration/%v", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set a single configuration value. Not all configuration values can be updated so
// the list of accepted configuration variables is small.
func (r *ConfigurationService) UpdateValue(ctx context.Context, name ConfigurationUpdateValueParamsName, params ConfigurationUpdateValueParams, opts ...option.RequestOption) (res *ConfigurationSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/configuration/%v", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Title of the configuration value.
type ConfigValueFilter string

const (
	ConfigValueFilterConfigurationIsDemoSite            ConfigValueFilter = "configuration.is_demo_site"
	ConfigValueFilterConfigurationPermissionUpdateCheck ConfigValueFilter = "configuration.permission_update_check"
	ConfigValueFilterConfigurationLastUpdateCheck       ConfigValueFilter = "configuration.last_update_check"
	ConfigValueFilterConfigurationSingleUserMode        ConfigValueFilter = "configuration.single_user_mode"
	ConfigValueFilterFireflyVersion                     ConfigValueFilter = "firefly.version"
	ConfigValueFilterFireflyDefaultLocation             ConfigValueFilter = "firefly.default_location"
	ConfigValueFilterFireflyAccountToTransaction        ConfigValueFilter = "firefly.account_to_transaction"
	ConfigValueFilterFireflyAllowedOpposingTypes        ConfigValueFilter = "firefly.allowed_opposing_types"
	ConfigValueFilterFireflyAccountRoles                ConfigValueFilter = "firefly.accountRoles"
	ConfigValueFilterFireflyValidLiabilities            ConfigValueFilter = "firefly.valid_liabilities"
	ConfigValueFilterFireflyInterestPeriods             ConfigValueFilter = "firefly.interest_periods"
	ConfigValueFilterFireflyEnableExternalMap           ConfigValueFilter = "firefly.enable_external_map"
	ConfigValueFilterFireflyExpectedSourceTypes         ConfigValueFilter = "firefly.expected_source_types"
	ConfigValueFilterAppTimezone                        ConfigValueFilter = "app.timezone"
	ConfigValueFilterFireflyBillPeriods                 ConfigValueFilter = "firefly.bill_periods"
	ConfigValueFilterFireflyCreditCardTypes             ConfigValueFilter = "firefly.credit_card_types"
	ConfigValueFilterFireflyLanguages                   ConfigValueFilter = "firefly.languages"
	ConfigValueFilterFireflyValidViewRanges             ConfigValueFilter = "firefly.valid_view_ranges"
	ConfigValueFilterCerEnabled                         ConfigValueFilter = "cer.enabled"
	ConfigValueFilterFireflyPreselectedAccounts         ConfigValueFilter = "firefly.preselected_accounts"
	ConfigValueFilterFireflyRuleActions                 ConfigValueFilter = "firefly.rule-actions"
	ConfigValueFilterFireflyContextRuleActions          ConfigValueFilter = "firefly.context-rule-actions"
	ConfigValueFilterSearchOperators                    ConfigValueFilter = "search.operators"
	ConfigValueFilterWebhookTriggers                    ConfigValueFilter = "webhook.triggers"
	ConfigValueFilterWebhookResponses                   ConfigValueFilter = "webhook.responses"
	ConfigValueFilterWebhookDeliveries                  ConfigValueFilter = "webhook.deliveries"
)

type Configuration struct {
	// If this config variable can be edited by the user
	Editable bool `json:"editable" api:"required"`
	// Title of the configuration value.
	//
	// Any of "configuration.is_demo_site", "configuration.permission_update_check",
	// "configuration.last_update_check", "configuration.single_user_mode",
	// "firefly.version", "firefly.default_location", "firefly.account_to_transaction",
	// "firefly.allowed_opposing_types", "firefly.accountRoles",
	// "firefly.valid_liabilities", "firefly.interest_periods",
	// "firefly.enable_external_map", "firefly.expected_source_types", "app.timezone",
	// "firefly.bill_periods", "firefly.credit_card_types", "firefly.languages",
	// "firefly.valid_view_ranges", "cer.enabled", "firefly.preselected_accounts",
	// "firefly.rule-actions", "firefly.context-rule-actions", "search.operators",
	// "webhook.triggers", "webhook.responses", "webhook.deliveries".
	Title ConfigValueFilter        `json:"title" api:"required"`
	Value PolymorphicPropertyUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Editable    respjson.Field
		Title       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Configuration) RawJSON() string { return r.JSON.raw }
func (r *Configuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigurationSingle struct {
	Data Configuration `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationSingle) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolymorphicPropertyUnion contains all possible properties and values from
// [bool], [string], [map[string]any], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfString OfPolymorphicPropertyMapItem OfStringArray]
type PolymorphicPropertyUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfPolymorphicPropertyMapItem any `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfBool                       respjson.Field
		OfString                     respjson.Field
		OfPolymorphicPropertyMapItem respjson.Field
		OfStringArray                respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PolymorphicPropertyUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolymorphicPropertyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolymorphicPropertyUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PolymorphicPropertyUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PolymorphicPropertyUnion) RawJSON() string { return u.JSON.raw }

func (r *PolymorphicPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolymorphicPropertyUnion to a
// PolymorphicPropertyUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolymorphicPropertyUnionParam.Overrides()
func (r PolymorphicPropertyUnion) ToParam() PolymorphicPropertyUnionParam {
	return param.Override[PolymorphicPropertyUnionParam](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PolymorphicPropertyUnionParam struct {
	OfBool        param.Opt[bool]   `json:",omitzero,inline"`
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap      map[string]any    `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PolymorphicPropertyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfString, u.OfAnyMap, u.OfStringArray)
}
func (u *PolymorphicPropertyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type ConfigurationGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ConfigurationGetValueParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ConfigurationUpdateValueParams struct {
	Value    PolymorphicPropertyUnionParam `json:"value,omitzero" api:"required"`
	XTraceID param.Opt[string]             `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ConfigurationUpdateValueParams) MarshalJSON() (data []byte, err error) {
	type shadow ConfigurationUpdateValueParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigurationUpdateValueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigurationUpdateValueParamsName string

const (
	ConfigurationUpdateValueParamsNameConfigurationIsDemoSite            ConfigurationUpdateValueParamsName = "configuration.is_demo_site"
	ConfigurationUpdateValueParamsNameConfigurationPermissionUpdateCheck ConfigurationUpdateValueParamsName = "configuration.permission_update_check"
	ConfigurationUpdateValueParamsNameConfigurationLastUpdateCheck       ConfigurationUpdateValueParamsName = "configuration.last_update_check"
	ConfigurationUpdateValueParamsNameConfigurationSingleUserMode        ConfigurationUpdateValueParamsName = "configuration.single_user_mode"
	ConfigurationUpdateValueParamsNameConfigurationEnableExchangeRates   ConfigurationUpdateValueParamsName = "configuration.enable_exchange_rates"
	ConfigurationUpdateValueParamsNameConfigurationUseRunningBalance     ConfigurationUpdateValueParamsName = "configuration.use_running_balance"
	ConfigurationUpdateValueParamsNameConfigurationEnableExternalMap     ConfigurationUpdateValueParamsName = "configuration.enable_external_map"
	ConfigurationUpdateValueParamsNameConfigurationEnableExternalRates   ConfigurationUpdateValueParamsName = "configuration.enable_external_rates"
	ConfigurationUpdateValueParamsNameConfigurationAllowWebhooks         ConfigurationUpdateValueParamsName = "configuration.allow_webhooks"
	ConfigurationUpdateValueParamsNameConfigurationValidURLProtocols     ConfigurationUpdateValueParamsName = "configuration.valid_url_protocols"
)
