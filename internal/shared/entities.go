package shared

import "time"

type APODEvent struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Explanation string    `json:"explanation"`
	Title       string    `json:"title"`
	Picture     string    `json:"picture"`
}
