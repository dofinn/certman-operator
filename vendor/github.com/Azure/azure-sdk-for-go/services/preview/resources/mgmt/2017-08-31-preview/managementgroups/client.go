// Package managementgroups implements the Azure ARM Managementgroups service API version 2017-08-31-preview.
//
// The Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
//
package managementgroups

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/satori/go.uuid"
)

const (
	// DefaultBaseURI is the default URI used for the service Managementgroups
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Managementgroups.
type BaseClient struct {
	autorest.Client
	BaseURI string
	GroupID uuid.UUID
}

// New creates an instance of the BaseClient client.
func New(groupID uuid.UUID) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, groupID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, groupID uuid.UUID) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
		GroupID: groupID,
	}
}
