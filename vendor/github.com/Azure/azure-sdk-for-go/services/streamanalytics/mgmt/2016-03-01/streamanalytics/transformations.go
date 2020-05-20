package streamanalytics

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

// TransformationsClient is the stream Analytics Client
type TransformationsClient struct {
	BaseClient
}

// NewTransformationsClient creates an instance of the TransformationsClient client.
func NewTransformationsClient(subscriptionID string) TransformationsClient {
	return NewTransformationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTransformationsClientWithBaseURI creates an instance of the TransformationsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTransformationsClientWithBaseURI(baseURI string, subscriptionID string) TransformationsClient {
	return TransformationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrReplace creates a transformation or replaces an already existing transformation under an existing streaming
// job.
// Parameters:
// transformation - the definition of the transformation that will be used to create a new transformation or
// replace the existing one under the streaming job.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// transformationName - the name of the transformation.
// ifMatch - the ETag of the transformation. Omit this value to always overwrite the current transformation.
// Specify the last-seen ETag value to prevent accidentally overwriting concurrent changes.
// ifNoneMatch - set to '*' to allow a new transformation to be created, but to prevent updating an existing
// transformation. Other values will result in a 412 Pre-condition Failed response.
func (client TransformationsClient) CreateOrReplace(ctx context.Context, transformation Transformation, resourceGroupName string, jobName string, transformationName string, ifMatch string, ifNoneMatch string) (result Transformation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransformationsClient.CreateOrReplace")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrReplacePreparer(ctx, transformation, resourceGroupName, jobName, transformationName, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.TransformationsClient", "CreateOrReplace", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrReplaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.TransformationsClient", "CreateOrReplace", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrReplaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.TransformationsClient", "CreateOrReplace", resp, "Failure responding to request")
	}

	return
}

// CreateOrReplacePreparer prepares the CreateOrReplace request.
func (client TransformationsClient) CreateOrReplacePreparer(ctx context.Context, transformation Transformation, resourceGroupName string, jobName string, transformationName string, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":            autorest.Encode("path", jobName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"transformationName": autorest.Encode("path", transformationName),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", pathParameters),
		autorest.WithJSON(transformation),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrReplaceSender sends the CreateOrReplace request. The method will close the
// http.Response Body if it receives an error.
func (client TransformationsClient) CreateOrReplaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrReplaceResponder handles the response to the CreateOrReplace request. The method always
// closes the http.Response Body.
func (client TransformationsClient) CreateOrReplaceResponder(resp *http.Response) (result Transformation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets details about the specified transformation.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// transformationName - the name of the transformation.
func (client TransformationsClient) Get(ctx context.Context, resourceGroupName string, jobName string, transformationName string) (result Transformation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransformationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, jobName, transformationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.TransformationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.TransformationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.TransformationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TransformationsClient) GetPreparer(ctx context.Context, resourceGroupName string, jobName string, transformationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":            autorest.Encode("path", jobName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"transformationName": autorest.Encode("path", transformationName),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TransformationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TransformationsClient) GetResponder(resp *http.Response) (result Transformation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing transformation under an existing streaming job. This can be used to partially update (ie.
// update one or two properties) a transformation without affecting the rest the job or transformation definition.
// Parameters:
// transformation - a Transformation object. The properties specified here will overwrite the corresponding
// properties in the existing transformation (ie. Those properties will be updated). Any properties that are
// set to null here will mean that the corresponding property in the existing transformation will remain the
// same and not change as a result of this PATCH operation.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// jobName - the name of the streaming job.
// transformationName - the name of the transformation.
// ifMatch - the ETag of the transformation. Omit this value to always overwrite the current transformation.
// Specify the last-seen ETag value to prevent accidentally overwriting concurrent changes.
func (client TransformationsClient) Update(ctx context.Context, transformation Transformation, resourceGroupName string, jobName string, transformationName string, ifMatch string) (result Transformation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransformationsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, transformation, resourceGroupName, jobName, transformationName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.TransformationsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.TransformationsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.TransformationsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client TransformationsClient) UpdatePreparer(ctx context.Context, transformation Transformation, resourceGroupName string, jobName string, transformationName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":            autorest.Encode("path", jobName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"transformationName": autorest.Encode("path", transformationName),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", pathParameters),
		autorest.WithJSON(transformation),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client TransformationsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client TransformationsClient) UpdateResponder(resp *http.Response) (result Transformation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
