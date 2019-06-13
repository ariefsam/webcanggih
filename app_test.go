package main

import "testing"

func TestLoadConfig(t *testing.T) {
	config := LoadConfig()
	if config.Appname != "Webcanggih" {
		t.Errorf("App name tidak sama dengan 'Webcanggih'")
	}
	if config.Code != "12xx4" {
		t.Errorf("Code tidak sama")
	}
}
