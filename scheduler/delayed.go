package scheduler

import (
	"github.com/Nightgunner5/yet-another-game-engine/iface"
	"time"
)

type delayedThink struct {
	thinker iface.Thinker
	delay   time.Duration
}

func (t *delayedThink) Think(delta time.Duration) (delay time.Duration) {
	if t.delay == 0 {
		return t.thinker.Think(delta)
	}

	delay = t.delay
	t.delay = 0
	return
}
