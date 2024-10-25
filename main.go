package main

import (
	"context"
	"fmt"
	"io"
	"math/rand/v2"
	"os"

	"github.com/ucarion/cli"
)

type args struct {
	Dingleberry string  `cli:"-d,--dingleberry" usage:"dingleberry string (default 'dingleberry')"`
	Bias        float64 `cli:"-b,--bias" usage:"likelihood of printing dingleberry string per character (default 0.01)"`
}

func (args) ExtendedDescription() string {
	return "echoes stdin, randomly inserting 'dingleberry'"
}

func main() {
	cli.Run(context.Background(), func(ctx context.Context, args args) error {
		if args.Dingleberry == "" {
			args.Dingleberry = "dingleberry"
		}
		if args.Bias == 0 {
			args.Bias = 0.01
		}

		if rand.Float64() < args.Bias {
			fmt.Print(args.Dingleberry)
		}

		for {
			var b [1]byte
			_, err := os.Stdin.Read(b[:])
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return fmt.Errorf("read stdin: %w", err)
			}

			if _, err := os.Stdout.Write(b[:]); err != nil {
				return fmt.Errorf("write stdout: %w", err)
			}

			if rand.Float64() < args.Bias {
				fmt.Print(args.Dingleberry)
			}
		}
	})
}
