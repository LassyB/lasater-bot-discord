package dnd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Spell struct {
	Name        string   `json:"name"`
	Description []string `json:"desc"`
	Range       string   `json:"range"`
	Ritual      bool     `json:"ritual"`
	//Damage      Damage `json:"damage"`
}

type Damage struct {
}

func (m Spell) printSpell() string {
	v := reflect.ValueOf(m)
	typeOfS := v.Type()
	var sb strings.Builder
	for i := 0; i < v.NumField(); i++ {
		sb.WriteString(fmt.Sprintf("%s:\t%v\n", typeOfS.Field(i).Name, v.Field(i).Interface()))
	}
	return sb.String()
}

func HandleMessage(category, searchTerm string) (message string) {
	resp, err := http.Get(fmt.Sprintf("https://www.dnd5eapi.co/api/%s/%s", category, searchTerm))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	switch category {
	case "spells":
		spell := Spell{}
		json.NewDecoder(resp.Body).Decode(&spell)
		return spell.printSpell()
	}
	return "I can't do that yet"
}
