## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/faker/README.md)
- [官方扩展控制文件 (pgpm-faker.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/faker/pgpm-faker.control)
- [官方扩展 SQL (pgpm-faker--0.15.5.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/faker/sql/pgpm-faker--0.15.5.sql)

`pgpm-faker` — @pgpm/faker 提供了一套全面的生成假数据的功能直接在 PostgreSQL 中实现。适用于填充测试数据库、创建演示数据和开发环境。所有函数都是纯 plpgsql 实现，返回看起来很真实的数据，无需外部依赖。当 SQL 需要这些特殊功能或聚合时可以使用它。上游将此功能描述为实验性的。

### 核心工作流

```sql
CREATE EXTENSION "pgpm-faker";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `faker.address(state text DEFAULT NULL, city text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `faker.attachment(mime text DEFAULT NULL)` 是一个扩展函数，返回 `attachment`。
- `faker.birthdate(min int DEFAULT 1, max int DEFAULT 100)` 是一个扩展函数，返回 `date`。
- `faker.boolean()` 是一个扩展函数，返回 `boolean`。
- `faker.business()` 是一个扩展函数，返回 `text`。
- `faker.city(state text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `faker.date(min int DEFAULT 1, max int DEFAULT 100, future boolean DEFAULT false)` 是一个扩展函数，返回 `date`。
- `faker.email()` 是一个扩展函数，返回 `text`。
- `faker.ext(mime text DEFAULT faker.mime())` 是一个扩展函数，返回 `text`。
- `faker.file(mime text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `faker.float(min double precision DEFAULT 0, max double precision DEFAULT 100)` 是一个扩展函数，返回 `double`。
- `faker.fullname(gender text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `faker.gender(gender text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `faker.hostname()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.15.5`。
- 先安装确认的扩展依赖项：`citext`, `pgcrypto`, `plpgsql`, `pgpm-types`, `pgpm-verify`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 上游将项目的一部分或全部标记为实验性的。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
