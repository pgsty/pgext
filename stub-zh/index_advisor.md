## 用法

来源：

- [Official source](https://github.com/supabase/index_advisor/blob/b41d1f5f42f8909091edd6752d1f5006d6535db3/README.md)
- [Official source](https://github.com/supabase/index_advisor/blob/b41d1f5f42f8909091edd6752d1f5006d6535db3/index_advisor.control)
- [Official source](https://github.com/supabase/index_advisor/blob/b41d1f5f42f8909091edd6752d1f5006d6535db3/index_advisor--0.2.0--0.2.1.sql)

`index_advisor` 0.2.1 使用 HypoPG 比较假想索引建立前后的规划器成本。它返回建议，不保证实际执行时间，也不会创建建议的物理索引。

### 基本流程

先以所需权限安装依赖 `hypopg`。该 SQL 扩展不需要预加载或重启。

```sql
CREATE EXTENSION hypopg;
CREATE EXTENSION index_advisor;
CREATE TABLE book (id bigint PRIMARY KEY, title text NOT NULL);
SELECT * FROM index_advisor('SELECT id FROM book WHERE title = $1');
```

`index_advisor(query text)` 接受带参数的查询，返回建立索引前后的启动成本、总成本、`index_statements` 和 `errors`。创建索引前应检查建议 DDL 和实际负载。视图及物化视图支持取决于顾问能否解析到底层表。

### 0.2.1 修复与边界

0.2.1 修复 HypoPG 安装在其他模式时的函数限定、注释移除的非贪婪匹配，以及仅清理顾问自身预备语句的行为。安装新 SQL 文件后，执行 `ALTER EXTENSION index_advisor UPDATE TO '0.2.1'` 更新已有数据库。

应以明确的表访问权限和有代表性的统计信息进行分析。规划成本不能覆盖全部运维代价；额外索引消耗磁盘并增加写入开销，上线前应针对实际查询形式验证。
