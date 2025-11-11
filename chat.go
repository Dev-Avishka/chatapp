package main

import "github.com/gin-gonic/gin"

func start(c *gin.Context) {
	var messages = ReadJSON()
	c.JSON(200, gin.H{
		"messages": messages,
	})
}

func AddToChat(c *gin.Context) {
	var newMessage ReqMessage
	if err := c.BindJSON(&newMessage); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
		return
	}
	var lastID = GetLastID()
	var finalNewMessage = Message{
		ID:       lastID + 1,
		Content:  newMessage.Content,
		UserName: newMessage.UserName,
	}
	if err := AppendJSON(finalNewMessage); err != nil {
		c.JSON(500, gin.H{
			"error": "Could not save message",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Message added successfully",
	})
}

func Get_Last_Message(c *gin.Context) {
	lastMessage := GetLastObject()
	c.JSON(200, gin.H{
		"message": lastMessage,
	})
}
