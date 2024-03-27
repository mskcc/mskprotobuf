# IndicatorQueryResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlleleExist** | Pointer to **bool** |  | [optional] 
**DataVersion** | Pointer to **string** |  | [optional] 
**DiagnosticImplications** | Pointer to [**[]Implication**](Implication.md) |  | [optional] 
**DiagnosticSummary** | Pointer to **string** |  | [optional] 
**GeneExist** | Pointer to **bool** |  | [optional] 
**GeneSummary** | Pointer to **string** |  | [optional] 
**HighestDiagnosticImplicationLevel** | Pointer to **string** |  | [optional] 
**HighestPrognosticImplicationLevel** | Pointer to **string** |  | [optional] 
**HighestResistanceLevel** | Pointer to **string** |  | [optional] 
**HighestSensitiveLevel** | Pointer to **string** |  | [optional] 
**Hotspot** | Pointer to **bool** |  | [optional] 
**LastUpdate** | Pointer to **string** |  | [optional] 
**MutationEffect** | Pointer to [**MutationEffectResp**](MutationEffectResp.md) |  | [optional] 
**Oncogenic** | Pointer to **string** |  | [optional] 
**OtherSignificantResistanceLevels** | Pointer to **[]string** |  | [optional] 
**OtherSignificantSensitiveLevels** | Pointer to **[]string** |  | [optional] 
**PrognosticImplications** | Pointer to [**[]Implication**](Implication.md) |  | [optional] 
**PrognosticSummary** | Pointer to **string** |  | [optional] 
**Query** | Pointer to [**Query**](Query.md) |  | [optional] 
**Treatments** | Pointer to [**[]IndicatorQueryTreatment**](IndicatorQueryTreatment.md) |  | [optional] 
**TumorTypeSummary** | Pointer to **string** |  | [optional] 
**VariantExist** | Pointer to **bool** |  | [optional] 
**VariantSummary** | Pointer to **string** |  | [optional] 
**Vus** | Pointer to **bool** |  | [optional] 

## Methods

### NewIndicatorQueryResp

`func NewIndicatorQueryResp() *IndicatorQueryResp`

NewIndicatorQueryResp instantiates a new IndicatorQueryResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorQueryRespWithDefaults

`func NewIndicatorQueryRespWithDefaults() *IndicatorQueryResp`

NewIndicatorQueryRespWithDefaults instantiates a new IndicatorQueryResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlleleExist

`func (o *IndicatorQueryResp) GetAlleleExist() bool`

GetAlleleExist returns the AlleleExist field if non-nil, zero value otherwise.

### GetAlleleExistOk

`func (o *IndicatorQueryResp) GetAlleleExistOk() (*bool, bool)`

GetAlleleExistOk returns a tuple with the AlleleExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlleleExist

`func (o *IndicatorQueryResp) SetAlleleExist(v bool)`

SetAlleleExist sets AlleleExist field to given value.

### HasAlleleExist

`func (o *IndicatorQueryResp) HasAlleleExist() bool`

HasAlleleExist returns a boolean if a field has been set.

### GetDataVersion

`func (o *IndicatorQueryResp) GetDataVersion() string`

GetDataVersion returns the DataVersion field if non-nil, zero value otherwise.

### GetDataVersionOk

`func (o *IndicatorQueryResp) GetDataVersionOk() (*string, bool)`

GetDataVersionOk returns a tuple with the DataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVersion

`func (o *IndicatorQueryResp) SetDataVersion(v string)`

SetDataVersion sets DataVersion field to given value.

### HasDataVersion

`func (o *IndicatorQueryResp) HasDataVersion() bool`

HasDataVersion returns a boolean if a field has been set.

### GetDiagnosticImplications

`func (o *IndicatorQueryResp) GetDiagnosticImplications() []Implication`

GetDiagnosticImplications returns the DiagnosticImplications field if non-nil, zero value otherwise.

### GetDiagnosticImplicationsOk

`func (o *IndicatorQueryResp) GetDiagnosticImplicationsOk() (*[]Implication, bool)`

GetDiagnosticImplicationsOk returns a tuple with the DiagnosticImplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticImplications

`func (o *IndicatorQueryResp) SetDiagnosticImplications(v []Implication)`

SetDiagnosticImplications sets DiagnosticImplications field to given value.

### HasDiagnosticImplications

`func (o *IndicatorQueryResp) HasDiagnosticImplications() bool`

HasDiagnosticImplications returns a boolean if a field has been set.

### GetDiagnosticSummary

`func (o *IndicatorQueryResp) GetDiagnosticSummary() string`

GetDiagnosticSummary returns the DiagnosticSummary field if non-nil, zero value otherwise.

### GetDiagnosticSummaryOk

`func (o *IndicatorQueryResp) GetDiagnosticSummaryOk() (*string, bool)`

