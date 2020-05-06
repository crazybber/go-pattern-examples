# Golang工程模式示例集合(Go Patterns Examples)

**包括了[go-patterns](https://github.com/tmrts/go-patterns) 和[golang-design-pattern](https://github.com/senghoo/golang-design-pattern)中的全部模式**

目前包括了**39种Go中常用的、面向工程化和最佳实践的模式/套路**，自然也包含常见的23种设计模式,重点是这里全部是例子、通俗易懂，甚至每个模式下的例子，改一下名字，稍微再增加几行代码就可以直接用在项目和工程中了。

每一种设计模式都有其特定的应用场景和要解决的问题，了解模式的关键点就在于弄清这些目标场景和问题，千万不要纠结于：为什么这个模式叫这个名字，这个模式为啥要这样用？

**这些模式不是你总结的，也不是我的总结的，如果是你的写的，你可以按照自己的喜欢的感觉给这些套路取名字，让别人去费劲想。**

## 了解一下姿势 Ways

+ 所谓模式就是套路,如功夫,招有定式
+ 学习模式，就是学习套路，弄清楚套路要解决的目标场景，这很重要.
+ 这里就是以实际代码示例展示设计模式,通俗易懂
+ 除了常见的23种普适的设计模式,Go也有一些属于自己的模式

## 行为型模式 Behavior Patterns

+ [x] [备忘录模式(Memento)](./behavior/09_memento)
+ [x] [中介者模式(Mediator)](./behavior/01_mediator)
+ [x] [闭包选项模式(Function Option)](./behavior/02_option)
+ [x] [观察者模式(Observer)](./behavior/03_observer)
+ [x] [命令模式(Command)](./behavior/11_command)
+ [x] [迭代器模式(Iterator)](./behavior/04_iterator)
+ [x] [模板方法模式(Template Method)](./behavior/05_template_method)
+ [x] [策略模式(Strategy)](./behavior/12_strategy)
+ [x] [状态模式(State)](./behavior/behavior16_state)
+ [x] [访问者模式(Visitor)](./behavior/07_visitor)
+ [x] [解释器模式(Interpreter)](./behavior/08_interpreter)
+ [x] [职责链模式(Chain of Responsibility)](./behavior/06_chain_of_responsibility)

## 创建型模式 Creation Patterns

+ [x] [New模式(New)](./creation/01_new)
+ [x] [简单工厂模式(Simple Factory)](./creation/02_simple_factory)
+ [x] [创建者模式(Builder)](./creation/03_builder)
+ [x] [单例模式(Singleton)](./creation/06_singleton)
+ [x] [对象池模式(Object Pool)](./creation/04_object_pool)
+ [x] [工厂方法模式(Factory Method)](./creation/05_factory_method)
+ [x] [抽象工厂模式(Abstract Factory)](./creation/08_abstract_factory)
+ [x] [原型模式(Prototype)](./creation/07_prototype)

## 结构型模式 Structure Patterns

+ [x] [外观模式(Facade)](./structure/01_facade)
+ [x] [适配器模式(Adapter)](./structure/02_adapter)
+ [x] [桥模式(Bridge)](./structure/03_bridge)
+ [x] [复合模式(Composite)](./structure/05_composite)
+ [x] [享元模式(Flyweight)](./structure/04_flyweight)
+ [x] [装饰器模式(Decorator)](./structure/06_decorator)
+ [x] [代理模式(Proxy)](./structure/07_proxy)

## 更多模式 Go More Patterns

+ [x] [发布订阅模式(Pub-Sub)](./gomore/01_messages)
+ [x] [时差模式(Time Profile)](./gomore/02_profiles)
+ [x] [上下文模式(Context)](./gomore/03_context)
+ [ ] [WIP][淡入模式(Fan-In)](./gomore/04_fan_in)
+ [ ] [WIP][淡出模式(Fan-Out)](./gomore/05_fan_out)
+ [ ] [WIP][熔断模式(circuit breaker)](./gomore/06_circuit_breaker)
+ [ ] [WIP][限流模式(rate limiting))](./gomore/07_rate_limiting)
+ [ ] [WIP][信号量模式(Semaphore)](./gomore/08_semaphore)
+ [ ] [WIP][并行模式(Parallelism)](./gomore/09_parallelism)
+ [ ] [WIP][生成器模式(Generators)](./gomore/10_generators)
+ [ ] [WIP][屏障模式(N-Barrier)](./gomore/11_n_barrier)
+ [ ] [WIP][有限并行模式(Bounded Parallelism)](./gomore/12_bounded_parallelism)

## 参考资料(Design patters Articles)

| Patterns | Design Patterns | Status |
|:-------:|:----------- |:------:|
| [go-patterns](https://github.com/crazybber/go-patterns) | [design-pattern-tutorial](https://www.runoob.com/design-pattern/design-pattern-tutorial.html) |p|
| [design_pattern](http://c.biancheng.net/design_pattern)|[golang-design-pattern](https://github.com/senghoo/golang-design-pattern) |p|
[go-resiliency](https://github.com/eapache/go-resiliency) | [Behavioral](https://github.com/AlexanderGrom/go-patterns/tree/master/Behavioral)|v|
| [go-patterns](https://github.com/sevenelevenlee/go-patterns) | [go_design_pattern](https://github.com/monochromegane/go_design_pattern)|p|

## 更多 More

需要重新温习下Go基础?看这里

[go-fucking-exercises](https://github.com/crazybber/go-fucking-exercise)
