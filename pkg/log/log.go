package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rp1s/chzcake/pkg/path"
)

type Logger struct {
	PathFile string
	file     *os.File

	bs *strings.Builder
}

func NewLogger(pathFile string) (*Logger, error) {
	normalizedPath, err := path.NormalizePath(pathFile)
	if err != nil {
		return nil, fmt.Errorf("could not normalize path: %w", err)
	}

	dir := filepath.Dir(normalizedPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("could not create log directory: %w", err)
	}

	file, err := os.OpenFile(
		normalizedPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	return &Logger{
		PathFile: normalizedPath,
		file:     file,
		bs:       &strings.Builder{},
	}, nil
}

func (l *Logger) Close() {
	l.file.Close()
}

func (l *Logger) Log(message string) error {
	_, err := l.file.WriteString(message)
	if err != nil {
		return fmt.Errorf("could not write to log file: %w", err)
	}
	return nil
}

func (l *Logger) Logf(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	return l.Log(message)
}

func (l *Logger) Slog(message string) error {
	_, err := l.bs.WriteString(message)
	return err
}

func (l *Logger) Slogf(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	_, err := l.bs.WriteString(message)
	return err
}

// out log stack
func (l *Logger) Prints() error {
	_, err := l.file.WriteString(l.bs.String())
	l.bs.Reset()
	return err
}

func Errorp(err error) {
	if err != nil {
		panic(err.Error())
	}
}
