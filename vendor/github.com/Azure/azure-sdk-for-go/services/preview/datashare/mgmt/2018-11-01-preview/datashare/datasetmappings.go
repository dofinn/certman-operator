package datashare

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DataSetMappingsClient is the creates a Microsoft.DataShare management client.
type DataSetMappingsClient struct {
	BaseClient
}

// NewDataSetMappingsClient creates an instance of the DataSetMappingsClient client.
func NewDataSetMappingsClient(subscriptionID string) DataSetMappingsClient {
	return NewDataSetMappingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataSetMappingsClientWithBaseURI creates an instance of the DataSetMappingsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDataSetMappingsClientWithBaseURI(baseURI string, subscriptionID string) DataSetMappingsClient {
	return DataSetMappingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a DataSetMapping
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareSubscriptionName - the name of the share subscription which will hold the data set sink.
// dataSetMappingName - the name of the data set mapping to be created.
// dataSetMapping - destination data set configuration details.
func (client DataSetMappingsClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, dataSetMapping BasicDataSetMapping) (result DataSetMappingModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataSetMappingsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, shareSubscriptionName, dataSetMappingName, dataSetMapping)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client DataSetMappingsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, dataSetMapping BasicDataSetMapping) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"dataSetMappingName":    autorest.Encode("path", dataSetMappingName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"shareSubscriptionName": autorest.Encode("path", shareSubscriptionName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings/{dataSetMappingName}", pathParameters),
		autorest.WithJSON(dataSetMapping),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client DataSetMappingsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DataSetMappingsClient) CreateResponder(resp *http.Response) (result DataSetMappingModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a DataSetMapping in a shareSubscription
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareSubscriptionName - the name of the shareSubscription.
// dataSetMappingName - the name of the dataSetMapping.
func (client DataSetMappingsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataSetMappingsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, shareSubscriptionName, dataSetMappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DataSetMappingsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"dataSetMappingName":    autorest.Encode("path", dataSetMappingName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"shareSubscriptionName": autorest.Encode("path", shareSubscriptionName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings/{dataSetMappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DataSetMappingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DataSetMappingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a DataSetMapping in a shareSubscription
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareSubscriptionName - the name of the shareSubscription.
// dataSetMappingName - the name of the dataSetMapping.
func (client DataSetMappingsClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string) (result DataSetMappingModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataSetMappingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, shareSubscriptionName, dataSetMappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DataSetMappingsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"dataSetMappingName":    autorest.Encode("path", dataSetMappingName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"shareSubscriptionName": autorest.Encode("path", shareSubscriptionName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings/{dataSetMappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DataSetMappingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DataSetMappingsClient) GetResponder(resp *http.Response) (result DataSetMappingModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByShareSubscription list DataSetMappings in a share subscription
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareSubscriptionName - the name of the share subscription.
// skipToken - continuation token
func (client DataSetMappingsClient) ListByShareSubscription(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result DataSetMappingListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataSetMappingsClient.ListByShareSubscription")
		defer func() {
			sc := -1
			if result.dsml.Response.Response != nil {
				sc = result.dsml.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByShareSubscriptionNextResults
	req, err := client.ListByShareSubscriptionPreparer(ctx, resourceGroupName, accountName, shareSubscriptionName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "ListByShareSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByShareSubscriptionSender(req)
	if err != nil {
		result.dsml.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "ListByShareSubscription", resp, "Failure sending request")
		return
	}

	result.dsml, err = client.ListByShareSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "ListByShareSubscription", resp, "Failure responding to request")
	}

	return
}

// ListByShareSubscriptionPreparer prepares the ListByShareSubscription request.
func (client DataSetMappingsClient) ListByShareSubscriptionPreparer(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"shareSubscriptionName": autorest.Encode("path", shareSubscriptionName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByShareSubscriptionSender sends the ListByShareSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client DataSetMappingsClient) ListByShareSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByShareSubscriptionResponder handles the response to the ListByShareSubscription request. The method always
// closes the http.Response Body.
func (client DataSetMappingsClient) ListByShareSubscriptionResponder(resp *http.Response) (result DataSetMappingList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByShareSubscriptionNextResults retrieves the next set of results, if any.
func (client DataSetMappingsClient) listByShareSubscriptionNextResults(ctx context.Context, lastResults DataSetMappingList) (result DataSetMappingList, err error) {
	req, err := lastResults.dataSetMappingListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "listByShareSubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByShareSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "listByShareSubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByShareSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.DataSetMappingsClient", "listByShareSubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByShareSubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client DataSetMappingsClient) ListByShareSubscriptionComplete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result DataSetMappingListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataSetMappingsClient.ListByShareSubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByShareSubscription(ctx, resourceGroupName, accountName, shareSubscriptionName, skipToken)
	return
}
