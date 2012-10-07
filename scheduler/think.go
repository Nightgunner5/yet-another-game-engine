package scheduler

import (
	"github.com/Nightgunner5/yet-another-game-engine/iface"
	"time"
)

func think(thinker iface.Thinker) {
	var delay time.Duration
	last := time.Now()
	for {
		current := time.Now()
		delay = thinker.Think(current.Sub(last))
		if delay == 0 {
			return
		}
		last = current

		time.Sleep(delay)
	}
}
