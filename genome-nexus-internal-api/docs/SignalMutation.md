# SignalMutation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiallelicCountsByTumorType** | Pointer to [**[]CountByTumorType**](CountByTumorType.md) | Biallelic Counts by Tumor Type | [optional] 
**Chromosome** | Pointer to **string** | Chromosome | [optional] 
**CountsByTumorType** | Pointer to [**[]CountByTumorType**](CountByTumorType.md) | Counts by Tumor Type | [optional] 
**EndPosition** | Pointer to **int64** | End Position | [optional] 
**GeneralPopulationStats** | Pointer to [**GeneralPopulationStats**](GeneralPopulationStats.md) |  | [optional] 
**HugoGeneSymbol** | Pointer to **string** | Hugo Gene Symbol | [optional] 
**MskExperReview** | Pointer to **bool** | Msk Expert Review | [optional] 
**MutationStatus** | Pointer to **string** | Mutation Status | [optional] 
**OverallNumberOfGermlineHomozygous** | Pointer to **int32** |  | [optional] 
**Pathogenic** | Pointer to **string** | Pathogenic | [optional] 
**Penetrance** | Pointer to **string** | Penetrance | [optional] 
**QcPassCountsByTumorType** | Pointer to [**[]CountByTumorType**](CountByTumorType.md) | QC Pass Counts by Tumor Type | [optional] 
**ReferenceAllele** | Pointer to **string** | Reference Allele | [optional] 
**StartPosition** | Pointer to **int64** | Start Position | [optional] 
**StatsByTumorType** | Pointer to [**[]StatsByTumorType**](StatsByTumorType.md) | Stats By Tumor Type | [optional] 
**VariantAllele** | Pointer to **string** | Variant Allele | [optional] 

## Methods

### NewSignalMutation

`func NewSignalMutation() *SignalMutation`

NewSignalMutation instantiates a new SignalMutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalMutationWithDefaults

`func NewSignalMutationWithDefaults() *SignalMutation`

NewSignalMutationWithDefaults instantiates a new SignalMutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiallelicCountsByTumorType

`func (o *SignalMutation) GetBiallelicCountsByTumorType() []CountByTumorType`

GetBiallelicCountsByTumorType returns the BiallelicCountsByTumorType field if non-nil, zero value otherwise.

### GetBiallelicCountsByTumorTypeOk

`func (o *SignalMutation) GetBiallelicCountsByTumorTypeOk() (*[]CountByTumorType, bool)`

GetBiallelicCountsByTumorTypeOk returns a tuple with the BiallelicCountsByTumorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiallelicCountsByTumorType

`func (o *SignalMutation) SetBiallelicCountsByTumorType(v []CountByTumorType)`

SetBiallelicCountsByTumorType sets BiallelicCountsByTumorType field to given value.

### HasBiallelicCountsByTumorType

`func (o *SignalMutation) HasBiallelicCountsByTumorType() bool`

HasBiallelicCountsByTumorType returns a boolean if a field has been set.

### GetChromosome

`func (o *SignalMutation) GetChromosome() string`

GetChromosome returns the Chromosome field if non-nil, zero value otherwise.

### GetChromosomeOk

`func (o *SignalMutation) GetChromosomeOk() (*string, bool)`

GetChromosomeOk returns a tuple with the Chromosome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosome

`func (o *SignalMutation) SetChromosome(v string)`

SetChromosome sets Chromosome field to given value.

### HasChromosome

`func (o *SignalMutation) HasChromosome() bool`

HasChromosome returns a boolean if a field has been set.

### GetCountsByTumorType

`func (o *SignalMutation) GetCountsByTumorType() []CountByTumorType`

GetCountsByTumorType returns the CountsByTumorType field if non-nil, zero value otherwise.

### GetCountsByTumorTypeOk

`func (o *SignalMutation) GetCountsByTumorTypeOk() (*[]CountByTumorType, bool)`

GetCountsByTumorTypeOk returns a tuple with the CountsByTumorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsByTumorType

`func (o *SignalMutation) SetCountsByTumorType(v []CountByTumorType)`

SetCountsByTumorType sets CountsByTumorType field to given value.

### HasCountsByTumorType

`func (o *SignalMutation) HasCountsByTumorType() bool`

HasCountsByTumorType returns a boolean if a field has been set.

### GetEndPosition

`func (o *SignalMutation) GetEndPosition() int64`

GetEndPosition returns the EndPosition field if non-nil, zero value otherwise.

### GetEndPositionOk

