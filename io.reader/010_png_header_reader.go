package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Run with: go run io.reader/010_png_header_reader.go
func main() {
	data := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, // PNG sig
		0x00, 0x00, 0x00, 0x0D, // len 13
		'I', 'H', 'D', 'R', // type
		0x00, 0x00, 0x02, 0x00, // width 512
		0x00, 0x00, 0x01, 0x00, // height 256
		0x08, 0x02, 0x00, 0x00, 0x00, // other IHDR fields
		0xDE, 0xAD, 0xBE, 0xEF, // CRC (fake)
	}
	r := bytes.NewReader(data)

	sig := make([]byte, 8)
	if _, err := r.Read(sig); err != nil { panic(err) }
	fmt.Printf("sig: %x\n", sig)

	var length uint32
	binary.Read(r, binary.BigEndian, &length)
	fmt.Println("len:", length)

	typ := make([]byte, 4)
	r.Read(typ)
	fmt.Println("type:", string(typ))

	var w, h uint32
	binary.Read(r, binary.BigEndian, &w)
	binary.Read(r, binary.BigEndian, &h)
	fmt.Println("wxh:", w, h)

	// skip rest of IHDR data
	r.Seek(int64(length-8), io.SeekCurrent)

	var crc uint32
	binary.Read(r, binary.BigEndian, &crc)
	fmt.Printf("crc: 0x%X\n", crc)
}

