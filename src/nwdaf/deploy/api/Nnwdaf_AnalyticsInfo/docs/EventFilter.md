# EventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnySlice** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**Snssais** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | Identification(s) of network slice to which the subscription belongs. | [optional] 
**AppIds** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**Dnns** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**Dnais** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**NetworkArea** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**NfInstanceIds** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**NfSetIds** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**NfTypes** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**NsiIdInfos** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**QosRequ** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**NwPerfTypes** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**BwRequs** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**ExcepIds** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**ExptAnaType** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**ExptUeBehav** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 

## Methods

### NewEventFilter

`func NewEventFilter() *EventFilter`

NewEventFilter instantiates a new EventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventFilterWithDefaults

`func NewEventFilterWithDefaults() *EventFilter`

NewEventFilterWithDefaults instantiates a new EventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnySlice

`func (o *EventFilter) GetAnySlice() map[string]interface{}`

GetAnySlice returns the AnySlice field if non-nil, zero value otherwise.

### GetAnySliceOk

`func (o *EventFilter) GetAnySliceOk() (*map[string]interface{}, bool)`

GetAnySliceOk returns a tuple with the AnySlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnySlice

`func (o *EventFilter) SetAnySlice(v map[string]interface{})`

SetAnySlice sets AnySlice field to given value.

### HasAnySlice

`func (o *EventFilter) HasAnySlice() bool`

HasAnySlice returns a boolean if a field has been set.

### GetSnssais

`func (o *EventFilter) GetSnssais() []map[string]interface{}`

GetSnssais returns the Snssais field if non-nil, zero value otherwise.

### GetSnssaisOk

`func (o *EventFilter) GetSnssaisOk() (*[]map[string]interface{}, bool)`

GetSnssaisOk returns a tuple with the Snssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssais

`func (o *EventFilter) SetSnssais(v []map[string]interface{})`

SetSnssais sets Snssais field to given value.

### HasSnssais

`func (o *EventFilter) HasSnssais() bool`

HasSnssais returns a boolean if a field has been set.

### GetAppIds

