package main

import "fmt"

func main() {
	var gaji_pokok, potongan, bonus_lembur, jam_lembur, gaji_bersih int

	fmt.Scan(&gaji_pokok, &jam_lembur)
	bonus_lembur = 45000 * jam_lembur
	potongan = (2*gaji_pokok)/100 + (35*gaji_pokok)/1000
	gaji_bersih = gaji_pokok + bonus_lembur - potongan

	fmt.Println(gaji_bersih)
}
