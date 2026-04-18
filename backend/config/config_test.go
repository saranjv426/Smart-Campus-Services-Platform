package config

import "testing"

func TestDatabasePathDefaultValue(t *testing.T) {
	if DatabasePath == "" {
		t.Fatal("expected DatabasePath to be configured")
	}

	if DatabasePath != "data/smart_campus.db" {
		t.Fatalf("expected default database path %q, got %q", "data/smart_campus.db", DatabasePath)
	}
}
