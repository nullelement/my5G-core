# NetworkPerfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkArea** | [**map[string]interface{}**](object.md) |  | 
**NwPerfType** | [**NetworkPerfType**](NetworkPerfType.md) |  | 
**RelativeRatio** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**AbsoluteNum** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**Confidence** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 

## Methods

### NewNetworkPerfInfo

`func NewNetworkPerfInfo(networkArea map[string]interface{}, nwPerfType NetworkPerfType, ) *NetworkPerfInfo`

NewNetworkPerfInfo instantiates a new NetworkPerfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPerfInfoWithDefaults

`func NewNetworkPerfInfoWithDefaults() *NetworkPerfInfo`

NewNetworkPerfInfoWithDefaults instantiates a new NetworkPerfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkArea

`func (o *NetworkPerfInfo) GetNetworkArea() map[string]interface{}`

GetNetworkArea returns the NetworkArea field if non-nil, zero value otherwise.

### GetNetworkAreaOk

`func (o *NetworkPerfInfo) GetNetworkAreaOk() (*map[string]interface{}, bool)`

GetNetworkAreaOk returns a tuple with the NetworkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkArea

`func (o *NetworkPerfInfo) SetNetworkArea(v map[string]interface{})`

SetNetworkArea sets NetworkArea field to given value.


### GetNwPerfType

`func (o *NetworkPerfInfo) GetNwPerfType() NetworkPerfType`

GetNwPerfType returns the NwPerfType field if non-nil, zero value otherwise.

### GetNwPerfTypeOk

`func (o *NetworkPerfInfo) GetNwPerfTypeOk() (*NetworkPerfType, bool)`

GetNwPerfTypeOk returns a tuple with the NwPerfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfType

`func (o *NetworkPerfInfo) SetNwPerfType(v NetworkPerfType)`

SetNwPerfType sets NwPerfType field to given value.


### GetRelativeRatio

`func (o *NetworkPerfInfo) GetRelativeRatio() map[string]interface{}`

GetRelativeRatio returns the RelativeRatio field if non-nil, zero value otherwise.

### GetRelativeRatioOk

`func (o *NetworkPerfInfo) GetRelativeRatioOk() (*map[string]interface{}, bool)`

GetRelativeRatioOk returns a tuple with the RelativeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeRatio

`func (o *NetworkPerfInfo) SetRelativeRatio(v map[string]interface{})`

SetRelativeRatio sets RelativeRatio field to given value.

### HasRelativeRatio

`func (o *NetworkPerfInfo) HasRelativeRatio() bool`

HasRelativeRatio returns a boolean if a field has been set.

### GetAbsoluteNum

`func (o *NetworkPerfInfo) GetAbsoluteNum() map[string]interface{}`

GetAbsoluteNum returns the AbsoluteNum field if non-nil, zero value otherwise.

### GetAbsoluteNumOk

`func (o *NetworkPerfInfo) GetAbsoluteNumOk() (*map[string]interface{}, bool)`

GetAbsoluteNumOk returns a tuple with the AbsoluteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteNum

`func (o *NetworkPerfInfo) SetAbsoluteNum(v map[string]interface{})`

SetAbsoluteNum sets AbsoluteNum field to given value.

### HasAbsoluteNum

`func (o *NetworkPerfInfo) HasAbsoluteNum() bool`

HasAbsoluteNum returns a boolean if a field has been set.

### GetConfidence

`func (o *NetworkPerfInfo) GetConfidence() map[string]interface{}`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *NetworkPerfInfo) GetConfidenceOk() (*map[string]interface{}, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *NetworkPerfInfo) SetConfidence(v map[string]interface{})`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *NetworkPerfInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


