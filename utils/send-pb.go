package utils

import (
	"database/sql"
	"log"
	"nessus-parser/models"
	"sync"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
	pb "github.com/pocketbase/pocketbase"
)

// ... (diğer fonksiyonlar)

func DataPrepForPocketBase(wg *sync.WaitGroup, db *sql.DB, report *models.ReportHost) {
	defer wg.Done()

	pbClient := pb.NewClient("YOUR_POCKETBASE_API_KEY")

	// ... (hashmap oluşturma ve veri hazırlama)

	// HostInfo collection'a veri eklemek
	err := InsertData(pbClient, "HostInfo", hostInfoData)
	if err != nil {
		log.Printf("Error inserting into HostInfo (PocketBase): %v", err)
	}

	// ... (Host ID'yi alma)

	// Vulnerabilities collection'a veri eklemek
	for _, item := range report.ReportItemList {
		for _, cve := range item.CVE {
			// ... (vulnerabilityData hazırlama)
			err := InsertData(pbClient, "Vulnerabilities", vulnerabilityData)
			if err != nil {
				log.Printf("Error inserting into Vulnerabilities (PocketBase): %v", err)
			}
		}
	}
}
