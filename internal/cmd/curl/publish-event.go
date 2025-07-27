package curl

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jsnfwlr/vexil/internal/api/handlers"
	"github.com/jsnfwlr/vexil/internal/db"
	"github.com/spf13/cobra"
)

func init() {
	BaseCmd.AddCommand(PublishEventCmd)
}

var PublishEventCmd = &cobra.Command{
	Use:   "publish-event",
	Short: "publish a fake flag event",
	Run:   PublishEventRun,
}

func PublishEventRun(cmd *cobra.Command, args []string) {
	client := &http.Client{}

	flag := handlers.Flag{
		Name:  "SOME_STRING",
		Type:  db.FlagTypeString,
		Value: "3cm of red three-stand nylon",
		Env:   "dev",
	}

	b, err := json.Marshal(flag)
	if err != nil {
		panic(err)
	}

	body := bytes.NewBuffer(b)

	req, err := http.NewRequest("POST", "http://localhost:9765/api/event", body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
