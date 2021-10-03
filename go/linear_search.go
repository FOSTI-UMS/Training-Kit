package main


// funcgsi mencari angka di slice integer retuen index jika ditemukan jika tidak ada retuen -1
func linearSearch(data []int, target int) int {
	for index, v := range data {
		if v == target {
			return index
		}
	}
	return -1
}