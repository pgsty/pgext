## 用法

来源：

- [0.3.0 README](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/README.md)
- [Control file](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/vchord_bm25.control)
- [0.3.0 SQL objects](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/sql/install/vchord_bm25--0.3.0.sql)
- [Query settings](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/src/guc.rs)
- [0.3.0 migration](https://github.com/supervc-stack/VectorChord-bm25/blob/0.3.0/sql/vchord_bm25--0.2.2--0.3.0.sql)
- [0.3.0 release](https://github.com/supervc-stack/VectorChord-bm25/releases/tag/0.3.0)
- [Tokenizer installation](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/01-installation.md)
- [Tokenizer models](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/06-model.md)

`vchord_bm25` 使用稀疏词元频率类型和 `bm25` 索引访问方法提供 BM25 排序。分词由独立组件提供，通常使用 pg_tokenizer。扩展对象安装在固定的 `bm25_catalog` 模式中，创建时需要超级用户权限。

### 核心流程

示例使用 pg_tokenizer，它要求预加载并重启。修改预加载列表时保留已有条目：

```conf
shared_preload_libraries = 'pg_tokenizer'
```

```sql
CREATE EXTENSION pg_tokenizer;
CREATE EXTENSION vchord_bm25;
SET search_path = public, tokenizer_catalog, bm25_catalog;

SELECT create_tokenizer('english', $$
model = "bert_base_uncased"
$$);
CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    passage text,
    embedding bm25vector
);
INSERT INTO documents(passage) VALUES ('PostgreSQL full text search');
UPDATE documents SET embedding = tokenize(passage, 'english')::bm25vector;
CREATE INDEX documents_bm25 ON documents USING bm25 (embedding bm25_ops);

SELECT id, passage,
       embedding <&> to_bm25query('documents_bm25',
           tokenize('PostgreSQL', 'english')::bm25vector) AS score
FROM documents
ORDER BY score
LIMIT 10;
```

索引为 `to_bm25query` 提供语料统计。`<&>` 返回负分数，因此升序排列将相关性更高的结果放在前面。文档与查询应使用相同的分词器和模型。源文本变化时需要更新存储的词元向量，也可使用分词器提供的维护触发器辅助函数。词汇表变化后，应先重新分词存量文档，再重建索引。

### 类型、函数与搜索限制

- `bm25vector` 保存词元 ID 和频率，整数数组转换会合并重复 ID，并丢弃词元顺序。
- `bm25query` 将查询向量绑定到索引，由 `to_bm25query(regclass, bm25vector)` 构造；`bm25_ops` 是索引运算符类。
- `bm25_catalog.bm25_limit` 默认为 100，限制索引返回的候选数量。较大的 SQL 限制或严格过滤需要增加此参数，仅改变 SQL LIMIT 不会增加候选预算。
- `bm25_catalog.enable_index` 控制是否使用索引，`bm25_catalog.enable_prefilter` 控制预过滤，两者默认均为 true。
- `bm25_catalog.segment_growing_max_page_size` 默认为 4096 页，超过后将增长段封存。

```sql
SET bm25_catalog.bm25_limit = 1000;
```

访问方法名称在数据库中是全局的，不能与其他创建同名 bm25 访问方法的扩展共存，包括 pg_textsearch 和 pg_search 的兼容别名。稀疏频率不保留短语匹配所需的位置。中文可使用带 Jieba 预分词器的自定义语料模型，日文 Lindera 支持取决于分词器的构建选项和词典配置。

### 升级到 0.3.0

```sql
ALTER EXTENSION vchord_bm25 UPDATE TO '0.3.0';
```

0.2.2 到 0.3.0 的迁移新增 `bm25_page_inspect(regclass, integer)`，返回页面诊断文本。本次发布改变小词元在封存段中的页面分配，未声明强制重建索引要求。更新数据库对象前应安装匹配的扩展文件；替换预加载的分词器库还需要重启。应保持分词与排序组件的升级兼容，并检查代表性查询结果。
