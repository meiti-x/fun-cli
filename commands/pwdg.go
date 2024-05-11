package commands

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"math/rand"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+{}[]|;:,<.>?/"

// The init function in Go is a special function that serves as a constructor for a package.
//It is called automatically when a package is imported, before any other code in the package is executed.
//You can use the init function to perform setup tasks and initialize package-level variables.
func init() {
	// random function return a random param whenever the seed param changed
	rand.Seed(time.Now().UnixNano())
}

func Pwdg(context *cli.Context) error {
	length := context.Int("length")
	password := make([]byte, length)
	if length < 8 {
		return errors.New("length must be at least 8 characters")
	}

	for i := 0; i < length; i++ {
		fmt.Println(rand.Intn(length))
		password[i] += chars[rand.Intn(length)*i]
	}
	fmt.Println(string(password))
	return nil
}
