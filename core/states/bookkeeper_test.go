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
package states

import (
	"testing"

	"github.com/TesraSupernet/tesracrypto/keypair"
	"github.com/TesraSupernet/Tesra/common"
	"github.com/stretchr/testify/assert"
)

func TestBookkeeper_Deserialize_Serialize(t *testing.T) {
	_, pubKey1, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	_, pubKey2, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	_, pubKey3, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)
	_, pubKey4, _ := keypair.GenerateKeyPair(keypair.PK_ECDSA, keypair.P256)

	bk := BookkeeperState{
		StateBase:      StateBase{(byte)(1)},
		CurrBookkeeper: []keypair.PublicKey{pubKey1, pubKey2},
		NextBookkeeper: []keypair.PublicKey{pubKey3, pubKey4},
	}

	sink := common.NewZeroCopySink(nil)
	bk.Serialization(sink)
	bs := sink.Bytes()

	var bk2 BookkeeperState
	source := common.NewZeroCopySource(bs)
	bk2.Deserialization(source)
	assert.Equal(t, bk, bk2)

	source = common.NewZeroCopySource(bs[:len(bs)-1])
	err := bk2.Deserialization(source)
	assert.NotNil(t, err)
}
