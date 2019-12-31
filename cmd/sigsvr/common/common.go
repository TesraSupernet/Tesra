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

package common

import (
	"encoding/json"
	"fmt"
	"github.com/TesraSupernet/Tesra/account"
	"github.com/TesraSupernet/Tesra/cmd/sigsvr/store"
)

var DefWalletStore *store.WalletStore

type CliRpcRequest struct {
	Qid     string          `json:"qid"`
	Params  json.RawMessage `json:"params"`
	Account string          `json:"account"`
	Pwd     string          `json:"pwd"`
	Method  string          `json:"method"`
}

func (this *CliRpcRequest) GetAccount() (*account.Account, error) {
	var acc *account.Account
	var err error

	pwd := []byte(this.Pwd)
	if this.Pwd == "" {
		return nil, fmt.Errorf("pwd cannot empty")
	}
	if this.Account == "" {
		return nil, fmt.Errorf("account cannot empty")
	}
	acc, err = DefWalletStore.GetAccountByAddress(this.Account, pwd)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, fmt.Errorf("cannot find account by address: %s", this.Account)
	}
	return acc, nil
}

type CliRpcResponse struct {
	Qid       string      `json:"qid"`
	Method    string      `json:"method"`
	Result    interface{} `json:"result"`
	ErrorCode int         `json:"error_code"`
	ErrorInfo string      `json:"error_info"`
}
