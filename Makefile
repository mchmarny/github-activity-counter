GCP_PROJECT=s9-demo
GCP_REGION=us-central1
PUBSUB_EVENTS_TOPIC=github-events
GCP_FN_NAME=github-event-handler
BQ_SCHEMA_NAME=github
BQ_TABLE_NAME=events
RUN_ID := $(shell /bin/date "+%Y-%m-%d-%H-%M-%S")
# HOOK_SECRET=some-super-long-secret-string //defined in envvars

all: url

deploy:
	gcloud alpha functions deploy $(GCP_FN_NAME) \
		--entry-point GitHubEventHandler \
		--set-env-vars HOOK_SECRET=$(HOOK_SECRET),PUBSUB_EVENTS_TOPIC=$(PUBSUB_EVENTS_TOPIC) \
		--memory 128MB \
		--region $(GCP_REGION) \
		--runtime go111 \
		--trigger-http

policy:
	gcloud alpha functions add-iam-policy-binding $(GCP_FN_NAME) \
		--region $(GCP_REGION) \
		--member allUsers \
		--role roles/cloudfunctions.invoker

url:
	gcloud alpha functions describe github-event-handler \
		--region $(GCP_REGION) \
		--format='value(httpsTrigger.url)'

topic:
	gcloud beta pubsub topics create ${PUBSUB_EVENTS_TOPIC}

table:
	bq mk $(BQ_SCHEMA_NAME)
	bq mk --schema id:string,repo:string,type:string,actor:string,event_time:timestamp,countable:boolean -t $(BQ_SCHEMA_NAME).$(BQ_TABLE_NAME)

job:
	gcloud beta dataflow jobs run $(GCP_FN_NAME)-$(RUN_ID) \
  		--gcs-location gs://dataflow-templates/pubsub-to-bigquery/template_file \
  		--parameters="topic=projects/${GCP_PROJECT}/topics/${PUBSUB_EVENTS_TOPIC}","table=${GCP_PROJECT}:$(BQ_SCHEMA_NAME).$(BQ_TABLE_NAME)"

test:
	go test ./... -v

cover:
	go test ./... -cover
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

deps:
	go mod tidy

docs:
	godoc -http=:8888 &
	open http://localhost:8888/pkg/github.com/mchmarny/github-activity-counter/
	# killall -9 godoc
