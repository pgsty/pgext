## Usage

Sources:

- [Official documentation](https://github.com/verticalbarHQ/pg_ocpm/blob/bb025c78b4eb45c7a007b8868aedb43f7a4d5cf7/README.md)
- [Extension control file](https://github.com/verticalbarHQ/pg_ocpm/blob/bb025c78b4eb45c7a007b8868aedb43f7a4d5cf7/pg_ocpm.control)
- [Official repository](https://github.com/verticalbarHQ/pg_ocpm)

`pg_ocpm` Object-centric process-mining storage, traversal, aggregation, and compact analytical capsules.

### Enablement

Install the files for the intended server, then create `pg_ocpm` in the target database:

```sql
CREATE EXTENSION pg_ocpm;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE EXTENSION pg_ocpm;
SELECT ocpm.version();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `ocpm.version` | FUNCTION | Callable function from the reviewed install surface. |
| `ocpm.dataset` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `ocpm.dataset_id` | FUNCTION | Callable function from the reviewed install surface. |
| `ocpm.case_bucket` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `ocpm.rebuild_binding_index` | FUNCTION | Callable function from the reviewed install surface. |
| `ocpm.activity_profile` | FUNCTION | Callable function from the reviewed install surface. |
| `ocpm.adjacency_links` | FUNCTION | Callable function from the reviewed install surface. |
| `ocpm.adjacency_neighborhood` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 13, 14, 15, 16, 17, 18; do not infer unlisted majors.
- The extension fixes or creates schema objects under `ocpm`; include them in privilege and backup review.
