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

package validation

import (
	"errors"
	"fmt"

	"github.com/TesraSupernet/Tesra/common"
	"github.com/TesraSupernet/Tesra/common/constants"
	"github.com/TesraSupernet/Tesra/common/log"
	"github.com/TesraSupernet/Tesra/core/ledger"
	"github.com/TesraSupernet/Tesra/core/payload"
	"github.com/TesraSupernet/Tesra/core/signature"
	"github.com/TesraSupernet/Tesra/core/types"
	tstErrors "github.com/TesraSupernet/Tesra/errors"
	"github.com/TesraSupernet/Tesra/smartcontract/service/wasmvm"
)

// VerifyTransaction verifys received single transaction
func VerifyTransaction(tx *types.Transaction) tstErrors.ErrCode {
	if err := checkTransactionSignatures(tx); err != nil {
		log.Info("transaction verify error:", err)
		return tstErrors.ErrVerifySignature
	}

	if err := checkTransactionPayload(tx); err != nil {
		log.Warn("[VerifyTransaction],", err)
		return tstErrors.ErrTransactionPayload
	}

	return tstErrors.ErrNoError
}

func VerifyTransactionWithLedger(tx *types.Transaction, ledger *ledger.Ledger) tstErrors.ErrCode {
	//TODO: replay check
	return tstErrors.ErrNoError
}

func checkTransactionSignatures(tx *types.Transaction) error {
	hash := tx.Hash()

	lensig := len(tx.Sigs)
	if lensig > constants.TX_MAX_SIG_SIZE {
		return fmt.Errorf("transaction signature number %d execced %d", lensig, constants.TX_MAX_SIG_SIZE)
	}

	address := make(map[common.Address]bool, len(tx.Sigs))
	for _, sigdata := range tx.Sigs {
		sig, err := sigdata.GetSig()
		if err != nil {
			return err
		}

		m := int(sig.M)
		kn := len(sig.PubKeys)
		sn := len(sig.SigData)

		if kn > constants.MULTI_SIG_MAX_PUBKEY_SIZE || sn < m || m > kn || m <= 0 {
			return errors.New("wrong tx sig param length")
		}

		if kn == 1 {
			err := signature.Verify(sig.PubKeys[0], hash[:], sig.SigData[0])
			if err != nil {
				return errors.New("signature verification failed")
			}

			address[types.AddressFromPubKey(sig.PubKeys[0])] = true
		} else {
			if err := signature.VerifyMultiSignature(hash[:], sig.PubKeys, m, sig.SigData); err != nil {
				return err
			}

			addr, err := types.AddressFromMultiPubKeys(sig.PubKeys, m)
			if err != nil {
				return err
			}
			address[addr] = true
		}
	}

	// check payer in address
	if address[tx.Payer] == false {
		return errors.New("signature missing for payer: " + tx.Payer.ToBase58())
	}

	addrList := make([]common.Address, 0, len(address))
	for addr := range address {
		addrList = append(addrList, addr)
	}

	tx.SignedAddr = addrList

	return nil
}

func checkTransactionPayload(tx *types.Transaction) error {

	switch pld := tx.Payload.(type) {
	case *payload.DeployCode:
		deploy := tx.Payload.(*payload.DeployCode)
		if deploy.VmType() == payload.WASMVM_TYPE {
			_, err := wasmvm.ReadWasmModule(deploy.GetRawCode(), true)
			if err != nil {
				return err
			}
		}
		return nil
	case *payload.InvokeCode:
		return nil
	default:
		return errors.New(fmt.Sprint("[txValidator], unimplemented transaction payload type.", pld))
	}
	return nil
}
