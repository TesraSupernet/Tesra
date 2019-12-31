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
package tstid

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/TesraSupernet/Tesra/core/states"
	"github.com/TesraSupernet/Tesra/core/store/leveldbstore"
	"github.com/TesraSupernet/Tesra/core/store/overlaydb"
	"github.com/TesraSupernet/Tesra/smartcontract/service/native"
	"github.com/TesraSupernet/Tesra/smartcontract/storage"
)

func TestDeserializeGroup(t *testing.T) {
	id0 := []byte("did:tst:ARY2ekof1eCSetcimGdjqyzUYaVDDPVWmw")
	id1 := []byte("did:tst:ASbxtSqrpmydpjqCUGDiQp2mzsfd4zFArs")
	id2 := []byte("did:tst:AGxc3cdeB6QFvmZXzWhGwzuvohNtqaaaDw")
	g_ := &Group{
		Threshold: 1,
		Members: []interface{}{
			id0,
			&Group{
				Threshold: 2,
				Members: []interface{}{
					id1,
					id2,
				},
			},
		},
	}

	data, _ := hex.DecodeString("01022a6469643a6f6e743a41525932656b6f6631654353657463696d47646a71797a5559615644445056576d775a01022a6469643a6f6e743a4153627874537172706d7964706a7143554744695170326d7a736664347a464172732a6469643a6f6e743a414778633363646542365146766d5a587a576847777a75766f684e7471616161447701020101")

	g, err := deserializeGroup(data)
	if err != nil {
		t.Fatal(err)
	}

	err = groupCmp(g_, g)
	if err != nil {
		t.Fatal(err)
	}

	memback, _ := leveldbstore.NewMemLevelDBStore()
	overlay := overlaydb.NewOverlayDB(memback)
	cache := storage.NewCacheDB(overlay)

	srvc := new(native.NativeService)
	srvc.CacheDB = cache

	key, _ := encodeID(id0)
	insertPk(srvc, key, []byte("test pk"))
	cache.Put(key, states.GenRawStorageItem([]byte{flag_valid}))
	key, _ = encodeID(id1)
	insertPk(srvc, key, []byte("test pk"))
	cache.Put(key, states.GenRawStorageItem([]byte{flag_valid}))
	key, _ = encodeID(id2)
	insertPk(srvc, key, []byte("test pk"))
	cache.Put(key, states.GenRawStorageItem([]byte{flag_valid}))

	err = validateMembers(srvc, g)
	if err != nil {
		t.Fatal("validateMembers failed")
	}
}

func groupCmp(a, b *Group) error {
	if a.Threshold != b.Threshold {
		return fmt.Errorf("error threshold")
	}
	if len(a.Members) != len(b.Members) {
		return fmt.Errorf("error number of members")
	}
	for i := 0; i < len(a.Members); i++ {
		switch ma := a.Members[i].(type) {
		case []byte:
			mb, ok := b.Members[i].([]byte)
			if !ok {
				return fmt.Errorf("m%d: type error, tst id expected", i)
			}
			if !bytes.Equal(ma, mb) {
				return fmt.Errorf("m%d: mismatched id", i)
			}
		case *Group:
			mb, ok := b.Members[i].(*Group)
			if !ok {
				return fmt.Errorf("m%d: type error, subgroup expected", i)
			}
			err := groupCmp(ma, mb)
			if err != nil {
				return fmt.Errorf("m%d:%s", i, err)
			}
		default:
			return fmt.Errorf("error type")
		}
	}
	return nil
}

func TestDeserializeGroup1(t *testing.T) {
	data, _ := hex.DecodeString("01022a6469643a6f6e743a4153627874537172706d7964706a7143554744695170326d7a736664347a464172732a6469643a6f6e743a414778633363646542365146766d5a587a576847777a75766f684e747161616144770103")
	_, err := deserializeGroup(data)
	if err == nil {
		t.Fatal("deserializeGroup should fail due to the invalid threshold")
	}
}

func TestDeserializeGroup2(t *testing.T) {
	data, _ := hex.DecodeString("010203646964086469643a6f6e740101")
	_, err := deserializeGroup(data)
	if err == nil {
		t.Fatal("deserializeGroup should fail due to invalid member data")
	}
}

func TestSigners(t *testing.T) {
	id0 := []byte("did:tst:ARY2ekof1eCSetcimGdjqyzUYaVDDPVWmw")
	id1 := []byte("did:tst:ASbxtSqrpmydpjqCUGDiQp2mzsfd4zFArs")
	id2 := []byte("did:tst:AGxc3cdeB6QFvmZXzWhGwzuvohNtqaaaDw")
	g := &Group{
		Threshold: 1,
		Members: []interface{}{
			id0,
			&Group{
				Threshold: 2,
				Members: []interface{}{
					id1,
					id2,
				},
			},
		},
	}

	data, _ := hex.DecodeString("01022a6469643a6f6e743a4153627874537172706d7964706a7143554744695170326d7a736664347a4641727301012a6469643a6f6e743a414778633363646542365146766d5a587a576847777a75766f684e747161616144770101")
	signers, err := deserializeSigners(data)
	if err != nil {
		t.Fatal(err)
	}

	if !verifyThreshold(g, signers) {
		t.Fatal("verifyThreshold failed")
	}

	data, _ = hex.DecodeString("01012a6469643a6f6e743a4153627874537172706d7964706a7143554744695170326d7a736664347a464172730101")
	signers, err = deserializeSigners(data)
	if err != nil {
		t.Fatal(err)
	}

	if verifyThreshold(g, signers) {
		t.Fatal("verifyThreshold should fail")
	}
}
