package option

import (
	"testing"
)

func TestFileFunctionOptions(t *testing.T) {

	Introduce("tom", Gender(true), Company("land company"))

	Introduce("lily", Company("sky commnay"), UID(123))

	Introduce("admin", Company("risky commnay"), UID(883))

}
