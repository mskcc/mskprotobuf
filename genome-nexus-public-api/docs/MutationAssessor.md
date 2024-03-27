# MutationAssessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodonStartPosition** | Pointer to **string** | Codon start position | [optional] 
**CosmicCount** | Pointer to **int32** | Number of mutations in COSMIC for this protein | [optional] 
**FunctionalImpact** | Pointer to **string** | Functional impact | [optional] 
**FunctionalImpactScore** | Pointer to **float64** | Functional impact score | [optional] 
**Hgvs** | Pointer to **string** |  | [optional] 
**HugoSymbol** | Pointer to **string** | Hugo gene symbol | [optional] 
**Input** | **string** | User-input variants | 
**MappingIssue** | Pointer to **string** | Mapping issue info | [optional] 
**MsaGaps** | Pointer to **float64** | Portion of gaps in variant position in multiple sequence alignment | [optional] 
**MsaHeight** | Pointer to **int32** | Number of diverse sequences in multiple sequence alignment (identical or highly similar sequences filtered out) | [optional] 
**MsaLink** | Pointer to **string** | Link to multiple sequence alignment | [optional] 
**PdbLink** | Pointer to **string** | Link to 3d structure browser | [optional] 
**ReferenceGenomeVariant** | Pointer to **string** | Reference genome variant | [optional] 
**ReferenceGenomeVariantType** | Pointer to **string** | Reference genome variant type | [optional] 
**RefseqId** | Pointer to **string** | Refseq protein ID | [optional] 
**RefseqPosition** | Pointer to **int32** | Variant position in Refseq protein, can be different from the one in Uniprot | [optional] 
**RefseqResidue** | Pointer to **string** | Reference residue in Refseq protein, can be different from the one in Uniprot | [optional] 
**SnpCount** | Pointer to **int32** | Number of SNPs in dbSNP for this protein | [optional] 
**UniprotId** | Pointer to **string** | Uniprot protein accession ID | [optional] 
**UniprotPosition** | Pointer to **int32** | Variant position in Uniprot protein, can be different from the one in Refseq | [optional] 
**UniprotResidue** | Pointer to **string** | Reference residue in Uniprot protein, can be different from the one in Refseq | [optional] 
**Variant** | Pointer to **string** | Amino acid substitution | [optional] 
**VariantConservationScore** | Pointer to **float64** | Variant conservation score | [optional] 
**VariantSpecificityScore** | Pointer to **float64** | Variant specificity score | [optional] 

## Methods

### NewMutationAssessor

`func NewMutationAssessor(input string, ) *MutationAssessor`

NewMutationAssessor instantiates a new MutationAssessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutationAssessorWithDefaults

`func NewMutationAssessorWithDefaults() *MutationAssessor`

NewMutationAssessorWithDefaults instantiates a new MutationAssessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodonStartPosition

`func (o *MutationAssessor) GetCodonStartPosition() string`

GetCodonStartPosition returns the CodonStartPosition field if non-nil, zero value otherwise.

### GetCodonStartPositionOk

`func (o *MutationAssessor) GetCodonStartPositionOk() (*string, bool)`

GetCodonStartPositionOk returns a tuple with the CodonStartPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodonStartPosition

`func (o *MutationAssessor) SetCodonStartPosition(v string)`

SetCodonStartPosition sets CodonStartPosition field to given value.

### HasCodonStartPosition

`func (o *MutationAssessor) HasCodonStartPosition() bool`

HasCodonStartPosition returns a boolean if a field has been set.

### GetCosmicCount

`func (o *MutationAssessor) GetCosmicCount() int32`

GetCosmicCount returns the CosmicCount field if non-nil, zero value otherwise.

### GetCosmicCountOk

`func (o *MutationAssessor) GetCosmicCountOk() (*int32, bool)`

GetCosmicCountOk returns a tuple with the CosmicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmicCount

`func (o *MutationAssessor) SetCosmicCount(v int32)`

SetCosmicCount sets CosmicCount field to given value.

### HasCosmicCount

`func (o *MutationAssessor) HasCosmicCount() bool`

HasCosmicCount returns a boolean if a field has been set.

### GetFunctionalImpact

`func (o *MutationAssessor) GetFunctionalImpact() string`

GetFunctionalImpact returns the FunctionalImpact field if non-nil, zero value otherwise.

### GetFunctionalImpactOk

`func (o *MutationAssessor) GetFunctionalImpactOk() (*string, bool)`

