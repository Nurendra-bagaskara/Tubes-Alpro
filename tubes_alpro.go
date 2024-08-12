package main

import "fmt"

const NMAX int = 10

type DataPeserta struct {
	nama, alamat           string
	idPeserta, umur, paket int
}
type Paket struct {
	destinasi              string
	idPaket, durasi, harga int
	fasilitas              [NMAX]string
}

type tabData [NMAX]DataPeserta
type tabPaket [NMAX]Paket

func main() {
	var dataPeserta, inputPeserta tabData
	var dataPaket, inputPaket tabPaket
	var menu, aktivitas, nPeserta, nPaket int
	nPeserta = 0

	// data dummy peserta
	dataPeserta[0] = DataPeserta{
		idPeserta: 0,
		nama:      "Bagas",
		alamat:    "Jombang",
		umur:      25,
		paket:     1,
	}
	dataPeserta[1] = DataPeserta{
		idPeserta: 1,
		nama:      "Vega",
		alamat:    "Gubeng",
		umur:      27,
		paket:     2,
	}
	dataPeserta[2] = DataPeserta{
		idPeserta: 2,
		nama:      "Ajenk",
		alamat:    "Bandung",
		umur:      35,
		paket:     3,
	}
	dataPeserta[3] = DataPeserta{
		idPeserta: 3,
		nama:      "Fito",
		alamat:    "Malang",
		umur:      22,
		paket:     1,
	}
	dataPeserta[4] = DataPeserta{
		idPeserta: 4,
		nama:      "Risa",
		alamat:    "Malang",
		umur:      27,
		paket:     3,
	}

	//untuk mendapatkan nPeserta ketika menggunakan data dummy
	if dataPeserta[0].nama != "" {
		for i := 0; i < NMAX; i++ {
			if dataPeserta[i].nama != "" {
				nPeserta = nPeserta + 1
			}
		}
	}

	//data dummy paket
	dataPaket[0] = Paket{
		idPaket:   0,
		destinasi: "Jogja",
		durasi:    5,
		harga:     7000000,
		fasilitas: [NMAX]string{"Hotel", "Transportasi", "Makan", "Guide"},
	}
	dataPaket[1] = Paket{
		idPaket:   1,
		destinasi: "Bali",
		durasi:    7,
		harga:     8000000,
		fasilitas: [NMAX]string{"Hotel", "Transportasi", "Makan", "Spa", "Guide"},
	}
	dataPaket[2] = Paket{
		idPaket:   2,
		destinasi: "Dubai",
		durasi:    12,
		harga:     15000000,
		fasilitas: [NMAX]string{"Hotel", "Transportasi", "Makan", "Safari Tour", "Guide"},
	}
	dataPaket[3] = Paket{
		idPaket:   3,
		destinasi: "Lombok",
		durasi:    5,
		harga:     5000000,
		fasilitas: [NMAX]string{"Hotel", "Transportasi", "Makan", "Snorkeling", "Guide"},
	}
	dataPaket[4] = Paket{
		idPaket:   4,
		destinasi: "Borobudur",
		durasi:    3,
		harga:     2000000,
		fasilitas: [NMAX]string{"Hotel", "Transportasi", "Breakfast", "Guide"},
	}
	//untuk mendapatkan nPaket ketika menggunakan data dummy
	if dataPaket[0].destinasi != "" {
		for i := 0; i < NMAX; i++ {
			if dataPaket[i].destinasi != "" {
				nPaket = nPaket + 1
			}
		}
	}

	menu = -1
	for menu != 0 {
		fmt.Println("===============================================================")
		fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<  PILIH MENU  >>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("_______________________________________________________________")
		fmt.Println("1. Data Peserta           ")
		fmt.Println("2. Data Paket Wisata      ")
		fmt.Println("0. Exit                   ")
		fmt.Println("===============================================================")
		fmt.Print("Masukan pilihanmu : ")
		fmt.Scan(&menu)
		fmt.Println()
		if menu == 1 {
			aktivitas = -1
			for aktivitas != 0 {
				readPeserta(dataPeserta, nPeserta)
				fmt.Println("===============================================================")
				fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<  MENU PESERTA  >>>>>>>>>>>>>>>>>>>>>>>")
				fmt.Println("_______________________________________________________________")
				fmt.Println("1. ADD Data Peserta           	                                ")
				fmt.Println("2. Update Data Peserta                                         ")
				fmt.Println("3. Delete Data Peserta                                         ")
				fmt.Println("4. Search Data Peserta                                         ")
				fmt.Println("0. back to menu                                                ")
				fmt.Println("===============================================================")
				fmt.Print("Masukan pilihanmu : ")
				fmt.Scan(&aktivitas)
				fmt.Println()
				if aktivitas == 1 {
					addPeserta(&dataPeserta, &inputPeserta, dataPaket, nPaket, &nPeserta)
				} else if aktivitas == 2 {
					updatePeserta(&dataPeserta, dataPaket, nPaket, &nPeserta)
				} else if aktivitas == 3 {
					deletePeserta(&dataPeserta, &nPeserta)
				} else if aktivitas == 4 {
					cariPesertaInt(dataPeserta, nPeserta)
				}
			}
		} else if menu == 2 {
			aktivitas = -1
			for aktivitas != 0 {
				readPaket(dataPaket, nPaket)
				fmt.Println("===============================================================")
				fmt.Println("<<<<<<<<<<<<<<<<<<<<<<  MENU PAKET WISATA  >>>>>>>>>>>>>>>>>>>>")
				fmt.Println("_______________________________________________________________")
				fmt.Println("1. ADD Data Paket                                       	    ")
				fmt.Println("2. Update Data Paket                                           ")
				fmt.Println("3. Delete Data Paket                                           ")
				fmt.Println("4. Delete Fasilitas Paket                                      ")
				fmt.Println("5. Search Data Paket                                           ")
				fmt.Println("6. Shorting Paket(Ascending)                                   ")
				fmt.Println("7. Shorting Paket(Discanding)                                  ")
				fmt.Println("0. back to menu                                                ")
				fmt.Println("===============================================================")
				fmt.Print("Masukan pilihanmu : ")
				fmt.Scan(&aktivitas)
				fmt.Println()
				if aktivitas == 1 {
					addPaketWisata(&dataPaket, &inputPaket, &nPaket)
				} else if aktivitas == 2 {
					updatePaket(&dataPaket, &nPaket)
				} else if aktivitas == 3 {
					deletePaket(&dataPaket, &nPaket)
				} else if aktivitas == 4 {
					deleteFasilitasPaket(&dataPaket, &nPaket)
				} else if aktivitas == 5 {
					cariPaket(dataPaket, nPaket)
				} else if aktivitas == 6 {
					SelectionShortPaket(&dataPaket, nPaket)
				} else if aktivitas == 7 {
					insertionShortPaket(&dataPaket, nPaket)
				}
			}

		}
	}
}

