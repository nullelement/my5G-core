# \IndividualNWDAFEventsSubscriptionDocumentApi

All URIs are relative to *https://example.com/nnwdaf-eventssubscription/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNWDAFEventsSubscription**](IndividualNWDAFEventsSubscriptionDocumentApi.md#DeleteNWDAFEventsSubscription) | **Delete** /subscriptions/{subscriptionId} | Delete an existing Individual NWDAF Events Subscription
[**UpdateNWDAFEventsSubscription**](IndividualNWDAFEventsSubscriptionDocumentApi.md#UpdateNWDAFEventsSubscription) | **Put** /subscriptions/{subscriptionId} | Update an existing Individual NWDAF Events Subscription



## DeleteNWDAFEventsSubscription

> DeleteNWDAFEventsSubscription(ctx, subscriptionId).Execute()

Delete an existing Individual NWDAF Events Subscription

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
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Nnwdaf_EventsSubscription Service

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndividualNWDAFEventsSubscriptionDocumentApi.DeleteNWDAFEventsSubscription(context.Background(), subscriptionId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNWDAFEventsSubscriptionDocumentApi.DeleteNWDAFEventsSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Nnwdaf_EventsSubscription Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNWDAFEventsSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNWDAFEventsSubscription

> NnwdafEventsSubscription UpdateNWDAFEventsSubscription(ctx, subscriptionId).NnwdafEventsSubscription(nnwdafEventsSubscription).Execute()

Update an existing Individual NWDAF Events Subscription

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
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Nnwdaf_EventsSubscription Service
    nnwdafEventsSubscription := *openapiclient.NewNnwdafEventsSubscription([]openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewNwdafEvent())}) // NnwdafEventsSubscription | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndividualNWDAFEventsSubscriptionDocumentApi.UpdateNWDAFEventsSubscription(context.Background(), subscriptionId).NnwdafEventsSubscription(nnwdafEventsSubscription).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNWDAFEventsSubscriptionDocumentApi.UpdateNWDAFEventsSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNWDAFEventsSubscription`: NnwdafEventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualNWDAFEventsSubscriptionDocumentApi.UpdateNWDAFEventsSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Nnwdaf_EventsSubscription Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNWDAFEventsSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nnwdafEventsSubscription** | [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | 

### Return type

[**NnwdafEventsSubscription**](NnwdafEventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

