## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/README.md)
- [Version 0.0.1 schema](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/sql/pgAutomator--0.0.1.sql)
- [Extension control file](https://github.com/Tostino/pgAutomator-extension/blob/d1bb736bce2f046dffa7abc9f9ac81aeb57d32dd/pgAutomator.control)

`pgAutomator` installs the database-side catalog for the pgAutomator job scheduler. Because the control filename contains an uppercase letter, quote the extension name on case-sensitive systems.

```sql
CREATE EXTENSION "pgAutomator";
SELECT job_category_id, category
FROM pgautomator.job_category
ORDER BY job_category_id;
```

The script creates the fixed `pgautomator` schema, scheduling enums, agent and SMTP configuration, jobs, steps, schedules, connection details, and execution logs. The separate pgAutomator agent is required to poll and execute jobs; installing this extension alone does not run a scheduler.

Version 0.0.1 is an early 2017 schema with a one-line README. It stores optional SMTP and database passwords in ordinary tables and documents unsupported schedule modes. Restrict schema access, review credential storage, and test the matching external agent before considering any deployment.
