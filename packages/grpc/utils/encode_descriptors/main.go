package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Which Service Descriptor to encode: ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	commentPath := fmt.Sprintf("../../comment/comment_descriptor.pb")
	partyPath := fmt.Sprintf("../../party/party_descriptor.pb")
	storyPath := fmt.Sprintf("../../story/story_descriptor.pb")
	userPath := fmt.Sprintf("../../user/user_descriptor.pb")

	commentFile, err := os.Create("../../comment/comment_descriptor_bs64.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	partyFile, err := os.Create("../../party/party_descriptor_bs64.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	storyFile, err := os.Create("../../story/story_descriptor_bs64.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	userFile, err := os.Create("../../user/user_descriptor_bs64.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	switch input {
	case "all":
		commentData, err := os.ReadFile(commentPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		commentBs64 := base64.StdEncoding.EncodeToString(commentData)
		commentFile.WriteString(commentBs64)

		partyData, err := os.ReadFile(partyPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		partyBs64 := base64.StdEncoding.EncodeToString(partyData)
		partyFile.WriteString(partyBs64)

		storyData, err := os.ReadFile(storyPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		storyBs64 := base64.StdEncoding.EncodeToString(storyData)
		storyFile.WriteString(storyBs64)

		userData, err := os.ReadFile(userPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		userBs64 := base64.StdEncoding.EncodeToString(userData)
		userFile.WriteString(userBs64)
	case "comment":
		commentData, err := os.ReadFile(commentPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		commentBs64 := base64.StdEncoding.EncodeToString(commentData)
		commentFile.WriteString(commentBs64)
	case "party":
		partyData, err := os.ReadFile(partyPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		partyBs64 := base64.StdEncoding.EncodeToString(partyData)
		partyFile.WriteString(partyBs64)
	case "story":
		storyData, err := os.ReadFile(storyPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		storyBs64 := base64.StdEncoding.EncodeToString(storyData)
		storyFile.WriteString(storyBs64)
	case "user":
		userData, err := os.ReadFile(userPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		userBs64 := base64.StdEncoding.EncodeToString(userData)
		userFile.WriteString(userBs64)
	default:
		log.Fatal("Not a valid input")
	}
}