//add data peserta pada array
func addPeserta(DataPeserta, input *tabData, DataPaket tabPaket, nPaket int, nData *int) {
	var cekPaket, alert int
	cekPaket = 0
	fmt.Println("===============================================================")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<  INPUT DATA PESERTA  >>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("_______________________________________________________________")
	readPaket(DataPaket, nPaket)
	fmt.Print("Input Nama: ")
	fmt.Scan(&input[0].nama)
	fmt.Print("Input Alamat: ")
	fmt.Scan(&input[0].alamat)
	fmt.Print("Input Umur: ")
	fmt.Scan(&input[0].umur)
	// start untuk menyesuaikan data paket yang diinput
	for cekPaket != 1 {
		fmt.Print("Input Paket: ")
		fmt.Scan(&input[0].paket)
		for j := 0; j < nPaket; j++ {
			if input[0].paket == DataPaket[j].idPaket {
				cekPaket = 1
				alert = 1
			}
		}
		if alert != 1 {
			fmt.Println("Data Paket tidak ditemukan, Isi sesuai Data Paket yang ada!!")
		}
	}
	// end untuk menyesuaikan data paket yang diinput
	cek := -1
	for i := 0; i < NMAX; i++ {
		if DataPeserta[i].nama == "" && cek == -1 {
			DataPeserta[i] = input[0]
			DataPeserta[i].idPeserta = i
			*nData = *nData + 1
			cek = 0
			i = NMAX
		}
	}
	fmt.Println("Insert Data Successfully!!")
	fmt.Println("_______________________________________________________________")
}