`func (o *EventFilter) GetAppIds() []map[string]interface{}`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *EventFilter) GetAppIdsOk() (*[]map[string]interface{}, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *EventFilter) SetAppIds(v []map[string]interface{})`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *EventFilter) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetDnns

`func (o *EventFilter) GetDnns() []map[string]interface{}`

GetDnns returns the Dnns field if non-nil, zero value otherwise.

### GetDnnsOk

`func (o *EventFilter) GetDnnsOk() (*[]map[string]interface{}, bool)`

GetDnnsOk returns a tuple with the Dnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnns

`func (o *EventFilter) SetDnns(v []map[string]interface{})`

SetDnns sets Dnns field to given value.

### HasDnns

`func (o *EventFilter) HasDnns() bool`

HasDnns returns a boolean if a field has been set.

### GetDnais

`func (o *EventFilter) GetDnais() []map[string]interface{}`

GetDnais returns the Dnais field if non-nil, zero value otherwise.

### GetDnaisOk

`func (o *EventFilter) GetDnaisOk() (*[]map[string]interface{}, bool)`

GetDnaisOk returns a tuple with the Dnais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnais

`func (o *EventFilter) SetDnais(v []map[string]interface{})`

SetDnais sets Dnais field to given value.

### HasDnais

`func (o *EventFilter) HasDnais() bool`

HasDnais returns a boolean if a field has been set.

### GetNetworkArea

`func (o *EventFilter) GetNetworkArea() map[string]interface{}`

GetNetworkArea returns the NetworkArea field if non-nil, zero value otherwise.

### GetNetworkAreaOk

`func (o *EventFilter) GetNetworkAreaOk() (*map[string]interface{}, bool)`

GetNetworkAreaOk returns a tuple with the NetworkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkArea

`func (o *EventFilter) SetNetworkArea(v map[string]interface{})`

SetNetworkArea sets NetworkArea field to given value.

### HasNetworkArea

`func (o *EventFilter) HasNetworkArea() bool`

HasNetworkArea returns a boolean if a field has been set.

### GetNfInstanceIds

`func (o *EventFilter) GetNfInstanceIds() []map[string]interface{}`

GetNfInstanceIds returns the NfInstanceIds field if non-nil, zero value otherwise.

### GetNfInstanceIdsOk

`func (o *EventFilter) GetNfInstanceIdsOk() (*[]map[string]interface{}, bool)`

GetNfInstanceIdsOk returns a tuple with the NfInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceIds

`func (o *EventFilter) SetNfInstanceIds(v []map[string]interface{})`

SetNfInstanceIds sets NfInstanceIds field to given value.

### HasNfInstanceIds

`func (o *EventFilter) HasNfInstanceIds() bool`

HasNfInstanceIds returns a boolean if a field has been set.

### GetNfSetIds

`func (o *EventFilter) GetNfSetIds() []map[string]interface{}`

GetNfSetIds returns the NfSetIds field if non-nil, zero value otherwise.

### GetNfSetIdsOk

`func (o *EventFilter) GetNfSetIdsOk() (*[]map[string]interface{}, bool)`

GetNfSetIdsOk returns a tuple with the NfSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetIds

`func (o *EventFilter) SetNfSetIds(v []map[string]interface{})`

SetNfSetIds sets NfSetIds field to given value.

### HasNfSetIds

`func (o *EventFilter) HasNfSetIds() bool`

HasNfSetIds returns a boolean if a field has been set.

### GetNfTypes

`func (o *EventFilter) GetNfTypes() []map[string]interface{}`

GetNfTypes returns the NfTypes field if non-nil, zero value otherwise.

### GetNfTypesOk

`func (o *EventFilter) GetNfTypesOk() (*[]map[string]interface{}, bool)`

GetNfTypesOk returns a tuple with the NfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfTypes

`func (o *EventFilter) SetNfTypes(v []map[string]interface{})`

SetNfTypes sets NfTypes field to given value.

### HasNfTypes

`func (o *EventFilter) HasNfTypes() bool`

HasNfTypes returns a boolean if a field has been set.

### GetNsiIdInfos

`func (o *EventFilter) GetNsiIdInfos() []map[string]interface{}`

GetNsiIdInfos returns the NsiIdInfos field if non-nil, zero value otherwise.

### GetNsiIdInfosOk

`func (o *EventFilter) GetNsiIdInfosOk() (*[]map[string]interface{}, bool)`

GetNsiIdInfosOk returns a tuple with the NsiIdInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiIdInfos

`func (o *EventFilter) SetNsiIdInfos(v []map[string]interface{})`

SetNsiIdInfos sets NsiIdInfos field to given value.

### HasNsiIdInfos

`func (o *EventFilter) HasNsiIdInfos() bool`

HasNsiIdInfos returns a boolean if a field has been set.

### GetQosRequ

`func (o *EventFilter) GetQosRequ() map[string]interface{}`

GetQosRequ returns the QosRequ field if non-nil, zero value otherwise.

### GetQosRequOk

`func (o *EventFilter) GetQosRequOk() (*map[string]interface{}, bool)`

GetQosRequOk returns a tuple with the QosRequ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosRequ

`func (o *EventFilter) SetQosRequ(v map[string]interface{})`

SetQosRequ sets QosRequ field to given value.

### HasQosRequ

`func (o *EventFilter) HasQosRequ() bool`

HasQosRequ returns a boolean if a field has been set.

### GetNwPerfTypes

`func (o *EventFilter) GetNwPerfTypes() []map[string]interface{}`

GetNwPerfTypes returns the NwPerfTypes field if non-nil, zero value otherwise.

### GetNwPerfTypesOk

`func (o *EventFilter) GetNwPerfTypesOk() (*[]map[string]interface{}, bool)`

GetNwPerfTypesOk returns a tuple with the NwPerfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfTypes

`func (o *EventFilter) SetNwPerfTypes(v []map[string]interface{})`

SetNwPerfTypes sets NwPerfTypes field to given value.

### HasNwPerfTypes

`func (o *EventFilter) HasNwPerfTypes() bool`

HasNwPerfTypes returns a boolean if a field has been set.

### GetBwRequs

`func (o *EventFilter) GetBwRequs() []map[string]interface{}`

GetBwRequs returns the BwRequs field if non-nil, zero value otherwise.

### GetBwRequsOk

`func (o *EventFilter) GetBwRequsOk() (*[]map[string]interface{}, bool)`

GetBwRequsOk returns a tuple with the BwRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwRequs

`func (o *EventFilter) SetBwRequs(v []map[string]interface{})`

SetBwRequs sets BwRequs field to given value.

### HasBwRequs

`func (o *EventFilter) HasBwRequs() bool`

HasBwRequs returns a boolean if a field has been set.

### GetExcepIds

`func (o *EventFilter) GetExcepIds() []map[string]interface{}`

GetExcepIds returns the ExcepIds field if non-nil, zero value otherwise.

### GetExcepIdsOk

`func (o *EventFilter) GetExcepIdsOk() (*[]map[string]interface{}, bool)`

GetExcepIdsOk returns a tuple with the ExcepIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepIds

`func (o *EventFilter) SetExcepIds(v []map[string]interface{})`

SetExcepIds sets ExcepIds field to given value.

### HasExcepIds

`func (o *EventFilter) HasExcepIds() bool`

HasExcepIds returns a boolean if a field has been set.

### GetExptAnaType

`func (o *EventFilter) GetExptAnaType() map[string]interface{}`

GetExptAnaType returns the ExptAnaType field if non-nil, zero value otherwise.

### GetExptAnaTypeOk

`func (o *EventFilter) GetExptAnaTypeOk() (*map[string]interface{}, bool)`

GetExptAnaTypeOk returns a tuple with the ExptAnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptAnaType

`func (o *EventFilter) SetExptAnaType(v map[string]interface{})`

SetExptAnaType sets ExptAnaType field to given value.

### HasExptAnaType

`func (o *EventFilter) HasExptAnaType() bool`

HasExptAnaType returns a boolean if a field has been set.

### GetExptUeBehav

`func (o *EventFilter) GetExptUeBehav() map[string]interface{}`

GetExptUeBehav returns the ExptUeBehav field if non-nil, zero value otherwise.

### GetExptUeBehavOk

`func (o *EventFilter) GetExptUeBehavOk() (*map[string]interface{}, bool)`

GetExptUeBehavOk returns a tuple with the ExptUeBehav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptUeBehav

`func (o *EventFilter) SetExptUeBehav(v map[string]interface{})`

SetExptUeBehav sets ExptUeBehav field to given value.

### HasExptUeBehav

`func (o *EventFilter) HasExptUeBehav() bool`

HasExptUeBehav returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


