package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpv4Only(t *testing.T) {
	t.Run("single ipv4 network", func(t *testing.T) {
		filtered := Ipv4Only([]string {
			"127.0.0.1/32",
		})
		assert.Len(t, filtered, 1)
		assert.Equal(t, "127.0.0.1/32", filtered[0])
	})
	t.Run("single ipv6 network", func(t *testing.T) {
		filtered := Ipv4Only([]string {
			"::1/128",
		})
		assert.Len(t, filtered, 0)
	})
	t.Run("mixed ipv4/ipv6 networks", func(t *testing.T) {
		filtered := Ipv4Only([]string {
			"127.0.0.1/32",
			"::1/128",
		})
		assert.Len(t, filtered, 1)
		assert.Equal(t, "127.0.0.1/32", filtered[0])
	})
}

func TestIpv6Only(t *testing.T) {
	t.Run("single ipv4 network", func(t *testing.T) {
		filtered := Ipv6Only([]string {
			"127.0.0.1/32",
		})
		assert.Len(t, filtered, 0)
	})
	t.Run("single ipv6 network", func(t *testing.T) {
		filtered := Ipv6Only([]string {
			"::1/128",
		})
		assert.Len(t, filtered, 1)
		assert.Equal(t, "::1/128", filtered[0])
	})
	t.Run("mixed ipv4/ipv6 networks", func(t *testing.T) {
		filtered := Ipv6Only([]string {
			"127.0.0.1/32",
			"::1/128",
		})
		assert.Len(t, filtered, 1)
		assert.Equal(t, "::1/128", filtered[0])
	})
}