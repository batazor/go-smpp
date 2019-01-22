package smpp

import "github.com/linxGnu/gosmpp"

type worker struct {
	name       string
	connection map[string]*gosmpp.TCPIPConnection
}
