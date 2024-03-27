# \PtmControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPostTranslationalModificationsByPtmFilterPOST**](PtmControllerApi.md#FetchPostTranslationalModificationsByPtmFilterPOST) | **Post** /ptm/experimental | Retrieves PTM entries by Ensembl Transcript IDs
[**FetchPostTranslationalModificationsGET**](PtmControllerApi.md#FetchPostTranslationalModificationsGET) | **Get** /ptm/experimental | Retrieves PTM entries by Ensembl Transcript ID



## FetchPostTranslationalModificationsByPtmFilterPOST

> []PostTranslationalModification FetchPostTranslationalModificationsByPtmFilterPOST(ctx).PtmFilter(ptmFilter).Execute()

Retrieves PTM entries by Ensembl Transcript IDs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ptmFilter := *openapiclient.NewPtmFilter() // PtmFilter | List of Ensembl transcript IDs. For example [\"ENST00000420316\", \"ENST00000646891\", \"ENST00000371953\"]

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PtmControllerApi.FetchPostTranslationalModificationsByPtmFilterPOST(context.Background()).PtmFilter(ptmFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PtmControllerApi.FetchPostTranslationalModificationsByPtmFilterPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPostTranslationalModificationsByPtmFilterPOST`: []PostTranslationalModification
    fmt.Fprintf(os.Stdout, "Response from `PtmControllerApi.FetchPostTranslationalModificationsByPtmFilterPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchPostTranslationalModificationsByPtmFilterPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ptmFilter** | [**PtmFilter**](PtmFilter.md) | List of Ensembl transcript IDs. For example [\&quot;ENST00000420316\&quot;, \&quot;ENST00000646891\&quot;, \&quot;ENST00000371953\&quot;] | 

### Return type

[**[]PostTranslationalModification**](PostTranslationalModification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPostTranslationalModificationsGET

> []PostTranslationalModification FetchPostTranslationalModificationsGET(ctx).EnsemblTranscriptId(ensemblTranscriptId).Execute()

Retrieves PTM entries by Ensembl Transcript ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ensemblTranscriptId := "ensemblTranscriptId_example" // string | Ensembl Transcript ID. For example ENST00000646891 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PtmControllerApi.FetchPostTranslationalModificationsGET(context.Background()).EnsemblTranscriptId(ensemblTranscriptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PtmControllerApi.FetchPostTranslationalModificationsGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPostTranslationalModificationsGET`: []PostTranslationalModification
    fmt.Fprintf(os.Stdout, "Response from `PtmControllerApi.FetchPostTranslationalModificationsGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchPostTranslationalModificationsGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ensemblTranscriptId** | **string** | Ensembl Transcript ID. For example ENST00000646891 | 

### Return type

[**[]PostTranslationalModification**](PostTranslationalModification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

