package cmd

// LambdaHandler is a hook for lambda functions. All arguments should be supplied as environment variables
func LambdaHandler(command string) func() {
	return func() {
		rootCmd.SetArgs([]string{command})
		rootCmd.Execute()
	}
}
