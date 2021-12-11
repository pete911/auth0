package cmd

import (
	"strconv"
	"time"
)

func StringValue(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}

func TimeValue(in *time.Time) string {
	if in == nil {
		return ""
	}
	return in.Format(time.RFC822)
}

func Int64Value(in *int64) string {
	if in == nil {
		return ""
	}
	return strconv.Itoa(int(*in))
}
