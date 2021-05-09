package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"
import "strings"

type MoexColumn struct {
	Type string `json:"type"`
	Bytes int `json:"bytes"`
	MaxSize int `json:"max_size"`
}

type MoexEngines struct {
	Metadata map[string]MoexColumn `json:"metadata"`
	Columns []string `json:"columns"`
}

type MoexResult struct {
	Engines MoexEngines `json:"engines"`
}

func main() {

	resp, err := http.Get("https://iss.moex.com/iss/engines.json")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	result := MoexResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("type MoexEngine struct {\n")

	for k, v := range result.Engines.Metadata {
		name := strings.Title(k)
		fmt.Printf("    %s %s `json:\"%s\"`\n", name, v.Type, k)
	}

	fmt.Printf("}\n\n")

}

