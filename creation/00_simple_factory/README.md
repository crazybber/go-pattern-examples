# 简单工厂模式

go 可以使用自定义的New()来初始化相关类。
New 函数返回接口时就是简单工厂模式，也是golang一般推荐做法。

simplefactory包中只有Mouth接口和New()函数可见。
