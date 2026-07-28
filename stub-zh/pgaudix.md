## 用法

来源：

- [官方上游 README](https://github.com/sipesistemas/pgaudix/blob/95d386c4329e76e78ee67e2478a6660673193960/README.md)
- [官方扩展控制文件 (pgaudix.control)](https://github.com/sipesistemas/pgaudix/blob/95d386c4329e76e78ee67e2478a6660673193960/pgaudix.control)
- [官方扩展 SQL (pgaudix--0.1.0--0.2.0.sql)](https://github.com/sipesistemas/pgaudix/blob/95d386c4329e76e78ee67e2478a6660673193960/pgaudix--0.1.0--0.2.0.sql)

`pgaudix` — 一个原生的 PostgreSQL 扩展，用于自动表审计。它将表列镜像到审计表中，并在源表结构发生变化时自动保持同步。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgaudix;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgaudix.audit_trigger()` 是一个扩展函数，返回 `trigger`。
- `pgaudix.ddl_sync()` 是一个扩展函数，返回 `event_trigger`。
- `pgaudix.disable(target_table regclass, drop_data boolean DEFAULT false)` 是一个扩展函数，返回 `void`。
- `pgaudix.drop_cleanup()` 是一个扩展函数，返回 `event_trigger`。
- `pgaudix.enable(target_table regclass)` 是一个扩展函数，返回 `void`。
- `pgaudix.status()` 是一个扩展函数，返回 `TABLE`。
- `pgaudix.truncate_trigger()` 是一个扩展函数，返回 `trigger`。
- `pgaudix.monitored_tables` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `0.2.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
