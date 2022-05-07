package descriptor

import "testing"

var scheduleDescriptorsTestCases = []struct {
	Name       string
	Schedule   Schedule
	ShouldFail bool
}{
	{
		Name:     "working - minute",
		Schedule: Schedule{Title: "minute", Value: "*/15", Min: 0, Max: 59},
	},
	{
		Name:     "working - hour",
		Schedule: Schedule{Title: "hour", Value: "0", Min: 0, Max: 23},
	},
	{
		Name:     "working - day of month",
		Schedule: Schedule{Title: "day of month", Value: "1,15", Min: 1, Max: 31},
	},
	{
		Name:     "working - month",
		Schedule: Schedule{Title: "month", Value: "*", Min: 1, Max: 12},
	},
	{
		Name:     "working - day of week",
		Schedule: Schedule{Title: "day of week", Value: "1-5", Min: 1, Max: 7},
	},
}

func TestScheduleDescriptor(t *testing.T) {
	for _, test := range scheduleDescriptorsTestCases {
		t.Run(test.Name, func(t *testing.T) {
			if test.Schedule.GetTitle() != test.Schedule.Title {
				t.Error("invalid title returned", "received", test.Schedule.GetTitle())
			}
			if test.Schedule.GetMin() != test.Schedule.Min {
				t.Error("min value should be", test.Schedule.Min, "received", test.Schedule.GetMin())
			}
			if test.Schedule.GetMax() != test.Schedule.Max {
				t.Error("max value should be", test.Schedule.Max, "received", test.Schedule.GetMax())
			}
			if test.Schedule.Expression() != test.Schedule.Value {
				t.Error("value provide is not persisted", "received", test.Schedule.Expression())
			}
		})
	}
}