// Read Data Peserta atau menampilkan data
func readPeserta(DataPeserta tabData, nData int) {
	fmt.Println()
	fmt.Println("_______________________________________________________________")
	fmt.Println("                         Data Peserta                          ")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("%-10s %-15s %-15s %-10s %-10s\n", "ID", "Nama", "Alamat", "Umur", "Paket")
	fmt.Println("---------------------------------------------------------------")
	for i := 0; i <= nData; i++ {
		if DataPeserta[i].nama != "" {
			fmt.Printf("%-10d %-15s %-15s %-10d %-10d\n", DataPeserta[i].idPeserta, DataPeserta[i].nama, DataPeserta[i].alamat, DataPeserta[i].umur, DataPeserta[i].paket)
		}
	}
	fmt.Println("_______________________________________________________________")
	fmt.Println()
}

//Konsep Binary Search
func findBinaryPeserta(DataPeserta tabData, nData int, PId int) int {
	kiri := 0
	kanan := nData - 1
	ketemu := -1
	for kiri <= kanan && ketemu == -1 {
		tengah := (kiri + kanan) / 2
		if PId < DataPeserta[tengah].idPeserta {
			kanan = tengah - 1
		} else if PId > DataPeserta[tengah].idPeserta {
			kiri = tengah + 1
		} else {
			ketemu = tengah
		}
	}
	return ketemu
}

// update data peserta pada array
func updatePeserta(DataPeserta *tabData, DataPaket tabPaket, nPaket int, nData *int) {
	var cekPaket, alert int
	fmt.Println("===============================================================")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<< UPDATE DATA PESERTA >>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("_______________________________________________________________")
	var pId int
	cek := 0
	for cek != -1 {
		fmt.Print("Masukan Id yang akan di Update: ")
		fmt.Scan(&pId)
		index := findBinaryPeserta(*DataPeserta, *nData, pId)
		if index == -1 {
			fmt.Println("ID tidak ditemukan, silakan coba lagi.")
		} else {
			// Update data peserta
			fmt.Print("Update Nama: ")
			fmt.Scan(&DataPeserta[index].nama)
			fmt.Print("Update Alamat: ")
			fmt.Scan(&DataPeserta[index].alamat)
			fmt.Print("Update Umur: ")
			fmt.Scan(&DataPeserta[index].umur)
			// start untuk menyesuaikan data paket yang diinput
			for cekPaket != 1 {
				fmt.Print("Update Paket: ")
				fmt.Scan(&DataPeserta[index].paket)
				for j := 0; j < nPaket; j++ {
					if DataPeserta[index].paket == DataPaket[j].idPaket {
						cekPaket = 1
						alert = 1
					}
				}
				if alert != 1 {
					fmt.Println("Data Paket tidak ditemukan, Isi sesuai Data Paket yang ada!!")
				}
			}
			// end untuk menyesuaikan data paket yang diinput
			cek = -1
		}
	}
	fmt.Println("Update Successfully!!")
	fmt.Println("_______________________________________________________________")
}

// delete data peserta pada array
func deletePeserta(DataPeserta *tabData, nData *int) {
	fmt.Println("===============================================================")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<< DELETE DATA PESERTA >>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("_______________________________________________________________")
	var pId int
	cek := 0
	//masukan pId yang mau di hapus
	for cek != -1 {
		fmt.Print("Masukan Id yang akan di delete : ")
		fmt.Scan(&pId)
		index := findBinaryPeserta(*DataPeserta, *nData, pId)
		if index == -1 {
			fmt.Println("ID tidak ditemukan, silakan coba lagi.")
		} else {
			for i := 0; i <= *nData; i++ {
				if i >= index {
					DataPeserta[i] = DataPeserta[i+1]
				}
			}
			cek = -1
		}
	}
	fmt.Println("Deleted Successfully")
	fmt.Println("_______________________________________________________________")
}

