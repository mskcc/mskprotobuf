# TranscriptConsequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AminoAcids** | Pointer to **string** | Amino acids | [optional] 
**Canonical** | Pointer to **string** | Canonical transcript indicator | [optional] 
**Codons** | Pointer to **string** | Codons | [optional] 
**ConsequenceTerms** | Pointer to **[]string** | List of consequence terms | [optional] 
**Exon** | Pointer to **string** |  | [optional] 
**GeneId** | Pointer to **string** | Ensembl gene id | [optional] 
**GeneSymbol** | Pointer to **string** | Hugo gene symbol | [optional] 
**HgncId** | Pointer to **string** | HGNC id | [optional] 
**Hgvsc** | Pointer to **string** | HGVSc | [optional] 
**Hgvsg** | Pointer to **string** | HGVSg | [optional] 
**Hgvsp** | Pointer to **string** | HGVSp | [optional] 
**PolyphenPrediction** | Pointer to **string** | Polyphen Prediction | [optional] 
**PolyphenScore** | Pointer to **float64** | Polyphen Score | [optional] 
**ProteinEnd** | Pointer to **int32** | Protein end position | [optional] 
**ProteinId** | Pointer to **string** | Ensembl protein id | [optional] 
**ProteinStart** | Pointer to **int32** | Protein start position | [optional] 
**RefseqTranscriptIds** | Pointer to **[]string** | List of RefSeq transcript ids | [optional] 
**SiftPrediction** | Pointer to **string** | Sift Prediction | [optional] 
**SiftScore** | Pointer to **float64** | Sift Score | [optional] 
**TranscriptId** | **string** | Ensembl transcript id | 
**UniprotId** | Pointer to **string** |  | [optional] 
**VariantAllele** | Pointer to **string** | Variant allele | [optional] 

## Methods

### NewTranscriptConsequence

`func NewTranscriptConsequence(transcriptId string, ) *TranscriptConsequence`

NewTranscriptConsequence instantiates a new TranscriptConsequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptConsequenceWithDefaults

`func NewTranscriptConsequenceWithDefaults() *TranscriptConsequence`

NewTranscriptConsequenceWithDefaults instantiates a new TranscriptConsequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAminoAcids

`func (o *TranscriptConsequence) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *TranscriptConsequence) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *TranscriptConsequence) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *TranscriptConsequence) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetCanonical

`func (o *TranscriptConsequence) GetCanonical() string`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *TranscriptConsequence) GetCanonicalOk() (*string, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *TranscriptConsequence) SetCanonical(v string)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *TranscriptConsequence) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetCodons

`func (o *TranscriptConsequence) GetCodons() string`

GetCodons returns the Codons field if non-nil, zero value otherwise.

### GetCodonsOk

`func (o *TranscriptConsequence) GetCodonsOk() (*string, bool)`

GetCodonsOk returns a tuple with the Codons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodons

`func (o *TranscriptConsequence) SetCodons(v string)`

SetCodons sets Codons field to given value.

### HasCodons

`func (o *TranscriptConsequence) HasCodons() bool`

HasCodons returns a boolean if a field has been set.

### GetConsequenceTerms

`func (o *TranscriptConsequence) GetConsequenceTerms() []string`

GetConsequenceTerms returns the ConsequenceTerms field if non-nil, zero value otherwise.

### GetConsequenceTermsOk

`func (o *TranscriptConsequence) GetConsequenceTermsOk() (*[]string, bool)`

GetConsequenceTermsOk returns a tuple with the ConsequenceTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsequenceTerms

`func (o *TranscriptConsequence) SetConsequenceTerms(v []string)`

SetConsequenceTerms sets ConsequenceTerms field to given value.

### HasConsequenceTerms

`func (o *TranscriptConsequence) HasConsequenceTerms() bool`

HasConsequenceTerms returns a boolean if a field has been set.

### GetExon

`func (o *TranscriptConsequence) GetExon() string`

GetExon returns the Exon field if non-nil, zero value otherwise.

### GetExonOk

`func (o *TranscriptConsequence) GetExonOk() (*string, bool)`

GetExonOk returns a tuple with the Exon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExon

`func (o *TranscriptConsequence) SetExon(v string)`

SetExon sets Exon field to given value.

### HasExon

`func (o *TranscriptConsequence) HasExon() bool`

HasExon returns a boolean if a field has been set.

### GetGeneId

`func (o *TranscriptConsequence) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *TranscriptConsequence) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *TranscriptConsequence) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *TranscriptConsequence) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetGeneSymbol

