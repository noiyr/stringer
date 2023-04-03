package utils

import (
	"strings"
	_"strconv"
)

type IPAddr [4]byte
type IPv6 [6]byte

func (ip IPAddr) String() string {
	s := make([]string, len(ip))
	
	
	return strings.Join(s, ".")
}

