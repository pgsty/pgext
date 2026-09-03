## 用法

来源：

- [官方文档](https://github.com/FHE-Postgres/pg_fhe/blob/e15047067c197f07dc7b85bd42e72a0cb9ec932e/pg_fhe/README.md)
- [扩展控制文件](https://github.com/FHE-Postgres/pg_fhe/blob/e15047067c197f07dc7b85bd42e72a0cb9ec932e/pg_fhe/sql/pg_fhe.control)
- [官方仓库](https://github.com/FHE-Postgres/pg_fhe)

`pg_fhe` 基于 Microsoft SEAL 对 PostgreSQL 中的 CKKS 密文执行同态运算。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_fhe`：

```sql
CREATE EXTENSION pg_fhe;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE TABLE test_ckks_mult as
    SELECT id, ckks_mult(data, 2.0) as data FROM test;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `ckks_mult` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
