package storage

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// FileExists - checks if the file exist
func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsExist(err) {
		return true
	}

	return false
}

// CreateFolder - creates a folder by a given name
func CreateFolder(folderName string) {
	os.MkdirAll(folderName, os.ModePerm)
}

// CreateFile - creates a file if it does not exist
func CreateFile(fileName string) *os.File {
	fp, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())

		return nil
	}

	return fp
}

// OpenFile - opens file for read and write
func OpenFile(fileName string) *os.File {
	fp, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, os.ModeAppend)

	if err != nil {
		fmt.Println(err.Error())

		return nil
	}

	return fp
}

// ReadUInt16 - Reads 2 bytes starting at position {from}
func ReadUInt16(fp *os.File, from int64) uint16 {
	fp.Seek(from, 0)
	buff := make([]byte, 2)

	io.ReadFull(fp, buff)

	var result = binary.BigEndian.Uint16(buff)

	return result
}

// ReadUInt32 - Reads 4 bytes starting at position {from}
func ReadUInt32(fp *os.File, from int64) uint32 {
	fp.Seek(from, 0)
	buff := make([]byte, 4)

	io.ReadFull(fp, buff)

	var result = binary.BigEndian.Uint32(buff)

	return result
}

// ReadString - reads a string starting from and continuing to
func ReadString(fp *os.File, from int64, to int64) string {
	fp.Seek(from, 0)

	buff := make([]byte, to)
	io.ReadFull(fp, buff)

	var result = string(buff)

	return result
}
