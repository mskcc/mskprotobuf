# TranscriptConsequenceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AminoAcidAlt** | Pointer to **string** | Alt Amino Acid | [optional] 
**AminoAcidRef** | Pointer to **string** | Reference Amino Acid | [optional] 
**AminoAcids** | Pointer to **string** | Amino acids change | [optional] 
**CodonChange** | Pointer to **string** | Codon change | [optional] 
**ConsequenceTerms** | Pointer to **string** | Consequence terms (comma separated) | [optional] 
**EntrezGeneId** | Pointer to **string** | Entrez gene id | [optional] 
**Exon** | Pointer to **string** |  | [optional] 
**Hgvsc** | Pointer to **string** | HGVSc | [optional] 
**Hgvsp** | Pointer to **string** | HGVSp | [optional] 
**HgvspShort** | Pointer to **string** | HGVSp short | [optional] 
**HugoGeneSymbol** | Pointer to **string** | Hugo gene symbol | [optional] 
**IsVue** | Pointer to **bool** |  | [optional] 
**PolyphenPrediction** | Pointer to **string** | Polyphen Prediction | [optional] 
**PolyphenScore** | Pointer to **float64** | Polyphen Score | [optional] 
**ProteinPosition** | Pointer to [**IntegerRange**](IntegerRange.md) |  | [optional] 
**RefSeq** | Pointer to **string** | RefSeq id | [optional] 
**SiftPrediction** | Pointer to **string** | Sift Prediction | [optional] 
**SiftScore** | Pointer to **float64** | Sift Score | [optional] 
**TranscriptId** | **string** | Transcript id | 
**UniprotId** | Pointer to **string** | Uniprot ID | [optional] 
**VariantClassification** | Pointer to **string** | Variant classification | [optional] 

## Methods

### NewTranscriptConsequenceSummary

`func NewTranscriptConsequenceSummary(transcriptId string, ) *TranscriptConsequenceSummary`

NewTranscriptConsequenceSummary instantiates a new TranscriptConsequenceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptConsequenceSummaryWithDefaults

`func NewTranscriptConsequenceSummaryWithDefaults() *TranscriptConsequenceSummary`

NewTranscriptConsequenceSummaryWithDefaults instantiates a new TranscriptConsequenceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAminoAcidAlt

`func (o *TranscriptConsequenceSummary) GetAminoAcidAlt() string`

GetAminoAcidAlt returns the AminoAcidAlt field if non-nil, zero value otherwise.

### GetAminoAcidAltOk

`func (o *TranscriptConsequenceSummary) GetAminoAcidAltOk() (*string, bool)`

GetAminoAcidAltOk returns a tuple with the AminoAcidAlt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidAlt

`func (o *TranscriptConsequenceSummary) SetAminoAcidAlt(v string)`

SetAminoAcidAlt sets AminoAcidAlt field to given value.

### HasAminoAcidAlt

`func (o *TranscriptConsequenceSummary) HasAminoAcidAlt() bool`

HasAminoAcidAlt returns a boolean if a field has been set.

### GetAminoAcidRef

`func (o *TranscriptConsequenceSummary) GetAminoAcidRef() string`

GetAminoAcidRef returns the AminoAcidRef field if non-nil, zero value otherwise.

### GetAminoAcidRefOk

`func (o *TranscriptConsequenceSummary) GetAminoAcidRefOk() (*string, bool)`

GetAminoAcidRefOk returns a tuple with the AminoAcidRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidRef

`func (o *TranscriptConsequenceSummary) SetAminoAcidRef(v string)`

SetAminoAcidRef sets AminoAcidRef field to given value.

### HasAminoAcidRef

`func (o *TranscriptConsequenceSummary) HasAminoAcidRef() bool`

HasAminoAcidRef returns a boolean if a field has been set.

### GetAminoAcids

