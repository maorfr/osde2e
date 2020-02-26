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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// AccessTokenListBuilder contains the data and logic needed to build
// 'access_token' objects.
type AccessTokenListBuilder struct {
	items []*AccessTokenBuilder
}

// NewAccessTokenList creates a new builder of 'access_token' objects.
func NewAccessTokenList() *AccessTokenListBuilder {
	return new(AccessTokenListBuilder)
}

// Items sets the items of the list.
func (b *AccessTokenListBuilder) Items(values ...*AccessTokenBuilder) *AccessTokenListBuilder {
	b.items = make([]*AccessTokenBuilder, len(values))
	copy(b.items, values)
	return b
}

// Copy copies the items of the given list into this builder, discarding any previous items.
func (b *AccessTokenListBuilder) Copy(list *AccessTokenList) *AccessTokenListBuilder {
	if list == nil || list.items == nil {
		b.items = nil
	} else {
		b.items = make([]*AccessTokenBuilder, len(list.items))
		for i, v := range list.items {
			b.items[i] = NewAccessToken().Copy(v)
		}
	}
	return b
}

// Build creates a list of 'access_token' objects using the
// configuration stored in the builder.
func (b *AccessTokenListBuilder) Build() (list *AccessTokenList, err error) {
	items := make([]*AccessToken, len(b.items))
	for i, item := range b.items {
		items[i], err = item.Build()
		if err != nil {
			return
		}
	}
	list = new(AccessTokenList)
	list.items = items
	return
}
