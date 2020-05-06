// Package state is an example of the State Pattern.
package state

////////////////////////////////
//闹装有两种状态，震铃状态，非震铃状态

// AlertStater provides a common interface for various states.
type AlertStater interface {
	Alert() string
}

// Alert implements an alert depending on its state.
type Alert struct {
	state AlertStater
}

// Alert returns a alert string 代表振铃
func (a *Alert) Alert() string {
	return a.state.Alert()
}

// SetState changes state
func (a *Alert) SetState(state AlertStater) {
	a.state = state
}

// NewAlert is the Alert constructor，默认振铃是震动
func NewAlert() *Alert {
	return &Alert{state: &AlertVibration{}}
}

// AlertVibration implements vibration alert
type AlertVibration struct {
}

// Alert returns a alert string ，默认振铃
func (a *AlertVibration) Alert() string {
	return "vibrating humming ... vibrating humming...vibrating humming..."
}

// AlertSong implements beep alert
type AlertSong struct {
}

// Alert returns a new alert string 歌曲振铃，设置这个状态模式，闹钟只会唱歌
func (a *AlertSong) Alert() string {
	return "sun rise ,get up ,get up get up..."
}
