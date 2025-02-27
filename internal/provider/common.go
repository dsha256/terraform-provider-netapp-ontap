package provider

import (
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NameResourceModel Name module for names
type NameResourceModel struct {
	Name types.String `tfsdk:"name"`
}

// ExpontentialBackoff is a function that takes in a sleepTime and maxSleepTime and returns a new sleepTime
func ExpontentialBackoff(sleepTime int, maxSleepTime int) int {
	if sleepTime < maxSleepTime {
		sleepTime = sleepTime * 2
	}
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return sleepTime
}

// StringInSlice checks if a string is in a slice of strings
func StringInSlice(str string, list []types.String) bool {
	for _, v := range list {
		if v.ValueString() == str {
			return true
		}
	}
	return false
}
