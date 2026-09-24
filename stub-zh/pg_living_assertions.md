## 用法

来源：

- [PGXN 0.4.2](https://pgxn.org/dist/pg_living_assertions/0.4.2/)

`pg_living_assertions` 保存 SQL 检查、结果、核验时间和定义替换历史。它是无需预加载的纯 SQL 扩展，本包支持 PostgreSQL 14–18。发行包 0.4.2 保留 SQL 扩展版本 0.4.1。

### 登记与核验

```sql
CREATE EXTENSION pg_living_assertions;
SELECT living_assertions.declare(
  'simple_check', 'one equals one',
  $$SELECT 1 = 1 AS holds, 'arithmetic check'::text AS detail$$);
SELECT living_assertions.run('simple_check');
SELECT name, state, age FROM living_assertions.status;
```

检查必须返回一行，其中包含布尔结果和可选的详情字符串。STABLE 求值器拒绝写入。检查按请求运行，并非在每次数据变化时执行的 SQL ASSERTION 约束。

### 结果与历史

`living_assertions.run_all()` 执行已登记检查。`living_assertions.state()` 区分成立、失效、未知、检查报错、未检查、已停用和未登记。`living_assertions.stale()` 查找过期结果，并单独标记从未运行的检查。

只有受信任的管理员才应登记或修改检查 SQL，因为后续调用者会以自己的权限执行这些 SQL。定义通过填写原因来替换，不会静默改写；判断成功与否时也应考虑结果年龄。数据库转储包含登记信息及已有结果。
