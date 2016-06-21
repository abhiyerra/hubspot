package hubspot

import (
	"time"
)

// Timestamp is used
// http://developers.hubspot.com/docs/faq/how-should-timestamps-be-formatted-for-hubspots-apis
func Timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
