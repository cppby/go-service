package space

import (
	"errors"

	"github.com/giant-tech/go-service/framework/idata"
	"github.com/giant-tech/go-service/framework/msgdef"
)

type iPropsSender interface {
	SendFullProps() error
}

var errPropsNil = errors.New("Props num is zero")

// SendFullProps 发送完整AOI属性信息
func (e *Entity) SendFullProps() error {
	num, data := e.GetAOIProp()
	if num == 0 {
		return errPropsNil
	}

	msg := &msgdef.PropsSyncClient{
		EntityID: e.GetEntityID(),
		Num:      uint32(num),
		Data:     data,
	}
	e.CastMsgToMe(msg)
	return nil
}

// GetAOIProp 获得进入其它人AOI范围内需要收到的属性数据
func (e *Entity) GetAOIProp() (uint32, []byte) {
	return e.PackProps(uint32(idata.ServiceClient))
}

// GetBaseProps 获取基础属性
func (e *Entity) GetBaseProps() []byte {
	// msg := e.genBasePropsMsg()
	// data := make([]byte, msg.Size())
	// msg.MarshalTo(data)
	// return data
	return nil
}

func (e *Entity) genBasePropsMsg() *msgdef.EntityBaseProps {
	msg := &msgdef.EntityBaseProps{}

	msg.EntityID = e.GetEntityID()

	return msg
}

// setBasePropsDirty 设置基础属性变化
func (e *Entity) setBasePropsDirty(dirty bool) {
	e.basePropsDirty = dirty
}

// FlushBaseProps 刷新基础属性
func (e *Entity) FlushBaseProps() {
	if e.basePropsDirty {
		e.CastMsgToAllClient(e.genBasePropsMsg())
	}
}
