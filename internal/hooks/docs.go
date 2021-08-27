// Package hooks contains commands hooks
// to be reusable and organized under a package.
//
// Hooks follow the same signatures as actions but
// with the difference that they are executed
// before or after the core action the command is
// performing.
//
// Example:
//
// func Verbose(cmd *cobra.Command, args []string) {
// 	if v := utils.Debug(); v {
// 		err := cmd.Flags().Set("verbose", "true")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

package hooks
