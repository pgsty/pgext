## 用法

来源：

- [官方 README](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/README.md)
- [官方基础扩展 SQL](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/init.sql)
- [官方 RUM 支持 SQL](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/rum_support.sql)
- [官方扩展控制文件](https://github.com/postgrespro/tsvector2/blob/f3fd908c4b09729524052ca99790e30058aa112b/tsvector2.control)

`tsvector2` 是一种扩展全文搜索向量类型，旨在改善压缩并移除原始 `tsvector` 的 1 MB 大小限制。版本 `1.0` 对应常用 PostgreSQL 文本搜索操作，并支持 B-tree、GIN、GiST 与可选 RUM 索引。

### 核心流程

存储生成的 `tsvector2` 值，并使用扩展提供的 GIN 操作符类创建索引：

```sql
CREATE EXTENSION tsvector2;

CREATE TABLE article (
  id bigint PRIMARY KEY,
  body text,
  search tsvector2
);

UPDATE article
SET search = to_tsvector2('english'::regconfig, body);

CREATE INDEX article_search_idx
  ON article USING gin (search gin_tsvector2_ops);

SELECT id
FROM article
WHERE search @@ plainto_tsquery('english', 'database search');
```

该类型支持比较、通过 `||` 连接、通过 `@@` 匹配，以及排名、权重、过滤、词素删除和从文本、数组、JSON、JSONB 转换。

### 重要对象与索引

扩展专用入口包括 `to_tsvector2`、`array_to_tsvector2`、`tsvector2_to_array`、`tsvector2_stat`、`jsonb_to_tsvector2`、`json_to_tsvector2`、`tsvector2_update_trigger` 与 `tsvector2_update_trigger_column`。`strip`、`unnest`、`length`、`setweight`、`ts_rank`、`ts_rank_cd`、`ts_delete` 和 `ts_filter` 等通用名称针对新类型提供了重载。

默认操作符类是 `bt_tsvector2_ops`、`gin_tsvector2_ops` 与 `gist_tsvector2_ops`。如果创建 `tsvector2` 时 RUM 扩展已经存在，安装过程会通过 `install_rum_for_tsvector2()` 添加 RUM 距离、评分及操作符类；需要这些支持时应先安装 RUM。

扩展固定在 `public` 模式中，且不可重定位。上游说明支持 PostgreSQL 9.6 或更高版本，并指出 PostgreSQL 10 与 11 的 JSON 转换行为差异与核心文本搜索相同。替换核心 `tsvector` 存储前，应在准确的服务器版本上测试应用查询、生成列或触发器，以及索引重建流程。
