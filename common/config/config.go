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

package config

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"github.com/TesraSupernet/Tesra/common"
	"github.com/TesraSupernet/Tesra/common/constants"
	"github.com/TesraSupernet/Tesra/common/log"
	"github.com/TesraSupernet/Tesra/errors"
	"github.com/TesraSupernet/tesracrypto/keypair"
)

var Version = "" //Set value when build project

const (
	DEFAULT_CONFIG_FILE_NAME = "./tesranode.json"
	DEFAULT_WALLET_FILE_NAME = "./twallet.dat"
	MIN_GEN_BLOCK_TIME       = 2
	DEFAULT_GEN_BLOCK_TIME   = 6
	DBFT_MIN_NODE_NUM        = 4 //min node number of dbft consensus
	SOLO_MIN_NODE_NUM        = 1 //min node number of solo consensus
	VBFT_MIN_NODE_NUM        = 4 //min node number of vbft consensus

	CONSENSUS_TYPE_DBFT = "dbft"
	CONSENSUS_TYPE_SOLO = "solo"
	CONSENSUS_TYPE_VBFT = "vbft"

	DEFAULT_LOG_LEVEL                       = log.InfoLog
	DEFAULT_MAX_LOG_SIZE                    = 100         //MByte
	DEFAULT_NODE_PORT                       = uint(25766) //uint(20338)
	DEFAULT_CONSENSUS_PORT                  = uint(25767) //uint(20339)
	DEFAULT_RPC_PORT                        = uint(25768) //uint(20336)
	DEFAULT_RPC_LOCAL_PORT                  = uint(25769) //uint(20337)
	DEFAULT_REST_PORT                       = uint(25770) //uint(20334)
	DEFAULT_WS_PORT                         = uint(25771) //uint(20335)
	DEFAULT_REST_MAX_CONN                   = uint(1024)
	DEFAULT_MAX_CONN_IN_BOUND               = uint(1024)
	DEFAULT_MAX_CONN_OUT_BOUND              = uint(1024)
	DEFAULT_MAX_CONN_IN_BOUND_FOR_SINGLE_IP = uint(16)
	DEFAULT_HTTP_INFO_PORT                  = uint(0)
	DEFAULT_MAX_TX_IN_BLOCK                 = 60000
	DEFAULT_MAX_SYNC_HEADER                 = 500
	DEFAULT_ENABLE_CONSENSUS                = true
	DEFAULT_ENABLE_EVENT_LOG                = true
	DEFAULT_CLI_RPC_PORT                    = uint(20000)
	DEFUALT_CLI_RPC_ADDRESS                 = "127.0.0.1"
	DEFAULT_GAS_LIMIT                       = 20000
	DEFAULT_GAS_PRICE                       = 500
	DEFAULT_WASM_GAS_FACTOR                 = uint64(10)
	DEFAULT_WASM_MAX_STEPCOUNT              = uint64(8000000)

	DEFAULT_DATA_DIR      = "./Chain"
	DEFAULT_RESERVED_FILE = "./peers.rsv"
)

const (
	WASM_GAS_FACTOR = "WASM_GAS_FACTOR"
)

const (
	NETWORK_ID_MAIN_NET      = 1
	NETWORK_ID_SCORPIO_NET   = 2
	NETWORK_ID_SOLO_NET      = 3
	NETWORK_NAME_MAIN_NET    = "tesranode"
	NETWORK_NAME_SCORPIO_NET = "scorpio"
	NETWORK_NAME_SOLO_NET    = "testmode"
)

var NETWORK_MAGIC = map[uint32]uint32{
	NETWORK_ID_MAIN_NET:    constants.NETWORK_MAGIC_MAINNET, //Network main
	NETWORK_ID_SCORPIO_NET: constants.NETWORK_MAGIC_SCORPIO, //Network scorpio
	NETWORK_ID_SOLO_NET:    0,                               //Network solo
}

var NETWORK_NAME = map[uint32]string{
	NETWORK_ID_MAIN_NET:    NETWORK_NAME_MAIN_NET,
	NETWORK_ID_SCORPIO_NET: NETWORK_NAME_SCORPIO_NET,
	NETWORK_ID_SOLO_NET:    NETWORK_NAME_SOLO_NET,
}

func GetNetworkMagic(id uint32) uint32 {
	nid, ok := NETWORK_MAGIC[id]
	if ok {
		return nid
	}
	return id
}

