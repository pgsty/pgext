


## Usage

> [pg_jobmon: Autonomous job logging and monitoring for PostgreSQL](https://github.com/omniti-labs/pg_jobmon)

`pg_jobmon` provides autonomous (non-transactional) logging for PostgreSQL transactions and functions. If a function fails, all log information written up to that point is preserved rather than rolled back.

```sql
CREATE SCHEMA jobmon;
CREATE EXTENSION pg_jobmon SCHEMA jobmon;
```

### Setup

The extension uses dblink to connect back to the same database (for non-transactional logging). Add credentials:

```sql
INSERT INTO jobmon.dblink_mapping_jobmon (username, pwd) VALUES ('rolename', 'rolepassword');
```

For non-standard ports:

```sql
INSERT INTO jobmon.dblink_mapping_jobmon (host, username, pwd, port)
VALUES ('localhost', 'rolename', 'rolepassword', '5999');
```

### Core Logging Functions

```sql
-- Start a new job
SELECT jobmon.add_job('My Job Name');

-- Add a step to the job
SELECT jobmon.add_step(job_id, 'Step description');

-- Update step status
SELECT jobmon.update_step(step_id, 'OK', 'Step completed successfully');
SELECT jobmon.update_step(step_id, 'WARNING', 'Something unexpected');

-- Close the job
SELECT jobmon.close_job(job_id);

-- Or fail the job
SELECT jobmon.fail_job(job_id);
```

### Monitoring Functions

```sql
-- Check for failed jobs
SELECT * FROM jobmon.check_job_status();

-- View job history
SELECT * FROM jobmon.job_log ORDER BY start_time DESC;

-- View step details
SELECT * FROM jobmon.job_detail WHERE job_id = 123;
```

The autonomous logging ensures that even if the parent transaction rolls back, the job log entries are preserved for troubleshooting.
