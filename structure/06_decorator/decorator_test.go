package decorator

import (
	"fmt"
	"testing"
)

func TestDecorateNatureGirl(t *testing.T) {

	//准备天然的妹子一枚
	origin := &NatureGirl{faceValue: 6}

	fmt.Println("face looks ", origin.FaceLooks())

	//只需要略施粉黛
	makeupGirl := NewGirlWithMakeup(origin, 2)

	fmt.Println("after makeup face looks ", makeupGirl.FaceLooks())

	fmt.Println("girl's real face ", makeupGirl.(*GirlWithMakeups).FaceReal())

}
