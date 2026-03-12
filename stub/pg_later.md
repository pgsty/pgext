

## Usage

> [pg_later: Execute SQL now and get the results later](https://github.com/tembo-io/pg_later)

A PostgreSQL extension to execute queries asynchronously. Built on [pgmq](https://github.com/pgmq/pgmq).

### Getting Started

Initialize the extension's backend:

```sql
CREATE EXTENSION pg_later CASCADE;
SELECT pglater.init();
```

Execute a SQL query now:

```sql
SELECT pglater.exec(
  'SELECT * FROM pg_available_extensions ORDER BY name LIMIT 2'
) AS job_id;
```

```text
 job_id
--------
     1
```

Come back at some later time, and retrieve the results by providing the job id:

```sql
SELECT pglater.fetch_results(1);
```

```json
{
  "query": "select * from pg_available_extensions order by name limit 2",
  "job_id": 1,
  "result": [
    {
      "name": "adminpack",
      "comment": "administrative functions for PostgreSQL",
      "default_version": "2.1",
      "installed_version": null
    },
    {
      "name": "amcheck",
      "comment": "functions for verifying relation integrity",
      "default_version": "1.3",
      "installed_version": null
    }
  ],
  "status": "success"
}
```
