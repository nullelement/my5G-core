# NnwdafEventsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSubscriptions** | [**[]EventSubscription**](EventSubscription.md) | Subscribed events | 
**EvtReq** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**NotificationURI** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**SupportedFeatures** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**EventNotifications** | Pointer to [**[]EventNotification**](EventNotification.md) |  | [optional] 
**FailEventReports** | Pointer to [**[]FailureEventInfo**](FailureEventInfo.md) |  | [optional] 

## Methods

### NewNnwdafEventsSubscription

`func NewNnwdafEventsSubscription(eventSubscriptions []EventSubscription, ) *NnwdafEventsSubscription`

NewNnwdafEventsSubscription instantiates a new NnwdafEventsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNnwdafEventsSubscriptionWithDefaults

`func NewNnwdafEventsSubscriptionWithDefaults() *NnwdafEventsSubscription`

NewNnwdafEventsSubscriptionWithDefaults instantiates a new NnwdafEventsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSubscriptions

`func (o *NnwdafEventsSubscription) GetEventSubscriptions() []EventSubscription`

GetEventSubscriptions returns the EventSubscriptions field if non-nil, zero value otherwise.

### GetEventSubscriptionsOk

`func (o *NnwdafEventsSubscription) GetEventSubscriptionsOk() (*[]EventSubscription, bool)`

GetEventSubscriptionsOk returns a tuple with the EventSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscriptions

`func (o *NnwdafEventsSubscription) SetEventSubscriptions(v []EventSubscription)`

SetEventSubscriptions sets EventSubscriptions field to given value.


### GetEvtReq

`func (o *NnwdafEventsSubscription) GetEvtReq() map[string]interface{}`

GetEvtReq returns the EvtReq field if non-nil, zero value otherwise.

### GetEvtReqOk

`func (o *NnwdafEventsSubscription) GetEvtReqOk() (*map[string]interface{}, bool)`

GetEvtReqOk returns a tuple with the EvtReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvtReq

`func (o *NnwdafEventsSubscription) SetEvtReq(v map[string]interface{})`

SetEvtReq sets EvtReq field to given value.

### HasEvtReq

`func (o *NnwdafEventsSubscription) HasEvtReq() bool`

HasEvtReq returns a boolean if a field has been set.

### GetNotificationURI

`func (o *NnwdafEventsSubscription) GetNotificationURI() map[string]interface{}`

GetNotificationURI returns the NotificationURI field if non-nil, zero value otherwise.

### GetNotificationURIOk

`func (o *NnwdafEventsSubscription) GetNotificationURIOk() (*map[string]interface{}, bool)`

GetNotificationURIOk returns a tuple with the NotificationURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationURI

`func (o *NnwdafEventsSubscription) SetNotificationURI(v map[string]interface{})`

SetNotificationURI sets NotificationURI field to given value.

### HasNotificationURI

`func (o *NnwdafEventsSubscription) HasNotificationURI() bool`

HasNotificationURI returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NnwdafEventsSubscription) GetSupportedFeatures() map[string]interface{}`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NnwdafEventsSubscription) GetSupportedFeaturesOk() (*map[string]interface{}, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NnwdafEventsSubscription) SetSupportedFeatures(v map[string]interface{})`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NnwdafEventsSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetEventNotifications

`func (o *NnwdafEventsSubscription) GetEventNotifications() []EventNotification`

GetEventNotifications returns the EventNotifications field if non-nil, zero value otherwise.

### GetEventNotificationsOk

`func (o *NnwdafEventsSubscription) GetEventNotificationsOk() (*[]EventNotification, bool)`

GetEventNotificationsOk returns a tuple with the EventNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifications

`func (o *NnwdafEventsSubscription) SetEventNotifications(v []EventNotification)`

SetEventNotifications sets EventNotifications field to given value.

### HasEventNotifications

`func (o *NnwdafEventsSubscription) HasEventNotifications() bool`

HasEventNotifications returns a boolean if a field has been set.

### GetFailEventReports

`func (o *NnwdafEventsSubscription) GetFailEventReports() []FailureEventInfo`

GetFailEventReports returns the FailEventReports field if non-nil, zero value otherwise.

### GetFailEventReportsOk

`func (o *NnwdafEventsSubscription) GetFailEventReportsOk() (*[]FailureEventInfo, bool)`

GetFailEventReportsOk returns a tuple with the FailEventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailEventReports

`func (o *NnwdafEventsSubscription) SetFailEventReports(v []FailureEventInfo)`

SetFailEventReports sets FailEventReports field to given value.

### HasFailEventReports

`func (o *NnwdafEventsSubscription) HasFailEventReports() bool`

HasFailEventReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


