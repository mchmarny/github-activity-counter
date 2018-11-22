# github-activity-counter

Simple GitHub activity event counter to get a real-time visibility into the repo collaboration activity. It captures series of GitHub WebHook events and extracts normalized activity data.

>> Warning, this readme is still not complete

## Supported Events

* issue_comment
* commit_comment
* issues
* pull_request
* pull_request_review_comment
* pull_request_review
* push

## Extracted Data

* ID (string) - WebHook delivery ID, immutable even when the same event is submitted multiple times
* Type (string) - GitHUb event type, e.g. commit_comment
* EventAt (time.Time) - True event time, not the WebHook processing time (with exception of push which doesn't have push time and could include multiple commits)
* Repo (string) - Fully-qualified name of the repository, e.g. mchmarny/github-activity-counter
* Actor (string) - GitHub username of the actor who initialized that event, e.g. PR author vs the PR merger who could be a automation tool like prow
* Raw (json.RawMessage) - Full content fo the GitHub WebHook payload
* Countable (bool) - Indicator whether event was parsed or not one of the types that are counted (e.g. check_run)

## Setup

To setup `github-activity-counter` you will have to:

* Deploy the code to runtime (e.g. Google Cloud Functions)
* Setup GitHub WebHook

### Deploy the code

>> Assumes you already configured GCP account, project and gcloud

First, edit the `FN_SECRET` variable in `Makefile` to some auto-generated, opaque string. Then you should be able to deploy the function to GCF using

```shell
make deploy
```

The response from the deployment will be

```shell
Deploying function (may take a while - up to 2 minutes)
```

Followed by metadata about your function. The one we need to capture will be the `httpsTrigger`

```shell
httpsTrigger:
  url: https://us-central1-s9-demo.cloudfunctions.net/github-event-handler
```

If you ever forget that, you can look up the URL of your function using

```shell
make url-lookup
```

The first time you deploy, your function will be private by default. To expose it to world:

```shell
make policy
```

>> Your GitHub WebHook will include secret so only the GitHub activity will be counted


### Setup GitHub WebHook

GitHub has good [instructions](https://developer.github.com/webhooks/creating/) on how to setup your WebHook. In short it amounts to:

* Signing to GitHub, and navigating to repo or org settings
* Clicking "Webhooks" on the left panel
* Click on the "Add WebHook" button
* Pasting your deployed handler's URL (from the deployment step)
* Selecting "application/json" as the content type
* Select "Send me everything" or select individual events you want to count (see supported events)
* Leave the "Active" checkbox checked
* Click on "Add Webhook" to save your settings

