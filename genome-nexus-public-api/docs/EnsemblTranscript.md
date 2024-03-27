# EnsemblTranscript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniprotId** | Pointer to **string** |  | [optional] 
**TranscriptId** | **string** | Ensembl transcript id | 
**GeneId** | **string** | Ensembl gene id | 
**ProteinId** | **string** | Ensembl protein id | 
**ProteinLength** | Pointer to **int32** | Length of protein | [optional] 
**PfamDomains** | Pointer to [**[]PfamDomainRange**](PfamDomainRange.md) | Pfam domains | [optional] 
**HugoSymbols** | Pointer to **[]string** | Hugo symbols | [optional] 
**RefseqMrnaId** | Pointer to **string** | RefSeq mRNA ID | [optional] 
**CcdsId** | Pointer to **string** | Consensus CDS (CCDS) ID | [optional] 
**Exons** | Pointer to [**[]Exon**](Exon.md) | Exon information | [optional] 
**Utrs** | Pointer to [**[]UntranslatedRegion**](UntranslatedRegion.md) | UTR information | [optional] 

## Methods

### NewEnsemblTranscript

`func NewEnsemblTranscript(transcriptId string, geneId string, proteinId string, ) *EnsemblTranscript`

NewEnsemblTranscript instantiates a new EnsemblTranscript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnsemblTranscriptWithDefaults

`func NewEnsemblTranscriptWithDefaults() *EnsemblTranscript`

NewEnsemblTranscriptWithDefaults instantiates a new EnsemblTranscript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniprotId

`func (o *EnsemblTranscript) GetUniprotId() string`

GetUniprotId returns the UniprotId field if non-nil, zero value otherwise.

### GetUniprotIdOk

`func (o *EnsemblTranscript) GetUniprotIdOk() (*string, bool)`

GetUniprotIdOk returns a tuple with the UniprotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotId

`func (o *EnsemblTranscript) SetUniprotId(v string)`

SetUniprotId sets UniprotId field to given value.

### HasUniprotId

`func (o *EnsemblTranscript) HasUniprotId() bool`

HasUniprotId returns a boolean if a field has been set.

### GetTranscriptId

`func (o *EnsemblTranscript) GetTranscriptId() string`

GetTranscriptId returns the TranscriptId field if non-nil, zero value otherwise.

### GetTranscriptIdOk

`func (o *EnsemblTranscript) GetTranscriptIdOk() (*string, bool)`

GetTranscriptIdOk returns a tuple with the TranscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptId

`func (o *EnsemblTranscript) SetTranscriptId(v string)`

SetTranscriptId sets TranscriptId field to given value.


### GetGeneId

