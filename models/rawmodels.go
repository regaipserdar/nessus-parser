package models

import "encoding/xml"

type ReportHost struct {
	XMLName        xml.Name       `xml:"ReportHost"`
	HostProperties HostProperties `xml:"HostProperties"`
	ReportItemList []ReportItem   `xml:"ReportItem"`
}

type HostProperties struct {
	XMLName xml.Name   `xml:"HostProperties"`
	Info    []InfoList `xml:"tag"`
}

type InfoList struct {
	Key   string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type ReportItem struct {
	CVE                        []string `xml:"cve"`
	CvssBaseScore              float64  `xml:"cvss_base_score"`
	Desciption                 string   `xml:"description"`
	ExploitAvailable           bool     `xml:"exploit_available"`
	ExploitFrameworkCanvas     bool     `xml:"exploit_framework_canvas"`
	ExploitFrameworkCore       bool     `xml:"exploit_framework_core"`
	ExploitFrameworkMetasploit bool     `xml:"exploit_framework_metasploit"`
	ExploitEase                string   `xml:"exploit_ease"`
	MetasploitName             string   `xml:"metasploit_name"`
	PluginName                 string   `xml:"plugin_name"`
	RiskFactor                 string   `xml:"risk_factor"`
	Solution                   string   `xml:"solution"`
	Synopsis                   string   `xml:"synopsis"`
	PluginOutput               string   `xml:"plugin_output"`
	SeeAlso                    string   `xml:"see_also"`
	ExploitedMalware           bool     `xml:"exploited_by_malware"`
	Severity                   int      `xml:"severity,attr"`
	PluginID                   int      `xml:"pluginID,attr"`
}

type ReadyData struct {
	IP                string
	OS                string
	MAC               string
	FQDN              string
	Netbios           string
	CVE               []string
	CVSS              float64
	Desc              string
	ExploitAvailable  bool
	ExploitCanvas     bool
	ExploitCore       bool
	ExploitMetasploit bool
	ExploitEase       string
	MetasploitName    string
	PluginName        string
	Risk              string
	Solution          string
	Synopsis          string
	PluginOutput      string
	SeeAlso           string
	ExploitedMalware  bool
	Severity          int
	PluginID          int
}
