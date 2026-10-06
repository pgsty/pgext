## 用法

来源：

- [packages/faker/readme.md](https://github.com/pyramation/faker/blob/114980fb274f487142f9170a24c452427fd1ef51/packages/faker/readme.md)
- [packages/faker/sql/launchql-faker--0.1.0.sql](https://github.com/pyramation/faker/blob/114980fb274f487142f9170a24c452427fd1ef51/packages/faker/sql/launchql-faker--0.1.0.sql)
- [packages/faker/Makefile](https://github.com/pyramation/faker/blob/114980fb274f487142f9170a24c452427fd1ef51/packages/faker/Makefile)
- [packages/faker/launchql-faker.control](https://github.com/pyramation/faker/blob/114980fb274f487142f9170a24c452427fd1ef51/packages/faker/launchql-faker.control)

`launchql-faker` 提供 `faker` 模式、词典表，以及名称、地址、文本、标识符和数值的随机生成器。

### 核心用法

```sql
CREATE EXTENSION "launchql-faker" CASCADE;
SELECT faker.fullname(), faker.email(), faker.integer(1, 100);
```

### 运行边界

`faker.name`、`faker.fullname`、`faker.email`、`faker.address`、`faker.sentence`、`faker.integer` 与 `faker.uuid` 是主要生成器。取值依赖内置词典和随机数，不能视为已验证的联系资料，也不能将生成的密码或令牌当作安全保证。此历史版本未声明当前 PostgreSQL 主版本矩阵。

须先安装声明的依赖：`citext`, `pgcrypto`, `plpgsql`, `uuid-ossp`, `launchql-ext-types`。这是 SQL/PLpgSQL 代码，没有自己的共享库，也不要求预加载。控制文件允许非超级用户安装，但没有标记为 trusted；仍须满足依赖和模式权限。
