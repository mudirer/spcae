package suid

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func New() (string,error){

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return "",err
	}
	id := node.Generate()
	return id.String(),nil
}