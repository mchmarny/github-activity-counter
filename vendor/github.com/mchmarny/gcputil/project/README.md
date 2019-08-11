# gcputil/project

Many of the GPC client libraries still require `projectID` as an input parameter. For a long time the recommended practice was to set an environment variable (e.g. GCP_PROJECT, GOOGLE_CLOUD_PROJECT, GCLOUD_PROJECT or similar). Wile GCP now provides a metadata client to extract that data at runtime many libraries still demand it.

This utility exposes two simple functions to test for presence of project ID in environment variable (useful in local dev) and if not set, tries to obtain that data from the GCP metadata service. All of that wrapped in a single static function.

## Import

```shell
import "github.com/mchmarny/gcputil/project"
```

## Usage

```shell
p, err := project.GetID()
```

Or alternatively, fail if not set

```shell
p := project.GetIDOrFail()
```

Or gt the configured meta object

```shell
import "github.com/mchmarny/gcputil/meta"
```

...and get access to all the other metadata service methods, for example

```shell
name := meta.GetClient().InstanceName()
```

