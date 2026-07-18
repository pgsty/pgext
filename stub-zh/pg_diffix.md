## 用法

来源：

- [上游 README](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/README.md)
- [扩展 control 文件](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/pg_diffix.control)
- [管理员指南](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/docs/admin_guide.md)
- [分析人员指南](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/docs/analyst_guide.md)
- [SQL 安装脚本](https://github.com/diffix/pg_diffix/blob/94b4edff1e697de73f0649d7fab64890bdb52bc7/pg_diffix--fir.sql)

`pg_diffix` 为 PostgreSQL 实现动态查询匿名化。管理员把表标记为公开或个人表、指定一个或多个匿名标识列，并为角色分配 `direct`、`anonymized_trusted` 或 `anonymized_untrusted` 访问级别。此后对个人表的查询会受到限制，聚合结果会应用抑制与噪声。

### 最小管理流程

```sql
CREATE EXTENSION pg_diffix;

CALL diffix.mark_personal('public.customers', 'customer_id');
CALL diffix.mark_role('analyst', 'anonymized_untrusted');

ALTER DATABASE appdb
  SET session_preload_libraries TO 'pg_diffix';
```

设置预加载库后重新连接，再检查当前访问级别并执行允许的分组查询：

```sql
SELECT diffix.access_level();

SELECT city, count(*)
FROM public.customers
GROUP BY city;
```

匿名模式有意只允许有限的 SQL 子集。隐私效果取决于完整的表/列标签、合理的匿名标识、恰当的抑制设置、角色隔离，以及对 RLS 策略、用户自定义函数和其他所有已安装扩展的审查。被标为 `direct` 的角色会绕过匿名化。

管理员指南明确说明 `Fir` 代仍处于开发中，并要求 PostgreSQL 13 或更高版本。control 文件与目录使用非数字版本 `fir`，而同一提交的 `META.json` 报告发行版本 `1.2.0`；不要把这些标签当作升级序列。只有在针对实际 schema 和负载完成隐私审查与对抗测试后才应部署。
