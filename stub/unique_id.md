## Usage

Sources:

- [Official documentation](https://github.com/fabriziomello/unique_id/blob/95ac200ada23072cb0f4ad05741fd92c24c0470b/README.md)
- [Extension control file](https://github.com/fabriziomello/unique_id/blob/95ac200ada23072cb0f4ad05741fd92c24c0470b/unique_id.control)
- [Official repository](https://github.com/fabriziomello/unique_id)

`unique_id` Time-based K-sorted unique identifier generators inspired by Instagram and Sonyflake designs.

### Enablement

Install the files for the intended server, then create `unique_id` in the target database:

```sql
CREATE EXTENSION unique_id;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
fabrizio=# CREATE EXTENSION unique_id;
CREATE EXTENSION
fabrizio=# \dx unique_id
                                     List of installed extensions
   Name    | Version | Schema |                              Description
-----------+---------+--------+------------------------------------------------------------------------
 unique_id | 1.0     | public | Non-Standard Time-based K-sorted, Lexicographically Unique Identifiers
(1 row)

fabrizio=# CREATE SEQUENCE instagram_seq;
CREATE SEQUENCE
fabrizio=# -- Using default shard_id = 0
fabrizio=# SELECT unique_id_instagram('instagram_seq');
 unique_id_instagram
---------------------
 2563729919292997633
(1 row)

fabrizio=# -- Using default shard_id = 1
fabrizio=# SELECT unique_id_instagram('instagram_seq', 1);
 unique_id_instagram
---------------------
 2563729919292998658
(1 row)
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `unique_id_instagram` | FUNCTION | Callable function from the reviewed install surface. |
| `unique_id_sonyflake` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
