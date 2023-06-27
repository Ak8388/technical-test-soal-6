package main

import (
	"fmt"
	"strings"
)

func main() {
	//Soal 6 Bag 1
	for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}

	//Soal 6 Bag 2
	Palindrom("kasur ini Rusak")

	//Soal 6 bag 3
	Fakotrial(5)

	//soal 6 bag 4
	nilai := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 90}
	MinMax(nilai...)

	//Soal 6 bag 5
	Count("AZaz")
}

func FizzBuzz(nilai int) {
	if nilai%3 == 0 && nilai%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if nilai%3 == 0 {
		fmt.Println("Fizz")
	} else if nilai%5 == 0 {
		fmt.Println("Buzz")
	}
}

func Palindrom(kalimat string) {
	kalimat = strings.ToLower(kalimat)
	kalimat = strings.ReplaceAll(kalimat, " ", "")
	for i := 0; i < len(kalimat); i++ {
		if kalimat[i] != kalimat[len(kalimat)-(i+1)] {
			fmt.Println("Kalimat/Kata Bukan Merupakan Polindrom")
			return
		}
	}

	fmt.Println("Kalimat/Kata Merupakan Polindrom")
}

func Fakotrial(nilai int) {
	jumlah := 1
	for i := nilai; i >= 1; i-- {
		jumlah *= i
	}

	fmt.Println(jumlah)
}

func MinMax(nilai ...int) {
	max := 0

	for _, angka := range nilai {
		if angka > max {
			max = angka
		}
	}
	min := max
	for _, angka := range nilai {
		if angka < min {
			min = angka
		}
	}
	fmt.Printf("Nilai Min : %d\nNilai Max : %d\n", min, max)
}

func Count(kalimat string) {
	var angka, huruf int
	for i := 0; i < len(kalimat); i++ {
		if kalimat[i] >= 48 && kalimat[i] <= 57 {
			angka++
		} else if kalimat[i] >= 65 && kalimat[i] <= 90 || kalimat[i] >= 97 && kalimat[i] <= 122 {
			huruf++
		}
	}

	fmt.Printf("Jumlah Angka :%d\nJumlah Huruf :%d", angka, huruf)
}
