## Usage

Sources:

- [vector_recall.control](https://github.com/AlibabaIncubator/gpdb-faiss-vector/blob/f247fd9653361316c2e9fefed4b767d7811c75d5/vector_recall.control)
- [README.md](https://github.com/AlibabaIncubator/gpdb-faiss-vector/blob/f247fd9653361316c2e9fefed4b767d7811c75d5/README.md)
- [vector_recall--0.1.0.sql](https://github.com/AlibabaIncubator/gpdb-faiss-vector/blob/f247fd9653361316c2e9fefed4b767d7811c75d5/vector_recall--0.1.0.sql)
- [makefile](https://github.com/AlibabaIncubator/gpdb-faiss-vector/blob/f247fd9653361316c2e9fefed4b767d7811c75d5/makefile)

`vector_recall` integrates Faiss into Greenplum. Serialized indexes are stored as BYTEA; SQL functions create, train, add vectors and perform top-k or range searches. This is a Greenplum extension, with no stock PostgreSQL compatibility claim.

### Core Workflow

```sql
CREATE EXTENSION vector_recall;
SELECT faiss_index_create(3, 'Flat', 1);
```

### Operational Boundaries

The shared library and Faiss runtime must match the Greenplum deployment. Installation is privileged; the control does not require preload. Training and insertion return an updated serialized index, which callers must store. Search can cache deserialized indexes by a caller-supplied key: use keys that track index revisions. Cache reset/prune/charge functions and a top-k merge helper are provided. Serialized index input and server resources need to be trusted and bounded; upstream benchmarks do not establish workload guarantees.
