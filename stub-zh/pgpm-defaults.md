## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/defaults/README.md)
- [官方扩展控制文件 (pgpm-defaults.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/defaults/pgpm-defaults.control)
- [官方扩展 SQL (pgpm-defaults--0.15.5.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/defaults/sql/pgpm-defaults--0.15.5.sql)

`pgpm-defaults` — @pgpm/defaults 通过撤销默认的公共访问权限来为 PostgreSQL 数据库建立一个安全的基础配置。该包通过移除 PostgreSQL 的默认宽松设置并要求显式权限授予来实现最小权限原则。在实施相应的安全、审计或访问控制工作流时使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "pgpm-defaults";
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.15.5`。
- 先安装确认的扩展依赖项：`plpgsql`, `pgpm-verify`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，需确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
