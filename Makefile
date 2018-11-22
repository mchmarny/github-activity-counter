# Go parameters
GCP_REGION=us-central1
FN_NAME=github-event-handler
FN_SECRET=some-super-long-secret-string

all: test

deploy:
	gcloud alpha functions deploy $(FN_NAME) \
		--entry-point GitHubEventHandler \
		--set-env-vars HOOK_SECRET=$(FN_SECRET) \
		--memory 128MB \
		--region $(GCP_REGION) \
		--runtime go111 \
		--trigger-http

policy:
	gcloud alpha functions add-iam-policy-binding $(FN_NAME) \
		--region $(GCP_REGION) \
		--member allUsers \
		--role roles/cloudfunctions.invoker

test:
	go test -v ./...

deps:
	go mod tidy
