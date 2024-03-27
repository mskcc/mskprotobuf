# StatsByTumorType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeAtDx** | Pointer to **int32** | Median Age at Dx | [optional] 
**FBiallelic** | Pointer to **float64** | Frequency Of Biallelic | [optional] 
**FCancerTypeCount** | Pointer to **float64** | Frequency Of Cancer Type Count | [optional] 
**HrdScore** | Pointer to [**HrdScore**](HrdScore.md) |  | [optional] 
**MsiScore** | Pointer to **float64** | Msi Score | [optional] 
**NCancerTypeCount** | Pointer to **int32** | Number Of Cancer Type Count | [optional] 
**NumberOfGermlineHomozygous** | Pointer to **int32** | Number Of Germline Homozygous Per Tumor Type | [optional] 
**NumberWithSig** | Pointer to **int32** | Number of Variants with Signature | [optional] 
**Tmb** | Pointer to **float64** | Median TMB | [optional] 
**TumorType** | Pointer to **string** | Tumor Type | [optional] 

## Methods

### NewStatsByTumorType

`func NewStatsByTumorType() *StatsByTumorType`

NewStatsByTumorType instantiates a new StatsByTumorType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsByTumorTypeWithDefaults

`func NewStatsByTumorTypeWithDefaults() *StatsByTumorType`

NewStatsByTumorTypeWithDefaults instantiates a new StatsByTumorType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgeAtDx

`func (o *StatsByTumorType) GetAgeAtDx() int32`

GetAgeAtDx returns the AgeAtDx field if non-nil, zero value otherwise.

### GetAgeAtDxOk

`func (o *StatsByTumorType) GetAgeAtDxOk() (*int32, bool)`

GetAgeAtDxOk returns a tuple with the AgeAtDx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeAtDx

`func (o *StatsByTumorType) SetAgeAtDx(v int32)`

SetAgeAtDx sets AgeAtDx field to given value.

### HasAgeAtDx

`func (o *StatsByTumorType) HasAgeAtDx() bool`

HasAgeAtDx returns a boolean if a field has been set.

### GetFBiallelic

`func (o *StatsByTumorType) GetFBiallelic() float64`

GetFBiallelic returns the FBiallelic field if non-nil, zero value otherwise.

### GetFBiallelicOk

`func (o *StatsByTumorType) GetFBiallelicOk() (*float64, bool)`

GetFBiallelicOk returns a tuple with the FBiallelic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFBiallelic

`func (o *StatsByTumorType) SetFBiallelic(v float64)`

SetFBiallelic sets FBiallelic field to given value.

### HasFBiallelic

`func (o *StatsByTumorType) HasFBiallelic() bool`

HasFBiallelic returns a boolean if a field has been set.

### GetFCancerTypeCount

`func (o *StatsByTumorType) GetFCancerTypeCount() float64`

GetFCancerTypeCount returns the FCancerTypeCount field if non-nil, zero value otherwise.

### GetFCancerTypeCountOk

`func (o *StatsByTumorType) GetFCancerTypeCountOk() (*float64, bool)`

GetFCancerTypeCountOk returns a tuple with the FCancerTypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFCancerTypeCount

`func (o *StatsByTumorType) SetFCancerTypeCount(v float64)`

SetFCancerTypeCount sets FCancerTypeCount field to given value.

### HasFCancerTypeCount

`func (o *StatsByTumorType) HasFCancerTypeCount() bool`

HasFCancerTypeCount returns a boolean if a field has been set.

### GetHrdScore

`func (o *StatsByTumorType) GetHrdScore() HrdScore`

GetHrdScore returns the HrdScore field if non-nil, zero value otherwise.

### GetHrdScoreOk

`func (o *StatsByTumorType) GetHrdScoreOk() (*HrdScore, bool)`

GetHrdScoreOk returns a tuple with the HrdScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHrdScore

`func (o *StatsByTumorType) SetHrdScore(v HrdScore)`

SetHrdScore sets HrdScore field to given value.

### HasHrdScore

`func (o *StatsByTumorType) HasHrdScore() bool`

HasHrdScore returns a boolean if a field has been set.

