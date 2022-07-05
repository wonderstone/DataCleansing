package datacleaner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadData(t *testing.T) {
	assert.Equal(t, "Hello, World!", ReadData("/Users/wonderstone/go/tmp/finance/config"))

}
