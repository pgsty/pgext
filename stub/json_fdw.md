## Usage

Sources:

- [Official README](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/README.md)
- [Extension SQL](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/json_fdw--1.0.sql)
- [FDW implementation](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/json_fdw.c)

`json_fdw` in this repository is the `json_fdw2` fork: a PostgreSQL 9.4-era foreign data wrapper for newline-delimited JSON in local, gzip-compressed, or HTTP-fetched files. It supports nested-field and array projection; its experimental remote write mapping should not be treated as a transactional database interface.

### Core Read Workflow

Each physical line must contain one complete JSON document. Nested members are mapped with quoted dotted column names. Missing members and incompatible values can appear as nulls; set `max_error_count` only when deliberately accepting malformed documents.

```sql
CREATE EXTENSION json_fdw;
CREATE SERVER json_server FOREIGN DATA WRAPPER json_fdw;

CREATE FOREIGN TABLE customer_reviews (
    customer_id text,
    "review.date" date,
    "review.rating" integer,
    "product.id" text,
    "product.similar_ids" text[]
)
SERVER json_server
OPTIONS (
    filename '/srv/json/customer_reviews.ndjson.gz',
    max_error_count '0'
);

ANALYZE customer_reviews;

SELECT customer_id, "review.rating"
FROM customer_reviews
WHERE "review.date" >= date '2026-01-01';
```

`filename` may also be an HTTP URL. Such content is cached locally and revalidated with its ETag on every query; `http_post_vars` supplies form-style POST parameters. HTTPS behavior depends on the linked libcurl build and was not tested upstream.

### Remote Operations Mapping

For remote API operations, the mutually exclusive `rom_url` and `rom_path` options select a version-2 Remote Operations Mapping document. The implementation maps reads to GET and inserts or updates to PUT; DELETE remains unimplemented.

Remote writes are external HTTP side effects and are not rolled back with a PostgreSQL transaction. The implementation does not reliably turn every libcurl or HTTP failure into a SQL error, so SQL completion is not proof that the remote resource changed. Use idempotency, application-level response validation, audit logging, and reconciliation if testing this path at all.

### Security and Compatibility

Granting use of a server can expose server-readable local files through `filename` or permit requests to internal network endpoints. Restrict FDW/server ownership and `USAGE`, allowlist paths and destinations outside the extension, and never let untrusted users define table options.

This fork requires its specific `nkhorman/yajl` `json_path` branch, libcurl, and zlib, and upstream states that it works only with PostgreSQL 9.4. It has no maintained modern compatibility matrix. Prefer a maintained FDW or an ingestion service for production workloads.
