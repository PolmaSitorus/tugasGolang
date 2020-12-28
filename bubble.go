/*	mulai belajar go language
	======================

	Bubble Sort:

	Bubble Sort adalah metode pengurutan algoritma dengan cara melakukan penukaran data
	secara terus menerus sampai bisa dipastikan dalam suatu iterasi tertentu tidak ada lagi perubahan/penukaran.
	Algoritma ini menggunakan perbandingan dalam operasi antar elemennya.

	Berikut ini adalah gambaran dari algoritma bubble sort:

	1. Bandingkan nilai data ke-1 dan data ke-2
	2. Jika data ke-1 lebih besar dari data ke-2 maka tukar posisinya
	3. Kemudian data yg lebih besar tadi dibandingkan dengan data ke-3
	4. Lakukan langkah nomer 2 hingga selesai.


*/

package main

import (
	"fmt"
)

func main() {
	var data = [8]int{25, 57, 48, 37, 12, 92, 80, 33}

	fmt.Println("data sebelum di urutkan ", data)
	fmt.Println("---")
	urutkanData(data)

}

func urutkanData(dat [8]int) {
	var i, j, tmp, pjgdata int
	pjgdata = len(dat)

	for i = 1; i < pjgdata; i++ {
		for j = 0; j < pjgdata-i; j++ {
			if dat[j] > dat[j+1] {
				//tukar posisi buble
				tmp = dat[j]
				dat[j] = dat[j+1]
				dat[j+1] = tmp
			}
		}
	}
	fmt.Println("data setelah di urutkan ", dat)
}
