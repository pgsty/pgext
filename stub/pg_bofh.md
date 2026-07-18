## Usage

Sources:

- [Upstream README](https://github.com/rdunklau/pg_bofh/blob/master/README.md)
- [Extension control file](https://github.com/rdunklau/pg_bofh/blob/master/pg_bofh.control)

`pg_bofh` is a small planner-hook example that rejects queries without a `WHERE` clause. Its upstream version is `0.0.1`.

This is a headless module: the repository contains no versioned extension SQL, so there is no SQL activation step. Build and install the shared library, add it to the server configuration, and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_bofh'
```

Once loaded, statements without a `WHERE` clause raise an error instead of executing. That blanket rule affects more than accidental full-table changes, so treat the project as a planner-hook demonstration rather than a complete production authorization or query-safety system. Test application and maintenance workloads before enabling it on a cluster.
