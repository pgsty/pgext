## Usage

Sources:

- [README.md](https://github.com/dnewcome/pg-bitemporal-ext/blob/60ec92d9a8daee5dd0ea3339f0fbc067cafc672f/README.md)
- [sql/pg_bitemporal--0.1.sql](https://github.com/dnewcome/pg-bitemporal-ext/blob/60ec92d9a8daee5dd0ea3339f0fbc067cafc672f/sql/pg_bitemporal--0.1.sql)
- [pg_bitemporal.control](https://github.com/dnewcome/pg-bitemporal-ext/blob/60ec92d9a8daee5dd0ea3339f0fbc067cafc672f/pg_bitemporal.control)

`pg_bitemporal` combines interval-relation functions with triggers that retain system-time history. Full mode also stores application-supplied valid time; system-only mode tracks database changes.

### Core Workflow

```sql
CREATE EXTENSION btree_gist;
CREATE EXTENSION pg_bitemporal;
CREATE TABLE employee_history_demo (id integer NOT NULL, department text);
SELECT bitemporal.enable('employee_history_demo', ARRAY['id'], p_mode => 'system_only');
INSERT INTO employee_history_demo VALUES (1, 'Engineering');
SELECT id, department, sys_start, sys_end FROM employee_history_demo;
SELECT allen_relation('[1,4)'::int4range, '[3,8)'::int4range);
```

### Operational Boundaries

The README requires PostgreSQL 14 or later and `btree_gist`; no complete major-version test matrix is published. Native functions require administrator installation. `bitemporal.enable` alters the target table, adds temporal columns and indexes, installs exclusion constraints and a trigger, and requires the corresponding table privileges. Use the actual SQL parameter `p_mode`. Updates close the old version and insert another; deletes close the current version instead of physically removing it. Query current rows with `sys_end = 'infinity'`; historical rows remain in the table and require storage planning. Full mode adds `valid_from`, `valid_to` and `valid_range`, whose values the application supplies. `allen_relation`, `allen_relation_code` and the thirteen predicate functions operate on range values. The extension creates a fixed `bitemporal` schema even though its control declares relocatability.
