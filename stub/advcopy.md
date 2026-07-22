## Usage

Sources:

- [Official extension control file](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/advcopy.control)
- [Official extension SQL](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/advcopy--0.0.2.sql)
- [Official README](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/README.md)

`advcopy` 0.0.2 is an experimental PL/Python extension that moves CSV data between a PostgreSQL server, Amazon S3-compatible object storage, and a host reachable by `scp`. It executes all file, network, and shell operations from the database server, not from the SQL client.

### Core Workflow

Install the untrusted `plpython3u` language and make the Python `boto3` module available to the exact Python interpreter used by PostgreSQL. Only a superuser can create or execute untrusted PL/Python functions by default.

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION advcopy;

CREATE TABLE public.import_target (id bigint, payload text);

SELECT advcopy.import_from_s3(
    'public.import_target',
    's3://example-bucket/incoming/data.csv'
);
```

`import_from_s3` downloads the object to the database host and runs a server-side `COPY` into the supplied table. AWS credentials and region configuration are resolved by `boto3` in the PostgreSQL server process environment.

### User-Facing Functions

- `advcopy.import_from_s3(table_name text, endpoint_url text)` downloads one S3 object and imports it as CSV.
- `advcopy.import_to_s3(query text, file_name text, endpoint_url text)` is intended to export query output and upload it to S3.
- `advcopy.export_to_ip(query text, file_name text, ip text, folder text)` is intended to export query output and invoke `scp` from the database host.

### Operational Boundaries

The published 0.0.2 SQL has material defects. `advcopy.import_to_s3` assigns the URL parser to the wrong Python variable and then references undefined `urlparse`, so the documented S3 export path fails as shipped. `advcopy.export_to_ip` declares `RETURNS int` but has no return statement. Patch and test the upstream code before relying on either function.

Both S3 functions use predictable files under `/tmp`; `advcopy.import_from_s3` always uses `/tmp/s3_data.csv`. Concurrent calls can overwrite one another, failed calls can leave sensitive data behind, and the functions do not clean up temporary files. Table names and query text are interpolated into SQL, while the `scp` destination is assembled from caller input. Do not grant these functions to untrusted roles.

The extension does not define CSV options, transaction-safe object writes, retries, multipart uploads, credential management, or row-level error handling. Treat it as abandoned proof-of-concept code and prefer maintained client-side transfer tools for production data movement.
