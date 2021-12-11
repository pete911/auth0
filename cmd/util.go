package cmd

func StringValue(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}
