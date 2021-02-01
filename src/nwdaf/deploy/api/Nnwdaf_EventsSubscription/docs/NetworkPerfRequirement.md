# NetworkPerfRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NwPerfType** | [**NetworkPerfType**](NetworkPerfType.md) |  | 
**RelativeRatio** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**AbsoluteNum** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 

## Methods

### NewNetworkPerfRequirement

`func NewNetworkPerfRequirement(nwPerfType NetworkPerfType, ) *NetworkPerfRequirement`

NewNetworkPerfRequirement instantiates a new NetworkPerfRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPerfRequirementWithDefaults

`func NewNetworkPerfRequirementWithDefaults() *NetworkPerfRequirement`

NewNetworkPerfRequirementWithDefaults instantiates a new NetworkPerfRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNwPerfType

`func (o *NetworkPerfRequirement) GetNwPerfType() NetworkPerfType`

GetNwPerfType returns the NwPerfType field if non-nil, zero value otherwise.

### GetNwPerfTypeOk

`func (o *NetworkPerfRequirement) GetNwPerfTypeOk() (*NetworkPerfType, bool)`

GetNwPerfTypeOk returns a tuple with the NwPerfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfType

`func (o *NetworkPerfRequirement) SetNwPerfType(v NetworkPerfType)`

SetNwPerfType sets NwPerfType field to given value.


### GetRelativeRatio

`func (o *NetworkPerfRequirement) GetRelativeRatio() map[string]interface{}`

GetRelativeRatio returns the RelativeRatio field if non-nil, zero value otherwise.

### GetRelativeRatioOk

`func (o *NetworkPerfRequirement) GetRelativeRatioOk() (*map[string]interface{}, bool)`

GetRelativeRatioOk returns a tuple with the RelativeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeRatio

`func (o *NetworkPerfRequirement) SetRelativeRatio(v map[string]interface{})`

SetRelativeRatio sets RelativeRatio field to given value.

### HasRelativeRatio

`func (o *NetworkPerfRequirement) HasRelativeRatio() bool`

HasRelativeRatio returns a boolean if a field has been set.

### GetAbsoluteNum

`func (o *NetworkPerfRequirement) GetAbsoluteNum() map[string]interface{}`

GetAbsoluteNum returns the AbsoluteNum field if non-nil, zero value otherwise.

### GetAbsoluteNumOk

`func (o *NetworkPerfRequirement) GetAbsoluteNumOk() (*map[string]interface{}, bool)`

GetAbsoluteNumOk returns a tuple with the AbsoluteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteNum

`func (o *NetworkPerfRequirement) SetAbsoluteNum(v map[string]interface{})`

SetAbsoluteNum sets AbsoluteNum field to given value.

### HasAbsoluteNum

`func (o *NetworkPerfRequirement) HasAbsoluteNum() bool`

HasAbsoluteNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


