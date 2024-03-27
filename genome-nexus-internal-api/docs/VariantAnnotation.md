# VariantAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlleleString** | Pointer to **string** | Allele string (e.g: A/T) | [optional] 
**AnnotationJSON** | Pointer to **string** | Annotation data as JSON string | [optional] 
**AnnotationSummary** | Pointer to [**VariantAnnotationSummary**](VariantAnnotationSummary.md) |  | [optional] 
**AssemblyName** | Pointer to **string** | NCBI build number | [optional] 
**Clinvar** | Pointer to [**ClinvarAnnotation**](ClinvarAnnotation.md) |  | [optional] 
**ColocatedVariants** | Pointer to [**[]ColocatedVariant**](ColocatedVariant.md) |  | [optional] 
**End** | Pointer to **int32** | End position | [optional] 
**Hgvsg** | Pointer to **string** |  | [optional] 
**Hotspots** | Pointer to [**HotspotAnnotation**](HotspotAnnotation.md) |  | [optional] 
**Id** | **string** | Variant id | 
**IntergenicConsequences** | [**[]IntergenicConsequences**](IntergenicConsequences.md) | intergenicConsequences | 
**MostSevereConsequence** | Pointer to **string** | Most severe consequence | [optional] 
**MutationAssessor** | Pointer to [**MutationAssessorAnnotation**](MutationAssessorAnnotation.md) |  | [optional] 
**MyVariantInfo** | Pointer to [**MyVariantInfoAnnotation**](MyVariantInfoAnnotation.md) |  | [optional] 
**NucleotideContext** | Pointer to [**NucleotideContextAnnotation**](NucleotideContextAnnotation.md) |  | [optional] 
**Oncokb** | Pointer to [**OncokbAnnotation**](OncokbAnnotation.md) |  | [optional] 
**OriginalVariantQuery** | **string** | Original variant query | 
**Ptms** | Pointer to [**PtmAnnotation**](PtmAnnotation.md) |  | [optional] 
**SeqRegionName** | Pointer to **string** | Chromosome | [optional] 
**SignalAnnotation** | Pointer to [**SignalAnnotation**](SignalAnnotation.md) |  | [optional] 
**Start** | Pointer to **int32** | Start position | [optional] 
**Strand** | Pointer to **int32** | Strand (negative or positive) | [optional] 
**SuccessfullyAnnotated** | Pointer to **bool** | Status flag indicating whether variant was succesfully annotated | [optional] 
**TranscriptConsequences** | Pointer to [**[]TranscriptConsequence**](TranscriptConsequence.md) | List of transcripts | [optional] 
**Variant** | **string** | Variant key | 

## Methods

### NewVariantAnnotation

`func NewVariantAnnotation(id string, intergenicConsequences []IntergenicConsequences, originalVariantQuery string, variant string, ) *VariantAnnotation`

NewVariantAnnotation instantiates a new VariantAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantAnnotationWithDefaults

`func NewVariantAnnotationWithDefaults() *VariantAnnotation`

NewVariantAnnotationWithDefaults instantiates a new VariantAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlleleString

`func (o *VariantAnnotation) GetAlleleString() string`

GetAlleleString returns the AlleleString field if non-nil, zero value otherwise.

### GetAlleleStringOk

`func (o *VariantAnnotation) GetAlleleStringOk() (*string, bool)`

GetAlleleStringOk returns a tuple with the AlleleString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlleleString

`func (o *VariantAnnotation) SetAlleleString(v string)`

SetAlleleString sets AlleleString field to given value.

### HasAlleleString

`func (o *VariantAnnotation) HasAlleleString() bool`

HasAlleleString returns a boolean if a field has been set.

### GetAnnotationJSON

`func (o *VariantAnnotation) GetAnnotationJSON() string`

GetAnnotationJSON returns the AnnotationJSON field if non-nil, zero value otherwise.

### GetAnnotationJSONOk

`func (o *VariantAnnotation) GetAnnotationJSONOk() (*string, bool)`

GetAnnotationJSONOk returns a tuple with the AnnotationJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationJSON

`func (o *VariantAnnotation) SetAnnotationJSON(v string)`

SetAnnotationJSON sets AnnotationJSON field to given value.

### HasAnnotationJSON

`func (o *VariantAnnotation) HasAnnotationJSON() bool`

HasAnnotationJSON returns a boolean if a field has been set.

### GetAnnotationSummary

`func (o *VariantAnnotation) GetAnnotationSummary() VariantAnnotationSummary`

GetAnnotationSummary returns the AnnotationSummary field if non-nil, zero value otherwise.

