// MAP
// Map handles the index file of a database.
// A map contains the following information:
// 1. Database name
// 2. Version
// 3. Number of tables in the database
// 4. Structure of table with { Id, Name }
// 5. Pages with indecies poiting to the offset of the files in the actual table files
// Layout (Sequential)
// Header:
// 		[Magic (4 bytes)] [Version (2 bytes)] [TableCount (4 bytes)]
// Tables:
// 		[TableName (128 bytes)] [TableId (2 bytes)]
//		[TableName (128 bytes)] [TableId (2 bytes)]
// 		...
// Indecies:
//		[Table1 (2048 bytes)
// 			[Count (2 bytes)]
//			[Hash(4)] [Index(4)], [Hash(4)] [Index(4)], [Hash(4)] [Index(4)] ... (Max 255 kvps)
//		]
//		[Table2 (2048 bytes)
// 			[Count (2 bytes)]
//			[Hash(4)] [Index(4)], [Hash(4)] [Index(4)], [Hash(4)] [Index(4)] ... (Max 255 kvps)
//		]
// 		...

package db

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Vect0rZ/gibkv/storage"
)

// Table - Table info structure
type Table struct {
	ID        uint16
	TableName string
}

// Map - database map structure
type Map struct {
	DBName     string
	Version    uint16
	TableCount uint32
	Tables     []Table
}

var magic = []byte{1, 0, 0, 1}
var magic_tab = []byte{1, 0, 0, 2}
var version = []byte{0, 1}

// OpenIndexMap - reads the index map and returns a struct
func OpenIndexMap(dbname string) Map {
	fp := storage.OpenFile(dbname + "/kvmap")
	if fp == nil {
		panic("Cannot open kvmap")
	}

	var ver = storage.ReadUInt16(fp, 4)
	var tc = storage.ReadUInt32(fp, 6)

	var offset int64 = 10
	var tables = make([]Table, 0)
	for i := 0; i < int(tc); i++ {
		var name = storage.ReadString(fp, offset, 128)
		offset += 128

		var id = storage.ReadUInt16(fp, offset)
		offset += 2

		t := Table{id, name}
		tables = append(tables, t)
	}

	return Map{dbname, ver, tc, tables}
}

// CreateIndexMap - creates the initial index map for the database
func CreateIndexMap(dbname string) Map {
	if storage.FileExists(dbname + "/kvmap") {
		panic("Database already exist")
	}

	storage.CreateFolder(dbname)
	fp := storage.CreateFile(dbname + "/kvmap")
	defer fp.Close()

	if fp == nil {
		panic("Cannot create kvmap")
	}

	var data = make([]byte, 10)
	copy(data[0:4], magic)
	copy(data[4:6], version)
	copy(data[6:10], []byte{0, 0, 0, 0})

	fp.Write(data)
	fp.Sync()

	var ver = binary.BigEndian.Uint16(version)

	return Map{dbname, ver, 0, []Table{}}
}

// MapNewTable - maps the new table into the index map
func (m *Map) MapNewTable(tableName string) bool {
	fp := storage.OpenFile(m.DBName + "/kvmap")
	defer fp.Close()

	if fp == nil {
		panic("Cannot open kvmap")
	}

	fp.Seek(6, 0)
	tcbuff := make([]byte, 4)
	io.ReadFull(fp, tcbuff)

	// Update table count
	var tc = binary.BigEndian.Uint32(tcbuff)
	tc++
	binary.BigEndian.PutUint32(tcbuff, tc)
	fp.WriteAt(tcbuff, 6)

	// Write table data
	offset := 10 + ((tc - 1) * 130)
	var tableHeader = make([]byte, 128)
	copy(tableHeader[0:len(tableName)], []byte(tableName))

	tbid := make([]byte, 2)
	binary.BigEndian.PutUint16(tbid, uint16(tc))

	_, e := fp.WriteAt(tableHeader, int64(offset))
	if e != nil {
		fmt.Println(e)
		fmt.Println(e.Error())
	}
	offset += 128
	fp.WriteAt(tbid, int64(offset))
	offset += 2
	// Allocate 2048 buffer
	tableData := make([]byte, 2048)

	fp.WriteAt(tableData, int64(offset))

	err := fp.Sync()
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	}

	m.Tables = append(m.Tables, Table{uint16(m.TableCount + 1), tableName})
	m.TableCount++

	return true
}

// UpdateIndex - updates the index file with a newely inserted value
func (m *Map) UpdateIndex(tableName string, hash uint32, location int64) {

}
