package syncmap

import (
	"testing"
)

// TestLoad tests our ability to load generic data from the Map.
func TestLoad(t *testing.T) {
	testMap := New[string, string]()

	data, ok := testMap.Load("foo")
	if ok != false || data != "" {
		t.Error("load should not return data from empty map")
	}
}

// TestDelete tests our ability to delete generic data from the Map.
func TestDelete(t *testing.T) {
	testMap := New[string, string]()

	testMap.Store("foo", "bar")
	testMap.Delete("foo")
	data, ok := testMap.Load("foo")

	if ok != false || data != "" {
		t.Error("load should not return data from empty map")
	}
}

// TestStore tests our ability to store generic data into the Map.
func TestStore(t *testing.T) {
	testMap := New[string, string]()

	testMap.Store("foo", "bar")
	testMap.Store("baz", "buz")

	data, ok := testMap.Load("foo")
	if ok != true || data != "bar" {
		t.Error("load should return correct data from map")
	}

	data, ok = testMap.Load("baz")
	if ok != true || data != "buz" {
		t.Error("load should return correct data from map")
	}
}

// TestLoadOrStore tests our ability to load or store generic data from/into the Map.
func TestLoadOrStore(t *testing.T) {
	testMap := New[string, string]()

	testMap.Store("foo", "bar2")

	data, ok := testMap.LoadOrStore("foo", "bar")
	if ok != true || data != "bar2" {
		t.Error("load should return correct data from map")
	}

	data, ok = testMap.LoadOrStore("baz", "buz")
	if ok != false || data != "buz" {
		t.Error("load should return correct data from map")
	}
}

// TestLoadOrDelete tests our ability to load or delete generic data from the Map.
func TestLoadOrDelete(t *testing.T) {
	testMap := New[string, string]()

	testMap.Store("foo", "bar")

	data, ok := testMap.LoadAndDelete("foo")
	if ok != true || data != "bar" {
		t.Error("load should return correct data from map")
	}

	data, ok = testMap.Load("foo")
	if ok != false && data != "" {
		t.Error("data should not exist in map")
	}

	data, ok = testMap.LoadAndDelete("foo2")
	if ok != false || data != "" {
		t.Error("no data should return from map")
	}
}

// TestRange tests our ability to iterate over generic data from the Map.
func TestRange(t *testing.T) {
	testMap := New[string, string]()

	testMap.Store("foo", "bar")
	testMap.Store("baz", "buz")

	testMap.Range(func(a, b string) bool {
		if a == "foo" || a == "baz" {
			if b == "bar" || b == "buz" {
				return true
			}
		}

		t.Error("expected range data")

		return true
	})
}
