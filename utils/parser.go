package utils

import (
	"database/sql"
	"encoding/xml"
	"nessus-parser/models"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// XmlParse fonksiyonuna bir *sql.DB parametresi eklendi.
func XmlParse(wg *sync.WaitGroup, db *sql.DB, xmlFile *os.File) {
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "ReportHost" {
				var report models.ReportHost
				decoder.DecodeElement(&report, &se)
				wg.Add(1)                    // wg.Add(1) çağrısını DecodeElement'dan önce yapın.
				go DataPrep(wg, db, &report) // db parametresi eklendi.
			}
		}
	}
	wg.Done()
}
