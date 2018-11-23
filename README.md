# github-activity-counter [![Build Status](https://travis-ci.org/mchmarny/github-activity-counter.svg?branch=master)](https://travis-ci.org/mchmarny/github-activity-counter)

Simple GitHub activity counter to get a real-time visibility into the repo collaboration events. It captures series of GitHub WebHook events and extracts normalized activity data.

## Supported Events

* [issue_comment](https://developer.github.com/v3/activity/events/types/#issuecommentevent) - Issue comment is created, edited, or deleted
* [commit_comment](https://developer.github.com/v3/activity/events/types/#commitcommentevent) - Commit comment is created
* [issues](https://developer.github.com/v3/activity/events/types/#issuesevent) - Issue is opened, edited, deleted, transferred, closed, reopened, assigned, unassigned, labeled, unlabeled, milestoned, or demilestoned
* [pull_request](https://developer.github.com/v3/activity/events/types/#pullrequestevent) - Pull request is assigned, unassigned, labeled, unlabeled, opened, edited, closed, reopened, or synchronized. (Note, also triggered when a pull request review is requested/removed)
* [pull_request_review_comment](https://developer.github.com/v3/activity/events/types/#pullrequestreviewcommentevent) - Comment on a pull request's unified diff is created, edited, or deleted
* [pull_request_review](https://developer.github.com/v3/activity/events/types/#pullrequestreviewcommentevent) - Comment on a pull request's unified diff is created, edited, or deleted
* [push](https://developer.github.com/v3/activity/events/types/#pushevent) - Push to a repository branch (also repository tag pushes)

## Why

* Getting true repo activity is complex (e.g. PR comments by author vs committed which may be tool like prow)
* GitHub build-in tools/APIs don't expose data at the right granularity (e.g. user associated org grouped by month activity)
* Most readily available SDKs/Libs address only data retrieval, and have a lot of dependencies

## Extracted Data

| Data Element | Type   | Description                                                                                                                               |
| ------------ | ------ | ----------------------------------------------------------------------------------------------------------------------------------------- |
| ID           | string | WebHook delivery ID, immutable even when the same event is submitted multiple times                                                       |
| Type         | string | GitHUb event type, e.g. commit_comment                                                                                                    |
| EventAt      | time   | True event time, not the WebHook processing time (with exception of push which doesn't have push time and could include multiple commits) |
| Repo         | string | Fully-qualified name of the repository, e.g. mchmarny/github-activity-counter                                                             |
| Actor        | string | GitHub username of the actor who initialized that event, e.g. PR author vs the PR merger who could be a automation tool like prow         |
| Raw          | json   | Full content fo the GitHub WebHook payload (used for debugging and in reprocess operations)                                               |

## Setup

To setup `github-activity-counter` you will have to:

* Deploy the code to runtime (e.g. Google Cloud Functions)
* Setup GitHub WebHook

### Deploy the code

> Assumes you already configured GCP account, project and gcloud

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
make url
```

The first time you deploy, your function will be private by default. To expose it to world:

```shell
make policy
```

> Your GitHub WebHook will include secret so only the GitHub activity will be counted


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

