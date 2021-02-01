# \NWDAFEventsSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nnwdaf-eventssubscription/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNWDAFEventsSubscription**](NWDAFEventsSubscriptionsCollectionApi.md#CreateNWDAFEventsSubscription) | **Post** /subscriptions | Create a new Individual NWDAF Events Subscription



## CreateNWDAFEventsSubscription

> NnwdafEventsSubscription CreateNWDAFEventsSubscription(ctx).NnwdafEventsSubscription(nnwdafEventsSubscription).Execute()

Create a new Individual NWDAF Events Subscription

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
    nnwdafEventsSubscription := *openapiclient.NewNnwdafEventsSubscription([]openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewNwdafEvent())}) // NnwdafEventsSubscription | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NWDAFEventsSubscriptionsCollectionApi.CreateNWDAFEventsSubscription(context.Background()).NnwdafEventsSubscription(nnwdafEventsSubscription).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NWDAFEventsSubscriptionsCollectionApi.CreateNWDAFEventsSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNWDAFEventsSubscription`: NnwdafEventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `NWDAFEventsSubscriptionsCollectionApi.CreateNWDAFEventsSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNWDAFEventsSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nnwdafEventsSubscription** | [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | 

### Return type

[**NnwdafEventsSubscription**](NnwdafEventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

