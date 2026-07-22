## 用法

来源：

- [官方 dbdev 软件包页面](https://database.dev/supabase/dbdev)
- [官方软件包安装文档](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/website/content/docs/install-a-package.mdx)
- [官方项目 README](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/README.md)

`dbdev` 是 database.dev PostgreSQL 受信任语言扩展注册表的数据库内客户端。该目录条目对应注册表软件包 `supabase@dbdev`；它下载软件包 SQL，通过 `pg_tle` 注册，并把已安装的软件包暴露为 PostgreSQL 扩展。

### 核心流程

先安装 `pg_tle` 和 `http`，再按照官方软件包页面中的引导 SQL 注册并创建 `supabase@dbdev`。之后即可安装指定版本的软件包并将其启用为扩展：

```sql
CREATE EXTENSION IF NOT EXISTS http WITH SCHEMA extensions;
CREATE EXTENSION IF NOT EXISTS pg_tle;

-- Run the current bootstrap block from the official supabase@dbdev page.

SELECT dbdev.install('owner@package');
CREATE EXTENSION "owner@package"
    SCHEMA public
    VERSION '1.2.3';
```

主要 SQL API 是 `dbdev.install(package_name text)`。它从注册表获取可用版本与升级路径，调用 `pg_tle` 安装函数，并登记一次下载事件。为了确保可复现性，应在 `CREATE EXTENSION` 中固定扩展版本。

### 运维说明

版本 `0.0.5` 需要 `http` 与 `pg_tle` 扩展，以及访问 database.dev 的出站 HTTPS。安装软件包会在数据库内执行从远程注册表获取的 SQL，因此应在安装前审查所选软件包及版本，并控制谁可以调用 `dbdev.install`。官方文档警告，包含 TLE 的数据库可能无法完成逻辑恢复，并建议为该工作流使用物理备份。
