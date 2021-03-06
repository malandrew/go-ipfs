package cli

import (
	//"fmt"
	"testing"

	"github.com/jbenet/go-ipfs/commands"
)

func TestOptionParsing(t *testing.T) {
	cmd := &commands.Command{
		Options: []commands.Option{
			commands.Option{Names: []string{"b"}, Type: commands.String},
		},
		Subcommands: map[string]*commands.Command{
			"test": &commands.Command{},
		},
	}

	opts, input, err := parseOptions([]string{"--beep", "-boop=lol", "test2", "-c", "beep", "--foo=5"})
	/*for k, v := range opts {
	    fmt.Printf("%s: %s\n", k, v)
	  }
	  fmt.Printf("%s\n", input)*/
	if err != nil {
		t.Error("Should have passed")
	}
	if len(opts) != 4 || opts["beep"] != "" || opts["boop"] != "lol" || opts["c"] != "" || opts["foo"] != "5" {
		t.Error("Returned options were defferent than expected: %v", opts)
	}
	if len(input) != 2 || input[0] != "test2" || input[1] != "beep" {
		t.Error("Returned input was different than expected: %v", input)
	}

	_, _, err = parseOptions([]string{"-beep=1", "-boop=2", "-beep=3"})
	if err == nil {
		t.Error("Should have failed (duplicate option name)")
	}

	path, args := parsePath([]string{"test", "beep", "boop"}, cmd)
	if len(path) != 1 || path[0] != "test" {
		t.Error("Returned path was defferent than expected: %v", path)
	}
	if len(args) != 2 || args[0] != "beep" || args[1] != "boop" {
		t.Error("Returned args were different than expected: %v", args)
	}
}
