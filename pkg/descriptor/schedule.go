package descriptor

// Schedule is an implementation of a Descriptor
// and represents a single cron expression tab
// i.e. minute, hour, DOM, etc
type Schedule struct {
	Title string
	Min   int
	Max   int
	Value string
}

func (m Schedule) GetTitle() string {
	return m.Title
}

func (m Schedule) GetMin() int {
	return m.Min
}

func (m Schedule) GetMax() int {
	return m.Max
}

func (m Schedule) Expression() string {
	return m.Value
}
