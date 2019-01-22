package smpp

import (
	"bytes"
	"github.com/linxGnu/gosmpp"
	"github.com/linxGnu/gosmpp/Exception"
	"github.com/linxGnu/gosmpp/Utils"
)

func (w *worker) Add(addr, name string, port int) error {
	var err error
	w.name = name
	w.connection[w.name], err = gosmpp.NewTCPIPConnectionWithAddrPort(addr, port)
	return err
}

func (w *worker) Delete(name string) {
	delete(w.connection, name)
}

func (w *worker) Send(payload *bytes.Buffer) *Exception.Exception {
	var data *Utils.ByteBuffer
	data.Buffer = payload
	return w.connection[w.name].Send(data)
}
