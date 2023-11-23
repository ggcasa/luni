// Package luni key value store
package luni

import (
	"net"
	"time"
)

// Conn
type Conn interface {
	DialTcp(ip net.IP, p string) (b bool, err error)
	DialUdp(ip net.IP, p string) (b bool, err error)
}

// Zi
type Zi interface {
	Add(k []byte, t time.Time) (b bool, err error)
	Get(k []byte) (n int, err error)
	Find(k []byte, s string)
	Delete(k []byte) (s string, err error)
}
