package darknet

import (
	"bufio"
	"os"
	"strings"
)

// Section represents a section in the configuration file.
type Section struct {
	Type    string
	Options []string
}

// readCfg reads a configuration file and returns a list of sections.
func readCfg(filename string) ([]Section, error) {
	// Open the configuration file.
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a list of sections.
	var sections []Section
	var current *Section

	// Read each line in the file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		switch line[0] {
		case '[':
			if current != nil {
				sections = append(sections, *current)
			}
			current = &Section{Type: line}
		case '#', ';', '\x00':
			// Ignore comments and empty lines.
		default:
			if current == nil {
				current = &Section{}
			}
			current.Options = append(current.Options, line)
		}
	}
	if current != nil {
		sections = append(sections, *current)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sections, nil
}