//Search Data Peserta
func cariPesertaInt(DataPeserta tabData, nData int) {
	fmt.Println("===============================================================")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<< SEARCH DATA PESERTA >>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("_______________________________________________________________")
	var x, findInt, cari int
	var findString string
	x = -1
	fmt.Println("Kategori Pencarian : ")
	fmt.Println("1. Nama ")
	fmt.Println("2. Alamat ")
	fmt.Println("3. Umur")
	fmt.Println("4. Paket")
	fmt.Print("Masukan Kategori Pencarian : ")
	fmt.Scan(&cari)
	if cari == 1 || cari == 2 {
		fmt.Print("Masukan key yang dicari : ")
		fmt.Scan(&findString)
	} else if cari == 3 || cari == 4 {
		fmt.Print("Masukan key yang dicari : ")
		fmt.Scan(&findInt)
	}
	fmt.Println("_______________________________________________________________")
	fmt.Println("                      Hasil Pencarian                          ")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("%-10s %-10s %-10s %-10s %-10s\n", "ID", "Nama", "Alamat", "Umur", "Paket")
	fmt.Println("---------------------------------------------------------------")
	for i := 0; i < nData; i++ {
		if DataPeserta[i].umur == findInt || DataPeserta[i].paket == findInt || DataPeserta[i].nama == findString || DataPeserta[i].alamat == findString {
			x = 1
			fmt.Printf("%-10d %-10s %-10s %-10d %-10d\n", DataPeserta[i].idPeserta, DataPeserta[i].nama, DataPeserta[i].alamat, DataPeserta[i].umur, DataPeserta[i].paket)
		}
	}
	if x == -1 {
		fmt.Println("               Data Peserta tidak ditemukan!!                  ")
	}
	fmt.Println("_______________________________________________________________")
}

//Menambah Data paket wisata pada array
func addPaketWisata(DataPaket, input *tabPaket, nPaket *int) {
	fmt.Println("===============================================================")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<   INPUT DATA PAKET   >>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("_______________________________________________________________")
	fmt.Print("Destinasi : ")
	fmt.Scan(&input[0].destinasi)
	fmt.Print("Durasi : ")
	fmt.Scan(&input[0].durasi)
	fmt.Print("Harga : ")
	fmt.Scan(&input[0].harga)
	// start input fasilitas
	for j := 0; j < NMAX; j++ {
		input[0].fasilitas[j] = ""
	}
	cekFasilitas := -1
	for j := 0; j < NMAX && cekFasilitas == -1; j++ {
		var fasilitas string
		fmt.Printf("Masukkan fasilitas %d (atau ketik end untuk berhenti): ", j+1)
		fmt.Scan(&fasilitas)
		if fasilitas == "end" {
			cekFasilitas = 0
		} else {
			input[0].fasilitas[j] = fasilitas
		}
	}
	// end input fasilitas
	cek := -1
	for i := 0; i < NMAX; i++ {
		if DataPaket[i].destinasi == "" && cek == -1 {
			DataPaket[i] = input[0]
			DataPaket[i].idPaket = i
			*nPaket = *nPaket + 1
			cek = 0
			i = NMAX
		}
	}
	fmt.Println("Insert Data Successfully!!")
	fmt.Println("_______________________________________________________________")
}