### GetAnnotationSummaryOk

`func (o *VariantAnnotation) GetAnnotationSummaryOk() (*VariantAnnotationSummary, bool)`

GetAnnotationSummaryOk returns a tuple with the AnnotationSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationSummary

`func (o *VariantAnnotation) SetAnnotationSummary(v VariantAnnotationSummary)`

SetAnnotationSummary sets AnnotationSummary field to given value.

### HasAnnotationSummary

`func (o *VariantAnnotation) HasAnnotationSummary() bool`

HasAnnotationSummary returns a boolean if a field has been set.

### GetAssemblyName

`func (o *VariantAnnotation) GetAssemblyName() string`

GetAssemblyName returns the AssemblyName field if non-nil, zero value otherwise.

### GetAssemblyNameOk

`func (o *VariantAnnotation) GetAssemblyNameOk() (*string, bool)`

GetAssemblyNameOk returns a tuple with the AssemblyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyName

`func (o *VariantAnnotation) SetAssemblyName(v string)`

SetAssemblyName sets AssemblyName field to given value.

### HasAssemblyName

`func (o *VariantAnnotation) HasAssemblyName() bool`

HasAssemblyName returns a boolean if a field has been set.

### GetClinvar

`func (o *VariantAnnotation) GetClinvar() ClinvarAnnotation`

GetClinvar returns the Clinvar field if non-nil, zero value otherwise.

### GetClinvarOk

`func (o *VariantAnnotation) GetClinvarOk() (*ClinvarAnnotation, bool)`

GetClinvarOk returns a tuple with the Clinvar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinvar

`func (o *VariantAnnotation) SetClinvar(v ClinvarAnnotation)`

SetClinvar sets Clinvar field to given value.

### HasClinvar

`func (o *VariantAnnotation) HasClinvar() bool`

HasClinvar returns a boolean if a field has been set.

### GetColocatedVariants

`func (o *VariantAnnotation) GetColocatedVariants() []ColocatedVariant`

GetColocatedVariants returns the ColocatedVariants field if non-nil, zero value otherwise.

### GetColocatedVariantsOk

`func (o *VariantAnnotation) GetColocatedVariantsOk() (*[]ColocatedVariant, bool)`

GetColocatedVariantsOk returns a tuple with the ColocatedVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColocatedVariants

`func (o *VariantAnnotation) SetColocatedVariants(v []ColocatedVariant)`

SetColocatedVariants sets ColocatedVariants field to given value.

### HasColocatedVariants

`func (o *VariantAnnotation) HasColocatedVariants() bool`

HasColocatedVariants returns a boolean if a field has been set.

### GetEnd

