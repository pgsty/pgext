## Usage

Sources:

- [Current project README and architecture boundary](https://github.com/jasonmassie01/pg_sage/blob/492cd7b3ad840be3d8fd6ad1005a029f2396ecd7/README.md)
- [Version 0.5.0 extension SQL](https://github.com/jasonmassie01/pg_sage/blob/492cd7b3ad840be3d8fd6ad1005a029f2396ecd7/sql/pg_sage--0.5.0.sql)
- [Legacy extension entry point](https://github.com/jasonmassie01/pg_sage/blob/492cd7b3ad840be3d8fd6ad1005a029f2396ecd7/src/pg_sage.c)
- [Current reverse specification](https://github.com/jasonmassie01/pg_sage/blob/492cd7b3ad840be3d8fd6ad1005a029f2396ecd7/docs/REVERSE_SPEC.md)

pg_sage 0.5.0 is the cataloged, preload-based C extension from the pg_sage repository. It collects snapshots, analyzes database health, records findings and actions, and exposes control and JSON helper functions in the fixed `sage` schema. The current upstream product has since moved to a standalone Go sidecar and explicitly says no PostgreSQL extension is required; do not mix the two architectures.

### Core Workflow

For an existing 0.5.0 extension deployment, preload both its library and `pg_stat_statements`, restart PostgreSQL, then create the extension in the target database:

```ini
shared_preload_libraries = 'pg_stat_statements,pg_sage'
pg_sage.database = 'postgres'
```

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_sage;

SELECT sage.status();
SELECT id, severity, title, recommendation
FROM sage.findings
WHERE status = 'open'
ORDER BY last_seen DESC;
```

Use observation or recommendation behavior first, review every proposed SQL action, and verify emergency-stop handling before any autonomous mode.

### Important Objects

- `sage.snapshots`, `sage.findings`, and `sage.explain_cache` store collected evidence and analysis results.
- `sage.action_log` records executed SQL and rollback metadata.
- `sage.config` stores runtime configuration, including endpoints and autonomous mode.
- `sage.status`, `sage.emergency_stop`, and `sage.resume` expose operational control.
- `sage.explain`, `sage.suppress`, `sage.diagnose`, and `sage.briefing` expose analysis workflows.
- `sage.health_json`, `sage.findings_json`, `sage.schema_json`, and related helpers support the sidecar-facing JSON surface.

### Security and Lifecycle Notes

The legacy library installs executor hooks, shared memory, libcurl, and four background workers, and it requires superuser installation. It can collect query text and plans, contact an LLM endpoint, retain API-key-related configuration, and execute recommended SQL, so limit schema access and protect configuration and logs as secrets. The current README and reverse specification describe a wire-protocol sidecar instead of this extension. Treat 0.5.0 as a legacy deployment boundary and migrate deliberately rather than applying current sidecar instructions to the C module.
