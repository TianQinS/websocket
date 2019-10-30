// Code generated by automatic for 'time'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlibs

import (
	"reflect"
	"time"
)

func init() {
	Symbols["time"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ANSIC":                  reflect.ValueOf(time.ANSIC),
		"After":                  reflect.ValueOf(time.After),
		"AfterFunc":              reflect.ValueOf(time.AfterFunc),
		"April":                  reflect.ValueOf(time.April),
		"August":                 reflect.ValueOf(time.August),
		"Date":                   reflect.ValueOf(time.Date),
		"December":               reflect.ValueOf(time.December),
		"February":               reflect.ValueOf(time.February),
		"FixedZone":              reflect.ValueOf(time.FixedZone),
		"Friday":                 reflect.ValueOf(time.Friday),
		"Hour":                   reflect.ValueOf(int64(time.Hour)),
		"January":                reflect.ValueOf(time.January),
		"July":                   reflect.ValueOf(time.July),
		"June":                   reflect.ValueOf(time.June),
		"Kitchen":                reflect.ValueOf(time.Kitchen),
		"LoadLocation":           reflect.ValueOf(time.LoadLocation),
		"LoadLocationFromTZData": reflect.ValueOf(time.LoadLocationFromTZData),
		"Local":                  reflect.ValueOf(&time.Local).Elem(),
		"March":                  reflect.ValueOf(time.March),
		"May":                    reflect.ValueOf(time.May),
		"Microsecond":            reflect.ValueOf(time.Microsecond),
		"Millisecond":            reflect.ValueOf(time.Millisecond),
		"Minute":                 reflect.ValueOf(int64(time.Minute)),
		"Monday":                 reflect.ValueOf(time.Monday),
		"Nanosecond":             reflect.ValueOf(time.Nanosecond),
		"NewTicker":              reflect.ValueOf(time.NewTicker),
		"NewTimer":               reflect.ValueOf(time.NewTimer),
		"November":               reflect.ValueOf(time.November),
		"Now":                    reflect.ValueOf(time.Now),
		"October":                reflect.ValueOf(time.October),
		"Parse":                  reflect.ValueOf(time.Parse),
		"ParseDuration":          reflect.ValueOf(time.ParseDuration),
		"ParseInLocation":        reflect.ValueOf(time.ParseInLocation),
		"RFC1123":                reflect.ValueOf(time.RFC1123),
		"RFC1123Z":               reflect.ValueOf(time.RFC1123Z),
		"RFC3339":                reflect.ValueOf(time.RFC3339),
		"RFC3339Nano":            reflect.ValueOf(time.RFC3339Nano),
		"RFC822":                 reflect.ValueOf(time.RFC822),
		"RFC822Z":                reflect.ValueOf(time.RFC822Z),
		"RFC850":                 reflect.ValueOf(time.RFC850),
		"RubyDate":               reflect.ValueOf(time.RubyDate),
		"Saturday":               reflect.ValueOf(time.Saturday),
		"Second":                 reflect.ValueOf(time.Second),
		"September":              reflect.ValueOf(time.September),
		"Since":                  reflect.ValueOf(time.Since),
		"Sleep":                  reflect.ValueOf(time.Sleep),
		"Stamp":                  reflect.ValueOf(time.Stamp),
		"StampMicro":             reflect.ValueOf(time.StampMicro),
		"StampMilli":             reflect.ValueOf(time.StampMilli),
		"StampNano":              reflect.ValueOf(time.StampNano),
		"Sunday":                 reflect.ValueOf(time.Sunday),
		"Thursday":               reflect.ValueOf(time.Thursday),
		"Tick":                   reflect.ValueOf(time.Tick),
		"Tuesday":                reflect.ValueOf(time.Tuesday),
		"UTC":                    reflect.ValueOf(&time.UTC).Elem(),
		"Unix":                   reflect.ValueOf(time.Unix),
		"UnixDate":               reflect.ValueOf(time.UnixDate),
		"Until":                  reflect.ValueOf(time.Until),
		"Wednesday":              reflect.ValueOf(time.Wednesday),

		// type definitions
		"Duration":   reflect.ValueOf((*time.Duration)(nil)),
		"Location":   reflect.ValueOf((*time.Location)(nil)),
		"Month":      reflect.ValueOf((*time.Month)(nil)),
		"ParseError": reflect.ValueOf((*time.ParseError)(nil)),
		"Ticker":     reflect.ValueOf((*time.Ticker)(nil)),
		"Time":       reflect.ValueOf((*time.Time)(nil)),
		"Timer":      reflect.ValueOf((*time.Timer)(nil)),
		"Weekday":    reflect.ValueOf((*time.Weekday)(nil)),
	}
}
