package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const nUser int = 10

type dataUser struct {
	noData   int
	username string
	password string
}
type dataSementara [nUser]dataUser
type dataSekarang [nUser]dataUser

type email struct {
	pengirim string
	penerima string
	chat     string
	noEmail  int
}
type tabEmail [1000]email

var akun_Dummy dataSementara
var akun_Ori dataSekarang
var data_email tabEmail

var nData int = 0
var nReg int = 0
var nEmail int = 0

func main() {
	akun_Dummy[0].username = "ryan@."
	akun_Dummy[0].password = "ryan"
	akun_Dummy[0].noData = 1

	akun_Dummy[1].username = "agung@."
	akun_Dummy[1].password = "agung"
	akun_Dummy[1].noData = 2

	akun_Dummy[2].username = "bagas@."
	akun_Dummy[2].password = "bagas"
	akun_Dummy[2].noData = 3
	nReg = 3

	data_email[0].pengirim = "agus@."
	data_email[0].penerima = "dani@."
	data_email[0].chat = "halo dani"
	data_email[0].noEmail = 1

	data_email[1].pengirim = "dani@."
	data_email[1].penerima = "dani@."
	data_email[1].chat = "halo dani"
	data_email[1].noEmail = 2

	data_email[2].pengirim = "cantika@."
	data_email[2].penerima = "dani@."
	data_email[2].chat = "halo dani"
	data_email[2].noEmail = 3

	nEmail = 3

	akun_Ori[0].username = "agus@."
	akun_Ori[0].password = "agus"
	akun_Ori[0].noData = 1

	akun_Ori[1].username = "dani@."
	akun_Ori[1].password = "dani"
	akun_Ori[1].noData = 2

	akun_Ori[2].username = "cantika@."
	akun_Ori[2].password = "cantika"
	akun_Ori[2].noData = 3

	nData = 3
	var useroption int
	for {
		Main_menu(&useroption)
		switch useroption {
		case 1:
			UserLogin()
		case 2:
			UserRegister()
		default:
		}
		if useroption == 0 {
			break
		}
	}
}

// Cek Login
func UserLogin() {
	var id, pass, pilih string
	isLogin := false

	for !isLogin {
		Login_menu(&id, &pass)
		idx := cek_login(id)

		if id == "admin" && pass == "admin" {
			Admin_access()
			isLogin = true
		} else if idx < 0 {
			fmt.Println("alamat email tidak ditemukan")
			fmt.Print("Lanjut? (y/n): ")
			fmt.Scan(&pilih)
			if pilih == "n" {
				break
			}
		} else if akun_Ori[idx].password == pass {
			user := akun_Ori[idx].username
			User_access(user)
			isLogin = true
		} else {
			fmt.Println("Username dan password tidak cocok")
			fmt.Println("Lanjut? (y/n): ")
			fmt.Scan(&pilih)
			if pilih == "n" {
				break
			}
		}
	}
}

// Cek Register
func UserRegister() {
	var id, pass, pilih string
	var isValid bool = false

	for !isValid {
		Register_menu(&id, &pass)

		// cek format username email
		userValid := cek_valid_username(id)
		//cek apakah username email sudah ada
		userExist := cek_user_exist(id)

		if !userValid {
			fmt.Println()
			fmt.Println("              format email salah!             ")
			fmt.Println("           contoh: dwibagus@123.com           ")
			fmt.Println()

		}
		if !userExist {
			fmt.Println()
			fmt.Println("          alamt email tidak tersedia!         ")
			fmt.Println("       silahkan masukan alamat email lain     ")
			fmt.Println()
		}

		isValid = userValid && userExist
		fmt.Println("Lanjut? (y/n): ")
		fmt.Scan(&pilih)
	}
	if pilih == "n" {
		fmt.Println("Terimakasih silahkan menunggu pendaftaran anda di setujui")
	} else {
		UserRegister()
	}

	// Register the new user
	akun_Dummy[nReg].noData = nReg + 1
	akun_Dummy[nReg].username = id
	akun_Dummy[nReg].password = pass
	nReg++
}

// sub program dari function login
func Admin_access() {
	var pilih int
	for {
		Admin_menu(&pilih)
		switch pilih {
		case 1:
			cetak_user_request()
		case 2:
			cetak_user_aktif()
		default:
		}
		if pilih == 0 {
			break
		}
	}
}
func User_access(userid string) {
	var pilih int
	for {
		User_menu(&pilih)
		switch pilih {
		case 1:
			// Creat_email()
			buat_email_baru(userid)
		case 2:
			daftar_email_kirim(userid)
		case 3:
			daftar_email_masuk(userid)

		default:
		}
		if pilih == 0 {
			break
		}
	}
}

