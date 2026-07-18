## Usage

Sources:

- [Pinned upstream README](https://github.com/Der-Henning/postgres_s3_io/blob/b9a84c9e3ee5df1821ae3ae61f89871ed6f05d2d/README.md)
- [Pinned Rust implementation](https://github.com/Der-Henning/postgres_s3_io/blob/b9a84c9e3ee5df1821ae3ae61f89871ed6f05d2d/src/lib.rs)
- [Pinned Cargo metadata](https://github.com/Der-Henning/postgres_s3_io/blob/b9a84c9e3ee5df1821ae3ae61f89871ed6f05d2d/Cargo.toml)
- [Pinned extension control file](https://github.com/Der-Henning/postgres_s3_io/blob/b9a84c9e3ee5df1821ae3ae61f89871ed6f05d2d/s3_io.control)

s3_io version 0.1.0 provides basic S3-compatible object operations as in-process Rust functions: create a bucket, put and get a complete object, and test object existence with a head request. It does not provide listing, deletion, streaming, multipart transfer, or table import/export abstractions.

### Environment and example

Before PostgreSQL starts, its service environment must supply S3_ENDPOINT_URL, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, and optionally AWS_SESSION_TOKEN. The implementation defaults the region to us-east-1 and forces path-style addressing.

```sql
CREATE EXTENSION s3_io;

SELECT s3_create_bucket('analytics');

SELECT s3_put_object(
    'analytics',
    'hello.txt',
    convert_to('hello from PostgreSQL', 'UTF8'),
    content_type => 'text/plain'
);

SELECT convert_from(
    s3_get_object('analytics', 'hello.txt'),
    'UTF8'
);

SELECT s3_object_exists_lazy('analytics', 'hello.txt');
```

The put function returns the service ETag. Get collects the entire response into one PostgreSQL bytea value.

### Current implementation risks

Each backend creates a single-thread Tokio runtime and synchronously blocks on network calls. Large objects consume backend memory, and slow or unavailable endpoints hold a database session. Apply strict object-size, statement-timeout, concurrency, and network-egress controls; this API is unsuitable for bulk transfer.

Credentials and clients are cached in backend-process memory. SQL credentials can also leak through activity views, statistics, logs, or error reports, so prefer a protected service environment and revoke public execution from all functions. Grant access only through audited wrappers that fix the endpoint and bucket.

The reviewed source uses eager default evaluation, so even supplying explicit endpoint or key arguments still attempts to read the corresponding environment variables. Its client cache key also omits the session token, allowing a later call with otherwise identical settings to reuse the earlier client. Treat these as version 0.1.0 defects. The build declares pgrx features for PostgreSQL 13 through 18 and requires superuser installation, but each target version and S3-compatible service needs separate testing.
