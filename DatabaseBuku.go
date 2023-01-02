package main

import "fmt"

type buku struct {
	judul, penulis string
	tahun          int
}

type TabBuku [5]buku

func tambahBuku(kardus *TabBuku, atas *int) {
	if *atas == 4 {
		fmt.Println("kardus penuh")
	} else {
		*atas++
		fmt.Scan(&kardus[*atas].judul, &kardus[*atas].penulis, &kardus[*atas].tahun)
	}
}

func ambilBuku(kardus *TabBuku, atas *int, ambil *buku) {
	if *atas == -1 {
		fmt.Println("kardus kosong")
	} else {
		*ambil = kardus[*atas]
		*atas--
	}
}

func cariBuku(kardus *TabBuku, atas *int, X string) {
	var found bool = false
	var ambil buku
	for !found && *atas != -1 {
		ambilBuku(kardus, atas, &ambil)
		fmt.Println("Judul buku yang dikeluarkan:", ambil.judul)
		found = ambil.judul == X
	}

	if found {
		fmt.Println("KETEMU")
	} else {
		fmt.Println("TIDAK KETEMU")
	}
}
func main() {
	var kardus TabBuku
	var atas int

	//1. buat kardus kosong
	atas = -1

	//2. tambah 4 buku
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)

	//3. Cari buku C
	cariBuku(&kardus, &atas, "C")

	//4. tambah buku sampai kardus penuh
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)

	//5. Cari buku yang tidak ada di kardus
	cariBuku(&kardus, &atas, "ALPRO")

}
