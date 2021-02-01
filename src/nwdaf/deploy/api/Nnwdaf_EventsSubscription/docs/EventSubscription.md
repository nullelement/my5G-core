# EventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnySlice** | Pointer to **bool** | FALSE represents not applicable for all slices. TRUE represents applicable for all slices. | [optional] 
**AppIds** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | Identification(s) of application to which the subscription applies. | [optional] 
**Dnns** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | Identification(s) of DNN to which the subscription applies. | [optional] 
**Dnais** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**Event** | [**NwdafEvent**](NwdafEvent.md) |  | 
**ExtraReportReq** | Pointer to [**EventReportingRequirement**](EventReportingRequirement.md) |  | [optional] 
**LoadLevelThreshold** | Pointer to **int32** | Indicates that the NWDAF shall report the corresponding network slice load level to the NF service consumer where the load level of the network slice instance identified by snssais is reached. | [optional] 
**NotificationMethod** | Pointer to [**NotificationMethod**](NotificationMethod.md) |  | [optional] 
**MatchingDir** | Pointer to [**MatchingDirection**](MatchingDirection.md) |  | [optional] 
**NfLoadLvlThds** | Pointer to [**[]ThresholdLevel**](ThresholdLevel.md) | Shall be supplied in order to start reporting when an average load level is reached. | [optional] 
**NfInstanceIds** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**NfSetIds** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**NfTypes** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**NetworkArea** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**NsiIdInfos** | Pointer to [**[]NsiIdInfo**](NsiIdInfo.md) |  | [optional] 
**NsiLevelThrds** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**QosRequ** | Pointer to [**QosRequirement**](QosRequirement.md) |  | [optional] 
**QosFlowRetThds** | Pointer to [**[]RetainabilityThreshold**](RetainabilityThreshold.md) |  | [optional] 
**RanUeThrouThds** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**RepetitionPeriod** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 
**Snssaia** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | Identification(s) of network slice to which the subscription applies. When subscribed event is \&quot;SLICE_LOAD_LEVEL\&quot;, either information about slice(s) identified by snssai, or anySlice set to \&quot;TRUE\&quot; shall be included. It corresponds to snssais in the data model definition of 3GPP TS 29.520. When subscribed is “QOS_SUSTAINABILITY”, the identifications of network slices is optional. | [optional] 
**TgtUe** | Pointer to [**TargetUeInformation**](TargetUeInformation.md) |  | [optional] 
**CongThresholds** | Pointer to [**[]ThresholdLevel**](ThresholdLevel.md) |  | [optional] 
**NwPerfRequs** | Pointer to [**[]NetworkPerfRequirement**](NetworkPerfRequirement.md) |  | [optional] 
**BwRequs** | Pointer to [**[]BwRequirement**](BwRequirement.md) |  | [optional] 
**ExcepRequs** | Pointer to [**[]Exception**](Exception.md) |  | [optional] 
**ExptAnaType** | Pointer to [**ExpectedAnalyticsType**](ExpectedAnalyticsType.md) |  | [optional] 
**ExptUeBehav** | Pointer to [**map[string]interface{}**](object.md) |  | [optional] 

## Methods

### NewEventSubscription

`func NewEventSubscription(event NwdafEvent, ) *EventSubscription`

NewEventSubscription instantiates a new EventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionWithDefaults

`func NewEventSubscriptionWithDefaults() *EventSubscription`

NewEventSubscriptionWithDefaults instantiates a new EventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnySlice

`func (o *EventSubscription) GetAnySlice() bool`

GetAnySlice returns the AnySlice field if non-nil, zero value otherwise.

### GetAnySliceOk

`func (o *EventSubscription) GetAnySliceOk() (*bool, bool)`

GetAnySliceOk returns a tuple with the AnySlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnySlice

`func (o *EventSubscription) SetAnySlice(v bool)`

SetAnySlice sets AnySlice field to given value.

### HasAnySlice

`func (o *EventSubscription) HasAnySlice() bool`

HasAnySlice returns a boolean if a field has been set.

### GetAppIds

