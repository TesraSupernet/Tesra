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

package test

import (
	"os"
	"testing"

	"github.com/TesraSupernet/Tesra/account"
	"github.com/TesraSupernet/Tesra/cmd/utils"
	"github.com/TesraSupernet/Tesra/core/payload"
	"github.com/TesraSupernet/Tesra/core/store/ledgerstore"
	"github.com/stretchr/testify/assert"
)

func TestPreExecuteContractWasmDeploy(t *testing.T) {
	acct := account.NewAccount("")
	testLedgerStore, err := ledgerstore.NewLedgerStore("test/ledgerfortmp", 0)
	code := []byte{0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00, 0x01, 0x09, 0x02, 0x60, 0x00, 0x00, 0x60, 0x02, 0x7f, 0x7f, 0x00, 0x02, 0x14, 0x01, 0x03, 0x65, 0x6e, 0x76, 0x0c, 0x6f, 0x6e, 0x74, 0x69, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x00, 0x01, 0x03, 0x02, 0x01, 0x00, 0x05, 0x03, 0x01, 0x00, 0x01, 0x07, 0x0a, 0x01, 0x06, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x00, 0x01, 0x0a, 0x12, 0x01, 0x10, 0x00, 0x41, 0x00, 0x42, 0xae, 0x11, 0x37, 0x03, 0x00, 0x41, 0x00, 0x41, 0x08, 0x10, 0x00, 0x0b}
	mutable, _ := utils.NewDeployCodeTransaction(0, 100000000, code, payload.NEOVM_TYPE, "name", "version",
		"author", "email", "desc")
	_ = utils.SignTransaction(acct, mutable)
	tx, err := mutable.IntoImmutable()
	_, err = testLedgerStore.PreExecuteContract(tx)
	assert.EqualError(t, err, "this code is wasm binary. can not deployed as neo contract")

	mutable, _ = utils.NewDeployCodeTransaction(0, 100000000, code, payload.WASMVM_TYPE, "name", "version",
		"author", "email", "desc")
	_ = utils.SignTransaction(acct, mutable)
	tx, err = mutable.IntoImmutable()
	_, err = testLedgerStore.PreExecuteContract(tx)
	assert.Nil(t, err)

	_ = os.RemoveAll("./test")
}