var STATE_HASH_CHECK_HEIGHT = map[uint32]uint32{
	NETWORK_ID_MAIN_NET:    constants.STATE_HASH_HEIGHT_MAINNET, //Network main
	NETWORK_ID_SCORPIO_NET: constants.STATE_HASH_HEIGHT_SCORPIO, //Network scorpio
	NETWORK_ID_SOLO_NET:    0,                                   //Network solo
}

func GetStateHashCheckHeight(id uint32) uint32 {
	return STATE_HASH_CHECK_HEIGHT[id]
}

var OPCODE_HASKEY_ENABLE_HEIGHT = map[uint32]uint32{
	NETWORK_ID_MAIN_NET:    constants.OPCODE_HEIGHT_UPDATE_FIRST_MAINNET, //Network main
	NETWORK_ID_SCORPIO_NET: constants.OPCODE_HEIGHT_UPDATE_FIRST_SCORPIO, //Network scorpio
	NETWORK_ID_SOLO_NET:    0,                                            //Network solo
}

func GetOpcodeUpdateCheckHeight(id uint32) uint32 {
	return OPCODE_HASKEY_ENABLE_HEIGHT[id]
}

func GetNetworkName(id uint32) string {
	name, ok := NETWORK_NAME[id]
	if ok {
		return name
	}
	return fmt.Sprintf("%d", id)
}

var ScorpioConfig = &GenesisConfig{
	SeedList: []string{
		"dapp1.tesra.me:25766",
		"dapp2.tesra.me:25766",
		"dapp3.tesra.me:25766",
		"dapp4.tesra.me:25766"},
	ConsensusType: CONSENSUS_TYPE_VBFT,
	VBFT: &VBFTConfig{
		N:                    7,
		C:                    2,
		K:                    7,
		L:                    112,
		BlockMsgDelay:        10000,
		HashMsgDelay:         10000,
		PeerHandshakeTimeout: 10,
		MaxBlockChangeView:   3000,
		AdminTstID:           "did:tst:AMAx993nE6NEqZjwBssUfopxnnvTdob9ij",
		MinInitStake:         10000,
		VrfValue:             "1c9810aa9822e511d5804a9c4db9dd08497c31087b0daafa34d768a3253441fa20515e2f30f81741102af0ca3cefc4818fef16adb825fbaa8cad78647f3afb590e",
		VrfProof:             "c57741f934042cb8d8b087b44b161db56fc3ffd4ffb675d36cd09f83935be853d8729f3f5298d12d6fd28d45dde515a4b9d7f67682d182ba5118abf451ff1988",
		Peers: []*VBFTPeerStakeInfo{
			{
				Index:      1,
				PeerPubkey: "037c9e6c6a446b6b296f89b722cbf686b81e0a122444ef05f0f87096777663284b",
				Address:    "AXmQDzzvpEtPkNwBEFsREzApTTDZFW6frD",
				InitPos:    10000,
			},
			{
				Index:      2,
				PeerPubkey: "03dff4c63267ae5e23da44ace1bc47d0da1eb8d36fd71181dcccf0e872cb7b31fa",
				Address:    "AY5W6p4jHeZG2jjW6nS1p4KDUhcqLkU6jz",
				InitPos:    20000,
			},
			{
				Index:      3,
				PeerPubkey: "0205bc592aa9121428c4144fcd669ece1fa73fee440616c75624967f83fb881050",
				Address:    "ALZVrZrFqoSvqyi38n7mpPoeDp7DMtZ9b6",
				InitPos:    30000,
			},
			{
				Index:      4,
				PeerPubkey: "030a34dcb075d144df1f65757b85acaf053395bb47b019970607d2d1cdd222525c",
				Address:    "AMogjmLf2QohTcGST7niV75ekZfj44SKme",
				InitPos:    40000,
			},
			{
				Index:      5,
				PeerPubkey: "021844159f97d81da71da52f84e8451ee573c83b296ff2446387b292e44fba5c98",
				Address:    "AZzQTkZvjy7ih9gjvwU8KYiZZyNoy6jE9p",
				InitPos:    30000,
			},
			{
				Index:      6,
				PeerPubkey: "020cc76feb375d6ea8ec9ff653bab18b6bbc815610cecc76e702b43d356f885835",
				Address:    "AKEqQKmxCsjWJz8LPGryXzb6nN5fkK1WDY",
				InitPos:    20000,
			},
			{
				Index:      7,
				PeerPubkey: "03aa4d52b200fd91ca12deff46505c4608a0f66d28d9ae68a342c8a8c1266de0f9",
				Address:    "AQNpGWz4oHHFBejtBbakeR43DHfen7cm8L",
				InitPos:    10000,
			},
		},
	},
	DBFT: &DBFTConfig{},
	SOLO: &SOLOConfig{},
}

