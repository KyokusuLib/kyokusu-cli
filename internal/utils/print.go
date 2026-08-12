package utils

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/lanxre/kyokusu-cli/internal/models"
)

func PrintInit(w io.Writer, rootMessage models.Root) {
	fmt.Fprintln(w, rootMessage.Long)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Usage:\n\t%-3s %s\n\n", rootMessage.Usage, "[options]")
}

func PrintInputJSON(in models.Input) {
	bytes, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
