package emails

import (
	"context"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"log"
	"planetbi/config"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var sender = config.Sender

func Send(to, subject, html string) error {
	ctx := context.Background()

	// load the shared AWS configuration (~/.aws/config)
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	svc := pinpointemail.NewFromConfig(cfg)

	input := &pinpointemail.SendEmailInput{
		Destination: &types.Destination{
			CcAddresses: []string{},
			ToAddresses: []string{
				to,
			},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Html: &types.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(html),
					},
				},
				Subject: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(subject),
				},
			},
		},
		FromEmailAddress: aws.String(sender),
	}

	_, err = svc.SendEmail(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
