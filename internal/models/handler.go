package models

import "context"

type Handler func(
	ctx context.Context,
	input Input,
) error