`func (o *SignalMutation) GetEndPositionOk() (*int64, bool)`

GetEndPositionOk returns a tuple with the EndPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPosition

`func (o *SignalMutation) SetEndPosition(v int64)`

SetEndPosition sets EndPosition field to given value.

### HasEndPosition

`func (o *SignalMutation) HasEndPosition() bool`

HasEndPosition returns a boolean if a field has been set.

### GetGeneralPopulationStats

`func (o *SignalMutation) GetGeneralPopulationStats() GeneralPopulationStats`

GetGeneralPopulationStats returns the GeneralPopulationStats field if non-nil, zero value otherwise.

### GetGeneralPopulationStatsOk

`func (o *SignalMutation) GetGeneralPopulationStatsOk() (*GeneralPopulationStats, bool)`

GetGeneralPopulationStatsOk returns a tuple with the GeneralPopulationStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralPopulationStats

`func (o *SignalMutation) SetGeneralPopulationStats(v GeneralPopulationStats)`

SetGeneralPopulationStats sets GeneralPopulationStats field to given value.

### HasGeneralPopulationStats

`func (o *SignalMutation) HasGeneralPopulationStats() bool`

HasGeneralPopulationStats returns a boolean if a field has been set.

### GetHugoGeneSymbol

`func (o *SignalMutation) GetHugoGeneSymbol() string`

GetHugoGeneSymbol returns the HugoGeneSymbol field if non-nil, zero value otherwise.

### GetHugoGeneSymbolOk

`func (o *SignalMutation) GetHugoGeneSymbolOk() (*string, bool)`

GetHugoGeneSymbolOk returns a tuple with the HugoGeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugoGeneSymbol

`func (o *SignalMutation) SetHugoGeneSymbol(v string)`

SetHugoGeneSymbol sets HugoGeneSymbol field to given value.

### HasHugoGeneSymbol

`func (o *SignalMutation) HasHugoGeneSymbol() bool`

HasHugoGeneSymbol returns a boolean if a field has been set.

### GetMskExperReview

`func (o *SignalMutation) GetMskExperReview() bool`

GetMskExperReview returns the MskExperReview field if non-nil, zero value otherwise.

### GetMskExperReviewOk

`func (o *SignalMutation) GetMskExperReviewOk() (*bool, bool)`

GetMskExperReviewOk returns a tuple with the MskExperReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMskExperReview

`func (o *SignalMutation) SetMskExperReview(v bool)`

SetMskExperReview sets MskExperReview field to given value.

### HasMskExperReview

`func (o *SignalMutation) HasMskExperReview() bool`

HasMskExperReview returns a boolean if a field has been set.

### GetMutationStatus

`func (o *SignalMutation) GetMutationStatus() string`

GetMutationStatus returns the MutationStatus field if non-nil, zero value otherwise.

### GetMutationStatusOk

`func (o *SignalMutation) GetMutationStatusOk() (*string, bool)`

GetMutationStatusOk returns a tuple with the MutationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutationStatus

`func (o *SignalMutation) SetMutationStatus(v string)`

SetMutationStatus sets MutationStatus field to given value.

### HasMutationStatus

`func (o *SignalMutation) HasMutationStatus() bool`

HasMutationStatus returns a boolean if a field has been set.

### GetOverallNumberOfGermlineHomozygous

`func (o *SignalMutation) GetOverallNumberOfGermlineHomozygous() int32`

GetOverallNumberOfGermlineHomozygous returns the OverallNumberOfGermlineHomozygous field if non-nil, zero value otherwise.

### GetOverallNumberOfGermlineHomozygousOk

`func (o *SignalMutation) GetOverallNumberOfGermlineHomozygousOk() (*int32, bool)`

GetOverallNumberOfGermlineHomozygousOk returns a tuple with the OverallNumberOfGermlineHomozygous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallNumberOfGermlineHomozygous

`func (o *SignalMutation) SetOverallNumberOfGermlineHomozygous(v int32)`

SetOverallNumberOfGermlineHomozygous sets OverallNumberOfGermlineHomozygous field to given value.

### HasOverallNumberOfGermlineHomozygous

`func (o *SignalMutation) HasOverallNumberOfGermlineHomozygous() bool`

HasOverallNumberOfGermlineHomozygous returns a boolean if a field has been set.

### GetPathogenic

`func (o *SignalMutation) GetPathogenic() string`

GetPathogenic returns the Pathogenic field if non-nil, zero value otherwise.

### GetPathogenicOk

