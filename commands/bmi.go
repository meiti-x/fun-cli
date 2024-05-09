package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func BMI(context *cli.Context) error {
	weight := context.Int("weight")
	height := context.Float64("height")
	if weight < 1 || height < 1 {
		return fmt.Errorf("invalid weight or height")
	}
	fmt.Printf("your BMI is: %d", int(float64(weight)/(height*height)))
	return nil
}
