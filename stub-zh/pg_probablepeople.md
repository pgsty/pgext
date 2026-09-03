## 用法

来源：

- [官方文档](https://github.com/513analytics/pg_probablepeople/blob/a0587406b3d843e39f0bde52cc8587a972f317aa/README.md)
- [扩展控制文件](https://github.com/513analytics/pg_probablepeople/blob/a0587406b3d843e39f0bde52cc8587a972f317aa/pg_probablepeople.control)
- [官方仓库](https://github.com/513analytics/pg_probablepeople)

`pg_probablepeople` 使用 CRF 将个人或组织名称解析为结构化组成部分。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_probablepeople`：

```sql
CREATE EXTENSION pg_probablepeople;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT * FROM parse_name('Mr. John Doe');
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `parse_name` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `tag_name` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 目录生命周期为 preview；生产使用前应测试升级、备份恢复与服务器兼容性。
