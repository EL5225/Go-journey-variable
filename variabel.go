package main //Inisiasi untuk tempat execute program. Jadi menjalankan program pasti di package main

import "fmt" //package umum untuk I/O

func main() {
	var i int //definisi variabel metode biasa
	i = 13
	var j int = 4 // definisi varibel langsung di assignment
	k := 12       // definisi variabel menggunakan ':=' tanpa tipe datanya. Tipe data akan langsung terbaca secara otomatis sesuai value

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