`func (o *EventSubscription) GetAppIds() []map[string]interface{}`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *EventSubscription) GetAppIdsOk() (*[]map[string]interface{}, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *EventSubscription) SetAppIds(v []map[string]interface{})`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *EventSubscription) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetDnns

`func (o *EventSubscription) GetDnns() []map[string]interface{}`

GetDnns returns the Dnns field if non-nil, zero value otherwise.

### GetDnnsOk

`func (o *EventSubscription) GetDnnsOk() (*[]map[string]interface{}, bool)`

GetDnnsOk returns a tuple with the Dnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnns

`func (o *EventSubscription) SetDnns(v []map[string]interface{})`

SetDnns sets Dnns field to given value.

### HasDnns

`func (o *EventSubscription) HasDnns() bool`

HasDnns returns a boolean if a field has been set.

### GetDnais

`func (o *EventSubscription) GetDnais() []map[string]interface{}`

GetDnais returns the Dnais field if non-nil, zero value otherwise.

### GetDnaisOk

`func (o *EventSubscription) GetDnaisOk() (*[]map[string]interface{}, bool)`

GetDnaisOk returns a tuple with the Dnais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnais

`func (o *EventSubscription) SetDnais(v []map[string]interface{})`

SetDnais sets Dnais field to given value.

### HasDnais

`func (o *EventSubscription) HasDnais() bool`

HasDnais returns a boolean if a field has been set.

### GetEvent

`func (o *EventSubscription) GetEvent() NwdafEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventSubscription) GetEventOk() (*NwdafEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventSubscription) SetEvent(v NwdafEvent)`

SetEvent sets Event field to given value.


### GetExtraReportReq

`func (o *EventSubscription) GetExtraReportReq() EventReportingRequirement`

GetExtraReportReq returns the ExtraReportReq field if non-nil, zero value otherwise.

### GetExtraReportReqOk

`func (o *EventSubscription) GetExtraReportReqOk() (*EventReportingRequirement, bool)`

GetExtraReportReqOk returns a tuple with the ExtraReportReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraReportReq

`func (o *EventSubscription) SetExtraReportReq(v EventReportingRequirement)`

SetExtraReportReq sets ExtraReportReq field to given value.

### HasExtraReportReq

`func (o *EventSubscription) HasExtraReportReq() bool`

HasExtraReportReq returns a boolean if a field has been set.

### GetLoadLevelThreshold

`func (o *EventSubscription) GetLoadLevelThreshold() int32`

GetLoadLevelThreshold returns the LoadLevelThreshold field if non-nil, zero value otherwise.

### GetLoadLevelThresholdOk

`func (o *EventSubscription) GetLoadLevelThresholdOk() (*int32, bool)`

GetLoadLevelThresholdOk returns a tuple with the LoadLevelThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadLevelThreshold

`func (o *EventSubscription) SetLoadLevelThreshold(v int32)`

SetLoadLevelThreshold sets LoadLevelThreshold field to given value.

### HasLoadLevelThreshold

`func (o *EventSubscription) HasLoadLevelThreshold() bool`

HasLoadLevelThreshold returns a boolean if a field has been set.

### GetNotificationMethod

`func (o *EventSubscription) GetNotificationMethod() NotificationMethod`

GetNotificationMethod returns the NotificationMethod field if non-nil, zero value otherwise.

### GetNotificationMethodOk

`func (o *EventSubscription) GetNotificationMethodOk() (*NotificationMethod, bool)`

GetNotificationMethodOk returns a tuple with the NotificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMethod

`func (o *EventSubscription) SetNotificationMethod(v NotificationMethod)`

SetNotificationMethod sets NotificationMethod field to given value.

### HasNotificationMethod

`func (o *EventSubscription) HasNotificationMethod() bool`

HasNotificationMethod returns a boolean if a field has been set.

### GetMatchingDir

`func (o *EventSubscription) GetMatchingDir() MatchingDirection`

GetMatchingDir returns the MatchingDir field if non-nil, zero value otherwise.

### GetMatchingDirOk

`func (o *EventSubscription) GetMatchingDirOk() (*MatchingDirection, bool)`

GetMatchingDirOk returns a tuple with the MatchingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingDir

`func (o *EventSubscription) SetMatchingDir(v MatchingDirection)`

SetMatchingDir sets MatchingDir field to given value.

### HasMatchingDir

`func (o *EventSubscription) HasMatchingDir() bool`

HasMatchingDir returns a boolean if a field has been set.

### GetNfLoadLvlThds

`func (o *EventSubscription) GetNfLoadLvlThds() []ThresholdLevel`

GetNfLoadLvlThds returns the NfLoadLvlThds field if non-nil, zero value otherwise.

### GetNfLoadLvlThdsOk

`func (o *EventSubscription) GetNfLoadLvlThdsOk() (*[]ThresholdLevel, bool)`

GetNfLoadLvlThdsOk returns a tuple with the NfLoadLvlThds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfLoadLvlThds

`func (o *EventSubscription) SetNfLoadLvlThds(v []ThresholdLevel)`

SetNfLoadLvlThds sets NfLoadLvlThds field to given value.

### HasNfLoadLvlThds

`func (o *EventSubscription) HasNfLoadLvlThds() bool`

HasNfLoadLvlThds returns a boolean if a field has been set.

### GetNfInstanceIds

`func (o *EventSubscription) GetNfInstanceIds() []map[string]interface{}`

GetNfInstanceIds returns the NfInstanceIds field if non-nil, zero value otherwise.

### GetNfInstanceIdsOk

`func (o *EventSubscription) GetNfInstanceIdsOk() (*[]map[string]interface{}, bool)`

GetNfInstanceIdsOk returns a tuple with the NfInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceIds

`func (o *EventSubscription) SetNfInstanceIds(v []map[string]interface{})`

SetNfInstanceIds sets NfInstanceIds field to given value.

### HasNfInstanceIds

`func (o *EventSubscription) HasNfInstanceIds() bool`

HasNfInstanceIds returns a boolean if a field has been set.

### GetNfSetIds

`func (o *EventSubscription) GetNfSetIds() []map[string]interface{}`

GetNfSetIds returns the NfSetIds field if non-nil, zero value otherwise.

### GetNfSetIdsOk

`func (o *EventSubscription) GetNfSetIdsOk() (*[]map[string]interface{}, bool)`

GetNfSetIdsOk returns a tuple with the NfSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetIds

`func (o *EventSubscription) SetNfSetIds(v []map[string]interface{})`

SetNfSetIds sets NfSetIds field to given value.

### HasNfSetIds

`func (o *EventSubscription) HasNfSetIds() bool`

HasNfSetIds returns a boolean if a field has been set.

### GetNfTypes

`func (o *EventSubscription) GetNfTypes() []map[string]interface{}`

GetNfTypes returns the NfTypes field if non-nil, zero value otherwise.

### GetNfTypesOk

`func (o *EventSubscription) GetNfTypesOk() (*[]map[string]interface{}, bool)`

GetNfTypesOk returns a tuple with the NfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfTypes

`func (o *EventSubscription) SetNfTypes(v []map[string]interface{})`

SetNfTypes sets NfTypes field to given value.

### HasNfTypes

`func (o *EventSubscription) HasNfTypes() bool`

HasNfTypes returns a boolean if a field has been set.

### GetNetworkArea

`func (o *EventSubscription) GetNetworkArea() map[string]interface{}`

GetNetworkArea returns the NetworkArea field if non-nil, zero value otherwise.

### GetNetworkAreaOk

`func (o *EventSubscription) GetNetworkAreaOk() (*map[string]interface{}, bool)`

GetNetworkAreaOk returns a tuple with the NetworkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkArea

`func (o *EventSubscription) SetNetworkArea(v map[string]interface{})`

SetNetworkArea sets NetworkArea field to given value.

### HasNetworkArea

`func (o *EventSubscription) HasNetworkArea() bool`

HasNetworkArea returns a boolean if a field has been set.

### GetNsiIdInfos

`func (o *EventSubscription) GetNsiIdInfos() []NsiIdInfo`

GetNsiIdInfos returns the NsiIdInfos field if non-nil, zero value otherwise.

### GetNsiIdInfosOk

`func (o *EventSubscription) GetNsiIdInfosOk() (*[]NsiIdInfo, bool)`

GetNsiIdInfosOk returns a tuple with the NsiIdInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiIdInfos

`func (o *EventSubscription) SetNsiIdInfos(v []NsiIdInfo)`

SetNsiIdInfos sets NsiIdInfos field to given value.

### HasNsiIdInfos

`func (o *EventSubscription) HasNsiIdInfos() bool`

HasNsiIdInfos returns a boolean if a field has been set.

### GetNsiLevelThrds

`func (o *EventSubscription) GetNsiLevelThrds() []map[string]interface{}`

GetNsiLevelThrds returns the NsiLevelThrds field if non-nil, zero value otherwise.

### GetNsiLevelThrdsOk

`func (o *EventSubscription) GetNsiLevelThrdsOk() (*[]map[string]interface{}, bool)`

GetNsiLevelThrdsOk returns a tuple with the NsiLevelThrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiLevelThrds

`func (o *EventSubscription) SetNsiLevelThrds(v []map[string]interface{})`

SetNsiLevelThrds sets NsiLevelThrds field to given value.

### HasNsiLevelThrds

`func (o *EventSubscription) HasNsiLevelThrds() bool`

HasNsiLevelThrds returns a boolean if a field has been set.

### GetQosRequ

`func (o *EventSubscription) GetQosRequ() QosRequirement`

GetQosRequ returns the QosRequ field if non-nil, zero value otherwise.

### GetQosRequOk

`func (o *EventSubscription) GetQosRequOk() (*QosRequirement, bool)`

GetQosRequOk returns a tuple with the QosRequ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosRequ

`func (o *EventSubscription) SetQosRequ(v QosRequirement)`

SetQosRequ sets QosRequ field to given value.

### HasQosRequ

`func (o *EventSubscription) HasQosRequ() bool`

HasQosRequ returns a boolean if a field has been set.

### GetQosFlowRetThds

`func (o *EventSubscription) GetQosFlowRetThds() []RetainabilityThreshold`

GetQosFlowRetThds returns the QosFlowRetThds field if non-nil, zero value otherwise.

### GetQosFlowRetThdsOk

`func (o *EventSubscription) GetQosFlowRetThdsOk() (*[]RetainabilityThreshold, bool)`

GetQosFlowRetThdsOk returns a tuple with the QosFlowRetThds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowRetThds

`func (o *EventSubscription) SetQosFlowRetThds(v []RetainabilityThreshold)`

SetQosFlowRetThds sets QosFlowRetThds field to given value.

### HasQosFlowRetThds

`func (o *EventSubscription) HasQosFlowRetThds() bool`

HasQosFlowRetThds returns a boolean if a field has been set.

### GetRanUeThrouThds

`func (o *EventSubscription) GetRanUeThrouThds() []map[string]interface{}`

GetRanUeThrouThds returns the RanUeThrouThds field if non-nil, zero value otherwise.

### GetRanUeThrouThdsOk

`func (o *EventSubscription) GetRanUeThrouThdsOk() (*[]map[string]interface{}, bool)`

GetRanUeThrouThdsOk returns a tuple with the RanUeThrouThds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUeThrouThds

`func (o *EventSubscription) SetRanUeThrouThds(v []map[string]interface{})`

SetRanUeThrouThds sets RanUeThrouThds field to given value.

### HasRanUeThrouThds

`func (o *EventSubscription) HasRanUeThrouThds() bool`

HasRanUeThrouThds returns a boolean if a field has been set.

### GetRepetitionPeriod

`func (o *EventSubscription) GetRepetitionPeriod() map[string]interface{}`

GetRepetitionPeriod returns the RepetitionPeriod field if non-nil, zero value otherwise.

### GetRepetitionPeriodOk

`func (o *EventSubscription) GetRepetitionPeriodOk() (*map[string]interface{}, bool)`

GetRepetitionPeriodOk returns a tuple with the RepetitionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionPeriod

`func (o *EventSubscription) SetRepetitionPeriod(v map[string]interface{})`

SetRepetitionPeriod sets RepetitionPeriod field to given value.

### HasRepetitionPeriod

`func (o *EventSubscription) HasRepetitionPeriod() bool`

HasRepetitionPeriod returns a boolean if a field has been set.

### GetSnssaia

`func (o *EventSubscription) GetSnssaia() []map[string]interface{}`

GetSnssaia returns the Snssaia field if non-nil, zero value otherwise.

### GetSnssaiaOk

`func (o *EventSubscription) GetSnssaiaOk() (*[]map[string]interface{}, bool)`

GetSnssaiaOk returns a tuple with the Snssaia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaia

`func (o *EventSubscription) SetSnssaia(v []map[string]interface{})`

SetSnssaia sets Snssaia field to given value.

### HasSnssaia

`func (o *EventSubscription) HasSnssaia() bool`

HasSnssaia returns a boolean if a field has been set.

### GetTgtUe

`func (o *EventSubscription) GetTgtUe() TargetUeInformation`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *EventSubscription) GetTgtUeOk() (*TargetUeInformation, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *EventSubscription) SetTgtUe(v TargetUeInformation)`

