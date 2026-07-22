## Usage

Sources:

- [Official README](https://github.com/umitanuki/s3_fdw/blob/ee5b25a10d0d760423571fbe3895345cb0644df5/doc/s3_fdw.md)
- [Official extension SQL](https://github.com/umitanuki/s3_fdw/blob/ee5b25a10d0d760423571fbe3895345cb0644df5/s3_fdw.sql)
- [Official PGXN distribution page](https://pgxn.org/dist/s3_fdw/)

`s3_fdw` is an experimental, read-only foreign data wrapper that streams one Amazon S3 object into a foreign table using PostgreSQL `COPY` parsing options. The reviewed release is `0.1.0` from 2011; its transport and authentication design are obsolete, so treat it as historical code rather than a production S3 integration.

### Core Workflow

Create the extension, register its foreign-data wrapper, and store the S3 credentials in a user mapping:

```sql
CREATE EXTENSION s3_fdw;

CREATE SERVER amazon_s3
  FOREIGN DATA WRAPPER s3_fdw;

CREATE USER MAPPING FOR CURRENT_USER
  SERVER amazon_s3
  OPTIONS (
    accesskey 'example-access-key',
    secretkey 'example-secret-key'
  );
```

Map one object to a foreign table. The table column definitions must agree with the object contents and the chosen `COPY` options:

```sql
CREATE FOREIGN TABLE s3_events (
  event_time timestamptz,
  payload text
)
SERVER amazon_s3
OPTIONS (
  hostname 's3-ap-northeast-1.amazonaws.com',
  bucketname 'example-bucket',
  filename 'events.tsv',
  delimiter E'\t'
);

SELECT * FROM s3_events;
```

### Options and Behavior

- `hostname`, `bucketname`, and `filename` identify the S3 endpoint, bucket, and object; all three foreign-table options are required.
- `accesskey` and `secretkey` are user-mapping options used by the legacy S3 request signer.
- Other foreign-table options are passed to PostgreSQL's `COPY` input parser, allowing settings such as `delimiter` to describe the object format.
- The wrapper implements scan callbacks only. It does not support `INSERT`, `UPDATE`, or `DELETE`, and it exposes no server-level options.

### Operational Caveats

The implementation constructs an `http://` URL and uses legacy AWS access-key signing; it does not provide modern HTTPS, region, session-token, or IAM-role handling. The secret key is stored in PostgreSQL catalog-visible user-mapping options. It also relies on a forked helper process and a Unix FIFO, and the upstream documentation explicitly calls the release unstable and unpredictable. Do not expose credentials or sensitive data through this code. For contemporary systems, use a maintained S3 FDW or an ingestion pipeline with current TLS and AWS authentication support.
