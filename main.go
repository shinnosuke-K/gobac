package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shinnosuke-K/gobac/job"
	"github.com/shinnosuke-K/gobac/queue"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create session: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(queue.ListJobQueue(context.Background(), sess))
	fmt.Println(job.ListJob(context.Background(), sess))
}
