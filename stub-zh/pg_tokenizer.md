

## 用法

> [GitHub: tensorchord/pg_tokenizer.rs](https://github.com/tensorchord/pg_tokenizer.rs)

`pg_tokenizer` 是一个为全文搜索提供分词器的 PostgreSQL 扩展。它设计与 [VectorChord-bm25](https://github.com/tensorchord/VectorChord-bm25) 配合使用，提供原生 BM25 排序索引支持。

## 快速开始

```sql
CREATE EXTENSION pg_tokenizer;

-- 使用 LLMLingua2 模型创建分词器
SELECT create_tokenizer('tokenizer1', $$
model = "llmlingua2"
$$);

-- 分词文本
SELECT tokenize('PostgreSQL is a powerful, open-source object-relational database system. It has over 15 years of active development.', 'tokenizer1');
```

## 分词器模型

pg_tokenizer 支持多种分词器模型，适用于不同语言和场景：

| 模型 | 语言 | 说明 |
|------|------|------|
| `llmlingua2` | 英语 | 基于 BERT 的 LLMLingua2 分词器 |
| `jieba` | 中文 | 结巴中文分词 |
| `lindera/ipadic` | 日语 | 带 IPADIC 词典的 Lindera 分词器 |
| 自定义模型 | 任意 | 用户训练的领域特定文本模型 |

### 创建分词器

```sql
-- 英文分词器
SELECT create_tokenizer('en_tokenizer', $$
model = "llmlingua2"
$$);

-- 中文分词器
SELECT create_tokenizer('zh_tokenizer', $$
model = "jieba"
$$);

-- 日文分词器
SELECT create_tokenizer('ja_tokenizer', $$
model = "lindera/ipadic"
$$);
```

### 文本分词

```sql
-- 分词英文文本
SELECT tokenize('full text search in PostgreSQL', 'en_tokenizer');

-- 分词中文文本
SELECT tokenize('PostgreSQL是一个强大的数据库系统', 'zh_tokenizer');
```

## 文本分析器

pg_tokenizer 还提供文本分析器功能，将分词与额外的文本处理步骤结合。详细的文本分析器用法请参见[文本分析器文档](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/05-text-analyzer.md)。

## 与 VectorChord-BM25 集成

pg_tokenizer 通常与 VectorChord-BM25 配合使用以获得完整的 BM25 排序支持：

```sql
CREATE EXTENSION IF NOT EXISTS pg_tokenizer CASCADE;
CREATE EXTENSION IF NOT EXISTS vchord_bm25 CASCADE;

-- 创建分词器
SELECT create_tokenizer('my_tokenizer', $$
model = "llmlingua2"
$$);

-- 将文本分词为 bm25vector 用于索引和搜索
SELECT tokenize('your search query', 'my_tokenizer');
```

## 文档

更多详情请参见完整文档：

- [安装](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/01-installation.md)
- [示例](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/03-examples.md)
- [用法](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/04-usage.md)
- [文本分析器](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/05-text-analyzer.md)
- [模型参考](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/06-model.md)
- [限制](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/07-limitation.md)
