package main
import "fmt"
func main(){
	var hari, penumpang string
	var normal, keluaran int
	fmt.Scan(&hari, &penumpang, &normal)
	keluaran = hargaTiket(hari, penumpang, normal)
	fmt.Print(keluaran)
}
func hargaTiket(hari, penumpang string, normal int)int{
	var hasil float64
	if hari == "selasa" || hari == "rabu" || hari == "kamis"{
		if penumpang == "pelajar"{
			hasil = float64(normal)*0.9
		}else{
			hasil = float64(normal)*0.95
		}
	}else{
		hasil = float64(normal)
	}
	return int(hasil)
}