GetDiagnosticSummaryOk returns a tuple with the DiagnosticSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticSummary

`func (o *IndicatorQueryResp) SetDiagnosticSummary(v string)`

SetDiagnosticSummary sets DiagnosticSummary field to given value.

### HasDiagnosticSummary

`func (o *IndicatorQueryResp) HasDiagnosticSummary() bool`

HasDiagnosticSummary returns a boolean if a field has been set.

### GetGeneExist

`func (o *IndicatorQueryResp) GetGeneExist() bool`

GetGeneExist returns the GeneExist field if non-nil, zero value otherwise.

### GetGeneExistOk

`func (o *IndicatorQueryResp) GetGeneExistOk() (*bool, bool)`

GetGeneExistOk returns a tuple with the GeneExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneExist

`func (o *IndicatorQueryResp) SetGeneExist(v bool)`

SetGeneExist sets GeneExist field to given value.

### HasGeneExist

`func (o *IndicatorQueryResp) HasGeneExist() bool`

HasGeneExist returns a boolean if a field has been set.

### GetGeneSummary

`func (o *IndicatorQueryResp) GetGeneSummary() string`

GetGeneSummary returns the GeneSummary field if non-nil, zero value otherwise.

### GetGeneSummaryOk

`func (o *IndicatorQueryResp) GetGeneSummaryOk() (*string, bool)`

GetGeneSummaryOk returns a tuple with the GeneSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSummary

`func (o *IndicatorQueryResp) SetGeneSummary(v string)`

SetGeneSummary sets GeneSummary field to given value.

### HasGeneSummary

`func (o *IndicatorQueryResp) HasGeneSummary() bool`

HasGeneSummary returns a boolean if a field has been set.

### GetHighestDiagnosticImplicationLevel

`func (o *IndicatorQueryResp) GetHighestDiagnosticImplicationLevel() string`

GetHighestDiagnosticImplicationLevel returns the HighestDiagnosticImplicationLevel field if non-nil, zero value otherwise.

### GetHighestDiagnosticImplicationLevelOk

`func (o *IndicatorQueryResp) GetHighestDiagnosticImplicationLevelOk() (*string, bool)`

GetHighestDiagnosticImplicationLevelOk returns a tuple with the HighestDiagnosticImplicationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestDiagnosticImplicationLevel

`func (o *IndicatorQueryResp) SetHighestDiagnosticImplicationLevel(v string)`

SetHighestDiagnosticImplicationLevel sets HighestDiagnosticImplicationLevel field to given value.

### HasHighestDiagnosticImplicationLevel

`func (o *IndicatorQueryResp) HasHighestDiagnosticImplicationLevel() bool`

HasHighestDiagnosticImplicationLevel returns a boolean if a field has been set.

### GetHighestPrognosticImplicationLevel

`func (o *IndicatorQueryResp) GetHighestPrognosticImplicationLevel() string`

GetHighestPrognosticImplicationLevel returns the HighestPrognosticImplicationLevel field if non-nil, zero value otherwise.

### GetHighestPrognosticImplicationLevelOk

`func (o *IndicatorQueryResp) GetHighestPrognosticImplicationLevelOk() (*string, bool)`

GetHighestPrognosticImplicationLevelOk returns a tuple with the HighestPrognosticImplicationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestPrognosticImplicationLevel

`func (o *IndicatorQueryResp) SetHighestPrognosticImplicationLevel(v string)`

SetHighestPrognosticImplicationLevel sets HighestPrognosticImplicationLevel field to given value.

### HasHighestPrognosticImplicationLevel

`func (o *IndicatorQueryResp) HasHighestPrognosticImplicationLevel() bool`

HasHighestPrognosticImplicationLevel returns a boolean if a field has been set.

### GetHighestResistanceLevel

`func (o *IndicatorQueryResp) GetHighestResistanceLevel() string`

GetHighestResistanceLevel returns the HighestResistanceLevel field if non-nil, zero value otherwise.

### GetHighestResistanceLevelOk

`func (o *IndicatorQueryResp) GetHighestResistanceLevelOk() (*string, bool)`

GetHighestResistanceLevelOk returns a tuple with the HighestResistanceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestResistanceLevel

`func (o *IndicatorQueryResp) SetHighestResistanceLevel(v string)`

SetHighestResistanceLevel sets HighestResistanceLevel field to given value.

### HasHighestResistanceLevel

`func (o *IndicatorQueryResp) HasHighestResistanceLevel() bool`

HasHighestResistanceLevel returns a boolean if a field has been set.

### GetHighestSensitiveLevel

`func (o *IndicatorQueryResp) GetHighestSensitiveLevel() string`

GetHighestSensitiveLevel returns the HighestSensitiveLevel field if non-nil, zero value otherwise.

### GetHighestSensitiveLevelOk

