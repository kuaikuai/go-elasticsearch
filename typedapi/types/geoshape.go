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

// GeoShape type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/_types/Geo.ts#L56-L57
type GeoShape interface{}

// GeoShapeBuilder holds GeoShape struct and provides a builder API.
type GeoShapeBuilder struct {
	v GeoShape
}

// NewGeoShape provides a builder for the GeoShape struct.
func NewGeoShapeBuilder() *GeoShapeBuilder {
	return &GeoShapeBuilder{}
}

// Build finalize the chain and returns the GeoShape struct
func (b *GeoShapeBuilder) Build() GeoShape {
	return b.v
}

func (b *GeoShapeBuilder) GeoShape(value GeoShape) *GeoShapeBuilder {
	b.v = value
	return b
}
