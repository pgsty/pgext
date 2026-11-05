## 用法

来源：

- [上游来源 1](https://github.com/veridit/jsonb_stats/blob/97d7b3dd5997a0d1b2ae5d39357f4e3cb7a797d3/README.md)
- [上游来源 2](https://github.com/veridit/jsonb_stats/blob/97d7b3dd5997a0d1b2ae5d39357f4e3cb7a797d3/jsonb_stats.control)

`jsonb_stats` 3.2.0 为动态命名的 JSONB 变量生成可合并的统计摘要。

### 核心流程

`stat` 构造带类型的值，`jsonb_stats_agg` 聚合统计信息，`jsonb_stats_merge_agg` 合并已有摘要。数值状态包含数量、总和、平方和及最小最大值，3.x 表示采用精确算术，避免依赖合并顺序的浮点累加。

```sql
CREATE EXTENSION jsonb_stats;
SELECT stat(42::integer);
```

### 运行边界

此版本的构建特性面向 PostgreSQL 18。控制文件允许非超级用户安装，但安装角色仍需创建相关对象的数据库和模式权限。升级时应安装匹配文件，重新连接以加载新库，再执行 ALTER EXTENSION UPDATE。不同大版本的持久化摘要不可直接混用，应遵循上游迁移说明。
