package state

import "testing"

func TestState(t *testing.T) {

	expect := "vibrating humming ... vibrating humming...vibrating humming..." +
		"vibrating humming ... vibrating humming...vibrating humming..." +
		"sun rise ,get up ,get up get up..."

	mobile := NewMobileAlert()

	result := mobile.Alert()
	result += mobile.Alert()

	mobile.SetState(&MobileAlertSong{})

	result += mobile.Alert()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
func TestWeeks(t *testing.T) {
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}

	for i := 0; i < 8; i++ {
		todayAndNext()
	}
	// Output:
	// Sunday
	// Monday
	// Tuesday
	// Wednesday
	// Thursday
	// Friday
	// Saturday
	// Sunday
}
