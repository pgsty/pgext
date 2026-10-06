## 用法

来源：

- [packages/uuid/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/uuid/README.md)
- [packages/uuid/sql/pgpm-uuid--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/uuid/sql/pgpm-uuid--0.47.0.sql)
- [packages/uuid/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/uuid/Makefile)
- [packages/uuid/pgpm-uuid.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/uuid/pgpm-uuid.control)

`pgpm-uuid` 在 `uuids` 中安装伪有序 UUID 生成器及触发器辅助函数，可使标识符按时间或种子聚集，同时保留随机分量。

### 核心用法

```sql
CREATE EXTENSION "pgpm-uuid" CASCADE;
SELECT uuids.pseudo_order_uuid();
SELECT uuids.pseudo_order_seed_uuid('tenant-a');
```

### 运行边界

生成器不是 UUID v7，也不是严格单调序列。`uuids.trigger_set_uuid_seed` 与 `uuids.trigger_set_uuid_related_field` 要求相应触发器参数和字段。模式与历史 `launchql-uuid` 重叠，不应同时安装。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`pgcrypto`、`plpgsql`、`uuid-ossp`、`hstore`、`pgpm-verify`。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
