package helper

import (
	"github.com/spf13/viper"
	"regexp"
)

// AllowedOrigin - Used to add the cors in header
func AllowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}
