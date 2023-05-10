package cmd

//import (
//	"context"
//	"github.com/spf13/cobra"
//	"sync"
//)
//
//// startCmd represents the start command
//var startCmd = &cobra.Command{
//	Use:   "start",
//	Short: "Start the service",
//	Long:  `This is a generated code starter`,
//	Run: func(cmd *cobra.Command, args []string) {
//		ctx := context.Background()
//
//		var wg sync.WaitGroup
//
//		wg.Add(1)
//		go runner.NewAPI(Config, Instance).Go(ctx, &wg)
//
//		wg.Add(1)
//		go runner.NewGRPC(config, instance).Go(ctx, &wg)
//
//		wg.Wait()
//		return nil
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(startCmd)
//}
