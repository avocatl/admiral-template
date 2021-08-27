// Package actions contains a set of executable functions
// that commands use to fulfill their purpose.
//
// An action is a function that receives a pointer to the
// command object and a slice of arguments.
//
// ```golang
// func(cmd *cobra.Command, args []string) {}
// ```
// When actions must return an error, the signature changes
// to the following.
//
// ```golang
// func(cmd *cobra.Command, args []string) error {}
// ```
package actions
