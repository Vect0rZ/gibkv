// TABLE
// Table handles KVP data
// A table contains the following information
// 1. Id
// 2. Version
// 3. Number of records in the table
// 4. Key value pairs
// Layout (Sequential)
// Header:
// 		[Magic (4 bytes)] [Version (2 bytes)] [Id (2 bytes)]
// Data:
// 		[Length (4 bytes)] [Data (Len)]
// 		[Length (4 bytes)] [Data (Len)]
// 		[Length (4 bytes)] [Data (Len)]
//		...

package db

import (
	"encoding/binary"

	"github.com/Vect0rZ/gibkv/storage"
	"github.com/Vect0rZ/gibkv/util"
)

// CreateTable - creates a table for a given database
func (m *Map) CreateTable(tableName string) bool {
	fp := storage.CreateFile(m.DBName + "/" + tableName)
	defer fp.Close()

	var data = make([]byte, 8)
	copy(data[0:4], magic_tab)
	copy(data[4:6], version)

	idbuff := make([]byte, 2)
	binary.BigEndian.PutUint16(idbuff, uint16(m.TableCount+1))

	copy(data[6:8], idbuff)

	fp.Write(data)
	fp.Sync()

	m.MapNewTable(tableName)

	return true
}

func (m *Map) Insert(tableName, key, value string) {
	fp := storage.OpenFile(m.DBName + "/" + tableName)
	defer fp.Close()

	fi, _ := fp.Stat()

	kh := util.Hash(key)
	vd := []byte(value)

	var data = make([]byte, 4+len(vd))
	copy(data[0:4], util.UInt32ToBytes(kh))
	copy(data[4:len(vd)], vd)

	fp.Write(data)

	m.UpdateIndex(tableName, kh, fi.Size())
}
