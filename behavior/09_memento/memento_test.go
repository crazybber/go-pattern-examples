package memento

import (
	"testing"
	"time"
)

func TestGameArchive(t *testing.T) {

	gamerole := GamePlayer{hp: 1000, mp: 232, level: 20}

	datakeeper := RoleStatusCaretaker{memens: make(map[string]*RoleStatusMemento)}

	archive1 := gamerole.Create("第一次存档")

	//交给管数据的人，存起来
	datakeeper.SaveStatus(archive1)

	//模拟,随机玩会儿游戏
	time.Sleep(time.Millisecond * 1132)

	//更新角色当前状态
	gamerole = GamePlayer{hp: 500, mp: 10, level: 30}

	//看一下状态
	gamerole.Status()

	archive2 := gamerole.Create("第二次存档")

	//交给管数据的人，存起来
	datakeeper.SaveStatus(archive2)

	//准备恢复第一次的存档

	//查找档案
	restore1 := datakeeper.RetriveStatus("第一次存档")

	//载入档案
	gamerole.Load(restore1)

	//看一下状态
	gamerole.Status()

}
