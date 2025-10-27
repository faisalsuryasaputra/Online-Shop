package main
import (
	"fmt"
	"os"
	"time"
	"runtime"
	"os/exec"
	"strings"
	"strconv"
)
const N int = 10000
type Tanggal struct{
	tahun,bulan, hari int
}
type RecType struct{
	createuser,createpass, merk, nama, jenis, id string
	hargabeli, hargajual, jumharga, stok float64
	
	tanggal Tanggal
}
type ArrType [N]RecType

var orang ArrType

var index int
var clear map[string]func() 

func urut2(orang *ArrType,index int ){
	var urut,maxindex,j int
	var tempat RecType
	urut = 1
	for urut <= index {

		maxindex = urut
		j = urut + 1
		for j <= index {
			if orang[maxindex].merk > orang[j].merk{
				maxindex = j
			}
			j++
		}
		tempat = orang[urut]
		orang[urut] = orang[maxindex]
		orang[maxindex] = tempat
		urut++
	}
}


func urut1(orang *ArrType,index int){
	var urut,inx int
	var tempat RecType
	urut = 1
	for urut <= index {
		tempat = orang[urut]
		inx  = urut - 1
		for inx >= 0 && orang[inx].nama > tempat.nama {
			orang[inx+1] = orang[inx]
			inx = inx - 1
		}
		orang[inx+1] = tempat
		urut++
	} 
}

func urut3(orang *ArrType,index int){
	var urut,inx int
	var tempat RecType
	urut = 1
	for urut <= index {
		tempat = orang[urut]
		inx  = urut - 1
		for inx >= 0 && orang[inx].hargajual > tempat.hargajual {
			orang[inx+1] = orang[inx]
			inx = inx - 1
		}
		orang[inx+1] = tempat
		urut++
	} 
}

func urut4(orang *ArrType,index int){
	var urut,maxindex,j int
	var tempat RecType
	urut = 1
	for urut <= index {

		maxindex = urut
		j = urut + 1
		for j <= index {
			if (orang[maxindex].tanggal.tahun > orang[j].tanggal.tahun) {
				maxindex = j
			} else if  (orang[maxindex].tanggal.tahun == orang[j].tanggal.tahun && orang[maxindex].tanggal.bulan > orang[j].tanggal.bulan) {
				maxindex = j
			} else if (orang[maxindex].tanggal.tahun == orang[j].tanggal.tahun && orang[maxindex].tanggal.bulan ==  orang[j].tanggal.bulan && orang[maxindex].tanggal.hari  >orang[j].tanggal.hari){
				maxindex = j
			}
			j++
		}
		tempat = orang[urut]
		orang[urut] = orang[maxindex]
		orang[maxindex] = tempat
		urut++
	}
}

