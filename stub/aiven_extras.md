## Usage

Sources:

- [Aiven Extras README at the reviewed commit](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/README.md)
- [Control template at the reviewed commit](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/aiven_extras.control.in)
- [Build version at the reviewed commit](https://github.com/aiven/aiven-extras/blob/ed2a4a676034b1ca3edb767d5805c482262d3f9b/Makefile)
- [Aiven Extras tags](https://github.com/aiven/aiven-extras/tags)

`aiven_extras` is an Aiven-maintained extension that lets non-superusers perform a selected set of otherwise privileged operations after a privileged operator installs it. Its main use cases are logical-replication administration, session-local `auto_explain` configuration, selected pgaudit settings, and controlled dblink execution.

The replication wrappers include `pg_create_publication`, `pg_create_subscription`, and subscription enable, disable, refresh, list, and drop operations. The calling role still needs the `REPLICATION` privilege; the extension does not turn an arbitrary role into a superuser.

### Basic Setup

```sql
CREATE EXTENSION aiven_extras;

SELECT aiven_extras.auto_explain_load();
SELECT aiven_extras.set_auto_explain_log_min_duration('2000');
SELECT aiven_extras.set_auto_explain_log_analyze('on');
```

The example loads and configures `auto_explain` only for the current database session. Publication and subscription functions should be called with fully qualified table names and carefully scoped connection strings.

### Caveats

- A privileged operator must install the extension first. Logical-replication operations still require `REPLICATION` on the calling role.
- The extension deliberately exposes only selected operations; it is not a general privilege-bypass mechanism.
- The current `main` Makefile identifies version `1.1.20`, matching the catalog, while the latest visible upstream tag is `1.1.15`. Pin the source commit or provider package whose behavior you have tested.
- Subscription connection strings can contain credentials. Keep them out of logs and query history and prefer provider-supported secret handling.
