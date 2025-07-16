package event

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jsnfwlr/vexil/internal/api/handlers"
	"github.com/jsnfwlr/vexil/internal/db"
	"github.com/spf13/cobra"
)

func init() {
	BaseCmd.AddCommand(PublishCmd)
}

var PublishCmd = &cobra.Command{
	Use:   "publish",
	Short: "run migrations",
	Run:   PublishRun,
}

func PublishRun(cmd *cobra.Command, args []string) {
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
