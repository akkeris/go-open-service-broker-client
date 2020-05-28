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

import (
	"testing"
)

func TestAtLeast(t *testing.T) {
	v2_15 := Version2_15()
	v2_14 := Version2_14()
	v2_13 := Version2_13()
	v2_12 := Version2_12()
	v2_11 := Version2_11()

	if !v2_15.AtLeast(v2_14) {
		t.Error("Expected 2.15 >= 2.14")
	}

	if !v2_14.AtLeast(v2_13) {
		t.Error("Expected 2.14 >= 2.13")
	}

	if !v2_13.AtLeast(v2_12) {
		t.Error("Expected 2.13 >= 2.12")
	}

	if !v2_12.AtLeast(v2_11) {
		t.Error("Expected 2.12 >= 2.11")
	}

	if v2_11.AtLeast(v2_12) {
		t.Error("Expected 2.11 < 2.12")
	}
}

func TestLatestAPIVersion(t *testing.T) {

	if LatestAPIVersion() != Version2_15() {
		t.Error("Unexpected Latest API Version--expected 2.15")
	}
}
