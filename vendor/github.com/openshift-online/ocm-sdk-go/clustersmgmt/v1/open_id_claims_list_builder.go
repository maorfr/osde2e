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

// OpenIDClaimsListBuilder contains the data and logic needed to build
// 'open_ID_claims' objects.
type OpenIDClaimsListBuilder struct {
	items []*OpenIDClaimsBuilder
}

// NewOpenIDClaimsList creates a new builder of 'open_ID_claims' objects.
func NewOpenIDClaimsList() *OpenIDClaimsListBuilder {
	return new(OpenIDClaimsListBuilder)
}

// Items sets the items of the list.
func (b *OpenIDClaimsListBuilder) Items(values ...*OpenIDClaimsBuilder) *OpenIDClaimsListBuilder {
	b.items = make([]*OpenIDClaimsBuilder, len(values))
	copy(b.items, values)
	return b
}

// Copy copies the items of the given list into this builder, discarding any previous items.
func (b *OpenIDClaimsListBuilder) Copy(list *OpenIDClaimsList) *OpenIDClaimsListBuilder {
	if list == nil || list.items == nil {
		b.items = nil
	} else {
		b.items = make([]*OpenIDClaimsBuilder, len(list.items))
		for i, v := range list.items {
			b.items[i] = NewOpenIDClaims().Copy(v)
		}
	}
	return b
}

// Build creates a list of 'open_ID_claims' objects using the
// configuration stored in the builder.
func (b *OpenIDClaimsListBuilder) Build() (list *OpenIDClaimsList, err error) {
	items := make([]*OpenIDClaims, len(b.items))
	for i, item := range b.items {
		items[i], err = item.Build()
		if err != nil {
			return
		}
	}
	list = new(OpenIDClaimsList)
	list.items = items
	return
}
