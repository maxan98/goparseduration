package goparseduration

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {

	some, _:=ParseDuration(`2w3d4m14h`)
	fmt.Printf("%dw\n%dd\n%dh\n%dm\n%ds\n",some.weeks,some.days,some.hours,some.minutes,some.seconds)
}
