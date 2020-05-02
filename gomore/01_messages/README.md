# 发布订阅模式

发布-订阅是一种消息传递模式，基本设计原则是将消息发布者和提供者分开，常常用于在不同组件之间传递消息。

![图示关系](../../images/pub-sub-pattern-0.png)
发布订阅模式往往实现为<消息代理，消息中间件，消息队列>等

该模式中的三个关键类型：消息本身、消息主题、订阅用户。
![图示关系](../../images/pub-sub-pattern-1.png)

图片来源：[pubsub-pattern.md](https://github.com/imsardine/dev-notes/blob/source/docs/pubsub-pattern.md)

现实生活中的各种信息平台，就是很好的发布订阅的例子，比如某八戒、某无忧等

这里演示，一个拼车例子，车主发布拼车(Topic)消息，消息推送到订阅拼车(Topic)信息的所有用户.

并模拟以下情形：

+ 车主(Topic)发布拼车消息
+ 拼车用户订阅拼车消息(Topic)
+ 拼车用户处理收到的拼车消息
