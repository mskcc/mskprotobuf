# MyVariantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClinVar** | Pointer to [**ClinVar**](ClinVar.md) |  | [optional] 
**Cosmic** | Pointer to [**Cosmic**](Cosmic.md) |  | [optional] 
**Dbsnp** | Pointer to [**Dbsnp**](Dbsnp.md) |  | [optional] 
**GnomadExome** | Pointer to [**Gnomad**](Gnomad.md) |  | [optional] 
**GnomadGenome** | Pointer to [**Gnomad**](Gnomad.md) |  | [optional] 
**Hgvs** | Pointer to **string** | hgvs | [optional] 
**Mutdb** | Pointer to [**Mutdb**](Mutdb.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Snpeff** | Pointer to [**Snpeff**](Snpeff.md) |  | [optional] 
**Variant** | Pointer to **string** | variant | [optional] 
**Vcf** | Pointer to [**Vcf**](Vcf.md) |  | [optional] 
**Version** | Pointer to **int32** | version | [optional] 

## Methods

### NewMyVariantInfo

`func NewMyVariantInfo() *MyVariantInfo`

NewMyVariantInfo instantiates a new MyVariantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyVariantInfoWithDefaults

`func NewMyVariantInfoWithDefaults() *MyVariantInfo`

NewMyVariantInfoWithDefaults instantiates a new MyVariantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClinVar

`func (o *MyVariantInfo) GetClinVar() ClinVar`

GetClinVar returns the ClinVar field if non-nil, zero value otherwise.

### GetClinVarOk

`func (o *MyVariantInfo) GetClinVarOk() (*ClinVar, bool)`

GetClinVarOk returns a tuple with the ClinVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinVar

`func (o *MyVariantInfo) SetClinVar(v ClinVar)`

SetClinVar sets ClinVar field to given value.

### HasClinVar

`func (o *MyVariantInfo) HasClinVar() bool`

HasClinVar returns a boolean if a field has been set.

### GetCosmic

`func (o *MyVariantInfo) GetCosmic() Cosmic`

GetCosmic returns the Cosmic field if non-nil, zero value otherwise.

### GetCosmicOk

`func (o *MyVariantInfo) GetCosmicOk() (*Cosmic, bool)`

GetCosmicOk returns a tuple with the Cosmic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmic

`func (o *MyVariantInfo) SetCosmic(v Cosmic)`

SetCosmic sets Cosmic field to given value.

### HasCosmic

`func (o *MyVariantInfo) HasCosmic() bool`

HasCosmic returns a boolean if a field has been set.

### GetDbsnp

`func (o *MyVariantInfo) GetDbsnp() Dbsnp`

GetDbsnp returns the Dbsnp field if non-nil, zero value otherwise.

### GetDbsnpOk

`func (o *MyVariantInfo) GetDbsnpOk() (*Dbsnp, bool)`

GetDbsnpOk returns a tuple with the Dbsnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbsnp

`func (o *MyVariantInfo) SetDbsnp(v Dbsnp)`

SetDbsnp sets Dbsnp field to given value.

### HasDbsnp

`func (o *MyVariantInfo) HasDbsnp() bool`

HasDbsnp returns a boolean if a field has been set.

### GetGnomadExome

`func (o *MyVariantInfo) GetGnomadExome() Gnomad`

GetGnomadExome returns the GnomadExome field if non-nil, zero value otherwise.

### GetGnomadExomeOk

`func (o *MyVariantInfo) GetGnomadExomeOk() (*Gnomad, bool)`

GetGnomadExomeOk returns a tuple with the GnomadExome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnomadExome

`func (o *MyVariantInfo) SetGnomadExome(v Gnomad)`

SetGnomadExome sets GnomadExome field to given value.

### HasGnomadExome

`func (o *MyVariantInfo) HasGnomadExome() bool`

HasGnomadExome returns a boolean if a field has been set.

### GetGnomadGenome

`func (o *MyVariantInfo) GetGnomadGenome() Gnomad`

GetGnomadGenome returns the GnomadGenome field if non-nil, zero value otherwise.

### GetGnomadGenomeOk

`func (o *MyVariantInfo) GetGnomadGenomeOk() (*Gnomad, bool)`

GetGnomadGenomeOk returns a tuple with the GnomadGenome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnomadGenome

`func (o *MyVariantInfo) SetGnomadGenome(v Gnomad)`

SetGnomadGenome sets GnomadGenome field to given value.

### HasGnomadGenome

`func (o *MyVariantInfo) HasGnomadGenome() bool`

HasGnomadGenome returns a boolean if a field has been set.

### GetHgvs

`func (o *MyVariantInfo) GetHgvs() string`

GetHgvs returns the Hgvs field if non-nil, zero value otherwise.

### GetHgvsOk

`func (o *MyVariantInfo) GetHgvsOk() (*string, bool)`

GetHgvsOk returns a tuple with the Hgvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvs

`func (o *MyVariantInfo) SetHgvs(v string)`

SetHgvs sets Hgvs field to given value.

### HasHgvs

`func (o *MyVariantInfo) HasHgvs() bool`

HasHgvs returns a boolean if a field has been set.

### GetMutdb

`func (o *MyVariantInfo) GetMutdb() Mutdb`

GetMutdb returns the Mutdb field if non-nil, zero value otherwise.

### GetMutdbOk

`func (o *MyVariantInfo) GetMutdbOk() (*Mutdb, bool)`

GetMutdbOk returns a tuple with the Mutdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutdb

`func (o *MyVariantInfo) SetMutdb(v Mutdb)`

SetMutdb sets Mutdb field to given value.

### HasMutdb

`func (o *MyVariantInfo) HasMutdb() bool`

HasMutdb returns a boolean if a field has been set.

### GetQuery

`func (o *MyVariantInfo) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MyVariantInfo) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MyVariantInfo) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MyVariantInfo) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSnpeff

`func (o *MyVariantInfo) GetSnpeff() Snpeff`

GetSnpeff returns the Snpeff field if non-nil, zero value otherwise.

### GetSnpeffOk

`func (o *MyVariantInfo) GetSnpeffOk() (*Snpeff, bool)`

GetSnpeffOk returns a tuple with the Snpeff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnpeff

`func (o *MyVariantInfo) SetSnpeff(v Snpeff)`

SetSnpeff sets Snpeff field to given value.

### HasSnpeff

`func (o *MyVariantInfo) HasSnpeff() bool`

HasSnpeff returns a boolean if a field has been set.

### GetVariant

`func (o *MyVariantInfo) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *MyVariantInfo) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *MyVariantInfo) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *MyVariantInfo) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetVcf

`func (o *MyVariantInfo) GetVcf() Vcf`

GetVcf returns the Vcf field if non-nil, zero value otherwise.

### GetVcfOk

`func (o *MyVariantInfo) GetVcfOk() (*Vcf, bool)`

GetVcfOk returns a tuple with the Vcf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcf

`func (o *MyVariantInfo) SetVcf(v Vcf)`

SetVcf sets Vcf field to given value.

### HasVcf

`func (o *MyVariantInfo) HasVcf() bool`

HasVcf returns a boolean if a field has been set.

### GetVersion

`func (o *MyVariantInfo) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MyVariantInfo) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MyVariantInfo) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MyVariantInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


