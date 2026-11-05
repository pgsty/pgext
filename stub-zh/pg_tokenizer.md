## 用法

来源：

- [安装、升级、权限与备份](https://vchord.pgsty.com/docs/tokenizer/installation/)
- [分词与原生全文检索用法](https://vchord.pgsty.com/docs/tokenizer/overview/)
- [文本分析流水线](https://vchord.pgsty.com/docs/tokenizer/text-analyzer/)
- [模型与语料过滤](https://vchord.pgsty.com/docs/tokenizer/models/)
- [多语言分词](https://vchord.pgsty.com/docs/tokenizer/multilingual/)
- [SQL 与 TOML 参考](https://vchord.pgsty.com/docs/tokenizer/reference/)
- [1.0.0 控制文件](https://github.com/pgsty/pg_tokenizer/blob/v1.0.0/pg_tokenizer.control)

`pg_tokenizer` 1.0.0 为 PostgreSQL 14–18 提供可配置的文本分析和模型分词，可返回字符串词元、整数词元 ID，或供 GIN/GiST 与兼容 BM25 索引使用的原生 `tsvector`。它没有必需的扩展依赖，要求共享预加载，安装在固定的 `tokenizer_catalog` 模式中。

### 分析文本与建立索引

由管理员将库加入已有预加载列表，然后重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_tokenizer'
```

```sql
CREATE EXTENSION pg_tokenizer;
SET search_path = public, tokenizer_catalog;

SELECT create_text_analyzer('fts_words', $$
character_filters = ["to_lowercase"]
pre_tokenizer = "unicode_segmentation"
$$);
SELECT apply_text_analyzer('Cat dog cat', 'fts_words');
SELECT analyze_to_tsvector('Cat dog cat', 'fts_words');

CREATE TABLE fts_documents (id bigint PRIMARY KEY, passage text, terms tsvector);
INSERT INTO fts_documents VALUES
    (1, 'Cat dog cat', analyze_to_tsvector('Cat dog cat', 'fts_words'));
CREATE INDEX fts_documents_gin ON fts_documents USING gin (terms);
SELECT id FROM fts_documents WHERE terms @@ 'cat <-> dog'::tsquery;
```

`apply_text_analyzer` 返回 `text[]`，`analyze_to_tsvector` 将这些字符串转成带位置信息的词元。分析器定义来自目录表，因此这些函数是 `STABLE`。应将输出存入普通列，并在写入时维护；不能直接用于表达式索引或存储生成列。

```sql
CREATE FUNCTION fts_documents_set_terms() RETURNS trigger
LANGUAGE plpgsql AS $$
BEGIN
    NEW.terms := tokenizer_catalog.analyze_to_tsvector(NEW.passage, 'fts_words');
    RETURN NEW;
END;
$$;
CREATE TRIGGER fts_documents_terms
BEFORE INSERT OR UPDATE OF passage ON fts_documents
FOR EACH ROW EXECUTE FUNCTION fts_documents_set_terms();
```

短语查询示例使用已知的分析器词元。扩展没有提供通用的 `tsquery` 构造函数，查询词必须与文档采用相同的分析流程和表示方式。位置对应最终词元序列，而非原文字符偏移；过滤器可能改变短语距离。PostgreSQL 的限制仍然适用：过长词元会被跳过，超过 16,383 的位置会被截断，每个词元最多保留 256 个位置。

### 中文分析与模型分词

Jieba 是分析器中的预分词器，无需训练模型即可生成原生中文词元：

```sql
SELECT create_text_analyzer('zh_words', $$
[pre_tokenizer.jieba]
mode = "search"
enable_hmm = true
$$);
SELECT analyze_to_tsvector('PostgreSQL全文检索', 'zh_words');

SELECT create_tokenizer('bert_tokens', $$
model = "bert_base_uncased"
$$);
SELECT tokenize('hello hello', 'bert_tokens');
SELECT tokenize_to_tsvector('hello hello', 'bert_tokens');
```

`tokenize` 返回 `integer[]`，`tokenize_to_tsvector` 将模型 ID 作为十进制词元存储。两个原生向量函数都会保留重复的输出词元，不会对模型解码，也不会再应用 PostgreSQL 的文本规范化。`plainto_tsquery` 这类普通语言查询不会生成模型 ID。应向 BM25 查询构造函数传入匹配的向量，或用对应词元构造原生查询。

### 函数与配置索引

| 接口 | 用途 |
| --- | --- |
| `create_text_analyzer`、`drop_text_analyzer` | 管理由字符过滤器、一个预分词器和词元过滤器组成的 TOML 流水线。 |
| `create_tokenizer`、`drop_tokenizer` | 将分析器与词表或子词模型组合。 |
| `create_custom_model_tokenizer_and_trigger` | 训练自定义词表、回填整数词元 ID 并创建维护触发器；可选的第七个 `filter` 参数用于筛选语料行。 |
| `create_huggingface_model`、`create_lindera_model` | 导入 Hugging Face 分词器定义，或配置 Lindera 词典。 |
| `create_stopwords`、`create_synonym` | 为分析器过滤器创建命名词典。 |
| `add_preload_model`、`remove_preload_model`、`list_preload_models` | 管理或查看模型预加载；修改需要获得授权的管理员，并在重启后生效。 |
| `llmlingua2`、`bert_base_uncased`、`wiki_tocken`、`gemma2b` | 内置模型名称。 |

自定义模型辅助函数仍填充整数数组，而非原生向量。它的过滤条件是针对源列的可信 SQL 布尔表达式：只有结果为真的行参与处理，词表只追加，离开过滤范围的行保留原目标值。不能传入不可信 SQL。使用原生全文索引时，应另外维护向量列。

### 升级、权限与备份

安装 1.0.0 文件并重启 PostgreSQL，然后更新各个数据库：

```sql
ALTER EXTENSION pg_tokenizer UPDATE TO '1.0.0';
```

从 0.1.0 和 0.1.1 升级会保留目录行、自定义模型 ID 和六参数辅助函数。如果表达式索引、生成列、分区键或扩展统计仍依赖旧版不可变分析器函数，升级会拒绝执行；应先改为维护普通列。已提交的定义修改现在会按 SQL 快照使后端缓存失效，但不会重写已存储的词元数组或向量。

Lindera 6 改变了词典的词元 ID。应使用 Lindera 6 重建外部词典，重新生成受影响的词元或向量，并重建索引。原生 RPM/DEB 构建不嵌入 Lindera 词典，需要自行提供外部词典。逻辑备份之外，还应单独备份这些文件及集群的 `pg_tokenizer/preload_models` 目录。

`pg_database_owner` 获得数据库内目录表和函数的访问权限，应用角色需要显式授权；删除对象仍受所有权限制。集群级预加载修改不参与事务，并单独限制权限。引入新词元的词表写入按模型串行执行；在可重复读或可串行化隔离级别下遇到 SQLSTATE `40001` 时应重试。修改定义的 API 不允许两阶段事务准备。

普通 `pg_dump` 现在会包含分析器、分词器、模型、停用词和同义词定义，以及自定义词表、序列和源表触发器。随扩展提供的 `lucene_english`、`nltk_english` 和 `iso_english` 列表是保留名称，由扩展安装过程重新创建；自定义列表应使用其他名称。
