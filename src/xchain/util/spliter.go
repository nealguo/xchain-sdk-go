package util

import (
	"log"
	"math"
	"strconv"
	"strings"
	"xchain-sdk-go/src/xchain/config"
)

func SplitToNodes(value string) []*config.Node {
	if value == "" {
		return nil
	}
	var parts []string
	if strings.Contains(value, ";") {
		parts = strings.Split(value, ";")
	} else {
		parts = []string{value}
	}
	if parts == nil || len(parts) == 0 {
		return nil
	}
	var nodes []*config.Node
	for _, part := range parts {
		if strings.Contains(part, ":") {
			innerParts := strings.Split(part, ":")
			portValue, err := strconv.Atoi(strings.TrimSpace(innerParts[1]))
			if err != nil {
				log.Printf("Wrong port of node:%s, err:%v", part, err)
				continue
			}
			if portValue > math.MaxUint16 {
				log.Printf("Wrong port of node:%s, port is too large, port:%v", part, portValue)
				continue
			}
			host := strings.TrimSpace(innerParts[0])
			port := uint16(portValue)
			node := &config.Node{Host: host, Port: port}
			nodes = append(nodes, node)
		}
	}
	return nodes
}