// ADMIN
func cetak_user_request() {
	var idx int
	var pilih int

	fmt.Println("         DATA REQUEST MASUK          ")
	fmt.Println("-------------------------------------")
	fmt.Printf("%-5s %-20s %-10s\n", "No.", "Alamat email", "Password")
	for i := 0; i < nReg; i++ {
		fmt.Printf("%-5d %-20s %-10s\n", akun_Dummy[i].noData, akun_Dummy[i].username, akun_Dummy[i].password)
		fmt.Println("-------------------------------------")
	}
	fmt.Print("Silahkan pilih no data yang akan anda setujui :")

	fmt.Scan(&pilih)
	if pilih > 0 && pilih <= nReg {
		for i := 0; i < nReg; i++ {
			if akun_Dummy[i].noData == pilih-1 {
				akun_Ori[i].username = akun_Dummy[i].username
				akun_Ori[i].password = akun_Dummy[i].password
				idx = i
				nData++
			}
		}
		for i := idx; i < nReg-1; i++ {
			akun_Dummy[i] = akun_Dummy[i+1]
			akun_Dummy[i].noData--
		}
		nReg = nReg - 1
	} else {
		fmt.Println("data tidak valid")
	}
}
func cetak_user_aktif() {
	urut_data_Ascending()
	for {
		fmt.Println("           DATA USER AKTIF           ")
		fmt.Println("-------------------------------------")
		fmt.Printf("%-5s %-20s %-10s\n", "No.", "Alamat email", "Password")
		fmt.Println("-------------------------------------")
		for i := 0; i <= nData; i++ {
			fmt.Printf("%-5d %-20s %-10s\n", i+1, akun_Ori[i].username, akun_Ori[i].password)
			fmt.Println("-------------------------------------")
		}
		var pilih string
		fmt.Print("kembali? (y/n) : ")
		fmt.Scan(&pilih)
		if pilih == "y" {
			break
		}
	}
}

// USER
func buat_email_baru(userid string) {
	var pilih string
	var username string
	var chat string
	for {
		fmt.Println("           CREATE NEW EMAIL           ")
		fmt.Println("--------------------------------------")
		fmt.Print("Masukan email tujuan :")

		fmt.Scan(&username)
		if cek := cek_valid_email(username); cek {
			fmt.Print("Tulis pesan anda : ")

			fmt.Scanln(&chat)
			reader := bufio.NewReader(os.Stdin)
			chat, _ = reader.ReadString('\n')
			chat = strings.TrimSpace(chat)

			data_email[nEmail].chat = chat
			data_email[nEmail].pengirim = userid
			data_email[nEmail].penerima = username
			data_email[nEmail].noEmail = nData + 1

			fmt.Println()
			fmt.Println("terimakasih email anda sudah terkirim")
			fmt.Println("-------------------------------------")

			nEmail += 1
		} else {
			fmt.Println("email yang anda tuju tidak tersedia")
		}

		fmt.Println("buat lagi ? (y/n) : ")
		fmt.Scan(&pilih)
		if pilih == "n" {
			break
		}
	}
}

func daftar_email_kirim(userid string) {
	var pilih string
	for {
		fmt.Println("       EMAIL TERKIRIM       ")
		fmt.Println("----------------------------")
		for i := 0; i < nEmail; i++ {
			if data_email[i].pengirim == userid {
				fmt.Printf("%-6s %-20s\n", "kepada    :", data_email[i].penerima)
				fmt.Printf("%-6s %-20s\n", "isi pesan :", data_email[i].chat)
				fmt.Println("----------------------------")

			}
		}
		fmt.Print("kembali ? (y/n): ")
		fmt.Scan(&pilih)
		if pilih == "y" {
			break
		}
	}
}
func daftar_email_masuk(userid string) {
	var pilih int
	var no int = 1

	//mengurutkan data email user secara Descending
	urut_data_Descending()

	fmt.Println("     DAFTAR EMAIL MASUK     ")
	fmt.Println("----------------------------")
	for i := 0; i <= nEmail; i++ {
		if data_email[i].penerima == userid {
			fmt.Printf("%-5d %-20s\n", no, data_email[i].pengirim)
			fmt.Println("----------------------------")
			no++
		}

	}
	fmt.Println("silahkan pilih no email yang akan anda lihat ")
	fmt.Println("ketik angka 0 untuk kembali : ")
	fmt.Scan(&pilih)
	if pilih == 0 {
		fmt.Print()
	} else {
		cetak_email_masuk(pilih-1, userid)
	}
}
func cetak_email_masuk(pilih int, userid string) {
	var idx_pesan int
	var balas, hapus string

	for i := 0; i <= nEmail; i++ {
		if data_email[i].penerima == userid {
			if pilih == 0 {
				fmt.Println("          PESAN MASUK         ")
				fmt.Println("----------------------------")
				fmt.Printf("%-6s %-20s\n", "dari      :", data_email[i].pengirim)
				fmt.Printf("%-6s %-20s\n", "isi pesan :", data_email[i].chat)
				fmt.Println("----------------------------")

				idx_pesan = i
				fmt.Print("balas pesan ? (y/n): ")
				fmt.Scan(&balas)
				if balas == "y" {
					balas_email(idx_pesan)
					pilih = 0
				}
				if balas == "n" {
					fmt.Print("hapus pesan ? (y/n): ")
					fmt.Scan(&hapus)
					if hapus == "y" {
						hapus_email(idx_pesan)
					} else {
						break
					}
				}
			}
			pilih--
		}
	}
}

