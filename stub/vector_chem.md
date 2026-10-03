## Usage

Sources:

- [vector_chem.control](https://github.com/leotaku/pgvector_chem/blob/d87e9c8688faca1afbf825095107452b3b0e52cb/vector_chem.control)
- [README.md](https://github.com/leotaku/pgvector_chem/blob/d87e9c8688faca1afbf825095107452b3b0e52cb/README.md)
- [sql/vector_chem.sql](https://github.com/leotaku/pgvector_chem/blob/d87e9c8688faca1afbf825095107452b3b0e52cb/sql/vector_chem.sql)
- [Makefile](https://github.com/leotaku/pgvector_chem/blob/d87e9c8688faca1afbf825095107452b3b0e52cb/Makefile)

`vector_chem` adds Tanimoto/Jaccard distance to the pgvector type for dichotomous values. It exposes the `<^>` operator and `tanimoto_distance` function.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION vector_chem;
SELECT tanimoto_distance('[1,0,1]'::vector, '[1,0,0]'::vector);
```

### Operational Boundaries

Install the matching pgvector library and enable `vector` first; the SQL uses its type even though the control omits the dependency. Installation requires superuser privileges. No preload is documented. Continuous-valued vectors and approximate-nearest-neighbor indexing are not supported by this extension; it is a scalar distance implementation.
