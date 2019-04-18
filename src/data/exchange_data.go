package data

import (
	"strings"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-rabbitmq/src/data/consts"
)

// ExchangeData is the representation of the exchanges endpoint
type ExchangeData struct {
	Name         string
	Vhost        string
	MessageStats struct {
		PublishIn        *int64 `json:"publish_in" metric_name:"exchange.messagesPublishedPerChannel" source_type:"gauge"`
		PublishOut       *int64 `json:"publish_out" metric_name:"exchange.messagesPublishedQueue" source_type:"gauge"`
		PublishInDetails struct {
			Rate *float64 `metric_name:"exchange.messagesPublishedPerChannelPerSecond" source_type:"gauge"`
		} `json:"publish_in_details"`
		PublishOutDetails struct {
			Rate *float64 `metric_name:"exchange.messagesPublishedQueuePerSecond" source_type:"gauge"`
		} `json:"publish_out_details"`
	} `json:"message_stats"`
	Type       string
	Durable    bool
	AutoDelete bool `json:"auto_delete"`
	Arguments  map[string]interface{}
}

// CollectInventory collects inventory data and reports it to the integration.Entity
func (e *ExchangeData) CollectInventory(entity *integration.Entity, bindingStats BindingStats) {
	SetInventoryItem(entity, strings.TrimPrefix(consts.ExchangeType, "ra-"), "type", e.Type)
	SetInventoryItem(entity, strings.TrimPrefix(consts.ExchangeType, "ra-"), "durable", ConvertBoolToInt(e.Durable))
	SetInventoryItem(entity, strings.TrimPrefix(consts.ExchangeType, "ra-"), "auto_delete", ConvertBoolToInt(e.AutoDelete))
	setInventoryMap(entity, strings.TrimPrefix(consts.ExchangeType, "ra-"), "arguments", e.Arguments)
	setInventoryBindings(entity, e, bindingStats)
}

// GetEntity creates an integration.Entity for this ExchangeData
func (e *ExchangeData) GetEntity(integration *integration.Integration, clusterName string) (*integration.Entity, []metric.Attribute, error) {
	return CreateEntity(integration, e.Name, strings.TrimPrefix(consts.ExchangeType, "ra-"), e.Vhost, clusterName)
}

// EntityType returns the type of this entity
func (e *ExchangeData) EntityType() string {
	return consts.ExchangeType
}

// EntityName returns the main name of this entity
func (e *ExchangeData) EntityName() string {
	return e.Name
}

// EntityVhost returns the vhost of this entity
func (e *ExchangeData) EntityVhost() string {
	return e.Vhost
}
