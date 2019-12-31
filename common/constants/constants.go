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

package constants

import (
	"time"
)

// genesis constants
var (
	//TODO: modify this when on mainnet
	GENESIS_BLOCK_TIMESTAMP = uint32(time.Date(2019, time.December, 30, 0, 0, 0, 0, time.UTC).Unix())
)

// tst constants
const (
	TST_NAME         = "TST Token"
	TST_SYMBOL       = "TST"
	TST_DECIMALS     = 0
	TST_TOTAL_SUPPLY = uint64(1000000000)
)

// tsg constants
const (
	TSG_NAME         = "TSG Token"
	TSG_SYMBOL       = "TSG"
	TSG_DECIMALS     = 9
	TSG_TOTAL_SUPPLY = uint64(1000000000000000000)
)

// tst/tsg unbound model constants
const UNBOUND_TIME_INTERVAL = uint32(31536000)

var UNBOUND_GENERATION_AMOUNT = [18]uint64{5, 4, 3, 3, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

// the end of unbound timestamp offset from genesis block's timestamp
var UNBOUND_DEADLINE = (func() uint32 {
	count := uint64(0)
	for _, m := range UNBOUND_GENERATION_AMOUNT {
		count += m
	}
	count *= uint64(UNBOUND_TIME_INTERVAL)

	numInterval := len(UNBOUND_GENERATION_AMOUNT)

	if UNBOUND_GENERATION_AMOUNT[numInterval-1] != 1 ||
		!(count-uint64(UNBOUND_TIME_INTERVAL) < TST_TOTAL_SUPPLY && TST_TOTAL_SUPPLY <= count) {
		panic("incompatible constants setting")
	}

	return UNBOUND_TIME_INTERVAL*uint32(numInterval) - uint32(count-uint64(TST_TOTAL_SUPPLY))
})()

// multi-sig constants
const MULTI_SIG_MAX_PUBKEY_SIZE = 16

// transaction constants
const TX_MAX_SIG_SIZE = 16

// network magic number
const (
	NETWORK_MAGIC_MAINNET = 0x73ac67c8 //0x8c77ab60
	NETWORK_MAGIC_SCORPIO = 0x2d8829df
)

// ledger state hash check height
const STATE_HASH_HEIGHT_MAINNET = 3000000
const STATE_HASH_HEIGHT_SCORPIO = 850000

// neovm opcode update check height
const OPCODE_HEIGHT_UPDATE_FIRST_MAINNET = 6300000
const OPCODE_HEIGHT_UPDATE_FIRST_SCORPIO = 2100000
