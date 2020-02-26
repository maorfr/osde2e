/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterRegistration represents the values of the 'cluster_registration' type.
//
// Registration of a new cluster to the service.
type ClusterRegistration struct {
	externalID     *string
	subscriptionID *string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterRegistration) Empty() bool {
	return o == nil || (o.externalID == nil &&
		o.subscriptionID == nil &&
		true)
}

// ExternalID returns the value of the 'external_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Identifier of the cluster generated by the installer.
//
// For example, to register a cluster that has been provisioned outside
// of this service, send a a request like this:
//
// [source,http]
// ----
// POST /api/clusters_mgmt/v1/register_cluster HTTP/1.1
// ----
//
// With a request body like this:
//
// [source,json]
// ----
// {
//   "external_id": "d656aecf-11a6-4782-ad86-8f72638449ba"
// }
// ----
func (o *ClusterRegistration) ExternalID() string {
	if o != nil && o.externalID != nil {
		return *o.externalID
	}
	return ""
}

// GetExternalID returns the value of the 'external_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Identifier of the cluster generated by the installer.
//
// For example, to register a cluster that has been provisioned outside
// of this service, send a a request like this:
//
// [source,http]
// ----
// POST /api/clusters_mgmt/v1/register_cluster HTTP/1.1
// ----
//
// With a request body like this:
//
// [source,json]
// ----
// {
//   "external_id": "d656aecf-11a6-4782-ad86-8f72638449ba"
// }
// ----
func (o *ClusterRegistration) GetExternalID() (value string, ok bool) {
	ok = o != nil && o.externalID != nil
	if ok {
		value = *o.externalID
	}
	return
}

// SubscriptionID returns the value of the 'subscription_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Subscription identifier of the cluster generated by the account
// manager.
func (o *ClusterRegistration) SubscriptionID() string {
	if o != nil && o.subscriptionID != nil {
		return *o.subscriptionID
	}
	return ""
}

// GetSubscriptionID returns the value of the 'subscription_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Subscription identifier of the cluster generated by the account
// manager.
func (o *ClusterRegistration) GetSubscriptionID() (value string, ok bool) {
	ok = o != nil && o.subscriptionID != nil
	if ok {
		value = *o.subscriptionID
	}
	return
}

// ClusterRegistrationListKind is the name of the type used to represent list of objects of
// type 'cluster_registration'.
const ClusterRegistrationListKind = "ClusterRegistrationList"

// ClusterRegistrationListLinkKind is the name of the type used to represent links to list
// of objects of type 'cluster_registration'.
const ClusterRegistrationListLinkKind = "ClusterRegistrationListLink"

// ClusterRegistrationNilKind is the name of the type used to nil lists of objects of
// type 'cluster_registration'.
const ClusterRegistrationListNilKind = "ClusterRegistrationListNil"

// ClusterRegistrationList is a list of values of the 'cluster_registration' type.
type ClusterRegistrationList struct {
	href  *string
	link  bool
	items []*ClusterRegistration
}

// Len returns the length of the list.
func (l *ClusterRegistrationList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterRegistrationList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterRegistrationList) Get(i int) *ClusterRegistration {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterRegistrationList) Slice() []*ClusterRegistration {
	var slice []*ClusterRegistration
	if l == nil {
		slice = make([]*ClusterRegistration, 0)
	} else {
		slice = make([]*ClusterRegistration, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterRegistrationList) Each(f func(item *ClusterRegistration) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ClusterRegistrationList) Range(f func(index int, item *ClusterRegistration) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
