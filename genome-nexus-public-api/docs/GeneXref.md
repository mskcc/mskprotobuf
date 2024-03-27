# GeneXref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbDisplayName** | **string** | Database display name | 
**Dbname** | **string** | Database name | 
**Description** | **string** | Description | 
**DisplayId** | **string** | Display id | 
**EnsemblGeneId** | Pointer to **string** |  | [optional] 
**InfoText** | Pointer to **string** | Database info text | [optional] 
**InfoTypes** | Pointer to **string** | Database info type | [optional] 
**PrimaryId** | **string** | Primary id | 
**Synonyms** | Pointer to **[]string** | Synonyms | [optional] 
**Version** | **string** | Version | 

## Methods

### NewGeneXref

`func NewGeneXref(dbDisplayName string, dbname string, description string, displayId string, primaryId string, version string, ) *GeneXref`

NewGeneXref instantiates a new GeneXref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneXrefWithDefaults

`func NewGeneXrefWithDefaults() *GeneXref`

NewGeneXrefWithDefaults instantiates a new GeneXref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbDisplayName

`func (o *GeneXref) GetDbDisplayName() string`

GetDbDisplayName returns the DbDisplayName field if non-nil, zero value otherwise.

### GetDbDisplayNameOk

`func (o *GeneXref) GetDbDisplayNameOk() (*string, bool)`

GetDbDisplayNameOk returns a tuple with the DbDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbDisplayName

`func (o *GeneXref) SetDbDisplayName(v string)`

SetDbDisplayName sets DbDisplayName field to given value.


### GetDbname

`func (o *GeneXref) GetDbname() string`

GetDbname returns the Dbname field if non-nil, zero value otherwise.

### GetDbnameOk

`func (o *GeneXref) GetDbnameOk() (*string, bool)`

GetDbnameOk returns a tuple with the Dbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbname

`func (o *GeneXref) SetDbname(v string)`

SetDbname sets Dbname field to given value.


### GetDescription

`func (o *GeneXref) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeneXref) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeneXref) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayId

`func (o *GeneXref) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *GeneXref) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *GeneXref) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.


### GetEnsemblGeneId

`func (o *GeneXref) GetEnsemblGeneId() string`

GetEnsemblGeneId returns the EnsemblGeneId field if non-nil, zero value otherwise.

### GetEnsemblGeneIdOk

`func (o *GeneXref) GetEnsemblGeneIdOk() (*string, bool)`

GetEnsemblGeneIdOk returns a tuple with the EnsemblGeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblGeneId

`func (o *GeneXref) SetEnsemblGeneId(v string)`

SetEnsemblGeneId sets EnsemblGeneId field to given value.

### HasEnsemblGeneId

`func (o *GeneXref) HasEnsemblGeneId() bool`

HasEnsemblGeneId returns a boolean if a field has been set.

### GetInfoText

`func (o *GeneXref) GetInfoText() string`

GetInfoText returns the InfoText field if non-nil, zero value otherwise.

### GetInfoTextOk

`func (o *GeneXref) GetInfoTextOk() (*string, bool)`

GetInfoTextOk returns a tuple with the InfoText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoText

`func (o *GeneXref) SetInfoText(v string)`

SetInfoText sets InfoText field to given value.

### HasInfoText

`func (o *GeneXref) HasInfoText() bool`

HasInfoText returns a boolean if a field has been set.

### GetInfoTypes

`func (o *GeneXref) GetInfoTypes() string`

GetInfoTypes returns the InfoTypes field if non-nil, zero value otherwise.

### GetInfoTypesOk

`func (o *GeneXref) GetInfoTypesOk() (*string, bool)`

GetInfoTypesOk returns a tuple with the InfoTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoTypes

`func (o *GeneXref) SetInfoTypes(v string)`

SetInfoTypes sets InfoTypes field to given value.

### HasInfoTypes

`func (o *GeneXref) HasInfoTypes() bool`

HasInfoTypes returns a boolean if a field has been set.

### GetPrimaryId

`func (o *GeneXref) GetPrimaryId() string`

GetPrimaryId returns the PrimaryId field if non-nil, zero value otherwise.

### GetPrimaryIdOk

`func (o *GeneXref) GetPrimaryIdOk() (*string, bool)`

GetPrimaryIdOk returns a tuple with the PrimaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryId

`func (o *GeneXref) SetPrimaryId(v string)`

SetPrimaryId sets PrimaryId field to given value.


### GetSynonyms

`func (o *GeneXref) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *GeneXref) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *GeneXref) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *GeneXref) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetVersion

`func (o *GeneXref) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GeneXref) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GeneXref) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


