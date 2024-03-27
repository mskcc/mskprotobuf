# \EnsemblControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET**](EnsemblControllerApi.md#FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET) | **Get** /ensembl/canonical-gene/entrez/{entrezGeneId} | Retrieves Ensembl canonical gene id by Entrez Gene Id
[**FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST**](EnsemblControllerApi.md#FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST) | **Post** /ensembl/canonical-gene/entrez | Retrieves canonical Ensembl Gene ID by Entrez Gene Ids
[**FetchCanonicalEnsemblGeneIdByHugoSymbolGET**](EnsemblControllerApi.md#FetchCanonicalEnsemblGeneIdByHugoSymbolGET) | **Get** /ensembl/canonical-gene/hgnc/{hugoSymbol} | Retrieves Ensembl canonical gene id by Hugo Symbol
[**FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST**](EnsemblControllerApi.md#FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST) | **Post** /ensembl/canonical-gene/hgnc | Retrieves canonical Ensembl Gene ID by Hugo Symbols
[**FetchCanonicalEnsemblTranscriptByHugoSymbolGET**](EnsemblControllerApi.md#FetchCanonicalEnsemblTranscriptByHugoSymbolGET) | **Get** /ensembl/canonical-transcript/hgnc/{hugoSymbol} | Retrieves Ensembl canonical transcript by Hugo Symbol
[**FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST**](EnsemblControllerApi.md#FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST) | **Post** /ensembl/canonical-transcript/hgnc | Retrieves Ensembl canonical transcripts by Hugo Symbols
[**FetchEnsemblTranscriptByTranscriptIdGET**](EnsemblControllerApi.md#FetchEnsemblTranscriptByTranscriptIdGET) | **Get** /ensembl/transcript/{transcriptId} | Retrieves the transcript by an Ensembl transcript ID
[**FetchEnsemblTranscriptsByEnsemblFilterPOST**](EnsemblControllerApi.md#FetchEnsemblTranscriptsByEnsemblFilterPOST) | **Post** /ensembl/transcript | Retrieves Ensembl Transcripts by Ensembl transcript IDs, hugo Symbols, protein IDs, or gene IDs
[**FetchEnsemblTranscriptsGET**](EnsemblControllerApi.md#FetchEnsemblTranscriptsGET) | **Get** /ensembl/transcript | Retrieves Ensembl Transcripts by protein ID, and gene ID. Retrieves all transcripts in case no query parameter provided
[**FetchGeneXrefsGET**](EnsemblControllerApi.md#FetchGeneXrefsGET) | **Get** /ensembl/xrefs | Perform lookups of Ensembl identifiers and retrieve their external references in other databases



## FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET

> EnsemblGene FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET(ctx, entrezGeneId).Execute()

Retrieves Ensembl canonical gene id by Entrez Gene Id

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
    entrezGeneId := "entrezGeneId_example" // string | An Entrez Gene Id. For example 23140

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET(context.Background(), entrezGeneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET`: EnsemblGene
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByEntrezGeneIdGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entrezGeneId** | **string** | An Entrez Gene Id. For example 23140 | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCanonicalEnsemblGeneIdByEntrezGeneIdGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnsemblGene**](EnsemblGene.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST

> []EnsemblGene FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST(ctx).EntrezGeneIds(entrezGeneIds).Execute()

Retrieves canonical Ensembl Gene ID by Entrez Gene Ids

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
    entrezGeneIds := []string{"Property_example"} // []string | List of Entrez Gene Ids. For example [\"23140\",\"26009\",\"100131879\"]

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST(context.Background()).EntrezGeneIds(entrezGeneIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST`: []EnsemblGene
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCanonicalEnsemblGeneIdByEntrezGeneIdsPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entrezGeneIds** | **[]string** | List of Entrez Gene Ids. For example [\&quot;23140\&quot;,\&quot;26009\&quot;,\&quot;100131879\&quot;] | 

### Return type

[**[]EnsemblGene**](EnsemblGene.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCanonicalEnsemblGeneIdByHugoSymbolGET

> EnsemblGene FetchCanonicalEnsemblGeneIdByHugoSymbolGET(ctx, hugoSymbol).Execute()

Retrieves Ensembl canonical gene id by Hugo Symbol

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
    hugoSymbol := "hugoSymbol_example" // string | A Hugo Symbol. For example TP53

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByHugoSymbolGET(context.Background(), hugoSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByHugoSymbolGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCanonicalEnsemblGeneIdByHugoSymbolGET`: EnsemblGene
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByHugoSymbolGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hugoSymbol** | **string** | A Hugo Symbol. For example TP53 | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCanonicalEnsemblGeneIdByHugoSymbolGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnsemblGene**](EnsemblGene.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST

> []EnsemblGene FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST(ctx).HugoSymbols(hugoSymbols).Execute()

Retrieves canonical Ensembl Gene ID by Hugo Symbols

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
    hugoSymbols := []string{"Property_example"} // []string | List of Hugo Symbols. For example [\"TP53\",\"PIK3CA\",\"BRCA1\"]

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST(context.Background()).HugoSymbols(hugoSymbols).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST`: []EnsemblGene
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchCanonicalEnsemblGeneIdByHugoSymbolsPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCanonicalEnsemblGeneIdByHugoSymbolsPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hugoSymbols** | **[]string** | List of Hugo Symbols. For example [\&quot;TP53\&quot;,\&quot;PIK3CA\&quot;,\&quot;BRCA1\&quot;] | 

### Return type

[**[]EnsemblGene**](EnsemblGene.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCanonicalEnsemblTranscriptByHugoSymbolGET

> EnsemblTranscript FetchCanonicalEnsemblTranscriptByHugoSymbolGET(ctx, hugoSymbol).IsoformOverrideSource(isoformOverrideSource).Execute()

Retrieves Ensembl canonical transcript by Hugo Symbol

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
    hugoSymbol := "hugoSymbol_example" // string | A Hugo Symbol. For example TP53
    isoformOverrideSource := "isoformOverrideSource_example" // string | Isoform override source. For example uniprot (optional) (default to "uniprot")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchCanonicalEnsemblTranscriptByHugoSymbolGET(context.Background(), hugoSymbol).IsoformOverrideSource(isoformOverrideSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchCanonicalEnsemblTranscriptByHugoSymbolGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCanonicalEnsemblTranscriptByHugoSymbolGET`: EnsemblTranscript
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchCanonicalEnsemblTranscriptByHugoSymbolGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hugoSymbol** | **string** | A Hugo Symbol. For example TP53 | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCanonicalEnsemblTranscriptByHugoSymbolGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isoformOverrideSource** | **string** | Isoform override source. For example uniprot | [default to &quot;uniprot&quot;]

### Return type

[**EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST

> []EnsemblTranscript FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST(ctx).HugoSymbols(hugoSymbols).IsoformOverrideSource(isoformOverrideSource).Execute()

Retrieves Ensembl canonical transcripts by Hugo Symbols

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
    hugoSymbols := []string{"Property_example"} // []string | List of Hugo Symbols. For example [\"TP53\",\"PIK3CA\",\"BRCA1\"]
    isoformOverrideSource := "isoformOverrideSource_example" // string | Isoform override source. For example uniprot (optional) (default to "uniprot")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST(context.Background()).HugoSymbols(hugoSymbols).IsoformOverrideSource(isoformOverrideSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST`: []EnsemblTranscript
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchCanonicalEnsemblTranscriptsByHugoSymbolsPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCanonicalEnsemblTranscriptsByHugoSymbolsPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hugoSymbols** | **[]string** | List of Hugo Symbols. For example [\&quot;TP53\&quot;,\&quot;PIK3CA\&quot;,\&quot;BRCA1\&quot;] | 
 **isoformOverrideSource** | **string** | Isoform override source. For example uniprot | [default to &quot;uniprot&quot;]

### Return type

[**[]EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEnsemblTranscriptByTranscriptIdGET

> EnsemblTranscript FetchEnsemblTranscriptByTranscriptIdGET(ctx, transcriptId).Execute()

Retrieves the transcript by an Ensembl transcript ID

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
    transcriptId := "transcriptId_example" // string | An Ensembl transcript ID. For example ENST00000361390

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchEnsemblTranscriptByTranscriptIdGET(context.Background(), transcriptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchEnsemblTranscriptByTranscriptIdGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEnsemblTranscriptByTranscriptIdGET`: EnsemblTranscript
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchEnsemblTranscriptByTranscriptIdGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transcriptId** | **string** | An Ensembl transcript ID. For example ENST00000361390 | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchEnsemblTranscriptByTranscriptIdGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEnsemblTranscriptsByEnsemblFilterPOST

> []EnsemblTranscript FetchEnsemblTranscriptsByEnsemblFilterPOST(ctx).EnsemblFilter(ensemblFilter).Execute()

Retrieves Ensembl Transcripts by Ensembl transcript IDs, hugo Symbols, protein IDs, or gene IDs

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
    ensemblFilter := *openapiclient.NewEnsemblFilter() // EnsemblFilter | List of Ensembl transcript IDs. For example [\"ENST00000361390\", \"ENST00000361453\", \"ENST00000361624\"]<br>OR<br>List of Hugo Symbols. For example [\"TP53\", \"PIK3CA\", \"BRCA1\"]<br>OR<br>List of Ensembl protein IDs. For example [\"ENSP00000439985\", \"ENSP00000478460\", \"ENSP00000346196\"]<br>OR<br>List of Ensembl gene IDs. For example [\"ENSG00000136999\", \"ENSG00000272398\", \"ENSG00000198695\"]

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchEnsemblTranscriptsByEnsemblFilterPOST(context.Background()).EnsemblFilter(ensemblFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchEnsemblTranscriptsByEnsemblFilterPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEnsemblTranscriptsByEnsemblFilterPOST`: []EnsemblTranscript
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchEnsemblTranscriptsByEnsemblFilterPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchEnsemblTranscriptsByEnsemblFilterPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ensemblFilter** | [**EnsemblFilter**](EnsemblFilter.md) | List of Ensembl transcript IDs. For example [\&quot;ENST00000361390\&quot;, \&quot;ENST00000361453\&quot;, \&quot;ENST00000361624\&quot;]&lt;br&gt;OR&lt;br&gt;List of Hugo Symbols. For example [\&quot;TP53\&quot;, \&quot;PIK3CA\&quot;, \&quot;BRCA1\&quot;]&lt;br&gt;OR&lt;br&gt;List of Ensembl protein IDs. For example [\&quot;ENSP00000439985\&quot;, \&quot;ENSP00000478460\&quot;, \&quot;ENSP00000346196\&quot;]&lt;br&gt;OR&lt;br&gt;List of Ensembl gene IDs. For example [\&quot;ENSG00000136999\&quot;, \&quot;ENSG00000272398\&quot;, \&quot;ENSG00000198695\&quot;] | 

### Return type

[**[]EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEnsemblTranscriptsGET

> []EnsemblTranscript FetchEnsemblTranscriptsGET(ctx).GeneId(geneId).ProteinId(proteinId).HugoSymbol(hugoSymbol).Execute()

Retrieves Ensembl Transcripts by protein ID, and gene ID. Retrieves all transcripts in case no query parameter provided

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
    geneId := "geneId_example" // string | An Ensembl gene ID. For example ENSG00000136999 (optional)
    proteinId := "proteinId_example" // string | An Ensembl protein ID. For example ENSP00000439985 (optional)
    hugoSymbol := "hugoSymbol_example" // string | A Hugo Symbol For example ARF5 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchEnsemblTranscriptsGET(context.Background()).GeneId(geneId).ProteinId(proteinId).HugoSymbol(hugoSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchEnsemblTranscriptsGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEnsemblTranscriptsGET`: []EnsemblTranscript
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchEnsemblTranscriptsGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchEnsemblTranscriptsGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **geneId** | **string** | An Ensembl gene ID. For example ENSG00000136999 | 
 **proteinId** | **string** | An Ensembl protein ID. For example ENSP00000439985 | 
 **hugoSymbol** | **string** | A Hugo Symbol For example ARF5 | 

### Return type

[**[]EnsemblTranscript**](EnsemblTranscript.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchGeneXrefsGET

> []GeneXref FetchGeneXrefsGET(ctx).Accession(accession).Execute()

Perform lookups of Ensembl identifiers and retrieve their external references in other databases

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
    accession := "accession_example" // string | Ensembl gene accession. For example ENSG00000169083

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnsemblControllerApi.FetchGeneXrefsGET(context.Background()).Accession(accession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnsemblControllerApi.FetchGeneXrefsGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchGeneXrefsGET`: []GeneXref
    fmt.Fprintf(os.Stdout, "Response from `EnsemblControllerApi.FetchGeneXrefsGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchGeneXrefsGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accession** | **string** | Ensembl gene accession. For example ENSG00000169083 | 

### Return type

[**[]GeneXref**](GeneXref.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

