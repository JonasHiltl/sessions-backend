package consumer

import "time"

func nullableFloatToFloat(f *float32) float32 {
	if f == nil {
		return 0
	}
	return *f
}

func nullableStringToStr(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func nullableTimeToStr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.UTC().Format(time.RFC3339)
}

func nullableBoolToBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
