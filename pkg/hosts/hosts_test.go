/*
Copyright 2017 The Kubernetes Authors.

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

package hosts

import (
	"reflect"
	"testing"
)

func etcHostTests(file []byte, t *testing.T) {
	expectedOut := []string{"", "# Begin host entries managed by etcd-manager[etcd] - do not edit", "10.0.0.1\tetc1.local.test.com etc1.local.test.net", "# End host entries managed by etcd-manager[etcd]", ""}
	addrToHosts := make(map[string][]string)
	addrToHosts["10.0.0.1"] = []string{"etc1.local.test.com", "etc1.local.test.net"}
	data, err := updateHostsFileWithRecords(file, "etcd-manager[etcd]", addrToHosts)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(data, expectedOut) {
		t.Fatal("Expected Output Did Not Match")
	}
	t.Log(data)
}

func TestEmptyFile(t *testing.T) {
	var emptyFile []byte
	etcHostTests(emptyFile, t)
}
func TestParseEmptyEctdBlock(t *testing.T) {
	file, err := readHostsFile("./testdata/emptyetcdblock")
	if err != nil {
		t.Fatal(err.Error())
	}
	etcHostTests(file, t)
}

func TestParseMultipleEmptyEctdBlock(t *testing.T) {
	file, err := readHostsFile("./testdata/multipleemptyetcdblocks")
	if err != nil {
		t.Fatal(err.Error())
	}
	etcHostTests(file, t)
}

func TestBadEctdFile(t *testing.T) {
	file, err := readHostsFile("./testdata/badhostfile")
	if err != nil {
		t.Fatal(err.Error())
	}
	etcHostTests(file, t)
}
