// Package counter github-activity-counter
//
// The `github-activity-counter` is a simple GitHub activity counter to get a real-time
// visibility into the repo collaboration events. It captures series of GitHub WebHook
// events and extracts normalized activity data for configurable persistence.
//
// Terms Of Service:
//
// There are no TOS at this moment, use at your own risk authors take no responsibility.
//
//     Schemes: http, https
//     Host: 127.0.0.1:8080
//     BasePath: /
//     Version: 0.1
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
package counter
