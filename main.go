package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ymotongpoo/goltsv"
	"github.com/yuya-takeyama/argf"
)

const lf = byte('\n')

func main() {
	var displayHelp = flag.Bool("h", false, "display help")
	var displayVersion = flag.Bool("v", false, "display version")
	flag.Parse()

	stdin := os.Stdin
	stdout := os.Stdout

	if *displayHelp {
		fmt.Fprintln(stdout, "Usage: ltsv2json [FILE]...")
		return
	}

	if *displayVersion {
		fmt.Fprintf(stdout, "ltsv2json v%s, build %s\n", Version, GitCommit)
		return
	}

	ltsv2json(ltsvReader(stdin, flag.Args()), stdout)
}

func ltsvReader(stdin io.Reader, args []string) *goltsv.LTSVReader {
	reader, err := argf.From(args)
	if err != nil {
		panic(err)
	}

	return goltsv.NewReader(reader)
}

func ltsv2json(reader *goltsv.LTSVReader, stdout io.Writer) {
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		b, err := json.Marshal(record)
		if err != nil {
			panic(err)
		}

		stdout.Write(append(b, lf))
	}
}
