## Usage

Sources:

- [Official documentation](https://github.com/lgrosz/pg_climb/blob/8dad0d7239e0cb6761a729f5be45e6698ecae733/README.md)
- [Extension control file](https://github.com/lgrosz/pg_climb/blob/8dad0d7239e0cb6761a729f5be45e6698ecae733/pg_climb.control)
- [Official repository](https://github.com/lgrosz/pg_climb)

`pg_climb` Rock-climbing grade conversion and utility functions.

### Enablement

Install the files for the intended server, then create `pg_climb` in the target database:

```sql
CREATE EXTENSION pg_climb;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE TABLE routes (grade grade);
INSERT INTO routes VALUES ('V5'::grade), ('F7A+'::grade);

SELECT grade, grade_type(grade)
FROM routes
ORDER BY grade;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `btree_grade_ops` | OPERATOR CLASS | Operator class used by indexes. |
| `grade` | TYPE | User-facing data type created by the extension. |
| `grade_cmp` | FUNCTION | Callable function from the reviewed install surface. |
| `grade_eq` | FUNCTION | Callable function from the reviewed install surface. |
| `grade_ge` | FUNCTION | Callable function from the reviewed install surface. |
| `grade_gt` | FUNCTION | Callable function from the reviewed install surface. |
| `grade_in` | FUNCTION | Callable function from the reviewed install surface. |
| `grade_le` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 17; do not infer unlisted majors.
