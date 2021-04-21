// Copyright 2020 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package okteto

import (
	"context"
	"fmt"
)

// CreateBody top body answer
type CreateBody struct {
	Namespace Namespace `json:"createSpace" yaml:"createSpace"`
}

// DeleteBody top body answer
type DeleteBody struct {
	Namespace Namespace `json:"deleteSpace" yaml:"deleteSpace"`
}

//Spaces represents an Okteto list of spaces
type Spaces struct {
	Spaces []Namespace `json:"spaces" yaml:"spaces"`
}

//Namespace represents an Okteto k8s namespace
type Namespace struct {
	ID string `json:"id" yaml:"id"`
}

// CreateNamespace creates a namespace
func CreateNamespace(ctx context.Context, namespace string) (string, error) {
	q := fmt.Sprintf(`mutation{
		createSpace(name: "%s"){
			id
		},
	}`, namespace)

	var body CreateBody
	if err := query(ctx, q, &body); err != nil {
		return "", err
	}

	return body.Namespace.ID, nil
}

// ListNamespaces list namespaces
func ListNamespaces(ctx context.Context) ([]Namespace, error) {
	q := `query{
		spaces{
			id
		},
	}`

	var body Spaces
	if err := query(ctx, q, &body); err != nil {
		return nil, err
	}

	return body.Spaces, nil
}

// AddNamespaceMembers adds members to a namespace
func AddNamespaceMembers(ctx context.Context, namespace string, members []string) error {
	m := membersToString(members)

	q := fmt.Sprintf(`mutation{
		updateSpace(id: "%s", members: [%s]){
			id
		},
	}`, namespace, m)

	var body CreateBody
	return query(ctx, q, &body)
}

func membersToString(members []string) string {
	m := ""
	for _, mm := range members {
		if len(m) > 0 {
			m += ","
		}

		m += fmt.Sprintf(`"%s"`, mm)
	}

	return m
}

// DeleteNamespace deletes a namespace
func DeleteNamespace(ctx context.Context, namespace string) error {
	q := fmt.Sprintf(`mutation{
		deleteSpace(id: "%s"){
			id
		},
	}`, namespace)

	var body DeleteBody
	return query(ctx, q, &body)
}
