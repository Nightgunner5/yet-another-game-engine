package scheduler

import (
	"testing"
	"time"
)

type thinkOnce chan time.Duration

func (t thinkOnce) Think(delta time.Duration) (delay time.Duration) {
	t <- delta

	close(t)

	return 0
}

func TestImmediateThink(t *testing.T) {
	thinker := make(thinkOnce)
	ScheduleThink(thinker)

	time.Sleep(5 * time.Millisecond)
	select {
	case d := <-thinker:
		if d > 5*time.Millisecond {
			t.Errorf("Delta was longer than 5ms (%s)", d)
		}
	default:
		t.Error("Thinker did not think within 5ms, but was supposed to think immediately")
	}
}

func TestDelayedThink(t *testing.T) {
	thinker := make(thinkOnce)
	ScheduleThinkDelayed(thinker, 10*time.Millisecond)

	time.Sleep(5 * time.Millisecond)
	select {
	case d := <-thinker:
		t.Errorf("Think should not have happened yet (delta %s)", d)
	default:
		// Expected
	}

	time.Sleep(10 * time.Millisecond)
	select {
	case d := <-thinker:
		if d < 10*time.Millisecond {
			t.Errorf("Think was only delayed by %s, not 10ms as requested", d)
		}
	default:
		t.Error("Think scheduled for 10ms has not happened after 15ms")
	}
}
