package helpers

import (
	"fmt"
	"handler/cmd"
)

func GetProperties() ([]cmd.Command, error) {
    if len(cmd.CommandRegistry) == 0 {
        return nil, fmt.Errorf("no commands registered")
    }
    return cmd.CommandRegistry, nil
}
