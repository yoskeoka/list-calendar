package util

import (
	"reflect"
	"testing"
	"time"
)

func TestGetStartDatetime(t *testing.T) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	tests := []struct {
		name string
		t    time.Time
		want time.Time
	}{
		{"time.Now()", time.Date(2020, 9, 7, 0, 0, 0, 0, jst), time.Date(2020, 9, 7, 0, 0, 0, 0, jst)},
		{"23:59:59", time.Date(2020, 1, 1, 23, 59, 59, 9, time.UTC), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"00:00:00", time.Date(2020, 9, 7, 0, 0, 0, 0, time.UTC), time.Date(2020, 9, 7, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStartDatetime(tt.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStartDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}
