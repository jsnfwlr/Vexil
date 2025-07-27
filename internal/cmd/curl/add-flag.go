package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
	"github.com/jsnfwlr/vexil/internal/db"
	"github.com/spf13/cobra"
)

func init() {
	BaseCmd.AddCommand(AddFlagCmd)
}

var AddFlagCmd = &cobra.Command{
	Use:   "add-flag",
	Short: "add a flag",
	Run:   AddFlagRun,
}

func AddFlagRun(cmd *cobra.Command, args []string) {
	client := &http.Client{}

	vt, _ := db.FlagTypeString.ToAPIEnum()
	flag := oapi.CreateFlagJSONRequestBody{
		Name:         "SOME_STRING",
		Type:         vt,
		DefaultValue: "3cm of red three-stand nylon",
	}

	b, err := json.Marshal(flag)
	if err != nil {
		panic(err)
	}

	body := bytes.NewBuffer(b)

	req, err := http.NewRequest("POST", "http://localhost:9765/api/flag", body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 300 {
		panic(fmt.Sprintf("request failed %d", resp.StatusCode))
	}
}
