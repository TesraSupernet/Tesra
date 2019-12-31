/*
 * Copyright (C) 2019 The TesraSupernet Authors
 * This file is part of The TesraSupernet library.
 *
 * The TesraSupernet is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The TesraSupernet is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The TesraSupernet.  If not, see <http://www.gnu.org/licenses/>.
 */
package types

import (
	"github.com/TesraSupernet/tesracrypto/keypair"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressFromBookkeepers(t *testing.T) {
	_, pubKey1, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	_, pubKey2, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	_, pubKey3, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	pubkeys := []keypair.PublicKey{pubKey1, pubKey2, pubKey3}

	addr, _ := AddressFromBookkeepers(pubkeys)
	addr2, _ := AddressFromMultiPubKeys(pubkeys, 3)
	assert.Equal(t, addr, addr2)

	pubkeys = []keypair.PublicKey{pubKey3, pubKey2, pubKey1}
	addr3, _ := AddressFromMultiPubKeys(pubkeys, 3)

	assert.Equal(t, addr2, addr3)
}
