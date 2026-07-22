## Usage

Sources:

- [AWS setup guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL-Lambda.html)
- [AWS function and parameter reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL-Lambda-functions.html)
- [AWS invocation examples](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL-Lambda-examples.html)

`aws_lambda` is an Amazon RDS for PostgreSQL extension that invokes AWS Lambda functions from SQL. It is an AWS-managed, provider-only component: networking and IAM must be configured on the RDS instance before the SQL interface can work.

### Core Workflow

Configure outbound HTTPS access to Lambda, attach an RDS IAM role with `lambda:InvokeFunction` for the permitted functions, and associate that role with the DB instance using the `Lambda` feature. Then connect as a user with `rds_superuser` privileges and install the extension with its `aws_commons` dependency:

```sql
CREATE EXTENSION IF NOT EXISTS aws_lambda CASCADE;
```

Build a function identifier and make a synchronous call:

```sql
SELECT *
FROM aws_lambda.invoke(
  aws_commons.create_lambda_function_arn(
    'arn:aws:lambda:us-west-1:444455556666:function:simple'
  ),
  '{"body":"Hello from PostgreSQL"}'::json,
  'RequestResponse'
);
```

The result contains `status_code`, `payload`, `executed_version`, and `log_result`. Grant non-superusers only the SQL access they need:

```sql
GRANT USAGE ON SCHEMA aws_lambda TO app_user;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA aws_lambda TO app_user;
```

### Main Interface

- `aws_commons.create_lambda_function_arn(function_name, region)` returns the composite function identifier accepted by `aws_lambda.invoke`.
- `aws_lambda.invoke` has `json` and `jsonb` overloads and accepts either a text function name/ARN or the `aws_commons` composite value.
- `invocation_type` is case-sensitive: `RequestResponse` waits for and returns the payload, `Event` queues an asynchronous invocation, and `DryRun` checks access without running the function.
- `log_type = 'Tail'` returns the last 4 KB of execution log in Base64; the default is `None`.
- `context` passes client context, while `qualifier` selects a Lambda version or alias.

### Configuration and Caveats

`aws_lambda.connect_timeout_ms` and `aws_lambda.request_timeout_ms` are dynamic timeout parameters. `aws_lambda.endpoint_override` is static and requires a database restart when changed.

AWS currently documents support for PostgreSQL 12.6+, 13.2+, 14.1+, and all RDS PostgreSQL 15–18 releases. Availability is determined by the exact RDS engine build, not by a portable community package.

SQL grants do not replace IAM authorization: the DB instance role still limits which Lambda functions can run. Private instances need NAT or a Lambda VPC endpoint; the latter also requires `rds.custom_dns_resolution = 1` and a reboot. Treat payloads, returned logs, network latency, Lambda retries, and asynchronous side effects as external-service behavior rather than transactional PostgreSQL work.
