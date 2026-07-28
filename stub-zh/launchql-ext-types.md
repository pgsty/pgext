## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/types/README.md)
- [官方扩展控制文件 (launchql-ext-types.control)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/types/launchql-ext-types.control)
- [官方扩展 SQL (launchql-ext-types--0.4.5.sql)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/types/sql/launchql-ext-types--0.4.5.sql)

`launchql-ext-types` — PostgreSQL 扩展，提供带有内置验证的自定义域类型。此扩展包含一组常用的数据类型及其验证约束，便于在数据库级别直接强制执行数据完整性。当应用程序数据需要此类类型、域或其操作符时，请使用此扩展。在安装此扩展之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "launchql-ext-types";

-- Create a table using the custom domain types
CREATE TABLE users (
  id serial PRIMARY KEY,
  email email NOT NULL,
  website url,
  profile_image image,
  origin origin
);

-- Insert data with automatic validation
INSERT INTO users (email, website, profile_image, origin)
VALUES (
  'user@example.com',
  'https://example.com',
  '{"url": "https://example.com/profile.jpg", "mime": "image/jpeg"}',
  'https://example.com'
);

-- Invalid data will be rejected automatically
INSERT INTO users (email) VALUES ('not-an-email'); -- Fails validation
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `attachment` 是扩展定义的域。
- `email` 是扩展定义的域。
- `hostname` 是扩展定义的域。
- `image` 是扩展定义的域。
- `multiple_select` 是扩展定义的域。
- `origin` 是扩展定义的域。
- `single_select` 是扩展定义的域。
- `upload` 是扩展定义的域。
- `url` 是扩展定义的域。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.4.5`。
- 请先安装确认的扩展依赖项：`plpgsql`, `citext`。
- 控制文件将此扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
