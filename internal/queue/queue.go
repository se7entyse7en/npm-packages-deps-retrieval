package queue

import (
	"bufio"
	"context"
	"os"
)

type Queue interface {
	Add(context.Context, string) error
	Remove(context.Context) (string, error)
	Len(context.Context) (int, error)
}

type MemoryQueue struct {
	q []string
}

func (sq *MemoryQueue) Add(ctx context.Context, element string) error {
	sq.q = append(sq.q, element)
	return nil
}

func (sq *MemoryQueue) Remove(ctx context.Context) (string, error) {
	el := sq.q[0]
	sq.q = sq.q[1:]
	return el, nil
}

func (sq *MemoryQueue) Len(ctx context.Context) (int, error) {
	return len(sq.q), nil
}

func NewMemoryQueueFromFile(fileName string) (*MemoryQueue, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &MemoryQueue{q: lines}, nil
}
