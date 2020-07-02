package library

import (
	"fmt"
)

type mahasiswa struct {
	nama string
	umur int
}

func ShowMahasiswa() {
	var x1 mahasiswa

	x1.nama = "Zawi"
	x1.umur = 27

	x1.insertMahasiswa()
}

func (m mahasiswa) insertMahasiswa() {
	fmt.Println("Nama Saya : ", m.nama)
	fmt.Println("Umur Saya : ", m.umur)
}
