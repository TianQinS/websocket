// Code generated by automatic for 'github.com/TianQinS/evio'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"github.com/TianQinS/evio"
	"net"
	"reflect"
)

func init() {
	Symbols["github.com/TianQinS/evio"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Close":            reflect.ValueOf(evio.Close),
		"Detach":           reflect.ValueOf(evio.Detach),
		"Handshake":        reflect.ValueOf(evio.Handshake),
		"LeastConnections": reflect.ValueOf(evio.LeastConnections),
		"None":             reflect.ValueOf(evio.None),
		"Random":           reflect.ValueOf(evio.Random),
		"RoundRobin":       reflect.ValueOf(evio.RoundRobin),
		"Serve":            reflect.ValueOf(evio.Serve),
		"Shutdown":         reflect.ValueOf(evio.Shutdown),

		// type definitions
		"Action":      reflect.ValueOf((*evio.Action)(nil)),
		"Conn":        reflect.ValueOf((*evio.Conn)(nil)),
		"Events":      reflect.ValueOf((*evio.Events)(nil)),
		"InputStream": reflect.ValueOf((*evio.InputStream)(nil)),
		"LoadBalance": reflect.ValueOf((*evio.LoadBalance)(nil)),
		"Options":     reflect.ValueOf((*evio.Options)(nil)),
		"Server":      reflect.ValueOf((*evio.Server)(nil)),

		// interface wrapper definitions
		"_Conn": reflect.ValueOf((*_github_com_TianQinS_evio_Conn)(nil)),
	}
}

// _github_com_TianQinS_evio_Conn is an interface wrapper for Conn type
type _github_com_TianQinS_evio_Conn struct {
	WAddrIndex  func() int
	WContext    func() interface{}
	WLocalAddr  func() net.Addr
	WRemoteAddr func() net.Addr
	WSetContext func(a0 interface{})
	WWake       func()
	WWrite      func(p []byte) (n int, err error)
}

func (W _github_com_TianQinS_evio_Conn) AddrIndex() int                    { return W.WAddrIndex() }
func (W _github_com_TianQinS_evio_Conn) Context() interface{}              { return W.WContext() }
func (W _github_com_TianQinS_evio_Conn) LocalAddr() net.Addr               { return W.WLocalAddr() }
func (W _github_com_TianQinS_evio_Conn) RemoteAddr() net.Addr              { return W.WRemoteAddr() }
func (W _github_com_TianQinS_evio_Conn) SetContext(a0 interface{})         { W.WSetContext(a0) }
func (W _github_com_TianQinS_evio_Conn) Wake()                             { W.WWake() }
func (W _github_com_TianQinS_evio_Conn) Write(p []byte) (n int, err error) { return W.WWrite(p) }