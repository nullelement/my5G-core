# UeCommunication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommDur** | [**map[string]interface{}**](object.md) |  | 
**CommDurVariance** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**PerioTime** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**PerioTimeVariance** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**Ts** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**TsVariance** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**RecurringTime** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**TrafChar** | [**TrafficCharacterization**](TrafficCharacterization.md) |  | 
**Ratio** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**Confidence** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 

## Methods

### NewUeCommunication

`func NewUeCommunication(commDur map[string]interface{}, trafChar TrafficCharacterization, ) *UeCommunication`

NewUeCommunication instantiates a new UeCommunication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeCommunicationWithDefaults

`func NewUeCommunicationWithDefaults() *UeCommunication`

NewUeCommunicationWithDefaults instantiates a new UeCommunication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommDur

`func (o *UeCommunication) GetCommDur() map[string]interface{}`

GetCommDur returns the CommDur field if non-nil, zero value otherwise.

### GetCommDurOk

`func (o *UeCommunication) GetCommDurOk() (*map[string]interface{}, bool)`

GetCommDurOk returns a tuple with the CommDur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommDur

`func (o *UeCommunication) SetCommDur(v map[string]interface{})`

SetCommDur sets CommDur field to given value.


### GetCommDurVariance

`func (o *UeCommunication) GetCommDurVariance() map[string]interface{}`

GetCommDurVariance returns the CommDurVariance field if non-nil, zero value otherwise.

### GetCommDurVarianceOk

`func (o *UeCommunication) GetCommDurVarianceOk() (*map[string]interface{}, bool)`

GetCommDurVarianceOk returns a tuple with the CommDurVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommDurVariance

`func (o *UeCommunication) SetCommDurVariance(v map[string]interface{})`

SetCommDurVariance sets CommDurVariance field to given value.

### HasCommDurVariance

`func (o *UeCommunication) HasCommDurVariance() bool`

HasCommDurVariance returns a boolean if a field has been set.

### GetPerioTime

`func (o *UeCommunication) GetPerioTime() map[string]interface{}`

GetPerioTime returns the PerioTime field if non-nil, zero value otherwise.

### GetPerioTimeOk

`func (o *UeCommunication) GetPerioTimeOk() (*map[string]interface{}, bool)`

GetPerioTimeOk returns a tuple with the PerioTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerioTime

`func (o *UeCommunication) SetPerioTime(v map[string]interface{})`

SetPerioTime sets PerioTime field to given value.

### HasPerioTime

`func (o *UeCommunication) HasPerioTime() bool`

HasPerioTime returns a boolean if a field has been set.

### GetPerioTimeVariance

`func (o *UeCommunication) GetPerioTimeVariance() map[string]interface{}`

GetPerioTimeVariance returns the PerioTimeVariance field if non-nil, zero value otherwise.

### GetPerioTimeVarianceOk

`func (o *UeCommunication) GetPerioTimeVarianceOk() (*map[string]interface{}, bool)`

GetPerioTimeVarianceOk returns a tuple with the PerioTimeVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerioTimeVariance

`func (o *UeCommunication) SetPerioTimeVariance(v map[string]interface{})`

SetPerioTimeVariance sets PerioTimeVariance field to given value.

### HasPerioTimeVariance

`func (o *UeCommunication) HasPerioTimeVariance() bool`

HasPerioTimeVariance returns a boolean if a field has been set.

### GetTs

`func (o *UeCommunication) GetTs() map[string]interface{}`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *UeCommunication) GetTsOk() (*map[string]interface{}, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *UeCommunication) SetTs(v map[string]interface{})`

SetTs sets Ts field to given value.

### HasTs

`func (o *UeCommunication) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTsVariance

`func (o *UeCommunication) GetTsVariance() map[string]interface{}`

GetTsVariance returns the TsVariance field if non-nil, zero value otherwise.

### GetTsVarianceOk

`func (o *UeCommunication) GetTsVarianceOk() (*map[string]interface{}, bool)`

GetTsVarianceOk returns a tuple with the TsVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsVariance

`func (o *UeCommunication) SetTsVariance(v map[string]interface{})`

SetTsVariance sets TsVariance field to given value.

### HasTsVariance

`func (o *UeCommunication) HasTsVariance() bool`

HasTsVariance returns a boolean if a field has been set.

### GetRecurringTime

`func (o *UeCommunication) GetRecurringTime() map[string]interface{}`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *UeCommunication) GetRecurringTimeOk() (*map[string]interface{}, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *UeCommunication) SetRecurringTime(v map[string]interface{})`

SetRecurringTime sets RecurringTime field to given value.

### HasRecurringTime

`func (o *UeCommunication) HasRecurringTime() bool`

HasRecurringTime returns a boolean if a field has been set.

### GetTrafChar

`func (o *UeCommunication) GetTrafChar() TrafficCharacterization`

GetTrafChar returns the TrafChar field if non-nil, zero value otherwise.

### GetTrafCharOk

`func (o *UeCommunication) GetTrafCharOk() (*TrafficCharacterization, bool)`

GetTrafCharOk returns a tuple with the TrafChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafChar

`func (o *UeCommunication) SetTrafChar(v TrafficCharacterization)`

SetTrafChar sets TrafChar field to given value.


### GetRatio

`func (o *UeCommunication) GetRatio() map[string]interface{}`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *UeCommunication) GetRatioOk() (*map[string]interface{}, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *UeCommunication) SetRatio(v map[string]interface{})`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *UeCommunication) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetConfidence

`func (o *UeCommunication) GetConfidence() map[string]interface{}`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *UeCommunication) GetConfidenceOk() (*map[string]interface{}, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *UeCommunication) SetConfidence(v map[string]interface{})`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *UeCommunication) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


