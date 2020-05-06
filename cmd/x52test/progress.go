package main

import (
	"fmt"
	"time"

	pb "github.com/schollz/progressbar/v3"
	"nirenjan.org/saitek-x52/x52"
)

func progressBar(name string, max int) *pb.ProgressBar {
	return pb.NewOptions(max,
		pb.OptionFullWidth(),
		pb.OptionSetDescription(name),
		pb.OptionSetPredictTime(true),
		pb.OptionSetRenderBlankState(true),
		pb.OptionOnCompletion(func() {
			fmt.Println("")
		}),
	)
}

func updateDev(ctx *x52.Context, bar *pb.ProgressBar) error {
	if !mockTests {
		err := ctx.Update()
		if err != nil {
			bar.Clear()
			return err
		}
	}

	return nil
}

func delayMs(ms int64) {
	duration := ms * int64(time.Millisecond)
	if mockTests {
		duration /= 10
	}

	time.Sleep(time.Duration(duration))
}
