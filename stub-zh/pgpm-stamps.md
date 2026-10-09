## 用法

来源：

- [packages/stamps/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/stamps/README.md)
- [packages/stamps/sql/pgpm-stamps--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/stamps/sql/pgpm-stamps--0.47.0.sql)
- [packages/stamps/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/stamps/Makefile)
- [packages/stamps/pgpm-stamps.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/stamps/pgpm-stamps.control)

`pgpm-stamps` 提供 `stamps.timestamps()` 与 `stamps.peoplestamps()` 触发器，分别维护创建/更新时间以及通过 `jwt_public.current_user_id()` 取得的创建/更新用户 UUID。

### 核心用法

```sql
CREATE EXTENSION "pgpm-stamps";
CREATE TABLE stamp_demo (
  id integer PRIMARY KEY,
  body text,
  created_at timestamptz,
  updated_at timestamptz
);
CREATE TRIGGER stamp_demo_timestamps
BEFORE INSERT OR UPDATE ON stamp_demo
FOR EACH ROW EXECUTE FUNCTION stamps.timestamps();
```

### 运行边界

触发器应在 INSERT 或 UPDATE 之前执行。时间记录要求存在 `created_at` 与 `updated_at`，用户记录要求存在 `created_by` 与 `updated_by`。更新时保留创建字段，重新填写更新字段。JWT 身份由应用提供；触发器维护的是行内元数据，并非独立且不可修改的历史记录。 依赖为 `plpgsql`, `pgpm-jwt-claims`, `pgpm-verify`。 安装脚本会向预先存在的角色 `authenticated`, `anonymous` 授权。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
