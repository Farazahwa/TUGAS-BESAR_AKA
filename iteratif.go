package main

import "fmt"

// Struktur data untuk Mata Kuliah
type Course struct {
	Name string
	SKS  int
}

// Fungsi iteratif untuk menyusun jadwal KRS
func scheduleKRSIterative(courseList [10]Course, maxSKS int) [10]Course {
	var schedule [10]Course
	currentSKS := 0
	count := 0
	i := 0

	// Iterasi manual tanpa for
	for i < 10 {
		course := courseList[i]
		// Memeriksa apakah mata kuliah bisa dimasukkan
		if currentSKS+course.SKS <= maxSKS {
			// Menambahkan mata kuliah ke jadwal
			schedule[count] = course
			currentSKS += course.SKS
			count++
		}
		i++
	}

	return schedule
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
	schedule := scheduleKRSIterative(courseList, maxSKS)
	printSchedule(schedule)
}