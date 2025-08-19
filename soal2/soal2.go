package soal2

import "fmt"

type LoginRequest struct {
	Email    string
	Password string
}

func simulateLoginUser(user LoginRequest) bool {
	if user.Email == "habibganteng@gmail.com" && user.Password == "Habibganteng123" {
		return true
	}
	return false

	// if user.Username != "" && user.Password == "Habibganteng123" {
	// 	return true
	// }
	// return false
}

func Main() {
	// Ketika login benar
	userLoginSuccess := LoginRequest{
		Email:    "habibganteng@gmail.com",
		Password: "Habibganteng123",
	}

	// Ketika login salah
	userLoginFailed := LoginRequest{
		Email:    "habibjelek@gmail.com",
		Password: "Habibjelek123",
	}

	// Hasil dari Login benar
	if simulateLoginUser(userLoginSuccess) {
		fmt.Println("✅ Login berhasil! Selamat datang, Habibku")
	} else {
		fmt.Println("❌ Login gagal. Periksa kembali username dan password anda ganteng.")
	}

	// Hasil dari login salah
	if simulateLoginUser(userLoginFailed) {
		fmt.Println("✅ Login berhasil! Selamat datang, Habibku")
	} else {
		fmt.Println("❌ Login gagal. Periksa kembali username dan password anda ganteng.")
	}
}