`func (o *VariantAnnotation) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *VariantAnnotation) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *VariantAnnotation) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *VariantAnnotation) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetHgvsg

`func (o *VariantAnnotation) GetHgvsg() string`

GetHgvsg returns the Hgvsg field if non-nil, zero value otherwise.

### GetHgvsgOk

`func (o *VariantAnnotation) GetHgvsgOk() (*string, bool)`

GetHgvsgOk returns a tuple with the Hgvsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvsg

`func (o *VariantAnnotation) SetHgvsg(v string)`

SetHgvsg sets Hgvsg field to given value.

### HasHgvsg

`func (o *VariantAnnotation) HasHgvsg() bool`

HasHgvsg returns a boolean if a field has been set.

### GetHotspots

`func (o *VariantAnnotation) GetHotspots() HotspotAnnotation`

GetHotspots returns the Hotspots field if non-nil, zero value otherwise.

### GetHotspotsOk

`func (o *VariantAnnotation) GetHotspotsOk() (*HotspotAnnotation, bool)`

GetHotspotsOk returns a tuple with the Hotspots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspots

`func (o *VariantAnnotation) SetHotspots(v HotspotAnnotation)`

SetHotspots sets Hotspots field to given value.

### HasHotspots

`func (o *VariantAnnotation) HasHotspots() bool`

HasHotspots returns a boolean if a field has been set.

### GetId

`func (o *VariantAnnotation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariantAnnotation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariantAnnotation) SetId(v string)`

SetId sets Id field to given value.


### GetIntergenicConsequences

`func (o *VariantAnnotation) GetIntergenicConsequences() []IntergenicConsequences`

GetIntergenicConsequences returns the IntergenicConsequences field if non-nil, zero value otherwise.

### GetIntergenicConsequencesOk

`func (o *VariantAnnotation) GetIntergenicConsequencesOk() (*[]IntergenicConsequences, bool)`

GetIntergenicConsequencesOk returns a tuple with the IntergenicConsequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntergenicConsequences

`func (o *VariantAnnotation) SetIntergenicConsequences(v []IntergenicConsequences)`

SetIntergenicConsequences sets IntergenicConsequences field to given value.


### GetMostSevereConsequence

`func (o *VariantAnnotation) GetMostSevereConsequence() string`

GetMostSevereConsequence returns the MostSevereConsequence field if non-nil, zero value otherwise.

### GetMostSevereConsequenceOk

`func (o *VariantAnnotation) GetMostSevereConsequenceOk() (*string, bool)`

GetMostSevereConsequenceOk returns a tuple with the MostSevereConsequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostSevereConsequence

`func (o *VariantAnnotation) SetMostSevereConsequence(v string)`

SetMostSevereConsequence sets MostSevereConsequence field to given value.

### HasMostSevereConsequence

`func (o *VariantAnnotation) HasMostSevereConsequence() bool`

HasMostSevereConsequence returns a boolean if a field has been set.

### GetMutationAssessor

`func (o *VariantAnnotation) GetMutationAssessor() MutationAssessorAnnotation`

GetMutationAssessor returns the MutationAssessor field if non-nil, zero value otherwise.

### GetMutationAssessorOk

`func (o *VariantAnnotation) GetMutationAssessorOk() (*MutationAssessorAnnotation, bool)`

GetMutationAssessorOk returns a tuple with the MutationAssessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutationAssessor

`func (o *VariantAnnotation) SetMutationAssessor(v MutationAssessorAnnotation)`

SetMutationAssessor sets MutationAssessor field to given value.

### HasMutationAssessor

`func (o *VariantAnnotation) HasMutationAssessor() bool`

HasMutationAssessor returns a boolean if a field has been set.

### GetMyVariantInfo

`func (o *VariantAnnotation) GetMyVariantInfo() MyVariantInfoAnnotation`

GetMyVariantInfo returns the MyVariantInfo field if non-nil, zero value otherwise.

### GetMyVariantInfoOk

`func (o *VariantAnnotation) GetMyVariantInfoOk() (*MyVariantInfoAnnotation, bool)`

GetMyVariantInfoOk returns a tuple with the MyVariantInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyVariantInfo

`func (o *VariantAnnotation) SetMyVariantInfo(v MyVariantInfoAnnotation)`

SetMyVariantInfo sets MyVariantInfo field to given value.

### HasMyVariantInfo

`func (o *VariantAnnotation) HasMyVariantInfo() bool`

HasMyVariantInfo returns a boolean if a field has been set.

### GetNucleotideContext

`func (o *VariantAnnotation) GetNucleotideContext() NucleotideContextAnnotation`

GetNucleotideContext returns the NucleotideContext field if non-nil, zero value otherwise.

### GetNucleotideContextOk

`func (o *VariantAnnotation) GetNucleotideContextOk() (*NucleotideContextAnnotation, bool)`

GetNucleotideContextOk returns a tuple with the NucleotideContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotideContext

`func (o *VariantAnnotation) SetNucleotideContext(v NucleotideContextAnnotation)`

SetNucleotideContext sets NucleotideContext field to given value.

### HasNucleotideContext

`func (o *VariantAnnotation) HasNucleotideContext() bool`

HasNucleotideContext returns a boolean if a field has been set.

### GetOncokb

`func (o *VariantAnnotation) GetOncokb() OncokbAnnotation`

GetOncokb returns the Oncokb field if non-nil, zero value otherwise.

### GetOncokbOk

`func (o *VariantAnnotation) GetOncokbOk() (*OncokbAnnotation, bool)`

GetOncokbOk returns a tuple with the Oncokb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOncokb

`func (o *VariantAnnotation) SetOncokb(v OncokbAnnotation)`

SetOncokb sets Oncokb field to given value.

### HasOncokb

`func (o *VariantAnnotation) HasOncokb() bool`

HasOncokb returns a boolean if a field has been set.

### GetOriginalVariantQuery

`func (o *VariantAnnotation) GetOriginalVariantQuery() string`

GetOriginalVariantQuery returns the OriginalVariantQuery field if non-nil, zero value otherwise.

### GetOriginalVariantQueryOk

`func (o *VariantAnnotation) GetOriginalVariantQueryOk() (*string, bool)`

GetOriginalVariantQueryOk returns a tuple with the OriginalVariantQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVariantQuery

`func (o *VariantAnnotation) SetOriginalVariantQuery(v string)`

SetOriginalVariantQuery sets OriginalVariantQuery field to given value.


### GetPtms

`func (o *VariantAnnotation) GetPtms() PtmAnnotation`

GetPtms returns the Ptms field if non-nil, zero value otherwise.

### GetPtmsOk

`func (o *VariantAnnotation) GetPtmsOk() (*PtmAnnotation, bool)`

GetPtmsOk returns a tuple with the Ptms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtms

`func (o *VariantAnnotation) SetPtms(v PtmAnnotation)`

SetPtms sets Ptms field to given value.

### HasPtms

`func (o *VariantAnnotation) HasPtms() bool`

HasPtms returns a boolean if a field has been set.

### GetSeqRegionName

`func (o *VariantAnnotation) GetSeqRegionName() string`

GetSeqRegionName returns the SeqRegionName field if non-nil, zero value otherwise.

### GetSeqRegionNameOk

`func (o *VariantAnnotation) GetSeqRegionNameOk() (*string, bool)`

GetSeqRegionNameOk returns a tuple with the SeqRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqRegionName

`func (o *VariantAnnotation) SetSeqRegionName(v string)`

SetSeqRegionName sets SeqRegionName field to given value.

### HasSeqRegionName

`func (o *VariantAnnotation) HasSeqRegionName() bool`

HasSeqRegionName returns a boolean if a field has been set.

### GetSignalAnnotation

`func (o *VariantAnnotation) GetSignalAnnotation() SignalAnnotation`

GetSignalAnnotation returns the SignalAnnotation field if non-nil, zero value otherwise.

### GetSignalAnnotationOk

`func (o *VariantAnnotation) GetSignalAnnotationOk() (*SignalAnnotation, bool)`

GetSignalAnnotationOk returns a tuple with the SignalAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalAnnotation

`func (o *VariantAnnotation) SetSignalAnnotation(v SignalAnnotation)`

SetSignalAnnotation sets SignalAnnotation field to given value.

### HasSignalAnnotation

`func (o *VariantAnnotation) HasSignalAnnotation() bool`

HasSignalAnnotation returns a boolean if a field has been set.

### GetStart

`func (o *VariantAnnotation) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *VariantAnnotation) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *VariantAnnotation) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *VariantAnnotation) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStrand

