// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Represents a single entry in a list of agents. AgentListEntry returns an
// array that contains a list of agents when the ListAgents operation is called.
type AgentListEntry struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent.
	AgentArn *string `type:"string"`

	// The name of the agent.
	Name *string `min:"1" type:"string"`

	// The status of the agent.
	Status AgentStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s AgentListEntry) String() string {
	return awsutil.Prettify(s)
}

// The subnet and the security group that DataSync uses to access target EFS
// file system. The subnet must have at least one mount target for that file
// system. The security group that you provide needs to be able to communicate
// with the security group on the mount target in the subnet specified.
type Ec2Config struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Names (ARNs) of the security groups that are configured
	// for the Amazon EC2 resource.
	//
	// SecurityGroupArns is a required field
	SecurityGroupArns []string `min:"1" type:"list" required:"true"`

	// The ARN of the subnet and the security group that DataSync uses to access
	// the target EFS file system.
	//
	// SubnetArn is a required field
	SubnetArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Ec2Config) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Ec2Config) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Ec2Config"}

	if s.SecurityGroupArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityGroupArns"))
	}
	if s.SecurityGroupArns != nil && len(s.SecurityGroupArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityGroupArns", 1))
	}

	if s.SubnetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies which files, folders and objects to include or exclude when transferring
// files from source to destination.
type FilterRule struct {
	_ struct{} `type:"structure"`

	// The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN
	// rule type.
	FilterType FilterType `type:"string" enum:"true"`

	// A single filter string that consists of the patterns to include or exclude.
	// The patterns are delimited by "|" (that is, a pipe), for example: /folder1|/folder2
	Value *string `type:"string"`
}

// String returns the string representation
func (s FilterRule) String() string {
	return awsutil.Prettify(s)
}

// Represents a single entry in a list of locations. LocationListEntry returns
// an array that contains a list of locations when the ListLocations operation
// is called.
type LocationListEntry struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the location. For Network File System (NFS)
	// or Amazon EFS, the location is the export path. For Amazon S3, the location
	// is the prefix path that you want to mount and use as the root of the location.
	LocationArn *string `type:"string"`

	// Represents a list of URLs of a location. LocationUri returns an array that
	// contains a list of locations when the ListLocations operation is called.
	//
	// Format: TYPE://GLOBAL_ID/SUBDIR.
	//
	// TYPE designates the type of location. Valid values: NFS | EFS | S3.
	//
	// GLOBAL_ID is the globally unique identifier of the resource that backs the
	// location. An example for EFS is us-east-2.fs-abcd1234. An example for Amazon
	// S3 is the bucket name, such as myBucket. An example for NFS is a valid IPv4
	// address or a host name compliant with Domain Name Service (DNS).
	//
	// SUBDIR is a valid file system path, delimited by forward slashes as is the
	// *nix convention. For NFS and Amazon EFS, it's the export path to mount the
	// location. For Amazon S3, it's the prefix path that you mount to and treat
	// as the root of the location.
	LocationUri *string `type:"string"`
}

// String returns the string representation
func (s LocationListEntry) String() string {
	return awsutil.Prettify(s)
}

// Represents the mount options that are available for DataSync to access an
// NFS location.
type NfsMountOptions struct {
	_ struct{} `type:"structure"`

	// The specific NFS version that you want DataSync to use to mount your NFS
	// share. If the server refuses to use the version specified, the sync will
	// fail. If you don't specify a version, DataSync defaults to AUTOMATIC. That
	// is, DataSync automatically selects a version based on negotiation with the
	// NFS server.
	//
	// You can specify the following NFS versions:
	//
	//    * NFSv3 (https://tools.ietf.org/html/rfc1813) - stateless protocol version
	//    that allows for asynchronous writes on the server.
	//
	//    * NFSv4.0 (https://tools.ietf.org/html/rfc3530) - stateful, firewall-friendly
	//    protocol version that supports delegations and pseudo filesystems.
	//
	//    * NFSv4.1 (https://tools.ietf.org/html/rfc5661) - stateful protocol version
	//    that supports sessions, directory delegations, and parallel data processing.
	//    Version 4.1 also includes all features available in version 4.0.
	Version NfsVersion `type:"string" enum:"true"`
}

