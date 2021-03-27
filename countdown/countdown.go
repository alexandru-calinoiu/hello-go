package countdown

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i -= 1 {
		s.Sleep()
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprint(w, finalWord)
}
