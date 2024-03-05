package models

import "gorm.io/gorm"

type ZapWebSite struct {
	gorm.Model
	Uid                   uint   `json:"uid" gorm:"index"`
	Title                 string `json:"title"`
	Description           string `json:"description"`
	Domain                string `json:"domain"`
	DomainNames           string `json:"domain_names"` // domain names   (空格分开)
	WwwRoot               string `json:"www_root"`
	Application           string `json:"application"`             // static or php
	ApplicationId         int    `json:"application_id"`          // apps_id
	ApplicationVersion    string `json:"application_version"`     // app version
	ApplicationExpose     string `json:"application_expose"`      // app expose
	ApplicationExposeName string `json:"application_expose_name"` // app expose name  (unix | http | https )
	IndexFiles            string `json:"index_files"`
	RunDirectory          string `json:"run_directory"`
	SslCertificate        string `json:"ssl_certificate"`
	SslCertificateKey     string `json:"ssl_certificate_key"`
	SslOcsp               string `json:"ssl_ocsp"`
	ForceHttps            string `json:"force_https"`
	SslProtocols          string `json:"ssl_protocols"`
	Settings              string `json:"settings"`             // properties
	Status                string `json:"status" gorm:"index;"` // 状态 active
}
