// Package main provides ...
package manager

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type Entrypoint struct {
	Exec string
	Args []string
}

type Service struct {
	Name       string     `json:"name"`
	Path       string     `json:"path"`
	Entrypoint Entrypoint `json:"entrypoint"`
	DependsOn  []*Service `json:"dependsOn"`
}

func (s *Service) AddDependent(dependent *Service) error {

	if dependent == s {
		return errors.New("cannot use the same reference")
	}

	s.DependsOn = append(s.DependsOn, dependent)

	return nil
}

func (s *Service) Execute(ctx context.Context) error {

	var dir string

	if s.Path != "" {

		fileStat, err := os.Stat(s.Path)

		if err != nil {
			return err
		}

		if !fileStat.IsDir() {
			return errors.New("Need to be a dir")
		}

		if err := os.Chdir(s.Path); err != nil {
			return err
		}

		dir = s.Path
	}

	cmd := exec.CommandContext(ctx, s.Entrypoint.Exec, s.Entrypoint.Args...)

	cmd.Dir = dir

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println(cmd.Process.Pid)

	return nil
}
