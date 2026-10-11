## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/README.md)
- [pg_living_assertions.control](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/pg_living_assertions.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/CHANGELOG.md)
- [pg_living_assertions--0.5.9--0.5.10.sql](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/pg_living_assertions--0.5.9--0.5.10.sql)

`pg_living_assertions` 0.5.10 保存具名 SQL 断言、结果与替换历史。检查按需运行，并非每次写入都执行的约束。这是无需预加载的纯 SQL 扩展。

### 核心用法

```sql
CREATE EXTENSION pg_living_assertions;
SELECT living_assertions.declare(
  'simple_check', 'one equals one',
  $$SELECT 1 = 1 AS holds, 'arithmetic check'::text AS detail$$);
SELECT living_assertions.run('simple_check');
SELECT name, state, age FROM living_assertions.status;
```

### 运行边界

检查返回一个布尔列 `holds` 和可选文本列 `detail`。`living_assertions.run_all()` 执行注册表中的检查；`living_assertions.state()`、`living_assertions.status` 和 `living_assertions.stale()` 区分失效、未知、报错，以及过期或未执行的检查。`living_assertions.declare_unchanged()` 对规范化表达式计算指纹，原有值变成 NULL 时判为失效。检查以记录的作者身份，在只读且回滚的子事务中执行，采用作者搜索路径并将临时模式放在最后。调用者须有权 `SET ROLE` 到该作者。求值时取得的会话咨询锁会被释放，但外部副作用和临时序列变化不在沙箱控制范围内，只有可信作者应登记 SQL。注册和检查时间由服务器确定，表拒绝 TRUNCATE，退役与替换仅限作者或有权代理的角色。升级须安装完整匹配的 SQL 链；旧断言保留原有指纹算法。当前版本采用 Apache-2.0，截至 0.5.3 的已发布版本仍保留原来的 PostgreSQL License。

### 0.5.10 的角色隔离

0.5.10 在需要角色隔离时，将检查放入由记录作者拥有的临时 `SECURITY DEFINER` 函数。PostgreSQL 会拒绝在该执行环境中重置角色或会话授权；函数与只读检查在同一子事务中创建并回滚。需要该执行环境时，作者必须拥有数据库的 `TEMP` 权限，调用者仍须有权代理该作者。`SECURITY DEFINER` 调用者现可凭此权限执行其他作者的检查。这修复了 0.5.8/0.5.9 中重置角色的逃逸路径，但不会使外部副作用变得安全。
