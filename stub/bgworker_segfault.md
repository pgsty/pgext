## Usage

Sources:

- [Repository](https://github.com/optionfactory/bgworker_segfault)
- [Extension control file at the reviewed commit](https://github.com/optionfactory/bgworker_segfault/blob/5c239e488936efd3068c1679bc933693fed9dc42/bgworker_segfault.control)
- [Cargo package and PostgreSQL feature matrix](https://github.com/optionfactory/bgworker_segfault/blob/5c239e488936efd3068c1679bc933693fed9dc42/Cargo.toml)
- [Background-worker test source](https://github.com/optionfactory/bgworker_segfault/blob/5c239e488936efd3068c1679bc933693fed9dc42/src/lib.rs)

`bgworker_segfault` is a minimal pgrx reproduction project, not an application extension. Its repository description says it demonstrates a segfault when dynamic background-worker allocation exhausts `max_worker_processes`. The source exports a `bgworker_sleep` background-worker entry point and contains a test that allocates one worker beyond the estimated available capacity.

### Intended Use

Use this project only to reproduce or investigate that pgrx/PostgreSQL failure mode in an isolated test cluster. There is no documented end-user SQL API, and `bgworker_sleep` is a C background-worker entry point rather than a SQL function.

The control file marks the extension `superuser`, untrusted, and non-relocatable. The package is version `0.0.0`, uses pgrx `0.11.0`, and declares build features for PostgreSQL 11 through 16 with PostgreSQL 13 as the default feature.

### Caveats

- The test intentionally drives background-worker capacity to its limit and is associated with a server crash reproduction. Never run it on a production or shared cluster.
- Installing the extension does not provide an operational background-worker service; the useful behavior exists in the test harness.
- The source has no README, releases, user-facing configuration, or compatibility claims beyond its Cargo feature list. Treat it as a frozen diagnostic fixture, not a supported PostgreSQL component.
