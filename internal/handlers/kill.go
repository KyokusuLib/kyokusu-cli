package handlers

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/lanxre/kyokusu-cli/internal/commands"
	"github.com/lanxre/kyokusu-cli/internal/models"
)

func RunKill(ctx context.Context, input models.Input) error {
	port := commands.FindOption("p", input.Command.Options)

	var pids []string
	var err error

	if port != nil {
		n, err := strconv.Atoi(port.Value)
		if err != nil || n < 1 || n > 65535 {
			return fmt.Errorf(
				"kill: invalid port %q: must be a number from 1 to 65535",
				port.Value,
			)
		}

		pids, err = collectPids(ctx, "lsof", "-t", "-i", ":"+port.Value)
		if err != nil {
			return fmt.Errorf("kill: find process on port %s: %w", port.Value, err)
		}
	} else {
		pids, err = collectPids(ctx, "pgrep", "kyokusu-")
		if err != nil {
			return fmt.Errorf("kill: find kyokusu processes: %w", err)
		}
	}

	if len(pids) == 0 {
		return nil
	}

	args := append([]string{"-9"}, pids...)
	if err := exec.CommandContext(ctx, "kill", args...).Run(); err != nil {
		return fmt.Errorf("kill: %w", err)
	}

	return nil
}

func collectPids(ctx context.Context, name string, args ...string) ([]string, error) {
	out, err := exec.CommandContext(ctx, name, args...).Output()
	pids := strings.Fields(string(out))
	if err != nil && len(pids) == 0 {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == 1 {
			return nil, nil
		}
		return nil, err
	}

	return pids, nil
}
