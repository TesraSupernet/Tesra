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

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TesraSupernet/Tesra/common/constants"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strings"
)

const (
	PRECISION_TSG = 9
	PRECISION_TST = 0
)

//FormatAssetAmount return asset amount multiplied by math.Pow10(precision) to raw float string
//For example 1000000000123456789 => 1000000000.123456789
func FormatAssetAmount(amount uint64, precision byte) string {
	if precision == 0 {
		return fmt.Sprintf("%d", amount)
	}
	divisor := math.Pow10(int(precision))
	intPart := amount / uint64(divisor)
	fracPart := amount - intPart*uint64(divisor)
	if fracPart == 0 {
		return fmt.Sprintf("%d", intPart)
	}
	bf := new(big.Float).SetUint64(fracPart)
	bf.Quo(bf, new(big.Float).SetFloat64(math.Pow10(int(precision))))
	bf.Add(bf, new(big.Float).SetUint64(intPart))
	return bf.Text('f', -1)
}

//ParseAssetAmount return raw float string to uint64 multiplied by math.Pow10(precision)
//For example 1000000000.123456789 => 1000000000123456789
func ParseAssetAmount(rawAmount string, precision byte) uint64 {
	bf, ok := new(big.Float).SetString(rawAmount)
	if !ok {
		return 0
	}
	bf.Mul(bf, new(big.Float).SetFloat64(math.Pow10(int(precision))))
	amount, _ := bf.Uint64()
	return amount
}

func FormatTsg(amount uint64) string {
	return FormatAssetAmount(amount, PRECISION_TSG)
}

func ParseTsg(rawAmount string) uint64 {
	return ParseAssetAmount(rawAmount, PRECISION_TSG)
}

func FormatTst(amount uint64) string {
	return FormatAssetAmount(amount, PRECISION_TST)
}

func ParseTst(rawAmount string) uint64 {
	return ParseAssetAmount(rawAmount, PRECISION_TST)
}

func CheckAssetAmount(asset string, amount uint64) error {
	switch strings.ToLower(asset) {
	case "tst":
		if amount > constants.TST_TOTAL_SUPPLY {
			return fmt.Errorf("amount:%d larger than TST total supply:%d", amount, constants.TST_TOTAL_SUPPLY)
		}
	case "tsg":
		if amount > constants.TSG_TOTAL_SUPPLY {
			return fmt.Errorf("amount:%d larger than TSG total supply:%d", amount, constants.TSG_TOTAL_SUPPLY)
		}
	default:
		return fmt.Errorf("unknown asset:%s", asset)
	}
	return nil
}

func GetJsonObjectFromFile(filePath string, jsonObject interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	// Remove the UTF-8 Byte Order Mark
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	err = json.Unmarshal(data, jsonObject)
	if err != nil {
		return fmt.Errorf("json.Unmarshal %s error:%s", data, err)
	}
	return nil
}

func GetStoreDirPath(dataDir, networkName string) string {
	return dataDir + string(os.PathSeparator) + networkName
}

func GenExportBlocksFileName(name string, start, end uint32) string {
	index := strings.LastIndex(name, ".")
	fileName := ""
	fileExt := ""
	if index < 0 {
		fileName = name
	} else {
		fileName = name[0:index]
		if index < len(name)-1 {
			fileExt = name[index+1:]
		}
	}
	fileName = fmt.Sprintf("%s_%d_%d", fileName, start, end)
	if index > 0 {
		fileName = fileName + "." + fileExt
	}
	return fileName
}
