package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnumerateCidr(t *testing.T) {
	t.Run("single ip", func(t *testing.T) {
		ips, err := EnumerateCidr("192.168.0.1/32")
		assert.Nil(t, err)
		assert.Len(t, ips, 1)
		assert.Equal(t,"192.168.0.1", ips[0])
	})
	t.Run("slash 30", func(t *testing.T) {
		ips, err := EnumerateCidr("172.16.0.8/30")
		assert.Nil(t, err)
		assert.Len(t, ips, 4)
		assert.Equal(t,"172.16.0.8", ips[0])
		assert.Equal(t,"172.16.0.9", ips[1])
		assert.Equal(t,"172.16.0.10", ips[2])
		assert.Equal(t,"172.16.0.11", ips[3])
	})
}

func Test_EnumeratePorts(t *testing.T) {
	t.Run("single port", func(t *testing.T) {
		ports, err := EnumeratePorts("3333")
		assert.Nil(t, err)
		assert.Len(t, ports, 1)
		assert.Equal(t, 3333, ports[0])
	})
	t.Run("comma seperated ports", func(t *testing.T) {
		ports, err := EnumeratePorts("3333,4444")
		assert.Nil(t, err)
		assert.Len(t, ports, 2)
		assert.Equal(t, 3333, ports[0])
		assert.Equal(t, 4444, ports[1])
	})
	t.Run("block", func(t *testing.T) {
		ports, err := EnumeratePorts("1000-1005")
		assert.Nil(t, err)
		assert.Len(t, ports, 6)
		assert.Equal(t, 1000, ports[0])
		assert.Equal(t, 1001, ports[1])
		assert.Equal(t, 1002, ports[2])
		assert.Equal(t, 1003, ports[3])
		assert.Equal(t, 1004, ports[4])
		assert.Equal(t, 1005, ports[5])
	})
	t.Run("mixed single and block", func(t *testing.T) {
		ports, err := EnumeratePorts("1000-1005,3333")
		assert.Nil(t, err)
		assert.Len(t, ports, 7)
		assert.Equal(t, 1000, ports[0])
		assert.Equal(t, 1001, ports[1])
		assert.Equal(t, 1002, ports[2])
		assert.Equal(t, 1003, ports[3])
		assert.Equal(t, 1004, ports[4])
		assert.Equal(t, 1005, ports[5])
		assert.Equal(t, 3333, ports[6])
	})
	t.Run("remove duplicates", func(t *testing.T) {
		ports, err := EnumeratePorts("1000-1005,1003")
		assert.Nil(t, err)
		assert.Len(t, ports, 6)
		assert.Equal(t, 1000, ports[0])
		assert.Equal(t, 1001, ports[1])
		assert.Equal(t, 1002, ports[2])
		assert.Equal(t, 1003, ports[3])
		assert.Equal(t, 1004, ports[4])
		assert.Equal(t, 1005, ports[5])
	})
}