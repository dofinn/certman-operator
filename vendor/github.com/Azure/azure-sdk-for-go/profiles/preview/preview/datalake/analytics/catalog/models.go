// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package catalog

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/datalake/analytics/2015-10-01-preview/catalog"
)

const (
	DefaultAdlaCatalogDNSSuffix = original.DefaultAdlaCatalogDNSSuffix
)

type FileType = original.FileType

const (
	Assembly FileType = original.Assembly
	Resource FileType = original.Resource
)

type BaseClient = original.BaseClient
type Client = original.Client
type DataLakeAnalyticsCatalogSecretCreateOrUpdateParameters = original.DataLakeAnalyticsCatalogSecretCreateOrUpdateParameters
type DdlName = original.DdlName
type EntityID = original.EntityID
type ExternalTable = original.ExternalTable
type Item = original.Item
type ItemList = original.ItemList
type TypeFieldInfo = original.TypeFieldInfo
type USQLAssembly = original.USQLAssembly
type USQLAssemblyClr = original.USQLAssemblyClr
type USQLAssemblyDependencyInfo = original.USQLAssemblyDependencyInfo
type USQLAssemblyFileInfo = original.USQLAssemblyFileInfo
type USQLAssemblyList = original.USQLAssemblyList
type USQLAssemblyListIterator = original.USQLAssemblyListIterator
type USQLAssemblyListPage = original.USQLAssemblyListPage
type USQLCredential = original.USQLCredential
type USQLCredentialList = original.USQLCredentialList
type USQLCredentialListIterator = original.USQLCredentialListIterator
type USQLCredentialListPage = original.USQLCredentialListPage
type USQLDatabase = original.USQLDatabase
type USQLDatabaseList = original.USQLDatabaseList
type USQLDatabaseListIterator = original.USQLDatabaseListIterator
type USQLDatabaseListPage = original.USQLDatabaseListPage
type USQLDirectedColumn = original.USQLDirectedColumn
type USQLDistributionInfo = original.USQLDistributionInfo
type USQLExternalDataSource = original.USQLExternalDataSource
type USQLExternalDataSourceList = original.USQLExternalDataSourceList
type USQLExternalDataSourceListIterator = original.USQLExternalDataSourceListIterator
type USQLExternalDataSourceListPage = original.USQLExternalDataSourceListPage
type USQLIndex = original.USQLIndex
type USQLProcedure = original.USQLProcedure
type USQLProcedureList = original.USQLProcedureList
type USQLProcedureListIterator = original.USQLProcedureListIterator
type USQLProcedureListPage = original.USQLProcedureListPage
type USQLSchema = original.USQLSchema
type USQLSchemaList = original.USQLSchemaList
type USQLSchemaListIterator = original.USQLSchemaListIterator
type USQLSchemaListPage = original.USQLSchemaListPage
type USQLSecret = original.USQLSecret
type USQLTable = original.USQLTable
type USQLTableColumn = original.USQLTableColumn
type USQLTableList = original.USQLTableList
type USQLTableListIterator = original.USQLTableListIterator
type USQLTableListPage = original.USQLTableListPage
type USQLTablePartition = original.USQLTablePartition
type USQLTablePartitionList = original.USQLTablePartitionList
type USQLTablePartitionListIterator = original.USQLTablePartitionListIterator
type USQLTablePartitionListPage = original.USQLTablePartitionListPage
type USQLTableStatistics = original.USQLTableStatistics
type USQLTableStatisticsList = original.USQLTableStatisticsList
type USQLTableStatisticsListIterator = original.USQLTableStatisticsListIterator
type USQLTableStatisticsListPage = original.USQLTableStatisticsListPage
type USQLTableType = original.USQLTableType
type USQLTableTypeList = original.USQLTableTypeList
type USQLTableTypeListIterator = original.USQLTableTypeListIterator
type USQLTableTypeListPage = original.USQLTableTypeListPage
type USQLTableValuedFunction = original.USQLTableValuedFunction
type USQLTableValuedFunctionList = original.USQLTableValuedFunctionList
type USQLTableValuedFunctionListIterator = original.USQLTableValuedFunctionListIterator
type USQLTableValuedFunctionListPage = original.USQLTableValuedFunctionListPage
type USQLType = original.USQLType
type USQLTypeList = original.USQLTypeList
type USQLTypeListIterator = original.USQLTypeListIterator
type USQLTypeListPage = original.USQLTypeListPage
type USQLView = original.USQLView
type USQLViewList = original.USQLViewList
type USQLViewListIterator = original.USQLViewListIterator
type USQLViewListPage = original.USQLViewListPage

