package bitnamialt

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed alternatives.json
var alternatesRaw []byte

var Alternatives Entries

func init() {
	if err := json.Unmarshal(alternatesRaw, &Alternatives); err != nil {
		log.Fatal(err)
	}
}
