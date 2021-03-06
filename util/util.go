package util

import (
	"net"
	"strings"
	"time"
)

const loopbackIP = "127.0.0.1"

// SelectIntOpt takes an option and a default value and returns the default value if
// the option is equal to zero, and the option otherwise.
func SelectIntOpt(opt, def int) int {
	if opt == 0 {
		return def
	}
	return opt
}

// SelectDurationOpt takes an option and a default value and returns the default value if
// the option is equal to zero, and the option otherwise.
func SelectDurationOpt(opt, def time.Duration) time.Duration {
	if opt == time.Duration(0) {
		return def
	}
	return opt
}

// GetLocalIP returns the local IP address.
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return loopbackIP
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return loopbackIP
}

// SafeSplit splits the give string by space and handles quotation marks
func SafeSplit(s string) []string {
	split := strings.Split(s, " ")

	var result []string
	var inquote string
	var block string
	for _, i := range split {
		if inquote == "" {
			if strings.HasPrefix(i, "'") || strings.HasPrefix(i, "\"") {
				inquote = string(i[0])
				block = strings.TrimPrefix(i, inquote) + " "
			} else {
				result = append(result, i)
			}
		} else {
			if !strings.HasSuffix(i, inquote) {
				block += i + " "
			} else {
				block += strings.TrimSuffix(i, inquote)
				inquote = ""
				result = append(result, block)
				block = ""
			}
		}
	}

	return result
}
