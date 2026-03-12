


## Usage

> [pgagent: A PostgreSQL job scheduler](https://www.pgadmin.org/docs/pgadmin4/development/pgagent.html)

pgAgent is a job scheduling agent for PostgreSQL, capable of running multi-step batch/shell scripts and SQL tasks on complex schedules. It runs as a daemon and stores job definitions in the database.

### Core Concepts

- **Job**: A named schedulable unit containing one or more steps and schedules
- **Step**: An individual action (SQL script or OS batch/shell command) within a job
- **Schedule**: Defines when a job runs, with cron-like flexibility

### Job Management via SQL

pgAgent stores its configuration in the `pgagent` schema. Jobs can be managed through pgAdmin or directly via SQL.

```sql
-- View all jobs
SELECT jobid, jobname, jobenabled, jobdesc
FROM pgagent.pga_job;

-- View job steps
SELECT jstid, jstjobid, jstname, jstenabled, jstkind, jstcode
FROM pgagent.pga_jobstep;

-- View job schedules
SELECT jscid, jscjobid, jscname, jscenabled,
       jscstart, jscend, jscminutes, jschours,
       jscweekdays, jscmonthdays, jscmonths
FROM pgagent.pga_schedule;

-- View job execution log
SELECT * FROM pgagent.pga_joblog
WHERE jlgjobid = 1 ORDER BY jlgstart DESC;

-- View step execution log
SELECT * FROM pgagent.pga_jobsteplog
WHERE jsljlgid IN (SELECT jlgid FROM pgagent.pga_joblog WHERE jlgjobid = 1)
ORDER BY jslstart DESC;
```

### Step Types

| Kind | Description |
|------|-------------|
| `s` | SQL script executed against a database |
| `b` | Batch/shell command executed on the OS |

### Schedule Fields

| Field | Description |
|-------|-------------|
| `jscstart` / `jscend` | Valid date range for the schedule |
| `jscminutes` | Boolean array[60]: which minutes to run |
| `jschours` | Boolean array[24]: which hours to run |
| `jscweekdays` | Boolean array[7]: which days of week |
| `jscmonthdays` | Boolean array[32]: which days of month |
| `jscmonths` | Boolean array[12]: which months |

### Security

The pgAgent daemon connects to the database using a stored connection string. Only database superusers or users granted appropriate privileges on the `pgagent` schema tables should manage jobs.
