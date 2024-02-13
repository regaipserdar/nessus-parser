package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"nessus-parser/utils"
	"os"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite driver import
)

var (
	verbose int
	dbPath  string
)

func main() {
	startTime := time.Now() // Programın başlama zamanını kaydet

	fileOpt := flag.String("file", "", "file to parse into db") // Varsayılan değeri boş string olarak güncelledim.
	coreOpt := flag.Int("core", 1, "Number of Cores to use")
	verboseOpt := flag.Int("verbose", 0, "Verbose level 0,1,2")
	flag.StringVar(&dbPath, "db", "default.db", "Path to the SQLite database file")
	flag.Parse()

	utils.CoreCheck(coreOpt, verboseOpt)
	verbose = *verboseOpt

	// Veritabanı bağlantısını aç
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Error opening database:", err)
		return
	}
	defer db.Close()

	// Tabloları oluştur
	err = utils.CreateTables(db)
	if err != nil {
		log.Fatal("Error creating tables:", err)
		return
	}

	file := *fileOpt
	var wg sync.WaitGroup
	if file != "" { // "xmlFile" kontrolünü dosya adı boş olup olmadığına göre güncelledim.
		xmlFile, err := os.Open(file)
		if err != nil {
			log.Fatal("Error opening XML file:", err)
			return
		}
		defer xmlFile.Close()

		wg.Add(1)
		// utils.XmlParse fonksiyonuna veritabanı bağlantısını da geçir.
		go utils.XmlParse(&wg, db, xmlFile) // db parametresi eklendi.
	}
	wg.Wait()

	endTime := time.Now()                                // Programın bitiş zamanını kaydet
	elapsed := endTime.Sub(startTime).Round(time.Second) // Geçen süreyi hesapla

	fmt.Printf("Program started at: %s\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Program ended at: %s\n", endTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