// String returns the string representation
func (s NfsMountOptions) String() string {
	return awsutil.Prettify(s)
}

// A list of Amazon Resource Names (ARNs) of agents to use for a Network File
// System (NFS) location.
type OnPremConfig struct {
	_ struct{} `type:"structure"`

	// ARNs)of the agents to use for an NFS location.
	//
	// AgentArns is a required field
	AgentArns []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s OnPremConfig) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *OnPremConfig) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "OnPremConfig"}

	if s.AgentArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("AgentArns"))
	}
	if s.AgentArns != nil && len(s.AgentArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AgentArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the options that are available to control the behavior of a StartTaskExecution
// operation. Behavior includes preserving metadata such as user ID (UID), group
// ID (GID), and file permissions, and also overwriting files in the destination,
// data integrity verification, and so on.
//
// A task has a set of default options associated with it. If you don't specify
// an option in StartTaskExecution, the default value is used. You can override
// the defaults options on each task execution by specifying an overriding Options
// value to StartTaskExecution.
type Options struct {
	_ struct{} `type:"structure"`

	// A file metadata value that shows the last time a file was accessed (that
	// is, when the file was read or written to). If you set Atime to BEST_EFFORT,
	// DataSync attempts to preserve the original Atime attribute on all source
	// files (that is, the version before the PREPARING phase). However, Atime's
	// behavior is not fully standard across platforms, so AWS DataSync can only
	// do this on a best-effort basis.
	//
	// Default value: BEST_EFFORT.
	//
	// BEST_EFFORT: Attempt to preserve the per-file Atime value (recommended).
	//
	// NONE: Ignore Atime.
	//
	// If Atime is set to BEST_EFFORT, Mtime must be set to PRESERVE.
	//
	// If Atime is set to NONE, Mtime must also be NONE.
	Atime Atime `type:"string" enum:"true"`

	// A value that limits the bandwidth used by AWS DataSync. For example, if you
	// want AWS DataSync to use a maximum of 1 MB, set this value to 1048576 (=1024*1024).
	BytesPerSecond *int64 `type:"long"`

	// The group ID (GID) of the file's owners.
	//
	// Default value: INT_VALUE. This preserves the integer value of the ID.
	//
	// INT_VALUE: Preserve the integer value of user ID (UID) and GID (recommended).
	//
	// NONE: Ignore UID and GID.
	Gid Gid `type:"string" enum:"true"`

	// A value that indicates the last time that a file was modified (that is, a
	// file was written to) before the PREPARING phase.
	//
	// Default value: PRESERVE.
	//
	// PRESERVE: Preserve original Mtime (recommended)
	//
	// NONE: Ignore Mtime.
	//
	// If Mtime is set to PRESERVE, Atime must be set to BEST_EFFORT.
	//
	// If Mtime is set to NONE, Atime must also be set to NONE.
	Mtime Mtime `type:"string" enum:"true"`

	// A value that determines whether files at the destination should be overwritten
	// or preserved when copying files. If set to NEVER a destination file will
	// not be replaced by a source file, even if the destination file differs from
	// the source file. If you modify files in the destination and you sync the
	// files, you can use this value to protect against overwriting those changes.
	//
	// Some storage classes have specific behaviors that can affect your S3 storage
	// cost. For detailed information, see using-storage-classes in the AWS DataSync
	// User Guide.
	OverwriteMode OverwriteMode `type:"string" enum:"true"`

	// A value that determines which users or groups can access a file for a specific
	// purpose such as reading, writing, or execution of the file.
	//
	// Default value: PRESERVE.
	//
	// PRESERVE: Preserve POSIX-style permissions (recommended).
	//
	// NONE: Ignore permissions.
	//
	// AWS DataSync can preserve extant permissions of a source location.
	PosixPermissions PosixPermissions `type:"string" enum:"true"`

	// A value that specifies whether files in the destination that don't exist
	// in the source file system should be preserved. This option can affect your
	// storage cost. If your task deletes objects, you might incur minimum storage
	// duration charges for certain storage classes. For detailed information, see
	// using-storage-classes in the AWS DataSync User Guide.
	//
	// Default value: PRESERVE.
	//
	// PRESERVE: Ignore such destination files (recommended).
	//
	// REMOVE: Delete destination files that aren’t present in the source.
	PreserveDeletedFiles PreserveDeletedFiles `type:"string" enum:"true"`

	// A value that determines whether AWS DataSync should preserve the metadata
	// of block and character devices in the source file system, and recreate the
	// files with that device name and metadata on the destination.
	//
	// AWS DataSync can't sync the actual contents of such devices, because they
	// are nonterminal and don't return an end-of-file (EOF) marker.
	//
	// Default value: NONE.
	//
	// NONE: Ignore special devices (recommended).
	//
	// PRESERVE: Preserve character and block device metadata. This option isn't
	// currently supported for Amazon EFS.
	PreserveDevices PreserveDevices `type:"string" enum:"true"`

	// A value that determines whether tasks should be queued before executing the
	// tasks. If set to Enabled, the tasks will queued. The default is Enabled.
	//
	// If you use the same agent to run multiple tasks you can enable the tasks
	// to run in series. For more information see task-queue.
	TaskQueueing TaskQueueing `type:"string" enum:"true"`

	// The user ID (UID) of the file's owner.
	//
	// Default value: INT_VALUE. This preserves the integer value of the ID.
	//
	// INT_VALUE: Preserve the integer value of UID and group ID (GID) (recommended).
	//
	// NONE: Ignore UID and GID.
	Uid Uid `type:"string" enum:"true"`

	// A value that determines whether a data integrity verification should be performed
	// at the end of a task execution after all data and metadata have been transferred.
	//
	// Default value: POINT_IN_TIME_CONSISTENT.
	//
	// POINT_IN_TIME_CONSISTENT: Perform verification (recommended).
	//
	// ONLY_FILES_TRANSFERRED: Perform verification on only files that were transferred.
	//
	// NONE: Skip verification.
	VerifyMode VerifyMode `type:"string" enum:"true"`
}

// String returns the string representation
func (s Options) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Options) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Options"}
	if s.BytesPerSecond != nil && *s.BytesPerSecond < -1 {
		invalidParams.Add(aws.NewErrParamMinValue("BytesPerSecond", -1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The VPC endpoint, subnet and security group that an agent uses to access
// IP addresses in a VPC (Virtual Private Cloud).
type PrivateLinkConfig struct {
	_ struct{} `type:"structure"`

	// The private endpoint that is configured for an agent that has access to IP
	// addresses in a PrivateLink (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-service.html).
	// An agent that is configured with this endpoint will not be accessible over
	// the public Internet.
	PrivateLinkEndpoint *string `min:"7" type:"string"`

	// The Amazon Resource Names (ARNs) of the security groups that are configured
	// for the EC2 resource that hosts an agent activated in a VPC or an agent that
	// has access to a VPC endpoint.
	SecurityGroupArns []string `min:"1" type:"list"`

	// The Amazon Resource Names (ARNs) of the subnets that are configured for an
	// agent activated in a VPC or an agent that has access to a VPC endpoint.
	SubnetArns []string `min:"1" type:"list"`

	// The ID of the VPC endpoint that is configured for an agent. An agent that
	// is configured with a VPC endpoint will not be accessible over the public
	// Internet.
	VpcEndpointId *string `type:"string"`
}

// String returns the string representation
func (s PrivateLinkConfig) String() string {
	return awsutil.Prettify(s)
}

// The Amazon Resource Name (ARN) of the AWS Identity and Access Management
// (IAM) role that is used to access an Amazon S3 bucket.
//
// For detailed information about using such a role, see Creating a Location
// for Amazon S3 in the AWS DataSync User Guide.
type S3Config struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 bucket to access. This bucket is used as a parameter in the
	// CreateLocationS3 operation.
	//
	// BucketAccessRoleArn is a required field
	BucketAccessRoleArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s S3Config) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *S3Config) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "S3Config"}

	if s.BucketAccessRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("BucketAccessRoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the mount options that are available for DataSync to access an
// SMB location.
type SmbMountOptions struct {
	_ struct{} `type:"structure"`

	// The specific SMB version that you want DataSync to use to mount your SMB
	// share. If you don't specify a version, DataSync defaults to AUTOMATIC. That
	// is, DataSync automatically selects a version based on negotiation with the
	// SMB server.
	Version SmbVersion `type:"string" enum:"true"`
}

// String returns the string representation
func (s SmbMountOptions) String() string {
	return awsutil.Prettify(s)
}

// Represents a single entry in a list of AWS resource tags. TagListEntry returns
// an array that contains a list of tasks when the ListTagsForResource operation
// is called.
type TagListEntry struct {
	_ struct{} `type:"structure"`

	// The key for an AWS resource tag.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The value for an AWS resource tag.
	Value *string `min:"1" type:"string"`
}

// String returns the string representation
func (s TagListEntry) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagListEntry) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagListEntry"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}
	if s.Value != nil && len(*s.Value) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Value", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a single entry in a list of task executions. TaskExecutionListEntry
