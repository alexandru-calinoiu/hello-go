package countdown

import (
	"reflect"
	"testing"
)

type CountdownOperationSpy struct {
	calls []string
}

const sleep = "sleep"
const write = "write"

func (s *CountdownOperationSpy) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
}

func TestCountdown(t *testing.T) {
	spy := CountdownOperationSpy{}

	Countdown(&spy, &spy)

	want := []string{
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want, spy.calls) {
		t.Errorf("wanted calls %v got %v", want, spy.calls)
	}
}