func balas_email(n int) {
	var pesan string

	fmt.Print("Tulis balasan anda :")
	fmt.Scanln(&pesan)
	reader := bufio.NewReader(os.Stdin)
	pesan, _ = reader.ReadString('\n')
	pesan = strings.TrimSpace(pesan)

	data_email[nEmail].chat = pesan
	data_email[nEmail].pengirim = data_email[n].penerima
	data_email[nEmail].penerima = data_email[n].pengirim
	data_email[nEmail].noEmail = nData + 1

	fmt.Println("terimakasih email anda sudah terkirim")
	fmt.Println("-------------------------------------")
	nEmail++
}
func hapus_email(n int) {
	for i := n; i < nEmail; i++ {
		data_email[i] = data_email[i+1]
	}
	nEmail--
}

func cek_valid_email(id string) bool {
	for i := 0; i < nData; i++ {
		if id == akun_Ori[i].username {
			return true
		}
	}
	return false
}

func cek_valid_username(id string) bool {
	ada1, ada2 := false, false
	for i := 0; i < len(id); i++ {
		if id[i] == '@' {
			ada1 = true
		}
		if ada1 && id[i] == '.' {
			ada2 = true
		}
	}
	return ada1 && ada2
}
func cek_user_exist(id string) bool {
	for i := 0; i < nUser; i++ {
		if id == akun_Ori[i].username || akun_Ori[i].username == "admin" {
			return false
		}
	}
	return true
}
func cek_login(id string) int {
	//squential search
	idx := -1
	for i := 0; i < nData; i++ {
		if akun_Ori[i].username == id {
			idx = i
		}
	}
	return idx
}

func urut_data_Ascending() {
	//selection sort Ascending
	for a := 0; a < nData; a++ {
		tempA := akun_Ori[a]
		tempB := akun_Ori[a].username
		pos := a

		for b := a + 1; b < nData; b++ {
			if akun_Ori[b].username < tempB {
				tempB = akun_Ori[b].username
				tempA = akun_Ori[b]
				pos = b
			}
		}
		akun_Ori[pos] = akun_Ori[a]
		akun_Ori[a] = tempA

	}
}

func urut_data_Descending() {
	var i, pass int
	var temp email
	pass = 1
	for pass <= nEmail-1 {
		i = pass
		temp = data_email[pass]
		for i > 0 && temp.noEmail > data_email[i-1].noEmail {
			data_email[i] = data_email[i-1]
			i = i - 1
		}
		data_email[i] = temp
		pass = pass + 1
	}
}

func Main_menu(useroption *int) {
	fmt.Println("--------------------------")
	fmt.Println("          M E N U         ")
	fmt.Println("--------------------------")
	fmt.Println("1. Login        		   ")
	fmt.Println("2. Register               ")
	fmt.Println("0. Exit                   ")
	fmt.Println("--------------------------")
	fmt.Print("Pilih (1/2/0): ")
	fmt.Scan(useroption)
}
func Admin_menu(useroption *int) {
	fmt.Println("--------------------------")
	fmt.Println("         A D M I N        ")
	fmt.Println("--------------------------")
	fmt.Println("1. Daftar request user    ")
	fmt.Println("2. Cek user aktif         ")
	fmt.Println("0. Logout                 ")
	fmt.Println("--------------------------")
	fmt.Print("Pilih (1/2/0): ")
	fmt.Scan(useroption)
}
func User_menu(useroption *int) {
	fmt.Println("--------------------------")
	fmt.Println("          U S E R         ")
	fmt.Println("--------------------------")
	fmt.Println("1. Create new email	   ")
	fmt.Println("2. Daftar email terkirim  ")
	fmt.Println("3. Daftar email masuk	   ")
	fmt.Println("0. Logout                 ")
	fmt.Println("--------------------------")
	fmt.Println("Pilih (1/2/0): ")
	fmt.Scan(useroption)
}
func Login_menu(id, pass *string) {
	fmt.Println()
	fmt.Println("Silahkan masukkan alamat email dan password anda")
	fmt.Println()
	fmt.Print("Username: ")
	fmt.Scan(id)
	fmt.Print("Password: ")
	fmt.Scan(pass)
}
func Register_menu(id, pass *string) {
	fmt.Println("           Registrasi Akun Baru              ")
	fmt.Println("---------------------------------------------")
	fmt.Print("Alamat email : ")
	fmt.Scan(id)
	fmt.Print("Password     : ")
	fmt.Scan(pass)
	fmt.Println("---------------------------------------------")

}