// var MainNetConfig = &GenesisConfig{
// 	SeedList: []string{
// 		"127.0.0.1:25666",
// 		"127.0.0.1:25665",
// 		"127.0.0.1:25664",
// 		"127.0.0.1:25663",
// 		"127.0.0.1:25662",
// 		"127.0.0.1:25661",
// 		"127.0.0.1:25660"},
// 	ConsensusType: CONSENSUS_TYPE_VBFT,
// 	VBFT: &VBFTConfig{
// 		N:                    7,
// 		C:                    2,
// 		K:                    7,
// 		L:                    112,
// 		BlockMsgDelay:        10000,
// 		HashMsgDelay:         10000,
// 		PeerHandshakeTimeout: 10,
// 		MaxBlockChangeView:   3000,
// 		AdminTstID:           "did:tst:TYhUu9bidgc7MbkWeS55sywjct5VYrqxdW",
// 		MinInitStake:         10000,
// 		VrfValue:             "1c9810aa9822e511d5804a9c4db9dd08497c31087b0daafa34d768a3253441fa20515e2f30f81741102af0ca3cefc4818fef16adb825fbaa8cad78647f3afb590e",
// 		VrfProof:             "c57741f934042cb8d8b087b44b161db56fc3ffd4ffb675d36cd09f83935be853d8729f3f5298d12d6fd28d45dde515a4b9d7f67682d182ba5118abf451ff1988",
// 		Peers: []*VBFTPeerStakeInfo{
// 			{
// 				Index:      1,
// 				PeerPubkey: "03c5e12f54e1104ce5fb9a10c32850776864b03e6bc2ff2285efbb1abb9d9ba36a",
// 				Address:    "AXKnGhh9WyoB1G5aXHDdzgSq4Zv3CJRNNs",
// 				InitPos:    10000,
// 			},
// 			{
// 				Index:      2,
// 				PeerPubkey: "031dadf8fa327b9cc4c03bb8796d32c7048695b574cf37d7d4d30c6647958cddc9",
// 				Address:    "AbLD3PEEg8ohKX2QP5YUoA8hj4hgnaKH5d",
// 				InitPos:    20000,
// 			},
// 			{
// 				Index:      3,
// 				PeerPubkey: "0247cecfbdfb75ed9da031196dde8579d3f0df3f13fca05a82ab708c43b1fa9b5b",
// 				Address:    "AQ31oFNS9HYtXYk4tAzajrZYCRQQRBBzcr",
// 				InitPos:    30000,
// 			},
// 			{
// 				Index:      4,
// 				PeerPubkey: "020a3da2aa5a5ccead145088e60a2897ebf1834a01445ff71fc6dc651b9a078709",
// 				Address:    "AboGDWUUkP2J6msQh9GhXgPvupdbzmkE2G",
// 				InitPos:    40000,
// 			},
// 			{
// 				Index:      5,
// 				PeerPubkey: "02cb0af023e79618fefa1dcc75cce869f67a31c1e316b5b2d6fd98e58984befdec",
// 				Address:    "Ab6YW8QKuVD3FVi3sUL1vu5fd2ge2KLVjC",
// 				InitPos:    30000,
// 			},
// 			{
// 				Index:      6,
// 				PeerPubkey: "027e910e3ae3605307d25617a2cf5b04605d7d7d0e6ad87a92fc8437fe9186e9fc",
// 				Address:    "ATErqzubJQxVtbHBVJJUfAahtqDmqt8Nkc",
// 				InitPos:    20000,
// 			},
// 			{
// 				Index:      7,
// 				PeerPubkey: "02d7c46aaa3badb0f0f36434c6c6c5b823e11266c6d0febd63e0fb097d3eb7d5ac",
// 				Address:    "Aafm5VJy8eiquJ7cthfvw2fPxsrwjT2Ucq",
// 				InitPos:    10000,
// 			},
// 		},
// 	},
// 	DBFT: &DBFTConfig{},
// 	SOLO: &SOLOConfig{},
// }