`func (o *TranscriptConsequence) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *TranscriptConsequence) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *TranscriptConsequence) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *TranscriptConsequence) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### GetHgncId

`func (o *TranscriptConsequence) GetHgncId() string`

GetHgncId returns the HgncId field if non-nil, zero value otherwise.

### GetHgncIdOk

`func (o *TranscriptConsequence) GetHgncIdOk() (*string, bool)`

GetHgncIdOk returns a tuple with the HgncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgncId

`func (o *TranscriptConsequence) SetHgncId(v string)`

SetHgncId sets HgncId field to given value.

### HasHgncId

`func (o *TranscriptConsequence) HasHgncId() bool`

HasHgncId returns a boolean if a field has been set.

### GetHgvsc

`func (o *TranscriptConsequence) GetHgvsc() string`

GetHgvsc returns the Hgvsc field if non-nil, zero value otherwise.

### GetHgvscOk

`func (o *TranscriptConsequence) GetHgvscOk() (*string, bool)`

GetHgvscOk returns a tuple with the Hgvsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsc

`func (o *TranscriptConsequence) SetHgvsc(v string)`

SetHgvsc sets Hgvsc field to given value.

### HasHgvsc

`func (o *TranscriptConsequence) HasHgvsc() bool`

HasHgvsc returns a boolean if a field has been set.

### GetHgvsg

`func (o *TranscriptConsequence) GetHgvsg() string`

GetHgvsg returns the Hgvsg field if non-nil, zero value otherwise.

### GetHgvsgOk

`func (o *TranscriptConsequence) GetHgvsgOk() (*string, bool)`

GetHgvsgOk returns a tuple with the Hgvsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsg

`func (o *TranscriptConsequence) SetHgvsg(v string)`

SetHgvsg sets Hgvsg field to given value.

### HasHgvsg

`func (o *TranscriptConsequence) HasHgvsg() bool`

HasHgvsg returns a boolean if a field has been set.

### GetHgvsp

`func (o *TranscriptConsequence) GetHgvsp() string`

GetHgvsp returns the Hgvsp field if non-nil, zero value otherwise.

### GetHgvspOk

`func (o *TranscriptConsequence) GetHgvspOk() (*string, bool)`

GetHgvspOk returns a tuple with the Hgvsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsp

`func (o *TranscriptConsequence) SetHgvsp(v string)`

SetHgvsp sets Hgvsp field to given value.

### HasHgvsp

`func (o *TranscriptConsequence) HasHgvsp() bool`

HasHgvsp returns a boolean if a field has been set.

### GetPolyphenPrediction

`func (o *TranscriptConsequence) GetPolyphenPrediction() string`

GetPolyphenPrediction returns the PolyphenPrediction field if non-nil, zero value otherwise.

### GetPolyphenPredictionOk

`func (o *TranscriptConsequence) GetPolyphenPredictionOk() (*string, bool)`

GetPolyphenPredictionOk returns a tuple with the PolyphenPrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyphenPrediction

`func (o *TranscriptConsequence) SetPolyphenPrediction(v string)`

SetPolyphenPrediction sets PolyphenPrediction field to given value.

### HasPolyphenPrediction

`func (o *TranscriptConsequence) HasPolyphenPrediction() bool`

HasPolyphenPrediction returns a boolean if a field has been set.

### GetPolyphenScore

`func (o *TranscriptConsequence) GetPolyphenScore() float64`

GetPolyphenScore returns the PolyphenScore field if non-nil, zero value otherwise.

### GetPolyphenScoreOk

`func (o *TranscriptConsequence) GetPolyphenScoreOk() (*float64, bool)`

GetPolyphenScoreOk returns a tuple with the PolyphenScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyphenScore

`func (o *TranscriptConsequence) SetPolyphenScore(v float64)`

SetPolyphenScore sets PolyphenScore field to given value.

### HasPolyphenScore

`func (o *TranscriptConsequence) HasPolyphenScore() bool`

HasPolyphenScore returns a boolean if a field has been set.

### GetProteinEnd

`func (o *TranscriptConsequence) GetProteinEnd() int32`

GetProteinEnd returns the ProteinEnd field if non-nil, zero value otherwise.

### GetProteinEndOk

`func (o *TranscriptConsequence) GetProteinEndOk() (*int32, bool)`

GetProteinEndOk returns a tuple with the ProteinEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinEnd

`func (o *TranscriptConsequence) SetProteinEnd(v int32)`

SetProteinEnd sets ProteinEnd field to given value.

### HasProteinEnd

`func (o *TranscriptConsequence) HasProteinEnd() bool`

HasProteinEnd returns a boolean if a field has been set.

### GetProteinId

`func (o *TranscriptConsequence) GetProteinId() string`

GetProteinId returns the ProteinId field if non-nil, zero value otherwise.

### GetProteinIdOk

`func (o *TranscriptConsequence) GetProteinIdOk() (*string, bool)`

GetProteinIdOk returns a tuple with the ProteinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinId

`func (o *TranscriptConsequence) SetProteinId(v string)`

SetProteinId sets ProteinId field to given value.

### HasProteinId

`func (o *TranscriptConsequence) HasProteinId() bool`

HasProteinId returns a boolean if a field has been set.

### GetProteinStart

`func (o *TranscriptConsequence) GetProteinStart() int32`

GetProteinStart returns the ProteinStart field if non-nil, zero value otherwise.

### GetProteinStartOk

`func (o *TranscriptConsequence) GetProteinStartOk() (*int32, bool)`

GetProteinStartOk returns a tuple with the ProteinStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinStart

`func (o *TranscriptConsequence) SetProteinStart(v int32)`

SetProteinStart sets ProteinStart field to given value.

### HasProteinStart

`func (o *TranscriptConsequence) HasProteinStart() bool`

HasProteinStart returns a boolean if a field has been set.

### GetRefseqTranscriptIds

`func (o *TranscriptConsequence) GetRefseqTranscriptIds() []string`

GetRefseqTranscriptIds returns the RefseqTranscriptIds field if non-nil, zero value otherwise.

### GetRefseqTranscriptIdsOk

`func (o *TranscriptConsequence) GetRefseqTranscriptIdsOk() (*[]string, bool)`

GetRefseqTranscriptIdsOk returns a tuple with the RefseqTranscriptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqTranscriptIds

`func (o *TranscriptConsequence) SetRefseqTranscriptIds(v []string)`

SetRefseqTranscriptIds sets RefseqTranscriptIds field to given value.

### HasRefseqTranscriptIds

`func (o *TranscriptConsequence) HasRefseqTranscriptIds() bool`

HasRefseqTranscriptIds returns a boolean if a field has been set.

### GetSiftPrediction

`func (o *TranscriptConsequence) GetSiftPrediction() string`

GetSiftPrediction returns the SiftPrediction field if non-nil, zero value otherwise.

### GetSiftPredictionOk

`func (o *TranscriptConsequence) GetSiftPredictionOk() (*string, bool)`

GetSiftPredictionOk returns a tuple with the SiftPrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiftPrediction

`func (o *TranscriptConsequence) SetSiftPrediction(v string)`

SetSiftPrediction sets SiftPrediction field to given value.

### HasSiftPrediction

`func (o *TranscriptConsequence) HasSiftPrediction() bool`

HasSiftPrediction returns a boolean if a field has been set.

### GetSiftScore

`func (o *TranscriptConsequence) GetSiftScore() float64`

GetSiftScore returns the SiftScore field if non-nil, zero value otherwise.

### GetSiftScoreOk

`func (o *TranscriptConsequence) GetSiftScoreOk() (*float64, bool)`

GetSiftScoreOk returns a tuple with the SiftScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiftScore

`func (o *TranscriptConsequence) SetSiftScore(v float64)`

SetSiftScore sets SiftScore field to given value.

### HasSiftScore

`func (o *TranscriptConsequence) HasSiftScore() bool`

HasSiftScore returns a boolean if a field has been set.

### GetTranscriptId

`func (o *TranscriptConsequence) GetTranscriptId() string`

GetTranscriptId returns the TranscriptId field if non-nil, zero value otherwise.

### GetTranscriptIdOk

`func (o *TranscriptConsequence) GetTranscriptIdOk() (*string, bool)`

GetTranscriptIdOk returns a tuple with the TranscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptId

`func (o *TranscriptConsequence) SetTranscriptId(v string)`

SetTranscriptId sets TranscriptId field to given value.


### GetUniprotId

`func (o *TranscriptConsequence) GetUniprotId() string`

GetUniprotId returns the UniprotId field if non-nil, zero value otherwise.

### GetUniprotIdOk

`func (o *TranscriptConsequence) GetUniprotIdOk() (*string, bool)`

GetUniprotIdOk returns a tuple with the UniprotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotId

`func (o *TranscriptConsequence) SetUniprotId(v string)`

SetUniprotId sets UniprotId field to given value.

### HasUniprotId

`func (o *TranscriptConsequence) HasUniprotId() bool`

HasUniprotId returns a boolean if a field has been set.

### GetVariantAllele

`func (o *TranscriptConsequence) GetVariantAllele() string`

GetVariantAllele returns the VariantAllele field if non-nil, zero value otherwise.

### GetVariantAlleleOk

`func (o *TranscriptConsequence) GetVariantAlleleOk() (*string, bool)`

GetVariantAlleleOk returns a tuple with the VariantAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantAllele

`func (o *TranscriptConsequence) SetVariantAllele(v string)`

SetVariantAllele sets VariantAllele field to given value.

### HasVariantAllele

`func (o *TranscriptConsequence) HasVariantAllele() bool`

HasVariantAllele returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


