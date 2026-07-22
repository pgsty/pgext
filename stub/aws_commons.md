## Usage

Sources:

- [Official Amazon RDS S3 function reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.S3Import.Reference.html#USER_PostgreSQL.S3Import.create_s3_uri)
- [Official Amazon RDS S3 import guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.S3Import.html)

`aws_commons` is an Amazon RDS for PostgreSQL helper extension that constructs the composite S3 URI and credential values consumed by `aws_s3`. It is provider-only: these objects are supplied by RDS, not by community PostgreSQL packages, and version `1.2` availability depends on the RDS engine release.

### Core Workflow

Enable the provider extensions with an RDS privileged account, then build an S3 object descriptor:

```sql
CREATE EXTENSION aws_commons;
CREATE EXTENSION aws_s3;

SELECT aws_commons.create_s3_uri(
  'example-bucket',
  'imports/events.csv',
  'us-east-1'
);
```

Pass that value to an `aws_s3` import call. The import options use PostgreSQL `COPY` syntax:

```sql
SELECT aws_s3.table_import_from_s3(
  'public.events',
  '',
  '(format csv, header true)',
  aws_commons.create_s3_uri(
    'example-bucket',
    'imports/events.csv',
    'us-east-1'
  )
);
```

### Important Objects

- `aws_commons.create_s3_uri(bucket, file_path, region)` returns an `aws_commons._s3_uri_1` value identifying one S3 object.
- `aws_commons.create_aws_credentials(access_key, secret_key, session_token)` returns an `aws_commons._aws_credentials_1` value for the credential-taking `aws_s3` overload.

### Operational Notes

Prefer an IAM role attached to the RDS instance so SQL does not contain access keys. If explicit credentials are unavoidable, use temporary credentials, restrict SQL visibility and logging, and never persist the returned composite value. Bucket policy, IAM permissions, network reachability, object format, and the target table remain separate prerequisites. Consult the documentation for the exact RDS major/minor engine because the extension cannot be installed on self-managed PostgreSQL and its functions may vary by provider release.