var MainNetConfig = &GenesisConfig{
	SeedList: []string{
		"dapp1.tesra.me:25766",
		"dapp2.tesra.me:25766",
		"dapp3.tesra.me:25766",
		"dapp4.tesra.me:25766",
		"www.tesra.me:25766",
		"explorer.tesra.me:25766",
		"explorer.tesra.dev:25766"},
	ConsensusType: CONSENSUS_TYPE_VBFT,
	VBFT: &VBFTConfig{
		N:                    7,
		C:                    2,
		K:                    7,
		L:                    112,
		BlockMsgDelay:        10000,
		HashMsgDelay:         10000,
		PeerHandshakeTimeout: 10,
		MaxBlockChangeView:   120000,
		AdminTstID:           "did:tst:AdjfcJgwru2FD8kotCPvLDXYzRjqFjc9Tb",
		MinInitStake:         100000,
		VrfValue:             "1c9810aa9822e511d5804a9c4db9dd08497c31087b0daafa34d768a3253441fa20515e2f30f81741102af0ca3cefc4818fef16adb825fbaa8cad78647f3afb590e",
		VrfProof:             "c57741f934042cb8d8b087b44b161db56fc3ffd4ffb675d36cd09f83935be853d8729f3f5298d12d6fd28d45dde515a4b9d7f67682d182ba5118abf451ff1988",
		Peers: []*VBFTPeerStakeInfo{
			{
				Index:      1,
				PeerPubkey: "02e24e41966a8ba5d7a3b77175d5d45df5a452e4237c16815ff638a8bdeddceb9d",
				Address:    "AWFT2u2nZuVZdfi5veSa6vt1YGR8fauqE4",
				InitPos:    10,
			},
			{
				Index:      2,
				PeerPubkey: "02941ff96e8946ea737e4f829e184dca7033065a4dd3af6a28c511052cdf2a267a",
				Address:    "ANxoZ4h8iEoEApXrVDCcrzYTi1GjDnQsVm",
			},
			{
				Index:      3,
				PeerPubkey: "03ced9dec9b78f436cd1e3d96888074d209e22efabca7c46475cc0e9e138f53a83",
				Address:    "AKhVMSSNT62ZJWWeVvw4SLY9wDXUxzcDsk",
			},
			{
				Index:      4,
				PeerPubkey: "031aa3cade9e75bf2171a3e89a0ae326c3f36d821322caf2f5009dcb21e9802e35",
				Address:    "AeaJu9bHU4HG5xCnehjmdo9fazLivQ7o1x",
			},
			{
				Index:      5,
				PeerPubkey: "02c3c12fc599306db62c0edcb3e13806274ba3eadb7316b2a5f70b3fba9a5405cd",
				Address:    "AarzRfeck21CWP3B5hqeCvmKd7W8pDbbXc",
			},
			{
				Index:      6,
				PeerPubkey: "03cbc34b9306d466b467bd1dcdfa726f378c2807480610d52d05149e97fb08517b",
				Address:    "APppNcRbGU5nPG2gsPj9SpZzgHGqfFurQS",
			},
			{
				Index:      7,
				PeerPubkey: "0208810415c5cf8751ec042ddf3fef731231778826f4952d8609581a1c3335e67b",
				Address:    "ATJKpmsvh3bUsafXhrvA19nNm8ZQgG6jsU",
			},
		},
	},
	DBFT: &DBFTConfig{},
	SOLO: &SOLOConfig{},
}

var DefConfig = NewTesranodeConfig()

type GenesisConfig struct {
	SeedList      []string
	ConsensusType string
	VBFT          *VBFTConfig
	DBFT          *DBFTConfig
	SOLO          *SOLOConfig
}

func NewGenesisConfig() *GenesisConfig {
	return &GenesisConfig{
		SeedList:      make([]string, 0),
		ConsensusType: CONSENSUS_TYPE_DBFT,
		VBFT:          &VBFTConfig{},
		DBFT:          &DBFTConfig{},
		SOLO:          &SOLOConfig{},
	}
}

