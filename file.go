package dadjokes

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
)

type FileJoker struct {
	filePath string
	jokes    map[int]string
	mu       sync.RWMutex
	loaded   bool
}

func NewFileJokerDefault() (*FileJoker, error) {
	return NewFileJoker("data/jokes.txt")
}

func NewFileJoker(filePath string) (*FileJoker, error) {
	j := &FileJoker{
		filePath: filePath,
		jokes:    make(map[int]string),
	}
	if err := j.loadJokes(); err != nil {
		return nil, fmt.Errorf("failed to load jokes: %w", err)
	}
	return j, nil
}

func (j *FileJoker) loadJokes() error {
	j.mu.Lock()
	defer j.mu.Unlock()

	file, err := os.Open(j.filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		joke := scanner.Text()
		if joke != "" { // Skip empty lines
			j.jokes[i] = joke
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	if len(j.jokes) == 0 {
		return fmt.Errorf("no jokes found in file")
	}

	j.loaded = true
	return nil
}

func (j *FileJoker) GetJoke() (opener string, punch string) {
	j.mu.RLock()
	defer j.mu.RUnlock()

	if !j.loaded || len(j.jokes) == 0 {
		return "No jokes available"
	}

	// Get random index
	idx := rand.Intn(len(j.jokes))
	return j.jokes[idx]
}

// ReloadJokes allows refreshing jokes from file
func (j *FileJoker) ReloadJokes() error {
	return j.loadJokes()
}
