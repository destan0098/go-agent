package go_agent

import (
	"bufio"
	_ "embed"
	"math/rand"
	"strings"
	"sync"
	"time"
)

//go:embed internal/UserAgent.txt
var textFileContent string

var (
	agents []string
	once   sync.Once

	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	mu  sync.Mutex
)

func initAgents() {
	sc := bufio.NewScanner(strings.NewReader(textFileContent))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		agents = append(agents, line)
	}
}

// Useragent وقتی rnd=true باشد یک یوزر ایجنت رندوم از فایل می‌دهد.
// اگر rnd=false باشد مقدار fallback را برمی‌گرداند.
func Useragent(rnd bool) string {
	once.Do(initAgents)

	if len(agents) == 0 {
		return "BasicAgent"
	}

	if rnd {
		mu.Lock()
		idx := rng.Intn(len(agents))
		mu.Unlock()
		return agents[idx]
	}

	// fallback ثابت
	return "BasicAgent"
}
