// Package counter main
//
// Simple Cloud Run service you can configure as a target for GitHub event
// WebHook to monitor repository (or organization) activity.Besides capturing
// the real-time event throughput metrics in Stackdriver, this service also
// normalizes the GitHub activity data and stores the results in an easy to
// query BigQuery table.
//
// Terms Of Service:
//
// There are no TOS at this moment, use at your own risk authors take no responsibility.
//
//     Schemes: http, https
//     Host: 127.0.0.1:8080
//     BasePath: /v1/github
//     Version: 0.4.3
//     License: Apache 2.0  http://www.apache.org/licenses/
//     Contact: Mark Chmarny<mark@chmarny.com> https://github.com/mchmarny/github-activity-counter
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basicAuth: []
//
//     SecurityDefinitions:
//     basicAuth:
//       type: basic
//       description:  HTTP basic authentication.
//
package main
