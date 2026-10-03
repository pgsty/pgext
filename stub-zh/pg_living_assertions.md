## 用法

来源：

- [README.md](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/README.md)
- [pg_living_assertions.control](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/pg_living_assertions.control)
- [pg_living_assertions--0.4.1--0.5.0.sql](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/pg_living_assertions--0.4.1--0.5.0.sql)
- [pg_living_assertions--0.5.0--0.5.1.sql](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/pg_living_assertions--0.5.0--0.5.1.sql)
- [test/sql/read_only.sql](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/test/sql/read_only.sql)

`pg_living_assertions` 0.5.1 保存 SQL 检查、结论、核验时间与替换历史。检查按需运行，并非每次写入都求值的 SQL ASSERTION 约束。此扩展为纯 SQL 实现，无需预加载。

### 注册与核验

```sql
CREATE EXTENSION pg_living_assertions;
SELECT living_assertions.declare(
  'simple_check', 'one equals one',
  $$SELECT 1 = 1 AS holds, 'arithmetic check'::text AS detail$$);
SELECT living_assertions.run('simple_check');
SELECT name, state, age FROM living_assertions.status;
```

### 结果与历史

每个检查必须返回恰好一行，包含布尔列 `holds` 和可选文本列 `detail`。`living_assertions.run_all()` 执行已注册检查。`living_assertions.state()` 区分成立、失效、未知、报错、未检查、已退役与未注册；`living_assertions.stale()` 区分从未检查与结果过期。`living_assertions.declare_unchanged()` 保存表达式以供后续文本比较，因此作者必须自行规范化输出。定义通过附带原因的替换保留历史，结果与注册表数据包含在数据库备份中。

### 执行与权限

从 0.5.0 起，求值器以只读方式在始终回滚的子事务中执行，并保留检查结论。这修复了旧求值器仅依赖 STABLE、无法阻止易变函数副作用的问题。它不是不可信 SQL 的沙箱：临时序列变更、会话级咨询锁与外部副作用仍可能保留。仅应授权可信管理员注册检查；检查使用后续调用者的权限执行。注册表属于扩展所有者，写入函数默认撤销 `PUBLIC` 执行权限。

### 升级

安装匹配的文件后，执行 `ALTER EXTENSION pg_living_assertions UPDATE TO '0.5.1'`。0.4.1→0.5.0→0.5.1 升级链替换求值函数，不修改注册表结构。最后的补丁在 `run()` 中限定行类型名称，防止类型缓存失效后在断言自身的无关搜索路径中重新解析类型。

全新安装也使用较早的基础 SQL 脚本并依次应用包内升级链，因此必须安装完整且版本匹配的脚本集。
