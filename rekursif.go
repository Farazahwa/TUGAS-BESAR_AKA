package main

import "fmt"

// Struktur data untuk Mata Kuliah
type Course struct {
	Name string
	SKS  int
}

// Fungsi rekursif untuk menyusun jadwal KRS
func scheduleKRSRecursive(courseList [10]Course, maxSKS int, currentSKS int, index int, schedule [10]Course, count int) [10]Course {
	if index >= 10 { // Kondisi berhenti jika sudah sampai akhir daftar
		return schedule
	}

	course := courseList[index]

	// Memeriksa apakah mata kuliah bisa dimasukkan
	if currentSKS+course.SKS <= maxSKS {
		// Menambahkan mata kuliah ke jadwal
		schedule[count] = course
		return scheduleKRSRecursive(courseList, maxSKS, currentSKS+course.SKS, index+1, schedule, count+1)
	}

	// Melanjutkan ke mata kuliah berikutnya tanpa menambahkannya
	return scheduleKRSRecursive(courseList, maxSKS, currentSKS, index+1, schedule, count)
}

// Fungsi untuk mencetak jadwal KRS
func printSchedule(schedule [10]Course) {
	fmt.Println("Jadwal KRS:")
	i := 0
	for i < 10 {
		if schedule[i].Name != "" {
			fmt.Printf("Mata Kuliah: %s, SKS: %d\n", schedule[i].Name, schedule[i].SKS)
		}
		i++
	}
}

func main() {
	// Daftar mata kuliah Informatika
	courseList := [10]Course{
		{"Pemrograman Dasar", 3}, {"Struktur Data", 3}, {"Algoritma", 3},
		{"Sistem Operasi", 3}, {"Jaringan Komputer", 3}, {"Pemrograman Web", 3},
		{"Kecerdasan Buatan", 3}, {"Pemrograman Mobile", 3}, {"Kriptografi", 3},
		{"Machine Learning", 3},
	}

	maxSKS := 12
	var schedule [10]Course
	schedule = scheduleKRSRecursive(courseList, maxSKS, 0, 0, schedule, 0)
	printSchedule(schedule)
}