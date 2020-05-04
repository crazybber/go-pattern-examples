package visitor

import "fmt"

//IGameObject 游戏上下文对象
type IGameObject interface {
	Accept(IGameVisitor)
}

//Player 代表其他玩家，因为用户角色需要和其他玩家交互
type Player struct {
	Name  string
	Level int
}

//Accept 提供交互接口
func (p Player) Accept(v IGameVisitor) {
	v.VisitPlayer(p)
}

//NPC 被方法对象
type NPC struct {
	Name       string
	IsImmortal bool //是否可以被打死
}

//Accept 接受聚能NPC访问能力的对象的访问
func (n NPC) Accept(v IGameVisitor) {
	v.VisitNPC(n)
}

//SystemEnv 环境对象
type SystemEnv struct {
	Mark    string //环境标识
	Version string //环境版本
}

//Accept 提供对环境的访问
func (s SystemEnv) Accept(v IGameVisitor) {
	v.VisitSystemEnv(s)
}

//IGameVisitor 游戏提供的环境访问能力
type IGameVisitor interface {
	VisitPlayer(Player)
	VisitNPC(NPC)
	VisitSystemEnv(SystemEnv)
}

// SettingVisitor 只提供Setting的能力
type SettingVisitor struct{}

//VisitPlayer 提供交互的第三方对象的信息
func (SettingVisitor) VisitPlayer(p Player) {
	fmt.Printf("Game Player: Name:%s ,Level:%d\n", p.Name, p.Level)
}

//VisitNPC 提供NPC的信息
func (SettingVisitor) VisitNPC(n NPC) {
	fmt.Printf("Game NPC: Name:%s ,Immortal:%v\n", n.Name, n.IsImmortal)
}

//VisitSystemEnv 提供游戏环境信息
func (SettingVisitor) VisitSystemEnv(s SystemEnv) {
	fmt.Printf("Game Env: Mark:%s ,Version:%s\n", s.Mark, s.Version)
}

// Attacker 攻击者
type Attacker struct{ name string }

//VisitPlayer 攻击其他玩家
func (a Attacker) VisitPlayer(p Player) {
	fmt.Printf("%s Attack Player : %s\n", a.name, p.Name)
}

//VisitNPC 攻击NPC
func (a Attacker) VisitNPC(n NPC) {
	fmt.Printf("%s Attack NPC: %s\n", a.name, n.Name)
}

//VisitSystemEnv 攻击环境，如石头，大门，墙壁
func (a Attacker) VisitSystemEnv(s SystemEnv) {
	fmt.Printf("Unsupported target %s\n", "game env")
}
