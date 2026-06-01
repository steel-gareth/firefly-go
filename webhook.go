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

// These endpoints can be used to manage the user&#039;s webhooks and triggers them
// if necessary.
//
// WebhookService contains methods and other services that help with interacting
// with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	options []option.RequestOption
	// These endpoints can be used to manage the user&#039;s webhooks and triggers them
	// if necessary.
	Messages WebhookMessageService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.options = opts
	r.Messages = NewWebhookMessageService(opts...)
	return
}

// Creates a new webhook. The data required can be submitted as a JSON body or as a
// list of parameters. The webhook will be given a random secret.
func (r *WebhookService) New(ctx context.Context, params WebhookNewParams, opts ...option.RequestOption) (res *WebhookSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Gets all info of a single webhook.
func (r *WebhookService) Get(ctx context.Context, id string, query WebhookGetParams, opts ...option.RequestOption) (res *WebhookSingle, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing webhook's information. If you wish to reset the secret,
// submit any value as the "secret". Firefly III will take this as a hint and reset
// the secret of the webhook.
func (r *WebhookService) Update(ctx context.Context, id string, params WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookSingle, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all the user's webhooks.
func (r *WebhookService) List(ctx context.Context, params WebhookListParams, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	path := "v1/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a webhook.
func (r *WebhookService) Delete(ctx context.Context, id string, body WebhookDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/webhooks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// This endpoint will submit any open messages for this webhook. This is an
// asynchronous operation, so you can't see the result. Refresh the webhook message
// and/or the webhook message attempts to see the results. This may take some time
// if the webhook receiver is slow.
func (r *WebhookService) Submit(ctx context.Context, id string, body WebhookSubmitParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", body.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/webhooks/%s/submit", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// This endpoint will execute this webhook for a given transaction ID. This is an
// asynchronous operation, so you can't see the result. Refresh the webhook message
// and/or the webhook message attempts to see the results. This may take some time
// if the webhook receiver is slow.
func (r *WebhookService) TriggerTransaction(ctx context.Context, transactionID string, params WebhookTriggerTransactionParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return err
	}
	path := fmt.Sprintf("v1/webhooks/%s/trigger-transaction/%s", url.PathEscape(params.ID), url.PathEscape(transactionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type Webhook struct {
	ID         string            `json:"id" api:"required"`
	Attributes WebhookAttributes `json:"attributes" api:"required"`
	Links      ObjectLink        `json:"links" api:"required"`
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
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookAttributes struct {
	Delivery any `json:"delivery" api:"required"`
	Response any `json:"response" api:"required"`
	// A title for the webhook for easy recognition.
	Title   string `json:"title" api:"required"`
	Trigger any    `json:"trigger" api:"required"`
	// The URL of the webhook. Has to start with `https`.
	URL string `json:"url" api:"required"`
	// Boolean to indicate if the webhook is active
	Active     bool              `json:"active"`
	CreatedAt  time.Time         `json:"created_at" format:"date-time"`
	Deliveries []WebhookDelivery `json:"deliveries"`
	Responses  []WebhookResponse `json:"responses"`
	// A 24-character secret for the webhook. It's generated by Firefly III when saving
	// a new webhook. If you submit a new secret through the PUT endpoint it will
	// generate a new secret for the selected webhook, a new secret bearing no relation
	// to whatever you just submitted.
	Secret    string           `json:"secret"`
	Triggers  []WebhookTrigger `json:"triggers"`
	UpdatedAt time.Time        `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delivery    respjson.Field
		Response    respjson.Field
		Title       respjson.Field
		Trigger     respjson.Field
		URL         respjson.Field
		Active      respjson.Field
		CreatedAt   respjson.Field
		Deliveries  respjson.Field
		Responses   respjson.Field
		Secret      respjson.Field
		Triggers    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookAttributes) RawJSON() string { return r.JSON.raw }
func (r *WebhookAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format of the delivered response.
type WebhookDelivery string

const (
	WebhookDeliveryJson WebhookDelivery = "JSON"
)

// Indicator for what Firefly III will deliver to the webhook URL.
type WebhookResponse string

const (
	WebhookResponseTransactions WebhookResponse = "TRANSACTIONS"
	WebhookResponseAccounts     WebhookResponse = "ACCOUNTS"
	WebhookResponseBudget       WebhookResponse = "BUDGET"
	WebhookResponseRelevant     WebhookResponse = "RELEVANT"
	WebhookResponseNone         WebhookResponse = "NONE"
)

type WebhookSingle struct {
	Data Webhook `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSingle) RawJSON() string { return r.JSON.raw }
func (r *WebhookSingle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The trigger for the webhook.
type WebhookTrigger string

const (
	WebhookTriggerAny                    WebhookTrigger = "ANY"
	WebhookTriggerStoreTransaction       WebhookTrigger = "STORE_TRANSACTION"
	WebhookTriggerUpdateTransaction      WebhookTrigger = "UPDATE_TRANSACTION"
	WebhookTriggerDestroyTransaction     WebhookTrigger = "DESTROY_TRANSACTION"
	WebhookTriggerStoreBudget            WebhookTrigger = "STORE_BUDGET"
	WebhookTriggerUpdateBudget           WebhookTrigger = "UPDATE_BUDGET"
	WebhookTriggerDestroyBudget          WebhookTrigger = "DESTROY_BUDGET"
	WebhookTriggerStoreUpdateBudgetLimit WebhookTrigger = "STORE_UPDATE_BUDGET_LIMIT"
)

type WebhookListResponse struct {
	Data  []Webhook `json:"data" api:"required"`
	Links PageLink  `json:"links" api:"required"`
	Meta  Meta      `json:"meta" api:"required"`
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
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	Delivery any `json:"delivery,omitzero" api:"required"`
	Response any `json:"response,omitzero" api:"required"`
	// A title for the webhook for easy recognition.
	Title   string `json:"title" api:"required"`
	Trigger any    `json:"trigger,omitzero" api:"required"`
	// The URL of the webhook. Has to start with `https`.
	URL string `json:"url" api:"required"`
	// Boolean to indicate if the webhook is active
	Active     param.Opt[bool]   `json:"active,omitzero"`
	XTraceID   param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	Deliveries []WebhookDelivery `json:"deliveries,omitzero"`
	Responses  []WebhookResponse `json:"responses,omitzero"`
	Triggers   []WebhookTrigger  `json:"triggers,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type WebhookUpdateParams struct {
	// Boolean to indicate if the webhook is active
	Active param.Opt[bool] `json:"active,omitzero"`
	// A 24-character secret for the webhook. It's generated by Firefly III when saving
	// a new webhook. If you submit a new secret through the PUT endpoint it will
	// generate a new secret for the selected webhook, a new secret bearing no relation
	// to whatever you just submitted.
	Secret param.Opt[string] `json:"secret,omitzero"`
	// A title for the webhook for easy recognition.
	Title param.Opt[string] `json:"title,omitzero"`
	// The URL of the webhook. Has to start with `https`.
	URL        param.Opt[string] `json:"url,omitzero"`
	XTraceID   param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	Deliveries []WebhookDelivery `json:"deliveries,omitzero"`
	Responses  []WebhookResponse `json:"responses,omitzero"`
	Triggers   []WebhookTrigger  `json:"triggers,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListParams struct {
	// Number of items per page. The default pagination is per 50 items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number. The default pagination is per 50 items.
	Page     param.Opt[int64]  `query:"page,omitzero" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeleteParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type WebhookSubmitParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type WebhookTriggerTransactionParams struct {
	ID       string            `path:"id" api:"required" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
