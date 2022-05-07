package descriptor

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	specialCharacters = []string{"/", "-", ",", "*"}
	maxColumnLen      = 14
)

// Descriptor defines the contract needed for
// an expression to be calculated
type Descriptor interface {
	GetTitle() string
	GetMin() int
	GetMax() int
	Expression() string
}

// Describe builds the output combination of the
// title and the provided expression for a Descriptor
func Describe(descriptor Descriptor) string {
	return fmt.Sprintf("%s %s", describeTitle(descriptor), describeExpression(descriptor))
}

func describeTitle(descriptor Descriptor) string {
	title := descriptor.GetTitle()
	result := ""
	for i := 0; i <= maxColumnLen; i++ {
		if i <= len(title)-1 {
			result += string(title[i])
		} else {
			result += " "
		}
	}
	return result
}

// describeExpression converts a cron expression into
// a human readable string. All the conversion work happens
// in this method
func describeExpression(descriptor Descriptor) string {
	expression := descriptor.Expression()
	// check if the descriptor is of type Command
	// and if so just return the expression for the command
	switch descriptor.(type) {
	case Command:
		return expression
	}

	result := ""
	// check if it contains any special char and act on it
	if strings.ContainsAny(expression, strings.Join(specialCharacters, "")) {
		if strings.Contains(expression, "*") {
			if strings.Contains(expression, "*/") {
				segments := strings.Split(expression, "*/")
				val, err := strconv.Atoi(segments[1])
				if err != nil {
					return "error"
				}
				for i := descriptor.GetMin(); i <= descriptor.GetMax(); i++ {
					if i%val == 0 {
						result += fmt.Sprint(i) + " "
					}
				}
				return result
			} else {
				for i := descriptor.GetMin(); i <= descriptor.GetMax(); i++ {
					result += fmt.Sprint(i) + " "
				}
				return result
			}
		} else if strings.Contains(expression, ",") {
			for _, val := range strings.Split(expression, ",") {
				seg, err := strconv.Atoi(val)
				if err != nil {
					return "error"
				}
				if seg >= descriptor.GetMin() || seg <= descriptor.GetMax() {
					result += fmt.Sprint(seg) + " "
				}
			}
			return result
		} else if strings.Contains(expression, "-") {
			segments := strings.Split(expression, "-")
			start, err := strconv.Atoi(segments[0])
			if err != nil {
				return "error"
			}
			end, err := strconv.Atoi(segments[1])
			if err != nil {
				return "error"
			}
			if start < descriptor.GetMin() {
				start = descriptor.GetMin()
			}
			if end > descriptor.GetMax() {
				end = descriptor.GetMax()
			}
			for i := start; i <= end; i++ {
				result += fmt.Sprint(i) + " "
			}
			return result
		}
	} else {
		val, err := strconv.Atoi(expression)
		if err != nil {
			// is an expression and not just a single number
			return expression
		}
		// below we make sure that the value provided falls within
		// the min and max values of the descriptor and if not we
		// default to the min/max respectively for the descriptor
		if val < descriptor.GetMin() {
			val = descriptor.GetMin()
		} else if val > descriptor.GetMax() {
			val = descriptor.GetMax()
		}
		return fmt.Sprint(val)
	}
	// by default always return the expression
	return expression
}
