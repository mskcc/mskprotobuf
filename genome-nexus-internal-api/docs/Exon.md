# Exon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExonId** | **string** | Exon id | 
**ExonStart** | **int32** | Start position of exon | 
**ExonEnd** | **int32** | End position of exon | 
**Rank** | **int32** | Number of exon in transcript | 
**Strand** | **int32** | Strand exon is on, -1 for - and 1 for + | 
**Version** | **int32** | Exon version | 

## Methods

### NewExon

`func NewExon(exonId string, exonStart int32, exonEnd int32, rank int32, strand int32, version int32, ) *Exon`

NewExon instantiates a new Exon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExonWithDefaults

`func NewExonWithDefaults() *Exon`

NewExonWithDefaults instantiates a new Exon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExonId

`func (o *Exon) GetExonId() string`

GetExonId returns the ExonId field if non-nil, zero value otherwise.

### GetExonIdOk

`func (o *Exon) GetExonIdOk() (*string, bool)`

GetExonIdOk returns a tuple with the ExonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExonId

`func (o *Exon) SetExonId(v string)`

SetExonId sets ExonId field to given value.


### GetExonStart

`func (o *Exon) GetExonStart() int32`

GetExonStart returns the ExonStart field if non-nil, zero value otherwise.

### GetExonStartOk

`func (o *Exon) GetExonStartOk() (*int32, bool)`

GetExonStartOk returns a tuple with the ExonStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExonStart

`func (o *Exon) SetExonStart(v int32)`

SetExonStart sets ExonStart field to given value.


### GetExonEnd

`func (o *Exon) GetExonEnd() int32`

GetExonEnd returns the ExonEnd field if non-nil, zero value otherwise.

### GetExonEndOk

`func (o *Exon) GetExonEndOk() (*int32, bool)`

GetExonEndOk returns a tuple with the ExonEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExonEnd

`func (o *Exon) SetExonEnd(v int32)`

SetExonEnd sets ExonEnd field to given value.


### GetRank

`func (o *Exon) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Exon) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Exon) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetStrand

`func (o *Exon) GetStrand() int32`

GetStrand returns the Strand field if non-nil, zero value otherwise.

### GetStrandOk

`func (o *Exon) GetStrandOk() (*int32, bool)`

GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrand

`func (o *Exon) SetStrand(v int32)`

SetStrand sets Strand field to given value.


### GetVersion

`func (o *Exon) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Exon) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Exon) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


