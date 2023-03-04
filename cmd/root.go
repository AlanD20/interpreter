package cmd

import (
	"github.com/aland20/iad/internal"
	"github.com/spf13/cobra"
)

func getRootCommand() *cobra.Command {

	var input, output string

	var rootCmd = &cobra.Command{
		Use:   "interpreter",
		Short: "interpreter is a simple scripting language interpreter",
		Long:  `Learn more at https://github.com/AlanD20/interpreter`,
		Run: func(cmd *cobra.Command, args []string) {

			// Process input file
			internal.Run(input, output)
		},
	}

	rootCmd.Flags().StringVarP(&input, "input", "i", "", "Entry file path for interpreter.")

	rootCmd.Flags().StringVarP(&output, "output", "o", "", "Output file path.")

	return rootCmd
}

func Execute() error {

	rootCmd := getRootCommand()
	// rootCmd.AddCommand(NewExampleCommand())

	return rootCmd.Execute()
}
