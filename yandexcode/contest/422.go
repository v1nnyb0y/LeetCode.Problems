package contest

import (
	"fmt"
	"log"
	"math"
)

type TwoTeams struct {
	A, B, N int
}

func (t *TwoTeams) inputFunc() {
	inputNumber := func() int {
		var number int
		_, err := fmt.Scan(&number)
		if err != nil {
			log.Fatal(err)
		}
		return number
	}
	t.A = inputNumber()
	t.B = inputNumber()
	t.N = inputNumber()
}

func (t *TwoTeams) outputFunc() {
	maxTeamA := t.A
	var minTeamB int
	if t.B%t.N == 0 {
		minTeamB = t.B / t.N
	} else {
		minTeamB = int(math.Floor(float64(t.B/t.N))) + 1
	}
	if maxTeamA > minTeamB {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func TwoTeamsProcessing() {
	solution := &TwoTeams{}
	solution.inputFunc()
	solution.outputFunc()
}