//
// VBFT genesis config, from local config file
//
type VBFTConfig struct {
	N                    uint32               `json:"n"` // network size
	C                    uint32               `json:"c"` // consensus quorum
	K                    uint32               `json:"k"`
	L                    uint32               `json:"l"`
	BlockMsgDelay        uint32               `json:"block_msg_delay"`
	HashMsgDelay         uint32               `json:"hash_msg_delay"`
	PeerHandshakeTimeout uint32               `json:"peer_handshake_timeout"`
	MaxBlockChangeView   uint32               `json:"max_block_change_view"`
	MinInitStake         uint32               `json:"min_init_stake"`
	AdminTstID           string               `json:"admin_tst_id"`
	VrfValue             string               `json:"vrf_value"`
	VrfProof             string               `json:"vrf_proof"`
	Peers                []*VBFTPeerStakeInfo `json:"peers"`
}

func (self *VBFTConfig) Serialization(sink *common.ZeroCopySink) error {
	sink.WriteUint32(self.N)
	sink.WriteUint32(self.C)
	sink.WriteUint32(self.K)
	sink.WriteUint32(self.L)
	sink.WriteUint32(self.BlockMsgDelay)
	sink.WriteUint32(self.HashMsgDelay)
	sink.WriteUint32(self.PeerHandshakeTimeout)
	sink.WriteUint32(self.MaxBlockChangeView)
	sink.WriteUint32(self.MinInitStake)
	sink.WriteString(self.AdminTstID)
	sink.WriteString(self.VrfValue)
	sink.WriteString(self.VrfProof)
	sink.WriteVarUint(uint64(len(self.Peers)))
	for _, peer := range self.Peers {
		if err := peer.Serialization(sink); err != nil {
			return err
		}
	}

	return nil
}

func (this *VBFTConfig) Deserialization(source *common.ZeroCopySource) error {
	n, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize n error!")
	}
	c, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize c error!")
	}
	k, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize k error!")
	}
	l, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize l error!")
	}
	blockMsgDelay, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize blockMsgDelay error!")
	}
	hashMsgDelay, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize hashMsgDelay error!")
	}
	peerHandshakeTimeout, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize peerHandshakeTimeout error!")
	}
	maxBlockChangeView, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize maxBlockChangeView error!")
	}
	minInitStake, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize minInitStake error!")
	}
	adminTstID, _, irregular, eof := source.NextString()
	if irregular {
		return errors.NewDetailErr(common.ErrIrregularData, errors.ErrNoCode, "serialization.ReadString, deserialize adminTstID error!")
	}
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadString, deserialize adminTstID error!")
	}
	vrfValue, _, irregular, eof := source.NextString()
	if irregular {
		return errors.NewDetailErr(common.ErrIrregularData, errors.ErrNoCode, "serialization.ReadString, deserialize vrfValue error!")
	}
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadString, deserialize vrfValue error!")
	}
	vrfProof, _, irregular, eof := source.NextString()
	if irregular {
		return errors.NewDetailErr(common.ErrIrregularData, errors.ErrNoCode, "serialization.ReadString, deserialize vrfProof error!")
	}
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadString, deserialize vrfProof error!")
	}
	length, _, irregular, eof := source.NextVarUint()
	if irregular {
		return errors.NewDetailErr(common.ErrIrregularData, errors.ErrNoCode, "serialization.ReadVarUint, deserialize peer length error!")
	}
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadVarUint, deserialize peer length error!")
	}
	peers := make([]*VBFTPeerStakeInfo, 0)
	for i := 0; uint64(i) < length; i++ {
		peer := new(VBFTPeerStakeInfo)
		err := peer.Deserialization(source)
		if err != nil {
			return errors.NewDetailErr(err, errors.ErrNoCode, "deserialize peer error!")
		}
		peers = append(peers, peer)
	}
	this.N = n
	this.C = c
	this.K = k
	this.L = l
	this.BlockMsgDelay = blockMsgDelay
	this.HashMsgDelay = hashMsgDelay
	this.PeerHandshakeTimeout = peerHandshakeTimeout
	this.MaxBlockChangeView = maxBlockChangeView
	this.MinInitStake = minInitStake
	this.AdminTstID = adminTstID
	this.VrfValue = vrfValue
	this.VrfProof = vrfProof
	this.Peers = peers
	return nil
}

