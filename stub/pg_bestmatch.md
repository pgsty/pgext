
--------

## Usage

- repo: https://github.com/tensorchord/pg_bestmatch.rs
- benchmark: https://hazyresearch.stanford.edu/blog/2024-05-20-m2-bert-retrieval


### How does it work?

- Create an BM25 statistics based on your document set by `bm25_create(table_name, column_name, statistic_name);`. It will create a materilized view to record the stats.
- Generate document sparse vector by `bm25_document_to_svector(statistic_name, passage)`
- For query, generate query sparse vector `bm25_query_to_svector(statistic_name, query)`
- Calculate the score by dot product between the query sparse vector and the document sparse vector
- Currently we use huggingface tokenizer with `bert-base-uncased` vocabulary set to tokenize words. Might support more configuration on tokenizer in the future.


### Install

```sql
CREATE EXTENSION pg_bestmatch;
SET search_path TO public, bm_catalog;
```



--------

## Example

Here is an example workflow demonstrating the usage of this extension with the example of [Stanford LoCo benchmark](https://hazyresearch.stanford.edu/blog/2024-05-20-m2-bert-retrieval).

0. Load the dataset. Here is a script for you if you want to experience `pg_bestmatch` with the dataset.

```sh
wget https://huggingface.co/api/datasets/hazyresearch/LoCoV1-Documents/parquet/default/test/0.parquet -O documents.parquet
wget https://huggingface.co/api/datasets/hazyresearch/LoCoV1-Queries/parquet/default/test/0.parquet -O queries.parquet
```

```python
import pandas as pd
from sqlalchemy import create_engine
import numpy as np
from psycopg2.extensions import register_adapter, AsIs

def adapter_numpy_float64(numpy_float64):
    return AsIs(numpy_float64)

def adapter_numpy_int64(numpy_int64):
    return AsIs(numpy_int64)

def adapter_numpy_float32(numpy_float32):
    return AsIs(numpy_float32)

def adapter_numpy_int32(numpy_int32):
    return AsIs(numpy_int32)

def adapter_numpy_array(numpy_array):
    return AsIs(tuple(numpy_array))

register_adapter(np.float64, adapter_numpy_float64)
register_adapter(np.int64, adapter_numpy_int64)
register_adapter(np.float32, adapter_numpy_float32)
register_adapter(np.int32, adapter_numpy_int32)
register_adapter(np.ndarray, adapter_numpy_array)

db_url = "postgresql://localhost:5432/pg_bestmatch_test"
engine = create_engine(db_url)

def load_documents():
    df = pd.read_parquet("documents.parquet")
    df.to_sql("documents", engine, if_exists='replace', index=False)

def load_queries():
    df = pd.read_parquet("queries.parquet")
    df['answer_pids'] = df['answer_pids'].apply(lambda x: str(x[0]))    
    df.to_sql("queries", engine, if_exists='replace', index=False)

load_documents()
load_queries()
```

1. Create BM25 statistics for the `documents` table.

```sql
SELECT bm25_create('documents', 'passage', 'documents_passage_bm25', 0.75, 1.2);
```

2. Add an embedding column to the `documents` and `queries` tables and update the embeddings for documents and queries.

```sql
ALTER TABLE documents ADD COLUMN embedding svector; -- for pgvecto.rs users
ALTER TABLE documents ADD COLUMN embedding sparsevec; -- for pgvector users

UPDATE documents SET embedding = bm25_document_to_svector('documents_passage_bm25', passage)::svector; -- for pgvecto.rs users
UPDATE documents SET embedding = bm25_document_to_svector('documents_passage_bm25', passage, 'pgvector')::sparsevec; -- for pgvector users
```

3. (Optional) Create a vector index on the sparse vector column.

```sql
CREATE INDEX ON documents USING vectors (embedding svector_dot_ops); -- for pgvecto.rs users
CREATE INDEX ON documents USING ivfflat (embedding sparsevec_ip_ops); -- for pgvector users
```

4. Perform a vector search to find the most relevant documents for each query.

```sql
ALTER TABLE queries ADD COLUMN embedding svector; -- for pgvecto.rs users
ALTER TABLE queries ADD COLUMN embedding sparsevec; -- for pgvector users

UPDATE queries SET embedding = bm25_query_to_svector('documents_passage_bm25', query)::svector; -- for pgvecto.rs users
UPDATE queries SET embedding = bm25_query_to_svector('documents_passage_bm25', query, 'pgvector')::sparsevec; -- for pgvector users

SELECT sum((array[answer_pids] = array(SELECT pid FROM documents WHERE queries.dataset = documents.dataset ORDER BY queries.embedding <#> documents.embedding LIMIT 1))::int) FROM queries;
```

This workflow showcases how to leverage BM25 text queries and vector search in PostgreSQL using this extension. The Top 1 recall of BM25 on this dataset is `0.77`. If you reproduce the result, your operations are correct.


--------

## Comparison with pg_search

- `pg_bestmatch.rs` only provides methods for generating sparse vectors and does not support index-based search (which can be achieved by pgvecto.rs or pgvector).
- `pg_search` performs BM25 retrieval via the external `tantivy` engine, which may have limitations when combined with transactions, filters, or JOIN operations. Since `pg_bestmatch.rs` is entirely native to Postgres, it offers full compatibility with these operations inside postgres.


--------

## Reference

- `tokenize`
    - Description: Tokenizes an input string into individual tokens.
    - Example:
      ```sql
      SELECT tokenize('i have an apple'); -- result: {i,have,an,apple}
      ```
- `bm25_create`
    - Description: Creates BM25 statistics for a specified table and column.
    - Usage:
      ```sql
      SELECT bm25_create('documents', 'passage', 'documents_passage_bm25');
      ```
    - Parameters:
        - `table_name`: Name of the table.
        - `column_name`: Name of the column.
        - `stat_name`: Name of the BM25 statistics.
        - `b`: BM25 parameter (default 0.75).
        - `k`: BM25 parameter (default 1.2).
- `bm25_refresh`
    - Description: Updates the BM25 statistics to reflect any changes in the underlying data.
    - Usage:
      ```sql
      SELECT bm25_refresh('documents_passage_bm25');
      ```
    - Parameters:
        - `stat_name`: Name of the BM25 statistics to update.
- `bm25_drop`
    - Description: Deletes the BM25 statistics for a specified table and column.
    - Usage:
      ```sql
      SELECT bm25_drop('documents_passage_bm25');
      ```
    - Parameters:
        - `stat_name`: Name of the BM25 statistics to delete.
- `bm25_document_to_svector`
    - Description: Converts document text into a sparse vector representation.
    - Usage:
      ```sql
      SELECT bm25_document_to_svector('documents_passage_bm25', 'document_text');
      ```
    - Parameters:
        - `stat_name`: Name of the BM25 statistics.
        - `document_text`: The text of the document.
        - `style`: Emits `pgvecto.rs`-style sparse vector or `pgvector`-style sparse vector.
- `bm25_query_to_svector`
    - Description: Converts query text into a sparse vector representation.
    - Usage:
      ```sql
      SELECT bm25_query_to_svector('documents_passage_bm25', 'We begin, as always, with the text.');
      ```
    - Parameters:
        - `stat_name`: Name of the BM25 statistics.
        - `query_text`: The text of the query.
        - `style`: Emits `pgvecto.rs`-style sparse vector or `pgvector`-style sparse vector.