package cmd

import (
	"fmt"
	"gopkg.in/auth0.v5/management"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxStringLength = 50

func NewManagement() management.Management {
	m, err := management.New(ProfileConfig.Domain, management.WithClientCredentials(ProfileConfig.ClientId, ProfileConfig.ClientSecret))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return *m
}

func MapValue(in map[string]interface{}) string {
	var out []string
	for k, v := range in {
		out = append(out, fmt.Sprintf("%s: %+v", k, v))
	}
	sOut := strings.Join(out, ", ")
	return StringValue(&sOut)
}

func StringValue(in *string) string {
	if in == nil {
		return ""
	}
	v := *in
	if len(v) > maxStringLength {
		return fmt.Sprintf("%s...", v[:maxStringLength-3])
	}
	return v
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
