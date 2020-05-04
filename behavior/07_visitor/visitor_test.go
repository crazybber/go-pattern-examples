package visitor

import "testing"

func TestVisitor(t *testing.T) {

	//汽油提供给，制衣工厂
	g := gas{density: 100}

	//柴油，提供给军工厂
	d := diesel{energy: 897}

	//购买石油的客户
	m := &militaryFactory{}

	c := &clothFactory{}

	g.Accept(c)

	d.Accept(m)

}

func TestGameVisitorsList(t *testing.T) {

	retriveSetting := SettingVisitor{}
	attacker := Attacker{}

	pA := Player{"snow dance", 100} //角色名名：snow dance 100级
	pB := Player{"fire dragon", 120}
	npc := NPC{"groceries", true} //卖杂货的NPC，是能被打死的
	env := SystemEnv{"made by china", "v1.2.11"}

	//游戏对象
	gameObjects := []IGameContext{pA, npc, env, pB}

	for _, v := range gameObjects {
		v.Accept(retriveSetting)
	}

	t.Log("\n---- ！！！attack！！！- --")

	for _, v := range gameObjects {
		v.Accept(attacker)
	}

}
