package bridge

import "testing"

func TestSendMessage(t *testing.T) {

	//注意看这里,接口的实例关系在初始化时候是固定的
	ws := &WSMessage{&EmergencyWSMessage{}, 100}

	//注意看这里,接口的实例关系在初始化时候是固定的
	email := &EmailMessage{&EmergencyEmailMessage{}, 10}

	ws.NoticeUser("Miss White ,Let's Drink")
	email.NoticeUser("Miss White,Fire!,Fire!")

	ews := &EmergencyWSMessage{}
	eem := &EmergencyEmailMessage{}
	list := []IMessage{
		&WSMessage{ews, 50},
		&WSMessage{ews, 100},
		&EmailMessage{eem, 10},
		&EmailMessage{eem, 20},
	}
	for _, v := range list {
		v.NoticeUser("Let’s go for fun")
	}
}
