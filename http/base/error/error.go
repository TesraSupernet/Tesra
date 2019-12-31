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

// Package error privides error code for http
package error

import tstErrors "github.com/TesraSupernet/Tesra/errors"

const (
	SUCCESS            int64 = 0
	SESSION_EXPIRED    int64 = 41001
	SERVICE_CEILING    int64 = 41002
	ILLEGAL_DATAFORMAT int64 = 41003
	INVALID_VERSION    int64 = 41004

	INVALID_METHOD int64 = 42001
	INVALID_PARAMS int64 = 42002

	INVALID_TRANSACTION int64 = 43001
	INVALID_ASSET       int64 = 43002
	INVALID_BLOCK       int64 = 43003

	UNKNOWN_TRANSACTION int64 = 44001
	UNKNOWN_ASSET       int64 = 44002
	UNKNOWN_BLOCK       int64 = 44003
	UNKNOWN_CONTRACT    int64 = 44004

	INTERNAL_ERROR  int64 = 45001
	SMARTCODE_ERROR int64 = 47001
	PRE_EXEC_ERROR  int64 = 47002
)

var ErrMap = map[int64]string{
	SUCCESS:            "SUCCESS",
	SESSION_EXPIRED:    "SESSION EXPIRED",
	SERVICE_CEILING:    "SERVICE CEILING",
	ILLEGAL_DATAFORMAT: "ILLEGAL DATAFORMAT",
	INVALID_VERSION:    "INVALID VERSION",

	INVALID_METHOD: "INVALID METHOD",
	INVALID_PARAMS: "INVALID PARAMS",

	INVALID_TRANSACTION: "INVALID TRANSACTION",
	INVALID_ASSET:       "INVALID ASSET",
	INVALID_BLOCK:       "INVALID BLOCK",

	UNKNOWN_TRANSACTION: "UNKNOWN TRANSACTION",
	UNKNOWN_ASSET:       "UNKNOWN ASSET",
	UNKNOWN_BLOCK:       "UNKNOWN BLOCK",
	UNKNOWN_CONTRACT:    "UNKNOWN CONTRACT",

	INTERNAL_ERROR:                           "INTERNAL ERROR",
	SMARTCODE_ERROR:                          "SMARTCODE EXEC ERROR",
	PRE_EXEC_ERROR:                           "SMARTCODE PREPARE EXEC ERROR",
	int64(tstErrors.ErrNoCode):               "INTERNAL ERROR, ErrNoCode",
	int64(tstErrors.ErrUnknown):              "INTERNAL ERROR, ErrUnknown",
	int64(tstErrors.ErrDuplicatedTx):         "INTERNAL ERROR, ErrDuplicatedTx",
	int64(tstErrors.ErrDuplicateInput):       "INTERNAL ERROR, ErrDuplicateInput",
	int64(tstErrors.ErrAssetPrecision):       "INTERNAL ERROR, ErrAssetPrecision",
	int64(tstErrors.ErrTransactionBalance):   "INTERNAL ERROR, ErrTransactionBalance",
	int64(tstErrors.ErrAttributeProgram):     "INTERNAL ERROR, ErrAttributeProgram",
	int64(tstErrors.ErrTransactionContracts): "INTERNAL ERROR, ErrTransactionContracts",
	int64(tstErrors.ErrTransactionPayload):   "INTERNAL ERROR, ErrTransactionPayload",
	int64(tstErrors.ErrDoubleSpend):          "INTERNAL ERROR, ErrDoubleSpend",
	int64(tstErrors.ErrTxHashDuplicate):      "INTERNAL ERROR, ErrTxHashDuplicate",
	int64(tstErrors.ErrStateUpdaterVaild):    "INTERNAL ERROR, ErrStateUpdaterVaild",
	int64(tstErrors.ErrSummaryAsset):         "INTERNAL ERROR, ErrSummaryAsset",
	int64(tstErrors.ErrXmitFail):             "INTERNAL ERROR, ErrXmitFail",
	int64(tstErrors.ErrNoAccount):            "INTERNAL ERROR, ErrNoAccount",
}
