

## Usage

> [GitHub: tensorchord/pg_tokenizer.rs](https://github.com/tensorchord/pg_tokenizer.rs)

`pg_tokenizer` is a PostgreSQL extension that provides tokenizers for full-text search. It is designed to work with [VectorChord-bm25](https://github.com/tensorchord/VectorChord-bm25) for native BM25 ranking index support.

## Quick Start

```sql
CREATE EXTENSION pg_tokenizer;

-- Create a tokenizer using the LLMLingua2 model
SELECT create_tokenizer('tokenizer1', $$
model = "llmlingua2"
$$);

-- Tokenize text
SELECT tokenize('PostgreSQL is a powerful, open-source object-relational database system. It has over 15 years of active development.', 'tokenizer1');
```

## Tokenizer Models

pg_tokenizer supports multiple tokenizer models for different languages and use cases:

| Model | Language | Description |
|-------|----------|-------------|
| `llmlingua2` | English | BERT-based tokenizer from LLMLingua2 |
| `jieba` | Chinese | Jieba Chinese text segmentation |
| `lindera/ipadic` | Japanese | Lindera tokenizer with IPADIC dictionary |
| Custom models | Any | User-trained models for domain-specific text |

### Creating Tokenizers

```sql
-- English tokenizer
SELECT create_tokenizer('en_tokenizer', $$
model = "llmlingua2"
$$);

-- Chinese tokenizer
SELECT create_tokenizer('zh_tokenizer', $$
model = "jieba"
$$);

-- Japanese tokenizer
SELECT create_tokenizer('ja_tokenizer', $$
model = "lindera/ipadic"
$$);
```

### Tokenizing Text

```sql
-- Tokenize English text
SELECT tokenize('full text search in PostgreSQL', 'en_tokenizer');

-- Tokenize Chinese text
SELECT tokenize('PostgreSQL是一个强大的数据库系统', 'zh_tokenizer');
```

## Text Analyzer

pg_tokenizer also provides text analyzer functionality that combines tokenization with additional text processing steps. For detailed text analyzer usage, refer to the [Text Analyzer documentation](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/05-text-analyzer.md).

## Integration with VectorChord-BM25

pg_tokenizer is typically used together with VectorChord-BM25 for full BM25 ranking support:

```sql
CREATE EXTENSION IF NOT EXISTS pg_tokenizer CASCADE;
CREATE EXTENSION IF NOT EXISTS vchord_bm25 CASCADE;

-- Create a tokenizer
SELECT create_tokenizer('my_tokenizer', $$
model = "llmlingua2"
$$);

-- Tokenize text into bm25vectors for indexing and search
SELECT tokenize('your search query', 'my_tokenizer');
```

## Documentation

For more details, see the full documentation:

- [Installation](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/01-installation.md)
- [Examples](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/03-examples.md)
- [Usage](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/04-usage.md)
- [Text Analyzer](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/05-text-analyzer.md)
- [Model Reference](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/06-model.md)
- [Limitations](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/07-limitation.md)
