package plugin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const metadataInvocation = "chzcake:metadata/metadata.get@1.0.0()"

var waveRecordKey = regexp.MustCompile(`([,{]\s*)([a-z][a-z0-9-]*)(\s*:)`)

// Metadata is the Go representation of chzcake:metadata/metadata.info@1.0.0.
type Metadata struct {
	Name        string
	Version     string
	Description string
}

// Plugin is a component which has passed the metadata contract check.
type Plugin struct {
	Path     string
	Metadata Metadata
}

// Loader loads WebAssembly components through Wasmtime.
//
// Wasmtime performs Component Model validation and resolves get() by its full
// versioned WIT name. Calling the function as part of Load also verifies its
// result against the WIT record at runtime.
type Loader struct {
	wasmtime string
	run      commandRunner
}

type commandRunner interface {
	Run(ctx context.Context, executable string, args ...string) ([]byte, error)
}

type execRunner struct{}

func (execRunner) Run(ctx context.Context, executable string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, executable, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		message := strings.TrimSpace(stderr.String())
		if message == "" {
			message = err.Error()
		}
		return nil, errors.New(message)
	}

	return stdout.Bytes(), nil
}

// NewLoader locates wasmtime in PATH.
func NewLoader() (*Loader, error) {
	executable, err := exec.LookPath("wasmtime")
	if err != nil {
		return nil, fmt.Errorf("find wasmtime: %w", err)
	}

	return newLoader(executable, execRunner{}), nil
}

func newLoader(executable string, runner commandRunner) *Loader {
	return &Loader{
		wasmtime: executable,
		run:      runner,
	}
}

// Load validates, instantiates and invokes
// chzcake:metadata/metadata.get@1.0.0() on a component.
func (l *Loader) Load(ctx context.Context, path string) (*Plugin, error) {
	if l == nil || l.run == nil || l.wasmtime == "" {
		return nil, errors.New("plugin loader is not initialized")
	}
	if err := validatePluginFile(path); err != nil {
		return nil, err
	}

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve plugin path %q: %w", path, err)
	}

	output, err := l.run.Run(
		ctx,
		l.wasmtime,
		"run",
		"--invoke",
		metadataInvocation,
		absolutePath,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"plugin %q does not satisfy chzcake:metadata@1.0.0: %w",
			path,
			err,
		)
	}

	metadata, err := parseMetadata(output)
	if err != nil {
		return nil, fmt.Errorf("decode metadata returned by plugin %q: %w", path, err)
	}

	return &Plugin{
		Path:     absolutePath,
		Metadata: metadata,
	}, nil
}

func validatePluginFile(path string) error {
	if strings.TrimSpace(path) == "" {
		return errors.New("plugin path is empty")
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open plugin %q: %w", path, err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("stat plugin %q: %w", path, err)
	}
	if !info.Mode().IsRegular() {
		return fmt.Errorf("plugin %q is not a regular file", path)
	}

	var magic [4]byte
	if _, err := file.Read(magic[:]); err != nil {
		return fmt.Errorf("read plugin %q: %w", path, err)
	}
	if magic != [4]byte{0x00, 0x61, 0x73, 0x6d} {
		return fmt.Errorf("plugin %q is not a WebAssembly binary", path)
	}

	return nil
}

func parseMetadata(output []byte) (Metadata, error) {
	wave := strings.TrimSpace(string(output))
	jsonRecord := waveRecordKey.ReplaceAllString(wave, `${1}"${2}"${3}`)

	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte(jsonRecord), &fields); err != nil {
		return Metadata{}, fmt.Errorf("decode WAVE record: %w", err)
	}
	for _, name := range [...]string{"name", "version", "description"} {
		if _, ok := fields[name]; !ok {
			return Metadata{}, fmt.Errorf("metadata record has no %q field", name)
		}
	}

	var metadata Metadata
	decoder := json.NewDecoder(strings.NewReader(jsonRecord))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&metadata); err != nil {
		return Metadata{}, fmt.Errorf("decode WAVE record: %w", err)
	}
	return metadata, nil
}
