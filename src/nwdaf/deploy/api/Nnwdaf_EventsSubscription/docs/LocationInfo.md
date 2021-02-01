# LocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loc** | [**map[string]interface{}**](object.md) |  | 
**Ratio** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**Confidence** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 

## Methods

### NewLocationInfo

`func NewLocationInfo(loc map[string]interface{}, ) *LocationInfo`

NewLocationInfo instantiates a new LocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationInfoWithDefaults

`func NewLocationInfoWithDefaults() *LocationInfo`

NewLocationInfoWithDefaults instantiates a new LocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoc

`func (o *LocationInfo) GetLoc() map[string]interface{}`

GetLoc returns the Loc field if non-nil, zero value otherwise.

### GetLocOk

`func (o *LocationInfo) GetLocOk() (*map[string]interface{}, bool)`

GetLocOk returns a tuple with the Loc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoc

`func (o *LocationInfo) SetLoc(v map[string]interface{})`

SetLoc sets Loc field to given value.


### GetRatio

`func (o *LocationInfo) GetRatio() map[string]interface{}`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *LocationInfo) GetRatioOk() (*map[string]interface{}, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *LocationInfo) SetRatio(v map[string]interface{})`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *LocationInfo) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetConfidence

`func (o *LocationInfo) GetConfidence() map[string]interface{}`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *LocationInfo) GetConfidenceOk() (*map[string]interface{}, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *LocationInfo) SetConfidence(v map[string]interface{})`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *LocationInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


