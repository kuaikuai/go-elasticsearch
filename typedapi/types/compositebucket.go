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

import (
	"encoding/json"
)

// CompositeBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/_types/aggregations/Aggregate.ts#L594-L596
type CompositeBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"-"`
	DocCount     int64                       `json:"doc_count"`
	Key          map[string]interface{}      `json:"key"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s CompositeBucket) MarshalJSON() ([]byte, error) {
	type opt CompositeBucket
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Aggregations {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CompositeBucketBuilder holds CompositeBucket struct and provides a builder API.
type CompositeBucketBuilder struct {
	v *CompositeBucket
}

// NewCompositeBucket provides a builder for the CompositeBucket struct.
func NewCompositeBucketBuilder() *CompositeBucketBuilder {
	r := CompositeBucketBuilder{
		&CompositeBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
			Key:          make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the CompositeBucket struct
func (rb *CompositeBucketBuilder) Build() CompositeBucket {
	return *rb.v
}

func (rb *CompositeBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *CompositeBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *CompositeBucketBuilder) DocCount(doccount int64) *CompositeBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *CompositeBucketBuilder) Key(value map[string]interface{}) *CompositeBucketBuilder {
	rb.v.Key = value
	return rb
}
