package main

import (
	"fmt"
	"time"

	"github.com/rp1s/chzcake/pkg/log"
	"github.com/rp1s/chzcake/pkg/path"
	"github.com/spf13/cobra"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("doWork:", time.Since(start))
	}()

	rp, err := path.NormalizePath("~/chzcake/PR/cli.log")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Normalized Path:", rp)
	l, err := log.NewLogger(rp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer l.Close()

	log.Errorp(l.Log("Starting...\n"))
	log.Errorp(l.Logf("Created a log file: %s\n", rp))

	// CMD cli
	var flag string

	root := &cobra.Command{
		Use: "chz",
	}

	build := &cobra.Command{
		Use:   "build",
		Short: loadDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(flag)
		},
	}

	build.Flags().StringVar(
		&flag,
		"flag",
		"",
		"Динамическое описание флага",
	)

	gen := &cobra.Command{
		Use:   "gen",
		Short: loadDescription(),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(flag)
		},
	}

	gen.Flags().StringVar(
		&flag,
		"flag",
		"",
		"Динамическое описание флага",
	)

	root.AddCommand(build)
	root.AddCommand(gen)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}

func loadDescription() string {
	return "Динамическое описание build"
}
