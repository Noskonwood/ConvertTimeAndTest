package app

import "time"

func ConvertTime(s string) (string, error) {
	t, err := time.Parse("3:05:06PM", s)
	if err != nil {
		return "", err
	}
	return t.Format("15:05:06"), nil
}
