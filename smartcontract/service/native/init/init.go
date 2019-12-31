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

package init

import (
	"bytes"
	"math/big"

	"github.com/TesraSupernet/Tesra/common"
	"github.com/TesraSupernet/Tesra/smartcontract/service/native/auth"
	params "github.com/TesraSupernet/Tesra/smartcontract/service/native/global_params"
	"github.com/TesraSupernet/Tesra/smartcontract/service/native/governance"
	tsg "github.com/TesraSupernet/Tesra/smartcontract/service/native/tsg"
	tst "github.com/TesraSupernet/Tesra/smartcontract/service/native/tst"
	tstid "github.com/TesraSupernet/Tesra/smartcontract/service/native/tstid"
	"github.com/TesraSupernet/Tesra/smartcontract/service/native/utils"
	"github.com/TesraSupernet/Tesra/smartcontract/service/neovm"
	vm "github.com/TesraSupernet/Tesra/vm/neovm"
)

var (
	COMMIT_DPOS_BYTES = InitBytes(utils.GovernanceContractAddress, governance.COMMIT_DPOS)
)

func init() {
	tsg.InitTsg()
	tst.InitTst()
	params.InitGlobalParams()
	tstid.Init()
	auth.Init()
	governance.InitGovernance()
}

func InitBytes(addr common.Address, method string) []byte {
	bf := new(bytes.Buffer)
	builder := vm.NewParamsBuilder(bf)
	builder.EmitPushByteArray([]byte{})
	builder.EmitPushByteArray([]byte(method))
	builder.EmitPushByteArray(addr[:])
	builder.EmitPushInteger(big.NewInt(0))
	builder.Emit(vm.SYSCALL)
	builder.EmitPushByteArray([]byte(neovm.NATIVE_INVOKE_NAME))

	return builder.ToArray()
}
