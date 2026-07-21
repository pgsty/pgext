## Usage

Sources:

- [Official v2.0.1 README](https://github.com/jaiminpan/pg_jieba/blob/v2.0.1/README.md)
- [Extension control file](https://github.com/jaiminpan/pg_jieba/blob/v2.0.1/pg_jieba.control)
- [SQL parser and configuration definitions](https://github.com/jaiminpan/pg_jieba/blob/v2.0.1/pg_jieba.sql)

`pg_jieba` adds Jieba-based Chinese word segmentation to PostgreSQL full-text search. The upstream `v2.0.1` source release installs SQL extension version `1.1.0`, as recorded by its control file. It provides separate document and query parsers plus ready-to-use text-search configurations.

### Core Workflow

```sql
CREATE EXTENSION pg_jieba;

SELECT to_tsvector(
    'jiebacfg',
    '小明硕士毕业于中国科学院计算所，后在日本京都大学深造'
);

SELECT plainto_tsquery('jiebaqry', '云计算专家');
```

Use `jiebacfg` to build searchable document vectors and `jiebaqry` to segment user queries:

```sql
ALTER TABLE articles
ADD COLUMN search_vector tsvector
GENERATED ALWAYS AS (to_tsvector('jiebacfg', body)) STORED;

CREATE INDEX articles_search_idx
ON articles USING GIN (search_vector);

SELECT title
FROM articles
WHERE search_vector @@ plainto_tsquery('jiebaqry', '中文全文检索');
```

### Object Index

- `jieba`: document text-search parser.
- `jiebaqry`: query-oriented text-search parser.
- `jiebacfg`: document text-search configuration using `jieba` and `jieba_stem`.
- `jiebaqry`: text-search configuration of the same name using the query parser.
- `jieba_stem`: simple dictionary with Jieba stop words used for the parser's token categories.

### Custom Dictionary and Caveats

Upstream reads a custom dictionary named `jieba.user.dict.utf8` from PostgreSQL's `tsearch_data` directory. Entries may contain a word and optional part-of-speech tag:

```text
云计算
韩玉鉴赏
蓝翔 nz
```

- The v2.x source requires a C++11-capable compiler because of its bundled `cppjieba` dependency.
- Upstream's published compatibility testing is old and limited. Build and regression-test the package against the exact PostgreSQL major version used in production.
- Changing dictionaries changes tokenization. Recompute stored `tsvector` values and rebuild dependent indexes when dictionary output changes.
