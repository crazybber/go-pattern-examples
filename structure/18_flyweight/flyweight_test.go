/*
 * @Description: github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-01 17:24:28
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-01 19:46:30
 */
package flyweight

import (
	"fmt"
	"testing"
)

func TestDeliveryPackets(t *testing.T) {

	dc := &DeliverCompany{Employees: make(map[string]IDeliver)}

	dc.Hire("bob")
	dc.Hire("lily")
	dc.Hire("bob")

	deliver1 := dc.GetDeliver("lily")

	deliver2 := dc.GetDeliver("lily")

	//一次送货任务
	dc.DeliverTask("lily", []string{"box1", "box2"})

	dc.DeliverTask("lily", []string{"box6", "box7", "box8"})

	deliver2.DeliverPackets([]string{"box9", "box10", "box11"})

	deliver1.DeliverPackets([]string{"box12"})

}

func TestDeliverEmployee(t *testing.T) {
	dc := &DeliverCompany{Employees: make(map[string]IDeliver)}

	dc.Hire("bob")
	dc.Hire("lily")

	deliver1 := dc.GetDeliver("lily")

	deliver2 := dc.GetDeliver("lily")

	dp1 := fmt.Sprintf("%p", deliver1)

	dp2 := fmt.Sprintf("%p", deliver2)

	if dp1 != dp2 {
		t.Error(dp1, "not the same with", dp2)
	}
}
