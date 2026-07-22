## Usage

Sources:

- [Official README](https://github.com/pierreforstmann/pg_log_statements/blob/4ee44c244dc1fafe72d3f7060e8fe6970cc1f546/README.md)
- [Extension control file](https://github.com/pierreforstmann/pg_log_statements/blob/4ee44c244dc1fafe72d3f7060e8fe6970cc1f546/pg_log_statements.control)
- [Extension SQL for 0.0.2](https://github.com/pierreforstmann/pg_log_statements/blob/4ee44c244dc1fafe72d3f7060e8fe6970cc1f546/pg_log_statements--0.0.2.sql)
- [C implementation](https://github.com/pierreforstmann/pg_log_statements/blob/4ee44c244dc1fafe72d3f7060e8fe6970cc1f546/pg_log_statements.c)

`pg_log_statements` selectively forces `log_statement=all` for named server processes or for newly authenticated sessions matching a filter. It changes server-log behavior; it does not provide a queryable statement-history table.

### Core Workflow

Preload the library, restart PostgreSQL, install the SQL functions, and immediately restrict their execution to an audit administrator.

```conf
shared_preload_libraries = 'pg_log_statements'
```

```sql
CREATE EXTENSION pg_log_statements;

REVOKE ALL ON FUNCTION pgls_start(integer) FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_stop(integer) FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_start_filter(cstring, cstring) FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_stop_filter(cstring, cstring) FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_state() FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_conf() FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_start_debug() FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_stop_debug() FROM PUBLIC;

SELECT pgls_start(12345);
SELECT * FROM pgls_state() AS s(pid integer, status text);
SELECT pgls_stop(12345);

SELECT pgls_start_filter('application_name', 'billing-worker');
SELECT * FROM pgls_conf() AS c(filter_type text, filter_value text);
SELECT pgls_stop_filter('application_name', 'billing-worker');
```

PID mode uses `pgls_start(integer)` and `pgls_stop(integer)`. Filter mode accepts `application_name`, `user_name`, `hostname`, `ip_address`, or `database_name` and affects only sessions created after the filter is installed. `pgls_start_debug()` and `pgls_stop_debug()` toggle authentication-field logging. Because `pgls_state()` and `pgls_conf()` return anonymous records, callers must supply column definitions.

### Security and Logging Caveats

The extension functions contain no privilege check and receive public execute permission unless revoked. An untrusted user could enable statement logging for other sessions, expose SQL text and embedded secrets to server logs, and generate substantial log volume. Grant execution only to an audited administrator and secure log storage, retention, and access.

PID changes take effect when the target backend next reaches a hooked executor or utility event. Enabling forces `log_statement` to `all`; disabling forces it to `none` rather than restoring a previous session value. Filters are stored in shared memory, are limited in count and value length by the implementation, and disappear on restart. The cataloged extension version is 0.0.2 even though newer repository metadata mentions 0.0.4; match behavior and SQL to the installed package. Upstream reports validation through PostgreSQL 16, so separately test later majors.
