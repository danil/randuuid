// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randuuid_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/danil/randuuid"
	"github.com/google/uuid"
)

func line() int { _, _, l, _ := runtime.Caller(1); return l }

var NewTestCases = []struct {
	line      int
	seed      int64
	expected  uuid.UUID
	benchmark bool
}{
	{
		line:     line(),
		seed:     0,
		expected: uuid.Must(uuid.Parse("0194fdc2-fa2f-4cc0-81d3-ff12045b73c8")),
	},
	{
		line:     line(),
		seed:     42,
		expected: uuid.Must(uuid.Parse("538c7f96-b164-4f1b-97bb-9f4bb472e89f")),
	},
}

func TestNew(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range NewTestCases {
		tc := tc
		t.Run(fmt.Sprintf("seed %d", tc.seed), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			id, err := randuuid.New(tc.seed)
			if err != nil {
				t.Fatalf("new rand uuid function error: %s", err)
			}

			if id != tc.expected {
				t.Errorf("unexpected UUID, expected: %q, recieved: %q %s", tc.expected, id, linkToExample)
			}
		})
	}
}
