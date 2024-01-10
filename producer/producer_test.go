package producer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Mat-Greaves/go-modules-test/producer"
)

func TestProduce(t *testing.T) {
	assert.Equal(t, "produced!", producer.Produce())
}
