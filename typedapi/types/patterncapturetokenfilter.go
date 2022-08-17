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

// PatternCaptureTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/_types/analysis/token_filters.ts#L277-L281
type PatternCaptureTokenFilter struct {
	Patterns         []string       `json:"patterns"`
	PreserveOriginal *bool          `json:"preserve_original,omitempty"`
	Type             string         `json:"type,omitempty"`
	Version          *VersionString `json:"version,omitempty"`
}

// PatternCaptureTokenFilterBuilder holds PatternCaptureTokenFilter struct and provides a builder API.
type PatternCaptureTokenFilterBuilder struct {
	v *PatternCaptureTokenFilter
}

// NewPatternCaptureTokenFilter provides a builder for the PatternCaptureTokenFilter struct.
func NewPatternCaptureTokenFilterBuilder() *PatternCaptureTokenFilterBuilder {
	r := PatternCaptureTokenFilterBuilder{
		&PatternCaptureTokenFilter{},
	}

	r.v.Type = "pattern_capture"

	return &r
}

// Build finalize the chain and returns the PatternCaptureTokenFilter struct
func (rb *PatternCaptureTokenFilterBuilder) Build() PatternCaptureTokenFilter {
	return *rb.v
}

func (rb *PatternCaptureTokenFilterBuilder) Patterns(patterns ...string) *PatternCaptureTokenFilterBuilder {
	rb.v.Patterns = patterns
	return rb
}

func (rb *PatternCaptureTokenFilterBuilder) PreserveOriginal(preserveoriginal bool) *PatternCaptureTokenFilterBuilder {
	rb.v.PreserveOriginal = &preserveoriginal
	return rb
}

func (rb *PatternCaptureTokenFilterBuilder) Version(version VersionString) *PatternCaptureTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
