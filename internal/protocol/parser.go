package protocol

import (
	"errors"
	"strings"

	"github.com/sawada-naoya/mini-redis/internal/command"
)

var ErrEmptyCommand = errors.New("empty command")

func ParseLine(line string) (command.Command, error) {
	line = strings.TrimSpace(line)

	if line == "" {
		return command.Command{}, ErrEmptyCommand
	}

	parts := strings.Fields(line)
	if len(parts) == 0 {
		return command.Command{}, ErrEmptyCommand
	}

	return command.Command{
		Name: strings.ToUpper(parts[0]),
		Args: parts[1:],
	}, nil
}
