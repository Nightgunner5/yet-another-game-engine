package iface

import "time"

type Thinker interface {
	// Perform an operation, scaled by delta. Return value
	// should be the minimum time until the next think.
	// Returning zero means this function should not be
	// called again until it is re-added to the scheduler.
	Think(delta time.Duration) (delay time.Duration)
}
