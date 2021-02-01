# QosSustainabilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaInfo** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**StartTs** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**EndTs** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**QosFlowRetThd** | Pointer to [**RetainabilityThreshold**](RetainabilityThreshold.md) |  | [optional] 
**RanUeThrouThd** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**Snssai** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**Confidence** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 

## Methods

### NewQosSustainabilityInfo

`func NewQosSustainabilityInfo() *QosSustainabilityInfo`

NewQosSustainabilityInfo instantiates a new QosSustainabilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosSustainabilityInfoWithDefaults

`func NewQosSustainabilityInfoWithDefaults() *QosSustainabilityInfo`

NewQosSustainabilityInfoWithDefaults instantiates a new QosSustainabilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaInfo

`func (o *QosSustainabilityInfo) GetAreaInfo() map[string]interface{}`

GetAreaInfo returns the AreaInfo field if non-nil, zero value otherwise.

### GetAreaInfoOk

`func (o *QosSustainabilityInfo) GetAreaInfoOk() (*map[string]interface{}, bool)`

GetAreaInfoOk returns a tuple with the AreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaInfo

`func (o *QosSustainabilityInfo) SetAreaInfo(v map[string]interface{})`

SetAreaInfo sets AreaInfo field to given value.

### HasAreaInfo

`func (o *QosSustainabilityInfo) HasAreaInfo() bool`

HasAreaInfo returns a boolean if a field has been set.

### GetStartTs

`func (o *QosSustainabilityInfo) GetStartTs() map[string]interface{}`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *QosSustainabilityInfo) GetStartTsOk() (*map[string]interface{}, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *QosSustainabilityInfo) SetStartTs(v map[string]interface{})`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *QosSustainabilityInfo) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *QosSustainabilityInfo) GetEndTs() map[string]interface{}`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *QosSustainabilityInfo) GetEndTsOk() (*map[string]interface{}, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *QosSustainabilityInfo) SetEndTs(v map[string]interface{})`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *QosSustainabilityInfo) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetQosFlowRetThd

`func (o *QosSustainabilityInfo) GetQosFlowRetThd() RetainabilityThreshold`

GetQosFlowRetThd returns the QosFlowRetThd field if non-nil, zero value otherwise.

### GetQosFlowRetThdOk

`func (o *QosSustainabilityInfo) GetQosFlowRetThdOk() (*RetainabilityThreshold, bool)`

GetQosFlowRetThdOk returns a tuple with the QosFlowRetThd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowRetThd

`func (o *QosSustainabilityInfo) SetQosFlowRetThd(v RetainabilityThreshold)`

SetQosFlowRetThd sets QosFlowRetThd field to given value.

### HasQosFlowRetThd

`func (o *QosSustainabilityInfo) HasQosFlowRetThd() bool`

HasQosFlowRetThd returns a boolean if a field has been set.

### GetRanUeThrouThd

`func (o *QosSustainabilityInfo) GetRanUeThrouThd() map[string]interface{}`

GetRanUeThrouThd returns the RanUeThrouThd field if non-nil, zero value otherwise.

### GetRanUeThrouThdOk

`func (o *QosSustainabilityInfo) GetRanUeThrouThdOk() (*map[string]interface{}, bool)`

GetRanUeThrouThdOk returns a tuple with the RanUeThrouThd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUeThrouThd

`func (o *QosSustainabilityInfo) SetRanUeThrouThd(v map[string]interface{})`

SetRanUeThrouThd sets RanUeThrouThd field to given value.

### HasRanUeThrouThd

`func (o *QosSustainabilityInfo) HasRanUeThrouThd() bool`

HasRanUeThrouThd returns a boolean if a field has been set.

### GetSnssai

`func (o *QosSustainabilityInfo) GetSnssai() map[string]interface{}`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *QosSustainabilityInfo) GetSnssaiOk() (*map[string]interface{}, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *QosSustainabilityInfo) SetSnssai(v map[string]interface{})`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *QosSustainabilityInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetConfidence

`func (o *QosSustainabilityInfo) GetConfidence() map[string]interface{}`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *QosSustainabilityInfo) GetConfidenceOk() (*map[string]interface{}, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *QosSustainabilityInfo) SetConfidence(v map[string]interface{})`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *QosSustainabilityInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


