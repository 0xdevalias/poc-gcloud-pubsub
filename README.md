# test-gcloud-pubsub

Basic PoC/test code for using [Google Cloud Platform](https://cloud.google.com/) [PubSub](https://cloud.google.com/pubsub/) from [Go](https://golang.org/).

## Setup

You will need to create a `topic` and a `subscription` on that topic for the subscriber to use.

```
# Setup environment variables
source ./auth.sh

# Get go dependencies
dep ensure

# Create a topic
gcloud beta pubsub topics create my-test-topic --project testpubsubstuff

# Create a subscription
gcloud beta pubsub subscriptions create foo-subscription --topic my-test-topic --project testpubsubstuff

# List topics
gcloud beta pubsub topics list --project testpubsubstuff

# List subscriptions
gcloud beta pubsub subscriptions list --project testpubsubstuff
```

## Usage

```
# Setup environment variables
source ./auth.sh

# Run
go run subscriber.go
go run publisher.go
```

## Cleanup

```
# Setup environment variables
source ./auth.sh

# Remove subscription
gcloud beta pubsub subscriptions delete foo-subscription --project testpubsubstuff

# Remove topic
gcloud beta pubsub topics delete my-test-topic --project testpubsubstuff
```

## Misc

* https://cloud.google.com/docs/authentication/getting-started
    * https://console.cloud.google.com/apis/credentials/serviceaccountkey
    * `export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"`
* https://cloud.google.com/pubsub/docs/access_control
* https://cloud.google.com/pubsub/docs/quickstart-client-libraries#pubsub-client-libraries-go
* https://cloud.google.com/pubsub/docs/publisher
* https://cloud.google.com/pubsub/docs/emulator
* https://godoc.org/cloud.google.com/go/pubsub
* https://github.com/GoogleCloudPlatform/google-cloud-go
* https://github.com/GoogleCloudPlatform/golang-samples/blob/master/pubsub/pubsub_quickstart/main.go
