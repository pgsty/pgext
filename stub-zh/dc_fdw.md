## 用法

来源：

- [官方文档](https://github.com/ZhengYang/dc_fdw/blob/b362a57b8b7b934575688e1f428e8c5b3be2b861/README.md)
- [扩展控制文件](https://github.com/ZhengYang/dc_fdw/blob/b362a57b8b7b934575688e1f428e8c5b3be2b861/dc_fdw.control)
- [官方仓库](https://github.com/ZhengYang/dc_fdw)

`dc_fdw` 用于访问本地磁盘文档集合的历史外部数据包装器。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `dc_fdw`：

```sql
CREATE EXTENSION dc_fdw;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE SERVER dc_server
  FOREIGN DATA WRAPPER dc_fdw;

CREATE FOREIGN TABLE documents (id integer, content text)
SERVER dc_server
OPTIONS (
  data_dir '/srv/documents',
  index_dir '/srv/documents-index',
  index_method 'SPIM',
  buffer_size '10',
  id_col 'id',
  text_col 'content'
);

SELECT id, content FROM documents LIMIT 10;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `dc_fdw_handler` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `dc_fdw_validator` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 9.1, 9.2, 9.3；不要推断未列出的主版本。
- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
