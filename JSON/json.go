package JSON

import (
	"encoding/json"
	"fmt"
)

func Marshal(a interface{}, indent int) []byte {
	if indent < 0 {
		indent = 0
	}
	indentS := make([]byte, indent)
	for i := range indentS {
		indentS[i] = ' '
	}
	out, err := json.MarshalIndent(a, "", string(indentS))
	if err != nil {
		_ = fmt.Errorf("%#v", err)
	}
	return out
}

func Stringify(a interface{}, indent int) string {
	return string(Marshal(a, indent))
}

func Unmarshal(txt []byte, object interface{}) {
	if err := json.Unmarshal(txt, &object); err != nil {
		_ = fmt.Errorf("%#v", err)
	}
}

func Parse(txt string, object interface{}) {
	Unmarshal([]byte(txt), object)
}
