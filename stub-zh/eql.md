## 用法

来源：

- [官方 `cipherstash@eql` 2.2.1 包页面](https://database.dev/cipherstash/eql)
- [EQL 官方文档](https://cipherstash.com/docs)
- [当前上游 3.0.2 版本](https://github.com/cipherstash/encrypt-query-language/releases/tag/eql-3.0.2)

`eql`（包名为 `cipherstash@eql`）为可搜索密文提供 PostgreSQL 类型、函数、操作符与索引配置。目录中的包版本是 2.2.1，使用 `eql_v2` API。EQL 只提供数据库端表示与查询接口；实际加解密必须由 PostgreSQL 外部的 CipherStash Proxy 或 Protect.js 完成。

### 核心流程

生成并应用固定版本的 dbdev 迁移，然后确认已安装的 EQL API 版本：

```sh
dbdev add -o ./migrations -s extensions -v 2.2.1 package -n "cipherstash@eql"
psql -f ./migrations/*.sql
```

```sql
SELECT eql_v2.version();
```

使用持久化的 public 类型定义密文列，注册其明文类型，并按需增加可搜索索引配置：

```sql
CREATE TABLE users (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  encrypted_email public.eql_v2_encrypted
);

SELECT eql_v2.add_column('users', 'encrypted_email', 'text');
SELECT eql_v2.add_search_config(
  'users', 'encrypted_email', 'unique', 'text'
);
```

应通过配置好的 CipherStash 客户端读写该列。直接向 `public.eql_v2_encrypted` 插入明文并不会让 EQL 自动加密。

### 重要对象

- `public.eql_v2_encrypted` 是密文列类型，底层保存经过约束校验的 `jsonb` 数据。
- `public.eql_v2_configuration` 保存可搜索加密配置；`public.eql_v2_configuration_state` 是其状态类型。
- `eql_v2` 模式包含配置和查询密文所需的函数与操作符。
- `eql_v2.add_column(...)` 注册密文列及其逻辑明文类型。
- `eql_v2.add_search_config(...)` 增加 unique/equality 等搜索配置。
- `eql_v2.version()` 返回已安装的 API 版本。

### 权限、升级与版本边界

安装角色需要数据库和 public 模式上的 `CREATE` 权限。迁移角色还需要目标表上的 `ALTER` 以及 `public.eql_v2_configuration` 的读写权限；运行时角色通常只需要 `eql_v2` 的 `USAGE`、函数执行权限、配置表只读权限和普通业务表权限。

删除 `public.eql_v2_encrypted` 可能连带删除应用列。EQL 特意把持久化 public 类型和配置数据放在可替换的 `eql_v2` 模式之外，因此卸载和大版本升级必须按数据迁移处理。

GitHub 上游 3.0.2 已改用差异很大的 `eql_v3` domain API 与发布版 SQL 安装器。上游 README 明确提醒 dbdev 发布可能滞后于 GitHub。不要把 v3 示例用于目录中的 `cipherstash@eql` 2.2.1 包；应固定分发版本并使用与之匹配的文档。
