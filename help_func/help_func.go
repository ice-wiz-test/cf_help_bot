package help_func

import (
	"bytes"
	"fmt"
)

func Max(data []int) int {
	max := -10000000
	for i := 0; i < len(data); i++ {
		if max < data[i] {
			max = data[i]
		}
	}
	return max
}

func ConvertMaptToString(m map[string]int) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		valueConv := fmt.Sprint(value)
		fmt.Fprintf(b, "%s=\"%s\"\n", key, valueConv)
	}
	return b.String()
}

func Convert_lang_to_DB_like(lang string) int {
	var lang_db int
	if lang == "eng" {
		lang_db = 0
	} else if lang == "rus" {
		lang_db = 1
	}
	return lang_db
}
