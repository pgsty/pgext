## 用法

来源：

- [官方文档](https://github.com/pgmoneta/pgmoneta_ext/blob/4b642dd88f3324c80aa9b30759a18df4773afe13/README.md)
- [扩展控制文件](https://github.com/pgmoneta/pgmoneta_ext/blob/4b642dd88f3324c80aa9b30759a18df4773afe13/sql/pgmoneta_ext.control)
- [官方仓库](https://github.com/pgmoneta/pgmoneta_ext)

`pgmoneta_ext` 为 pgmoneta 提供服务端数据块与增量备份辅助功能。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pgmoneta_ext`：

```sql
CREATE EXTENSION pgmoneta_ext;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT pgmoneta_ext_version();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pgmoneta_ext_checkpoint` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgmoneta_ext_fips` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgmoneta_ext_get_file` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgmoneta_ext_get_files` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgmoneta_ext_get_oid` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgmoneta_ext_get_oids` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgmoneta_ext_promote` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pgmoneta_ext_receive_file_chunk` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 14；不要推断未列出的主版本。
