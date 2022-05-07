package descriptor

// Command represents the command Descriptor to execute
// on the provided cron schedule
type Command struct {
	Value string
}

func (m Command) GetTitle() string {
	return "command"
}

func (m Command) GetMin() int {
	return 0
}

func (m Command) GetMax() int {
	return 0
}

func (m Command) Expression() string {
	return m.Value
}
