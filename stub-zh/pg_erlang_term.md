## 用法

来源：

- [官方文档](https://github.com/flambard/pg_erlang_term/blob/32715a42e25ad75da709cdac61bb9af6c8617f61/README.md)
- [扩展控制文件](https://github.com/flambard/pg_erlang_term/blob/32715a42e25ad75da709cdac61bb9af6c8617f61/pg_erlang_term.control)
- [官方仓库](https://github.com/flambard/pg_erlang_term)

`pg_erlang_term` 用于编码、存储与解码 Erlang External Term Format 值的 PostgreSQL 类型。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_erlang_term`：

```sql
CREATE EXTENSION pg_erlang_term;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE TABLE messages (payload erlang_term);
INSERT INTO messages VALUES ('{ok,42}'::erlang_term);
SELECT payload FROM messages;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `erlang_term` | TYPE | 扩展创建的用户数据类型。 |
| `erlang_term_decode` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_term_encode` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_term_input` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_term_output` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_term_receive` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_term_send` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
