package mustgather

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/openshift/oc/pkg/cli/admin/inspect"
)

// MustGatherOptions struct creates genric IOStreams
type MustGatherOptions struct {
	genericclioptions.IOStreams
}

// NewMustGatherOptions returns Options
func NewMustGatherOptions(streams genericclioptions.IOStreams) *MustGatherOptions {
	return &MustGatherOptions{
		IOStreams: streams,
	}
}

// NewCmdMustGather returns Command
func NewCmdMustGather(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewMustGatherOptions(streams)

	cmd := &cobra.Command{
		Use:          "openshift-must-gather",
		Short:        "Gather debugging data for a given cluster operator",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.AddCommand(inspect.NewCmdInspect(streams))
	return cmd
}

// Complete returns error
func (o *MustGatherOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate returns error
func (o *MustGatherOptions) Validate() error {
	return nil
}

// Run returns error
func (o *MustGatherOptions) Run() error {
	return nil
}

func main() {
	flags := pflag.NewFlagSet("must-gather", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := NewCmdMustGather(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}

}
