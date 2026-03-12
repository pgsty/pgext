


## Usage

> [pg_crash: Send random signals to random processes](https://github.com/cybertec-postgresql/pg_crash)

pg_crash is a chaos engineering extension that periodically sends kill signals to PostgreSQL backend processes, useful for HA and failover testing. It must be added to `shared_preload_libraries`.

### Configuration

Add to `postgresql.conf`:

```
shared_preload_libraries = 'pg_crash'

# POSIX signals to send (space-separated)
crash.signals = '1 2 3'

# Delay in seconds between sending signals
crash.delay = 30
```

### Signal Reference

Common POSIX signals: `1` (SIGHUP), `2` (SIGINT), `3` (SIGQUIT), `9` (SIGKILL), `15` (SIGTERM).

After configuring, restart the server. The background worker will periodically send the configured signals to random backend processes at the specified interval.
