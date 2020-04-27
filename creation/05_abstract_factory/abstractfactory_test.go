package abstractfactory

import "testing"

func TestSQLFactory(t *testing.T) {

	factory := &SQLFactory{}
	orderWorker := factory.CreateOrderWorker()
	orderWorker.SaveOrder()
	detailWorker := factory.CreateOrderDetailWorker()
	detailWorker.SaveOrderDetail()
}

func TestNoSqlFactory(t *testing.T) {

	factory := &NoSQLFactory{}
	orderWorker := factory.CreateOrderWorker()
	orderWorker.SaveOrder()
	detailWorker := factory.CreateOrderDetailWorker()
	detailWorker.SaveOrderDetail()
}
