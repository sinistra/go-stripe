package main

import (
    "github.com/plaid/go-envvar/envvar"
)

// Config defines the settings/configuration of the application.
type Config struct {
    Application struct {
        Version       string `envvar:"APP_VERSION" default:"1.0.0"`
        HostName      string `envvar:"APP_HOSTNAME" default:"localhost"`
        Port          int    `envvar:"API_PORT"`
        FrontEnd      string `envvar:"FRONT_END_URL" default:"http://localhost:4000"`
        Environment   string `envvar:"ENVIRONMENT" default:"development"`
        SignerKey     string `envvar:"SIGNER_KEY" default:"bRWmrwNUTqNUuzckjxsFlHZjxHkjrzKP"`
        PrivateKey    string `envvar:"CONTAINER_PRIVATE_KEY" secret:"CONTAINER_PRIVATE_KEY" default:""`
        MigrationMode bool   `envvar:"MIGRATION_MODE" default:"false"`
    }
    Storage struct {
        MySQL struct {
            Username  string `envvar:"DB_USERNAME" default:"root"`
            Password  string `envvar:"DB_PASSWORD" default:"root"`
            Name      string `envvar:"DB_NAME" default:"widgets"`
            Host      string `envvar:"DB_HOST" default:"127.0.0.1"`
            Port      int    `envvar:"DB_PORT" default:"3306"`
            SSLMode   string `envvar:"DB_SSL" default:"false"`
            Migration struct {
                Version uint   `envvar:"DB_MIGRATION_VERSION" default:"0"`
                Path    string `envvar:"DB_MIGRATION_PATH" default:""`
            }
        }
    }
    Payment struct {
        Stripe struct {
            Key    string `envvar:"STRIPE_KEY" default:"pk_test_EIJcR19fDAzMEXJxbuNRSptv00XaKwXsKj"`
            Secret string `envvar:"STRIPE_SECRET" default:"sk_test_7RpSFsv4HUe2GTQCe5PfZGoc00eeclwLSN"`
        }
    }
    Document struct {
        // AssetsPath specifies the path to static assets for document generation.
        // It is the same for all deployed envs (as configured in Dockerfile)
        // but can be optionally overridden on local.
        AssetsPath    string `envvar:"DOCUMENT_ASSETS_PATH" default:"./assets"`
        TemplatesPath string `envvar:"DOCUMENT_TEMPLATES_PATH" default:"./assets/templates"`
    }
    Notification struct {
        Email struct {
            Host               string `envvar:"SMTP_HOST" default:"smtp.mailtrap.io"`
            Port               int    `envvar:"SMTP_PORT" default:"587"`
            User               string `envvar:"SMTP_USER"  default:"30980d8770302b02e"`
            Password           string `envvar:"SMTP_PASSWORD" default:"33c58457afcc76"`
            Sender             string `envvar:"SMTP_SENDER" default:"someone@nms.com.au"`
            ReturnEmailAddress string `envvar:"RETURN_EMAIL_ADDRESS" default:"reply@nms.com.au"`
        }
        EmailQueueName string `envvar:"SMTP_EMAIL_QUEUE_NAME" default:"not-in-use"`
    }
}

// LoadEnv performs the loading of configuration values from ENV variables.
func LoadEnv() (Config, error) {
    c := Config{}
    if err := envvar.Parse(&c); err != nil {
        return Config{}, err
    }

    return c, nil
}
