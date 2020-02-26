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

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalGithubIdentityProvider writes a value of the 'github_identity_provider' type to the given writer.
func MarshalGithubIdentityProvider(object *GithubIdentityProvider, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeGithubIdentityProvider(object, stream)
	stream.Flush()
	return stream.Error
}

// writeGithubIdentityProvider writes a value of the 'github_identity_provider' type to the given stream.
func writeGithubIdentityProvider(object *GithubIdentityProvider, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.ca != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("ca")
		stream.WriteString(*object.ca)
		count++
	}
	if object.clientID != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("client_id")
		stream.WriteString(*object.clientID)
		count++
	}
	if object.hostname != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("hostname")
		stream.WriteString(*object.hostname)
		count++
	}
	if object.teams != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("teams")
		writeStringList(object.teams, stream)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalGithubIdentityProvider reads a value of the 'github_identity_provider' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalGithubIdentityProvider(source interface{}) (object *GithubIdentityProvider, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readGithubIdentityProvider(iterator)
	err = iterator.Error
	return
}

// readGithubIdentityProvider reads a value of the 'github_identity_provider' type from the given iterator.
func readGithubIdentityProvider(iterator *jsoniter.Iterator) *GithubIdentityProvider {
	object := &GithubIdentityProvider{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "ca":
			value := iterator.ReadString()
			object.ca = &value
		case "client_id":
			value := iterator.ReadString()
			object.clientID = &value
		case "hostname":
			value := iterator.ReadString()
			object.hostname = &value
		case "teams":
			value := readStringList(iterator)
			object.teams = value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
