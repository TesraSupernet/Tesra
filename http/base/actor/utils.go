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

// Package actor privides communication with other actor
package actor

import (
	"github.com/TesraSupernet/Tesra/common"
	"github.com/TesraSupernet/Tesra/smartcontract/service/native/utils"
)

func updateNativeSCAddr(hash common.Address) common.Address {
	if hash == utils.TstContractAddress {
		hash = common.AddressFromVmCode(utils.TstContractAddress[:])
	} else if hash == utils.TsgContractAddress {
		hash = common.AddressFromVmCode(utils.TsgContractAddress[:])
	} else if hash == utils.TstIDContractAddress {
		hash = common.AddressFromVmCode(utils.TstIDContractAddress[:])
	} else if hash == utils.ParamContractAddress {
		hash = common.AddressFromVmCode(utils.ParamContractAddress[:])
	} else if hash == utils.AuthContractAddress {
		hash = common.AddressFromVmCode(utils.AuthContractAddress[:])
	} else if hash == utils.GovernanceContractAddress {
		hash = common.AddressFromVmCode(utils.GovernanceContractAddress[:])
	}
	return hash
}
