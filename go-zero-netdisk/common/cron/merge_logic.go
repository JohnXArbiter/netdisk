package cron

import "context"

func merge(ctx context.Context, ErrCallback func()) {
	var err error

	defer func() {
		if err != nil {
			ErrCallback()
		}
	}()
}
