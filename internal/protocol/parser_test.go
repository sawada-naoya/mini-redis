package protocol

import (
	"reflect"
	"testing"

	"github.com/sawada-naoya/mini-redis/internal/command"
)

func TestParseLine(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    command.Command
		wantErr bool
	}{
		{
			name:  "PING command",
			input: "PING\n",
			want: command.Command{
				Name: "PING",
				Args: []string{},
			},
			wantErr: false,
		},
		{
			name:  "SET command",
			input: "SET foo bar\n",
			want: command.Command{
				Name: "SET",
				Args: []string{"foo", "bar"},
			},
			wantErr: false,
		},
		{
			name:  "lowercase command is normalized",
			input: "get foo\n",
			want: command.Command{
				Name: "GET",
				Args: []string{"foo"},
			},
			wantErr: false,
		},
		{
			name:    "empty line",
			input:   "\n",
			wantErr: true,
		},
		{
			name:    "spaces only",
			input:   "   \n",
			wantErr: true,
		},
		{
			name:  "multiple spaces are ignored",
			input: "SET   foo    bar\n",
			want: command.Command{
				Name: "SET",
				Args: []string{"foo", "bar"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseLine(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got = %#v, want = %#v", got, tt.want)
			}
		})
	}
}
