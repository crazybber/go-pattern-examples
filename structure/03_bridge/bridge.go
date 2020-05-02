package bridge

import "fmt"

//IMessage 发送消息接口
type IMessage interface {
	NoticeUser(text string)
	Priority() int
}

//ISMSMessage send SMS MSG
type ISMSMessage interface {
	//延迟接口的实现到其他类中
	NoticeUser(text string, noticeMessage IMessage)
}

//WSMessage MSG
type WSMessage struct {
	Handler ISMSMessage //持有进一步实现的引用关系
	Level   int
}

//NoticeUser by SMS
func (w *WSMessage) NoticeUser(text string) {
	//转递消息给其他对象，相当于承上启下
	fmt.Println("Websocket Notice User...", text)
	//转递消息给其他对象，相当于承上启下，并且需要把上下文变量传递下去
	if w.Handler != nil {
		w.Handler.NoticeUser(text, w)
	}

}

//Priority of SMS
func (w *WSMessage) Priority() int {
	return w.Level
}

//EmailMessage MSG
type EmailMessage struct {
	Handler ISMSMessage
	Level   int
}

//NoticeUser by SMS
func (e *EmailMessage) NoticeUser(text string) {
	//转递消息给其他对象，相当于承上启下，并且需要把上下文变量传递下去
	fmt.Println("Email Notice User...", text)
	if e.Handler != nil {
		e.Handler.NoticeUser(text, e)
	}

}

//Priority of SMS
func (e *EmailMessage) Priority() int {
	return e.Level
}

///需要实现具体的消息发送行为

//EmergencyWSMessage 紧急的短信消息
type EmergencyWSMessage struct {
}

//NoticeUser by email
func (e *EmergencyWSMessage) NoticeUser(text string, noticeMessage IMessage) {
	fmt.Println("Notice User", text, " By Websocket:", "with Level: ", noticeMessage.Priority())
}

//EmergencyEmailMessage 紧急的短信消息
type EmergencyEmailMessage struct {
}

//NoticeUser by email
func (e *EmergencyEmailMessage) NoticeUser(text string, noticeMessage IMessage) {
	fmt.Println("Notice User:", text, " By Email:", "with Level: ", noticeMessage.Priority())

}
