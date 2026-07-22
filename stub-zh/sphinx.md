## 用法

来源：

- [官方 README](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/README.org)
- [0.3 扩展 SQL](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/sphinx--0.3.sql)
- [扩展 control 文件](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/sphinx.control)

`sphinx` 0.3 通过 PostgreSQL 函数连接 Sphinx Search 服务器。它经由 Sphinx SQL 接口支持搜索、替换和删除文档以及生成摘要；仓库已经归档，因此只适合维护兼容的遗留部署。

### 核心流程

以超级用户创建扩展，再在扩展管理的配置表中设置连接参数。

```sql
CREATE EXTENSION sphinx;

UPDATE sphinx_config SET value = '127.0.0.1' WHERE key = 'host';
UPDATE sphinx_config SET value = '9306' WHERE key = 'port';
UPDATE sphinx_config SET value = 'test_' WHERE key = 'prefix';

SELECT *
FROM sphinx_select('blog_posts', 'photo', '', '', 0, 20, '');
```

可选前缀会添加到索引名称前，因此示例实际访问名为 `test_blog_posts` 的 Sphinx 索引。

### 主要对象

- `sphinx_select(...)`：返回搜索结果的 `(id, weight)` 行。
- `sphinx_replace(index, id, data)`：替换文档；数据数组由键和值交替组成。
- `sphinx_delete(index, id)`：删除文档。
- `sphinx_snippet(...)` 与 `sphinx_snippet_options(...)`：生成带高亮的摘要。
- `sphinx_config`：保存 `host`、`port`、`username`、`password` 与 `prefix`，并包含在扩展转储中。

0.3 版本在安装时把 `sphinx_config` 的全部权限授予 public。应撤销该授权，只委派最低必需权限，尤其是表中保存凭据时。调用同步依赖外部 Sphinx 服务，而文档更新属于外部副作用，不参与 PostgreSQL 事务回滚。
