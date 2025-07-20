package greeting

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Greet(ctx context.Context, cmd *cli.Command) error {
	switch cmd.String("lang") {
	case "en":
		fmt.Println("Hello")
	case "fr":
		fmt.Println("Salut")
	case "es":
		fmt.Println("Hola")
	default:
		fmt.Println("χαῖρε!")
	}
	return nil
}
