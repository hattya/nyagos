package commands

import (
	"fmt"
	"os"
	"syscall"

	"../conio"
	. "../interpreter"
)

func cmd_del(cmd *Interpreter) (ErrorLevel, error) {
	n := len(cmd.Args)
	if n <= 1 {
		fmt.Fprintln(cmd.Stderr, "Usage: del   [/q] FILE(S)...")
		fmt.Fprintln(cmd.Stderr, "       erase [/q] FILE(S)...")
		return NOERROR, nil
	}
	all := false
	errorcount := ErrorLevel(0)
	i := 1
	for _, arg1 := range cmd.Args[1:] {
		if arg1 == "/q" {
			all = true
			n--
			continue
		}
		path := arg1
		stat, err := os.Lstat(path)
		if err != nil {
			fmt.Fprintf(cmd.Stdout, "(%d/%d) %s: %s\n", i, n-1, path, err)
			errorcount++
			continue
		}
		if mode := stat.Mode(); mode.IsDir() {
			fmt.Fprintf(cmd.Stdout, "(%d/%d) %s is directory and passed.\n",
				i, n-1, path)
			errorcount++
			continue
		}
		if all {
			fmt.Fprintf(cmd.Stdout, "(%d/%d) %s: Remove ", i, n-1, path)
		} else {
			fmt.Fprintf(cmd.Stdout,
				"(%d/%d) %s: Remove ? [Yes/No/All/Quit] ",
				i, n-1, path)
			ch := conio.GetCh()
			fmt.Fprintf(cmd.Stdout, "%c ", ch)
			switch ch {
			case 'q', 'Q':
				fmt.Fprintln(cmd.Stdout)
				return errorcount, nil
			case 'y', 'Y':
				break
			case 'a', 'A':
				all = true
			default: // for 'n','N'
				fmt.Println("-> canceled")
				continue
			}
		}
		err = syscall.Unlink(path)
		if err != nil {
			fmt.Printf("-> %s\n", err)
			errorcount++
			continue
		}
		fmt.Println("-> done.")
		i++
	}
	return errorcount, nil
}
