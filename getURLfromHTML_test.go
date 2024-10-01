package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "URLs with query parameters",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one?param=value">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one?param=value">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one?param=value", "https://other.com/path/one?param=value"},
		},
		{
			name:     "URLs with fragment",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one#fragment">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one#fragment">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one#fragment", "https://other.com/path/one#fragment"},
		},
		{
			name:     "URLs with different scheme",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="http://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "http://other.com/path/one"},
		},
		{
			name:     "URLs with no scheme and domain",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="/path/two">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
		},
		{
			name:     "URLs with relative scheme",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="//other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="/path/two">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
		},
		{
			name:     "URLs with relative scheme",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="//other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a>
			<span>Boot.dev></span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "bad HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev></span>
	</a>
</html body>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
	}
	for i, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			baseURL, err := url.Parse(tst.inputURL)
			if err != nil {
				t.Errorf("Test #%v - %s failed with unexpected error.\nError: %v\n----------\n", i, tst.name, err)
				return
			}
			actual, err := getURLsFromHTML(tst.inputBody, baseURL)
			if err != nil && !strings.Contains(err.Error(), tst.errorContains) {
				t.Errorf("Test #%v - %s failed with unexpected error.\nError: %v\n----------\n", i, tst.name, err)
				return
			} else if err != nil && tst.errorContains == "" {
				t.Errorf("Test #%v - %s failed with unexpected error, no error was expected.\nError: %v\n----------\n", i, tst.name, err)
				return
			} else if err == nil && tst.errorContains != "" {
				t.Errorf("Test #%v - %s failed with no error, an error was expected.\n----------\n", i, tst.name)
			}
			if reflect.DeepEqual(actual, tst.expected) != true {
				t.Errorf("Test #%v - %s failed.\nInput Body: %v\nInput baseURL: %v\nexpected URL: %v\nactualURL: %v\n----------\n", i, tst.name, tst.inputBody, tst.inputBody, tst.expected, actual)
			}
		})
	}
}
