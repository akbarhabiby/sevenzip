package sevenzip

import (
	"context"
	"os/exec"
)

func standardFactory(ctx context.Context, cmdLtr commandLetter, archive string, source []string, sw ...*Switches) *exec.Cmd {
	return run(ctx, append([]string{string(cmdLtr), archive}, source...), sw...)
}

func extractFactory(ctx context.Context, cmdLtr commandLetter, archive string, output string, sw ...*Switches) *exec.Cmd {
	for i := range sw {
		sw[i].OutputDirectory = &output // * Switches OutputDirectory is force replaced by output variable
	}
	return run(ctx, []string{string(cmdLtr), archive}, sw...)
}

func simplexFactory(ctx context.Context, cmdLtr commandLetter, target string, sw ...*Switches) *exec.Cmd {
	return run(ctx, []string{string(cmdLtr), target}, sw...)
}
