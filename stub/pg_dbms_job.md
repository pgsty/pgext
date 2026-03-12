

## Usage

> [pg_dbms_job: Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL](https://github.com/MigOpsRepos/pg_dbms_job)

### Enabling

```sql
CREATE EXTENSION pg_dbms_job;
```

A dedicated scheduler daemon must be running per database:

```bash
pg_dbms_job -c /etc/pg_dbms_job/mydb-dbms_job.conf
```

### SUBMIT - Schedule a Job

```sql
BEGIN;
-- Scheduled job: run a procedure every day
CALL dbms_job.submit(
    job       => jobid,
    what      => 'CALL my_procedure();',
    next_date => current_timestamp + interval '1 minute',
    interval  => 'current_timestamp + ''1 day''::interval'
);
COMMIT;
```

When `next_date` and `interval` are omitted, the job executes immediately and asynchronously.

### BROKEN - Disable/Enable a Job

```sql
BEGIN;
CALL dbms_job.broken(12345, true);   -- disable job
CALL dbms_job.broken(12345, false);  -- re-enable job
COMMIT;
```

### CHANGE - Modify a Job

```sql
BEGIN;
CALL dbms_job.change(12345, null, null, 'current_timestamp + ''3 days''::interval');
COMMIT;
```

### INTERVAL - Change Execution Interval

```sql
BEGIN;
CALL dbms_job.interval(12345, 'current_timestamp + ''1 hour''::interval');
COMMIT;
```

### NEXT_DATE - Change Next Execution Date

```sql
BEGIN;
CALL dbms_job.next_date(12345, current_timestamp + interval '30 minutes');
COMMIT;
```

### WHAT - Change Job Code

```sql
BEGIN;
CALL dbms_job.what(12345, 'CALL new_procedure();');
COMMIT;
```

### REMOVE - Delete a Job

```sql
BEGIN;
CALL dbms_job.remove(12345);
COMMIT;
```

### RUN - Execute Immediately

```sql
CALL dbms_job.run(12345);
```

### Viewing Jobs

```sql
SELECT * FROM dbms_job.all_jobs;
```

### Execution History

```sql
SELECT * FROM dbms_job.all_scheduler_job_run_details;
```
