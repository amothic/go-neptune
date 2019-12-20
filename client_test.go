package neptune

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	client, err := Open("https://neptune.amazonaws.com:8182/gremlin")
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
