## 用法

来源：

- [Control](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.7.0/softvisio_admin.control)
- [Version 1.7.0 SQL](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.7.0/softvisio_admin--1.7.0.sql)
- [Installed SQL files](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.7.0/Makefile)
- [Changelog](https://github.com/softvisio/postgresql-softvisio-admin/blob/v1.7.0/CHANGELOG.md)

`softvisio_admin` 1.7.0 用 SQL/PL/pgSQL 提供数据库与角色创建，以及跨库扩展盘点和升级功能。安装需要超级用户、`dblink` 和 `pgcrypto`，无需原生库或预加载。

### 创建与检查

实际 1.7.0 SQL 的 `create_database(name text)` 只接收一个参数。它为缺少的登录角色生成密码，从 `template0` 创建使用 ICU 区域 `und` 与 `C.UTF-8` 的 UTF8 数据库，再设置所有权并创建模式。服务器必须支持这些区域配置。旧 README 中可选的排序规则参数与此 SQL 版本不符。

```sql
CREATE EXTENSION softvisio_admin CASCADE;
CALL create_database('appdb');
SELECT * FROM outdated_extensions();
```

### 跨库操作

`outdated_extensions()` 通过 dblink 遍历允许连接的非模板数据库，比较已安装版本与默认版本。`update_extensions()` 跨这些库执行升级，并在适用时调用 PostGIS 升级辅助函数。远端错误可能以提示形式报告后继续执行，应检查返回的安装版本及提示。远程语句独立提交，失败不会撤销先前数据库或创建步骤的修改。

### 升级与权限边界

1.7.0 源码只安装新的基础 SQL，并移除了历史升级脚本，因此仅靠此发行本身不能为旧安装提供常规升级路径。修改服务器文件前，应检查 `pg_extension_update_paths('softvisio_admin')` 并准备经过审查的迁移方案，不要假定 `ALTER EXTENSION UPDATE` 一定成功。这些例程应仅开放给可信管理员：它们能执行权限很高的远程 SQL，新生成的密码还会通过 NOTICE 输出，应限制客户端与服务器日志的访问。调用跨库升级器前，应审查并测试每个受影响扩展的升级。
