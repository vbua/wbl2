package event

import (
	"time"
)

//type eventDate time.Time

type Event struct {
	ID    uint32    `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

//func (e *eventDate) UnmarshalJSON(b []byte) error {
//	s := strings.Trim(string(b), `"`)
//	newTime, err := time.ParseInLocation("2006-01-02", string(s), time.Local)
//	if err != nil {
//		return err
//	}
//	*e = eventDate(newTime)
//	return nil
//}
//
//func (e *eventDate) MarshalJSON() ([]byte, error) {
//	stamp := fmt.Sprintf("\"%s\"", time.Time(*e).Format("2006-01-02"))
//	return []byte(stamp), nil
//}
//
//func (e *eventDate) String() string {
//	return time.Time(*e).Format("2006-01-02")
//}