// //Menampilkan data Paket Wisata
func readPaket(DataPaket tabPaket, nPaket int) {
	fmt.Println()
	fmt.Println("_______________________________________________________________")
	fmt.Println("                         Data Paket                            ")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf(" %-10s %-10s %-10s %-10s %-10s\n", "ID", "Destinasi", "Durasi", "Harga", "Fasilitas")
	fmt.Println("---------------------------------------------------------------")
	for i := 0; i <= nPaket; i++ {
		if DataPaket[i].destinasi != "" {
			for j := 0; j < NMAX; j++ {
				if DataPaket[i].fasilitas[j] != "" {
					if j == 0 {
						fmt.Printf(" %-10d %-10s %-10d %-10d %-10s \n", DataPaket[i].idPaket, DataPaket[i].destinasi, DataPaket[i].durasi, DataPaket[i].harga, DataPaket[i].fasilitas[j])
					} else {
						fmt.Printf("%-45s%-10s\n", "", DataPaket[i].fasilitas[j])
					}
				}
			}
			fmt.Println()
		}
	}
	fmt.Println("_______________________________________________________________")
	fmt.Println()
}

// mencari panjang Fasilitas
func countFasilitas(fasilitas [NMAX]string) int {
	var nDataFas int
	nDataFas = 0
	for i := 0; i < NMAX; i++ {
		if fasilitas[i] != "" {
			nDataFas += 1
		}
	}
	return nDataFas
}

//Konsep Sequensial Search untuk paket
func findSequenPaket(DataPaket tabPaket, nPaket int, PId int) int {
	index := -1
	i := 0
	for index == -1 && i < nPaket {
		if DataPaket[i].idPaket == PId {
			index = i
		}
		i++
	}
	return index
}

// Konsep Sequensial Search untuk fasilitas
func findSequenFasilitas(DataPaket tabPaket, PId int, PString string) int {
	index := -1
	for j := 0; j < countFasilitas(DataPaket[PId].fasilitas) && index == -1; j++ {
		if DataPaket[PId].fasilitas[j] == PString {
			index = j
		}
	}
	// mengembalikan indeks karena dia ngecek dataPaket[pId].fasilitas[j] benar benar ada atau tidak
	return index
}

// mengupdate data Paket
func updatePaket(DataPaket *tabPaket, nPaket *int) {
	fmt.Println("===============================================================")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<   UPDATE DATA PAKET   >>>>>>>>>>>>>>>>>>>>")
	fmt.Println("_______________________________________________________________")
	var pId int
	//masukan pId yang mau di update
	cek := 0
	for cek != -1 {
		fmt.Print("Masukan Id yang akan di Update : ")
		fmt.Scan(&pId)

		index := findSequenPaket(*DataPaket, *nPaket, pId)
		if index == -1 {
			fmt.Println("ID tidak Ditemukan, Silakan Coba Lagi!!")
		} else {
			fmt.Print("Update Destinasi: ")
			fmt.Scan(&DataPaket[index].destinasi)
			fmt.Print("Update Durasi: ")
			fmt.Scan(&DataPaket[index].durasi)
			fmt.Print("Update Harga: ")
			fmt.Scan(&DataPaket[index].harga)
			// start input fasilitas
			cekFasilitas := -1
			for j := 0; j < NMAX && cekFasilitas == -1; j++ {
				var fasilitas string
				fmt.Printf("Masukkan Update fasilitas %d (atau ketik end untuk berhenti): ", j+1)
				fmt.Scan(&fasilitas)
				if fasilitas == "end" {
					cekFasilitas = 0
				} else {
					DataPaket[index].fasilitas[j] = fasilitas
				}
			}
			// end input fasilitas
			cek = -1
		}

	}
	fmt.Println("Update Successfully!!")
	fmt.Println("_______________________________________________________________")
}

// menghapus data Paket
func deletePaket(DataPaket *tabPaket, nPaket *int) {
	fmt.Println("===============================================================")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<   DELETE DATA PAKET   >>>>>>>>>>>>>>>>>>>>")
	fmt.Println("_______________________________________________________________")
	var pId int
	//masukan pId yang mau di hapus
	cek := 0
	for cek != -1 {
		fmt.Print("Masukan Id yang akan di delete : ")
		fmt.Scan(&pId)
		index := findSequenPaket(*DataPaket, *nPaket, pId)
		if index == -1 {
			fmt.Println("ID Tidak Ditemukan, Silakan Coba Lagi!!")
		} else {
			for i := 0; i <= *nPaket; i++ {
				if i >= index {
					DataPaket[i] = DataPaket[i+1]
				}
			}
			cek = -1
		}
	}
	fmt.Println("Delete Successfully!!")
	fmt.Println("_______________________________________________________________")
}

