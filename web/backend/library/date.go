package library

import "time"

type Time time.Time
type Date time.Time
type DateTime time.Time

const (
	datetimeFormat = "2006-01-02 15:04:05"
	dateFormat     = "2006-01-02"
	timeFormat     = "15:04:05"
)

func (t *Date) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dateFormat+`"`, string(data), time.Local)
	*t = Date(now)
	return
}

func (t Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateFormat)
	b = append(b, '"')
	return b, nil
}

func (t Date) String() string {
	return time.Time(t).Format(dateFormat)
}

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+datetimeFormat+`"`, string(data), time.Local)
	*t = DateTime(now)
	return
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(datetimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, datetimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t DateTime) String() string {
	return time.Time(t).Format(datetimeFormat)
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}
