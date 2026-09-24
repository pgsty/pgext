## 用法

来源：

- [Version 4.0.9 SQL](https://github.com/pgroonga/pgroonga/blob/4.0.9/data/pgroonga.sql)
- [Version 4.0.9 control](https://github.com/pgroonga/pgroonga/blob/4.0.9/pgroonga.control)
- [Official tutorial](https://pgroonga.github.io/tutorial/)
- [Version 4.0.9 release](https://github.com/pgroonga/pgroonga/releases/tag/4.0.9)
- [Upgrade guidance](https://pgroonga.github.io/upgrade/)

`pgroonga` 4.0.9 使用 Groonga 索引实现多语言全文检索，安装 `pgroonga` 访问方法及 SQL 操作符，普通使用不需要共享预加载。

### 核心流程

安装兼容的 PGroonga 与 Groonga 库后，由管理员创建扩展：

```sql
CREATE EXTENSION pgroonga;
CREATE TABLE search_notes (id bigint PRIMARY KEY, body text);
CREATE INDEX search_notes_body_idx ON search_notes USING pgroonga (body);
INSERT INTO search_notes VALUES (1, 'PostgreSQL supports full text search');
SELECT id, body FROM search_notes WHERE body &@ 'PostgreSQL';
SELECT id, body, pgroonga_score(tableoid, ctid) AS score
FROM search_notes WHERE body &@~ 'PostgreSQL OR Groonga'
ORDER BY score DESC;
```

### 主要对象

- `&@` 匹配关键词；`&@~` 接受 Groonga 查询语法。受支持的 LIKE/ILIKE 查询也可使用索引，并在需要时复查结果。
- `pgroonga_score(tableoid, ctid)` 取得搜索得分；按得分排序时，应确认实际执行了预期索引计划。
- `pgroonga_highlight_html()` 与 `pgroonga_query_extract_keywords()` 生成高亮结果；`pgroonga_snippet_html()` 提供关键词附近的文本。
- 4.0.9 新增 `pgroonga_physical_table_names(partitioned_index, prefix)`，以文本数组返回 Groonga 命令参数，标识各分区索引背后的物理表。该版本也开始为 PGroonga 扫描累加 `pg_stat_user_indexes.idx_scan`。

### 维护与权限

4.0.9 控制文件未声明受信任安装或可迁移模式，不应假定普通用户能安装扩展或将其移到其他模式。索引创建和查询仍需相应的表权限。应使扩展及 Groonga 库与目标 PostgreSQL 构建匹配，并在替换二进制前遵循上游升级说明。

PGroonga 除表数据外还管理派生索引文件，应据此规划磁盘容量和备份恢复流程。适用时使用 REINDEX 修复索引。独立的 `pgroonga_database` 模块用于恢复损坏的内部 Groonga 数据库，正常搜索并不需要它。在生产环境中，不要为了强制使用索引而全局关闭顺序扫描。
