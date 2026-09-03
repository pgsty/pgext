## 用法

来源：

- [官方文档](https://github.com/chebizarro/nostrpgx/blob/25ba7715e58dca4486cc6a5a81f8ad1493f58b2b/README.md)
- [扩展控制文件](https://github.com/chebizarro/nostrpgx/blob/25ba7715e58dca4486cc6a5a81f8ad1493f58b2b/nostr_filter.control)
- [官方仓库](https://github.com/chebizarro/nostrpgx)

`nostr_filter` 在 PostgreSQL 中提供 Nostr 事件类型、写入助手与 JSON 过滤计算。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `nostr_filter`：

```sql
CREATE EXTENSION nostr_filter;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT test_insert_nostr_event();
SELECT test_retrieve_nostr_event();
SELECT test_nostr_event_type();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `insert_nostr_event` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `insert_nostr_event_type` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `nostr_filter_search` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `nostr_event` | TYPE | 扩展创建的用户数据类型。 |
| `nostr_events` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `nostr_events` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `insert_nostr_event` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `insert_nostr_event_type` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `nostr_event` | TYPE | 扩展创建的用户数据类型。 |
| `nostr_filter_search` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
