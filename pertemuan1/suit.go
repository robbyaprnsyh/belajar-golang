package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var pemain1, pemain2 string
	var pilihanPemain1, pilihanPemain2 int

	fmt.Println("Pilih mau pilihan apa?")
	fmt.Println("ketik 1 untuk gunting")
	fmt.Println("ketik 2 untuk batu")
	fmt.Println("ketik 3 untuk kertas \n")

	fmt.Scanln(&pilihanPemain1)

	switch pilihanPemain1 {
	case 1:
		pemain1 = "Gunting"
	case 2:
		pemain1 = "Batu"
	case 3:
		pemain1 = "Kertas"
	default:
		fmt.Println("Ada pemain yang memilih di luar pilihan")
	}
	fmt.Println(pemain1)

	pilihanPemain2 = rand.Intn(3)

	switch pilihanPemain2 {
	case 1:
		pemain2 = "Gunting"
	case 2:
		pemain2 = "Batu"
	case 3:
		pemain2 = "Kertas"
	default:
		fmt.Println("Ada pemain yang memilih di luar pilihan")
	}
	fmt.Println(pilihanPemain2)
	fmt.Println(pemain2, "\n")

	fmt.Printf("pilihan pemain satu %s \n", pemain1)
	fmt.Printf("pilihan pemain dua %s \n \n", pemain2)

	if pemain1 == "kertas" && pemain2 == "batu" {
		fmt.Println("Pemain 1 menang")
	} else if pemain1 == "kertas" && pemain2 == "gunting" {
		fmt.Println("pemain 2 menang")
	} else if pemain1 == "gunting" && pemain2 == "kertas" {
		fmt.Println("Pemain 1 menang")
	} else if pemain1 == "batu" && pemain2 == "kertas" {
		fmt.Println("Pemain 2 menang")
	} else if pemain1 == "batu" && pemain2 == "gunting" {
		fmt.Println("Pemain 1 menang")
	} else if pemain1 == "gunting" && pemain2 == "batu" {
		fmt.Println("Pemain 2 menang")
	} else if pemain1 == pemain2 {
		fmt.Println("Hasilnya Draw")
	} else {
		fmt.Println("Ada pemain yang memlihih diluar angka yang disana")
	}
}
