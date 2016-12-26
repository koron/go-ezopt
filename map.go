package ezopt

import "fmt"

// Map represents sub commands map.
type Map map[string]interface{}

func (m Map) run(args ...string) error {
	if len(args) < 1 {
		return ErrNoSubCommand
	}
	cmd := args[0]
	v, ok := m[cmd]
	if !ok {
		return fmt.Errorf("unknown sub-command: %s", cmd)
	}
	return Run(v, args[1:]...)
}
