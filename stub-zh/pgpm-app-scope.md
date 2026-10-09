## 用法

来源：

- [packages/app-scope/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/app-scope/README.md)
- [packages/app-scope/sql/pgpm-app-scope--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/app-scope/sql/pgpm-app-scope--0.47.0.sql)
- [packages/app-scope/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/app-scope/Makefile)
- [packages/app-scope/pgpm-app-scope.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/app-scope/pgpm-app-scope.control)

`pgpm-app-scope` 在已配置的 Constructive 目录中生成有序作用域，供就近匹配使用。`app_scope.frames` 返回作用域、查找数据库和键，`app_scope.local_frames` 仅遍历单个数据库。

### 核心用法

```sql
CREATE EXTENSION "pgpm-app-scope";
SELECT scope, lookup_database_id, key_value
FROM app_scope.frames(app_scope.platform_database_id(), 'platform');
```

### 运行边界

未注册平台数据库时，`app_scope.platform_database_id` 会报错；执行作用域不能为 NULL。查找依赖已填充的 metaschema 与作用域投影。上游明确要求不要在行安全策略中调用作用域遍历器。多个辅助函数采用 `SECURITY DEFINER`，应限制所有者、执行权限及相关目录的访问权限。 依赖为 `plpgsql`, `pgpm-verify`, `metaschema-schema`, `metaschema-modules`。 安装脚本会向预先存在的角色 `administrator`, `authenticated` 授权。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
