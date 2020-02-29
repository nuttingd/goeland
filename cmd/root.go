package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "indigo run",
	Short: "Indigo is a simple rss to email program.",
	Long: `Indigo is a simple rss to email program.
	
It was inspired by rss2email, but is an alternative with some cool features, such as filters.

The simple way to use it is to copy the provided config.toml.sample file and customize it.

The filters available are:
	- all: default, include all entries
	- none: include no entries
	- today: only includes entries with are today
	- digest: combines all the entries into one
	- digest2: combines all the entires into one, but with one less level (<h2> instead of <h1>)
	- combine: like digest, but use the first item as the title of the digest,
	- links: rewrite src="// and href="// to have an https:// prefix
	- lebrief: fetch full articles for LeBrief by NextINpact
	- wikipedia: remove unnecessary text from wikipedia entries`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//      Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
