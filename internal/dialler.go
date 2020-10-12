package internal

import (
	"fmt"
	"net"
	"time"
)

type DialSpec struct {
	Ip string
	Port int
	Timeout int
}

func Check(spec *DialSpec) (bool, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d",spec.Ip,spec.Port),
		time.Duration(spec.Timeout) * time.Millisecond)
	if err != nil {
		e, ok := err.(net.Error)
		if !ok {
			return false, err
		}
		if e.Timeout() {
			return false, nil
		}
		return false, e
	}
	conn.Close()
	return true, nil
}
