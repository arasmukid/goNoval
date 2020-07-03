package main

import "fmt"

func main() {
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
