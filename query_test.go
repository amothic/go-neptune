package neptune

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Query(t *testing.T) {
	client, _ := Open("http://localhost:8182")
	tests := []struct {
		name    string
		query   string
		wantErr bool
	}{
		{
			name:  "success",
			query: `g.V().count()`,
		},
		{
			name:    "failure",
			query:   `g.V().count`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := client.Query(tt.query)
			if !tt.wantErr {
				assert.NotNil(t, res)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
