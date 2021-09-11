package main

import (
	"fmt"
	"strconv"
)

var jumlah, pilihan int
var next, ID, judul, rating, rate string
var id int = 0

var games = []map[string]string{}

func inputData() {

	id++

	fmt.Println("===========================================")
	fmt.Println("               INPUT DATA                  ")
	fmt.Println("===========================================")

	fmt.Print("Jumlah data yang akan diinput : ")
	fmt.Scan(&jumlah)
	for i := 0; i < jumlah; i++ {

		fmt.Print("Data Ke ", i+1, "\n")

		fmt.Print("Input Game: ")
		fmt.Scan(&judul)

		fmt.Print("Input rating : ")
		fmt.Scan(&rating)

		var ID = strconv.Itoa(id)
		var rating, err = strconv.ParseFloat(rating, 64)
		if err == nil {
			if rating > 4 {
				rate = "Good"
			} else if rating >= 2 && rating <= 4 {
				rate = "Average"
			} else if rating < 2 {
				rate = "Poor"
			}

			var rateStr = strconv.FormatFloat(rating, 'f', 2, 64)

			var beforeGames []map[string]string = games

			var afterGames = []map[string]string{
				{"ID": ID, "Judul Game": judul, "Rating Game": rate, "Rate": rateStr},
			}

			games = append(beforeGames, afterGames...)
			fmt.Println("Berhasil melakukan input data : ", games)
			fmt.Println(" ")

		}
	}

	fmt.Printf("Tekan y[enter] untuk kembali ke menu \n")
	fmt.Printf("Tekan 0[enter] untuk keluar dari program : ")
}

func deleteData() {
	var id int

	fmt.Println("=================================")
	fmt.Println("           HAPUS DATA            ")
	fmt.Println("=================================")

	if len(games) == 0 {
		fmt.Println("TIDAK ADA DATA")
	} else {
		fmt.Print("Masukan id yang akan dihapus : ")
		fmt.Scan(&id)

		var before = games[:id-1]
		var after = games[id:]

		games = append(before, after...)
		fmt.Println("Data berhasil dihapus", games)

	}
	fmt.Println("")
	fmt.Printf("Tekan y[enter] untuk kembali ke menu \n")
	fmt.Printf("Tekan 0[enter] untuk keluar dari program : ")
}

func viewData() {
	fmt.Println("========================")
	fmt.Println("Judul Game | Rating Game")
	fmt.Println("========================")
	for _, game := range games {
		fmt.Println("  ", game["Judul Game"], "        ", game["Rating Game"])
	}

	fmt.Println("Jumlah data yang tersimpan : ", len(games), "data")

	fmt.Printf("Tekan y[enter] untuk kembali ke menu \n")
	fmt.Printf("Tekan 0[enter] untuk keluar dari program : ")
}

func searchData() {
	search := 0
	fmt.Print("Masukan judul yang akan dicari : ")
	fmt.Scan(&judul)

	for _, game := range games {
		if game["Judul Game"] == judul {
			search++
			fmt.Println(game["Judul Game"], game["Rating Game"])
		}
	}

	if search == 0 {
		fmt.Println("Data Tidak Ditemukan")
	}

	fmt.Printf("Tekan y[enter] untuk kembali ke menu \n")
	fmt.Printf("Tekan 0[enter] untuk keluar dari program : ")
}

func viewTop() {
	var top int = 0
	var i float64 = 5.0

	for i >= 0 {
		for _, game := range games {
			var rateFloat, _ = strconv.ParseFloat(game["Rate"], 64)
			if i == rateFloat {
				fmt.Println("Judul : ", game["Judul Game"], "Rating : ", game["Rating Game"])
				top++
			}

			if top == 3 {
				break
			}
		}
		if top == 3 {
			break
		}
	}

	fmt.Printf("Tekan y[enter] untuk kembali ke menu \n")
	fmt.Printf("Tekan 0[enter] untuk keluar dari program : ")
}

func viewAll() {
	var rating4 int = 0

	fmt.Println("=====================================")
	fmt.Println("       Rating Games > 4.0            ")
	fmt.Println("=====================================")

	for _, game := range games {
		var rateFloat, _ = strconv.ParseFloat(game["Rate"], 64)
		if rateFloat > 4.0 {
			rating4++
			fmt.Println("Judul Game : ", game["Judul Game"], "Rate : ", game["Rate"], "Rating : ", game["Rating Game"])
		}

	}

	if rating4 == 0 {
		fmt.Println("Data tidak ditemukan")
	}

	fmt.Printf("Tekan y[enter] untuk kembali ke menu \n")
	fmt.Printf("Tekan 0[enter] untuk keluar dari program : ")

}

func menu() {

	fmt.Println("========================================================")
	fmt.Println("                      Daftar Menu                       ")
	fmt.Println("========================================================")
	fmt.Println("1. Input data game")
	fmt.Println("2. Menghapus data menurut id")
	fmt.Println("3. Menampilkan Keseluruhan Data")
	fmt.Println("4. Mencari data berdasarkan menurut judul game")
	fmt.Println("5. Menampilkan top 3 game favorit ")
	fmt.Println("6. Menampilkan seluruh data dengan rating diatas 4.0")
	fmt.Println("0. Keluar dari menu")
	fmt.Println("=======================================================")
	fmt.Print("Masukan pilihan menu [1-6] : ")
	fmt.Scan(&pilihan)

	if pilihan == 0 {
		fmt.Println("==================      END     ===================")

	} else if pilihan == 1 {
		inputData()

	} else if pilihan == 2 {
		deleteData()
	} else if pilihan == 3 {
		viewData()
	} else if pilihan == 4 {
		searchData()
	} else if pilihan == 5 {
		viewTop()
	} else if pilihan == 6 {
		viewAll()
	}

	if pilihan != 0 {
		fmt.Scan(&next)
		if next == "y" {
			main()
		} else {
			fmt.Println("=========== END ===========")
		}
	}

}

func main() {
	menu()
}
