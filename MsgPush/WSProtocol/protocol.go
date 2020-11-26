package WSProtocol

import (
	"Lovers_srv/helper/Utils"
	"encoding/binary"
	"errors"
	"github.com/gorilla/websocket"
)

const (
	PACK_UNKNOWN   	= 0
	PACK_HEARTBEAT	= 1
	PACK_MESSAGE	= 2
	PACK_AUTH_FAIL	= 3
	PACK_AUTH_SUCC	= 4
	PACK_CONTINUE_ERR	= 5
)
const (
	// MaxBodySize max proto body size
	MaxBodySize = int32(1 << 12)
)

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _opSize + _seqSize
	_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_opOffset     = _opSize
	_seqOffset    = _opOffset + _seqSize
)

var (
	// ErrProtoPackLen proto packet len error
	ErrProtoPackLen = errors.New("default server codec pack length error")
	// ErrProtoHeaderLen proto header len error
	ErrProtoHeaderLen = errors.New("default server codec header length error")
)


type Proto struct {
	Op   int32  `json:"op"`
	Seq  int32  `json:"seq"`
	Body []byte `json:"body"`
}

// ReadWebsocket read a proto from websocket connection.
func (p *Proto) ReadWebsocket(ws *websocket.Conn) (err error) {
	var (
		bodyLen   int
		packLen   int32
		buf       []byte
	)
	if _, buf, err = ws.ReadMessage(); err != nil {
		return
	}
	packLen = int32(len(buf))
	if  packLen < (_opSize + _seqSize){
		return ErrProtoPackLen
	}
	p.Op = int32(binary.BigEndian.Uint32(buf[_packOffset:_opOffset]))
	p.Seq = int32(binary.BigEndian.Uint32(buf[_opOffset:_seqOffset]))
	if packLen < 0 || packLen > _maxPackSize {
		return ErrProtoPackLen
	}

	if bodyLen = int(packLen - (_opSize + _seqSize)); bodyLen > 0{
		p.Body = buf[_seqOffset:]
	}else{
		p.Body = nil
	}

	return
}

// WriteWebsocket write a proto to websocket connection.
func (p *Proto) WriteWebsocket(ws *websocket.Conn) (err error) {
	var buflen int
	if p.Body != nil{
		buflen = _opSize + _seqSize + len(p.Body)
	}else{
		buflen = _opSize + _seqSize
	}

	buf := make([]byte, buflen)
	binary.BigEndian.PutUint32(buf[_packOffset:], uint32(p.Op))
	binary.BigEndian.PutUint32(buf[_opOffset:],uint32(p.Seq))
	if p.Body != nil{
		copy(buf[_seqOffset : buflen], p.Body)
	}

	err = ws.WriteMessage(websocket.BinaryMessage, buf)
	if err != nil{
		Utils.ErrorOutputf("write web socket message failed:%s", err.Error())
	}
	return
}

// WriteWebsocketHeart write websocket heartbeat with room online.
func (p *Proto) WriteWebsocketHeart(wr *websocket.Conn) (err error) {

	buflen := _opSize + _seqSize
	p.Op = PACK_HEARTBEAT
	buf := make([]byte, buflen)
	binary.BigEndian.PutUint32(buf[_packOffset:], uint32(p.Op))
	binary.BigEndian.PutUint32(buf[_opOffset:],uint32(p.Seq))
	err = wr.WriteMessage(websocket.BinaryMessage, buf)
	if err != nil{
		Utils.ErrorOutputf("write web socket message failed:%s", err.Error())
	}
	return
}

func (p *Proto) WriteWSStatusWithOp(w *websocket.Conn, wsOp int32, body string)(err error){
	p.Op =wsOp
	p.Seq = 1
	if len(body) <= 0 || body == ""{
		p.Body = nil
	}else{
		p.Body = []byte(body)
	}
	err = p.WriteWebsocket(w)
	return
}