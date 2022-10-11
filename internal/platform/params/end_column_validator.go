package params

import "fmt"

type EndColumnValidator struct {
	IsRequired bool
}

func (ecv EndColumnValidator) IsValid(value string) error {
	if ecv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	if len(value) > 2 {
		return fmt.Errorf("bad request values column")
	}

	return nil
}

func (ecv EndColumnValidator) KeyParam() string {
	return "end_column"
}
