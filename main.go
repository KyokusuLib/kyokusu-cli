package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lanxre/kyokusu-cli/internal"
	"github.com/lanxre/kyokusu-cli/internal/apperrors"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	if err := internal.Run(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}

		fmt.Fprintln(os.Stderr, err)
		os.Exit(int(apperrors.ErrApp))
	}
}