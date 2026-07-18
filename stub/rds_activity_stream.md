## Usage

Sources:

- [AWS: Starting an Aurora database activity stream](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.Enabling.html)
- [AWS: Aurora PostgreSQL extension version matrix](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`rds_activity_stream` is the provider-managed component behind Amazon Aurora PostgreSQL Database Activity Streams. The current AWS extension matrix lists version `1.7`. It captures configured access and change events and delivers each Aurora cluster's encrypted audit data to a Kinesis stream in the same AWS Region.

This is a headless managed feature with no SQL activation step. Enable it for the Aurora cluster through the RDS console, CLI, or `StartActivityStream` API. For example:

```shell
aws rds start-activity-stream \
    --mode async \
    --kms-key-id my-kms-key-arn \
    --resource-arn my-cluster-arn \
    --apply-immediately
```

Aurora PostgreSQL supports synchronous or asynchronous mode. Supply a customer-selected KMS key rather than the default key. Starting the stream with `--apply-immediately` restarts the cluster immediately; without it, activation waits for the maintenance window and the stream remains stopped until that restart. Check the supported engine patch and instance-class requirements before scheduling the change.
