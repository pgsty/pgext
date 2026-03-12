


## Usage

> [omni_aws: Amazon Web Services APIs (S3)](https://docs.omnigres.org/omni_aws/s3/)

The `omni_aws` extension provides S3-compatible API functions for PostgreSQL.

### S3 API Functions

**`s3_create_bucket(bucket, region)`** -- Creates an S3 bucket.

**`s3_list_objects_v2(bucket, path, continuation_token, delimiter, encoding_type, fetch_owner, max_keys, start_after, region)`** -- Lists objects in a bucket with pagination support.

**`s3_put_object(bucket, path, payload, content_type, region)`** -- Uploads an object. Default content type is `application/octet-stream`.

### Execution

**Single request:**

```sql
SELECT * FROM omni_aws.aws_execute(
    access_key_id => 'AKID',
    secret_access_key => 'SECRET',
    request => omni_aws.s3_put_object('my-bucket', '/path/to/file', 'content'::bytea)
);
```

**Batch requests:** Pass `requests` as an array for efficient batch operations. Returns a table with an `error` column.

### Pre-signed URLs

```sql
SELECT omni_aws.s3_presigned_url(
    bucket => 'my-bucket',
    path => '/my-file',
    access_key_id => 'AKID',
    secret_access_key => 'SECRET',
    expires => 604800,  -- 7 days (default)
    region => 'us-east-1'
);
```

### Custom Endpoints

- **DigitalOcean Spaces:** `digitalocean_s3_endpoint()`
- **Generic:** `s3_endpoint(url)` with `${region}` and `${bucket}` variables for MinIO and similar providers
