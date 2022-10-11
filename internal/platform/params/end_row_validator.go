package params

import "fmt"

type EndRowValidator struct {
	IsRequired bool
}

func (erv EndRowValidator) IsValid(value string) error {
	if erv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	return nil
}

func (erv EndRowValidator) KeyParam() string {
	return "end_row"
}