type VBFTPeerStakeInfo struct {
	Index      uint32 `json:"index"`
	PeerPubkey string `json:"peerPubkey"`
	Address    string `json:"address"`
	InitPos    uint64 `json:"initPos"`
}

func (this *VBFTPeerStakeInfo) Serialization(sink *common.ZeroCopySink) error {
	sink.WriteUint32(this.Index)
	sink.WriteString(this.PeerPubkey)

	address, err := common.AddressFromBase58(this.Address)
	if err != nil {
		return fmt.Errorf("serialize VBFTPeerStackInfo error: %v", err)
	}
	address.Serialization(sink)
	sink.WriteUint64(this.InitPos)
	return nil
}

func (this *VBFTPeerStakeInfo) Deserialization(source *common.ZeroCopySource) error {
	index, eof := source.NextUint32()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize index error!")
	}
	peerPubkey, _, irregular, eof := source.NextString()
	if irregular {
		return errors.NewDetailErr(common.ErrIrregularData, errors.ErrNoCode, "serialization.ReadUint32, deserialize peerPubkey error!")
	}
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize peerPubkey error!")
	}
	address := new(common.Address)
	err := address.Deserialization(source)
	if err != nil {
		return errors.NewDetailErr(err, errors.ErrNoCode, "address.Deserialize, deserialize address error!")
	}
	initPos, eof := source.NextUint64()
	if eof {
		return errors.NewDetailErr(io.ErrUnexpectedEOF, errors.ErrNoCode, "serialization.ReadUint32, deserialize initPos error!")
	}
	this.Index = index
	this.PeerPubkey = peerPubkey
	this.Address = address.ToBase58()
	this.InitPos = initPos
	return nil
}

type DBFTConfig struct {
	GenBlockTime uint
	Bookkeepers  []string
}

type SOLOConfig struct {
	GenBlockTime uint
	Bookkeepers  []string
}

type CommonConfig struct {
	LogLevel       uint
	NodeType       string
	EnableEventLog bool
	SystemFee      map[string]int64
	GasLimit       uint64
	GasPrice       uint64
	DataDir        string
}

type ConsensusConfig struct {
	EnableConsensus bool
	MaxTxInBlock    uint
}

type P2PRsvConfig struct {
	ReservedPeers []string `json:"reserved"`
	MaskPeers     []string `json:"mask"`
}

type P2PNodeConfig struct {
	ReservedPeersOnly         bool
	ReservedCfg               *P2PRsvConfig
	NetworkMagic              uint32
	NetworkId                 uint32
	NetworkName               string
	NodePort                  uint
	IsTLS                     bool
	CertPath                  string
	KeyPath                   string
	CAPath                    string
	HttpInfoPort              uint
	MaxHdrSyncReqs            uint
	MaxConnInBound            uint
	MaxConnOutBound           uint
	MaxConnInBoundForSingleIP uint
}

type RpcConfig struct {
	EnableHttpJsonRpc bool
	HttpJsonPort      uint
	HttpLocalPort     uint
}

type RestfulConfig struct {
	EnableHttpRestful  bool
	HttpRestPort       uint
	HttpMaxConnections uint
	HttpCertPath       string
	HttpKeyPath        string
}

type WebSocketConfig struct {
	EnableHttpWs bool
	HttpWsPort   uint
	HttpCertPath string
	HttpKeyPath  string
}

type TesranodeConfig struct {
	Genesis   *GenesisConfig
	Common    *CommonConfig
	Consensus *ConsensusConfig
	P2PNode   *P2PNodeConfig
	Rpc       *RpcConfig
	Restful   *RestfulConfig
	Ws        *WebSocketConfig
}

