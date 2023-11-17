package useragent_test

import (
	"fmt"
	"testing"

	ua "github.com/medama-io/go-useragent"
	"github.com/stretchr/testify/assert"
)

var matchResults = [][]string{
	// Windows
	{ua.Safari, ua.Chrome, ua.Windows},
	{ua.Safari, ua.Chrome, ua.Windows},
	{ua.Windows, ua.IE},
	{ua.Windows, ua.IE},
	{ua.IE, ua.Windows},
	{ua.Windows, ua.IE},
	{ua.Edge, ua.Safari, ua.Chrome, ua.Windows},
	// Mac
	{ua.Safari, ua.Version, ua.MacOS},
	{ua.Safari, ua.Chrome, ua.MacOS},
	{ua.Firefox, ua.MacOS},
	{ua.Vivaldi, ua.Safari, ua.Chrome, ua.MacOS},
	{ua.Edge, ua.Safari, ua.Chrome, ua.MacOS},
	// Linux
	{ua.Firefox, ua.Linux},
	{ua.Firefox, ua.Linux},
	{ua.Firefox, ua.Linux, ua.Desktop},
	{ua.Firefox, ua.Linux, ua.Desktop},
	{ua.Safari, ua.Chrome, ua.Linux},
}

func TestMatchTokenIndexes(t *testing.T) {
	// Refer to version_test.go for versionResults test cases
	for i, v := range versionResults {
		t.Run(fmt.Sprintf("Case:%d", i), func(t *testing.T) {
			match := ua.MatchTokenIndexes(v)

			if len(match) != len(matchResults[i]) {
				t.Errorf("Test Case: %s, expected %d matches, got %d\nMatch Index: %d", v, len(match), len(matchResults[i]), i)
				t.FailNow()
			}

			for j, m := range match {
				assert.Equal(t, matchResults[i][j], m.Match, "Test Case: %s\nMatch Number: %d\nExpected: %v", v, i, match)
			}
		})
	}
}
