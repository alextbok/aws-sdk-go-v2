// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rdsiface provides an interface to enable mocking the Amazon Relational Database Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rdsiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

// ClientAPI provides an interface to enable mocking the
// rds.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon RDS.
//    func myFunc(svc rdsiface.ClientAPI) bool {
//        // Make svc.AddRoleToDBCluster request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := rds.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        rdsiface.ClientPI
//    }
//    func (m *mockClientClient) AddRoleToDBCluster(input *rds.AddRoleToDBClusterInput) (*rds.AddRoleToDBClusterOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AddRoleToDBClusterRequest(*rds.AddRoleToDBClusterInput) rds.AddRoleToDBClusterRequest

	AddRoleToDBInstanceRequest(*rds.AddRoleToDBInstanceInput) rds.AddRoleToDBInstanceRequest

	AddSourceIdentifierToSubscriptionRequest(*rds.AddSourceIdentifierToSubscriptionInput) rds.AddSourceIdentifierToSubscriptionRequest

	AddTagsToResourceRequest(*rds.AddTagsToResourceInput) rds.AddTagsToResourceRequest

	ApplyPendingMaintenanceActionRequest(*rds.ApplyPendingMaintenanceActionInput) rds.ApplyPendingMaintenanceActionRequest

	AuthorizeDBSecurityGroupIngressRequest(*rds.AuthorizeDBSecurityGroupIngressInput) rds.AuthorizeDBSecurityGroupIngressRequest

	BacktrackDBClusterRequest(*rds.BacktrackDBClusterInput) rds.BacktrackDBClusterRequest

	CopyDBClusterParameterGroupRequest(*rds.CopyDBClusterParameterGroupInput) rds.CopyDBClusterParameterGroupRequest

	CopyDBClusterSnapshotRequest(*rds.CopyDBClusterSnapshotInput) rds.CopyDBClusterSnapshotRequest

	CopyDBParameterGroupRequest(*rds.CopyDBParameterGroupInput) rds.CopyDBParameterGroupRequest

	CopyDBSnapshotRequest(*rds.CopyDBSnapshotInput) rds.CopyDBSnapshotRequest

	CopyOptionGroupRequest(*rds.CopyOptionGroupInput) rds.CopyOptionGroupRequest

	CreateCustomAvailabilityZoneRequest(*rds.CreateCustomAvailabilityZoneInput) rds.CreateCustomAvailabilityZoneRequest

	CreateDBClusterRequest(*rds.CreateDBClusterInput) rds.CreateDBClusterRequest

	CreateDBClusterEndpointRequest(*rds.CreateDBClusterEndpointInput) rds.CreateDBClusterEndpointRequest

	CreateDBClusterParameterGroupRequest(*rds.CreateDBClusterParameterGroupInput) rds.CreateDBClusterParameterGroupRequest

	CreateDBClusterSnapshotRequest(*rds.CreateDBClusterSnapshotInput) rds.CreateDBClusterSnapshotRequest

	CreateDBInstanceRequest(*rds.CreateDBInstanceInput) rds.CreateDBInstanceRequest

	CreateDBInstanceReadReplicaRequest(*rds.CreateDBInstanceReadReplicaInput) rds.CreateDBInstanceReadReplicaRequest

	CreateDBParameterGroupRequest(*rds.CreateDBParameterGroupInput) rds.CreateDBParameterGroupRequest

	CreateDBSecurityGroupRequest(*rds.CreateDBSecurityGroupInput) rds.CreateDBSecurityGroupRequest

	CreateDBSnapshotRequest(*rds.CreateDBSnapshotInput) rds.CreateDBSnapshotRequest

	CreateDBSubnetGroupRequest(*rds.CreateDBSubnetGroupInput) rds.CreateDBSubnetGroupRequest

	CreateEventSubscriptionRequest(*rds.CreateEventSubscriptionInput) rds.CreateEventSubscriptionRequest

	CreateGlobalClusterRequest(*rds.CreateGlobalClusterInput) rds.CreateGlobalClusterRequest

	CreateOptionGroupRequest(*rds.CreateOptionGroupInput) rds.CreateOptionGroupRequest

	DeleteCustomAvailabilityZoneRequest(*rds.DeleteCustomAvailabilityZoneInput) rds.DeleteCustomAvailabilityZoneRequest

	DeleteDBClusterRequest(*rds.DeleteDBClusterInput) rds.DeleteDBClusterRequest

	DeleteDBClusterEndpointRequest(*rds.DeleteDBClusterEndpointInput) rds.DeleteDBClusterEndpointRequest

	DeleteDBClusterParameterGroupRequest(*rds.DeleteDBClusterParameterGroupInput) rds.DeleteDBClusterParameterGroupRequest

	DeleteDBClusterSnapshotRequest(*rds.DeleteDBClusterSnapshotInput) rds.DeleteDBClusterSnapshotRequest

	DeleteDBInstanceRequest(*rds.DeleteDBInstanceInput) rds.DeleteDBInstanceRequest

	DeleteDBInstanceAutomatedBackupRequest(*rds.DeleteDBInstanceAutomatedBackupInput) rds.DeleteDBInstanceAutomatedBackupRequest

	DeleteDBParameterGroupRequest(*rds.DeleteDBParameterGroupInput) rds.DeleteDBParameterGroupRequest

	DeleteDBSecurityGroupRequest(*rds.DeleteDBSecurityGroupInput) rds.DeleteDBSecurityGroupRequest

	DeleteDBSnapshotRequest(*rds.DeleteDBSnapshotInput) rds.DeleteDBSnapshotRequest

	DeleteDBSubnetGroupRequest(*rds.DeleteDBSubnetGroupInput) rds.DeleteDBSubnetGroupRequest

	DeleteEventSubscriptionRequest(*rds.DeleteEventSubscriptionInput) rds.DeleteEventSubscriptionRequest

	DeleteGlobalClusterRequest(*rds.DeleteGlobalClusterInput) rds.DeleteGlobalClusterRequest

	DeleteInstallationMediaRequest(*rds.DeleteInstallationMediaInput) rds.DeleteInstallationMediaRequest

	DeleteOptionGroupRequest(*rds.DeleteOptionGroupInput) rds.DeleteOptionGroupRequest

	DescribeAccountAttributesRequest(*rds.DescribeAccountAttributesInput) rds.DescribeAccountAttributesRequest

	DescribeCertificatesRequest(*rds.DescribeCertificatesInput) rds.DescribeCertificatesRequest

	DescribeCustomAvailabilityZonesRequest(*rds.DescribeCustomAvailabilityZonesInput) rds.DescribeCustomAvailabilityZonesRequest

	DescribeDBClusterBacktracksRequest(*rds.DescribeDBClusterBacktracksInput) rds.DescribeDBClusterBacktracksRequest

	DescribeDBClusterEndpointsRequest(*rds.DescribeDBClusterEndpointsInput) rds.DescribeDBClusterEndpointsRequest

	DescribeDBClusterParameterGroupsRequest(*rds.DescribeDBClusterParameterGroupsInput) rds.DescribeDBClusterParameterGroupsRequest

	DescribeDBClusterParametersRequest(*rds.DescribeDBClusterParametersInput) rds.DescribeDBClusterParametersRequest

	DescribeDBClusterSnapshotAttributesRequest(*rds.DescribeDBClusterSnapshotAttributesInput) rds.DescribeDBClusterSnapshotAttributesRequest

	DescribeDBClusterSnapshotsRequest(*rds.DescribeDBClusterSnapshotsInput) rds.DescribeDBClusterSnapshotsRequest

	DescribeDBClustersRequest(*rds.DescribeDBClustersInput) rds.DescribeDBClustersRequest

	DescribeDBEngineVersionsRequest(*rds.DescribeDBEngineVersionsInput) rds.DescribeDBEngineVersionsRequest

	DescribeDBInstanceAutomatedBackupsRequest(*rds.DescribeDBInstanceAutomatedBackupsInput) rds.DescribeDBInstanceAutomatedBackupsRequest

	DescribeDBInstancesRequest(*rds.DescribeDBInstancesInput) rds.DescribeDBInstancesRequest

	DescribeDBLogFilesRequest(*rds.DescribeDBLogFilesInput) rds.DescribeDBLogFilesRequest

	DescribeDBParameterGroupsRequest(*rds.DescribeDBParameterGroupsInput) rds.DescribeDBParameterGroupsRequest

	DescribeDBParametersRequest(*rds.DescribeDBParametersInput) rds.DescribeDBParametersRequest

	DescribeDBSecurityGroupsRequest(*rds.DescribeDBSecurityGroupsInput) rds.DescribeDBSecurityGroupsRequest

	DescribeDBSnapshotAttributesRequest(*rds.DescribeDBSnapshotAttributesInput) rds.DescribeDBSnapshotAttributesRequest

	DescribeDBSnapshotsRequest(*rds.DescribeDBSnapshotsInput) rds.DescribeDBSnapshotsRequest

	DescribeDBSubnetGroupsRequest(*rds.DescribeDBSubnetGroupsInput) rds.DescribeDBSubnetGroupsRequest

	DescribeEngineDefaultClusterParametersRequest(*rds.DescribeEngineDefaultClusterParametersInput) rds.DescribeEngineDefaultClusterParametersRequest

	DescribeEngineDefaultParametersRequest(*rds.DescribeEngineDefaultParametersInput) rds.DescribeEngineDefaultParametersRequest

	DescribeEventCategoriesRequest(*rds.DescribeEventCategoriesInput) rds.DescribeEventCategoriesRequest

	DescribeEventSubscriptionsRequest(*rds.DescribeEventSubscriptionsInput) rds.DescribeEventSubscriptionsRequest

	DescribeEventsRequest(*rds.DescribeEventsInput) rds.DescribeEventsRequest

	DescribeGlobalClustersRequest(*rds.DescribeGlobalClustersInput) rds.DescribeGlobalClustersRequest

	DescribeInstallationMediaRequest(*rds.DescribeInstallationMediaInput) rds.DescribeInstallationMediaRequest

	DescribeOptionGroupOptionsRequest(*rds.DescribeOptionGroupOptionsInput) rds.DescribeOptionGroupOptionsRequest

	DescribeOptionGroupsRequest(*rds.DescribeOptionGroupsInput) rds.DescribeOptionGroupsRequest

	DescribeOrderableDBInstanceOptionsRequest(*rds.DescribeOrderableDBInstanceOptionsInput) rds.DescribeOrderableDBInstanceOptionsRequest

	DescribePendingMaintenanceActionsRequest(*rds.DescribePendingMaintenanceActionsInput) rds.DescribePendingMaintenanceActionsRequest

	DescribeReservedDBInstancesRequest(*rds.DescribeReservedDBInstancesInput) rds.DescribeReservedDBInstancesRequest

	DescribeReservedDBInstancesOfferingsRequest(*rds.DescribeReservedDBInstancesOfferingsInput) rds.DescribeReservedDBInstancesOfferingsRequest

	DescribeSourceRegionsRequest(*rds.DescribeSourceRegionsInput) rds.DescribeSourceRegionsRequest

	DescribeValidDBInstanceModificationsRequest(*rds.DescribeValidDBInstanceModificationsInput) rds.DescribeValidDBInstanceModificationsRequest

	DownloadDBLogFilePortionRequest(*rds.DownloadDBLogFilePortionInput) rds.DownloadDBLogFilePortionRequest

	FailoverDBClusterRequest(*rds.FailoverDBClusterInput) rds.FailoverDBClusterRequest

	ImportInstallationMediaRequest(*rds.ImportInstallationMediaInput) rds.ImportInstallationMediaRequest

	ListTagsForResourceRequest(*rds.ListTagsForResourceInput) rds.ListTagsForResourceRequest

	ModifyCurrentDBClusterCapacityRequest(*rds.ModifyCurrentDBClusterCapacityInput) rds.ModifyCurrentDBClusterCapacityRequest

	ModifyDBClusterRequest(*rds.ModifyDBClusterInput) rds.ModifyDBClusterRequest

	ModifyDBClusterEndpointRequest(*rds.ModifyDBClusterEndpointInput) rds.ModifyDBClusterEndpointRequest

	ModifyDBClusterParameterGroupRequest(*rds.ModifyDBClusterParameterGroupInput) rds.ModifyDBClusterParameterGroupRequest

	ModifyDBClusterSnapshotAttributeRequest(*rds.ModifyDBClusterSnapshotAttributeInput) rds.ModifyDBClusterSnapshotAttributeRequest

	ModifyDBInstanceRequest(*rds.ModifyDBInstanceInput) rds.ModifyDBInstanceRequest

	ModifyDBParameterGroupRequest(*rds.ModifyDBParameterGroupInput) rds.ModifyDBParameterGroupRequest

	ModifyDBSnapshotRequest(*rds.ModifyDBSnapshotInput) rds.ModifyDBSnapshotRequest

	ModifyDBSnapshotAttributeRequest(*rds.ModifyDBSnapshotAttributeInput) rds.ModifyDBSnapshotAttributeRequest

	ModifyDBSubnetGroupRequest(*rds.ModifyDBSubnetGroupInput) rds.ModifyDBSubnetGroupRequest

	ModifyEventSubscriptionRequest(*rds.ModifyEventSubscriptionInput) rds.ModifyEventSubscriptionRequest

	ModifyGlobalClusterRequest(*rds.ModifyGlobalClusterInput) rds.ModifyGlobalClusterRequest

	ModifyOptionGroupRequest(*rds.ModifyOptionGroupInput) rds.ModifyOptionGroupRequest

	PromoteReadReplicaRequest(*rds.PromoteReadReplicaInput) rds.PromoteReadReplicaRequest

	PromoteReadReplicaDBClusterRequest(*rds.PromoteReadReplicaDBClusterInput) rds.PromoteReadReplicaDBClusterRequest

	PurchaseReservedDBInstancesOfferingRequest(*rds.PurchaseReservedDBInstancesOfferingInput) rds.PurchaseReservedDBInstancesOfferingRequest

	RebootDBInstanceRequest(*rds.RebootDBInstanceInput) rds.RebootDBInstanceRequest

	RemoveFromGlobalClusterRequest(*rds.RemoveFromGlobalClusterInput) rds.RemoveFromGlobalClusterRequest

	RemoveRoleFromDBClusterRequest(*rds.RemoveRoleFromDBClusterInput) rds.RemoveRoleFromDBClusterRequest

	RemoveRoleFromDBInstanceRequest(*rds.RemoveRoleFromDBInstanceInput) rds.RemoveRoleFromDBInstanceRequest

	RemoveSourceIdentifierFromSubscriptionRequest(*rds.RemoveSourceIdentifierFromSubscriptionInput) rds.RemoveSourceIdentifierFromSubscriptionRequest

	RemoveTagsFromResourceRequest(*rds.RemoveTagsFromResourceInput) rds.RemoveTagsFromResourceRequest

	ResetDBClusterParameterGroupRequest(*rds.ResetDBClusterParameterGroupInput) rds.ResetDBClusterParameterGroupRequest

	ResetDBParameterGroupRequest(*rds.ResetDBParameterGroupInput) rds.ResetDBParameterGroupRequest

	RestoreDBClusterFromS3Request(*rds.RestoreDBClusterFromS3Input) rds.RestoreDBClusterFromS3Request

	RestoreDBClusterFromSnapshotRequest(*rds.RestoreDBClusterFromSnapshotInput) rds.RestoreDBClusterFromSnapshotRequest

	RestoreDBClusterToPointInTimeRequest(*rds.RestoreDBClusterToPointInTimeInput) rds.RestoreDBClusterToPointInTimeRequest

	RestoreDBInstanceFromDBSnapshotRequest(*rds.RestoreDBInstanceFromDBSnapshotInput) rds.RestoreDBInstanceFromDBSnapshotRequest

	RestoreDBInstanceFromS3Request(*rds.RestoreDBInstanceFromS3Input) rds.RestoreDBInstanceFromS3Request

	RestoreDBInstanceToPointInTimeRequest(*rds.RestoreDBInstanceToPointInTimeInput) rds.RestoreDBInstanceToPointInTimeRequest

	RevokeDBSecurityGroupIngressRequest(*rds.RevokeDBSecurityGroupIngressInput) rds.RevokeDBSecurityGroupIngressRequest

	StartActivityStreamRequest(*rds.StartActivityStreamInput) rds.StartActivityStreamRequest

	StartDBClusterRequest(*rds.StartDBClusterInput) rds.StartDBClusterRequest

	StartDBInstanceRequest(*rds.StartDBInstanceInput) rds.StartDBInstanceRequest

	StopActivityStreamRequest(*rds.StopActivityStreamInput) rds.StopActivityStreamRequest

	StopDBClusterRequest(*rds.StopDBClusterInput) rds.StopDBClusterRequest

	StopDBInstanceRequest(*rds.StopDBInstanceInput) rds.StopDBInstanceRequest

	WaitUntilDBClusterSnapshotAvailable(context.Context, *rds.DescribeDBClusterSnapshotsInput, ...aws.WaiterOption) error

	WaitUntilDBClusterSnapshotDeleted(context.Context, *rds.DescribeDBClusterSnapshotsInput, ...aws.WaiterOption) error

	WaitUntilDBInstanceAvailable(context.Context, *rds.DescribeDBInstancesInput, ...aws.WaiterOption) error

	WaitUntilDBInstanceDeleted(context.Context, *rds.DescribeDBInstancesInput, ...aws.WaiterOption) error

	WaitUntilDBSnapshotAvailable(context.Context, *rds.DescribeDBSnapshotsInput, ...aws.WaiterOption) error

	WaitUntilDBSnapshotDeleted(context.Context, *rds.DescribeDBSnapshotsInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*rds.Client)(nil)
