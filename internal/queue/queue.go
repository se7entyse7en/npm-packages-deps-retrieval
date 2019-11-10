package queue

import (
	"bufio"
	"context"
	"io"
	"os"
)

type Queue interface {
	Add(context.Context, string) error
	Remove(context.Context) (string, error)
	Len(context.Context) (int, error)

	Open(context.Context) error
	Close(context.Context) error
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

func (sq *MemoryQueue) Open(ctx context.Context) error {
	return nil
}

func (sq *MemoryQueue) Close(ctx context.Context) error {
	return nil
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

type FileQueue struct {
	MemoryQueue
	path   string
	noInit bool
}

func NewFileQueue(path string, noInit bool) *FileQueue {
	return &FileQueue{path: path, noInit: noInit}
}

func (fq *FileQueue) Open(context.Context) error {
	if fq.noInit {
		return nil
	}

	file, err := os.Open(fq.path)
	if err != nil {
		return err
	}

	defer file.Close()

	var q []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		q = append(q, line)
	}

	if err != io.EOF {
		return err
	}

	fq.q = q
	return nil
}

func (fq *FileQueue) Close(context.Context) error {
	file, err := os.OpenFile(fq.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range fq.q {
		if _, err = writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	return writer.Flush()
}
