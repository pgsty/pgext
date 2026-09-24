## Usage

Sources:

- [Official release v0.1.1](https://github.com/pgsty/pgs3/releases/tag/v0.1.1)
- [Official README v0.1.1](https://github.com/pgsty/pgs3/blob/v0.1.1/README.md)
- [Extension control file](https://github.com/pgsty/pgs3/blob/v0.1.1/pgs3.control)
- [Configuration reference](https://github.com/pgsty/pgs3/blob/v0.1.1/docs/guc.md)
- [Operations guide](https://github.com/pgsty/pgs3/blob/v0.1.1/docs/operations.md)
- [Known limitations](https://github.com/pgsty/pgs3/blob/v0.1.1/docs/known-limitations.md)

`pgs3` 0.1.1 turns one PostgreSQL database into a path-style S3-compatible endpoint. PostgreSQL background workers authenticate SigV4 requests and execute versioned object operations against ordinary SQL tables, so object metadata, payloads, authorization, WAL, physical backup, and recovery remain inside PostgreSQL. It targets PostgreSQL 17 and 18 and remains early-alpha software.

### Core Workflow

Install the extension in the database that will own the object store, create a restricted tenant role and credential, then start the worker pool:

```sql
CREATE EXTENSION pgs3;

CREATE ROLE tenant_app
  NOLOGIN NOINHERIT NOSUPERUSER NOCREATEDB NOCREATEROLE
  NOREPLICATION NOBYPASSRLS;

SELECT pgs3.create_credential(
  'tenant-access-key', 'replace-with-a-secret', 'tenant_app'::name, true
);

SELECT pgs3.start();
TABLE pgs3.worker_state;
TABLE pgs3.stats;
```

For automatic startup, preload the library and configure the target database before restarting PostgreSQL:

```conf
shared_preload_libraries = 'pgs3'
pgs3.enabled = on
pgs3.target_database = 'artifacts'
pgs3.listen_addr = '127.0.0.1'
pgs3.port = 9000
pgs3.workers = 4
```

Adding or removing `shared_preload_libraries`, or changing `pgs3.target_database`, requires a PostgreSQL restart. Other documented settings use SIGHUP semantics, but operational readiness must be checked through `pgs3.worker_state` and logs rather than only an open TCP port. A manually started pool can be stopped with `pgs3.stop()`.

### Client and Storage Behavior

Clients must use path-style addressing and an explicit endpoint:

```bash
export PGS3_ENDPOINT='https://s3.example.com'
export AWS_ACCESS_KEY_ID='<access-key>'
export AWS_SECRET_ACCESS_KEY='<secret-key>'
export AWS_DEFAULT_REGION='us-east-1'

aws --endpoint-url "$PGS3_ENDPOINT" s3api list-buckets
```

Always pass `--endpoint-url`; otherwise an AWS client can silently send a request to AWS. The supported client paths include AWS CLI, boto3, rclone, s3fs, and DuckDB `httpfs`. Bucket and object operations include range and conditional reads/writes, ListObjectsV2, permanent version history, delete markers, CopyObject, and multipart upload.

The canonical payload lives in `pgs3.blob`. Object versions, CopyObject, SQL Restore, and metadata-only Fork operations can share the same blob rather than duplicating bytes. One endpoint serves one configured database.

### Operations and Security

Credential access keys map to PostgreSQL tenant roles. Keep those roles `NOLOGIN`, `NOINHERIT`, and `NOBYPASSRLS`; do not use `pgs3.server_role` as an application identity. Row-level security is the tenant-isolation boundary, while credential-management and worker-control functions remain operator-only.

SigV4 requires reversibly stored secrets, so database backups and replicas contain credential material and must be encrypted and access-controlled. `pgs3` serves cleartext HTTP; production deployments need an external TLS proxy that preserves the signed path, host, and headers. Use physical backups: logical dump/restore is not supported for the extension-owned object state.

### Compatibility and Limitations

The packaged and upstream-supported paths are PostgreSQL 17 and 18 with pgrx 0.19.2. Version 0.1.1 includes the tested `0.1.0 -> 0.1.1` extension upgrade edge.

`pgs3` is not a general-purpose production S3 replacement. Virtual-host addressing, built-in TLS, IAM or bucket-policy languages, ACLs, lifecycle rules, and cross-database routing are not implemented. Small-object GET/PUT targets and the 100,000-object Fork target are not met, and full-object GET currently materializes the response in memory. Set a deployment-specific object-size limit and review the upstream limitations before exposing an endpoint.
