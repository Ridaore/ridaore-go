package ridaore

import "fmt"

type RidaoreError struct {
	Message string;
	Fix string;
}

func (e *RidaoreError) Error() string {
	msg := fmt.Sprintf("ERROR! %s, Cause by %s\n", e.Message, e.Fix);
	return msg;
}