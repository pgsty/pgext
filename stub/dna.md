## Usage

Sources:

- [Official documentation](https://github.com/sid2364/dna-sequences-pg-extension/blob/dc6e4f105cbc8bdcacab8ca0ba1f3ae005c42efb/README.md)
- [Extension control file](https://github.com/sid2364/dna-sequences-pg-extension/blob/dc6e4f105cbc8bdcacab8ca0ba1f3ae005c42efb/dna.control)
- [Official repository](https://github.com/sid2364/dna-sequences-pg-extension)

`dna` Compact DNA sequence and k-mer types with comparison, validation, and analysis functions.

### Enablement

Install the files for the intended server, then create `dna` in the target database:

```sql
CREATE EXTENSION dna;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT k.kmer FROM generate_kmers('ACGTACGCACGT', 6) AS k(kmer) WHERE 'DNMSRN' @> k.kmer ;
--  kmer
----------
-- GTACGC
-- GCACGT
--(2 rows)
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `kmer` | TYPE | User-facing data type created by the extension. |
| `generate_kmers` | FUNCTION | Callable function from the reviewed install surface. |
| `length` | FUNCTION | Callable function from the reviewed install surface. |
| `contains` | FUNCTION | Callable function from the reviewed install surface. |
| `qkmer` | TYPE | User-facing data type created by the extension. |
| `starts_with` | FUNCTION | Callable function from the reviewed install surface. |
| `dna_construct` | FUNCTION | Callable function from the reviewed install surface. |
| `dna_in` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
