package main

import "fmt"

func BELAJARTIPEDATA() {
	//tipe data integer
	// integer positif
	var positiveNumber uint8 = 89
	var negativeNumber = -1233434343434 // otomatis diinference ke int
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	fmt.Printf("var negativeNumber dengan value %d bertipe data %T\n", negativeNumber, negativeNumber)

	//TTPE DATA FLOAT
	var decimalNumber = 2.35
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	// pembulatan 3 angka dibelakang koma
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	fmt.Printf("var decimalNumber dengan value %f bertipe data %T\n", decimalNumber, decimalNumber)

	//TIPE DATA BOOL
	var exist bool = true
	fmt.Printf("exist? var memiliki nilai boolean %t \n", exist)

	//TIPE DATA STRING
	// satu baris
	var message = "hello"
	fmt.Printf("var message memiliki nilai string: %s\n", message)

	// multi baris dg backtick ``
	var messagePanjang = `halo juga
	nama saya "ngga penting", yg pe
	nting tujuannya`
	fmt.Printf(messagePanjang)

	// NILAI NIL & ZERO VALUE
}

func BELAJARVARIABLE() {
	// deklarasi variable dg tipe datanya dan value
	// var firstName string = "john"

	// deklarasi variable tanpa value
	// var lastName string
	// assignment value kemudian
	// lastName = "wick"

	//------------------------------------------------
	// deklarasi variable tanpa tipe data
	// manifest typing
	var firstName string = "iko"

	// inference.. tanpa var.. dan type data dihandle go
	lastName := "uwais"

	// print format sehingga value variable di masukan ke string langsung
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	//---------------------------------------------
	// deklarasi multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println(first, second, third)

	// deklarasi multi variable 1 baris
	var four, five, six string = "empat", "lima", "enam"
	// tanpa var dan type data
	seven, eight, nine := "tujuh", "delapan", "sembilan"
	fmt.Println(four, five, six)
	fmt.Println(seven, eight, nine)

	// deklarasi multi variable dg type berbeda dg teknik inference
	one, isTwo, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one, isTwo, twoPointTwo, say)

	// deklarasi variable dg underscore.. utk var yg tidak akan dipakai
	// name, _ := "john", "wick"

	//------------------------------------------------------
	// deklarasi var dg new keyword
	// membuat variable pointer dg default value tergantung dg type datanya
	alamat := new(int)
	// shg ketika diprint akan menampilkan memory addresnya
	fmt.Println(alamat)
	// untuk memprint valuenya harus di dereference dahulu
	fmt.Printf("value dari alamat adalah %v", *alamat)
}

func bikinhelloworld() {
	fmt.Println("hello world")
	fmt.Println("Hello", "world!", "how", "are", "you")

}
