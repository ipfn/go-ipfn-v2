// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package digest

import (
	keccak "github.com/gxed/hashland/keccakpg"
	"github.com/minio/sha256-simd"
)

// SumKeccak256Bytes - Sums Keccak256 secure hash.
func SumKeccak256Bytes(data ...[]byte) []byte {
	return SumBytes(keccak.New256(), data...)
}

// SumSha256Bytes - Sums Sha256 secure hash.
func SumSha256Bytes(data ...[]byte) []byte {
	return SumBytes(sha256.New(), data...)
}

// SumKeccak256 - Sums Keccak256 secure hash.
func SumKeccak256(data ...[]byte) Digest {
	return Sum(keccak.New256(), data...)
}

// SumSha256 - Sums Sha256 secure hash.
func SumSha256(data ...[]byte) Digest {
	return Sum(sha256.New(), data...)
}