func init() {
	clear = make(map[string]func()) 
	clear["linux"] = func() { 
		cmd := exec.Command("clear") 
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")  
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] 
	if ok { 
		value()  
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func validPass(pass string) bool {
	var cekhuruf, cekangka, i int
	huruf := strings.Split(pass,"")
	i = 0
	for i < len(huruf){
		if (pass[i] >= 65 && pass[i] <= 90) || (pass[i] >= 97 && pass[i] <= 122){
			cekhuruf = cekhuruf + 1
		} else if (pass[i] >= 48 && pass[i] <= 57) {
			cekangka = cekangka + 1
		}
		i++
	}
	return (cekangka + cekhuruf >= 6) && (cekangka > 0) && (cekhuruf > 0)
}

func main () {

	var login, menu, trans  int
	var username, password, merk, nama string
	var hargabeli float64
	var createuser, createpass, loginuser, loginpass string
	var num1,num2,num3, a, pakaian, buku, mainan, mrk int
	var kode, id string
	var ibar, ibel, itrans, cekid , urut, lagi int
	var hargajual, jumharga, beli, stokawal, jumlah, stok, jum float64
	var waktrans, jumlah6bln int
	var tanggal Tanggal


	for login != 4 {
		fmt.Println("...LOADING...")
		time.Sleep(2 * time.Second)
		CallClear()
	
		fmt.Println("=======================================")
		fmt.Println("-----SELAMAT DATANG DI ONLINE SHOP-----")
		fmt.Println("=======================================")
		fmt.Println("1. Login Admin")
		fmt.Println("2. Non Login User")
		fmt.Println("3. Login User")
		fmt.Println("4. Exit")
		fmt.Println("=======================================")
		fmt.Println("Masukan Nomor Pilihan Login: ")
		fmt.Scan(&login)
		fmt.Println("=======================================")
		

		if login == 1 {
			fmt.Println("...LOADING...")
			time.Sleep(2 * time.Second)
			CallClear()
			fmt.Println("===================================")
			fmt.Println("-MASUKAN USERNAME & PASSWORD ADMIN-")
			fmt.Println("===================================")
			fmt.Print("USERNAME:")
			fmt.Scan(&username)
			fmt.Print("PASSWORD:")
			fmt.Scan(&password)
			fmt.Println("===================================")
			
			
			for username != "rifkiadipramana" || password != "098765" {
				fmt.Println("--INVALID USERNAME OR PASSWORD--")
				time.Sleep(2 * time.Second)
				CallClear()
				
				fmt.Println("====================================")
				fmt.Println("-MASUKAN USERNAME & PASSWORD ADMIN-")
				fmt.Println("====================================")
				fmt.Print("USERNAME:")
				fmt.Scan(&username)
				fmt.Print("PASSWORD:")
				fmt.Scan(&password)
				fmt.Println("====================================")
				
			}
			time.Sleep(2 * time.Second)
			CallClear()

			fmt.Println("=====================================")
			fmt.Println("===========LOGIN BERHASIL============")
			fmt.Println("=====================================")
			fmt.Println("1. PENCATATAN MASTER DATA BARANG") 
			fmt.Println("2. TAMBAH STOK BARANG")
			fmt.Println("3. UBAH HARGA BELI BARANG")
			fmt.Println("4. TAMPILKAN  JUMLAH & DATA BARANG YANG TIDAK MELAKUKAN TRANSAKSI STOK LEBIH DARI 6 BULAN BERURUT BERDASARKAN TANGGAL TRANSAKSI TERAKHIR")
			fmt.Println("5. TANPILKAN JUMLAH & DATA BARANG BERDASARKAN STOK, BERURUT BERDASARKAN MERK BARANG")
			fmt.Println("=====================================")
			fmt.Println("MASUKAN NOMOR PILIHAN MENU: ")
			fmt.Scan(&menu)
			fmt.Println("=====================================")

			if menu == 1 {
				index = index + 1 
				fmt.Println("=================================")
				fmt.Println("==PENCATATAN MASTER DATA BARANG==")
				fmt.Println("===============================")
				fmt.Print("merk: ")
				fmt.Scan(&merk)
				orang[index].merk = merk
				fmt.Print("nama: ")
				fmt.Scan(&nama)
				orang[index].nama = nama

				var jenis string
				fmt.Print("Jenis Barang: ")  
				fmt.Scan(&jenis)
				jenis = strings.ToUpper(jenis)

				for jenis != "PAKAIAN" && jenis != "BUKU" && jenis != "MAINAN" {
					fmt.Println("===========================================================")
					fmt.Println("====Tidak valid, silahkan masukan jenis Barang kembali=====")
					fmt.Println("===========================================================")
					fmt.Print("Jenis Barang: ")
					fmt.Scan(&jenis)
					jenis = strings.ToUpper(jenis)
					
				}
				orang[index].jenis =  jenis
				
				fmt.Print("harga beli: ")
				fmt.Scan(&hargabeli)
				for hargabeli < 0 {
					fmt.Print("harga beli: ")
					fmt.Scan(&hargabeli)
				}
				orang[index].hargabeli = hargabeli
				
				//harga jual 
				hargajual = hargabeli + hargabeli*(0.3)
				orang[index].hargajual = hargajual

				fmt.Print("stok awal:")
				fmt.Scan(&stokawal)
				for stokawal <= 0 {
					fmt.Print("stok awal:")
					fmt.Scan(&stokawal)
				}

				orang[index].stok = stokawal
				
				//PROSES OUTPUT REKENING
		
		
				if jenis == "PAKAIAN" {
					kode = "P-"
					pakaian = 1 + pakaian
					num1 = pakaian / 100
					a    = pakaian % 100
					num2 = a / 10
					num3 = a % 10
				} else if jenis == "BUKU" {
					kode = "B-"
					buku = 1 + buku
					num1 = buku / 100
					a    = buku % 100
					num2 = a / 10
					num3 = a % 10
				} else if jenis == "MAINAN"{
					kode = "M-"
					mainan = 1 + mainan
					num1 = mainan / 100
					a    = mainan % 100
					num2 = a / 10
					num3 = a % 10
				} 
				satu := strconv.Itoa(num1)
				dua  := strconv.Itoa(num2)
				tiga := strconv.Itoa(num3)
		
				nilai := []string{kode,satu,dua,tiga}
				id = strings.Join(nilai, "")
		
				orang[index].id = id
				//proses output tanggal
				waktu    := time.Now()
				tanggal.tahun = waktu.Year()
				tanggal.bulan = int(waktu.Month())
				tanggal.hari  = waktu.Day()

			
				orang[index].tanggal = tanggal
				//output semuanya
				fmt.Println("=================================")
				fmt.Println("=========PENCATATAN SUKSES=======")
				fmt.Println("=================================")
				fmt.Println("ID: ", orang[index].id)
				fmt.Println("merk: ", orang[index].merk)
				fmt.Println("nama:", orang[index].nama)
				fmt.Println("jenis barang: ", orang[index].jenis)
				fmt.Println("harga beli: ", orang[index].hargabeli)
				fmt.Println("harga jual: ", orang[index].hargajual)
				fmt.Println("stok awal:", orang[index].stok)
				fmt.Println("tanggal transaksi:", orang[index].tanggal.tahun,"-",orang[index].tanggal.bulan,"-",orang[index].tanggal.hari)
				
				fmt.Print("MASUKAN 0 untuk keluar : ")
				fmt.Scan(&menu)
				fmt.Println("=====================================")
				if menu == 0{}
			


			} else if menu == 2 {

				fmt.Println("================================")
				fmt.Println("=======TAMBAH STOK BARANG=======")
				fmt.Println("================================")
				fmt.Println("Masukan ID barang yang akan di tambah : ")
				fmt.Scan(&id)
				ibar = 1
				for ibar < N-1 && orang[ibar].id != id {
					ibar++	
				}
				if orang[ibar].id == id {
					cekid = ibar
				} else {
					cekid = -1
				}
				for cekid == -1 {
				
					fmt.Println("Masukan ID barang yang akan di tambah : ")
					fmt.Scan(&id)
					
					ibar = 1
					for ibar < N-1 && orang[ibar].id != id {
						ibar++
					}
					if orang[ibar].id == id {
						cekid = ibar
					} else {
						cekid = -1
					}
				}

				fmt.Print("tambahkan stok barang ", orang[ibar].nama, " : ")
				fmt.Scan(&jumlah)
				for jumlah <= 0 {
					fmt.Print("tambahkan stok barang", orang[ibar].nama, " : ")
					fmt.Scan(&jumlah)
				}
				orang[ibar].stok = orang[ibar].stok + jumlah

				waktu    := time.Now()
				tanggal.tahun = waktu.Year()
				tanggal.bulan = int(waktu.Month())
				tanggal.hari  = waktu.Day()

			
				orang[index].tanggal = tanggal

				fmt.Println("=================================")
				fmt.Println("=========PENCATATAN SUKSES=======")
				fmt.Println("=================================")
				fmt.Println("ID: ", orang[ibar].id)
				fmt.Println("merk: ", orang[ibar].merk)
				fmt.Println("nama:", orang[ibar].nama)
				fmt.Println("jenis barang: ", orang[ibar].jenis)
				fmt.Println("harga beli: ", orang[ibar].hargabeli)
				fmt.Println("harga jual: ", orang[ibar].hargajual)
				fmt.Println("stok awal:", orang[ibar].stok)
				fmt.Println("tanggal transaksi:", orang[ibar].tanggal.tahun,"-",orang[ibar].tanggal.bulan,"-",orang[ibar].tanggal.hari)
				
				fmt.Print("MASUKAN 0 untuk keluar : ")
				fmt.Scan(&menu)
				fmt.Println("=====================================")
				if menu == 0{}
			

				

			} else if menu == 3 {
				fmt.Println("================================")
				fmt.Println("=====UBAH HARGA BELI BARANG=====")
				fmt.Println("================================")
				fmt.Println("Masukan ID barang : ")
				fmt.Scan(&id)
				ibar = 1

				for ibel < N-1 && orang[ibel].id != id {
					ibel++	
				}
				if orang[ibel].id == id {
					cekid = ibel
				} else {
					cekid = -1
				}
				for cekid == -1 {
				
					fmt.Println("Masukan ID barang: ")
					fmt.Scan(&id)
					
					ibar = 1
					for ibel < N-1 && orang[ibel].id != id {
						ibel++
					}
					if orang[ibel].id == id {
						cekid = ibel
					} else {
						cekid = -1
					}
				}

				fmt.Print("Ubah Harga Beli barang ", orang[ibel].nama, " : ")
				fmt.Scan(&jumharga)
				for jumharga <= 0 {
					fmt.Print("Ubah Harga Beli barang ", orang[ibel].nama, " : ")
					fmt.Scan(&jumharga)
				}
				orang[ibel].hargabeli = jumharga

				hargajual = orang[ibel].hargabeli + orang[ibel].hargabeli*(0.3)
				orang[ibel].hargajual = hargajual

				fmt.Println("=====================================")
				fmt.Println("----HARGA BELI & HARRGA JUAL BARU----")
				fmt.Println("=====================================")
				fmt.Println("harga beli: ", orang[ibel].hargabeli)
				fmt.Println("harga jual: ", orang[ibel].hargajual)
			
			

			} else if menu == 4 {

				waktu := time.Now()
				waktrans = 1
				urut4(&orang, index)
				for waktrans <= index {
					tahunT := waktu.Year()
					bulanT := int(waktu.Month())
					if (tahunT - orang[waktrans].tanggal.tahun == 0 && bulanT - orang[waktrans].tanggal.bulan >= 8) || (tahunT - orang[waktrans].tanggal.tahun == 1 && orang[waktrans].tanggal.bulan - bulanT <= 6) || (tahunT - orang[waktrans].tanggal.tahun >= 2) {
						fmt.Println("ID: ", orang[waktrans].id)
						fmt.Println("merk: ", orang[waktrans].merk)
						fmt.Println("nama:", orang[waktrans].nama)
						fmt.Println("jenis barang: ", orang[waktrans].jenis)
						fmt.Println("harga beli: ", orang[waktrans].hargabeli)
						fmt.Println("harga jual: ", orang[waktrans].hargajual)
						fmt.Println("stok awal:", orang[waktrans].stok)
						fmt.Println("tanggal transaksi:", orang[waktrans].tanggal.tahun,"-",orang[waktrans].tanggal.bulan,"-",orang[waktrans].tanggal.hari)
						jumlah6bln = jumlah6bln + 1
					}
					waktrans++
				}
				if jumlah6bln == 0 {
					fmt.Println("TIDAK DI TEMUKAN")
				}
				fmt.Print("MASUKAN 0 untuk keluar : ")
				fmt.Scan(&menu)
				fmt.Println("=====================================")
				if menu == 0{}

			} else if menu == 5 {
				fmt.Print("masukan stok yang akan ditampilkan")
				fmt.Scan(&stok)
				urut2(&orang, index)
				urut = 1
				for urut <= index {
					if stok == orang[urut].stok {
						fmt.Println("ID: ", orang[urut].id)
						fmt.Println("merk: ", orang[urut].merk)
						fmt.Println("nama:", orang[urut].nama)
						fmt.Println("stok awal:", orang[urut].stok)
						fmt.Println("tanggal transaksi:", orang[urut].tanggal.tahun,"-",orang[urut].tanggal.bulan,"-",orang[urut].tanggal.hari)
						
					}
					urut++
					
				}
				fmt.Println("===========================================================")
				fmt.Println("ADA DATA SEJUMLAH ",index,"DATA")
				fmt.Println("===========================================================")
				fmt.Print("MASUKAN 0 untuk keluar : ")
				fmt.Scan(&menu)
				fmt.Println("=====================================")
				if menu == 0{}



			}


		} else if login == 2 {
			fmt.Println("...LOADING...")
			time.Sleep(2 * time.Second)
			CallClear()

			fmt.Println("=====================================")
			fmt.Println("------------CREATE AKUN -------------")
			fmt.Println("=====================================")

			fmt.Print("USERNAME: ")
			fmt.Scan(&createuser)
			for createuser == "" {
				fmt.Println("Username tidak boleh kosong, silahkan isi username anda")
				fmt.Print("USERNAME: ")
				fmt.Scan(&createuser)
			}

			fmt.Print("PASSWORD: ")
			fmt.Scan(&createpass)
			for !validPass (createpass) {
				fmt.Println("Password harus terdiri dari angka dan huruf")
				fmt.Print("PASSWORD: ")
				fmt.Scan(&createpass)
			}

			orang[index].createpass = createpass
			orang[index].createuser = createuser

			fmt.Println("=====================================")

			fmt.Println("AKUN BERHASIL DIBUAT")
			time.Sleep(2 * time.Second)
			CallClear()


		} else if login == 3 {
			fmt.Println("...LOADING...")
			time.Sleep(2 * time.Second)
			CallClear()
			fmt.Println("=====================================")
			fmt.Println("-----------SILAHKAN LOGIN------------")
			fmt.Println("=====================================")
			fmt.Print("USERNAME: ")
			fmt.Scan(&loginuser)
			fmt.Print("PASSWORD: ")
			fmt.Scan(&loginpass)
			fmt.Println("=====================================")

			for createuser != loginuser || createpass != loginpass {
				fmt.Println("--INVALID USERNAME OR PASSWORD--")
				time.Sleep(2 * time.Second)
				CallClear()

				fmt.Println("=====================================")
				fmt.Println("-----------SILAHKAN LOGIN------------")
				fmt.Println("=====================================")
				fmt.Print("USERNAME: ")
				fmt.Scan(&loginuser)
				fmt.Print("PASSWORD: ")
				fmt.Scan(&loginpass)
				fmt.Println("=====================================")
			}
			time.Sleep(2 * time.Second)
			CallClear()
			fmt.Println("=====================================")
			fmt.Println("===========LOGIN BERHASIL============")
			fmt.Println("=====================================")
			fmt.Println("1. Tampilkan data barang berdasarkan merk")
			fmt.Println("2. Tampilkan data barang berdasarkan harga jual")
			fmt.Println("3. Transaksi pembelian barang")
			fmt.Println("=====================================")
			fmt.Println("Masukan pilihan menu: ")
			fmt.Scan(&trans)
			fmt.Println("=====================================")

			if trans == 1 {
				fmt.Print("masukan merk yang akan ditampilkan : ")
				fmt.Scan(&merk)
				urut2(&orang, index)
				urut = 1
				for urut <= index {
					if merk == orang[urut].merk {
						fmt.Println("ID: ", orang[urut].id)
						fmt.Println("merk: ", orang[urut].merk)
						fmt.Println("nama:", orang[urut].nama)
						fmt.Println("jenis barang: ", orang[urut].jenis)
						fmt.Println("harga beli: ", orang[urut].hargabeli)
						fmt.Println("harga jual: ", orang[urut].hargajual)
						fmt.Println("stok awal:", orang[urut].stok)
						fmt.Println("tanggal transaksi:", orang[urut].tanggal.tahun,"-",orang[urut].tanggal.bulan,"-",orang[urut].tanggal.hari)
						mrk++
					}
					urut++
					
				}
				if mrk == 0 {
					fmt.Println("BARANG TIDAK DITEMUKAN")
				}
				
				fmt.Println("===========================================================")
				fmt.Print("MASUKAN 0 untuk keluar : ")
				fmt.Scan(&menu)
				fmt.Println("=====================================")
				if menu == 0{}

			} else if trans == 2 {
				urut3(&orang, index)
				urut = 1
				for urut <= index {
					fmt.Println("ID: ", orang[urut].id)
					fmt.Println("merk: ", orang[urut].merk)
					fmt.Println("nama:", orang[urut].nama)
					fmt.Println("jenis barang: ", orang[urut].jenis)
					fmt.Println("harga beli: ", orang[urut].hargabeli)
					fmt.Println("harga jual: ", orang[urut].hargajual)
					fmt.Println("stok awal:", orang[urut].stok)
					fmt.Println("tanggal transaksi:", orang[urut].tanggal.tahun,"-",orang[urut].tanggal.bulan,"-",orang[urut].tanggal.hari)
					urut++
				}
				fmt.Println("===========================================================")
				fmt.Print("MASUKAN 0 untuk keluar : ")
				fmt.Scan(&menu)
				fmt.Println("=====================================")
				if menu == 0{}

			} else if trans == 3 {
				for lagi != 2 {
					mrk = 0
					fmt.Println("=====================================")
					fmt.Println("========TRANSAKSI PEMBELIAN==========")
					fmt.Println("=====================================")
					fmt.Println("Masukan ID barang yang akan di Beli : ")
					fmt.Scan(&id)
					itrans = 1
					for itrans < N-1 && orang[itrans].id != id {
						itrans++	
					}
					if orang[itrans].id == id {
						cekid = itrans
					} else {
						cekid = -1
					}
				
					for cekid == -1 {
						fmt.Println("======BARANG TIDAK DITEMUKAN=====")
						fmt.Println("Masukan ID barang yang akan di Beli: ")
						fmt.Scan(&id)
						
						itrans = 1
						for itrans < N-1 && orang[itrans].id != id {
							itrans++
						}
						if orang[itrans].id == id {
							cekid = itrans
						} else {
							cekid = -1
						}
					}

					fmt.Println("=========BARANG DITEMUKAN========")
					fmt.Println("=================================")
					fmt.Println("ID: ", orang[itrans].id)
					fmt.Println("merk: ", orang[itrans].merk)
					fmt.Println("nama:", orang[itrans].nama)
					fmt.Println("jenis barang: ", orang[itrans].jenis)
					fmt.Println("harga beli: ", orang[itrans].hargabeli)
					fmt.Println("harga jual: ", orang[itrans].hargajual)
					fmt.Println("stok awal:", orang[itrans].stok)
					fmt.Println("tanggal transaksi:", orang[itrans].tanggal.tahun,"-",orang[itrans].tanggal.bulan,"-",orang[itrans].tanggal.hari)

					fmt.Println("=================================")
					fmt.Print("Masukan Jumlah Pembelian Barang:")
					fmt.Scan(&beli)

					for (beli <= 0) || (beli > orang[itrans].stok) {
						fmt.Println("=================================")
						fmt.Print("Masukan ulang Jumlah Pembelian Barang:")
						fmt.Scan(&beli)
					}
					
				
					orang[itrans].stok = orang[itrans].stok - beli
					fmt.Println("=================================")
					fmt.Println("Harga Pembelian: ", beli * orang[itrans].hargajual )
					fmt.Println("=================================")
					fmt.Println("==========UPDATE BARANG==========")
					fmt.Println("=================================")
					fmt.Println("ID: ", orang[itrans].id)
					fmt.Println("merk: ", orang[itrans].merk)
					fmt.Println("nama:", orang[itrans].nama)
					fmt.Println("jenis barang: ", orang[itrans].jenis)
					fmt.Println("harga beli: ", orang[itrans].hargabeli)
					fmt.Println("harga jual: ", orang[itrans].hargajual)
					fmt.Println("stok:", orang[itrans].stok)
					fmt.Println("tanggal transaksi:", orang[itrans].tanggal.tahun,"-",orang[itrans].tanggal.bulan,"-",orang[itrans].tanggal.hari)
					fmt.Println("=================================")
					jum = beli * orang[itrans].hargajual + jum
					fmt.Println("Total Harga Pembelian: ", jum)
					fmt.Println("===========================================================")
					fmt.Print("MASUKAN 2 untuk keluar & 1 untuk beli lagi: ")
					fmt.Scan(&lagi)
					fmt.Println("=====================================")
				}

			
			}

		 
		} else if login != 4 {
			fmt.Println("Masukan 1 , 2 , 3 atau 4")
		}
	}
}