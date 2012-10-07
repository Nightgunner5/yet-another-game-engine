package scheduler

import (
	"github.com/Nightgunner5/yet-another-game-engine/iface"
	"time"
)

// Schedule a Thinker to Think as soon as possible.
func ScheduleThink(thinker iface.Thinker) {
	go think(thinker)
}

// Schedule a Thinker to Think after waiting at least delay nanoseconds.
func ScheduleThinkDelayed(thinker iface.Thinker, delay time.Duration) {
	go think(&delayedThink{thinker, delay})
}