`func (o *VariantAnnotation) GetStrand() int32`

GetStrand returns the Strand field if non-nil, zero value otherwise.

### GetStrandOk

`func (o *VariantAnnotation) GetStrandOk() (*int32, bool)`

GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrand

`func (o *VariantAnnotation) SetStrand(v int32)`

SetStrand sets Strand field to given value.

### HasStrand

`func (o *VariantAnnotation) HasStrand() bool`

HasStrand returns a boolean if a field has been set.

### GetSuccessfullyAnnotated

`func (o *VariantAnnotation) GetSuccessfullyAnnotated() bool`

GetSuccessfullyAnnotated returns the SuccessfullyAnnotated field if non-nil, zero value otherwise.

### GetSuccessfullyAnnotatedOk

`func (o *VariantAnnotation) GetSuccessfullyAnnotatedOk() (*bool, bool)`

GetSuccessfullyAnnotatedOk returns a tuple with the SuccessfullyAnnotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfullyAnnotated

`func (o *VariantAnnotation) SetSuccessfullyAnnotated(v bool)`

SetSuccessfullyAnnotated sets SuccessfullyAnnotated field to given value.

### HasSuccessfullyAnnotated

`func (o *VariantAnnotation) HasSuccessfullyAnnotated() bool`

HasSuccessfullyAnnotated returns a boolean if a field has been set.

### GetTranscriptConsequences

`func (o *VariantAnnotation) GetTranscriptConsequences() []TranscriptConsequence`

GetTranscriptConsequences returns the TranscriptConsequences field if non-nil, zero value otherwise.

### GetTranscriptConsequencesOk

`func (o *VariantAnnotation) GetTranscriptConsequencesOk() (*[]TranscriptConsequence, bool)`

GetTranscriptConsequencesOk returns a tuple with the TranscriptConsequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptConsequences

`func (o *VariantAnnotation) SetTranscriptConsequences(v []TranscriptConsequence)`

SetTranscriptConsequences sets TranscriptConsequences field to given value.

### HasTranscriptConsequences

`func (o *VariantAnnotation) HasTranscriptConsequences() bool`

HasTranscriptConsequences returns a boolean if a field has been set.

### GetVariant

`func (o *VariantAnnotation) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *VariantAnnotation) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *VariantAnnotation) SetVariant(v string)`

SetVariant sets Variant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


