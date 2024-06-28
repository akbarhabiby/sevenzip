package sevenzip

import (
	"context"
	"os/exec"
	"syscall"
)

type commandLetter string

const (
	command_add          commandLetter = "a"
	command_benchmark    commandLetter = "b" // * unused
	command_delete       commandLetter = "d"
	command_extract      commandLetter = "e"
	command_hash         commandLetter = "h"
	command_info         commandLetter = "i" // * unused
	command_list         commandLetter = "l"
	command_rename       commandLetter = "rn"
	command_test         commandLetter = "t"
	command_update       commandLetter = "u"
	command_extract_full commandLetter = "x"
)

func run(ctx context.Context, args []string, sw ...*Switches) *exec.Cmd {
	binPath, err := findPath()
	if err != nil {
		return nil
	}

	// I personally want to use Adding multiple files via stdin but the 7z isn't supported yet.
	// https://sourceforge.net/p/sevenzip/discussion/45797/thread/128e8351/
	s := MergeSwitches(sw...)
	arg := append(args, s.Args()...)
	cmd := exec.CommandContext(ctx, binPath, arg...)

	// Starting detached child process.
	// https://groups.google.com/g/golang-nuts/c/shST-SDqIp4
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	return cmd
}

// Add files to archive
func Add(archive string, source []string, sw ...*Switches) *exec.Cmd {
	return AddContext(context.Background(), archive, source, sw...)
}

// Add files to archive with Context
func AddContext(ctx context.Context, archive string, source []string, sw ...*Switches) *exec.Cmd {
	return standardFactory(ctx, command_add, archive, source, sw...)
}

// Delete files from archive
func Delete(archive string, source []string, sw ...*Switches) *exec.Cmd {
	return DeleteContext(context.Background(), archive, source, sw...)
}

// Delete files from archive with Context
func DeleteContext(ctx context.Context, archive string, source []string, sw ...*Switches) *exec.Cmd {
	return standardFactory(ctx, command_delete, archive, source, sw...)
}

// Extract files from archive (without using directory names)
func Extract(archive string, output string, sw ...*Switches) *exec.Cmd {
	return ExtractContext(context.Background(), archive, output, sw...)
}

// Extract files from archive (without using directory names) with Context
func ExtractContext(ctx context.Context, archive string, output string, sw ...*Switches) *exec.Cmd {
	return extractFactory(ctx, command_extract, archive, output, sw...)
}

// eXtract files with full path
func ExtractFull(archive string, output string, sw ...*Switches) *exec.Cmd {
	return ExtractFullContext(context.Background(), archive, output, sw...)
}

// eXtract files with full path with Context
func ExtractFullContext(ctx context.Context, archive string, output string, sw ...*Switches) *exec.Cmd {
	return extractFactory(ctx, command_extract_full, archive, output, sw...)
}

// Calculate hash values for files
func Hash(target string, sw ...*Switches) *exec.Cmd {
	return HashContext(context.Background(), target, sw...)
}

// Calculate hash values for files with Context
func HashContext(ctx context.Context, target string, sw ...*Switches) *exec.Cmd {
	return simplexFactory(ctx, command_hash, target, sw...)
}

// List contents of archive
func List(target string, sw ...*Switches) *exec.Cmd {
	return ListContext(context.Background(), target, sw...)
}

// List contents of archive with Context
func ListContext(ctx context.Context, target string, sw ...*Switches) *exec.Cmd {
	return simplexFactory(ctx, command_list, target, sw...)
}

// Rename files in archive
func Rename(archive string, source []string, sw ...*Switches) *exec.Cmd {
	return RenameContext(context.Background(), archive, source, sw...)
}

// Rename files in archive with Context
func RenameContext(ctx context.Context, archive string, source []string, sw ...*Switches) *exec.Cmd {
	return standardFactory(ctx, command_rename, archive, source, sw...)
}

// Test integrity of archive
func Test(target string, sw ...*Switches) *exec.Cmd {
	return TestContext(context.Background(), target, sw...)
}

// Test integrity of archive with Context
func TestContext(ctx context.Context, target string, sw ...*Switches) *exec.Cmd {
	return simplexFactory(ctx, command_test, target, sw...)
}

// Update files to archive
func Update(archive string, source []string, sw ...*Switches) *exec.Cmd {
	return UpdateContext(context.Background(), archive, source, sw...)
}

// Update files to archive with Context
func UpdateContext(ctx context.Context, archive string, source []string, sw ...*Switches) *exec.Cmd {
	return standardFactory(ctx, command_update, archive, source, sw...)
}
