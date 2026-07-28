## 用法

来源：

- [官方上游 README](https://github.com/willibrandon/pgl_validate/blob/fff68897716dc1d089719b0366039c6c79df3481/README.md)
- [官方扩展控制文件 (pgl_validate.control)](https://github.com/willibrandon/pgl_validate/blob/fff68897716dc1d089719b0366039c6c79df3481/pgl_validate.control)
- [官方实现源代码](https://github.com/willibrandon/pgl_validate/blob/fff68897716dc1d089719b0366039c6c79df3481/src/lib.rs)

`pgl_validate` — pgl_validate 是一个用于验证通过 pglogical 和 PostgreSQL 逻辑复制拓扑中的表内容的 PostgreSQL 扩展。在移动、转换或整合相应的数据时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgl_validate;

SELECT *
FROM pgl_validate.compare_table('public.accounts'::regclass);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
