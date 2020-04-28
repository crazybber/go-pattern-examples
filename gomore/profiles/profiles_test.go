package profile

import (
	"log"
	"math/big"
	"time"
	"testing"
)


func TestBigintProfile(t *testing.T) {


	bigint :=	big.NewInt(14468)

	BigIntFactorial(bigint) 

	bigint = big.NewInt(24566)

	BigIntFactorial(bigint) 

}


//Duration for time differ
func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)

	log.Printf("%s lasted %s", name, elapsed)
}

//BigIntFactorial ...
func BigIntFactorial(input *big.Int) *big.Int {

	//关键点是这一句.
	defer Duration(time.Now(), "IntFactorial")

	x := input
	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
		y.Mul(y, x)
	}

	return x.Set(y)
}
