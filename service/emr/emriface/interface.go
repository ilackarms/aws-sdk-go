// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package emriface provides an interface for the Amazon Elastic MapReduce.
package emriface

import (
	"github.com/djannot/aws-sdk-go/aws/request"
	"github.com/djannot/aws-sdk-go/service/emr"
)

// EMRAPI is the interface type for emr.EMR.
type EMRAPI interface {
	AddInstanceGroupsRequest(*emr.AddInstanceGroupsInput) (*request.Request, *emr.AddInstanceGroupsOutput)

	AddInstanceGroups(*emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error)

	AddJobFlowStepsRequest(*emr.AddJobFlowStepsInput) (*request.Request, *emr.AddJobFlowStepsOutput)

	AddJobFlowSteps(*emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error)

	AddTagsRequest(*emr.AddTagsInput) (*request.Request, *emr.AddTagsOutput)

	AddTags(*emr.AddTagsInput) (*emr.AddTagsOutput, error)

	DescribeClusterRequest(*emr.DescribeClusterInput) (*request.Request, *emr.DescribeClusterOutput)

	DescribeCluster(*emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error)

	DescribeJobFlowsRequest(*emr.DescribeJobFlowsInput) (*request.Request, *emr.DescribeJobFlowsOutput)

	DescribeJobFlows(*emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error)

	DescribeStepRequest(*emr.DescribeStepInput) (*request.Request, *emr.DescribeStepOutput)

	DescribeStep(*emr.DescribeStepInput) (*emr.DescribeStepOutput, error)

	ListBootstrapActionsRequest(*emr.ListBootstrapActionsInput) (*request.Request, *emr.ListBootstrapActionsOutput)

	ListBootstrapActions(*emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error)

	ListBootstrapActionsPages(*emr.ListBootstrapActionsInput, func(*emr.ListBootstrapActionsOutput, bool) bool) error

	ListClustersRequest(*emr.ListClustersInput) (*request.Request, *emr.ListClustersOutput)

	ListClusters(*emr.ListClustersInput) (*emr.ListClustersOutput, error)

	ListClustersPages(*emr.ListClustersInput, func(*emr.ListClustersOutput, bool) bool) error

	ListInstanceGroupsRequest(*emr.ListInstanceGroupsInput) (*request.Request, *emr.ListInstanceGroupsOutput)

	ListInstanceGroups(*emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error)

	ListInstanceGroupsPages(*emr.ListInstanceGroupsInput, func(*emr.ListInstanceGroupsOutput, bool) bool) error

	ListInstancesRequest(*emr.ListInstancesInput) (*request.Request, *emr.ListInstancesOutput)

	ListInstances(*emr.ListInstancesInput) (*emr.ListInstancesOutput, error)

	ListInstancesPages(*emr.ListInstancesInput, func(*emr.ListInstancesOutput, bool) bool) error

	ListStepsRequest(*emr.ListStepsInput) (*request.Request, *emr.ListStepsOutput)

	ListSteps(*emr.ListStepsInput) (*emr.ListStepsOutput, error)

	ListStepsPages(*emr.ListStepsInput, func(*emr.ListStepsOutput, bool) bool) error

	ModifyInstanceGroupsRequest(*emr.ModifyInstanceGroupsInput) (*request.Request, *emr.ModifyInstanceGroupsOutput)

	ModifyInstanceGroups(*emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error)

	RemoveTagsRequest(*emr.RemoveTagsInput) (*request.Request, *emr.RemoveTagsOutput)

	RemoveTags(*emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error)

	RunJobFlowRequest(*emr.RunJobFlowInput) (*request.Request, *emr.RunJobFlowOutput)

	RunJobFlow(*emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error)

	SetTerminationProtectionRequest(*emr.SetTerminationProtectionInput) (*request.Request, *emr.SetTerminationProtectionOutput)

	SetTerminationProtection(*emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error)

	SetVisibleToAllUsersRequest(*emr.SetVisibleToAllUsersInput) (*request.Request, *emr.SetVisibleToAllUsersOutput)

	SetVisibleToAllUsers(*emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error)

	TerminateJobFlowsRequest(*emr.TerminateJobFlowsInput) (*request.Request, *emr.TerminateJobFlowsOutput)

	TerminateJobFlows(*emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error)
}

var _ EMRAPI = (*emr.EMR)(nil)
