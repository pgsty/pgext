## 用法

来源：

- [官方 README](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/README.md)
- [1.4 版 SQL API](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/sphinxlink--1.4.sql)
- [扩展控制文件](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/sphinxlink.control)

`sphinxlink` 通过 SphinxSearch 或 ManticoreSearch 的 MySQL 兼容协议发送 SQL，并将结果暴露为 PostgreSQL 记录集。它可以在后端中保留命名连接，也可以为一次查询打开连接，还能暴露上一次远程搜索的元数据。

### 连接与查询

由于 `sphinx_query()` 返回由调用方定义形状的记录，调用方必须提供与远程结果匹配的列定义列表。

```sql
CREATE EXTENSION sphinxlink;

SELECT sphinx_connect('search', '127.0.0.1', 9306);

SELECT *
FROM sphinx_query(
  'search',
  'SELECT id, weight() AS rank FROM docs WHERE MATCH(?)',
  'postgresql extension'
) AS result(id bigint, rank integer);

SELECT * FROM sphinx_meta('search');
SELECT sphinx_disconnect('search');
```

`sphinx_connections()` 列出已打开的命名连接。`sphinx_query_params(host, port, query)` 变体会自动连接。带 `match_clause` 的重载替换 `MATCH(?)` 值；其他查询结构仍属于远程 SQL 字符串。

### 远程系统边界

查询从 PostgreSQL 后端同步执行，且不受 PostgreSQL 事务回滚控制。网络阻塞、远端限制、结果形状变化、协议错误及 Manticore/Sphinx 版本差异都可能使数据库会话失败或被占用。应设置外部网络超时与语句超时、限制可访问主机，并在代表性数据上匹配每个返回列类型。

应把查询字符串视为远程 SQL，绝不能拼接不可信标识符、谓词或命令。match-clause 重载只保护其预期 `MATCH(?)` 值，并不是通用 SQL 参数化 API。应审查默认函数授权、只允许可信调用方使用、监控连接池和取消场景下的连接清理，并且不要假定上次查询元数据会跨不同后端连接保留。
