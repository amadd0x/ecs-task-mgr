package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/cenkalti/backoff"
)

var (
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	tdef string
)

func deleteTaskDef(tdef string) (err error) {
	svc := ecs.New(sess)

	// split the family and revision
	parts := strings.Split(tdef, ":")
	family := parts[0]
	revision, _ := strconv.Atoi(parts[1]) // convert revision to an integer

	// loop until revision is zero
	start := time.Now()
	for revision > 0 {
		const maxRetries = 10
		backoffWithRetries := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetries)

		tdi := &ecs.DeregisterTaskDefinitionInput{
			TaskDefinition: aws.String(tdef),
		}

		deregister := func() error {
			_, err := svc.DeregisterTaskDefinition(tdi)
			return err
		}

		// Retry the operation using exponential backoff
		err = backoff.Retry(deregister, backoffWithRetries)
		if err != nil {
			fmt.Println("Failed to deregister task definition:", err)
			return
		}

		dtdi := &ecs.DeleteTaskDefinitionsInput{
			TaskDefinitions: []*string{&tdef},
		}

		delete := func() error {
			_, err := svc.DeleteTaskDefinitions(dtdi)
			return err
		}

		// Retry the operation using exponential backoff
		err = backoff.Retry(delete, backoffWithRetries)
		if err != nil {
			fmt.Println("Failed to delete task definition:", err)
			return
		}

		duration := time.Now()
		fmt.Printf("deleted: %s\n--Run time: %v\n\n", tdef, duration.Sub(start).Round(time.Second))
		// decrement revision
		revision--
		// create new task definition string with updated revision
		tdef = fmt.Sprintf("%s:%d", family, revision)
		// time.Sleep(1 * time.Second)
	}
	fmt.Println("Deregistration and deletion complete.")
	return
}

func convertTime(seconds int) string {
	if seconds < 60 {
		return fmt.Sprintf("%d seconds", seconds)
	} else if seconds < 3600 {
		minutes := seconds / 60
		return fmt.Sprintf("%d minutes and %d seconds", minutes, seconds%60)
	} else {
		hours := seconds / 3600
		minutes := (seconds % 3600) / 60
		return fmt.Sprintf("%d hours, %d minutes, and %d seconds", hours, minutes, seconds%60)
	}
}
