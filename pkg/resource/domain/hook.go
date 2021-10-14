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

package domain

import (
	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
	svccommon "github.com/aws-controllers-k8s/sagemaker-controller/pkg/common"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

var (
	modifyingStatuses = []string{
		svcsdk.DomainStatusPending,
		svcsdk.DomainStatusUpdating,
		svcsdk.DomainStatusDeleting,
	}

	resourceName = resourceGK.Kind
)

// customDescribeDomainSetOutput sets the resource ResourceSynced condition to False if endpoint is
// being modified by AWS
func (rm *resourceManager) customDescribeDomainSetOutput(ko *svcapitypes.Domain) {
	svccommon.SetSyncedCondition(&resource{ko}, ko.Status.DomainStatus, &resourceName, &modifyingStatuses)
}
