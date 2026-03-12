

## 用法

> [GitHub: tensorchord/VectorChord-bm25](https://github.com/tensorchord/VectorChord-bm25)

VectorChord-BM25 是一个实现 BM25 排序算法的 PostgreSQL 扩展，基于 Block-WeakAnd 算法。它设计与 [pg_tokenizer](https://github.com/tensorchord/pg_tokenizer.rs) 配合使用，支持自定义文本分词。

## 架构

该扩展由三个主要组件组成：

1. **分词器**：将文本转换为 `bm25vector`（存储词汇 ID 和词频的稀疏向量）
2. **bm25vector**：用于存储分词后文本的自定义数据类型
3. **bm25vector 索引**：加速搜索和排序操作

## 快速开始

```sql
-- 启用所需扩展
CREATE EXTENSION IF NOT EXISTS pg_tokenizer CASCADE;
CREATE EXTENSION IF NOT EXISTS vchord_bm25 CASCADE;

-- 创建分词器（如用于英文的 LLMLingua2）
SELECT create_tokenizer('tokenizer1', $$
model = "llmlingua2"
$$);

-- 创建包含文本内容的表
CREATE TABLE documents (
  id SERIAL PRIMARY KEY,
  passage TEXT,
  embedding bm25vector
);

-- 将文本段落分词为 bm25vector
UPDATE documents SET embedding = tokenize(passage, 'tokenizer1');

-- 创建 BM25 索引
CREATE INDEX documents_embedding_bm25 ON documents USING bm25 (embedding bm25_ops);

-- 使用 BM25 排序查询
SELECT id, passage, embedding <&> to_bm25query('documents_embedding_bm25', tokenize('search query', 'tokenizer1')) AS score
FROM documents
ORDER BY score
LIMIT 10;
```

**注意**：VectorChord-BM25 中的 BM25 分数为负数，越负表示相关性越高。

## `<&>` 运算符

`<&>` 运算符计算存储的 `bm25vector` 与查询 `bm25vector` 之间的 BM25 相关性分数。查询必须用 `to_bm25query()` 包装，它接受索引名称和分词后的查询：

```sql
-- 基本搜索查询
-- to_bm25query(索引名称, 分词后的查询)
SELECT id, passage, embedding <&> to_bm25query('documents_embedding_bm25', tokenize('database system', 'tokenizer1')) AS score
FROM documents
ORDER BY score
LIMIT 10;
```

## 语言支持

VectorChord-BM25 通过不同的分词器配置支持多种语言：

| 语言 | 方式 | 模型/预分词器 |
|------|------|---------------|
| 英语 | 预训练模型 | `model = "llmlingua2"` 或 `model = "bert_base_uncased"` |
| 中文 | 带结巴预分词器的自定义模型 | `[pre_tokenizer.jieba]` |
| 日语 | 带 Lindera 预分词器的自定义模型 | Lindera + IPADIC 词典 |
| 自定义 | 通过文本分析器训练的用户模型 | `create_custom_model_tokenizer_and_trigger()` |

### 中文文本搜索示例

中文文本需要带结巴预分词器的自定义模型（而非预训练模型）：

```sql
-- 创建带结巴预分词器的文本分析器
SELECT create_text_analyzer('zh_text_analyzer', $$
[pre_tokenizer.jieba]
$$);

-- 创建在语料上训练的自定义模型分词器
SELECT create_custom_model_tokenizer_and_trigger(
    tokenizer_name => 'zh_tokenizer',
    model_name => 'zh_model',
    text_analyzer_name => 'zh_text_analyzer',
    table_name => 'documents',
    source_column => 'passage',
    target_column => 'embedding'
);
```

### 自定义分词器模型

对于领域特定术语，你可以创建带停用词、词干提取和其他过滤器的文本分析器，然后使用 `create_custom_model_tokenizer_and_trigger()` 在语料上训练自定义模型。

## 与替代方案的比较

| 特性 | VectorChord-BM25 | PostgreSQL tsvector + ts_rank |
|------|-------------------|-------------------------------|
| 排序算法 | BM25 | tf-idf 变体 |
| 自定义分词器 | 支持（通过 pg_tokenizer） | 仅限内置配置 |
| 索引类型 | 专用 BM25 索引 | GIN 索引 |
| 原生 PostgreSQL | 是（扩展） | 内置 |
| 语言支持 | 通过模型可扩展 | 通过文本搜索配置 |
