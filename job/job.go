package job

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
)

// ListJob returns a list of job definitions.
func ListJob(ctx context.Context, sess *session.Session) ([]string, error) {
	unique := map[string]struct{}{}
	if err := batch.New(sess).DescribeJobDefinitionsPagesWithContext(ctx, &batch.DescribeJobDefinitionsInput{}, func(page *batch.DescribeJobDefinitionsOutput, lastPage bool) bool {
		for _, r := range page.JobDefinitions {
			if _, ok := unique[*r.JobDefinitionName]; !ok {
				unique[*r.JobDefinitionName] = struct{}{}
			}
		}
		return lastPage
	}); err != nil {
		return nil, err
	}
	jobNames := make([]string, 0, len(unique))
	for k := range unique {
		jobNames = append(jobNames, k)
	}
	sort.SliceStable(jobNames, func(i, j int) bool {
		return jobNames[i] < jobNames[j]
	})
	return jobNames, nil
}
