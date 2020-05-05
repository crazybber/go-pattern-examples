package memento

import (
	"fmt"
	"time"
)

////////////////////////////////
//使用游戏玩家的角色存档和读取的例子

//GamePlayer 是一个Originator 提供当前的游戏状态
type GamePlayer struct {
	hp, mp, role, level int //血量，魔法值，当前关卡
}

//RoleStatusMemento 一条备忘数据，存放瞬时状态的数据结构,一个数据结构
type RoleStatusMemento struct {
	tag           string //存档记录本身的名称，以便下次识别读取
	hp, mp, level int    //血量，魔法值，角色类型，当前关卡，
	timeMark      string //存档的可视化时间
}

//RoleStatusCaretaker 负责保存角色当前的状态数据,提供存取能力
//RoleStatusCaretaker 也是占内存/存储的地方，如果不停的读取，IO压力会变大的很大
type RoleStatusCaretaker struct {
	memens map[string]*RoleStatusMemento
}

//SaveStatus 保存当前角色的游戏状态
func (r *RoleStatusCaretaker) SaveStatus(item *RoleStatusMemento) {
	r.memens[item.tag] = item
	fmt.Printf("Game File %s  Saved at %s\n", item.tag, item.timeMark)
}

//RetriveStatus 提供需要的状态
func (r *RoleStatusCaretaker) RetriveStatus(savedTag string) *RoleStatusMemento {
	return r.memens[savedTag]

}

//Create 创建游戏的当前档案存档
func (g *GamePlayer) Create(tagName string) *RoleStatusMemento {

	return &RoleStatusMemento{
		tag:      tagName,
		hp:       g.hp,
		mp:       g.mp,
		level:    g.level,
		timeMark: time.Now().String(),
	}
}

//Load 载入存档，恢复数据
func (g *GamePlayer) Load(rm *RoleStatusMemento) {
	g.mp = rm.mp
	g.hp = rm.hp
	g.level = rm.level

	fmt.Printf("Game Profile had been restored to %s : %s\n", rm.tag, rm.timeMark)
}

//Status 玩家角色的当前状态
func (g *GamePlayer) Status() {
	fmt.Printf("Current Level :%d HP:%d, MP:%d\n", g.level, g.hp, g.mp)
}
