package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	seperator := placeholder{
		" ",
		"░",
		" ",
		"░",
		" ",
	}
	empty := placeholder{
		" ",
		" ",
		" ",
		" ",
		" ",
	}

	digits := [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}

	// for line := 0; line < 5; line++ {
	// 	for digit := range digits {
	// 		fmt.Print(digits[digit][line], "  ")
	// 	}
	// 	fmt.Println()
	// }
	screen.Clear()
	isSep := true
	colon := empty

	for {
		screen.MoveTopLeft()
		now := time.Now()

		hour := now.Hour()
		minute := now.Minute()
		second := now.Second()
		if isSep {
			colon = seperator
		} else {
			colon = empty
		}
		isSep = !isSep

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[minute/10], digits[minute%10],
			colon,
			digits[second/10], digits[second%10],
		}
		for line := 0; line < 5; line++ {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}

		time.Sleep(1 * time.Second)
	}

}