SetTgtUe sets TgtUe field to given value.

### HasTgtUe

`func (o *EventSubscription) HasTgtUe() bool`

HasTgtUe returns a boolean if a field has been set.

### GetCongThresholds

`func (o *EventSubscription) GetCongThresholds() []ThresholdLevel`

GetCongThresholds returns the CongThresholds field if non-nil, zero value otherwise.

### GetCongThresholdsOk

`func (o *EventSubscription) GetCongThresholdsOk() (*[]ThresholdLevel, bool)`

GetCongThresholdsOk returns a tuple with the CongThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongThresholds

`func (o *EventSubscription) SetCongThresholds(v []ThresholdLevel)`

SetCongThresholds sets CongThresholds field to given value.

### HasCongThresholds

`func (o *EventSubscription) HasCongThresholds() bool`

HasCongThresholds returns a boolean if a field has been set.

### GetNwPerfRequs

`func (o *EventSubscription) GetNwPerfRequs() []NetworkPerfRequirement`

GetNwPerfRequs returns the NwPerfRequs field if non-nil, zero value otherwise.

### GetNwPerfRequsOk

`func (o *EventSubscription) GetNwPerfRequsOk() (*[]NetworkPerfRequirement, bool)`

GetNwPerfRequsOk returns a tuple with the NwPerfRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfRequs

