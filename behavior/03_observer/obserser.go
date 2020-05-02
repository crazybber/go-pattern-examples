package observer

import (
	"context"
	"fmt"
)

//使用RSS邮件订阅的例子
//某科技论坛更新时候，通过邮件通知订阅该版块的用户关注更新.

type updates struct {
	topic string
	order int
}

//TechBBS 科技论坛，是发送通知的主体
type TechBBS struct {
	mailObservers []IMailReceiver ///邮件订阅该板块动态的用户
	context       context.Context ///更新的上下消息，其实就是一堆参数了
}

//IMailReceiver 邮件接收者就是订阅更新的用户
type IMailReceiver interface {
	Notice(context.Context)
}

//Registry 提供给用户通过邮件注册获取/订阅更新动态的能力
func (t *TechBBS) Registry(receiver IMailReceiver) {
	t.mailObservers = append(t.mailObservers, receiver)
}

//SetConext 设置更新内容的上下文
func (t *TechBBS) SetConext(ctx context.Context) {
	t.context = ctx
}

//notifyUpdate 通知订阅用户，我更新了，你们可以来看了
func (t *TechBBS) noticeAllUpdate() {
	for _, m := range t.mailObservers {
		//逐个通知
		m.Notice(t.context)
	}
}

//User 用户
type User struct {
	name string
}

//Notice 用户收到订阅通知
func (u *User) Notice(ctx context.Context) {

	content := ctx.Value(updates{}).(updates)

	fmt.Printf("%s receive updates notice\n", u.name)

	fmt.Printf("updates order: %d content：%s\n", content.order, content.topic)
}
