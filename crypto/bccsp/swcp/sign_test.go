// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2016-2018 IBM Corp. All Rights Reserved.
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

package swcp

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ipfn/ipfn/pkg/crypto/bccsp"
	mocks2 "github.com/ipfn/ipfn/pkg/crypto/bccsp/mocks"
	"github.com/ipfn/ipfn/pkg/crypto/bccsp/swcp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	t.Parallel()

	expectedKey := &mocks2.MockKey{}
	expectetDigest := []byte{1, 2, 3, 4}
	expectedOpts := &mocks2.SignerOpts{}
	expectetValue := []byte{0, 1, 2, 3, 4}
	expectedErr := errors.New("Expected Error")

	signers := make(map[reflect.Type]bccsp.Signer)
	signers[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.Signer{
		KeyArg:    expectedKey,
		DigestArg: expectetDigest,
		OptsArg:   expectedOpts,
		Value:     expectetValue,
		Err:       nil,
	}
	csp := CSP{signers: signers}
	value, err := csp.Sign(expectedKey, expectetDigest, expectedOpts)
	assert.Equal(t, expectetValue, value)
	assert.Nil(t, err)

	signers = make(map[reflect.Type]bccsp.Signer)
	signers[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.Signer{
		KeyArg:    expectedKey,
		DigestArg: expectetDigest,
		OptsArg:   expectedOpts,
		Value:     nil,
		Err:       expectedErr,
	}
	csp = CSP{signers: signers}
	value, err = csp.Sign(expectedKey, expectetDigest, expectedOpts)
	assert.Nil(t, value)
	assert.Contains(t, err.Error(), expectedErr.Error())
}
