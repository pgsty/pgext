## 用法

来源：

- [PGXN 上的 pg_relation_sql 0.2.2](https://pgxn.org/dist/pg_relation_sql/0.2.2/)
- [pg_relation_sql 0.2.2 README](https://api.pgxn.org/src/pg_relation_sql/pg_relation_sql-0.2.2/README.md)
- [pg_relation_sql 0.2.2 SQL 脚本](https://api.pgxn.org/src/pg_relation_sql/pg_relation_sql-0.2.2/relation_sql.sql)
- [pg_relation_sql 0.2.2 执行计划对比](https://api.pgxn.org/src/pg_relation_sql/pg_relation_sql-0.2.2/EXPLAIN.md)

`pg_relation_sql` 0.2.2 根据 PostgreSQL 外键生成成对的 SQL 函数：lookup 函数沿外键找到被引用行，list 函数则返回反向引用当前行的记录。生成的 `LANGUAGE sql` 函数被设计为可由优化器内联，使查询可以沿已声明的关系导航，而无需反复书写连接条件。

上游有意只发布一个独立的 `relation_sql.sql` 文件，没有 control 文件。因此不存在 `CREATE EXTENSION pg_relation_sql`；需要在每个使用这些函数的数据库中执行软件包提供的脚本。

```bash
psql app -f /usr/pgsql-17/share/pg_relation_sql/relation_sql.sql
psql app -f /usr/share/postgresql/17/pg_relation_sql/relation_sql.sql
```

脚本会在当前模式中创建 `relation_sql(text)`，最后请求执行 `relation_sql('install')`。

### 生成并使用关系函数

```sql
CREATE TABLE profile (
  id bigint PRIMARY KEY,
  name text
);

CREATE TABLE address (
  id bigint PRIMARY KEY,
  profile_id bigint REFERENCES profile(id),
  city text
);

SELECT status, command FROM relation_sql('sync');

SELECT a.city, p.name
FROM address AS a, profile(a) AS p;

SELECT p.name, a.city
FROM profile AS p, address_list(p) AS a;
```

每个外键都会得到一个沿引用方向查询的 lookup 函数，以及一个通常带 `_list` 后缀的反向函数；一对一外键除外。复合外键、跨模式外键均受支持，指向同一目标的多个外键会获得带角色前缀的名称。

### 生成器模式

- `relation_sql()` 返回状态面板。
- `relation_sql('show')` 显示计算出的函数及可直接执行的同步命令，但不修改对象。
- `relation_sql('sync')` 根据当前外键创建、替换或删除带标记的关系函数。
- `relation_sql('install')` 添加 `ddl_command_end` 事件触发器并立即同步。
- `relation_sql('uninstall')` 删除事件触发器；`relation_sql('drop')` 删除生成的函数。

### 运维边界

- 创建事件触发器需要超级用户权限。权限不足时会产生警告，但一次性同步仍会使用调用者已有的对象权限执行。
- 应把生成器安装在 `search_path` 受控的可信模式中：自动模式会创建一个保留安装时搜索路径的 `SECURITY DEFINER` 事件触发器辅助函数。
- 生成的函数依赖表的行类型。删除被这些函数用作行类型的表时可能需要 `CASCADE`；执行破坏性 DDL 前应检查依赖关系。
- 生成的函数体使用 `SELECT *`，因此不能很好地配合列级 `SELECT` 授权；行级安全仍会生效。
- 对执行计划敏感的查询应把关系函数写在 `FROM` 中。选择列表中的属性记法会变成 `ProjectSet`，而 `NOT EXISTS (SELECT FROM relation_function(row))` 可能仍是逐行执行的相关子计划，而不是等价的反连接。
- 查询对生成函数的依赖与对视图的依赖相同。不使用事件触发器时，应在迁移流程中运行 `relation_sql('sync')`。
- 上游要求 PostgreSQL 11 或更高版本；Pigsty 软件包覆盖 PostgreSQL 14–18。
