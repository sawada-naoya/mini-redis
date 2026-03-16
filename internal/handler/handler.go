package handler

import (
	"fmt"

	"github.com/sawada-naoya/mini-redis/internal/command"

	"github.com/sawada-naoya/mini-redis/internal/store"
)

type Handler struct {
	store *store.Store
}

func New(s *store.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) Execute(cmd command.Command) string {
	switch cmd.Name {
	case "PING":
		if len(cmd.Args) != 0 {
			return "ERR wrong number of arguments for PING\n"
		}
		return "PONG\n"

	case "SET":
		if len(cmd.Args) != 2 {
			return "ERR wrong number of arguments for SET\n"
		}
		key := cmd.Args[0]
		value := cmd.Args[1]
		h.store.Set(key, value)
		return "OK\n"

	case "GET":
		if len(cmd.Args) != 1 {
			return "ERR wrong number of arguments for GET\n"
		}

		key := cmd.Args[0]
		value, ok := h.store.Get(key)
		if !ok {
			return "(nil)\n"
		}

		return fmt.Sprintf("%s\n", value)

	case "DEL":
		if len(cmd.Args) != 1 {
			return "ERR wrong number of arguments for DEL\n"
		}

		key := cmd.Args[0]
		if h.store.Del(key) {
			return "1\n"
		}

		return "0\n"

	case "EXISTS":
		if len(cmd.Args) != 1 {
			return "ERR wrong number of arguments for EXISTS\n"
		}

		key := cmd.Args[0]
		if h.store.Exists(key) {
			return "1\n"
		}

		return "0\n"

	default:
		return "ERR unknown command\n"
	}
}
