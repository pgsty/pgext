## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/sha/sha-1.1.0/README)
- [官方扩展 SQL (sha.sql)](https://api.pgxn.org/src/sha/sha-1.1.0/sql/sha.sql)

`sha` — 该模块实现了 sha1、sha224、sha256、sha384、sha512 和 md5hash 数据类型。您可以对这些类型应用基本的比较操作符，并使用它们创建索引。支持 btree 和 hash 索引。请参阅 sql/sha.sql 以获取用法示例。当应用程序需要此类型、域或其操作符时，请使用它。请使用上述锁定的上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

经过审核的分发包使用了过时的 SQL 或非控制文件安装布局，因此没有建立现代独立的 `CREATE EXTENSION` 和升级工作流程。请遵循锁定的上游安装机制，并在隔离数据库中验证安装对象。

### 要求与注意事项

- 该目录记录了版本信息 `1.1.0`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以与锁定的源代码进行比对。
