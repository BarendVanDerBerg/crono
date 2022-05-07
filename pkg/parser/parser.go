package parser

import (
	"errors"
	"strings"

	"github.com/barendvanderberg/crono/pkg/descriptor"
)

// ParseArguments accepts provided arguments, validates them,
// and converts them into the relavant Descriptor
func ParseArguments(args []string) ([]descriptor.Descriptor, error) {
	descriptors := []descriptor.Descriptor{}

	// validate the provided arguments
	err := validateArgs(args)
	if err != nil {
		return descriptors, err
	}

	// get and validate the string inputs
	inputs := strings.Split(args[1], " ")
	err = validateInputs(inputs)
	if err != nil {
		return descriptors, err
	}

	//  construct a list of Descriptors
	command := ""
	for i, input := range inputs {
		switch i {
		case 0:
			descriptors = append(descriptors, descriptor.Schedule{Title: "minute", Value: input, Min: 0, Max: 59})
		case 1:
			descriptors = append(descriptors, descriptor.Schedule{Title: "hour", Value: input, Min: 0, Max: 23})
		case 2:
			descriptors = append(descriptors, descriptor.Schedule{Title: "day of month", Value: input, Min: 1, Max: 31})
		case 3:
			descriptors = append(descriptors, descriptor.Schedule{Title: "month", Value: input, Min: 1, Max: 12})
		case 4:
			descriptors = append(descriptors, descriptor.Schedule{Title: "day of week", Value: input, Min: 1, Max: 7})
		default:
			// anything past th 5th argument we assume is part of the command
			command += input + " "
		}
	}

	// add the Command descriptor and return
	return append(descriptors, descriptor.Command{Value: command}), nil
}

// validateArgs ensures the provide arguments are of the correct length
func validateArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("not enough arguments provided")
	} else if len(args) > 2 {
		return errors.New("too many arguments provided")
	}
	return nil
}

// validateInputs ensures the provide inputs are of the correct length
func validateInputs(inputs []string) error {
	if len(inputs) < 6 {
		return errors.New("not enough inputs provided")
	}
	return nil
}
