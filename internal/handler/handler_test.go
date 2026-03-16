package handler

import (
	"testing"

	"github.com/sawada-naoya/mini-redis/internal/command"
	"github.com/sawada-naoya/mini-redis/internal/store"
)

func TestHandlerExecutePing(t *testing.T) {
	t.Parallel()

	h := New(store.New())

	got := h.Execute(command.Command{
		Name: "PING",
		Args: []string{},
	})

	if got != "PONG\n" {
		t.Fatalf("got %q, want %q", got, "PONG\n")
	}
}

func TestHandlerExecuteSetAndGet(t *testing.T) {
	t.Parallel()

	h := New(store.New())

	setRes := h.Execute(command.Command{
		Name: "SET",
		Args: []string{"foo", "bar"},
	})

	if setRes != "OK\n" {
		t.Fatalf("set response got %q, want %q", setRes, "OK\n")
	}

	getRes := h.Execute(command.Command{
		Name: "GET",
		Args: []string{"foo"},
	})

	if getRes != "bar\n" {
		t.Fatalf("get response got %q, want %q", getRes, "bar\n")
	}
}

func TestHandlerExecuteGetMissing(t *testing.T) {
	t.Parallel()

	h := New(store.New())

	got := h.Execute(command.Command{
		Name: "GET",
		Args: []string{"missing"},
	})

	if got != "(nil)\n" {
		t.Fatalf("got %q, want %q", got, "(nil)\n")
	}
}

func TestHandlerExecuteDel(t *testing.T) {
	t.Parallel()

	h := New(store.New())

	h.Execute(command.Command{
		Name: "SET",
		Args: []string{"foo", "bar"},
	})

	delResp := h.Execute(command.Command{
		Name: "DEL",
		Args: []string{"foo"},
	})

	if delResp != "1\n" {
		t.Fatalf("got %q, want %q", delResp, "1\n")
	}

	getResp := h.Execute(command.Command{
		Name: "GET",
		Args: []string{"foo"},
	})

	if getResp != "(nil)\n" {
		t.Fatalf("got %q, want %q", getResp, "(nil)\n")
	}
}

func TestHandlerExecuteExists(t *testing.T) {
	t.Parallel()

	h := New(store.New())

	h.Execute(command.Command{
		Name: "SET",
		Args: []string{"foo", "bar"},
	})

	existsResp := h.Execute(command.Command{
		Name: "EXISTS",
		Args: []string{"foo"},
	})

	if existsResp != "1\n" {
		t.Fatalf("got %q, want %q", existsResp, "1\n")
	}

	missingResp := h.Execute(command.Command{
		Name: "EXISTS",
		Args: []string{"missing"},
	})

	if missingResp != "0\n" {
		t.Fatalf("got %q, want %q", missingResp, "0\n")
	}
}

func TestHandlerExecuteUnknownCommand(t *testing.T) {
	t.Parallel()

	h := New(store.New())

	got := h.Execute(command.Command{
		Name: "UNKNOWN",
		Args: []string{},
	})

	if got != "ERR unknown command\n" {
		t.Fatalf("got %q, want %q", got, "ERR unknown command\n")
	}
}

func TestHandlerExecuteWrongArgs(t *testing.T) {
	t.Parallel()

	h := New(store.New())

	got := h.Execute(command.Command{
		Name: "SET",
		Args: []string{"foo"},
	})

	if got != "ERR wrong number of arguments for SET\n" {
		t.Fatalf("got %q, want %q", got, "ERR wrong number of arguments for SET\n")
	}
}
