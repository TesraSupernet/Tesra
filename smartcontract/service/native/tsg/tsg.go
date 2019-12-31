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

package tsg

import (
	"math/big"

	"fmt"

	"github.com/TesraSupernet/Tesra/common"
	"github.com/TesraSupernet/Tesra/common/constants"
	"github.com/TesraSupernet/Tesra/errors"
	"github.com/TesraSupernet/Tesra/smartcontract/service/native"
	tst "github.com/TesraSupernet/Tesra/smartcontract/service/native/tst"
	"github.com/TesraSupernet/Tesra/smartcontract/service/native/utils"
)

func InitTsg() {
	native.Contracts[utils.TsgContractAddress] = RegisterTsgContract
}

func RegisterTsgContract(native *native.NativeService) {
	native.Register(tst.INIT_NAME, TsgInit)
	native.Register(tst.TRANSFER_NAME, TsgTransfer)
	native.Register(tst.APPROVE_NAME, TsgApprove)
	native.Register(tst.TRANSFERFROM_NAME, TsgTransferFrom)
	native.Register(tst.NAME_NAME, TsgName)
	native.Register(tst.SYMBOL_NAME, TsgSymbol)
	native.Register(tst.DECIMALS_NAME, TsgDecimals)
	native.Register(tst.TOTALSUPPLY_NAME, TsgTotalSupply)
	native.Register(tst.BALANCEOF_NAME, TsgBalanceOf)
	native.Register(tst.ALLOWANCE_NAME, TsgAllowance)
}

func TsgInit(native *native.NativeService) ([]byte, error) {
	contract := native.ContextRef.CurrentContext().ContractAddress
	amount, err := utils.GetStorageUInt64(native, tst.GenTotalSupplyKey(contract))
	if err != nil {
		return utils.BYTE_FALSE, err
	}

	if amount > 0 {
		return utils.BYTE_FALSE, errors.NewErr("Init tsg has been completed!")
	}

	item := utils.GenUInt64StorageItem(constants.TSG_TOTAL_SUPPLY)
	native.CacheDB.Put(tst.GenTotalSupplyKey(contract), item.ToArray())
	native.CacheDB.Put(append(contract[:], utils.TstContractAddress[:]...), item.ToArray())
	tst.AddNotifications(native, contract, &tst.State{To: utils.TstContractAddress, Value: constants.TSG_TOTAL_SUPPLY})
	return utils.BYTE_TRUE, nil
}

func TsgTransfer(native *native.NativeService) ([]byte, error) {
	var transfers tst.Transfers
	source := common.NewZeroCopySource(native.Input)
	if err := transfers.Deserialization(source); err != nil {
		return utils.BYTE_FALSE, errors.NewDetailErr(err, errors.ErrNoCode, "[TsgTransfer] Transfers deserialize error!")
	}
	contract := native.ContextRef.CurrentContext().ContractAddress
	for _, v := range transfers.States {
		if v.Value == 0 {
			continue
		}
		if v.Value > constants.TSG_TOTAL_SUPPLY {
			return utils.BYTE_FALSE, fmt.Errorf("transfer tsg amount:%d over totalSupply:%d", v.Value, constants.TSG_TOTAL_SUPPLY)
		}
		if _, _, err := tst.Transfer(native, contract, &v); err != nil {
			return utils.BYTE_FALSE, err
		}
		tst.AddNotifications(native, contract, &v)
	}
	return utils.BYTE_TRUE, nil
}

func TsgApprove(native *native.NativeService) ([]byte, error) {
	var state tst.State
	source := common.NewZeroCopySource(native.Input)
	if err := state.Deserialization(source); err != nil {
		return utils.BYTE_FALSE, errors.NewDetailErr(err, errors.ErrNoCode, "[TsgApprove] state deserialize error!")
	}
	if state.Value > constants.TSG_TOTAL_SUPPLY {
		return utils.BYTE_FALSE, fmt.Errorf("approve tsg amount:%d over totalSupply:%d", state.Value, constants.TSG_TOTAL_SUPPLY)
	}
	if native.ContextRef.CheckWitness(state.From) == false {
		return utils.BYTE_FALSE, errors.NewErr("authentication failed!")
	}
	contract := native.ContextRef.CurrentContext().ContractAddress
	native.CacheDB.Put(tst.GenApproveKey(contract, state.From, state.To), utils.GenUInt64StorageItem(state.Value).ToArray())
	return utils.BYTE_TRUE, nil
}

func TsgTransferFrom(native *native.NativeService) ([]byte, error) {
	var state tst.TransferFrom
	source := common.NewZeroCopySource(native.Input)
	if err := state.Deserialization(source); err != nil {
		return utils.BYTE_FALSE, errors.NewDetailErr(err, errors.ErrNoCode, "[TstTransferFrom] State deserialize error!")
	}
	if state.Value == 0 {
		return utils.BYTE_FALSE, nil
	}
	if state.Value > constants.TSG_TOTAL_SUPPLY {
		return utils.BYTE_FALSE, fmt.Errorf("approve tsg amount:%d over totalSupply:%d", state.Value, constants.TSG_TOTAL_SUPPLY)
	}
	contract := native.ContextRef.CurrentContext().ContractAddress
	if _, _, err := tst.TransferedFrom(native, contract, &state); err != nil {
		return utils.BYTE_FALSE, err
	}
	tst.AddNotifications(native, contract, &tst.State{From: state.From, To: state.To, Value: state.Value})
	return utils.BYTE_TRUE, nil
}

func TsgName(native *native.NativeService) ([]byte, error) {
	return []byte(constants.TSG_NAME), nil
}

func TsgDecimals(native *native.NativeService) ([]byte, error) {
	return big.NewInt(int64(constants.TSG_DECIMALS)).Bytes(), nil
}

func TsgSymbol(native *native.NativeService) ([]byte, error) {
	return []byte(constants.TSG_SYMBOL), nil
}

func TsgTotalSupply(native *native.NativeService) ([]byte, error) {
	contract := native.ContextRef.CurrentContext().ContractAddress
	amount, err := utils.GetStorageUInt64(native, tst.GenTotalSupplyKey(contract))
	if err != nil {
		return utils.BYTE_FALSE, errors.NewDetailErr(err, errors.ErrNoCode, "[TstTotalSupply] get totalSupply error!")
	}
	return common.BigIntToNeoBytes(big.NewInt(int64(amount))), nil
}

func TsgBalanceOf(native *native.NativeService) ([]byte, error) {
	return tst.GetBalanceValue(native, tst.TRANSFER_FLAG)
}

func TsgAllowance(native *native.NativeService) ([]byte, error) {
	return tst.GetBalanceValue(native, tst.APPROVE_FLAG)
}
