package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ezrantn/vex/dsl"
	"github.com/gogap/go-pandoc/pandoc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var formatCmd = &cobra.Command{
	Use:   "format [source=to]",
	Short: "Format a specific file",
	Long:  `Can format your file to specific type of file`,
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		parser := dsl.NewParser(input)
		source, to, err := parser.ParseFormatType()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Fprintf(os.Stderr, "Expected syntax: source=to\n")
			return
		}

		targetPath := strings.TrimSuffix(source, filepath.Ext(source)) + "." + to

		pandocConfig := struct {
			Verbose         bool   `mapstructure:"verbose"`
			Trace           bool   `mapstructure:"trace"`
			DumpArgs        bool   `mapstructure:"dump-args"`
			IgnoreArgs      bool   `mapstructure:"ignore-args"`
			EnableFilter    bool   `mapstructure:"enable-filter"`
			EnableLuaFilter bool   `mapstructure:"enable-lua-filter"`
			SafeDir         string `mapstructure:"safe-dir"`
		}{}

		if err := viper.UnmarshalKey("pandoc", &pandocConfig); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing Pandoc config: %v\n", err)
			return
		}

		pndoc, err := pandoc.New(pandocConfig)

	},
}

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("conf")
	viper.AddConfigPath(".")

	// Read the configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}
}