GetFunctionalImpactOk returns a tuple with the FunctionalImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalImpact

`func (o *MutationAssessor) SetFunctionalImpact(v string)`

SetFunctionalImpact sets FunctionalImpact field to given value.

### HasFunctionalImpact

`func (o *MutationAssessor) HasFunctionalImpact() bool`

HasFunctionalImpact returns a boolean if a field has been set.

### GetFunctionalImpactScore

`func (o *MutationAssessor) GetFunctionalImpactScore() float64`

GetFunctionalImpactScore returns the FunctionalImpactScore field if non-nil, zero value otherwise.

### GetFunctionalImpactScoreOk

`func (o *MutationAssessor) GetFunctionalImpactScoreOk() (*float64, bool)`

GetFunctionalImpactScoreOk returns a tuple with the FunctionalImpactScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalImpactScore

`func (o *MutationAssessor) SetFunctionalImpactScore(v float64)`

SetFunctionalImpactScore sets FunctionalImpactScore field to given value.

### HasFunctionalImpactScore

`func (o *MutationAssessor) HasFunctionalImpactScore() bool`

HasFunctionalImpactScore returns a boolean if a field has been set.

### GetHgvs

`func (o *MutationAssessor) GetHgvs() string`

GetHgvs returns the Hgvs field if non-nil, zero value otherwise.

### GetHgvsOk

`func (o *MutationAssessor) GetHgvsOk() (*string, bool)`

GetHgvsOk returns a tuple with the Hgvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvs

`func (o *MutationAssessor) SetHgvs(v string)`

SetHgvs sets Hgvs field to given value.

### HasHgvs

`func (o *MutationAssessor) HasHgvs() bool`

HasHgvs returns a boolean if a field has been set.

### GetHugoSymbol

`func (o *MutationAssessor) GetHugoSymbol() string`

GetHugoSymbol returns the HugoSymbol field if non-nil, zero value otherwise.

### GetHugoSymbolOk

`func (o *MutationAssessor) GetHugoSymbolOk() (*string, bool)`

GetHugoSymbolOk returns a tuple with the HugoSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugoSymbol

`func (o *MutationAssessor) SetHugoSymbol(v string)`

SetHugoSymbol sets HugoSymbol field to given value.

### HasHugoSymbol

`func (o *MutationAssessor) HasHugoSymbol() bool`

HasHugoSymbol returns a boolean if a field has been set.

### GetInput

`func (o *MutationAssessor) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *MutationAssessor) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *MutationAssessor) SetInput(v string)`

SetInput sets Input field to given value.


### GetMappingIssue

`func (o *MutationAssessor) GetMappingIssue() string`

GetMappingIssue returns the MappingIssue field if non-nil, zero value otherwise.

### GetMappingIssueOk

`func (o *MutationAssessor) GetMappingIssueOk() (*string, bool)`

GetMappingIssueOk returns a tuple with the MappingIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingIssue

`func (o *MutationAssessor) SetMappingIssue(v string)`

SetMappingIssue sets MappingIssue field to given value.

### HasMappingIssue

`func (o *MutationAssessor) HasMappingIssue() bool`

HasMappingIssue returns a boolean if a field has been set.

### GetMsaGaps

`func (o *MutationAssessor) GetMsaGaps() float64`

GetMsaGaps returns the MsaGaps field if non-nil, zero value otherwise.

### GetMsaGapsOk

`func (o *MutationAssessor) GetMsaGapsOk() (*float64, bool)`

GetMsaGapsOk returns a tuple with the MsaGaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsaGaps

`func (o *MutationAssessor) SetMsaGaps(v float64)`

SetMsaGaps sets MsaGaps field to given value.

### HasMsaGaps

`func (o *MutationAssessor) HasMsaGaps() bool`

HasMsaGaps returns a boolean if a field has been set.

### GetMsaHeight

`func (o *MutationAssessor) GetMsaHeight() int32`

GetMsaHeight returns the MsaHeight field if non-nil, zero value otherwise.

### GetMsaHeightOk

`func (o *MutationAssessor) GetMsaHeightOk() (*int32, bool)`

GetMsaHeightOk returns a tuple with the MsaHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsaHeight

`func (o *MutationAssessor) SetMsaHeight(v int32)`

SetMsaHeight sets MsaHeight field to given value.

### HasMsaHeight

`func (o *MutationAssessor) HasMsaHeight() bool`

HasMsaHeight returns a boolean if a field has been set.

### GetMsaLink

