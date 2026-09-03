## 用法

来源：

- [官方文档](https://github.com/semantius/semantius/blob/eb80b5f1f596ad2d206a6688bd362b5db3aa85d3/extension/README.md)
- [扩展控制文件](https://github.com/semantius/semantius/blob/eb80b5f1f596ad2d206a6688bd362b5db3aa85d3/extension/pg_semantius.control)
- [官方仓库](https://github.com/semantius/semantius)

`pg_semantius` 以数据库为中心、集成 RBAC、RLS、数据字典与消息队列的应用后端。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_semantius`：

```sql
CREATE EXTENSION pg_semantius CASCADE;
```

经审查的控制文件或官方流程要求 `pgcrypto`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
BEGIN;
SET LOCAL ROLE authenticated;                              -- the identity role (member of semantius_user)
SELECT set_config('request.jwt.claims', $1::text, true);  -- LOCAL; inject BEFORE any rbac call
-- … queries …
COMMIT;
```

### 主要对象

官方来源通过动态方式或供应商工具定义扩展接口；授权前应检查实际安装版本。

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 18；不要推断未列出的主版本。
