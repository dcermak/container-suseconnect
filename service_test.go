//
// Copyright (c) 2015 SUSE LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"bytes"
	"testing"
	"strings"
)

var expectedRepoFile = `# generated by container-suseconnect

[SLES12-Updates]
name=SLES12-Updates for sle-12-x86_64
baseurl=https://smt.test.lan/repo/SUSE/Updates/SLE-SERVER/12/x86_64/update
autorefresh=1
enabled=1

[SLES12-Debuginfo-Updates]
name=SLES12-Debuginfo-Updates for sle-12-x86_64
baseurl=https://smt.test.lan/repo/SUSE/Updates/SLE-SERVER/12/x86_64/update_debug
autorefresh=1
enabled=0

[SLES12-Pool]
name=SLES12-Pool for sle-12-x86_64
baseurl=https://smt.test.lan/repo/SUSE/Products/SLE-SERVER/12/x86_64/product
autorefresh=0
enabled=1

[SLES12-Debuginfo-Pool]
name=SLES12-Debuginfo-Pool for sle-12-x86_64
baseurl=https://smt.test.lan/repo/SUSE/Products/SLE-SERVER/12/x86_64/product_debug
autorefresh=0
enabled=0

`

func TestServiceOutput(t *testing.T) {
	reader := strings.NewReader(sccReply)

	product, err := ParseProduct(reader)
	if (err != nil) {
		t.Errorf(err.Error())
	}

	buf := bytes.Buffer{}
	DumpRepositories(&buf, product)

	result := buf.String()
	if expectedRepoFile != result {
		t.Errorf(result)
	}
}