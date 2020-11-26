package SportProtocol

import (
	"Lovers_srv/MsgPush/WSProtocol"
	"Lovers_srv/helper/Utils"
	"encoding/binary"
)

const(
	_packOffset = 0
	_typeOffset = 0
	_statusOffset = 4
	_bodyOffset = 8
	_maxPackLen = _bodyOffset + WSProtocol.MaxBodySize
)

//运动协议
type SportProto struct{
	SportType	int32	`json:"sport_type"` //运动类型，跑步、骑行、其他运动
	SportStatus int32	`json:sport_status` //运动状态，开始，运动中，暂停，结束
	Body		[]byte	`json:"body"`
}

func (sp *SportProto)ReadSportProto(buf []byte) (err error){
	var (
		bodyLen   int
		packLen   int32
	)

	packLen = int32(len(buf))
	if  packLen < _bodyOffset{
		return Utils.ErrorOutputf("[SportProto] read sport proto failed, pack len (%d) invalid",packLen)
	}
	sp.SportType = int32(binary.BigEndian.Uint32(buf[_typeOffset:_statusOffset]))
	sp.SportStatus = int32(binary.BigEndian.Uint32(buf[_statusOffset:_bodyOffset]))
	if packLen < 0 || packLen > _maxPackLen {
		return Utils.ErrorOutputf("[SportProto] read sport proto failed, pack len (%d) invalid",packLen)
	}

	if bodyLen = int(packLen - _bodyOffset); bodyLen > 0{
		sp.Body = buf[_bodyOffset:]
	}else{
		sp.Body = nil
	}

	return
}

func (sp *SportProto)PrepareSportProto()([]byte){
	buflen := _bodyOffset + len(sp.Body)
	buf := make([]byte, buflen)
	binary.BigEndian.PutUint32(buf[_packOffset:], uint32(sp.SportType))
	binary.BigEndian.PutUint32(buf[_statusOffset:_bodyOffset],uint32(sp.SportStatus))
	copy(buf[_bodyOffset : buflen], sp.Body)
	return buf
}
