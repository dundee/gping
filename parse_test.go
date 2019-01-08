package main

import "testing"

func TestParseDestination(t *testing.T) {
	assertEquals(t, parseDestination("example.com"), "example.com")
	assertEquals(t, parseDestination("http://example.com"), "example.com")
	assertEquals(t, parseDestination("http://example.com/"), "example.com")
	assertEquals(t, parseDestination("http://example.com/?aaa=bbb"), "example.com")
	assertEquals(t, parseDestination("https://example.com/?aaa=bbb"), "example.com")
}

func assertEquals(t *testing.T, res string, expected string) {
	if res != expected {
		t.Errorf("Test failed: '%s' expected, but '%s' returned", expected, res)
	}
}
