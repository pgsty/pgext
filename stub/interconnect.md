## Usage

Sources:

- [Official upstream README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/contrib/interconnect/README.md)
- [Official extension control file (interconnect.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/contrib/interconnect/interconnect.control)
- [Official extension SQL (interconnect--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/contrib/interconnect/interconnect--1.0.sql)

`interconnect` — This extension provides cumulative interconnect statistics for Apache Cloudberry, including queue sizes, buffer usage, retransmits, packet errors, and other UDPIFC‑related metrics. Use it for the corresponding analytical or storage workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION interconnect;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gp_interconnect_stats` is an extension-defined view.
- `gp_interconnect_stats_per_host` is an extension-defined view.
- `gp_interconnect_stats_per_segment` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
