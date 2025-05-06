package aws

import (
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/leothemighty/oliveplay/utils/utils"
)

var (
	// Global AWS session
	awsSession *session.Session
	initOnce   sync.Once
)

// GetSession returns a shared AWS session, initializing it if needed
func GetSession() *session.Session {
	initOnce.Do(func() {
		// Check for AWS SSO profile
		// Create session with the specified profile
		sess, err := session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Profile:           utils.AWSSession,
			Config: aws.Config{
				Region: aws.String(utils.AWSRegion),
			},
		})

		if err != nil {
			// If profile-based auth fails, fall back to environment variables
			log.Fatalf("Failed to create AWS session with profile %s: %v", utils.AWSSession, err)
			sess = session.Must(session.NewSession())
		}

		awsSession = sess
	})

	return awsSession
}
