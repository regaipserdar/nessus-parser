package utils

import (
	"database/sql"
	"fmt"
	"log"
	"nessus-parser/models"
	"sync"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// DataPrep fonksiyonuna bir *sql.DB parametresi eklendi.
func DataPrep(wg *sync.WaitGroup, db *sql.DB, report *models.ReportHost) {
	defer wg.Done() // Bu satırı fonksiyonun başına taşıdım.

	hashmap := make(map[string]string)
	for _, host := range report.HostProperties.Info {
		hashmap[host.Key] = host.Value
	}

	// HostProperties'dan HostInfo'ya veri hazırlama
	hostInfoData := map[string]interface{}{
		"IP":             hashmap["host-ip"],
		"OS":             hashmap["operating-system"],
		"MAC":            hashmap["mac-address"],
		"FQDN":           hashmap["host-fqdn"],
		"NetbiosName":    hashmap["netbios-name"],
		"HostProperties": fmt.Sprintf("%v", hashmap), // JSON olarak saklamayı düşünebilirsiniz
	}

	// HostInfo tablosuna veri ekleme
	err := InsertData(db, "HostInfo", hostInfoData)
	if err != nil {
		log.Printf("Error inserting into HostInfo: %v", err)
		return
	}

	// Host ID'yi al (Örnek: en son eklenen ID'yi kullanarak)
	var hostID int
	err = db.QueryRow("SELECT last_insert_rowid()").Scan(&hostID)
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return
	}

	// ReportItem'lardan Vulnerabilities'ye veri hazırlama ve ekleme
	for _, item := range report.ReportItemList {
		for _, cve := range item.CVE {
			vulnerabilityData := map[string]interface{}{
				"HostID":                     hostID,
				"CVE":                        cve,
				"CvssBaseScore":              item.CvssBaseScore,
				"Description":                item.Desciption,
				"ExploitAvailable":           item.ExploitAvailable,
				"ExploitFrameworkCanvas":     item.ExploitFrameworkCanvas,
				"ExploitFrameworkCore":       item.ExploitFrameworkCore,
				"ExploitFrameworkMetasploit": item.ExploitFrameworkMetasploit,
				"ExploitEase":                item.ExploitEase,
				"MetasploitName":             item.MetasploitName,
				"PluginName":                 item.PluginName,
				"RiskFactor":                 item.RiskFactor,
				"Solution":                   item.Solution,
				"Synopsis":                   item.Synopsis,
				"PluginOutput":               item.PluginOutput,
				"SeeAlso":                    item.SeeAlso,
				"ExploitedMalware":           item.ExploitedMalware,
				"Severity":                   item.Severity,
				"PluginID":                   item.PluginID,
			}

			err := InsertData(db, "Vulnerabilities", vulnerabilityData)
			if err != nil {
				log.Printf("Error inserting into Vulnerabilities: %v", err)
				// Hata durumunda döngüden çıkabilir veya loglama yapabilirsiniz.
			}
		}
	}
}
