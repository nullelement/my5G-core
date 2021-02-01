# NnwdafEventsSubscriptionNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventNotifications** | [**[]EventNotification**](EventNotification.md) | Notifications about Individual Events | 
**SubscriptionId** | **string** | String identifying a subscription to the Nnwdaf_EventsSubscription Service | 

## Methods

### NewNnwdafEventsSubscriptionNotification

`func NewNnwdafEventsSubscriptionNotification(eventNotifications []EventNotification, subscriptionId string, ) *NnwdafEventsSubscriptionNotification`

NewNnwdafEventsSubscriptionNotification instantiates a new NnwdafEventsSubscriptionNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNnwdafEventsSubscriptionNotificationWithDefaults

`func NewNnwdafEventsSubscriptionNotificationWithDefaults() *NnwdafEventsSubscriptionNotification`

NewNnwdafEventsSubscriptionNotificationWithDefaults instantiates a new NnwdafEventsSubscriptionNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventNotifications

`func (o *NnwdafEventsSubscriptionNotification) GetEventNotifications() []EventNotification`

GetEventNotifications returns the EventNotifications field if non-nil, zero value otherwise.

### GetEventNotificationsOk

`func (o *NnwdafEventsSubscriptionNotification) GetEventNotificationsOk() (*[]EventNotification, bool)`

GetEventNotificationsOk returns a tuple with the EventNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifications

`func (o *NnwdafEventsSubscriptionNotification) SetEventNotifications(v []EventNotification)`

SetEventNotifications sets EventNotifications field to given value.


### GetSubscriptionId

`func (o *NnwdafEventsSubscriptionNotification) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *NnwdafEventsSubscriptionNotification) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *NnwdafEventsSubscriptionNotification) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


