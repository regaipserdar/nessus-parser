package utils

import (
	"database/sql"
	"log"
)

// CreateTables veritabanında gerekli tabloları oluşturur.
func CreateTables(db *sql.DB) error {
	// HostInfo tablosunu oluşturma
	createHostInfoTableSQL := `
	CREATE TABLE IF NOT EXISTS HostInfo (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		IP TEXT,
		OS TEXT,
		MAC TEXT,
		FQDN TEXT,
		NetbiosName TEXT,
		HostProperties TEXT
	);`

	// Vulnerabilities tablosunu oluşturma
	createVulnerabilityTableSQL := `
CREATE TABLE IF NOT EXISTS Vulnerabilities (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    HostID INTEGER,
    CVE TEXT,
    CvssBaseScore REAL,
    Description TEXT,
    ExploitAvailable BOOLEAN,
    ExploitFrameworkCanvas BOOLEAN,
    ExploitFrameworkCore BOOLEAN,
    ExploitFrameworkMetasploit BOOLEAN,
    ExploitEase TEXT,
    MetasploitName TEXT,
    PluginName TEXT,
    RiskFactor TEXT,
    Solution TEXT,
    Synopsis TEXT,
    PluginOutput TEXT,
    SeeAlso TEXT,
    ExploitedMalware BOOLEAN,
    Severity INTEGER,
    PluginID INTEGER,
    FOREIGN KEY(HostID) REFERENCES HostInfo(ID)
);`

	// Tüm tabloları oluştur
	tableCreationQueries := []string{createHostInfoTableSQL, createVulnerabilityTableSQL}
	for _, query := range tableCreationQueries {
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Error creating table: %v", err)
			return err
		}
	}

	return nil
}
