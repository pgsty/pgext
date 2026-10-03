## 用法

来源：

- [0.1.1 control](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/pg_tokenizer.control)
- [Installation and preload](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/01-installation.md)
- [Tokenizer usage](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/04-usage.md)
- [Reference and n-grams](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/00-reference.md)
- [Models](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/06-model.md)
- [Cache limitations](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/07-limitation.md)

`pg_tokenizer` 为搜索应用将文本转换成词元 ID，常与 VectorChord-BM25 配合使用。分词器由文本分析器和词汇模型组成。扩展要求共享预加载，并将 SQL 对象安装在固定的 `tokenizer_catalog` 模式中。

### 启用与分词

将库加入已有预加载列表并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_tokenizer'
```

```sql
CREATE EXTENSION pg_tokenizer;
SET search_path = public, tokenizer_catalog;

SELECT create_tokenizer('english', $$
model = "llmlingua2"
$$);
SELECT tokenize('PostgreSQL full text search', 'english');
```

`tokenize(text, text)` 返回整数词元 ID，不是相关性分数，也不是每个词元一行。安装 BM25 伴随扩展后，可以将返回数组转换为其稀疏向量类型。

### 分析中文文本

Jieba 是 **预分词器**，不是名为 jieba 的内置词汇模型。应为语料创建文本分析器和自定义模型：

```sql
CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    passage text,
    token_ids integer[]
);
SELECT create_text_analyzer('chinese', $$
[pre_tokenizer.jieba]
$$);
SELECT create_custom_model_tokenizer_and_trigger(
    tokenizer_name => 'zh_tokenizer',
    model_name => 'zh_model',
    text_analyzer_name => 'chinese',
    table_name => 'documents',
    source_column => 'passage',
    target_column => 'token_ids'
);
INSERT INTO documents(passage) VALUES ('PostgreSQL全文检索');
SELECT tokenize('数据库', 'zh_tokenizer');
```

辅助函数从源列学习词汇，并创建触发器维护词元 ID。文档与查询应使用相同的分词器和模型。日文需要显式创建并配置词典的 Lindera 模型，不能直接套用其他分词接口中的裸模型名称。

### 对象与配置索引

- `create_tokenizer`、`drop_tokenizer`、`tokenize`：管理和执行分词器。
- `create_text_analyzer`、`apply_text_analyzer`：执行字符过滤、预分词和词元过滤。
- `create_custom_model_tokenizer_and_trigger`、`create_lindera_model`、`create_huggingface_model`：创建语料模型或导入模型。
- `create_stopwords`、`create_synonym`：管理词典。
- 内置模型包含 `llmlingua2`、`bert_base_uncased`、`wiki_tocken` 和 `gemma2b`。
- 0.1.1 新增 `ngram` 词元过滤器，`min_gram` 和 `max_gram` 范围为 1 到 255，`preserve_original` 默认为 false。配置使用 TOML。

### 升级与缓存边界

```sql
ALTER EXTENSION pg_tokenizer UPDATE TO '0.1.1';
```

替换预加载库后，应先重启 PostgreSQL，再更新数据库对象。0.1.0 到 0.1.1 的迁移没有新增 SQL 对象，行为变化在库文件中。文本分析器、模型及分词器按连接缓存，缓存不遵循事务隔离或回滚。回滚后若对象仍留在缓存中，可以重新连接或使用对应删除函数清理。扩展创建后不可重定位。它提供分词，排序及搜索索引由消费这些词元的扩展实现。
