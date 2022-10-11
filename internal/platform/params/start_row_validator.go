package params

import "fmt"

type StartRowValidator struct {
	IsRequired bool
}

func (srv StartRowValidator) IsValid(value string) error {
	if srv.IsRequired && value == "" {
		return fmt.Errorf("is required")
	}

	return nil
}

func (srv StartRowValidator) KeyParam() string {
	return "start_row"
}
