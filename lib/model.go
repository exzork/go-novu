package lib

import (
	"io"
	"time"
)

type ChannelType string
type GeneralError error

const Version = "v1"

const (
	HTTPStatusOk       = 200
	HTTPStatusCreated  = 201
	HTTPStatusNoChange = 204
	HTTPRedirectOk     = 300
)

const (
	EMAIL  ChannelType = "EMAIL"
	SMS    ChannelType = "SMS"
	DIRECT ChannelType = "DIRECT"
)

type Data struct {
	Acknowledged bool   `json:"acknowledged"`
	Status       string `json:"status"`
}

type Response struct {
	Data Data `json:"data"`
}

type ITriggerPayloadOptions struct {
	To            interface{} `json:"to,omitempty"`
	Payload       interface{} `json:"payload,omitempty"`
	Overrides     interface{} `json:"overrides,omitempty"`
	TransactionId string      `json:"transactionId,omitempty"`
	Actor         interface{} `json:"actor,omitempty"`
}

type TriggerRecipientsTypeArray interface {
	[]string | []SubscriberPayload
}
type TriggerRecipientsTypeSingle interface {
	string | SubscriberPayload
}

type TriggerTopicRecipientsTypeSingle struct {
	TopicKey string `json:"topicKey,omitempty"`
	Type     string `json:"type,omitempty"`
}

