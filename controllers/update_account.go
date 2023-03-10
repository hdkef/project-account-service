package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"project/helper"
	"project/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func updatePassword(ID int, db *sql.DB, u *models.User) int {
	//input password
	var data string
	fmt.Print("\n")
	fmt.Print("Input password baru : ")
	_, err := fmt.Scanln(&data)
	if err != nil {
		fmt.Println("Update password gagal")
		return -1
	}

	//validasi password
	valid, msg := helper.ValidasiPassword(data)
	if !valid {
		fmt.Println(msg)
		return -1
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err != nil {
		fmt.Println("Update password gagal")
		return -1
	}
	_, err = db.Exec("UPDATE users SET password = ?, updated_at = now() WHERE id = ?", string(hashedPassword), ID)
	if err != nil {
		fmt.Println("Update password gagal")
		return -1
	}
	fmt.Println("Update password berhasil, menjadi ", data)
	return -1
}

func updateTanggalLahir(ID int, db *sql.DB, u *models.User) int {
	fmt.Print("\n")
	fmt.Print("Input tanggal lahir baru (tahun-bulan-tanggal) : ")
	in := bufio.NewReader(os.Stdin)
	data, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Update nama gagal")
		return -1
	}

	//parsing dob
	dob, err := time.Parse("2006-1-2\n", data)
	if err != nil {
		fmt.Println("Harap mengisi tanggal lahir dengan benar")
		return -1
	}

	//validasi tanggal lahir
	valid, msg := helper.ValidasiTanggalLahir(dob)
	if !valid {
		fmt.Println(msg)
		return -1
	}

	//update tanggal lahir
	_, err = db.Exec("UPDATE users SET date_of_birth = ?, updated_at = now() WHERE ID = ?", dob, ID)
	if err != nil {
		fmt.Println("Update tanggal lahir gagal")
		return -1
	}

	fmt.Println("Update tanggal lahir berhasil, menjadi ", dob.Format("January 2, 2006"))
	u.DateOfBirth = dob
	return -1
}

func updateTelepon(ID int, db *sql.DB, u *models.User) int {
	//input telepon
	var data string
	fmt.Print("\n")
	fmt.Print("Input telepon baru : ")
	_, err := fmt.Scanln(&data)
	if err != nil {
		fmt.Println("Update telepon gagal")
		return -1
	}

	//validasi telepon
	valid, msg := helper.ValidasiTelepon(data, db)
	if !valid {
		fmt.Println(msg)
		return -1
	}

	//update telepon
	_, err = db.Exec("UPDATE users SET phone = ?, updated_at = now() WHERE ID = ?", data, ID)
	if err != nil {
		fmt.Println("Update telepon gagal")
		return -1
	}

	fmt.Println("Update telepon berhasil, menjadi ", data)
	u.Phone = data
	return -1
}

func updateNama(ID int, db *sql.DB, u *models.User) int {
	//input nama
	fmt.Print("\n")
	fmt.Print("Input nama baru : ")
	in := bufio.NewScanner(os.Stdin)
	valid := in.Scan()
	if !valid {
		fmt.Println("Nama tidak valid")
		return -1
	}
	data := in.Text()

	//validasi nama
	valid, msg := helper.ValidasiNama(data)
	if !valid {
		fmt.Println(msg)
		return -1
	}

	//update nama
	_, err := db.Exec("UPDATE users SET name = ?, updated_at = now() WHERE ID = ?", data, ID)
	if err != nil {
		fmt.Println("Update nama gagal")
		return -1
	}

	fmt.Println("Update nama berhasil, menjadi ", data)
	u.Name = data
	return -1
}

func UpdateAccount(db *sql.DB, ID int, u *models.User) int {
	//input pilih data yang akan diupdate

	var opsi int = -1
	//pilih sesuai menu
out:
	for opsi != 5 {
		fmt.Print("\n")
		fmt.Print("Pilih data yang akan diupdate :\n\n1.Telepon\n2.Nama\n3.Tanggal Lahir\n4.Password\n5.Menu Utama\n6.Exit\n\nPilih Menu : ")
		fmt.Scanln(&opsi)
		switch opsi {
		case 1:
			opsi = updateTelepon(ID, db, u)
		case 2:
			opsi = updateNama(ID, db, u)
		case 3:
			opsi = updateTanggalLahir(ID, db, u)
		case 4:
			opsi = updatePassword(ID, db, u)
		case 5:
			break out
		case 6:
			return 9
		}
	}
	return -1
}
