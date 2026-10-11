## 用法

来源：

- [README.md](https://github.com/dnewcome/pg-bitemporal-ext/blob/60ec92d9a8daee5dd0ea3339f0fbc067cafc672f/README.md)
- [sql/pg_bitemporal--0.1.sql](https://github.com/dnewcome/pg-bitemporal-ext/blob/60ec92d9a8daee5dd0ea3339f0fbc067cafc672f/sql/pg_bitemporal--0.1.sql)
- [pg_bitemporal.control](https://github.com/dnewcome/pg-bitemporal-ext/blob/60ec92d9a8daee5dd0ea3339f0fbc067cafc672f/pg_bitemporal.control)

`pg_bitemporal` 将区间关系函数与保存系统时间历史的触发器结合起来。完整模式还保存由应用提供的有效时间，单系统时间模式则记录数据库中的变更。

### 核心用法

```sql
CREATE EXTENSION btree_gist;
CREATE EXTENSION pg_bitemporal;
CREATE TABLE employee_history_demo (id integer NOT NULL, department text);
SELECT bitemporal.enable('employee_history_demo', ARRAY['id'], p_mode => 'system_only');
INSERT INTO employee_history_demo VALUES (1, 'Engineering');
SELECT id, department, sys_start, sys_end FROM employee_history_demo;
SELECT allen_relation('[1,4)'::int4range, '[3,8)'::int4range);
```

### 运行边界

README 要求 PostgreSQL 14 或更新版本及 `btree_gist`，未提供完整的主版本测试矩阵。原生函数需要管理员安装。`bitemporal.enable` 会修改目标表、增加时间列与索引、建立排斥约束和触发器，因此需要相应的表权限。SQL 中实际的参数名是 `p_mode`。更新会关闭旧版本并插入新版本，删除则关闭当前版本而不物理删除。用 `sys_end = 'infinity'` 查询当前数据；历史行仍保存在表中，需要考虑存储增长。完整模式增加 `valid_from`、`valid_to` 和 `valid_range`，由应用提供有效时间。`allen_relation`、`allen_relation_code` 及十三个判断函数用于区间关系查询。虽然 control 声明可迁移，扩展 SQL 仍显式创建固定的 `bitemporal` 模式。
