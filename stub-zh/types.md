## 用法

来源：

- [官方上游 README](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/readme.md)
- [官方扩展控制文件 (types.control)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/types/types.control)
- [官方扩展 SQL (types--0.0.1.sql)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/types/sql/types--0.0.1.sql)

`types` — 可重用的附件、电子邮件地址、主机名、图像、上传和 URL 的域。当应用程序数据需要此类型、域或其操作符时，请使用它。在安装应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION types;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `attachment` 是扩展定义的域。
- `email` 是扩展定义的域。
- `hostname` 是扩展定义的域。
- `image` 是扩展定义的域。
- `upload` 是扩展定义的域。
- `url` 是扩展定义的域。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 请先安装并验证确认的扩展依赖项：`plpgsql`, `uuid-ossp`, `citext`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
