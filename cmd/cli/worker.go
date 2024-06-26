package cli

import (
	"log"
	"os"
)

func main() {
	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}
	logger := ctn.Resolve("logger").(*zap.SugaredLogger)
	app := ctn.Resolve("cli").(*cli.App)
	err = app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
