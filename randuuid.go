// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randuuid

import (
	"math/rand"

	"github.com/google/uuid"
)

// New returns a Random (Version 4) UUID.
//
// Copied from the github.com/google/uuid.
//
// The strength of the UUIDs is based on the strength of the crypto/rand
// package.
//
// A note about uniqueness derived from the UUID Wikipedia entry:
//
//  Randomly generated UUIDs have 122 random bits.  One's annual risk of being
//  hit by a meteorite is estimated to be one chance in 17 billion, that
//  means the probability is about 0.00000000006 (6 × 10−11),
//  equivalent to the odds of creating a few tens of trillions of UUIDs in a
//  year and having one duplicate.
//
func New(seed int64) (uuid.UUID, error) {
	var id uuid.UUID
	r := rand.New(rand.NewSource(seed))
	rander := make([]byte, 16)
	_, err := r.Read(rander)
	if err != nil {
		return uuid.Nil, err
	}
	copy(id[:], rander[:16])
	id[6] = (id[6] & 0x0f) | 0x40 // Version 4
	id[8] = (id[8] & 0x3f) | 0x80 // Variant is 10
	return id, nil
}
