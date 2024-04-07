package go_agent

import (
	"bufio"
	_ "embed"
	"log"
	"math/rand"
	"os"
	"time"
)

//go:embed internal/UserAgent.txt
var textFileContent string

func Useragent(rnd bool) string {
	var randomagent string
	var agent []string
	agentfile := textFileContent
	if rnd {
		file, err := os.Open(agentfile)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Println(err)
			}
		}(file)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			agent = append(agent, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		rand.Seed(time.Now().UnixNano())
		randIdx := rand.Intn(len(agent))
		randomagent = agent[randIdx]
	} else {
		randomagent = "BasicAgent"
	}

	return randomagent
}
