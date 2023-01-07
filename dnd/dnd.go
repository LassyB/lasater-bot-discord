package dnd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	validEndpoints = map[string]bool{
		"ability-scores":       true,
		"alignments":           true,
		"backgrounds":          true,
		"classes":              true,
		"conditions":           true,
		"damage-types":         true,
		"equipment":            true,
		"equipment-categories": true,
		"feats":                true,
		"features":             true,
		"languages":            true,
		"magic-items":          true,
		"magic-schools":        true,
		"monsters":             true,
		"proficiencies":        true,
		"races":                true,
		"rule-sections":        true,
		"rules":                true,
		"skills":               true,
		"spells":               true,
		"subclasses":           true,
		"subraces":             true,
		"traits":               true,
		"weapon-properties":    true,
	}
)

func HandleMessage(searchArray []string) string {
	url, err := generateUrl(searchArray)
	if err != nil {
		return err.Error()
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var prettyJson bytes.Buffer
	err = json.Indent(&prettyJson, body, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	return formatMessage(prettyJson.String())
}

func generateUrl(searchArray []string) (string, error) {
	url := ""
	err := checkEndpoints(searchArray)
	if err != nil {
		return "", err
	}
	switch len(searchArray) {
	case 0:
		url = "https://www.dnd5eapi.co/api/"
	case 1:
		url = fmt.Sprintf("https://www.dnd5eapi.co/api/%s", searchArray[0])
	case 2:
		url = fmt.Sprintf("https://www.dnd5eapi.co/api/%s/%s", searchArray[0], searchArray[1])
	case 3:
		url = fmt.Sprintf("https://www.dnd5eapi.co/api/%s/%s/%s", searchArray[0], searchArray[1], searchArray[2])
	case 4:
		url = fmt.Sprintf("https://www.dnd5eapi.co/api/%s/%s/%s/%s", searchArray[0], searchArray[1], searchArray[2], searchArray[3])
	default:
		return "", errors.New("Invalid input to DND API.")
	}
	return url, nil
}

func checkEndpoints(searchArray []string) error {
	if len(searchArray) > 0 && !validEndpoints[searchArray[0]] {
		return errors.New("Not a valid endpoint. Please use !dnd for a list.")
	}
	if len(searchArray) > 3 && !validEndpoints[searchArray[2]] {
		errorMessage := fmt.Sprintf("Not a valid endpoint. Please use `!dnd %s %s` for a list.", searchArray[0], searchArray[1])
		return errors.New(errorMessage)
	}
	return nil
}

func formatMessage(message string) (response string) {
	x := strings.Split(message, "\n")
	x = x[1:(len(x) - 1)]
	for i := range x {
		if strings.Contains(x[i], "\t") {
			x[i] = strings.Join(strings.Split(x[i], "\t")[1:], "\t")
		}
	}
	response = strings.Join(x, "\n")
	response = strings.ReplaceAll(response, "{", "")
	response = strings.ReplaceAll(response, "}", "")
	response = strings.ReplaceAll(response, "[", "")
	response = strings.ReplaceAll(response, "]", "")
	return
}
