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
	"github.com/stainless-sdks/emcees-prod-testing-5-go/internal/requestconfig"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/option"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/param"
	"github.com/stainless-sdks/emcees-prod-testing-5-go/packages/respjson"
)

// These endpoints can be used to manage the user&#039;s webhooks and triggers them
// if necessary.
//
// WebhookMessageService contains methods and other services that help with
// interacting with the firefly API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookMessageService] method instead.
type WebhookMessageService struct {
	options []option.RequestOption
	// These endpoints can be used to manage the user&#039;s webhooks and triggers them
	// if necessary.
	Attempts WebhookMessageAttemptService
}

// NewWebhookMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookMessageService(opts ...option.RequestOption) (r WebhookMessageService) {
	r = WebhookMessageService{}
	r.options = opts
	r.Attempts = NewWebhookMessageAttemptService(opts...)
	return
}

// When a webhook is triggered it will store the actual content of the webhook in a
// webhook message. You can view and analyse a single one using this endpoint.
func (r *WebhookMessageService) Get(ctx context.Context, messageID int64, params WebhookMessageGetParams, opts ...option.RequestOption) (res *WebhookMessageGetResponse, err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s/messages/%v", url.PathEscape(params.ID), messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// When a webhook is triggered the actual message that will be send is stored in a
// "message". You can view and analyse these messages.
func (r *WebhookMessageService) List(ctx context.Context, id string, query WebhookMessageListParams, opts ...option.RequestOption) (res *WebhookMessageListResponse, err error) {
	if !param.IsOmitted(query.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", query.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.api+json")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s/messages", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a webhook message. Any time a webhook is triggered the message is stored
// before it's sent. You can delete them before or after sending.
func (r *WebhookMessageService) Delete(ctx context.Context, messageID int64, params WebhookMessageDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XTraceID) {
		opts = append(opts, option.WithHeader("X-Trace-Id", fmt.Sprintf("%v", params.XTraceID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/webhooks/%s/messages/%v", url.PathEscape(params.ID), messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WebhookMessage struct {
	ID         string                   `json:"id" api:"required"`
	Attributes WebhookMessageAttributes `json:"attributes" api:"required"`
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
func (r WebhookMessage) RawJSON() string { return r.JSON.raw }
func (r *WebhookMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookMessageAttributes struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// If this message has errored out.
	Errored bool `json:"errored"`
	// The actual message that is sent or will be sent as JSON string.
	Message string `json:"message" api:"nullable"`
	// If this message is sent yet.
	Sent      bool      `json:"sent"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Long UUID string for identification of this webhook message.
	Uuid string `json:"uuid"`
	// The ID of the webhook this message belongs to.
	WebhookID string `json:"webhook_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Errored     respjson.Field
		Message     respjson.Field
		Sent        respjson.Field
		UpdatedAt   respjson.Field
		Uuid        respjson.Field
		WebhookID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookMessageAttributes) RawJSON() string { return r.JSON.raw }
func (r *WebhookMessageAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookMessageGetResponse struct {
	Data WebhookMessage `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookMessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookMessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookMessageListResponse struct {
	Data []WebhookMessage `json:"data" api:"required"`
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
func (r WebhookMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookMessageGetParams struct {
	ID       string            `path:"id" api:"required" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type WebhookMessageListParams struct {
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type WebhookMessageDeleteParams struct {
	ID       string            `path:"id" api:"required" json:"-"`
	XTraceID param.Opt[string] `header:"X-Trace-Id,omitzero" format:"uuid" json:"-"`
	paramObj
}