func deleteFasilitasPaket(DataPaket *tabPaket, nPaket *int) {
	fmt.Println("===============================================================")
	fmt.Println("<<<<<<<<<<<<<<<<<<   DELETE FASILITAS PAKET  >>>>>>>>>>>>>>>>>>")
	fmt.Println("_______________________________________________________________")
	var pId int
	var pString string
	//masukan pId yang mau di hapus
	cek := 0
	for cek != -1 {
		fmt.Print("Masukan Id yang akan di delete : ")
		fmt.Scan(&pId)
		index := findSequenPaket(*DataPaket, *nPaket, pId)
		if index == -1 {
			fmt.Println("ID Tidak Ditemukan, Silakan Coba Lagi!!")
		} else {
			for cek != -1 {
				fmt.Print("Masukan nama fasilitas yang akan di delete : ")
				fmt.Scan(&pString)
				indexFasilitas := findSequenFasilitas(*DataPaket, pId, pString)
				if indexFasilitas == -1 {
					fmt.Println("Fasilitas di Id", pId, "Tidak Ditemukan, Silakan Coba Lagi!!")
				} else {
					for i := 0; i <= *nPaket; i++ {
						if i >= indexFasilitas {
							DataPaket[index].fasilitas[i] = DataPaket[index].fasilitas[i+1]
						}
					}
					cek = -1
				}
			}
		}
	}
	fmt.Println("Delete Successfully!!")
	fmt.Println("_______________________________________________________________")
}

// Search data Paket
func cariPaket(DataPaket tabPaket, nPaket int) {
	fmt.Println("===============================================================")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<   SEARCH DATA PAKET   >>>>>>>>>>>>>>>>>>>>")
	fmt.Println("_______________________________________________________________")
	var x, findInt, cari int
	var findString string
	fmt.Println("Kategori Pencarian : ")
	fmt.Println("1. Destinasi ")
	fmt.Println("2. Durasi ")
	fmt.Println("3. Harga")
	fmt.Println("4. Fasilitas")
	fmt.Print("Masukan Kategori Pencarian : ")
	fmt.Scan(&cari)
	fmt.Print("Masukan Key yang dicari : ")
	if cari == 1 || cari == 4 {
		fmt.Scan(&findString)
	} else if cari == 2 || cari == 3 {
		fmt.Scan(&findInt)
	}
	fmt.Println("_______________________________________________________________")
	fmt.Println("                      Hasil Pencarian                          ")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("%-10s %-10s %-10s %-10s %-10s\n", "ID", "Destinasi", "Durasi", "Harga", "Fasilitas")
	fmt.Println("---------------------------------------------------------------")
	ketemu := false
	for i := 0; i < nPaket; i++ {
		cekData := false
		if cari == 1 && DataPaket[i].destinasi == findString {
			cekData = true
		} else if cari == 2 && DataPaket[i].durasi == findInt {
			cekData = true
		} else if cari == 3 && DataPaket[i].harga == findInt {
			cekData = true
		} else if cari == 4 {
			for j := 0; j < countFasilitas(DataPaket[i].fasilitas); j++ {
				if DataPaket[i].fasilitas[j] == findString {
					cekData = true
					x = j
					j = 999
				}
			}
		}
		if cekData {
			ketemu = true
			if cari == 4 {
				fmt.Printf(" %-10d %-10s %-10d %-10d %-10s \n", DataPaket[i].idPaket, DataPaket[i].destinasi, DataPaket[i].durasi, DataPaket[i].harga, DataPaket[i].fasilitas[x])
			} else if cari == 1 || cari == 2 || cari == 3 {
				for j := 0; j < NMAX; j++ {
					if DataPaket[i].fasilitas[j] != "" {
						if j == 0 {
							fmt.Printf(" %-10d %-10s %-10d %-10d %-10s \n", DataPaket[i].idPaket, DataPaket[i].destinasi, DataPaket[i].durasi, DataPaket[i].harga, DataPaket[i].fasilitas[j])
						} else {
							fmt.Printf("%-45s%-10s\n", "", DataPaket[i].fasilitas[j])
						}
					}
				}
				fmt.Println()
			}

		}
	}
	if !ketemu {
		fmt.Println("                  Data Paket tidak ditemukan!!                 ")
	}
	fmt.Println("Search Successfully!!")
	fmt.Println("_______________________________________________________________")
}