type SubscriberPayload struct {
	FirstName string                 `json:"first_name,omitempty"`
	LastName  string                 `json:"last_name,omitempty"`
	Email     string                 `json:"email,omitempty"`
	Phone     string                 `json:"phone,omitempty"`
	Avatar    string                 `json:"avatar,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

type TriggerRecipientsType interface {
	TriggerRecipientsTypeSingle | TriggerRecipientsTypeArray | TriggerTopicRecipientsTypeSingle | []TriggerTopicRecipientsTypeSingle
}

type ITriggerPayload interface {
	string | []string | bool | int64 | IAttachmentOptions | []IAttachmentOptions
}

type IAttachmentOptions struct {
	Mime     string        `json:"mime,omitempty"`
	File     io.Reader     `json:"file,omitempty"`
	Name     string        `json:"name,omitempty"`
	Channels []ChannelType `json:"channels,omitempty"`
}

type EventResponse struct {
	Data interface{} `json:"data"`
}

type EventRequest struct {
	Name          string      `json:"name"`
	To            interface{} `json:"to"`
	Payload       interface{} `json:"payload"`
	Overrides     interface{} `json:"overrides,omitempty"`
	TransactionId string      `json:"transactionId,omitempty"`
	Actor         interface{} `json:"actor,omitempty"`
}

type SubscriberResponse struct {
	Data interface{} `json:"data"`
}

type Template struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Critical bool   `json:"critical"`
}

type Preference struct {
	Enabled  bool    `json:"enabled"`
	Channels Channel `json:"channels"`
}

type Channel struct {
	Email bool `json:"email"`
	Sms   bool `json:"sms"`
	Chat  bool `json:"chat"`
	InApp bool `json:"in_app"`
	Push  bool `json:"push"`
}

type SubscriberPreferencesResponse struct {
	Data []struct {
		Template   Template   `json:"template"`
		Preference Preference `json:"preference"`
	} `json:"data"`
}

type UpdateSubscriberPreferencesChannel struct {
	Type    ChannelType `json:"type"`
	Enabled bool        `json:"enabled"`
}

type UpdateSubscriberPreferencesOptions struct {
	Channel []UpdateSubscriberPreferencesChannel `json:"channel,omitempty"`
	Enabled bool                                 `json:"enabled,omitempty"`
}

type ListTopicsResponse struct {
	Page       int                `json:"name"`
	PageSize   int                `json:"pageSize"`
	TotalCount int                `json:"totalCount"`
	Data       []GetTopicResponse `json:"data"`
}

type GetTopicResponse struct {
	Id             string   `json:"_id"`
	OrganizationId string   `json:"_organizationId"`
	EnvironmentId  string   `json:"_environmentId"`
	Key            string   `json:"key"`
	Name           string   `json:"name"`
	Subscribers    []string `json:"subscribers"`
}

type ListTopicsOptions struct {
	Page     *int    `json:"page,omitempty"`
	PageSize *int    `json:"pageSize,omitempty"`
	Key      *string `json:"key,omitempty"`
}

type CreateTopicRequest struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type RenameTopicRequest struct {
	Name string `json:"name"`
}

type SubscribersTopicRequest struct {
	Subscribers []string `json:"subscribers"`
}

type SubscriberNotificationFeedOptions struct {
	Page           int    `queryKey:"page"`
	FeedIdentifier string `queryKey:"feedIdentifier"`
	Seen           bool   `queryKey:"seen"`
}

type SubscriberUnseenCountOptions struct {
	Seen *bool `json:"seen"`
}

type SubscriberMarkMessageSeenOptions struct {
	MessageID string `json:"messageId"`
	Seen      bool   `json:"seen"`
	Read      bool   `json:"read"`
}

type NotificationFeedData struct {
	CTA              CTA       `json:"cta"`
	Channel          string    `json:"channel"`
	Content          string    `json:"content"`
	CreatedAt        time.Time `json:"createdAt"`
	Deleted          bool      `json:"deleted"`
	DeviceTokens     []string  `json:"deviceTokens"`
	DirectWebhookURL string    `json:"directWebhookUrl"`
	EnvironmentID    string    `json:"_environmentId"`
	ErrorID          string    `json:"errorId"`
	ErrorText        string    `json:"errorText"`
	FeedID           string    `json:"_feedId"`
	ID               string    `json:"_id"`
	JobID            string    `json:"_jobId"`
	LastReadDate     time.Time `json:"lastReadDate"`
	LastSeenDate     time.Time `json:"lastSeenDate"`
	MessageTemplate  string    `json:"_messageTemplateId"`
	NotificationID   string    `json:"_notificationId"`
	OrganizationID   string    `json:"_organizationId"`
	Payload          struct {
		UpdateMessage string `json:"updateMessage"`
	} `json:"payload"`
	ProviderID string `json:"providerId"`
	Read       bool   `json:"read"`
	ResponseID string `json:"id"`
	Seen       bool   `json:"seen"`
	Status     string `json:"status"`
	Subscriber struct {
		ID           string `json:"_id"`
		SubscriberID string `json:"subscriberId"`
	} `json:"subscriber"`
	SubscriberID       string    `json:"_subscriberId"`
	TemplateID         string    `json:"_templateId"`
	TemplateIdentifier string    `json:"templateIdentifier"`
	TransactionID      string    `json:"transactionId"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

type SubscriberNotificationFeedResponse struct {
	TotalCount int                    `json:"totalCount"`
	Data       []NotificationFeedData `json:"data"`
	PageSize   int                    `json:"pageSize"`
	Page       int                    `json:"page"`
}

type SubscriberUnseenCountResponse struct {
	Data struct {
		Count int `json:"count"`
	} `json:"data"`
}

type CTA struct {
	Type   string `json:"type"`
	Action struct {
		Status  string `json:"status"`
		Buttons struct {
			Type          string `json:"type"`
			Content       string `json:"content"`
			ResultContent string `json:"resultContent"`
		} `json:"buttons"`
		Result struct {
			Payload map[string]interface{} `json:"payload"`
			Type    string                 `json:"type"`
		} `json:"result"`
	}
}
type IntegrationCredentials struct {
	ApiKey           string                 `json:"apiKey,omitempty"`
	User             string                 `json:"user,omitempty"`
	SecretKey        string                 `json:"secretKey,omitempty"`
	Domain           string                 `json:"domain,omitempty"`
	Password         string                 `json:"password,omitempty"`
	Host             string                 `json:"host,omitempty"`
	Port             string                 `json:"port,omitempty"`
	Secure           bool                   `json:"secure,omitempty"`
	Region           string                 `json:"region,omitempty"`
	AccountSID       string                 `json:"accountSid,omitempty"`
	MessageProfileID string                 `json:"messageProfileId,omitempty"`
	Token            string                 `json:"token,omitempty"`
	From             string                 `json:"from,omitempty"`
	SenderName       string                 `json:"senderName,omitempty"`
	ProjectName      string                 `json:"projectName,omitempty"`
	ApplicationID    string                 `json:"applicationId,omitempty"`
	ClientID         string                 `json:"clientId,omitempty"`
	RequireTls       bool                   `json:"requireTls,omitempty"`
	IgnoreTls        bool                   `json:"ignoreTls,omitempty"`
	TlsOptions       map[string]interface{} `json:"tlsOptions,omitempty"`
}

type CreateIntegrationRequest struct {
	ProviderID  string                 `json:"providerId"`
	Channel     ChannelType            `json:"channel"`
	Credentials IntegrationCredentials `json:"credentials,omitempty"`
	Active      bool                   `json:"active"`
	Check       bool                   `json:"check"`
}

type UpdateIntegrationRequest struct {
	Credentials IntegrationCredentials `json:"credentials"`
	Active      bool                   `json:"active"`
	Check       bool                   `json:"check"`
}

type Integration struct {
	Id             string                 `json:"_id"`
	EnvironmentID  string                 `json:"_environmentId"`
	OrganizationID string                 `json:"_organizationId"`
	ProviderID     string                 `json:"providerId"`
	Channel        ChannelType            `json:"channel"`
	Credentials    IntegrationCredentials `json:"credentials"`
	Active         bool                   `json:"active"`
	Deleted        bool                   `json:"deleted"`
	UpdatedAt      string                 `json:"updatedAt"`
	DeletedAt      string                 `json:"deletedAt"`
	DeletedBy      string                 `json:"deletedBy"`
}

type IntegrationResponse struct {
	Data Integration `json:"data"`
}

type GetIntegrationsResponse struct {
	Data []Integration `json:"data"`
}

type BulkTriggerOptions struct {
	Name          interface{} `json:"name,omitempty"`
	To            interface{} `json:"to,omitempty"`
	Payload       interface{} `json:"payload,omitempty"`
	Overrides     interface{} `json:"overrides,omitempty"`
	TransactionId string      `json:"transactionId,omitempty"`
	Actor         interface{} `json:"actor,omitempty"`
}

type BulkTriggerEvent struct {
	Events []BulkTriggerOptions `json:"events"`
}

type BroadcastEventToAll struct {
	Name          interface{} `json:"name,omitempty"`
	Payload       interface{} `json:"payload,omitempty"`
	Overrides     interface{} `json:"overrides,omitempty"`
	TransactionId string      `json:"transactionId,omitempty"`
	Actor         interface{} `json:"actor,omitempty"`
}

type UpdateSubscriptionCredentialsOptions struct {
	ProviderID  string      `json:"providerId"`
	Credentials interface{} `json:"credentials"`
}
