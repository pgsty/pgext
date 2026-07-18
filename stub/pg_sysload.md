## Usage

Sources:

- [Pinned Rust implementation](https://github.com/boringSQL/pg_sysload/blob/7c4db3cef4df26dd12718d43c83a75f1f0b5508b/src/lib.rs)
- [Pinned Cargo metadata](https://github.com/boringSQL/pg_sysload/blob/7c4db3cef4df26dd12718d43c83a75f1f0b5508b/Cargo.toml)
- [Pinned extension control file](https://github.com/boringSQL/pg_sysload/blob/7c4db3cef4df26dd12718d43c83a75f1f0b5508b/pg_sysload.control)

pg_sysload version 0.0.0 exposes the first three fields of Linux procfs load average as a floating-point array. They represent host runnable-load averages over approximately one, five, and fifteen minutes.

### Query host load

```sql
CREATE EXTENSION pg_sysload;

SELECT sys_loadavg();
```

No preload setting is required. Loading the library fails if the backend cannot find /proc/loadavg.

### Interpretation and portability

Load average is host or container-namespace state, not PostgreSQL workload. It includes runnable and uninterruptible tasks from other processes and is not divided by CPU count. Record CPU allocation, cgroup boundaries, container procfs behavior, and sampling time before comparing values across servers. Use PostgreSQL and operating-system monitoring together to attribute load.

The function reads and parses /proc/loadavg on every SQL call. If an individual number cannot be parsed, the implementation silently drops that element, so callers should validate that the returned array has exactly three finite values. Procfs read failures raise an error in the calling backend.

The project has no README, user documentation, SQL tests, or compatibility policy. Its Cargo features cover PostgreSQL 11 through 16 using pgrx 0.11.4, not PostgreSQL 17 or 18. The repository last changed in 2024. Restrict execution if host telemetry is sensitive, rate-limit collection, and prefer a maintained node exporter or platform agent for production monitoring.