`func (o *MutationAssessor) GetMsaLink() string`

GetMsaLink returns the MsaLink field if non-nil, zero value otherwise.

### GetMsaLinkOk

`func (o *MutationAssessor) GetMsaLinkOk() (*string, bool)`

GetMsaLinkOk returns a tuple with the MsaLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsaLink

`func (o *MutationAssessor) SetMsaLink(v string)`

SetMsaLink sets MsaLink field to given value.

### HasMsaLink

`func (o *MutationAssessor) HasMsaLink() bool`

HasMsaLink returns a boolean if a field has been set.

### GetPdbLink

`func (o *MutationAssessor) GetPdbLink() string`

GetPdbLink returns the PdbLink field if non-nil, zero value otherwise.

### GetPdbLinkOk

`func (o *MutationAssessor) GetPdbLinkOk() (*string, bool)`

GetPdbLinkOk returns a tuple with the PdbLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbLink

`func (o *MutationAssessor) SetPdbLink(v string)`

SetPdbLink sets PdbLink field to given value.

### HasPdbLink

`func (o *MutationAssessor) HasPdbLink() bool`

HasPdbLink returns a boolean if a field has been set.

### GetReferenceGenomeVariant

`func (o *MutationAssessor) GetReferenceGenomeVariant() string`

GetReferenceGenomeVariant returns the ReferenceGenomeVariant field if non-nil, zero value otherwise.

### GetReferenceGenomeVariantOk

`func (o *MutationAssessor) GetReferenceGenomeVariantOk() (*string, bool)`

GetReferenceGenomeVariantOk returns a tuple with the ReferenceGenomeVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceGenomeVariant

`func (o *MutationAssessor) SetReferenceGenomeVariant(v string)`

SetReferenceGenomeVariant sets ReferenceGenomeVariant field to given value.

### HasReferenceGenomeVariant

`func (o *MutationAssessor) HasReferenceGenomeVariant() bool`

HasReferenceGenomeVariant returns a boolean if a field has been set.

### GetReferenceGenomeVariantType

`func (o *MutationAssessor) GetReferenceGenomeVariantType() string`

GetReferenceGenomeVariantType returns the ReferenceGenomeVariantType field if non-nil, zero value otherwise.

### GetReferenceGenomeVariantTypeOk

`func (o *MutationAssessor) GetReferenceGenomeVariantTypeOk() (*string, bool)`

GetReferenceGenomeVariantTypeOk returns a tuple with the ReferenceGenomeVariantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceGenomeVariantType

`func (o *MutationAssessor) SetReferenceGenomeVariantType(v string)`

SetReferenceGenomeVariantType sets ReferenceGenomeVariantType field to given value.

### HasReferenceGenomeVariantType

`func (o *MutationAssessor) HasReferenceGenomeVariantType() bool`

HasReferenceGenomeVariantType returns a boolean if a field has been set.

### GetRefseqId

`func (o *MutationAssessor) GetRefseqId() string`

GetRefseqId returns the RefseqId field if non-nil, zero value otherwise.

### GetRefseqIdOk

`func (o *MutationAssessor) GetRefseqIdOk() (*string, bool)`

GetRefseqIdOk returns a tuple with the RefseqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqId

`func (o *MutationAssessor) SetRefseqId(v string)`

SetRefseqId sets RefseqId field to given value.

### HasRefseqId

`func (o *MutationAssessor) HasRefseqId() bool`

HasRefseqId returns a boolean if a field has been set.

### GetRefseqPosition

`func (o *MutationAssessor) GetRefseqPosition() int32`

GetRefseqPosition returns the RefseqPosition field if non-nil, zero value otherwise.

### GetRefseqPositionOk

`func (o *MutationAssessor) GetRefseqPositionOk() (*int32, bool)`

GetRefseqPositionOk returns a tuple with the RefseqPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqPosition

`func (o *MutationAssessor) SetRefseqPosition(v int32)`

SetRefseqPosition sets RefseqPosition field to given value.

### HasRefseqPosition

`func (o *MutationAssessor) HasRefseqPosition() bool`

HasRefseqPosition returns a boolean if a field has been set.

### GetRefseqResidue

`func (o *MutationAssessor) GetRefseqResidue() string`

GetRefseqResidue returns the RefseqResidue field if non-nil, zero value otherwise.

### GetRefseqResidueOk

`func (o *MutationAssessor) GetRefseqResidueOk() (*string, bool)`

