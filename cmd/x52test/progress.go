package main

import (
	"fmt"

	pb "github.com/schollz/progressbar/v3"
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
