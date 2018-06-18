package cli

import (
	"fmt"
	"os"

	"github.com/containerum/chkit/pkg/cli/postrun"
	"github.com/containerum/chkit/pkg/cli/prerun"
	"github.com/containerum/chkit/pkg/cli/solution"
	"github.com/containerum/chkit/pkg/context"
	"github.com/containerum/chkit/pkg/util/angel"
	"github.com/spf13/cobra"
)

func Run(ctx *context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "run",
		Short: "Run solution from public template.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := prerun.PreRun(ctx); err != nil {
				angel.Angel(ctx, err)
				os.Exit(1)
			}
			if err := prerun.GetNamespaceByUserfriendlyID(ctx, cmd.Flags()); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
		PersistentPostRun: postrun.PostRunFunc(ctx),
	}
	command.AddCommand(
		clisolution.Run(ctx),
	)
	command.PersistentFlags().
		StringP("namespace", "n", ctx.Namespace.ID, "")
	command.PersistentFlags().
		BoolP("help", "h", false, "Print help for chkit")
	return command
}
