## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/README.md)
- [Version 0.0.1 database schema](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/sql/pgAutomator--0.0.1.sql)
- [Extension control file](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/pgAutomator.control)

`pgAutomator` installs the database-side catalog for an early pgAutomator job scheduler. It creates scheduling, agent, job, step, queue, connection, email, and execution-log objects in the `pgautomator` schema. The extension itself contains no background worker or executable agent, so installing it does not execute jobs; a matching external pgAutomator agent is required.

Because the canonical extension name contains an uppercase letter, quote it in SQL.

```sql
CREATE EXTENSION "pgAutomator";

SELECT pgautomator.schema_version();
SELECT job_category_id, category
FROM pgautomator.job_category
ORDER BY job_category_id;
```

### Object Model

- `pgautomator.agent` and `pgautomator.agent_active` track external agents and their polling heartbeats.
- `pgautomator.job`, `pgautomator.step`, `pgautomator.job_schedule`, and the schedule subtype tables describe work, control flow, and calendar timing.
- `pgautomator.job_execution_queue` stores due executions; `pgautomator.queue_job_execution()` populates it from enabled jobs and schedules.
- `pgautomator.job_log` and `pgautomator.step_log` record execution state. `pgautomator.job_kill` is a request table for an external agent to interpret.
- `pgautomator.step_connection`, `pgautomator.agent_smtp`, `pgautomator.job_email`, and `pgautomator.step_email` hold connection and notification configuration.

The install script seeds job categories and materializes `pgautomator.dim_date` for roughly one year starting at extension creation. A long-lived installation must refresh that materialized view and validate the scheduler's date logic before the initial range expires.

### Operational Notes

- Upstream version `0.0.1` dates from 2017 and has only a one-line README. There is no maintained deployment, agent-protocol, upgrade, or current PostgreSQL compatibility documentation.
- The `ON_IDLE` and `ON_TRIGGER_FILE` members exist in `pgautomator.schedule_type`, but the source comment says those modes are unsupported. Use only behavior verified with the matching agent.
- SMTP passwords and database passwords are stored as plain text columns in ordinary tables. Lock down the `pgautomator` schema, backups, replicas, and logs before storing credentials.
- SQL and batch steps can execute powerful commands through the external agent. Treat job-definition writes as privileged administration and constrain the agent operating-system account and database roles.
- The schema is an early, tightly coupled data model with circular deferred foreign keys and generated schedule state. Test creation, queueing, retries, time zones, daylight-saving transitions, failure recovery, and agent compatibility before any production use.