`func (o *EnsemblTranscript) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *EnsemblTranscript) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *EnsemblTranscript) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.


### GetProteinId

`func (o *EnsemblTranscript) GetProteinId() string`

GetProteinId returns the ProteinId field if non-nil, zero value otherwise.

### GetProteinIdOk

`func (o *EnsemblTranscript) GetProteinIdOk() (*string, bool)`

GetProteinIdOk returns a tuple with the ProteinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinId

`func (o *EnsemblTranscript) SetProteinId(v string)`

SetProteinId sets ProteinId field to given value.


### GetProteinLength

`func (o *EnsemblTranscript) GetProteinLength() int32`

GetProteinLength returns the ProteinLength field if non-nil, zero value otherwise.

### GetProteinLengthOk

`func (o *EnsemblTranscript) GetProteinLengthOk() (*int32, bool)`

GetProteinLengthOk returns a tuple with the ProteinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinLength

`func (o *EnsemblTranscript) SetProteinLength(v int32)`

SetProteinLength sets ProteinLength field to given value.

### HasProteinLength

`func (o *EnsemblTranscript) HasProteinLength() bool`

HasProteinLength returns a boolean if a field has been set.

### GetPfamDomains

`func (o *EnsemblTranscript) GetPfamDomains() []PfamDomainRange`

GetPfamDomains returns the PfamDomains field if non-nil, zero value otherwise.

### GetPfamDomainsOk

`func (o *EnsemblTranscript) GetPfamDomainsOk() (*[]PfamDomainRange, bool)`

GetPfamDomainsOk returns a tuple with the PfamDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfamDomains

`func (o *EnsemblTranscript) SetPfamDomains(v []PfamDomainRange)`

SetPfamDomains sets PfamDomains field to given value.

### HasPfamDomains

`func (o *EnsemblTranscript) HasPfamDomains() bool`

HasPfamDomains returns a boolean if a field has been set.

### GetHugoSymbols

`func (o *EnsemblTranscript) GetHugoSymbols() []string`

GetHugoSymbols returns the HugoSymbols field if non-nil, zero value otherwise.

### GetHugoSymbolsOk

`func (o *EnsemblTranscript) GetHugoSymbolsOk() (*[]string, bool)`

GetHugoSymbolsOk returns a tuple with the HugoSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugoSymbols

`func (o *EnsemblTranscript) SetHugoSymbols(v []string)`

SetHugoSymbols sets HugoSymbols field to given value.

### HasHugoSymbols

`func (o *EnsemblTranscript) HasHugoSymbols() bool`

HasHugoSymbols returns a boolean if a field has been set.

### GetRefseqMrnaId

`func (o *EnsemblTranscript) GetRefseqMrnaId() string`

GetRefseqMrnaId returns the RefseqMrnaId field if non-nil, zero value otherwise.

### GetRefseqMrnaIdOk

`func (o *EnsemblTranscript) GetRefseqMrnaIdOk() (*string, bool)`

GetRefseqMrnaIdOk returns a tuple with the RefseqMrnaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqMrnaId

`func (o *EnsemblTranscript) SetRefseqMrnaId(v string)`

SetRefseqMrnaId sets RefseqMrnaId field to given value.

### HasRefseqMrnaId

`func (o *EnsemblTranscript) HasRefseqMrnaId() bool`

HasRefseqMrnaId returns a boolean if a field has been set.

### GetCcdsId

`func (o *EnsemblTranscript) GetCcdsId() string`

GetCcdsId returns the CcdsId field if non-nil, zero value otherwise.

### GetCcdsIdOk

`func (o *EnsemblTranscript) GetCcdsIdOk() (*string, bool)`

GetCcdsIdOk returns a tuple with the CcdsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcdsId

`func (o *EnsemblTranscript) SetCcdsId(v string)`

SetCcdsId sets CcdsId field to given value.

### HasCcdsId

`func (o *EnsemblTranscript) HasCcdsId() bool`

HasCcdsId returns a boolean if a field has been set.

### GetExons

`func (o *EnsemblTranscript) GetExons() []Exon`

GetExons returns the Exons field if non-nil, zero value otherwise.

### GetExonsOk

`func (o *EnsemblTranscript) GetExonsOk() (*[]Exon, bool)`

GetExonsOk returns a tuple with the Exons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExons

`func (o *EnsemblTranscript) SetExons(v []Exon)`

SetExons sets Exons field to given value.

### HasExons

`func (o *EnsemblTranscript) HasExons() bool`

HasExons returns a boolean if a field has been set.

### GetUtrs

`func (o *EnsemblTranscript) GetUtrs() []UntranslatedRegion`

GetUtrs returns the Utrs field if non-nil, zero value otherwise.

### GetUtrsOk

`func (o *EnsemblTranscript) GetUtrsOk() (*[]UntranslatedRegion, bool)`

GetUtrsOk returns a tuple with the Utrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtrs

`func (o *EnsemblTranscript) SetUtrs(v []UntranslatedRegion)`

SetUtrs sets Utrs field to given value.

### HasUtrs

`func (o *EnsemblTranscript) HasUtrs() bool`

HasUtrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


