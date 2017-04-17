package cmd

import (
	"github.com/rai-project/docker/cuda"
	"github.com/spf13/cobra"
)

// cudaCmd represents the cuda command
var cudaCmd = &cobra.Command{
	Use:   "cuda",
	Short: "Start CUDA docker volumes",
	Run: func(cmd *cobra.Command, args []string) {
		cuda.Serve()
	},
}

func init() {
	RootCmd.AddCommand(cudaCmd)
}
