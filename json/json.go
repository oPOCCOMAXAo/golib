package json

import (
	"encoding/json"
)

func Marshal(a interface{}, indent int) ([]byte, error) {
	if indent < 0 {
		indent = 0
	}
	indentS := make([]byte, indent)
	for i := range indentS {
		indentS[i] = ' '
	}
	return json.MarshalIndent(a, "", string(indentS))
}

func Stringify(a interface{}, indent int) (string, error) {
	b, err := Marshal(a, indent)
	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func Unmarshal(txt []byte, object interface{}) error {
	return json.Unmarshal(txt, &object)
}

func Parse(txt string, object interface{}) error {
	return Unmarshal([]byte(txt), object)
}