`func (o *TranscriptConsequenceSummary) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *TranscriptConsequenceSummary) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *TranscriptConsequenceSummary) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *TranscriptConsequenceSummary) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetCodonChange

`func (o *TranscriptConsequenceSummary) GetCodonChange() string`

GetCodonChange returns the CodonChange field if non-nil, zero value otherwise.

### GetCodonChangeOk

`func (o *TranscriptConsequenceSummary) GetCodonChangeOk() (*string, bool)`

GetCodonChangeOk returns a tuple with the CodonChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodonChange

`func (o *TranscriptConsequenceSummary) SetCodonChange(v string)`

SetCodonChange sets CodonChange field to given value.

### HasCodonChange

`func (o *TranscriptConsequenceSummary) HasCodonChange() bool`

HasCodonChange returns a boolean if a field has been set.

### GetConsequenceTerms

`func (o *TranscriptConsequenceSummary) GetConsequenceTerms() string`

GetConsequenceTerms returns the ConsequenceTerms field if non-nil, zero value otherwise.

### GetConsequenceTermsOk

`func (o *TranscriptConsequenceSummary) GetConsequenceTermsOk() (*string, bool)`

GetConsequenceTermsOk returns a tuple with the ConsequenceTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsequenceTerms

`func (o *TranscriptConsequenceSummary) SetConsequenceTerms(v string)`

SetConsequenceTerms sets ConsequenceTerms field to given value.

### HasConsequenceTerms

`func (o *TranscriptConsequenceSummary) HasConsequenceTerms() bool`

HasConsequenceTerms returns a boolean if a field has been set.

### GetEntrezGeneId

`func (o *TranscriptConsequenceSummary) GetEntrezGeneId() string`

GetEntrezGeneId returns the EntrezGeneId field if non-nil, zero value otherwise.

### GetEntrezGeneIdOk

`func (o *TranscriptConsequenceSummary) GetEntrezGeneIdOk() (*string, bool)`

GetEntrezGeneIdOk returns a tuple with the EntrezGeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrezGeneId

`func (o *TranscriptConsequenceSummary) SetEntrezGeneId(v string)`

SetEntrezGeneId sets EntrezGeneId field to given value.

### HasEntrezGeneId

`func (o *TranscriptConsequenceSummary) HasEntrezGeneId() bool`

HasEntrezGeneId returns a boolean if a field has been set.

### GetExon

`func (o *TranscriptConsequenceSummary) GetExon() string`

GetExon returns the Exon field if non-nil, zero value otherwise.

### GetExonOk

`func (o *TranscriptConsequenceSummary) GetExonOk() (*string, bool)`

GetExonOk returns a tuple with the Exon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExon

`func (o *TranscriptConsequenceSummary) SetExon(v string)`

SetExon sets Exon field to given value.

### HasExon

`func (o *TranscriptConsequenceSummary) HasExon() bool`

HasExon returns a boolean if a field has been set.

### GetHgvsc

`func (o *TranscriptConsequenceSummary) GetHgvsc() string`

GetHgvsc returns the Hgvsc field if non-nil, zero value otherwise.

### GetHgvscOk

`func (o *TranscriptConsequenceSummary) GetHgvscOk() (*string, bool)`

GetHgvscOk returns a tuple with the Hgvsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsc

`func (o *TranscriptConsequenceSummary) SetHgvsc(v string)`

SetHgvsc sets Hgvsc field to given value.

### HasHgvsc

`func (o *TranscriptConsequenceSummary) HasHgvsc() bool`

HasHgvsc returns a boolean if a field has been set.

### GetHgvsp

`func (o *TranscriptConsequenceSummary) GetHgvsp() string`

GetHgvsp returns the Hgvsp field if non-nil, zero value otherwise.

### GetHgvspOk

`func (o *TranscriptConsequenceSummary) GetHgvspOk() (*string, bool)`

GetHgvspOk returns a tuple with the Hgvsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsp

`func (o *TranscriptConsequenceSummary) SetHgvsp(v string)`

SetHgvsp sets Hgvsp field to given value.

### HasHgvsp

`func (o *TranscriptConsequenceSummary) HasHgvsp() bool`

HasHgvsp returns a boolean if a field has been set.

### GetHgvspShort

`func (o *TranscriptConsequenceSummary) GetHgvspShort() string`

GetHgvspShort returns the HgvspShort field if non-nil, zero value otherwise.

### GetHgvspShortOk

`func (o *TranscriptConsequenceSummary) GetHgvspShortOk() (*string, bool)`

GetHgvspShortOk returns a tuple with the HgvspShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvspShort

`func (o *TranscriptConsequenceSummary) SetHgvspShort(v string)`

SetHgvspShort sets HgvspShort field to given value.

### HasHgvspShort

`func (o *TranscriptConsequenceSummary) HasHgvspShort() bool`

HasHgvspShort returns a boolean if a field has been set.

### GetHugoGeneSymbol

`func (o *TranscriptConsequenceSummary) GetHugoGeneSymbol() string`

GetHugoGeneSymbol returns the HugoGeneSymbol field if non-nil, zero value otherwise.

### GetHugoGeneSymbolOk

`func (o *TranscriptConsequenceSummary) GetHugoGeneSymbolOk() (*string, bool)`

GetHugoGeneSymbolOk returns a tuple with the HugoGeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugoGeneSymbol

`func (o *TranscriptConsequenceSummary) SetHugoGeneSymbol(v string)`

SetHugoGeneSymbol sets HugoGeneSymbol field to given value.

### HasHugoGeneSymbol

`func (o *TranscriptConsequenceSummary) HasHugoGeneSymbol() bool`

HasHugoGeneSymbol returns a boolean if a field has been set.

### GetIsVue

`func (o *TranscriptConsequenceSummary) GetIsVue() bool`

GetIsVue returns the IsVue field if non-nil, zero value otherwise.

### GetIsVueOk

`func (o *TranscriptConsequenceSummary) GetIsVueOk() (*bool, bool)`

GetIsVueOk returns a tuple with the IsVue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVue

`func (o *TranscriptConsequenceSummary) SetIsVue(v bool)`

SetIsVue sets IsVue field to given value.

### HasIsVue

`func (o *TranscriptConsequenceSummary) HasIsVue() bool`

HasIsVue returns a boolean if a field has been set.

### GetPolyphenPrediction

`func (o *TranscriptConsequenceSummary) GetPolyphenPrediction() string`

GetPolyphenPrediction returns the PolyphenPrediction field if non-nil, zero value otherwise.

### GetPolyphenPredictionOk

`func (o *TranscriptConsequenceSummary) GetPolyphenPredictionOk() (*string, bool)`

GetPolyphenPredictionOk returns a tuple with the PolyphenPrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyphenPrediction

`func (o *TranscriptConsequenceSummary) SetPolyphenPrediction(v string)`

SetPolyphenPrediction sets PolyphenPrediction field to given value.

### HasPolyphenPrediction

`func (o *TranscriptConsequenceSummary) HasPolyphenPrediction() bool`

HasPolyphenPrediction returns a boolean if a field has been set.

### GetPolyphenScore

`func (o *TranscriptConsequenceSummary) GetPolyphenScore() float64`

GetPolyphenScore returns the PolyphenScore field if non-nil, zero value otherwise.

### GetPolyphenScoreOk

`func (o *TranscriptConsequenceSummary) GetPolyphenScoreOk() (*float64, bool)`

GetPolyphenScoreOk returns a tuple with the PolyphenScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyphenScore

`func (o *TranscriptConsequenceSummary) SetPolyphenScore(v float64)`

SetPolyphenScore sets PolyphenScore field to given value.

### HasPolyphenScore

`func (o *TranscriptConsequenceSummary) HasPolyphenScore() bool`

HasPolyphenScore returns a boolean if a field has been set.

### GetProteinPosition

`func (o *TranscriptConsequenceSummary) GetProteinPosition() IntegerRange`

GetProteinPosition returns the ProteinPosition field if non-nil, zero value otherwise.

### GetProteinPositionOk

`func (o *TranscriptConsequenceSummary) GetProteinPositionOk() (*IntegerRange, bool)`

GetProteinPositionOk returns a tuple with the ProteinPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinPosition

`func (o *TranscriptConsequenceSummary) SetProteinPosition(v IntegerRange)`

SetProteinPosition sets ProteinPosition field to given value.

### HasProteinPosition

`func (o *TranscriptConsequenceSummary) HasProteinPosition() bool`

HasProteinPosition returns a boolean if a field has been set.

### GetRefSeq

`func (o *TranscriptConsequenceSummary) GetRefSeq() string`

GetRefSeq returns the RefSeq field if non-nil, zero value otherwise.

### GetRefSeqOk

`func (o *TranscriptConsequenceSummary) GetRefSeqOk() (*string, bool)`

GetRefSeqOk returns a tuple with the RefSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefSeq

`func (o *TranscriptConsequenceSummary) SetRefSeq(v string)`

SetRefSeq sets RefSeq field to given value.

### HasRefSeq

`func (o *TranscriptConsequenceSummary) HasRefSeq() bool`

HasRefSeq returns a boolean if a field has been set.

### GetSiftPrediction

`func (o *TranscriptConsequenceSummary) GetSiftPrediction() string`

GetSiftPrediction returns the SiftPrediction field if non-nil, zero value otherwise.

### GetSiftPredictionOk

`func (o *TranscriptConsequenceSummary) GetSiftPredictionOk() (*string, bool)`

GetSiftPredictionOk returns a tuple with the SiftPrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiftPrediction

`func (o *TranscriptConsequenceSummary) SetSiftPrediction(v string)`

SetSiftPrediction sets SiftPrediction field to given value.

### HasSiftPrediction

`func (o *TranscriptConsequenceSummary) HasSiftPrediction() bool`

HasSiftPrediction returns a boolean if a field has been set.

### GetSiftScore

`func (o *TranscriptConsequenceSummary) GetSiftScore() float64`

GetSiftScore returns the SiftScore field if non-nil, zero value otherwise.

### GetSiftScoreOk

`func (o *TranscriptConsequenceSummary) GetSiftScoreOk() (*float64, bool)`

GetSiftScoreOk returns a tuple with the SiftScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiftScore

`func (o *TranscriptConsequenceSummary) SetSiftScore(v float64)`

SetSiftScore sets SiftScore field to given value.

### HasSiftScore

`func (o *TranscriptConsequenceSummary) HasSiftScore() bool`

HasSiftScore returns a boolean if a field has been set.

### GetTranscriptId

`func (o *TranscriptConsequenceSummary) GetTranscriptId() string`

GetTranscriptId returns the TranscriptId field if non-nil, zero value otherwise.

### GetTranscriptIdOk

`func (o *TranscriptConsequenceSummary) GetTranscriptIdOk() (*string, bool)`

GetTranscriptIdOk returns a tuple with the TranscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptId

`func (o *TranscriptConsequenceSummary) SetTranscriptId(v string)`

SetTranscriptId sets TranscriptId field to given value.


### GetUniprotId

`func (o *TranscriptConsequenceSummary) GetUniprotId() string`

GetUniprotId returns the UniprotId field if non-nil, zero value otherwise.

### GetUniprotIdOk

`func (o *TranscriptConsequenceSummary) GetUniprotIdOk() (*string, bool)`

GetUniprotIdOk returns a tuple with the UniprotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotId

`func (o *TranscriptConsequenceSummary) SetUniprotId(v string)`

SetUniprotId sets UniprotId field to given value.

### HasUniprotId

`func (o *TranscriptConsequenceSummary) HasUniprotId() bool`

HasUniprotId returns a boolean if a field has been set.

### GetVariantClassification

`func (o *TranscriptConsequenceSummary) GetVariantClassification() string`

GetVariantClassification returns the VariantClassification field if non-nil, zero value otherwise.

### GetVariantClassificationOk

`func (o *TranscriptConsequenceSummary) GetVariantClassificationOk() (*string, bool)`

GetVariantClassificationOk returns a tuple with the VariantClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantClassification

`func (o *TranscriptConsequenceSummary) SetVariantClassification(v string)`

SetVariantClassification sets VariantClassification field to given value.

### HasVariantClassification

`func (o *TranscriptConsequenceSummary) HasVariantClassification() bool`

HasVariantClassification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


