package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var Json_path = "./chat.json"

// the json will follow this structure
/*
[
  {
    "id": 0,
    "content": "Welcome",
    "UserName": "System"
  }
...
]
*/
func ReadJSON() []Message {
	fullContent := []Message{}
	data, err := os.ReadFile(Json_path)
	if err != nil {
		if os.IsNotExist(err) {
			// create empty JSON array file so subsequent reads succeed
			if werr := os.WriteFile(Json_path, []byte("[]"), 0644); werr != nil {
				fmt.Printf("failed to create %s: %v\n", Json_path, werr)
			}
			return fullContent
		}
		fmt.Printf("error reading %s: %v\n", Json_path, err)
		return fullContent
	}
	if err := json.Unmarshal(data, &fullContent); err != nil {
		fmt.Printf("error unmarshaling %s: %v\n", Json_path, err)
		return fullContent
	}
	return fullContent
}

func AppendJSON(newMessage Message) error {
	messages := ReadJSON()
	messages = append(messages, newMessage)
	data, err := json.MarshalIndent(messages, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(Json_path, data, 0644); err != nil {
		return err
	}
	return nil
}
func GetLastObject() Message {
	messages := ReadJSON()
	if len(messages) == 0 {
		return Message{}
	}
	return messages[len(messages)-1]
}

func GetLastID() int {
	messages := ReadJSON()
	if len(messages) == 0 {
		return -1
	}
	return messages[len(messages)-1].ID
}