### GetMsiScore

`func (o *StatsByTumorType) GetMsiScore() float64`

GetMsiScore returns the MsiScore field if non-nil, zero value otherwise.

### GetMsiScoreOk

`func (o *StatsByTumorType) GetMsiScoreOk() (*float64, bool)`

GetMsiScoreOk returns a tuple with the MsiScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsiScore

`func (o *StatsByTumorType) SetMsiScore(v float64)`

SetMsiScore sets MsiScore field to given value.

### HasMsiScore

`func (o *StatsByTumorType) HasMsiScore() bool`

HasMsiScore returns a boolean if a field has been set.

### GetNCancerTypeCount

`func (o *StatsByTumorType) GetNCancerTypeCount() int32`

GetNCancerTypeCount returns the NCancerTypeCount field if non-nil, zero value otherwise.

### GetNCancerTypeCountOk

`func (o *StatsByTumorType) GetNCancerTypeCountOk() (*int32, bool)`

GetNCancerTypeCountOk returns a tuple with the NCancerTypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCancerTypeCount

`func (o *StatsByTumorType) SetNCancerTypeCount(v int32)`

SetNCancerTypeCount sets NCancerTypeCount field to given value.

### HasNCancerTypeCount

`func (o *StatsByTumorType) HasNCancerTypeCount() bool`

HasNCancerTypeCount returns a boolean if a field has been set.

### GetNumberOfGermlineHomozygous

`func (o *StatsByTumorType) GetNumberOfGermlineHomozygous() int32`

GetNumberOfGermlineHomozygous returns the NumberOfGermlineHomozygous field if non-nil, zero value otherwise.

### GetNumberOfGermlineHomozygousOk

`func (o *StatsByTumorType) GetNumberOfGermlineHomozygousOk() (*int32, bool)`

GetNumberOfGermlineHomozygousOk returns a tuple with the NumberOfGermlineHomozygous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfGermlineHomozygous

`func (o *StatsByTumorType) SetNumberOfGermlineHomozygous(v int32)`

SetNumberOfGermlineHomozygous sets NumberOfGermlineHomozygous field to given value.

### HasNumberOfGermlineHomozygous

`func (o *StatsByTumorType) HasNumberOfGermlineHomozygous() bool`

HasNumberOfGermlineHomozygous returns a boolean if a field has been set.

### GetNumberWithSig

`func (o *StatsByTumorType) GetNumberWithSig() int32`

GetNumberWithSig returns the NumberWithSig field if non-nil, zero value otherwise.

### GetNumberWithSigOk

`func (o *StatsByTumorType) GetNumberWithSigOk() (*int32, bool)`

GetNumberWithSigOk returns a tuple with the NumberWithSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberWithSig

`func (o *StatsByTumorType) SetNumberWithSig(v int32)`

SetNumberWithSig sets NumberWithSig field to given value.

### HasNumberWithSig

`func (o *StatsByTumorType) HasNumberWithSig() bool`

HasNumberWithSig returns a boolean if a field has been set.

### GetTmb

`func (o *StatsByTumorType) GetTmb() float64`

GetTmb returns the Tmb field if non-nil, zero value otherwise.

### GetTmbOk

`func (o *StatsByTumorType) GetTmbOk() (*float64, bool)`

GetTmbOk returns a tuple with the Tmb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmb

`func (o *StatsByTumorType) SetTmb(v float64)`

SetTmb sets Tmb field to given value.

### HasTmb

`func (o *StatsByTumorType) HasTmb() bool`

HasTmb returns a boolean if a field has been set.

### GetTumorType

`func (o *StatsByTumorType) GetTumorType() string`

GetTumorType returns the TumorType field if non-nil, zero value otherwise.

### GetTumorTypeOk

`func (o *StatsByTumorType) GetTumorTypeOk() (*string, bool)`

GetTumorTypeOk returns a tuple with the TumorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTumorType

`func (o *StatsByTumorType) SetTumorType(v string)`

SetTumorType sets TumorType field to given value.

### HasTumorType

`func (o *StatsByTumorType) HasTumorType() bool`

HasTumorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


