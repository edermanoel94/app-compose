package manager

import (
	"testing"
)

func TestEntrypoint_AddDependent(t *testing.T) {

	t.Run("should add service to be dependent", func(t *testing.T) {

		s := Service{}
		s1 := Service{}

		if err := s.AddDependent(&s1); err != nil {
			t.Error(err)
		}

		if len(s.DependsOn) != 1 {
			t.Errorf("expected %d, got %d", 1, len(s.DependsOn))
		}
	})

	t.Run("should not be the same reference", func(t *testing.T) {
		s := Service{}

		if err := s.AddDependent(&s); err == nil {
			t.Error(err)
		}
	})
}

func TestEntrypoint_Execute(t *testing.T) {

	t.Run("should execute a unix command of entrypoint with empty Path", func(t *testing.T) {

		args := []string{"-la"}
		s := Service{Name: "ls", Entrypoint: Entrypoint{Exec: "ls", Args: args}}

		if err := s.Execute(); err != nil {
			t.Error(err)
		}
	})

	t.Run("should execute a command of entrypoint in a current working directory", func(t *testing.T) {

		args := []string{"run"}

		s := Service{
			Name:       "autofin-backend",
			Path:       "/home/edercale/Workspace/autofin-backend",
			Entrypoint: Entrypoint{Exec: "make", Args: args},
		}

		if err := s.Execute(); err != nil {
			t.Error(err)
		}
	})
}
