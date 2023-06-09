// Code generated by goa v3.11.3, DO NOT EDIT.
//
// todo HTTP client CLI support package
//
// Command:
// $ goa gen todo/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	todoc "todo/gen/http/todo/client"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `todo (list-tasks|create-task|complete-task|revert-task)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` todo list-tasks` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		todoFlags = flag.NewFlagSet("todo", flag.ContinueOnError)

		todoListTasksFlags = flag.NewFlagSet("list-tasks", flag.ExitOnError)

		todoCreateTaskFlags    = flag.NewFlagSet("create-task", flag.ExitOnError)
		todoCreateTaskBodyFlag = todoCreateTaskFlags.String("body", "REQUIRED", "")

		todoCompleteTaskFlags  = flag.NewFlagSet("complete-task", flag.ExitOnError)
		todoCompleteTaskIDFlag = todoCompleteTaskFlags.String("id", "REQUIRED", "id of task")

		todoRevertTaskFlags  = flag.NewFlagSet("revert-task", flag.ExitOnError)
		todoRevertTaskIDFlag = todoRevertTaskFlags.String("id", "REQUIRED", "id of task")
	)
	todoFlags.Usage = todoUsage
	todoListTasksFlags.Usage = todoListTasksUsage
	todoCreateTaskFlags.Usage = todoCreateTaskUsage
	todoCompleteTaskFlags.Usage = todoCompleteTaskUsage
	todoRevertTaskFlags.Usage = todoRevertTaskUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "todo":
			svcf = todoFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "todo":
			switch epn {
			case "list-tasks":
				epf = todoListTasksFlags

			case "create-task":
				epf = todoCreateTaskFlags

			case "complete-task":
				epf = todoCompleteTaskFlags

			case "revert-task":
				epf = todoRevertTaskFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "todo":
			c := todoc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list-tasks":
				endpoint = c.ListTasks()
				data = nil
			case "create-task":
				endpoint = c.CreateTask()
				data, err = todoc.BuildCreateTaskPayload(*todoCreateTaskBodyFlag)
			case "complete-task":
				endpoint = c.CompleteTask()
				data, err = todoc.BuildCompleteTaskPayload(*todoCompleteTaskIDFlag)
			case "revert-task":
				endpoint = c.RevertTask()
				data, err = todoc.BuildRevertTaskPayload(*todoRevertTaskIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// todoUsage displays the usage of the todo command and its subcommands.
func todoUsage() {
	fmt.Fprintf(os.Stderr, `todo service for management tasks
Usage:
    %[1]s [globalflags] todo COMMAND [flags]

COMMAND:
    list-tasks: ListTasks implements listTasks.
    create-task: CreateTask implements createTask.
    complete-task: CompleteTask implements completeTask.
    revert-task: RevertTask implements revertTask.

Additional help:
    %[1]s todo COMMAND --help
`, os.Args[0])
}
func todoListTasksUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo list-tasks

ListTasks implements listTasks.

Example:
    %[1]s todo list-tasks
`, os.Args[0])
}

func todoCreateTaskUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo create-task -body JSON

CreateTask implements createTask.
    -body JSON: 

Example:
    %[1]s todo create-task --body '{
      "name": "Numquam distinctio."
   }'
`, os.Args[0])
}

func todoCompleteTaskUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo complete-task -id STRING

CompleteTask implements completeTask.
    -id STRING: id of task

Example:
    %[1]s todo complete-task --id "Totam voluptatibus adipisci eos vel."
`, os.Args[0])
}

func todoRevertTaskUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo revert-task -id STRING

RevertTask implements revertTask.
    -id STRING: id of task

Example:
    %[1]s todo revert-task --id "Eaque nihil dolore corrupti odit enim."
`, os.Args[0])
}
