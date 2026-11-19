## 用法

来源：

- [Official source](https://github.com/pgbigm/pg_bigm/blob/v1.2-20261118/docs/pg_bigm_en.md)
- [Official source](https://github.com/pgbigm/pg_bigm/blob/v1.2-20261118/pg_bigm.control)
- [Official source](https://github.com/pgbigm/pg_bigm/blob/v1.2-20261118/README.md)
- [Official source](https://github.com/pgbigm/pg_bigm/releases/tag/v1.2-20261118)

`pg_bigm` 使用 GIN 二元组索引加速 LIKE 与相似度搜索，包括多字节文本。1.2-20261118 发行版保留 SQL 版本 1.2，并增加 PostgreSQL 19 支持。

### 搜索

```sql
CREATE EXTENSION pg_bigm;
CREATE TABLE documents (id bigint PRIMARY KEY, content text);
INSERT INTO documents VALUES (1, 'PostgreSQL database search');
CREATE INDEX docs_bigm_idx ON documents USING gin (content gin_bigm_ops);
SELECT * FROM documents WHERE content LIKE likequery('database');
SELECT show_bigm('PostgreSQL');
SET pg_bigm.similarity_limit = 0.2;
SELECT * FROM documents WHERE content =% 'database search';
```

`likequery(text)` 将字面搜索关键字转义为 LIKE 条件；`show_bigm(text)` 显示二元组，`bigm_similarity(text,text)` 返回相似度。`=%` 运算符使用通常为 0.3 的 `pg_bigm.similarity_limit`。只有规划器选择索引时，索引才参与查询。

### 设置与维护

通常应保持 `pg_bigm.enable_recheck` 开启，关闭复查可能返回假阳性结果。`pg_bigm.gin_key_limit` 默认 0，表示不限制；正值限制搜索所使用的键数。`pg_gin_pending_stats(regclass)` 返回 GIN 待处理列表的页面数和元组数，`pg_bigm.last_update` 标识库修订。

20261118 构建修复多字节截断错误、LIKE 下划线处理及相似度匹配遗漏。截断修复要求使用 PostgreSQL 14.21、15.16、16.12、17.8、18.2 或更新补丁版本，或 PostgreSQL 19 的头文件编译；SQL 扩展版本相同不证明运行库已包含修复。应在所有服务器部署匹配库，并替换仍使用旧库的进程。基本用法无需预加载。
