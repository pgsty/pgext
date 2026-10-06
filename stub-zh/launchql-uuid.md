## 用法

来源：

- [packages/uuid/readme.md](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/uuid/readme.md)
- [packages/uuid/sql/launchql-uuid--0.4.5.sql](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/uuid/sql/launchql-uuid--0.4.5.sql)
- [packages/uuid/Makefile](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/uuid/Makefile)
- [packages/uuid/launchql-uuid.control](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/uuid/launchql-uuid.control)

`launchql-uuid` 在 `uuids` 中提供伪有序、种子前缀 UUID 生成器及行触发器辅助函数。

### 核心用法

```sql
CREATE EXTENSION "launchql-uuid" CASCADE;
SELECT uuids.pseudo_order_uuid();
SELECT uuids.pseudo_order_seed_uuid('tenant-a');
```

### 运行边界

`uuids.trigger_set_uuid_seed` 和 `uuids.trigger_set_uuid_related_field` 要求版本化 SQL 中定义的字段与参数。生成的 UUID 既不是 UUID v7，也不是严格单调序列。此历史分发与 `pgpm-uuid` 使用相同模式，不应同时安装。

须先安装声明的依赖：`pgcrypto`, `plpgsql`, `uuid-ossp`, `hstore`。这是 SQL/PLpgSQL 代码，没有自己的共享库，也不要求预加载。控制文件允许非超级用户安装，但没有标记为 trusted；仍须满足依赖和模式权限。
