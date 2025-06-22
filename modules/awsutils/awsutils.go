package awsutils

import (
    "os"
    "log/slog"
    "context"
    "gopkg.in/yaml.v3"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

    appconfig "github.com/mcdaigle1/slaslosli/config"
)

type AWSSecrets struct {
    PrometheusUsername string  `yaml:"prometheus_user_name"`
    PrometheusPassword string  `yaml:"prometheus_password"`
}

var Secrets AWSSecrets

func LoadSecrets() error {

    awsSecretARN := appconfig.Global.Aws.SecretARN
//    awsSecretName := appconfig.Global.Aws.SecretName
    ssoProfile := appconfig.Global.Aws.SsoProfile

    os.Setenv("AWS_PROFILE", ssoProfile)

	// Load the AWS config
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithSharedConfigProfile("Administrator-450287579526"))
	if err != nil {
		slog.Error("failed to load AWS config: ", "error", err)
        return err
	}

    // Create a Secrets Manager client
	svc := secretsmanager.NewFromConfig(cfg)

    slog.Debug("getting secret: " + awsSecretARN)
	// Get the secret value
	secrets, err := svc.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: &awsSecretARN,
	})
	if err != nil {
		slog.Error("failed to retrieve AWS secret: ", "error", err)
	}

    err = yaml.Unmarshal([]byte(*secrets.SecretString), &Secrets)
    if err != nil {
        slog.Error("Error unmarshalling config when initially populating config: %v", err)
    }

    return nil
}