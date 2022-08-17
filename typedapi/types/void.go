// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


package types

// Void type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/_spec_utils/VoidValue.ts#L20-L28
type Void struct{}

// VoidBuilder holds Void struct and provides a builder API.
type VoidBuilder struct {
	v Void
}

// NewVoid provides a builder for the Void struct.
func NewVoidBuilder() *VoidBuilder {
	return &VoidBuilder{}
}

// Build finalize the chain and returns the Void struct
func (b *VoidBuilder) Build() Void {
	return b.v
}

func (b *VoidBuilder) Void(value Void) *VoidBuilder {
	b.v = value
	return b
}
