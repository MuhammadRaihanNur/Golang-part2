package main

import (
	"fmt"
	"time"
)

// Worker Pool function
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d memulai tugas %d\n", id, job)
		time.Sleep(time.Second) // Simulasi tugas yang membutuhkan waktu
		fmt.Printf("Worker %d selesai mengerjakan tugas %d\n", id, job)
		results <- job * 2 // Hasil dari tugas (misalnya, perkalian angka)
	}
}

func main() {
	// Membuat 2 channel untuk jobs dan results
	jobs := make(chan int, 100)    // Buffer untuk tugas
	results := make(chan int, 100) // Buffer untuk hasil

	// Membuat 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Mengirim 5 pekerjaan ke channel jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Menutup channel jobs setelah mengirim semua tugas

	// Mengambil dan mencetak hasil dari channel results
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("Hasil dari tugas: %d\n", result)
	}
}
