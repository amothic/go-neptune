package neptune

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "http://localhost:8182",
			want: "http://localhost:8182/gremlin",
		},
		{
			in:   "https://neptune.amazonaws.com:8182",
			want: "https://neptune.amazonaws.com:8182/gremlin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			client, err := Open(tt.in)
			require.NoError(t, err)
			assert.Equal(t, tt.want, client.endpoint)
		})
	}
}
