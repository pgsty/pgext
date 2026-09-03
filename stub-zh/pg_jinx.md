## 用法

来源：

- [官方文档](https://github.com/sourcewave/pg_jinx/blob/651430bb03e981ecef9d3e3e1c4441ccdbf522db/README.md)
- [扩展控制文件](https://github.com/sourcewave/pg_jinx/blob/651430bb03e981ecef9d3e3e1c4441ccdbf522db/pg_jinx.control)
- [官方仓库](https://github.com/sourcewave/pg_jinx)

`pg_jinx` 通过 JNI 使用 Java 编写 PostgreSQL 存储过程、触发器与外部数据包装器。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_jinx`：

```sql
CREATE EXTENSION pg_jinx;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT * FROM jproperties;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `jinx.TestTrigger` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `jinx.countRows` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `jinx.example_rs` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `jinx.fdw_handler` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `jinx.fdw_validator` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `jinx.getGravatar` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `jinx.getOptions` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `jinx.inline_handler` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
