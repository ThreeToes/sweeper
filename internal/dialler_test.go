package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheck(t *testing.T) {
	t.Run("check up", func(t *testing.T) {
		up, err := Check(&DialSpec{
			Ip:      "1.1.1.1",
			Port:    53,
			Timeout: 100,
		})
		assert.Nil(t, err)
		assert.True(t, up)
	})
	t.Run("check valid server with closed port", func(t *testing.T) {
		up, err := Check(&DialSpec{
			Ip:      "1.1.1.1",
			Port:    60,
			Timeout: 100,
		})
		assert.Nil(t, err)
		assert.False(t, up)
	})
	t.Run("check invalid server", func(t *testing.T) {
		up, err := Check(&DialSpec{
			Ip:      "notanip",
			Port:    60,
			Timeout: 100,
		})
		assert.NotNil(t, err)
		assert.False(t, up)
	})
}