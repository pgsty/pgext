


## Usage

> [aws_s3: PostgreSQL extension to import/export data from/to S3](https://github.com/chimpler/postgres-aws-s3)

### Setup Credentials

Configure AWS credentials via PostgreSQL session variables:

```sql
SET aws_s3.access_key_id TO 'your_access_key';
SET aws_s3.secret_key TO 'your_secret_key';
SET aws_s3.session_token TO 'optional_session_token';  -- if using temporary credentials
```

For local development with LocalStack:

```sql
SET aws_s3.endpoint_url TO 'http://localhost:4566';
```

### Import Data from S3

```sql
CREATE EXTENSION aws_s3;

CREATE TABLE animals (
  name text,
  age int
);

SELECT aws_s3.table_import_from_s3(
  'animals',
  '',
  '(FORMAT CSV, DELIMITER '','', HEADER true)',
  'my-bucket',
  'animals.csv',
  'us-east-1'
);

SELECT * FROM animals;
```

**Parameters:** table name, column list (empty string for all), COPY options, S3 bucket, S3 key, AWS region.

### Export Data to S3

```sql
SELECT * FROM aws_s3.query_export_to_s3(
  'SELECT * FROM animals',
  'my-bucket',
  'export/animals.csv',
  'us-east-1',
  options := 'FORMAT CSV, HEADER true'
);
```

**Parameters:** SQL query, S3 bucket, S3 key, AWS region, COPY options.

### Features

- Files with `Content-Encoding=gzip` metadata are automatically decompressed during import
- Credentials can be passed directly as function arguments or via session variables
- Uses the same COPY format options as PostgreSQL (CSV, TEXT, BINARY, with all related settings like DELIMITER, HEADER, NULL, etc.)
