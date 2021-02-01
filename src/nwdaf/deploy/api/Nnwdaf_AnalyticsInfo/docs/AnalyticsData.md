# AnalyticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**Expiry** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**TimeStampGen** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**SliceLoadLevelInfos** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | The slices and their load level information. | [optional] 
**NsiLoadLevelInfos** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**NfLoadLevelInfos** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**NwPerfs** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**SvcExps** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**QosSustainInfos** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**UeMobs** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**UeComms** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**UserDataCongInfos** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**AbnorBehavrs** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**SuppFeat** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 

## Methods

### NewAnalyticsData

`func NewAnalyticsData() *AnalyticsData`

NewAnalyticsData instantiates a new AnalyticsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsDataWithDefaults

`func NewAnalyticsDataWithDefaults() *AnalyticsData`

NewAnalyticsDataWithDefaults instantiates a new AnalyticsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *AnalyticsData) GetStart() map[string]interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AnalyticsData) GetStartOk() (*map[string]interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AnalyticsData) SetStart(v map[string]interface{})`

SetStart sets Start field to given value.

### HasStart

`func (o *AnalyticsData) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetExpiry

`func (o *AnalyticsData) GetExpiry() map[string]interface{}`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AnalyticsData) GetExpiryOk() (*map[string]interface{}, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AnalyticsData) SetExpiry(v map[string]interface{})`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AnalyticsData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetTimeStampGen

`func (o *AnalyticsData) GetTimeStampGen() map[string]interface{}`

GetTimeStampGen returns the TimeStampGen field if non-nil, zero value otherwise.

### GetTimeStampGenOk

`func (o *AnalyticsData) GetTimeStampGenOk() (*map[string]interface{}, bool)`

GetTimeStampGenOk returns a tuple with the TimeStampGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStampGen

`func (o *AnalyticsData) SetTimeStampGen(v map[string]interface{})`

SetTimeStampGen sets TimeStampGen field to given value.

### HasTimeStampGen

`func (o *AnalyticsData) HasTimeStampGen() bool`

HasTimeStampGen returns a boolean if a field has been set.

### GetSliceLoadLevelInfos

`func (o *AnalyticsData) GetSliceLoadLevelInfos() []map[string]interface{}`

GetSliceLoadLevelInfos returns the SliceLoadLevelInfos field if non-nil, zero value otherwise.

### GetSliceLoadLevelInfosOk

`func (o *AnalyticsData) GetSliceLoadLevelInfosOk() (*[]map[string]interface{}, bool)`

GetSliceLoadLevelInfosOk returns a tuple with the SliceLoadLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceLoadLevelInfos

`func (o *AnalyticsData) SetSliceLoadLevelInfos(v []map[string]interface{})`

SetSliceLoadLevelInfos sets SliceLoadLevelInfos field to given value.

### HasSliceLoadLevelInfos

`func (o *AnalyticsData) HasSliceLoadLevelInfos() bool`

HasSliceLoadLevelInfos returns a boolean if a field has been set.

### GetNsiLoadLevelInfos

`func (o *AnalyticsData) GetNsiLoadLevelInfos() []map[string]interface{}`

GetNsiLoadLevelInfos returns the NsiLoadLevelInfos field if non-nil, zero value otherwise.

### GetNsiLoadLevelInfosOk

`func (o *AnalyticsData) GetNsiLoadLevelInfosOk() (*[]map[string]interface{}, bool)`

GetNsiLoadLevelInfosOk returns a tuple with the NsiLoadLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiLoadLevelInfos

`func (o *AnalyticsData) SetNsiLoadLevelInfos(v []map[string]interface{})`

SetNsiLoadLevelInfos sets NsiLoadLevelInfos field to given value.

### HasNsiLoadLevelInfos

`func (o *AnalyticsData) HasNsiLoadLevelInfos() bool`

HasNsiLoadLevelInfos returns a boolean if a field has been set.

### GetNfLoadLevelInfos

`func (o *AnalyticsData) GetNfLoadLevelInfos() []map[string]interface{}`

GetNfLoadLevelInfos returns the NfLoadLevelInfos field if non-nil, zero value otherwise.

### GetNfLoadLevelInfosOk

`func (o *AnalyticsData) GetNfLoadLevelInfosOk() (*[]map[string]interface{}, bool)`

GetNfLoadLevelInfosOk returns a tuple with the NfLoadLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfLoadLevelInfos

`func (o *AnalyticsData) SetNfLoadLevelInfos(v []map[string]interface{})`

SetNfLoadLevelInfos sets NfLoadLevelInfos field to given value.

### HasNfLoadLevelInfos

`func (o *AnalyticsData) HasNfLoadLevelInfos() bool`

HasNfLoadLevelInfos returns a boolean if a field has been set.

### GetNwPerfs

`func (o *AnalyticsData) GetNwPerfs() []map[string]interface{}`

GetNwPerfs returns the NwPerfs field if non-nil, zero value otherwise.

### GetNwPerfsOk

`func (o *AnalyticsData) GetNwPerfsOk() (*[]map[string]interface{}, bool)`

GetNwPerfsOk returns a tuple with the NwPerfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfs

`func (o *AnalyticsData) SetNwPerfs(v []map[string]interface{})`

SetNwPerfs sets NwPerfs field to given value.

### HasNwPerfs

`func (o *AnalyticsData) HasNwPerfs() bool`

HasNwPerfs returns a boolean if a field has been set.

### GetSvcExps

`func (o *AnalyticsData) GetSvcExps() []map[string]interface{}`

GetSvcExps returns the SvcExps field if non-nil, zero value otherwise.

### GetSvcExpsOk

`func (o *AnalyticsData) GetSvcExpsOk() (*[]map[string]interface{}, bool)`

GetSvcExpsOk returns a tuple with the SvcExps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExps

`func (o *AnalyticsData) SetSvcExps(v []map[string]interface{})`

SetSvcExps sets SvcExps field to given value.

### HasSvcExps

`func (o *AnalyticsData) HasSvcExps() bool`

HasSvcExps returns a boolean if a field has been set.

### GetQosSustainInfos

`func (o *AnalyticsData) GetQosSustainInfos() []map[string]interface{}`

GetQosSustainInfos returns the QosSustainInfos field if non-nil, zero value otherwise.

### GetQosSustainInfosOk

`func (o *AnalyticsData) GetQosSustainInfosOk() (*[]map[string]interface{}, bool)`

GetQosSustainInfosOk returns a tuple with the QosSustainInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSustainInfos

`func (o *AnalyticsData) SetQosSustainInfos(v []map[string]interface{})`

SetQosSustainInfos sets QosSustainInfos field to given value.

### HasQosSustainInfos

`func (o *AnalyticsData) HasQosSustainInfos() bool`

HasQosSustainInfos returns a boolean if a field has been set.

### GetUeMobs

`func (o *AnalyticsData) GetUeMobs() []map[string]interface{}`

GetUeMobs returns the UeMobs field if non-nil, zero value otherwise.

### GetUeMobsOk

`func (o *AnalyticsData) GetUeMobsOk() (*[]map[string]interface{}, bool)`

GetUeMobsOk returns a tuple with the UeMobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMobs

`func (o *AnalyticsData) SetUeMobs(v []map[string]interface{})`

SetUeMobs sets UeMobs field to given value.

### HasUeMobs

`func (o *AnalyticsData) HasUeMobs() bool`

HasUeMobs returns a boolean if a field has been set.

### GetUeComms

`func (o *AnalyticsData) GetUeComms() []map[string]interface{}`

GetUeComms returns the UeComms field if non-nil, zero value otherwise.

### GetUeCommsOk

`func (o *AnalyticsData) GetUeCommsOk() (*[]map[string]interface{}, bool)`

GetUeCommsOk returns a tuple with the UeComms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeComms

`func (o *AnalyticsData) SetUeComms(v []map[string]interface{})`

SetUeComms sets UeComms field to given value.

### HasUeComms

`func (o *AnalyticsData) HasUeComms() bool`

HasUeComms returns a boolean if a field has been set.

### GetUserDataCongInfos

`func (o *AnalyticsData) GetUserDataCongInfos() []map[string]interface{}`

GetUserDataCongInfos returns the UserDataCongInfos field if non-nil, zero value otherwise.

### GetUserDataCongInfosOk

`func (o *AnalyticsData) GetUserDataCongInfosOk() (*[]map[string]interface{}, bool)`

GetUserDataCongInfosOk returns a tuple with the UserDataCongInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataCongInfos

`func (o *AnalyticsData) SetUserDataCongInfos(v []map[string]interface{})`

SetUserDataCongInfos sets UserDataCongInfos field to given value.

### HasUserDataCongInfos

`func (o *AnalyticsData) HasUserDataCongInfos() bool`

HasUserDataCongInfos returns a boolean if a field has been set.

### GetAbnorBehavrs

`func (o *AnalyticsData) GetAbnorBehavrs() []map[string]interface{}`

GetAbnorBehavrs returns the AbnorBehavrs field if non-nil, zero value otherwise.

### GetAbnorBehavrsOk

`func (o *AnalyticsData) GetAbnorBehavrsOk() (*[]map[string]interface{}, bool)`

GetAbnorBehavrsOk returns a tuple with the AbnorBehavrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnorBehavrs

`func (o *AnalyticsData) SetAbnorBehavrs(v []map[string]interface{})`

SetAbnorBehavrs sets AbnorBehavrs field to given value.

### HasAbnorBehavrs

`func (o *AnalyticsData) HasAbnorBehavrs() bool`

HasAbnorBehavrs returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AnalyticsData) GetSuppFeat() map[string]interface{}`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AnalyticsData) GetSuppFeatOk() (*map[string]interface{}, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AnalyticsData) SetSuppFeat(v map[string]interface{})`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AnalyticsData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


