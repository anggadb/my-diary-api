package main

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Dokumen .env tidak ditemukan")
	}
}
func main() {}
