## Usage

Sources:

- [AWS Database Activity Streams overview](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/DBActivityStreams.html)
- [AWS: Starting an Aurora database activity stream](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.Enabling.html)
- [AWS Aurora PostgreSQL extension version matrix](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`rds_activity_stream` is an AWS-managed internal component behind Database Activity Streams for Aurora PostgreSQL. The AWS extension matrix lists version 1.7, but users do not install or call it with SQL. Enable the managed feature on an Aurora DB cluster to deliver encrypted database activity records to an automatically created Amazon Kinesis data stream in the same Region.

### Start the Managed Stream

Use the RDS console, API, or AWS CLI at cluster scope. The following starts asynchronous streaming immediately:

```shell
aws rds start-activity-stream \
  --mode async \
  --kms-key-id <kms-key-arn> \
  --resource-arn <cluster-arn> \
  --apply-immediately
```

Aurora PostgreSQL supports `sync` and `async` modes. Choose the mode according to the required audit guarantees and workload latency. Supply a customer-managed KMS key that the service can use; the default AWS KMS key is not accepted for this workflow.

### Restart and Scope

Starting Database Activity Streams requires a database-cluster restart. With `--apply-immediately`, AWS restarts the cluster immediately; otherwise the change waits for the next maintenance window and the stream remains stopped until the restart. The setting applies to all current and future instances in the cluster. For an Aurora global database, enable and manage the stream independently on each regional cluster.

### Operations and Security

Check the AWS engine-version, patch, Region, and instance-class prerequisites before scheduling the change. The feature creates a Kinesis stream and incurs Kinesis and related consumer/storage charges. Build and monitor a consumer, encryption access, retention, alerting, throughput, and failure handling; merely starting the stream does not create a complete audit pipeline. Treat activity records as sensitive data and grant KMS, RDS, Kinesis, and consumer permissions using least privilege. Do not issue `CREATE EXTENSION rds_activity_stream` or assume it exposes a supported SQL API.
