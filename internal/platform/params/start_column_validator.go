package params

import "fmt"

type StartColumnValidator struct {
	IsRequired bool
}

func (scv StartColumnValidator) IsValid(value string) error {
	if scv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}
	if len(value) > 2 {
		return fmt.Errorf("bad request values column")
	}

	return nil
}

func (scv StartColumnValidator) KeyParam() string {
	return "start_column"
}