func New() BaseClient {
	return original.New()
}
func NewClient() Client {
	return original.NewClient()
}
func NewUSQLAssemblyListIterator(page USQLAssemblyListPage) USQLAssemblyListIterator {
	return original.NewUSQLAssemblyListIterator(page)
}
func NewUSQLAssemblyListPage(getNextPage func(context.Context, USQLAssemblyList) (USQLAssemblyList, error)) USQLAssemblyListPage {
	return original.NewUSQLAssemblyListPage(getNextPage)
}
func NewUSQLCredentialListIterator(page USQLCredentialListPage) USQLCredentialListIterator {
	return original.NewUSQLCredentialListIterator(page)
}
func NewUSQLCredentialListPage(getNextPage func(context.Context, USQLCredentialList) (USQLCredentialList, error)) USQLCredentialListPage {
	return original.NewUSQLCredentialListPage(getNextPage)
}
func NewUSQLDatabaseListIterator(page USQLDatabaseListPage) USQLDatabaseListIterator {
	return original.NewUSQLDatabaseListIterator(page)
}
func NewUSQLDatabaseListPage(getNextPage func(context.Context, USQLDatabaseList) (USQLDatabaseList, error)) USQLDatabaseListPage {
	return original.NewUSQLDatabaseListPage(getNextPage)
}
func NewUSQLExternalDataSourceListIterator(page USQLExternalDataSourceListPage) USQLExternalDataSourceListIterator {
	return original.NewUSQLExternalDataSourceListIterator(page)
}
func NewUSQLExternalDataSourceListPage(getNextPage func(context.Context, USQLExternalDataSourceList) (USQLExternalDataSourceList, error)) USQLExternalDataSourceListPage {
	return original.NewUSQLExternalDataSourceListPage(getNextPage)
}
func NewUSQLProcedureListIterator(page USQLProcedureListPage) USQLProcedureListIterator {
	return original.NewUSQLProcedureListIterator(page)
}
func NewUSQLProcedureListPage(getNextPage func(context.Context, USQLProcedureList) (USQLProcedureList, error)) USQLProcedureListPage {
	return original.NewUSQLProcedureListPage(getNextPage)
}
func NewUSQLSchemaListIterator(page USQLSchemaListPage) USQLSchemaListIterator {
	return original.NewUSQLSchemaListIterator(page)
}
func NewUSQLSchemaListPage(getNextPage func(context.Context, USQLSchemaList) (USQLSchemaList, error)) USQLSchemaListPage {
	return original.NewUSQLSchemaListPage(getNextPage)
}
func NewUSQLTableListIterator(page USQLTableListPage) USQLTableListIterator {
	return original.NewUSQLTableListIterator(page)
}
func NewUSQLTableListPage(getNextPage func(context.Context, USQLTableList) (USQLTableList, error)) USQLTableListPage {
	return original.NewUSQLTableListPage(getNextPage)
}
func NewUSQLTablePartitionListIterator(page USQLTablePartitionListPage) USQLTablePartitionListIterator {
	return original.NewUSQLTablePartitionListIterator(page)
}
func NewUSQLTablePartitionListPage(getNextPage func(context.Context, USQLTablePartitionList) (USQLTablePartitionList, error)) USQLTablePartitionListPage {
	return original.NewUSQLTablePartitionListPage(getNextPage)
}
func NewUSQLTableStatisticsListIterator(page USQLTableStatisticsListPage) USQLTableStatisticsListIterator {
	return original.NewUSQLTableStatisticsListIterator(page)
}
func NewUSQLTableStatisticsListPage(getNextPage func(context.Context, USQLTableStatisticsList) (USQLTableStatisticsList, error)) USQLTableStatisticsListPage {
	return original.NewUSQLTableStatisticsListPage(getNextPage)
}
func NewUSQLTableTypeListIterator(page USQLTableTypeListPage) USQLTableTypeListIterator {
	return original.NewUSQLTableTypeListIterator(page)
}
func NewUSQLTableTypeListPage(getNextPage func(context.Context, USQLTableTypeList) (USQLTableTypeList, error)) USQLTableTypeListPage {
	return original.NewUSQLTableTypeListPage(getNextPage)
}
func NewUSQLTableValuedFunctionListIterator(page USQLTableValuedFunctionListPage) USQLTableValuedFunctionListIterator {
	return original.NewUSQLTableValuedFunctionListIterator(page)
}
func NewUSQLTableValuedFunctionListPage(getNextPage func(context.Context, USQLTableValuedFunctionList) (USQLTableValuedFunctionList, error)) USQLTableValuedFunctionListPage {
	return original.NewUSQLTableValuedFunctionListPage(getNextPage)
}
func NewUSQLTypeListIterator(page USQLTypeListPage) USQLTypeListIterator {
	return original.NewUSQLTypeListIterator(page)
}
func NewUSQLTypeListPage(getNextPage func(context.Context, USQLTypeList) (USQLTypeList, error)) USQLTypeListPage {
	return original.NewUSQLTypeListPage(getNextPage)
}
func NewUSQLViewListIterator(page USQLViewListPage) USQLViewListIterator {
	return original.NewUSQLViewListIterator(page)
}
func NewUSQLViewListPage(getNextPage func(context.Context, USQLViewList) (USQLViewList, error)) USQLViewListPage {
	return original.NewUSQLViewListPage(getNextPage)
}
func NewWithoutDefaults(adlaCatalogDNSSuffix string) BaseClient {
	return original.NewWithoutDefaults(adlaCatalogDNSSuffix)
}
func PossibleFileTypeValues() []FileType {
	return original.PossibleFileTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
