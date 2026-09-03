## 用法

来源：

- [官方文档](https://github.com/verticalbarHQ/pg_ocpm/blob/bb025c78b4eb45c7a007b8868aedb43f7a4d5cf7/README.md)
- [扩展控制文件](https://github.com/verticalbarHQ/pg_ocpm/blob/bb025c78b4eb45c7a007b8868aedb43f7a4d5cf7/pg_ocpm.control)
- [官方仓库](https://github.com/verticalbarHQ/pg_ocpm)

`pg_ocpm` 提供面向对象流程挖掘的存储、遍历、聚合与紧凑分析载荷。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_ocpm`：

```sql
CREATE EXTENSION pg_ocpm;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE EXTENSION pg_ocpm;
SELECT ocpm.version();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `ocpm.version` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ocpm.dataset` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `ocpm.dataset_id` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ocpm.case_bucket` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `ocpm.rebuild_binding_index` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ocpm.activity_profile` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ocpm.adjacency_links` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ocpm.adjacency_neighborhood` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 13, 14, 15, 16, 17, 18；不要推断未列出的主版本。
- 扩展会在 `ocpm` 下固定或创建模式对象；权限与备份审查应包含这些对象。
