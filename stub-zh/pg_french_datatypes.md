## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_french_datatypes/pg_french_datatypes-0.1.1/README)
- [官方扩展 SQL (pg_french_datatypes.sql)](https://api.pgxn.org/src/pg_french_datatypes/pg_french_datatypes-0.1.1/sql/pg_french_datatypes.sql)

`pg_french_datatypes` — 这个小项目旨在包含以法国为中心的数据类型，例如：。当应用程序需要这种类型、域或其操作符时，请使用它。请使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

经过审查的分发包使用了过时的 SQL 或非控制文件安装布局，因此没有建立现代独立的 `CREATE EXTENSION` 和升级工作流。请遵循上游固定版本的安装机制，并在隔离数据库中验证安装的对象。

### 重要对象

- `jour` 是扩展定义的类型。
- `mois` 是扩展定义的类型。
- `code_postal_fr` 是扩展定义的域。
- `numero_securite_sociale_fr` 是扩展定义的域。

### 要求与注意事项

- 请确认版本记录 `0.1.1`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源进行比对。
