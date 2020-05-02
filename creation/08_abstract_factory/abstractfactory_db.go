package abstractfactory

import "fmt"

///这里一个分库分表的场景
///主订单，放在常规MySQL数据库
//订单详情，单条数据量过多，所以放在NewSQL数据库

//Order : 订单主记录
type Order interface {
	SaveOrder()
	//.....BALABALA..可以继续写很多接口方法
	//每个接口方法做一件事
}

//OrderDetail 订单详情
type OrderDetail interface {
	SaveOrderDetail()
	//.....BALABALA..可以继续写很多接口方法
	//每个接口方法做一件事
}

//IRepository 是当抽象工厂模式例子中的关键接口
//IRepository 返回一组数据处理对象,用于处理不同的数据类型
//IRepository 本质是创建工作对象，但必须以接口方式返回
type IRepository interface {
	CreateOrderWorker() Order
	CreateOrderDetailWorker() OrderDetail
	//.....BALABALA..可以继续写很多接口方法
	//每个接口方法都要返回一个接口
}

////////////////////////////////
//接口定义好了，开始进行实现和应用
////////////////////////////////

//MySQLOrderWorker 处理主订单的worker
type MySQLOrderWorker struct{}

//SaveOrder 实现SaveOrder接口,保存订单
func (*MySQLOrderWorker) SaveOrder() {
	fmt.Print("MySQL save main order\n")
}

//NewSQLOrderWorker 处理订单详情的worker
type NewSQLOrderWorker struct{}

// SaveOrderDetail SaveOrderDetail接口,保存订单细节
func (*NewSQLOrderWorker) SaveOrderDetail() {
	fmt.Print("NewSQL save OrderDetail\n")
}

//SQLFactory 第一个工厂
//SQLFactory SQL关系型数据库工厂实现DAOFactory接口
type SQLFactory struct{}

//CreateOrderWorker 创建实现了能够保存主订单记录的工作对象
func (*SQLFactory) CreateOrderWorker() Order {
	return &MySQLOrderWorker{}
}

//CreateOrderDetailWorker  创建实现了能够保存订单详情的工作对象
func (*SQLFactory) CreateOrderDetailWorker() OrderDetail {
	return &NewSQLOrderWorker{}
}

///假设说：现在CTO说要迁移到非关系数据库,我们进行下一个实现
///主订单，放在常规MySQL数据库
//订单详情，单条数据量过多，所以放在NewSQL数据库

//MongoDBWorker 处理主订单
//MongoDBWorker No-SQL非关系型数据库工厂实现DAOFactory接口
type MongoDBWorker struct{}

//SaveOrder by MongoDBWorker ...
func (*MongoDBWorker) SaveOrder() {
	fmt.Print("MongoDB save main order\n")
}

//PouchDBWorker 处理订单细节
type PouchDBWorker struct{}

// SaveOrderDetail ...
func (*PouchDBWorker) SaveOrderDetail() {
	fmt.Print("PouchDBWorker save OrderDetail\n")
}

//NoSQLFactory 第二个工厂，用于处理非关系型存储
//NoSQLFactory No-SQL非关系型数据库工厂实现DAOFactory接口
type NoSQLFactory struct{}

//CreateOrderWorker 创建主订单工作对象
func (*NoSQLFactory) CreateOrderWorker() Order {
	return &MongoDBWorker{}
}

//CreateOrderDetailWorker 创建订单详情工作对象
func (*NoSQLFactory) CreateOrderDetailWorker() OrderDetail {
	return &PouchDBWorker{}
}
