//go:build windows && amd64

package w32

import (
	"errors"
	"syscall"
	"time"
)

// GetVariantDate converts COM Variant Time value to Go time.Time.
func GetVariantDate(value float64) (time.Time, error) {
	var st syscall.Systemtime
	if OleAutDll.VariantTimeToSystemTime(value, &st) {
		return time.Date(int(st.Year), time.Month(st.Month), int(st.Day), int(st.Hour), int(st.Minute), int(st.Second), int(st.Milliseconds/1000), time.UTC), nil
	}
	return time.Now(), errors.New("could not convert to time, passing current time")
}
