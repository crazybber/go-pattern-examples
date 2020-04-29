# 发布订阅模式

发布-订阅是一种消息传递模式，基本设计原则是将消息发布者和提供者分开，常常用于在不同组件之间传递消息。

![图示关系](../../images/pub-sub-pattern-0.png)
发布订阅模式往往实现为<消息代理，消息中间件，消息队列>等

该模式中的三个关键类型：消息本身、消息主题、订阅用户。
![图示关系](../../images/pub-sub-pattern-1.png)

图片来源：[pubsub-pattern.md](https://github.com/imsardine/dev-notes/blob/source/docs/pubsub-pattern.md)

示例演示消息队的某个主题(Topic)收到用户订阅之后，的处理过程和与用户之间的响应互动。

并模拟以下情形：

+ 主题向用户发送消息
+ 用户订阅主题
+ 用户发送某个主题消息