`func (o *EventSubscription) SetNwPerfRequs(v []NetworkPerfRequirement)`

SetNwPerfRequs sets NwPerfRequs field to given value.

### HasNwPerfRequs

`func (o *EventSubscription) HasNwPerfRequs() bool`

HasNwPerfRequs returns a boolean if a field has been set.

### GetBwRequs

`func (o *EventSubscription) GetBwRequs() []BwRequirement`

GetBwRequs returns the BwRequs field if non-nil, zero value otherwise.

### GetBwRequsOk

`func (o *EventSubscription) GetBwRequsOk() (*[]BwRequirement, bool)`

GetBwRequsOk returns a tuple with the BwRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwRequs

`func (o *EventSubscription) SetBwRequs(v []BwRequirement)`

SetBwRequs sets BwRequs field to given value.

### HasBwRequs

`func (o *EventSubscription) HasBwRequs() bool`

HasBwRequs returns a boolean if a field has been set.

### GetExcepRequs

`func (o *EventSubscription) GetExcepRequs() []Exception`

GetExcepRequs returns the ExcepRequs field if non-nil, zero value otherwise.

### GetExcepRequsOk

`func (o *EventSubscription) GetExcepRequsOk() (*[]Exception, bool)`

GetExcepRequsOk returns a tuple with the ExcepRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepRequs

