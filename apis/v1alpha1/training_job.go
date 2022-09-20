// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TrainingJobSpec defines the desired state of TrainingJob.
//
// Contains information about a training job.
type TrainingJobSpec struct {
	// The registry path of the Docker image that contains the training algorithm
	// and algorithm-specific metadata, including the input mode. For more information
	// about algorithms provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).
	// For information about providing your own algorithms, see Using Your Own Algorithms
	// with Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).
	// +kubebuilder:validation:Required
	AlgorithmSpecification *AlgorithmSpecification `json:"algorithmSpecification"`
	// Contains information about the output location for managed spot training
	// checkpoint data.
	CheckpointConfig *CheckpointConfig `json:"checkpointConfig,omitempty"`

	DebugHookConfig *DebugHookConfig `json:"debugHookConfig,omitempty"`
	// Configuration information for Debugger rules for debugging output tensors.
	DebugRuleConfigurations []*DebugRuleConfiguration `json:"debugRuleConfigurations,omitempty"`
	// To encrypt all communications between ML compute instances in distributed
	// training, choose True. Encryption provides greater security for distributed
	// training, but training might take longer. How long it takes depends on the
	// amount of communication between compute instances, especially if you use
	// a deep learning algorithm in distributed training. For more information,
	// see Protect Communications Between ML Compute Instances in a Distributed
	// Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html).
	EnableInterContainerTrafficEncryption *bool `json:"enableInterContainerTrafficEncryption,omitempty"`
	// To train models using managed spot training, choose True. Managed spot training
	// provides a fully managed and scalable infrastructure for training machine
	// learning models. this option is useful when training jobs can be interrupted
	// and when there is flexibility when the training job is run.
	//
	// The complete and intermediate results of jobs are stored in an Amazon S3
	// bucket, and can be used as a starting point to train models incrementally.
	// Amazon SageMaker provides metrics and logs in CloudWatch. They can be used
	// to see when managed spot training jobs are running, interrupted, resumed,
	// or completed.
	EnableManagedSpotTraining *bool `json:"enableManagedSpotTraining,omitempty"`
	// Isolates the training container. No inbound or outbound network calls can
	// be made, except for calls between peers within a training cluster for distributed
	// training. If you enable network isolation for training jobs that are configured
	// to use a VPC, SageMaker downloads and uploads customer data and model artifacts
	// through the specified VPC, but the training container does not have network
	// access.
	EnableNetworkIsolation *bool `json:"enableNetworkIsolation,omitempty"`
	// The environment variables to set in the Docker container.
	Environment map[string]*string `json:"environment,omitempty"`

	ExperimentConfig *ExperimentConfig `json:"experimentConfig,omitempty"`
	// Algorithm-specific parameters that influence the quality of the model. You
	// set hyperparameters before you start the learning process. For a list of
	// hyperparameters for each training algorithm provided by SageMaker, see Algorithms
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).
	//
	// You can specify a maximum of 100 hyperparameters. Each hyperparameter is
	// a key-value pair. Each key and value is limited to 256 characters, as specified
	// by the Length Constraint.
	HyperParameters map[string]*string `json:"hyperParameters,omitempty"`
	// An array of Channel objects. Each channel is a named input source. InputDataConfig
	// describes the input data and its location.
	//
	// Algorithms can accept input data from one or more channels. For example,
	// an algorithm might have two channels of input data, training_data and validation_data.
	// The configuration for each channel provides the S3, EFS, or FSx location
	// where the input data is stored. It also provides information about the stored
	// data: the MIME type, compression method, and whether the data is wrapped
	// in RecordIO format.
	//
	// Depending on the input mode that the algorithm supports, SageMaker either
	// copies input data files from an S3 bucket to a local directory in the Docker
	// container, or makes it available as input streams. For example, if you specify
	// an EFS location, input data files are available as input streams. They do
	// not need to be downloaded.
	InputDataConfig []*Channel `json:"inputDataConfig,omitempty"`
	// Specifies the path to the S3 location where you want to store model artifacts.
	// SageMaker creates subfolders for the artifacts.
	// +kubebuilder:validation:Required
	OutputDataConfig *OutputDataConfig `json:"outputDataConfig"`

	ProfilerConfig *ProfilerConfig `json:"profilerConfig,omitempty"`
	// Configuration information for Debugger rules for profiling system and framework
	// metrics.
	ProfilerRuleConfigurations []*ProfilerRuleConfiguration `json:"profilerRuleConfigurations,omitempty"`
	// The resources, including the ML compute instances and ML storage volumes,
	// to use for model training.
	//
	// ML storage volumes store model artifacts and incremental states. Training
	// algorithms might also use ML storage volumes for scratch space. If you want
	// SageMaker to use the ML storage volume to store the training data, choose
	// File as the TrainingInputMode in the algorithm specification. For distributed
	// training algorithms, specify an instance count greater than 1.
	// +kubebuilder:validation:Required
	ResourceConfig *ResourceConfig `json:"resourceConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to
	// perform tasks on your behalf.
	//
	// During model training, SageMaker needs your permission to read input data
	// from an S3 bucket, download a Docker image that contains training code, write
	// model artifacts to an S3 bucket, write logs to Amazon CloudWatch Logs, and
	// publish metrics to Amazon CloudWatch. You grant permissions for all of these
	// tasks to an IAM role. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).
	//
	// To be able to pass this role to SageMaker, the caller of this API must have
	// the iam:PassRole permission.
	// +kubebuilder:validation:Required
	RoleARN *string `json:"roleARN"`
	// Specifies a limit to how long a model training job can run. It also specifies
	// how long a managed Spot training job has to complete. When the job reaches
	// the time limit, SageMaker ends the training job. Use this API to cap model
	// training costs.
	//
	// To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delays
	// job termination for 120 seconds. Algorithms can use this 120-second window
	// to save the model artifacts, so the results of training are not lost.
	// +kubebuilder:validation:Required
	StoppingCondition *StoppingCondition `json:"stoppingCondition"`
	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags []*Tag `json:"tags,omitempty"`

	TensorBoardOutputConfig *TensorBoardOutputConfig `json:"tensorBoardOutputConfig,omitempty"`
	// The name of the training job. The name must be unique within an Amazon Web
	// Services Region in an Amazon Web Services account.
	// +kubebuilder:validation:Required
	TrainingJobName *string `json:"trainingJobName"`
	// A VpcConfig object that specifies the VPC that you want your training job
	// to connect to. Control access to and from your training container by configuring
	// the VPC. For more information, see Protect Training Jobs by Using an Amazon
	// Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).
	VPCConfig *VPCConfig `json:"vpcConfig,omitempty"`
}

// TrainingJobStatus defines the observed state of TrainingJob
type TrainingJobStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Evaluation status of Debugger rules for debugging on a training job.
	// +kubebuilder:validation:Optional
	DebugRuleEvaluationStatuses []*DebugRuleEvaluationStatus `json:"debugRuleEvaluationStatuses,omitempty"`
	// If the training job failed, the reason it failed.
	// +kubebuilder:validation:Optional
	FailureReason *string `json:"failureReason,omitempty"`
	// Information about the Amazon S3 location that is configured for storing model
	// artifacts.
	// +kubebuilder:validation:Optional
	ModelArtifacts *ModelArtifacts `json:"modelArtifacts,omitempty"`
	// Evaluation status of Debugger rules for profiling on a training job.
	// +kubebuilder:validation:Optional
	ProfilerRuleEvaluationStatuses []*ProfilerRuleEvaluationStatus `json:"profilerRuleEvaluationStatuses,omitempty"`
	// Provides detailed information about the state of the training job. For detailed
	// information on the secondary status of the training job, see StatusMessage
	// under SecondaryStatusTransition.
	//
	// SageMaker provides primary statuses and secondary statuses that apply to
	// each of them:
	//
	// InProgress
	//
	//    * Starting - Starting the training job.
	//
	//    * Downloading - An optional stage for algorithms that support File training
	//    input mode. It indicates that data is being downloaded to the ML storage
	//    volumes.
	//
	//    * Training - Training is in progress.
	//
	//    * Interrupted - The job stopped because the managed spot training instances
	//    were interrupted.
	//
	//    * Uploading - Training is complete and the model artifacts are being uploaded
	//    to the S3 location.
	//
	// Completed
	//
	//    * Completed - The training job has completed.
	//
	// Failed
	//
	//    * Failed - The training job has failed. The reason for the failure is
	//    returned in the FailureReason field of DescribeTrainingJobResponse.
	//
	// Stopped
	//
	//    * MaxRuntimeExceeded - The job stopped because it exceeded the maximum
	//    allowed runtime.
	//
	//    * MaxWaitTimeExceeded - The job stopped because it exceeded the maximum
	//    allowed wait time.
	//
	//    * Stopped - The training job has stopped.
	//
	// Stopping
	//
	//    * Stopping - Stopping the training job.
	//
	// Valid values for SecondaryStatus are subject to change.
	//
	// We no longer support the following secondary statuses:
	//
	//    * LaunchingMLInstances
	//
	//    * PreparingTraining
	//
	//    * DownloadingTrainingImage
	// +kubebuilder:validation:Optional
	SecondaryStatus *string `json:"secondaryStatus,omitempty"`
	// The status of the training job.
	//
	// SageMaker provides the following training job statuses:
	//
	//    * InProgress - The training is in progress.
	//
	//    * Completed - The training job has completed.
	//
	//    * Failed - The training job has failed. To see the reason for the failure,
	//    see the FailureReason field in the response to a DescribeTrainingJobResponse
	//    call.
	//
	//    * Stopping - The training job is stopping.
	//
	//    * Stopped - The training job has stopped.
	//
	// For more detailed information, see SecondaryStatus.
	// +kubebuilder:validation:Optional
	TrainingJobStatus *string `json:"trainingJobStatus,omitempty"`
}

// TrainingJob is the Schema for the TrainingJobs API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FAILURE-REASON",type=string,priority=1,JSONPath=`.status.failureReason`
// +kubebuilder:printcolumn:name="SECONDARY-STATUS",type=string,priority=0,JSONPath=`.status.secondaryStatus`
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.trainingJobStatus`
type TrainingJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrainingJobSpec   `json:"spec,omitempty"`
	Status            TrainingJobStatus `json:"status,omitempty"`
}

// TrainingJobList contains a list of TrainingJob
// +kubebuilder:object:root=true
type TrainingJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrainingJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrainingJob{}, &TrainingJobList{})
}