GetRefseqResidueOk returns a tuple with the RefseqResidue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqResidue

`func (o *MutationAssessor) SetRefseqResidue(v string)`

SetRefseqResidue sets RefseqResidue field to given value.

### HasRefseqResidue

`func (o *MutationAssessor) HasRefseqResidue() bool`

HasRefseqResidue returns a boolean if a field has been set.

### GetSnpCount

`func (o *MutationAssessor) GetSnpCount() int32`

GetSnpCount returns the SnpCount field if non-nil, zero value otherwise.

### GetSnpCountOk

`func (o *MutationAssessor) GetSnpCountOk() (*int32, bool)`

GetSnpCountOk returns a tuple with the SnpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnpCount

`func (o *MutationAssessor) SetSnpCount(v int32)`

SetSnpCount sets SnpCount field to given value.

### HasSnpCount

`func (o *MutationAssessor) HasSnpCount() bool`

HasSnpCount returns a boolean if a field has been set.

### GetUniprotId

`func (o *MutationAssessor) GetUniprotId() string`

GetUniprotId returns the UniprotId field if non-nil, zero value otherwise.

### GetUniprotIdOk

`func (o *MutationAssessor) GetUniprotIdOk() (*string, bool)`

GetUniprotIdOk returns a tuple with the UniprotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotId

`func (o *MutationAssessor) SetUniprotId(v string)`

SetUniprotId sets UniprotId field to given value.

### HasUniprotId

`func (o *MutationAssessor) HasUniprotId() bool`

HasUniprotId returns a boolean if a field has been set.

### GetUniprotPosition

`func (o *MutationAssessor) GetUniprotPosition() int32`

GetUniprotPosition returns the UniprotPosition field if non-nil, zero value otherwise.

### GetUniprotPositionOk

`func (o *MutationAssessor) GetUniprotPositionOk() (*int32, bool)`

GetUniprotPositionOk returns a tuple with the UniprotPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotPosition

`func (o *MutationAssessor) SetUniprotPosition(v int32)`

SetUniprotPosition sets UniprotPosition field to given value.

### HasUniprotPosition

`func (o *MutationAssessor) HasUniprotPosition() bool`

HasUniprotPosition returns a boolean if a field has been set.

### GetUniprotResidue

`func (o *MutationAssessor) GetUniprotResidue() string`

GetUniprotResidue returns the UniprotResidue field if non-nil, zero value otherwise.

### GetUniprotResidueOk

`func (o *MutationAssessor) GetUniprotResidueOk() (*string, bool)`

GetUniprotResidueOk returns a tuple with the UniprotResidue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniprotResidue

`func (o *MutationAssessor) SetUniprotResidue(v string)`

SetUniprotResidue sets UniprotResidue field to given value.

### HasUniprotResidue

`func (o *MutationAssessor) HasUniprotResidue() bool`

HasUniprotResidue returns a boolean if a field has been set.

### GetVariant

`func (o *MutationAssessor) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *MutationAssessor) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *MutationAssessor) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *MutationAssessor) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetVariantConservationScore

`func (o *MutationAssessor) GetVariantConservationScore() float64`

GetVariantConservationScore returns the VariantConservationScore field if non-nil, zero value otherwise.

### GetVariantConservationScoreOk

`func (o *MutationAssessor) GetVariantConservationScoreOk() (*float64, bool)`

GetVariantConservationScoreOk returns a tuple with the VariantConservationScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantConservationScore

`func (o *MutationAssessor) SetVariantConservationScore(v float64)`

SetVariantConservationScore sets VariantConservationScore field to given value.

### HasVariantConservationScore

`func (o *MutationAssessor) HasVariantConservationScore() bool`

HasVariantConservationScore returns a boolean if a field has been set.

### GetVariantSpecificityScore

`func (o *MutationAssessor) GetVariantSpecificityScore() float64`

GetVariantSpecificityScore returns the VariantSpecificityScore field if non-nil, zero value otherwise.

### GetVariantSpecificityScoreOk

`func (o *MutationAssessor) GetVariantSpecificityScoreOk() (*float64, bool)`

GetVariantSpecificityScoreOk returns a tuple with the VariantSpecificityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantSpecificityScore

`func (o *MutationAssessor) SetVariantSpecificityScore(v float64)`

SetVariantSpecificityScore sets VariantSpecificityScore field to given value.

### HasVariantSpecificityScore

`func (o *MutationAssessor) HasVariantSpecificityScore() bool`

HasVariantSpecificityScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


