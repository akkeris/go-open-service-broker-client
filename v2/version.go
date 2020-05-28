/*
Copyright 2019 The Kubernetes Authors.

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

package v2

// APIVersion represents a specific version of the OSB API.
type APIVersion struct {
	label string
	order byte
}

// AtLeast returns whether the API version is greater than or equal to the
// given API version.
func (v APIVersion) AtLeast(test APIVersion) bool {
	return v.order >= test.order
}

// HeaderValue returns the value that should be sent in the API version header
// for this API version.
func (v APIVersion) HeaderValue() string {
	return v.label
}

func (v APIVersion) String() string {
	return v.label
}

// LatestAPIVersion returns the latest supported API version in the current
// release of this library.
func LatestAPIVersion() APIVersion {
	return Version2_15()
}

// APIVersions returns a list of the APIVersions supported by this library, with
// no guarantees of ordering.
func APIVersions() []APIVersion {
	return []APIVersion{
		Version2_11(),
		Version2_12(),
		Version2_13(),
		Version2_14(),
		Version2_15(),
	}
}

const (
	// internalAPIVersion2_11 represents the 2.11 version of the Open Service
	// Broker API.
	internalAPIVersion2_11 = "2.11"

	// internalAPIVersion2_12 represents the 2.12 version of the Open Service
	// Broker API.
	internalAPIVersion2_12 = "2.12"

	// internalAPIVersion2_13 represents the 2.13 version of the Open Service
	// Broker API.
	internalAPIVersion2_13 = "2.13"

	// internalAPIVersion2_13 represents the 2.13 version of the Open Service
	// Broker API.
	internalAPIVersion2_14 = "2.14"

	// internalAPIVersion2_13 represents the 2.13 version of the Open Service
	// Broker API.
	internalAPIVersion2_15 = "2.15"
)

//Version2_11 returns an APIVersion struct with the internal API version set to "2.11"
func Version2_11() APIVersion {
	return APIVersion{label: internalAPIVersion2_11, order: 0}
}

//Version2_12 returns an APIVersion struct with the internal API version set to "2.12"
func Version2_12() APIVersion {
	return APIVersion{label: internalAPIVersion2_12, order: 1}
}

//Version2_13 returns an APIVersion struct with the internal API version set to "2.13"
func Version2_13() APIVersion {
	return APIVersion{label: internalAPIVersion2_13, order: 2}
}

//Version2_13 returns an APIVersion struct with the internal API version set to "2.14"
func Version2_14() APIVersion {
	return APIVersion{label: internalAPIVersion2_14, order: 3}
}

//Version2_13 returns an APIVersion struct with the internal API version set to "2.15"
func Version2_15() APIVersion {
	return APIVersion{label: internalAPIVersion2_15, order: 4}
}