`func (o *SignalMutation) GetPathogenicOk() (*string, bool)`

GetPathogenicOk returns a tuple with the Pathogenic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathogenic

`func (o *SignalMutation) SetPathogenic(v string)`

SetPathogenic sets Pathogenic field to given value.

### HasPathogenic

`func (o *SignalMutation) HasPathogenic() bool`

HasPathogenic returns a boolean if a field has been set.

### GetPenetrance

`func (o *SignalMutation) GetPenetrance() string`

GetPenetrance returns the Penetrance field if non-nil, zero value otherwise.

### GetPenetranceOk

`func (o *SignalMutation) GetPenetranceOk() (*string, bool)`

GetPenetranceOk returns a tuple with the Penetrance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenetrance

`func (o *SignalMutation) SetPenetrance(v string)`

SetPenetrance sets Penetrance field to given value.

### HasPenetrance

`func (o *SignalMutation) HasPenetrance() bool`

HasPenetrance returns a boolean if a field has been set.

### GetQcPassCountsByTumorType

`func (o *SignalMutation) GetQcPassCountsByTumorType() []CountByTumorType`

GetQcPassCountsByTumorType returns the QcPassCountsByTumorType field if non-nil, zero value otherwise.

### GetQcPassCountsByTumorTypeOk

`func (o *SignalMutation) GetQcPassCountsByTumorTypeOk() (*[]CountByTumorType, bool)`

GetQcPassCountsByTumorTypeOk returns a tuple with the QcPassCountsByTumorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQcPassCountsByTumorType

`func (o *SignalMutation) SetQcPassCountsByTumorType(v []CountByTumorType)`

SetQcPassCountsByTumorType sets QcPassCountsByTumorType field to given value.

### HasQcPassCountsByTumorType

`func (o *SignalMutation) HasQcPassCountsByTumorType() bool`

HasQcPassCountsByTumorType returns a boolean if a field has been set.

### GetReferenceAllele

`func (o *SignalMutation) GetReferenceAllele() string`

GetReferenceAllele returns the ReferenceAllele field if non-nil, zero value otherwise.

### GetReferenceAlleleOk

`func (o *SignalMutation) GetReferenceAlleleOk() (*string, bool)`

GetReferenceAlleleOk returns a tuple with the ReferenceAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAllele

`func (o *SignalMutation) SetReferenceAllele(v string)`

SetReferenceAllele sets ReferenceAllele field to given value.

### HasReferenceAllele

`func (o *SignalMutation) HasReferenceAllele() bool`

HasReferenceAllele returns a boolean if a field has been set.

### GetStartPosition

`func (o *SignalMutation) GetStartPosition() int64`

GetStartPosition returns the StartPosition field if non-nil, zero value otherwise.

### GetStartPositionOk

`func (o *SignalMutation) GetStartPositionOk() (*int64, bool)`

GetStartPositionOk returns a tuple with the StartPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPosition

`func (o *SignalMutation) SetStartPosition(v int64)`

SetStartPosition sets StartPosition field to given value.

### HasStartPosition

`func (o *SignalMutation) HasStartPosition() bool`

HasStartPosition returns a boolean if a field has been set.

### GetStatsByTumorType

`func (o *SignalMutation) GetStatsByTumorType() []StatsByTumorType`

GetStatsByTumorType returns the StatsByTumorType field if non-nil, zero value otherwise.

### GetStatsByTumorTypeOk

`func (o *SignalMutation) GetStatsByTumorTypeOk() (*[]StatsByTumorType, bool)`

GetStatsByTumorTypeOk returns a tuple with the StatsByTumorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsByTumorType

`func (o *SignalMutation) SetStatsByTumorType(v []StatsByTumorType)`

SetStatsByTumorType sets StatsByTumorType field to given value.

### HasStatsByTumorType

`func (o *SignalMutation) HasStatsByTumorType() bool`

HasStatsByTumorType returns a boolean if a field has been set.

### GetVariantAllele

`func (o *SignalMutation) GetVariantAllele() string`

GetVariantAllele returns the VariantAllele field if non-nil, zero value otherwise.

### GetVariantAlleleOk

`func (o *SignalMutation) GetVariantAlleleOk() (*string, bool)`

GetVariantAlleleOk returns a tuple with the VariantAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantAllele

`func (o *SignalMutation) SetVariantAllele(v string)`

SetVariantAllele sets VariantAllele field to given value.

### HasVariantAllele

`func (o *SignalMutation) HasVariantAllele() bool`

HasVariantAllele returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


