# EnsemblFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneIds** | Pointer to **[]string** | List of Ensembl gene IDs. For example [\&quot;ENSG00000136999\&quot;, \&quot;ENSG00000272398\&quot;, \&quot;ENSG00000198695\&quot;] | [optional] 
**HugoSymbols** | Pointer to **[]string** | List of Hugo Symbols. For example [\&quot;TP53\&quot;, \&quot;PIK3CA\&quot;, \&quot;BRCA1\&quot;] | [optional] 
**ProteinIds** | Pointer to **[]string** | List of Ensembl protein IDs. For example [\&quot;ENSP00000439985\&quot;, \&quot;ENSP00000478460\&quot;, \&quot;ENSP00000346196\&quot;] | [optional] 
**TranscriptIds** | Pointer to **[]string** | List of Ensembl transcript IDs. For example [\&quot;ENST00000361390\&quot;, \&quot;ENST00000361453\&quot;, \&quot;ENST00000361624\&quot;] | [optional] 

## Methods

### NewEnsemblFilter

`func NewEnsemblFilter() *EnsemblFilter`

NewEnsemblFilter instantiates a new EnsemblFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnsemblFilterWithDefaults

`func NewEnsemblFilterWithDefaults() *EnsemblFilter`

NewEnsemblFilterWithDefaults instantiates a new EnsemblFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneIds

`func (o *EnsemblFilter) GetGeneIds() []string`

GetGeneIds returns the GeneIds field if non-nil, zero value otherwise.

### GetGeneIdsOk

`func (o *EnsemblFilter) GetGeneIdsOk() (*[]string, bool)`

GetGeneIdsOk returns a tuple with the GeneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneIds

`func (o *EnsemblFilter) SetGeneIds(v []string)`

SetGeneIds sets GeneIds field to given value.

### HasGeneIds

`func (o *EnsemblFilter) HasGeneIds() bool`

HasGeneIds returns a boolean if a field has been set.

### GetHugoSymbols

`func (o *EnsemblFilter) GetHugoSymbols() []string`

GetHugoSymbols returns the HugoSymbols field if non-nil, zero value otherwise.

### GetHugoSymbolsOk

`func (o *EnsemblFilter) GetHugoSymbolsOk() (*[]string, bool)`

GetHugoSymbolsOk returns a tuple with the HugoSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugoSymbols

`func (o *EnsemblFilter) SetHugoSymbols(v []string)`

SetHugoSymbols sets HugoSymbols field to given value.

### HasHugoSymbols

`func (o *EnsemblFilter) HasHugoSymbols() bool`

HasHugoSymbols returns a boolean if a field has been set.

### GetProteinIds

`func (o *EnsemblFilter) GetProteinIds() []string`

GetProteinIds returns the ProteinIds field if non-nil, zero value otherwise.

### GetProteinIdsOk

`func (o *EnsemblFilter) GetProteinIdsOk() (*[]string, bool)`

GetProteinIdsOk returns a tuple with the ProteinIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinIds

`func (o *EnsemblFilter) SetProteinIds(v []string)`

SetProteinIds sets ProteinIds field to given value.

### HasProteinIds

`func (o *EnsemblFilter) HasProteinIds() bool`

HasProteinIds returns a boolean if a field has been set.

### GetTranscriptIds

`func (o *EnsemblFilter) GetTranscriptIds() []string`

GetTranscriptIds returns the TranscriptIds field if non-nil, zero value otherwise.

### GetTranscriptIdsOk

`func (o *EnsemblFilter) GetTranscriptIdsOk() (*[]string, bool)`

GetTranscriptIdsOk returns a tuple with the TranscriptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptIds

`func (o *EnsemblFilter) SetTranscriptIds(v []string)`

SetTranscriptIds sets TranscriptIds field to given value.

### HasTranscriptIds

`func (o *EnsemblFilter) HasTranscriptIds() bool`

HasTranscriptIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


