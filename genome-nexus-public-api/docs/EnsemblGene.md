# EnsemblGene

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneId** | **string** | Ensembl gene id | 
**HugoSymbol** | **string** | Approved Hugo symbol | 
**Synonyms** | Pointer to **[]string** | Hugo symbol synonyms | [optional] 
**PreviousSymbols** | Pointer to **[]string** | Previous Hugo symbols | [optional] 
**EntrezGeneId** | Pointer to **string** | Entrez Gene Id | [optional] 

## Methods

### NewEnsemblGene

`func NewEnsemblGene(geneId string, hugoSymbol string, ) *EnsemblGene`

NewEnsemblGene instantiates a new EnsemblGene object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnsemblGeneWithDefaults

`func NewEnsemblGeneWithDefaults() *EnsemblGene`

NewEnsemblGeneWithDefaults instantiates a new EnsemblGene object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneId

`func (o *EnsemblGene) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *EnsemblGene) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *EnsemblGene) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.


### GetHugoSymbol

`func (o *EnsemblGene) GetHugoSymbol() string`

GetHugoSymbol returns the HugoSymbol field if non-nil, zero value otherwise.

### GetHugoSymbolOk

`func (o *EnsemblGene) GetHugoSymbolOk() (*string, bool)`

GetHugoSymbolOk returns a tuple with the HugoSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugoSymbol

`func (o *EnsemblGene) SetHugoSymbol(v string)`

SetHugoSymbol sets HugoSymbol field to given value.


### GetSynonyms

`func (o *EnsemblGene) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *EnsemblGene) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *EnsemblGene) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *EnsemblGene) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetPreviousSymbols

`func (o *EnsemblGene) GetPreviousSymbols() []string`

GetPreviousSymbols returns the PreviousSymbols field if non-nil, zero value otherwise.

### GetPreviousSymbolsOk

`func (o *EnsemblGene) GetPreviousSymbolsOk() (*[]string, bool)`

GetPreviousSymbolsOk returns a tuple with the PreviousSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSymbols

`func (o *EnsemblGene) SetPreviousSymbols(v []string)`

SetPreviousSymbols sets PreviousSymbols field to given value.

### HasPreviousSymbols

`func (o *EnsemblGene) HasPreviousSymbols() bool`

HasPreviousSymbols returns a boolean if a field has been set.

### GetEntrezGeneId

`func (o *EnsemblGene) GetEntrezGeneId() string`

GetEntrezGeneId returns the EntrezGeneId field if non-nil, zero value otherwise.

### GetEntrezGeneIdOk

`func (o *EnsemblGene) GetEntrezGeneIdOk() (*string, bool)`

GetEntrezGeneIdOk returns a tuple with the EntrezGeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrezGeneId

`func (o *EnsemblGene) SetEntrezGeneId(v string)`

SetEntrezGeneId sets EntrezGeneId field to given value.

### HasEntrezGeneId

`func (o *EnsemblGene) HasEntrezGeneId() bool`

HasEntrezGeneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