func NewTesranodeConfig() *TesranodeConfig {
	return &TesranodeConfig{
		Genesis: MainNetConfig,
		Common: &CommonConfig{
			LogLevel:       DEFAULT_LOG_LEVEL,
			EnableEventLog: DEFAULT_ENABLE_EVENT_LOG,
			SystemFee:      make(map[string]int64),
			GasLimit:       DEFAULT_GAS_LIMIT,
			DataDir:        DEFAULT_DATA_DIR,
		},
		Consensus: &ConsensusConfig{
			EnableConsensus: true,
			MaxTxInBlock:    DEFAULT_MAX_TX_IN_BLOCK,
		},
		P2PNode: &P2PNodeConfig{
			ReservedCfg:               &P2PRsvConfig{},
			ReservedPeersOnly:         false,
			NetworkId:                 NETWORK_ID_MAIN_NET,
			NetworkName:               GetNetworkName(NETWORK_ID_MAIN_NET),
			NetworkMagic:              GetNetworkMagic(NETWORK_ID_MAIN_NET),
			NodePort:                  DEFAULT_NODE_PORT,
			IsTLS:                     false,
			CertPath:                  "",
			KeyPath:                   "",
			CAPath:                    "",
			HttpInfoPort:              DEFAULT_HTTP_INFO_PORT,
			MaxHdrSyncReqs:            DEFAULT_MAX_SYNC_HEADER,
			MaxConnInBound:            DEFAULT_MAX_CONN_IN_BOUND,
			MaxConnOutBound:           DEFAULT_MAX_CONN_OUT_BOUND,
			MaxConnInBoundForSingleIP: DEFAULT_MAX_CONN_IN_BOUND_FOR_SINGLE_IP,
		},
		Rpc: &RpcConfig{
			EnableHttpJsonRpc: true,
			HttpJsonPort:      DEFAULT_RPC_PORT,
			HttpLocalPort:     DEFAULT_RPC_LOCAL_PORT,
		},
		Restful: &RestfulConfig{
			EnableHttpRestful: true,
			HttpRestPort:      DEFAULT_REST_PORT,
		},
		Ws: &WebSocketConfig{
			EnableHttpWs: true,
			HttpWsPort:   DEFAULT_WS_PORT,
		},
	}
}

func (this *TesranodeConfig) GetBookkeepers() ([]keypair.PublicKey, error) {
	var bookKeepers []string
	switch this.Genesis.ConsensusType {
	case CONSENSUS_TYPE_VBFT:
		for _, peer := range this.Genesis.VBFT.Peers {
			bookKeepers = append(bookKeepers, peer.PeerPubkey)
		}
	case CONSENSUS_TYPE_DBFT:
		bookKeepers = this.Genesis.DBFT.Bookkeepers
	case CONSENSUS_TYPE_SOLO:
		bookKeepers = this.Genesis.SOLO.Bookkeepers
	default:
		return nil, fmt.Errorf("Does not support %s consensus", this.Genesis.ConsensusType)
	}

	pubKeys := make([]keypair.PublicKey, 0, len(bookKeepers))
	for _, key := range bookKeepers {
		pubKey, err := hex.DecodeString(key)
		k, err := keypair.DeserializePublicKey(pubKey)
		if err != nil {
			return nil, fmt.Errorf("Incorrectly book keepers key:%s", key)
		}
		pubKeys = append(pubKeys, k)
	}
	keypair.SortPublicKeys(pubKeys)
	return pubKeys, nil
}

func (this *TesranodeConfig) GetDefaultNetworkId() (uint32, error) {
	defaultNetworkId, err := this.getDefNetworkIDFromGenesisConfig(this.Genesis)
	if err != nil {
		return 0, err
	}
	mainNetId, err := this.getDefNetworkIDFromGenesisConfig(MainNetConfig)
	if err != nil {
		return 0, err
	}
	polaridId, err := this.getDefNetworkIDFromGenesisConfig(ScorpioConfig)
	if err != nil {
		return 0, err
	}
	switch defaultNetworkId {
	case mainNetId:
		return NETWORK_ID_MAIN_NET, nil
	case polaridId:
		return NETWORK_ID_SCORPIO_NET, nil
	}
	return defaultNetworkId, nil
}

func (this *TesranodeConfig) getDefNetworkIDFromGenesisConfig(genCfg *GenesisConfig) (uint32, error) {
	var configData []byte
	var err error
	switch this.Genesis.ConsensusType {
	case CONSENSUS_TYPE_VBFT:
		configData, err = json.Marshal(genCfg.VBFT)
	case CONSENSUS_TYPE_DBFT:
		configData, err = json.Marshal(genCfg.DBFT)
	case CONSENSUS_TYPE_SOLO:
		return NETWORK_ID_SOLO_NET, nil
	default:
		return 0, fmt.Errorf("unknown consensus type:%s", this.Genesis.ConsensusType)
	}
	if err != nil {
		return 0, fmt.Errorf("json.Marshal error:%s", err)
	}
	data := sha256.Sum256(configData)
	return binary.LittleEndian.Uint32(data[0:4]), nil
}
