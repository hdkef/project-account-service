package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func updatePassword(ID int, db *sql.DB) int {
	//input password
	var data string
	fmt.Println("input password baru :")
	_, err := fmt.Scanln(&data)
	if err != nil {
		fmt.Println("update password gagal")
		return -1
	}

	//validasi password
	//minimal 8 karakter
	if len(data) < 8 {
		fmt.Println("password minimal 5 huruf")
		return -1
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err != nil {
		fmt.Println("update password gagal")
		return -1
	}
	_, err = db.Exec("UPDATE users SET password = ?, updated_at = now() WHERE id = ?", string(hashedPassword), ID)
	if err != nil {
		fmt.Println("update password gagal")
		return -1
	}
	fmt.Println("update password berhasil, menjadi ", data)
	return -1
}

func updateTanggalLahir(ID int, db *sql.DB) int {
	fmt.Println("input tanggal (tahun-bulan-tanggal) :")
	in := bufio.NewReader(os.Stdin)
	data, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("update nama gagal")
		return -1
	}

	//parsing dob
	dob, err := time.Parse("2006-1-2\n", data)
	if err != nil {
		fmt.Println("harap mengisi tanggal lahir dengan benar")
		return -1
	}

	//validasi tanggal lahir
	if time.Since(dob).Hours()/24/365 < 17 {
		fmt.Println("minimal usia 17 tahun")
		return -1
	}

	//update tanggal lahir
	_, err = db.Exec("UPDATE users SET date_of_birth = ?, updated_at = now() WHERE ID = ?", dob, ID)
	if err != nil {
		fmt.Println("update tanggal lahir gagal")
		return -1
	}

	fmt.Println("update tanggal lahir berhasil, menjadi ", dob.String())
	return -1
}

func updateTelepon(ID int, db *sql.DB) int {
	//input telepon
	var data string
	fmt.Println("input telepon :")
	_, err := fmt.Scanln(&data)
	if err != nil {
		fmt.Println("update telepon gagal")
		return -1
	}

	//validasi telepon
	//minimal 10 karakter
	if len(data) < 10 || len(data) > 12 {
		fmt.Println("telepon minimal 10 karakter maksimal 12")
		return -1
	}

	//hanya angka
	if !regexp.MustCompile(`^[0-9]*$`).MatchString(data) {
		fmt.Println("telepon hanya terdiri dari angka")
		return -1
	}

	//update telepon
	_, err = db.Exec("UPDATE users SET phone = ?, updated_at = now() WHERE ID = ?", data, ID)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			fmt.Println("nomor telepon telah dipakai")
			return -1
		}
		fmt.Println("update telepon gagal")
		return -1
	}

	fmt.Println("update telepon berhasil, menjadi ", data)
	return -1
}

func updateNama(ID int, db *sql.DB) int {
	//input nama
	fmt.Println("input nama :")
	in := bufio.NewScanner(os.Stdin)
	valid := in.Scan()
	if !valid {
		fmt.Println("nama tidak valid")
		return -1
	}
	data := in.Text()

	//validasi nama
	//maksimal 50 karakter
	if len(data) > 50 {
		fmt.Println("nama maksimal 50 karakter")
		return -1
	}
	//hanya huruf dan spasi
	if !regexp.MustCompile(`^[a-zA-Z ]*$`).MatchString(data) {
		fmt.Println("nama hanya boleh alfabet atau spasi")
		return -1
	}

	//update nama
	_, err := db.Exec("UPDATE users SET name = ?, updated_at = now() WHERE ID = ?", data, ID)
	if err != nil {
		fmt.Println("update nama gagal")
		return -1
	}

	fmt.Println("update nama berhasil, menjadi ", data)
	return -1
}

func UpdateAccount(db *sql.DB, ID int) int {
	//input pilih data yang akan diupdate

	var opsi int = -1
	//pilih sesuai menu
out:
	for opsi != 5 {
		fmt.Println("pilih data yang akan diupdate\n1.Telepon\n2.Nama\n3.Tanggal Lahir\n4.Password\n5.Menu Utama\n6.Exit")
		fmt.Scanln(&opsi)
		switch opsi {
		case 1:
			opsi = updateTelepon(ID, db)
		case 2:
			opsi = updateNama(ID, db)
		case 3:
			opsi = updateTanggalLahir(ID, db)
		case 4:
			opsi = updatePassword(ID, db)
		case 5:
			break out
		case 6:
			return 9
		}
	}
	return -1
}
