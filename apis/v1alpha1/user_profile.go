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

// UserProfileSpec defines the desired state of UserProfile.
type UserProfileSpec struct {
	// The ID of the associated Domain.
	// +kubebuilder:validation:Required
	DomainID *string `json:"domainID"`
	// A specifier for the type of value specified in SingleSignOnUserValue. Currently,
	// the only supported value is "UserName". If the Domain's AuthMode is SSO,
	// this field is required. If the Domain's AuthMode is not SSO, this field cannot
	// be specified.
	SingleSignOnUserIdentifier *string `json:"singleSignOnUserIdentifier,omitempty"`
	// The username of the associated AWS Single Sign-On User for this UserProfile.
	// If the Domain's AuthMode is SSO, this field is required, and must match a
	// valid username of a user in your directory. If the Domain's AuthMode is not
	// SSO, this field cannot be specified.
	SingleSignOnUserValue *string `json:"singleSignOnUserValue,omitempty"`
	// Each tag consists of a key and an optional value. Tag keys must be unique
	// per resource.
	Tags []*Tag `json:"tags,omitempty"`
	// A name for the UserProfile.
	// +kubebuilder:validation:Required
	UserProfileName *string `json:"userProfileName"`
	// A collection of settings.
	UserSettings *UserSettings `json:"userSettings,omitempty"`
}

// UserProfileStatus defines the observed state of UserProfile
type UserProfileStatus struct {
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
}

// UserProfile is the Schema for the UserProfiles API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type UserProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserProfileSpec   `json:"spec,omitempty"`
	Status            UserProfileStatus `json:"status,omitempty"`
}

// UserProfileList contains a list of UserProfile
// +kubebuilder:object:root=true
type UserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UserProfile{}, &UserProfileList{})
}
