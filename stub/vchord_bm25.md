

## Usage

> [GitHub: tensorchord/VectorChord-bm25](https://github.com/tensorchord/VectorChord-bm25)

VectorChord-BM25 is a PostgreSQL extension for the BM25 ranking algorithm, implemented via Block-WeakAnd algorithms. It is designed to work together with [pg_tokenizer](https://github.com/tensorchord/pg_tokenizer.rs) for customized text tokenization.

## Architecture

The extension comprises three main components:

1. **Tokenizer**: Converts text into `bm25vector` (sparse vectors storing vocabulary IDs and term frequencies)
2. **bm25vector**: A custom data type for storing tokenized text
3. **bm25vector indexes**: Accelerate search and ranking operations

## Quick Start

```sql
-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS pg_tokenizer CASCADE;
CREATE EXTENSION IF NOT EXISTS vchord_bm25 CASCADE;

-- Create a tokenizer (e.g., LLMLingua2 for English)
SELECT create_tokenizer('tokenizer1', $$
model = "llmlingua2"
$$);

-- Create a table with text content
CREATE TABLE documents (
  id SERIAL PRIMARY KEY,
  passage TEXT,
  embedding bm25vector
);

-- Tokenize text passages into bm25vectors
UPDATE documents SET embedding = tokenize(passage, 'tokenizer1');

-- Create a BM25 index
CREATE INDEX documents_embedding_bm25 ON documents USING bm25 (embedding bm25_ops);

-- Query with BM25 ranking
SELECT id, passage, embedding <&> to_bm25query('documents_embedding_bm25', tokenize('search query', 'tokenizer1')) AS score
FROM documents
ORDER BY score
LIMIT 10;
```

**Note**: BM25 scores in VectorChord-BM25 are negative, with more negative scores indicating greater relevance.

## The `<&>` Operator

The `<&>` operator computes the BM25 relevance score between a stored `bm25vector` and a query `bm25vector`. Queries must be wrapped in `to_bm25query()` which takes the index name and the tokenized query:

```sql
-- Basic search query
-- to_bm25query(index_name, tokenized_query)
SELECT id, passage, embedding <&> to_bm25query('documents_embedding_bm25', tokenize('database system', 'tokenizer1')) AS score
FROM documents
ORDER BY score
LIMIT 10;
```

## Language Support

VectorChord-BM25 supports multiple languages through different tokenizer configurations:

| Language | Approach | Model/Pre-tokenizer |
|----------|----------|---------------------|
| English | Pre-trained model | `model = "llmlingua2"` or `model = "bert_base_uncased"` |
| Chinese | Custom model with Jieba pre-tokenizer | `[pre_tokenizer.jieba]` |
| Japanese | Custom model with Lindera pre-tokenizer | Lindera with IPADIC dictionary |
| Custom | User-trained models via text analyzers | `create_custom_model_tokenizer_and_trigger()` |

### Chinese Text Search Example

Chinese text requires a custom model with a Jieba pre-tokenizer (not a pre-trained model):

```sql
-- Create a text analyzer with Jieba pre-tokenizer
SELECT create_text_analyzer('zh_text_analyzer', $$
[pre_tokenizer.jieba]
$$);

-- Create a custom model tokenizer that trains on your corpus
SELECT create_custom_model_tokenizer_and_trigger(
    tokenizer_name => 'zh_tokenizer',
    model_name => 'zh_model',
    text_analyzer_name => 'zh_text_analyzer',
    table_name => 'documents',
    source_column => 'passage',
    target_column => 'embedding'
);
```

### Custom Tokenizer Models

For domain-specific terminology, you can create text analyzers with stopwords, stemming, and other filters, then train custom models on your corpus using `create_custom_model_tokenizer_and_trigger()`.

## Comparison with Alternatives

| Feature | VectorChord-BM25 | PostgreSQL tsvector + ts_rank |
|---------|-------------------|-------------------------------|
| Ranking algorithm | BM25 | tf-idf variant |
| Custom tokenizers | Yes (via pg_tokenizer) | Limited to built-in configs |
| Index type | Dedicated BM25 index | GIN index |
| Native PostgreSQL | Yes (extension) | Built-in |
| Language support | Extensible via models | Via text search configs |
