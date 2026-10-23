## 用法

来源：

- [官方资料](https://github.com/Mazuco/pg_synonym/blob/77fa1f92327d2c26f21000293b623941f946dd2f/README.md)
- [官方资料](https://github.com/Mazuco/pg_synonym/blob/77fa1f92327d2c26f21000293b623941f946dd2f/pg_synonym.control)
- [官方资料](https://github.com/Mazuco/pg_synonym/blob/77fa1f92327d2c26f21000293b623941f946dd2f/pg_synonym--1.0.sql)

`pg_synonym` 使用按调用者权限执行的视图维护表与视图别名。发行版 1.0.0 安装的扩展版本为 1.0；要求 PostgreSQL 15 或更高版本，安装需要超级用户。不需要预加载或重启服务器。

### 核心流程

```sql
CREATE EXTENSION pg_synonym;
CREATE SCHEMA app;
CREATE SCHEMA erp;
CREATE TABLE erp.customers (id bigint PRIMARY KEY, name text);
SELECT synonym.create('app.customer', 'erp.customers');
INSERT INTO app.customer VALUES (1, 'Ada');
SELECT * FROM app.customer;
SELECT * FROM synonym.list;
```

全部辅助接口位于 `synonym` 模式中。

### 操作与权限

`synonym.create` 可以在目标尚不存在时登记别名，随后由事件触发器完成绑定。目标重命名时按对象身份跟随，受支持的列变更会重建别名。`synonym.refresh()` 用于手工协调状态。`synonym.drop` 删除单个别名，可指定 `if_exists => true`。

通过别名访问仍受目标权限和行级安全策略约束。创建别名需要在别名所在模式中创建对象的权限，以及目标模式的使用权限。删除别名需要所有权、相应角色成员身份或超级用户权限。

### 使用边界

目标可以是普通表、分区表、外部表、视图与物化视图。不支持别名指向别名、跨数据库对象、函数或序列。导出数据应通过查询进行；别名不能替代目标对象的 DDL 或截断操作。通过视图执行合并需要 PostgreSQL 17 或更高版本。

手工修改别名视图可能破坏同步。依赖该别名的其他视图可能阻止破坏性的目标变更，重建别名也可能重置直接授予它的权限。删除扩展会保留普通别名视图，需另行安排清理。
