## 用法

来源：

- [官方文档](https://github.com/samrose/pg-erl/blob/03e81b7f10c083246456dc798f6228539c04ad05/README.md)
- [扩展控制文件](https://github.com/samrose/pg-erl/blob/03e81b7f10c083246456dc798f6228539c04ad05/erlang_cnode.control)
- [官方仓库](https://github.com/samrose/pg-erl)

`erlang_cnode` 将 PostgreSQL 会话连接到 Erlang 或 Elixir 节点的 C-node 桥接扩展。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `erlang_cnode`：

```sql
CREATE EXTENSION erlang_cnode;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Connect to an Erlang node
SELECT erlang_connect('testnode@127.0.1.1', 'cookie123');

-- Call a remote function
SELECT erlang_call('testnode@127.0.1.1', 'erlang', 'node', '[]'::jsonb);

-- Disconnect
SELECT erlang_disconnect('testnode@127.0.1.1');
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `erlang_call` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_connect` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_disconnect` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_cast` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_check_connection` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_pending_requests` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_ping` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `erlang_receive_async` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
