package config

import "strconv"

type Node struct {
	Host string
	Port uint16
}

func (node *Node) Address() string {
	addr := ""
	if node == nil {
		return addr
	}
	addr = node.Host + ":" + strconv.Itoa(int(node.Port))
	return addr
}
