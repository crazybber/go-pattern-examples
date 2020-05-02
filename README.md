# Go语言设计模式示例集合(Go Patterns Examples)

Go常用的、面向工程化和最佳实践的模式套路，包含常见的23种设计模式,重点是这里全部是例子、通俗易懂，每个模式，改一下名字就可以直接用在项目和工程中。

每一种设计模式都有其特定的应用场景和要解决的问题，了解模式的关键点就在于弄清这些目标场景和问题，千万不要纠结于：为什么这个模式叫这个名字，这个模式为啥要这样用？

记住：这些模式不是你总结的，也不是我的总结的，如果是你的写的，你可以按照自己的喜欢的感觉给这些套路取名字，让别人去费劲想。

## 姿势

+ 所谓模式就是套路,如功夫,招有定式
+ 学习模式，就是学习套路，弄清楚套路要解决的目标场景，这很重要.
+ 这里就是以实际代码示例展示设计模式,通俗易懂
+ 除了常见的23种普适的设计模式,Go也有一些属于自己的模式

## 创建型模式

+ [x] [简单工厂模式(Simple Factory)](./creation/00_simple_factory)
+ [x] [工厂方法模式(Factory Method)](./creation/04_factory_method)
+ [x] [抽象工厂模式(Abstract Factory)](./creation/05_abstract_factory)
+ [x] [创建者模式(Builder)](./creation/06_builder)
+ [x] [原型模式(Prototype)](./creation/07_prototype)
+ [x] [单例模式(Singleton)](./creation/03_singleton)
+ [ ] [对象池模式(Object Pool)](./creation/24_object_pool)
+ [x] [New模式(New)](./creation/25_new)

## 结构型模式

+ [x] [外观模式(Facade)](./structure/01_facade)
+ [x] [适配器模式(Adapter)](./structure/02_adapter)
+ [x] [代理模式(Proxy)](./structure/09_proxy)
+ [ ] [复合模式(Composite)](./structure/13_composite)
+ [x] [享元模式(Flyweight)](./structure/18_flyweight)
+ [ ] [装饰模式(Decorator)](./structure/20_decorator)
+ [x] [桥模式(Bridge)](./structure/22_bridge)

## 行为型模式

+ [x] [中介者模式(Mediator)](./behavior/08_mediator)
+ [ ] [观察者模式(Observer)](./behavior/10_observer)
+ [ ] [命令模式(Command)](./behavior/11_command)
+ [ ] [迭代器模式(Iterator)](./behavior/12_iterator)
+ [ ] [模板方法模式(Template Method)](./behavior/14_template_method)
+ [x] [策略模式(Strategy)](./behavior/15_strategy)
+ [ ] [状态模式(State)](./behavior/behavior16_state)
+ [ ] [备忘录模式(Memento)](./behavior/17_memento)
+ [ ] [解释器模式(Interpreter)](./behavior/19_interpreter)
+ [ ] [职责链模式(Chain of Responsibility)](./behavior/21_chain_of_responsibility)
+ [ ] [访问者模式(Visitor)](./behavior/23_visitor)
+ [x] [闭包选项模式(Function Option)](./behavior/26_option)

## Go More

+ [x] [发布订阅模式(Pub-Sub)](./gomore/27_messages)
+ [x] [时差模式(Time Profile)](./gomore/28_profiles)
+ [x] [上下文模式(Context)](./gomore/29_context)

## 参考资料(Design patters Articles)

[GO模式文档](https://github.com/nynicg/go-patterns)

[菜鸟教程—设计模式](https://www.runoob.com/design-pattern/design-pattern-tutorial.html)

[23-Pattern-in-Go](https://github.com/senghoo/golang-design-pattern)


## 更多

需要重新温习下Go基础?看这里

[go-exercise](https://github.com/crazybber/go-exercise)
