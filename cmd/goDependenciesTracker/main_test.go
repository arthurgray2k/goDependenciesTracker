package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	var stdout, stderr bytes.Buffer

	// Test success case (CLI format) pointing to root module directory
	exitCode := run([]string{"-dir", "../.."}, &stdout, &stderr)
	if exitCode != 0 {
		t.Errorf("expected exit code 0, got %d", exitCode)
	}

	// Test invalid directory
	stdout.Reset()
	stderr.Reset()
	exitCode = run([]string{"-dir", "/invalid/dir/does/not/exist/999"}, &stdout, &stderr)
	if exitCode == 0 {
		t.Errorf("expected non-zero exit code for invalid dir, got %d", exitCode)
	}

	// Test invalid flag
	stdout.Reset()
	stderr.Reset()
	exitCode = run([]string{"-invalidflag"}, &stdout, &stderr)
	if exitCode != 2 {
		t.Errorf("expected exit code 2 for invalid flag, got %d", exitCode)
	}

	// Test unknown format default
	stdout.Reset()
	stderr.Reset()
	exitCode = run([]string{"-dir", "../..", "-format", "unknown"}, &stdout, &stderr)
	if exitCode != 0 {
		t.Errorf("expected exit code 0 for unknown format (fallback to txt), got %d", exitCode)
	}
	if !bytes.Contains(stderr.Bytes(), []byte("Unknown format")) {
		t.Errorf("expected unknown format warning in stderr")
	}

	// Test formats
	formats := []string{"json", "csv", "md"}
	for _, format := range formats {
		stdout.Reset()
		stderr.Reset()
		if code := run([]string{"-dir", "../..", "-format", format}, &stdout, &stderr); code != 0 {
			t.Errorf("expected exit code 0 for format %s, got %d", format, code)
		}
		if stdout.Len() == 0 {
			t.Errorf("expected output for format %s", format)
		}
	}
}