// ascending
// Mengurutkan data Paket Dengan Selection Short
func SelectionShortPaket(DataPaket *tabPaket, nPaket int) {
	var idx, cari int
	var temp Paket
	fmt.Println("Kategori Pencarian ")
	fmt.Println("1. Destinasi ")
	fmt.Println("2. Durasi")
	fmt.Println("3. Harga")
	fmt.Println("4. fasilitas")
	fmt.Print("Masukan Kategori Pencarian : ")
	fmt.Scan(&cari)
	for pass := 1; pass < nPaket; pass++ {
		idx = pass - 1
		for i := pass; i < nPaket; i++ {
			if cari == 1 {
				if DataPaket[idx].destinasi > DataPaket[i].destinasi {
					idx = i
				}
			} else if cari == 2 {
				if DataPaket[idx].durasi > DataPaket[i].durasi {
					idx = i
				}
			} else if cari == 3 {
				if DataPaket[idx].harga > DataPaket[i].harga {
					idx = i
				}
			} else if cari == 4 {
				if countFasilitas(DataPaket[idx].fasilitas) > countFasilitas(DataPaket[i].fasilitas) || (countFasilitas(DataPaket[idx].fasilitas) == countFasilitas(DataPaket[i].fasilitas) && DataPaket[idx].harga > DataPaket[i].harga) {
					idx = i
				}
			}
		}
		temp = DataPaket[pass-1]
		DataPaket[pass-1] = DataPaket[idx]
		DataPaket[idx] = temp
	}
	fmt.Println("Hasil Pengurutan Data Dengan Selection Short Paket :")
}

// discanding
// Mengurutkan data Paket Dengan insertion Short secara DESC dengan kategori tertentu
func insertionShortPaket(DataPaket *tabPaket, nPaket int) {
	var i, cari int
	var temp Paket
	fmt.Println("Kategori Pencarian ")
	fmt.Println("1. Destinasi ")
	fmt.Println("2. Durasi")
	fmt.Println("3. Harga")
	fmt.Println("4. Fasilitas")
	fmt.Print("Masukan Kategori Pencarian : ")
	fmt.Scan(&cari)
	for pass := 1; pass < nPaket; pass++ {
		temp = DataPaket[pass]
		i = pass
		if cari == 1 {
			for i > 0 && temp.destinasi > DataPaket[i-1].destinasi {
				DataPaket[i] = DataPaket[i-1]
				i--
			}
		} else if cari == 2 {
			for i > 0 && temp.durasi > DataPaket[i-1].durasi {
				DataPaket[i] = DataPaket[i-1]
				i--
			}
		} else if cari == 3 {
			for i > 0 && temp.harga > DataPaket[i-1].harga {
				DataPaket[i] = DataPaket[i-1]
				i--
			}
		} else if cari == 4 {
			for i > 0 && (countFasilitas(temp.fasilitas) > countFasilitas(DataPaket[i-1].fasilitas) ||
				(countFasilitas(temp.fasilitas) == countFasilitas(DataPaket[i-1].fasilitas) && temp.harga > DataPaket[i-1].harga)) {
				DataPaket[i] = DataPaket[i-1]
				i--
			}
		}
		DataPaket[i] = temp
	}
	fmt.Println("Hasil Pengurutan Data Dengan Insertion Short Paket :")
}
