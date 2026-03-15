package store

import "testing"

func TestStoreSetAndGet(t *testing.T) {
	t.Parallel()

	s := New()
	s.Set("foo", "bar")

	got, ok := s.Get("foo")
	if !ok {
		t.Fatalf("expected key to exist")
	}

	if got != "bar" {
		t.Fatalf("got %q, want %q", got, "bar")
	}
}

func TestStoreGetNotFound(t *testing.T) {
	t.Parallel()

	s := New()

	got, ok := s.Get("missing")
	if ok {
		t.Fatalf("expected key not to exist, got value %q", got)
	}
}

func TestStoreDel(t *testing.T) {
	t.Parallel()

	s := New()
	s.Set("foo", "bar")

	deleted := s.Del("foo")
	if !deleted {
		t.Fatalf("expected delete to return true")
	}

	_, ok := s.Get("foo")
	if ok {
		t.Fatalf("expected key to be deleted")
	}
}

func TestStoreDelNotFound(t *testing.T) {
	t.Parallel()

	s := New()

	deleted := s.Del("missing")
	if deleted {
		t.Fatalf("expected delete to return false")
	}
}

func TestStoreExists(t *testing.T) {
	t.Parallel()

	s := New()
	s.Set("foo", "bar")

	if !s.Exists("foo") {
		t.Fatalf("expected key to exist")
	}

	if s.Exists("missing") {
		t.Fatalf("expected missing key not to exist")
	}
}
