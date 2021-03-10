package plugins

import (
	"github.com/Mrs4s/MiraiGo/client/pb/msg"
	"github.com/Mrs4s/MiraiGo/message"
)

// PluginInfo 插件的信息
type PluginInfo struct {
	Name        string
	Description string
}

type messageType string

const (
	//GroupMessage 群消息
	GroupMessage = messageType("group")
	//PrivateMessage 私聊消息
	PrivateMessage = messageType("private")
)

//MessageRequest 消息请求
type MessageRequest struct {
	MessageType    messageType
	ID             int32
	InternalID     int32
	GroupCode      int64
	GroupName      string
	Sender         *message.Sender
	Time           int32
	Elements       []message.IMessageElement
	OriginalObject *msg.Message
}

// MessageResponse 消息事件返回对象
type MessageResponse struct {
	Elements []message.IMessageElement
}

//Plugin 插件
type Plugin interface {
	//PluginInfo 获取插件的信息
	PluginInfo() *PluginInfo
	//PluginInit 插件初始化
	PluginInit()
	//OnMessage 监听消息
	OnMessage(msg *MessageRequest) *MessageResponse
}