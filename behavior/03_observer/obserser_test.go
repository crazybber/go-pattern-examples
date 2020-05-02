package observer

import (
	"context"
	"testing"
)

func TestObserver(t *testing.T) {

	//内容提供商，科技BBS
	techInfoProvider := TechBBS{}

	lily := User{"Lily"}

	jacky := User{"Jacky"}

	techInfoProvider.Registry(&lily)
	techInfoProvider.Registry(&jacky)

	updateKey := updates{}

	updateValue := updates{topic: "cosmos", order: 1001}

	updateContent := context.WithValue(context.Background(), updateKey, updateValue)
	techInfoProvider.SetConext(updateContent)

	techInfoProvider.noticeAllUpdate()

}
