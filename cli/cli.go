package cli

import "flag"

type CliArgs struct {
	File      *string
	Tape      string
	Head      int
	State     string
	FullSpeed bool
}

func ParseFlags() CliArgs {
	fullSpeed := flag.Bool("full", false, "turn on full speed")
	tape := flag.String("tape", "", "the initial tape")
	head := flag.Int("head", 0, "the initial head position (index, 0-based)")
	state := flag.String("state", "", "the initial state")
	flag.Parse()

	var file *string
	if positional := flag.Args(); len(positional) > 0 {
		file = &positional[0]
	}

	return CliArgs{
		File:      file,
		Tape:      *tape,
		Head:      *head,
		State:     *state,
		FullSpeed: *fullSpeed,
	}
}