// returns an array that contains a list of specific invocations of a task when
// ListTaskExecutions operation is called.
type TaskExecutionListEntry struct {
	_ struct{} `type:"structure"`

	// The status of a task execution.
	Status TaskExecutionStatus `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the task that was executed.
	TaskExecutionArn *string `type:"string"`
}

// String returns the string representation
func (s TaskExecutionListEntry) String() string {
	return awsutil.Prettify(s)
}

// Describes the detailed result of a TaskExecution operation. This result includes
// the time in milliseconds spent in each phase, the status of the task execution,
// and the errors encountered.
type TaskExecutionResultDetail struct {
	_ struct{} `type:"structure"`

	// Errors that AWS DataSync encountered during execution of the task. You can
	// use this error code to help troubleshoot issues.
	ErrorCode *string `type:"string"`

	// Detailed description of an error that was encountered during the task execution.
	// You can use this information to help troubleshoot issues.
	ErrorDetail *string `type:"string"`

	// The total time in milliseconds that AWS DataSync spent in the PREPARING phase.
	PrepareDuration *int64 `type:"long"`

	// The status of the PREPARING phase.
	PrepareStatus PhaseStatus `type:"string" enum:"true"`

	// The total time in milliseconds that AWS DataSync spent in the TRANSFERRING
	// phase.
	TransferDuration *int64 `type:"long"`

	// The status of the TRANSFERRING Phase.
	TransferStatus PhaseStatus `type:"string" enum:"true"`

	// The total time in milliseconds that AWS DataSync spent in the VERIFYING phase.
	VerifyDuration *int64 `type:"long"`

	// The status of the VERIFYING Phase.
	VerifyStatus PhaseStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s TaskExecutionResultDetail) String() string {
	return awsutil.Prettify(s)
}

// Represents a single entry in a list of tasks. TaskListEntry returns an array
// that contains a list of tasks when the ListTasks operation is called. A task
// includes the source and destination file systems to sync and the options
// to use for the tasks.
type TaskListEntry struct {
	_ struct{} `type:"structure"`

	// The name of the task.
	Name *string `min:"1" type:"string"`

	// The status of the task.
	Status TaskStatus `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string `type:"string"`
}

// String returns the string representation
func (s TaskListEntry) String() string {
	return awsutil.Prettify(s)
}
