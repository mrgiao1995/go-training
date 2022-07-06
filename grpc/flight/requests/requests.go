package requests

import "time"

type SearchFlightRequest struct {
	From string
	To   string
	Date time.Time
}
