## 用法

来源：

- [官方资料](https://github.com/aaaristo/pg_plan_override/blob/e0215040e42047f957d0e301ec40600c2b61ce61/README.md)
- [官方资料](https://github.com/aaaristo/pg_plan_override/blob/e0215040e42047f957d0e301ec40600c2b61ce61/src/pg_plan_override.control)
- [官方资料](https://github.com/aaaristo/pg_plan_override/blob/e0215040e42047f957d0e301ec40600c2b61ce61/src/pg_plan_override--1.0.sql)
- [官方资料](https://github.com/aaaristo/pg_plan_override/blob/e0215040e42047f957d0e301ec40600c2b61ce61/src/pg_plan_override.c)

`pg_plan_override` 1.0 按查询 ID 或 SQL 文本模式应用逐查询规划参数。上游声明支持 PostgreSQL 12 及后续版本，在 PostgreSQL 12 上测试，并为新版规划器钩子签名提供兼容分支；这里未独立验证后续大版本。

### 启用与添加规则

管理员配置预加载并重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_plan_override'
```

由超级用户安装，并创建匹配范围明确的规则：

```sql
CREATE EXTENSION pg_plan_override;
SELECT plan_override.add_by_pattern(
    '%slow_report%',
    '{"enable_seqscan": "off"}'::jsonb,
    'Review report plan'
);
SELECT plan_override.refresh_cache();
SELECT * FROM plan_override.override_rules;
```

`plan_override.add_by_query_id` 可改为匹配数值指纹。多个规则匹配时，采用优先级最高的已启用规则。查询 ID 是否可用取决于 PostgreSQL 版本与指纹配置；上游指出旧大版本需使用语句统计扩展，而新大版本具备原生计算能力。

### 作用范围与缓存

钩子在规划期间临时修改参数，随后恢复，错误路径也会恢复。这些参数不会保留到执行阶段。尤其是，通过规则修改执行阶段的资源参数，并不会为执行中的查询提供相应资源预算。

`pg_plan_override.enabled` 是总开关，`pg_plan_override.debug` 控制覆盖行为日志。`pg_plan_override.cache_ttl` 默认为 60 秒。每个后端各自缓存规则，刷新一个连接不会刷新其他连接；测试新规则时，已经缓存的准备计划还需要单独重新规划。

### 权限与维护

安装会将规则表及其模式的读取权限授予 `PUBLIC`，模式和描述中不能放入凭据。应仅允许可信管理员修改规则；无效或过宽的规则可能使计划退化，模式检查开销也随规则数增长。需通过实际执行计划验证效果，不再有益的规则应禁用或删除。固定模式为 `plan_override`，扩展不可迁移模式。
