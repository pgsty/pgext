## Usage

Sources:

- [Upstream README](https://github.com/mansueli/supa_queue/blob/e44daf4506ecc71c36e4071985b28d0f0553ec89/README.md)
- [Extension control file](https://github.com/mansueli/supa_queue/blob/e44daf4506ecc71c36e4071985b28d0f0553ec89/supa_queue.control)
- [Version 1.0.4 SQL](https://github.com/mansueli/supa_queue/blob/e44daf4506ecc71c36e4071985b28d0f0553ec89/supa_queue--1.0.4.sql)
- [dbdev package page](https://database.dev/mansueli/supa_queue)

`supa_queue` version `1.0.4` is a Supabase-oriented SQL queue built on `pg_net` and `pg_cron`. Inserts into `job_queue` trigger asynchronous GET, POST, or DELETE requests; tables track in-flight jobs and workers, while scheduled functions collect responses and retry failures. The consumer URL and service-role key are read from Vault secrets.

### Example

After reviewing the installation SQL, configuring Vault, and creating the required extensions:

```sql
CREATE EXTENSION pg_net;
CREATE EXTENSION pg_cron;
CREATE EXTENSION supa_queue;
INSERT INTO job_queue (http_verb, payload, url_path)
VALUES ('POST', '{"task":"refresh"}', '/jobs');
SELECT job_id, status, retry_count FROM job_queue ORDER BY job_id DESC;
```

Installation immediately creates cron jobs, and one scheduled function sleeps twice for 20 seconds. Several functions are `SECURITY DEFINER`, two are forced into `public`, and `request_wrapper(...)` can issue arbitrary HTTP requests unless execution is revoked. Tables have RLS enabled but no policies. Before installation, fix ownership/search paths and schema placement, revoke public execution, add grants/policies, constrain destinations, protect Vault access, and define retry/idempotency and retention behavior.
