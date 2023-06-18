package main

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
)

// ListJobQueue returns a list of job queues.
func ListJobQueue(ctx context.Context, sess *session.Session) ([]string, error) {
	result := []string{}
	if err := batch.New(sess).DescribeJobQueuesPagesWithContext(ctx, &batch.DescribeJobQueuesInput{}, func(page *batch.DescribeJobQueuesOutput, lastPage bool) bool {
		for _, r := range page.JobQueues {
			result = append(result, *r.JobQueueName)
		}
		return lastPage
	}); err != nil {
		return nil, err
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result, nil
}
