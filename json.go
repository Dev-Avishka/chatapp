package main

import (
	"encoding/json"
	"os"
)

var Json_path = "chat.json"

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
		return fullContent
	}
	if err := json.Unmarshal(data, &fullContent); err != nil {
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
