// Code generated by automatic for 'github.com/gobwas/ws'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"github.com/gobwas/ws"
	"io"
	"reflect"
)

func init() {
	Symbols["github.com/gobwas/ws"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CheckCloseFrameData":                   reflect.ValueOf(ws.CheckCloseFrameData),
		"CheckHeader":                           reflect.ValueOf(ws.CheckHeader),
		"Cipher":                                reflect.ValueOf(ws.Cipher),
		"CompileFrame":                          reflect.ValueOf(ws.CompileFrame),
		"CompiledClose":                         reflect.ValueOf(&ws.CompiledClose).Elem(),
		"CompiledCloseGoingAway":                reflect.ValueOf(&ws.CompiledCloseGoingAway).Elem(),
		"CompiledCloseInternalServerError":      reflect.ValueOf(&ws.CompiledCloseInternalServerError).Elem(),
		"CompiledCloseInvalidFramePayloadData":  reflect.ValueOf(&ws.CompiledCloseInvalidFramePayloadData).Elem(),
		"CompiledCloseMandatoryExt":             reflect.ValueOf(&ws.CompiledCloseMandatoryExt).Elem(),
		"CompiledCloseMessageTooBig":            reflect.ValueOf(&ws.CompiledCloseMessageTooBig).Elem(),
		"CompiledCloseNoMeaningYet":             reflect.ValueOf(&ws.CompiledCloseNoMeaningYet).Elem(),
		"CompiledCloseNormalClosure":            reflect.ValueOf(&ws.CompiledCloseNormalClosure).Elem(),
		"CompiledClosePolicyViolation":          reflect.ValueOf(&ws.CompiledClosePolicyViolation).Elem(),
		"CompiledCloseProtocolError":            reflect.ValueOf(&ws.CompiledCloseProtocolError).Elem(),
		"CompiledCloseTLSHandshake":             reflect.ValueOf(&ws.CompiledCloseTLSHandshake).Elem(),
		"CompiledCloseUnsupportedData":          reflect.ValueOf(&ws.CompiledCloseUnsupportedData).Elem(),
		"CompiledPing":                          reflect.ValueOf(&ws.CompiledPing).Elem(),
		"CompiledPong":                          reflect.ValueOf(&ws.CompiledPong).Elem(),
		"DefaultClientReadBufferSize":           reflect.ValueOf(ws.DefaultClientReadBufferSize),
		"DefaultClientWriteBufferSize":          reflect.ValueOf(ws.DefaultClientWriteBufferSize),
		"DefaultDialer":                         reflect.ValueOf(&ws.DefaultDialer).Elem(),
		"DefaultHTTPUpgrader":                   reflect.ValueOf(&ws.DefaultHTTPUpgrader).Elem(),
		"DefaultServerReadBufferSize":           reflect.ValueOf(ws.DefaultServerReadBufferSize),
		"DefaultServerWriteBufferSize":          reflect.ValueOf(ws.DefaultServerWriteBufferSize),
		"DefaultUpgrader":                       reflect.ValueOf(&ws.DefaultUpgrader).Elem(),
		"Dial":                                  reflect.ValueOf(ws.Dial),
		"ErrHandshakeBadConnection":             reflect.ValueOf(&ws.ErrHandshakeBadConnection).Elem(),
		"ErrHandshakeBadExtensions":             reflect.ValueOf(&ws.ErrHandshakeBadExtensions).Elem(),
		"ErrHandshakeBadHost":                   reflect.ValueOf(&ws.ErrHandshakeBadHost).Elem(),
		"ErrHandshakeBadMethod":                 reflect.ValueOf(&ws.ErrHandshakeBadMethod).Elem(),
		"ErrHandshakeBadProtocol":               reflect.ValueOf(&ws.ErrHandshakeBadProtocol).Elem(),
		"ErrHandshakeBadSecAccept":              reflect.ValueOf(&ws.ErrHandshakeBadSecAccept).Elem(),
		"ErrHandshakeBadSecKey":                 reflect.ValueOf(&ws.ErrHandshakeBadSecKey).Elem(),
		"ErrHandshakeBadSecVersion":             reflect.ValueOf(&ws.ErrHandshakeBadSecVersion).Elem(),
		"ErrHandshakeBadStatus":                 reflect.ValueOf(&ws.ErrHandshakeBadStatus).Elem(),
		"ErrHandshakeBadSubProtocol":            reflect.ValueOf(&ws.ErrHandshakeBadSubProtocol).Elem(),
		"ErrHandshakeBadUpgrade":                reflect.ValueOf(&ws.ErrHandshakeBadUpgrade).Elem(),
		"ErrHandshakeUpgradeRequired":           reflect.ValueOf(&ws.ErrHandshakeUpgradeRequired).Elem(),
		"ErrHeaderLengthMSB":                    reflect.ValueOf(&ws.ErrHeaderLengthMSB).Elem(),
		"ErrHeaderLengthUnexpected":             reflect.ValueOf(&ws.ErrHeaderLengthUnexpected).Elem(),
		"ErrMalformedRequest":                   reflect.ValueOf(&ws.ErrMalformedRequest).Elem(),
		"ErrMalformedResponse":                  reflect.ValueOf(&ws.ErrMalformedResponse).Elem(),
		"ErrNotHijacker":                        reflect.ValueOf(&ws.ErrNotHijacker).Elem(),
		"ErrProtocolContinuationExpected":       reflect.ValueOf(&ws.ErrProtocolContinuationExpected).Elem(),
		"ErrProtocolContinuationUnexpected":     reflect.ValueOf(&ws.ErrProtocolContinuationUnexpected).Elem(),
		"ErrProtocolControlNotFinal":            reflect.ValueOf(&ws.ErrProtocolControlNotFinal).Elem(),
		"ErrProtocolControlPayloadOverflow":     reflect.ValueOf(&ws.ErrProtocolControlPayloadOverflow).Elem(),
		"ErrProtocolInvalidUTF8":                reflect.ValueOf(&ws.ErrProtocolInvalidUTF8).Elem(),
		"ErrProtocolMaskRequired":               reflect.ValueOf(&ws.ErrProtocolMaskRequired).Elem(),
		"ErrProtocolMaskUnexpected":             reflect.ValueOf(&ws.ErrProtocolMaskUnexpected).Elem(),
		"ErrProtocolNonZeroRsv":                 reflect.ValueOf(&ws.ErrProtocolNonZeroRsv).Elem(),
		"ErrProtocolOpCodeReserved":             reflect.ValueOf(&ws.ErrProtocolOpCodeReserved).Elem(),
		"ErrProtocolStatusCodeApplicationLevel": reflect.ValueOf(&ws.ErrProtocolStatusCodeApplicationLevel).Elem(),
		"ErrProtocolStatusCodeNoMeaning":        reflect.ValueOf(&ws.ErrProtocolStatusCodeNoMeaning).Elem(),
		"ErrProtocolStatusCodeNotInUse":         reflect.ValueOf(&ws.ErrProtocolStatusCodeNotInUse).Elem(),
		"ErrProtocolStatusCodeUnknown":          reflect.ValueOf(&ws.ErrProtocolStatusCodeUnknown).Elem(),
		"HeaderSize":                            reflect.ValueOf(ws.HeaderSize),
		"MaskFrame":                             reflect.ValueOf(ws.MaskFrame),
		"MaskFrameInPlace":                      reflect.ValueOf(ws.MaskFrameInPlace),
		"MaskFrameInPlaceWith":                  reflect.ValueOf(ws.MaskFrameInPlaceWith),
		"MaskFrameWith":                         reflect.ValueOf(ws.MaskFrameWith),
		"MaxControlFramePayloadSize":            reflect.ValueOf(ws.MaxControlFramePayloadSize),
		"MaxHeaderSize":                         reflect.ValueOf(ws.MaxHeaderSize),
		"MinHeaderSize":                         reflect.ValueOf(ws.MinHeaderSize),
		"MustCompileFrame":                      reflect.ValueOf(ws.MustCompileFrame),
		"MustReadFrame":                         reflect.ValueOf(ws.MustReadFrame),
		"MustWriteFrame":                        reflect.ValueOf(ws.MustWriteFrame),
		"NewBinaryFrame":                        reflect.ValueOf(ws.NewBinaryFrame),
		"NewCloseFrame":                         reflect.ValueOf(ws.NewCloseFrame),
		"NewCloseFrameBody":                     reflect.ValueOf(ws.NewCloseFrameBody),
		"NewFrame":                              reflect.ValueOf(ws.NewFrame),
		"NewMask":                               reflect.ValueOf(ws.NewMask),
		"NewPingFrame":                          reflect.ValueOf(ws.NewPingFrame),
		"NewPongFrame":                          reflect.ValueOf(ws.NewPongFrame),
		"NewTextFrame":                          reflect.ValueOf(ws.NewTextFrame),
		"OpBinary":                              reflect.ValueOf(ws.OpBinary),
		"OpClose":                               reflect.ValueOf(ws.OpClose),
		"OpContinuation":                        reflect.ValueOf(ws.OpContinuation),
		"OpPing":                                reflect.ValueOf(ws.OpPing),
		"OpPong":                                reflect.ValueOf(ws.OpPong),
		"OpText":                                reflect.ValueOf(ws.OpText),
		"ParseCloseFrameData":                   reflect.ValueOf(ws.ParseCloseFrameData),
		"ParseCloseFrameDataUnsafe":             reflect.ValueOf(ws.ParseCloseFrameDataUnsafe),
		"PutCloseFrameBody":                     reflect.ValueOf(ws.PutCloseFrameBody),
		"PutReader":                             reflect.ValueOf(ws.PutReader),
		"ReadFrame":                             reflect.ValueOf(ws.ReadFrame),
		"ReadHeader":                            reflect.ValueOf(ws.ReadHeader),
		"RejectConnectionError":                 reflect.ValueOf(ws.RejectConnectionError),
		"RejectionHeader":                       reflect.ValueOf(ws.RejectionHeader),
		"RejectionReason":                       reflect.ValueOf(ws.RejectionReason),
		"RejectionStatus":                       reflect.ValueOf(ws.RejectionStatus),
		"Rsv":                                   reflect.ValueOf(ws.Rsv),
		"SelectEqual":                           reflect.ValueOf(ws.SelectEqual),
		"SelectFromSlice":                       reflect.ValueOf(ws.SelectFromSlice),
		"StateClientSide":                       reflect.ValueOf(ws.StateClientSide),
		"StateExtended":                         reflect.ValueOf(ws.StateExtended),
		"StateFragmented":                       reflect.ValueOf(ws.StateFragmented),
		"StateServerSide":                       reflect.ValueOf(ws.StateServerSide),
		"StatusAbnormalClosure":                 reflect.ValueOf(ws.StatusAbnormalClosure),
		"StatusGoingAway":                       reflect.ValueOf(ws.StatusGoingAway),
		"StatusInternalServerError":             reflect.ValueOf(ws.StatusInternalServerError),
		"StatusInvalidFramePayloadData":         reflect.ValueOf(ws.StatusInvalidFramePayloadData),
		"StatusMandatoryExt":                    reflect.ValueOf(ws.StatusMandatoryExt),
		"StatusMessageTooBig":                   reflect.ValueOf(ws.StatusMessageTooBig),
		"StatusNoMeaningYet":                    reflect.ValueOf(ws.StatusNoMeaningYet),
		"StatusNoStatusRcvd":                    reflect.ValueOf(ws.StatusNoStatusRcvd),
		"StatusNormalClosure":                   reflect.ValueOf(ws.StatusNormalClosure),
		"StatusPolicyViolation":                 reflect.ValueOf(ws.StatusPolicyViolation),
		"StatusProtocolError":                   reflect.ValueOf(ws.StatusProtocolError),
		"StatusRangeApplication":                reflect.ValueOf(&ws.StatusRangeApplication).Elem(),
		"StatusRangeNotInUse":                   reflect.ValueOf(&ws.StatusRangeNotInUse).Elem(),
		"StatusRangePrivate":                    reflect.ValueOf(&ws.StatusRangePrivate).Elem(),
		"StatusRangeProtocol":                   reflect.ValueOf(&ws.StatusRangeProtocol).Elem(),
		"StatusTLSHandshake":                    reflect.ValueOf(ws.StatusTLSHandshake),
		"StatusUnsupportedData":                 reflect.ValueOf(ws.StatusUnsupportedData),
		"Upgrade":                               reflect.ValueOf(ws.Upgrade),
		"UpgradeHTTP":                           reflect.ValueOf(ws.UpgradeHTTP),
		"WriteFrame":                            reflect.ValueOf(ws.WriteFrame),
		"WriteHeader":                           reflect.ValueOf(ws.WriteHeader),

		// type definitions
		"Dialer":                reflect.ValueOf((*ws.Dialer)(nil)),
		"Frame":                 reflect.ValueOf((*ws.Frame)(nil)),
		"HTTPUpgrader":          reflect.ValueOf((*ws.HTTPUpgrader)(nil)),
		"Handshake":             reflect.ValueOf((*ws.Handshake)(nil)),
		"HandshakeHeader":       reflect.ValueOf((*ws.HandshakeHeader)(nil)),
		"HandshakeHeaderBytes":  reflect.ValueOf((*ws.HandshakeHeaderBytes)(nil)),
		"HandshakeHeaderFunc":   reflect.ValueOf((*ws.HandshakeHeaderFunc)(nil)),
		"HandshakeHeaderHTTP":   reflect.ValueOf((*ws.HandshakeHeaderHTTP)(nil)),
		"HandshakeHeaderString": reflect.ValueOf((*ws.HandshakeHeaderString)(nil)),
		"Header":                reflect.ValueOf((*ws.Header)(nil)),
		"OpCode":                reflect.ValueOf((*ws.OpCode)(nil)),
		"ProtocolError":         reflect.ValueOf((*ws.ProtocolError)(nil)),
		"RejectOption":          reflect.ValueOf((*ws.RejectOption)(nil)),
		"State":                 reflect.ValueOf((*ws.State)(nil)),
		"StatusCode":            reflect.ValueOf((*ws.StatusCode)(nil)),
		"StatusCodeRange":       reflect.ValueOf((*ws.StatusCodeRange)(nil)),
		"StatusError":           reflect.ValueOf((*ws.StatusError)(nil)),
		"Upgrader":              reflect.ValueOf((*ws.Upgrader)(nil)),

		// interface wrapper definitions
		"_HandshakeHeader": reflect.ValueOf((*_github_com_gobwas_ws_HandshakeHeader)(nil)),
	}
}

// _github_com_gobwas_ws_HandshakeHeader is an interface wrapper for HandshakeHeader type
type _github_com_gobwas_ws_HandshakeHeader struct {
	WWriteTo func(w io.Writer) (n int64, err error)
}

func (W _github_com_gobwas_ws_HandshakeHeader) WriteTo(w io.Writer) (n int64, err error) {
	return W.WWriteTo(w)
}