`func (o *IndicatorQueryResp) GetHighestSensitiveLevelOk() (*string, bool)`

GetHighestSensitiveLevelOk returns a tuple with the HighestSensitiveLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestSensitiveLevel

`func (o *IndicatorQueryResp) SetHighestSensitiveLevel(v string)`

SetHighestSensitiveLevel sets HighestSensitiveLevel field to given value.

### HasHighestSensitiveLevel

`func (o *IndicatorQueryResp) HasHighestSensitiveLevel() bool`

HasHighestSensitiveLevel returns a boolean if a field has been set.

### GetHotspot

`func (o *IndicatorQueryResp) GetHotspot() bool`

GetHotspot returns the Hotspot field if non-nil, zero value otherwise.

### GetHotspotOk

`func (o *IndicatorQueryResp) GetHotspotOk() (*bool, bool)`

GetHotspotOk returns a tuple with the Hotspot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspot

`func (o *IndicatorQueryResp) SetHotspot(v bool)`

SetHotspot sets Hotspot field to given value.

### HasHotspot

`func (o *IndicatorQueryResp) HasHotspot() bool`

HasHotspot returns a boolean if a field has been set.

### GetLastUpdate

`func (o *IndicatorQueryResp) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *IndicatorQueryResp) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *IndicatorQueryResp) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *IndicatorQueryResp) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetMutationEffect

`func (o *IndicatorQueryResp) GetMutationEffect() MutationEffectResp`

GetMutationEffect returns the MutationEffect field if non-nil, zero value otherwise.

### GetMutationEffectOk

`func (o *IndicatorQueryResp) GetMutationEffectOk() (*MutationEffectResp, bool)`

GetMutationEffectOk returns a tuple with the MutationEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutationEffect

`func (o *IndicatorQueryResp) SetMutationEffect(v MutationEffectResp)`

SetMutationEffect sets MutationEffect field to given value.

### HasMutationEffect

`func (o *IndicatorQueryResp) HasMutationEffect() bool`

HasMutationEffect returns a boolean if a field has been set.

### GetOncogenic

`func (o *IndicatorQueryResp) GetOncogenic() string`

GetOncogenic returns the Oncogenic field if non-nil, zero value otherwise.

### GetOncogenicOk

`func (o *IndicatorQueryResp) GetOncogenicOk() (*string, bool)`

GetOncogenicOk returns a tuple with the Oncogenic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOncogenic

`func (o *IndicatorQueryResp) SetOncogenic(v string)`

SetOncogenic sets Oncogenic field to given value.

### HasOncogenic

`func (o *IndicatorQueryResp) HasOncogenic() bool`

HasOncogenic returns a boolean if a field has been set.

### GetOtherSignificantResistanceLevels

`func (o *IndicatorQueryResp) GetOtherSignificantResistanceLevels() []string`

GetOtherSignificantResistanceLevels returns the OtherSignificantResistanceLevels field if non-nil, zero value otherwise.

### GetOtherSignificantResistanceLevelsOk

`func (o *IndicatorQueryResp) GetOtherSignificantResistanceLevelsOk() (*[]string, bool)`

GetOtherSignificantResistanceLevelsOk returns a tuple with the OtherSignificantResistanceLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSignificantResistanceLevels

`func (o *IndicatorQueryResp) SetOtherSignificantResistanceLevels(v []string)`

SetOtherSignificantResistanceLevels sets OtherSignificantResistanceLevels field to given value.

### HasOtherSignificantResistanceLevels

`func (o *IndicatorQueryResp) HasOtherSignificantResistanceLevels() bool`

HasOtherSignificantResistanceLevels returns a boolean if a field has been set.

### GetOtherSignificantSensitiveLevels

`func (o *IndicatorQueryResp) GetOtherSignificantSensitiveLevels() []string`

GetOtherSignificantSensitiveLevels returns the OtherSignificantSensitiveLevels field if non-nil, zero value otherwise.

### GetOtherSignificantSensitiveLevelsOk

`func (o *IndicatorQueryResp) GetOtherSignificantSensitiveLevelsOk() (*[]string, bool)`

GetOtherSignificantSensitiveLevelsOk returns a tuple with the OtherSignificantSensitiveLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSignificantSensitiveLevels

`func (o *IndicatorQueryResp) SetOtherSignificantSensitiveLevels(v []string)`

SetOtherSignificantSensitiveLevels sets OtherSignificantSensitiveLevels field to given value.

### HasOtherSignificantSensitiveLevels

`func (o *IndicatorQueryResp) HasOtherSignificantSensitiveLevels() bool`

HasOtherSignificantSensitiveLevels returns a boolean if a field has been set.

### GetPrognosticImplications

`func (o *IndicatorQueryResp) GetPrognosticImplications() []Implication`

GetPrognosticImplications returns the PrognosticImplications field if non-nil, zero value otherwise.

### GetPrognosticImplicationsOk

`func (o *IndicatorQueryResp) GetPrognosticImplicationsOk() (*[]Implication, bool)`

GetPrognosticImplicationsOk returns a tuple with the PrognosticImplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrognosticImplications

`func (o *IndicatorQueryResp) SetPrognosticImplications(v []Implication)`

SetPrognosticImplications sets PrognosticImplications field to given value.

### HasPrognosticImplications

`func (o *IndicatorQueryResp) HasPrognosticImplications() bool`

HasPrognosticImplications returns a boolean if a field has been set.

### GetPrognosticSummary

`func (o *IndicatorQueryResp) GetPrognosticSummary() string`

GetPrognosticSummary returns the PrognosticSummary field if non-nil, zero value otherwise.

### GetPrognosticSummaryOk

`func (o *IndicatorQueryResp) GetPrognosticSummaryOk() (*string, bool)`

GetPrognosticSummaryOk returns a tuple with the PrognosticSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrognosticSummary

`func (o *IndicatorQueryResp) SetPrognosticSummary(v string)`

SetPrognosticSummary sets PrognosticSummary field to given value.

### HasPrognosticSummary

`func (o *IndicatorQueryResp) HasPrognosticSummary() bool`

HasPrognosticSummary returns a boolean if a field has been set.

### GetQuery

`func (o *IndicatorQueryResp) GetQuery() Query`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IndicatorQueryResp) GetQueryOk() (*Query, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IndicatorQueryResp) SetQuery(v Query)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IndicatorQueryResp) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTreatments

`func (o *IndicatorQueryResp) GetTreatments() []IndicatorQueryTreatment`

GetTreatments returns the Treatments field if non-nil, zero value otherwise.

### GetTreatmentsOk

`func (o *IndicatorQueryResp) GetTreatmentsOk() (*[]IndicatorQueryTreatment, bool)`

GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatments

`func (o *IndicatorQueryResp) SetTreatments(v []IndicatorQueryTreatment)`

SetTreatments sets Treatments field to given value.

### HasTreatments

`func (o *IndicatorQueryResp) HasTreatments() bool`

HasTreatments returns a boolean if a field has been set.

### GetTumorTypeSummary

`func (o *IndicatorQueryResp) GetTumorTypeSummary() string`

GetTumorTypeSummary returns the TumorTypeSummary field if non-nil, zero value otherwise.

### GetTumorTypeSummaryOk

`func (o *IndicatorQueryResp) GetTumorTypeSummaryOk() (*string, bool)`

GetTumorTypeSummaryOk returns a tuple with the TumorTypeSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTumorTypeSummary

`func (o *IndicatorQueryResp) SetTumorTypeSummary(v string)`

SetTumorTypeSummary sets TumorTypeSummary field to given value.

### HasTumorTypeSummary

`func (o *IndicatorQueryResp) HasTumorTypeSummary() bool`

HasTumorTypeSummary returns a boolean if a field has been set.

### GetVariantExist

`func (o *IndicatorQueryResp) GetVariantExist() bool`

GetVariantExist returns the VariantExist field if non-nil, zero value otherwise.

### GetVariantExistOk

`func (o *IndicatorQueryResp) GetVariantExistOk() (*bool, bool)`

GetVariantExistOk returns a tuple with the VariantExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantExist

`func (o *IndicatorQueryResp) SetVariantExist(v bool)`

SetVariantExist sets VariantExist field to given value.

### HasVariantExist

`func (o *IndicatorQueryResp) HasVariantExist() bool`

HasVariantExist returns a boolean if a field has been set.

### GetVariantSummary

`func (o *IndicatorQueryResp) GetVariantSummary() string`

GetVariantSummary returns the VariantSummary field if non-nil, zero value otherwise.

### GetVariantSummaryOk

`func (o *IndicatorQueryResp) GetVariantSummaryOk() (*string, bool)`

GetVariantSummaryOk returns a tuple with the VariantSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantSummary

`func (o *IndicatorQueryResp) SetVariantSummary(v string)`

SetVariantSummary sets VariantSummary field to given value.

### HasVariantSummary

`func (o *IndicatorQueryResp) HasVariantSummary() bool`

HasVariantSummary returns a boolean if a field has been set.

### GetVus

`func (o *IndicatorQueryResp) GetVus() bool`

GetVus returns the Vus field if non-nil, zero value otherwise.

### GetVusOk

`func (o *IndicatorQueryResp) GetVusOk() (*bool, bool)`

GetVusOk returns a tuple with the Vus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVus

`func (o *IndicatorQueryResp) SetVus(v bool)`

SetVus sets Vus field to given value.

### HasVus

`func (o *IndicatorQueryResp) HasVus() bool`

HasVus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


