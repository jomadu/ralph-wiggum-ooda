package prompt

import (
	"bytes"
	"fmt"
	"text/template"
)

// ProcessTemplate executes Go text/template processing on content with the provided parameters.
// Returns the processed content or an error if template parsing or execution fails.
func ProcessTemplate(content string, params map[string]interface{}) (string, error) {
	tmpl, err := template.New("fragment").Parse(content)
	if err != nil {
		return "", fmt.Errorf("template parse error: %v", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, params)
	if err != nil {
		return "", fmt.Errorf("template execution error: %v", err)
	}

	return buf.String(), nil
}
