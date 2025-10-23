package migrate

import (
	"os"
	"strings"
)

// Parses the commandline arguments for the migrate flag and params.
// For example `-migrate foo bar` will return:
//
//	Migrate{Cmd: "foo", Param: "bar"}
//
// If there are no extra arguments given, the fields will be empty strings.
func ParseMigrationFlags(flag string) (m Migrate) {
	args := os.Args[1:]
	pos := 0

	// Check for commandline arguments in reverse order.
	for i := len(args) - 1; i >= 0; i-- {
		arg := args[i]

		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			if strings.ToLower(arg[1:]) == flag {
				pos = i
			}
			break
		}
	}

	// Extract possible command and argument after the migrate flag.
	if pos+1 < len(args) {
		m.Cmd = strings.ToLower(args[pos+1])
		if pos+2 < len(args) {
			m.Param = strings.ToLower(args[pos+2])
		}
	}

	return
}
