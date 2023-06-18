package queue

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
)

// ListJobQueue returns a list of job queues.
func ListJobQueue(ctx context.Context, sess *session.Session) ([]string, error) {
	queueNames := []string{}
	if err := batch.New(sess).DescribeJobQueuesPagesWithContext(ctx, &batch.DescribeJobQueuesInput{}, func(page *batch.DescribeJobQueuesOutput, lastPage bool) bool {
		for _, r := range page.JobQueues {
			queueNames = append(queueNames, *r.JobQueueName)
		}
		return lastPage
	}); err != nil {
		return nil, err
	}
	sort.SliceStable(queueNames, func(i, j int) bool {
		return queueNames[i] < queueNames[j]
	})
	return queueNames, nil
}
