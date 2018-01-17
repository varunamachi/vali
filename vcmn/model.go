package vcmn

import (
	"fmt"
	"time"
)

//Version - represents version of the application
type Version struct {
	Major int `json:"major" bson:"major"`
	Minor int `json:"minor" bson:"minor"`
	Patch int `json:"patch" bson:"patch"`
}

//String - version to string
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

//ArrayMatcher - matches elements of an array. If MatchAll set to true all
//the elements of the Tags array needs to be matched, otherwise only one element
//needs to match (minimum)
type ArrayMatcher struct {
	MatchAll bool     `json:"matchAll" bson:"matchAll"`
	Tags     []string `json:"tags" bson:"tags"`
}

//DateRange - represents date ranges
type DateRange struct {
	From time.Time `json:"from" bson:"from"`
	To   time.Time `json:"from" bson:"from"`
}

//Filter - generic filter used to filter data in any mongodb collection
type Filter struct {
	Fields map[string]string `json:"fields" bson:"fields"`
	Dates  []DateRange       `json:"dates" bson:"dates"`
	Lists  []ArrayMatcher    `json:"lists" bson:"lists"`
}
