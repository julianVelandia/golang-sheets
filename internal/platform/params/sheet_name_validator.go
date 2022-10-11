package params

import "fmt"

type SheetsNameValidator struct {
	IsRequired bool
}

func (snv SheetsNameValidator) IsValid(value string) error {
	if snv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	return nil
}

func (snv SheetsNameValidator) KeyParam() string {
	return "sheet_name"
}
