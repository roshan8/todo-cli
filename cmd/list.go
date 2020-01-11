package cmd

import (
	"fmt"
	"os"

	"github.com/todo-cli/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks in your todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, error := db.AllTasks()
		if error != nil {
			fmt.Println("Something went wrong: ", error.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no task to complete! Take a break :p")
			return
		}
		fmt.Println("You have the following tasks")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