`func (o *EventSubscription) SetExcepRequs(v []Exception)`

SetExcepRequs sets ExcepRequs field to given value.

### HasExcepRequs

`func (o *EventSubscription) HasExcepRequs() bool`

HasExcepRequs returns a boolean if a field has been set.

### GetExptAnaType

`func (o *EventSubscription) GetExptAnaType() ExpectedAnalyticsType`

GetExptAnaType returns the ExptAnaType field if non-nil, zero value otherwise.

### GetExptAnaTypeOk

`func (o *EventSubscription) GetExptAnaTypeOk() (*ExpectedAnalyticsType, bool)`

GetExptAnaTypeOk returns a tuple with the ExptAnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptAnaType

`func (o *EventSubscription) SetExptAnaType(v ExpectedAnalyticsType)`

SetExptAnaType sets ExptAnaType field to given value.

### HasExptAnaType

`func (o *EventSubscription) HasExptAnaType() bool`

HasExptAnaType returns a boolean if a field has been set.

### GetExptUeBehav

`func (o *EventSubscription) GetExptUeBehav() map[string]interface{}`

GetExptUeBehav returns the ExptUeBehav field if non-nil, zero value otherwise.

### GetExptUeBehavOk

`func (o *EventSubscription) GetExptUeBehavOk() (*map[string]interface{}, bool)`

GetExptUeBehavOk returns a tuple with the ExptUeBehav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptUeBehav

`func (o *EventSubscription) SetExptUeBehav(v map[string]interface{})`

SetExptUeBehav sets ExptUeBehav field to given value.

### HasExptUeBehav

`func (o *EventSubscription) HasExptUeBehav() bool`

HasExptUeBehav returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


