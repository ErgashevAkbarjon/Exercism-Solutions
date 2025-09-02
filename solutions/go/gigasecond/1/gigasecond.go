// It calculates what date it will be if we
// add a thousand million seconds to it
package gigasecond

import "time"

// Adding 1 000 000 000 seconds to t
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
