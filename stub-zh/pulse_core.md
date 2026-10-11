## 用法

来源：

- [README.md](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/README.md)
- [extensions/pulse_core/pulse_core--0.1.0.sql](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/pulse_core--0.1.0.sql)
- [extensions/pulse_core/src/pulse_core.c](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/src/pulse_core.c)
- [extensions/pulse_core/src/sqlapi.c](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/src/sqlapi.c)
- [extensions/pulse_core/test/pulse_core.conf](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/test/pulse_core.conf)
- [extensions/pulse_core/pulse_core.control](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/pulse_core.control)

`pulse_core` 是 pg_pulse 原型中已经实现的查询统计模块。执行器钩子在共享内存中收集各查询的调用次数、耗时、行数和最近出现时间。

### 核心用法

```ini
shared_preload_libraries = 'pulse_core'
compute_query_id = on
```

```sql
CREATE EXTENSION pulse_core;
SELECT * FROM pulse.queries() ORDER BY total_ms DESC;
```

### 运行边界

项目目标为 PostgreSQL 18。管理员安装前需要将 `pulse_core` 加入 `shared_preload_libraries` 并重启，后期加载会被拒绝。应按上游回归测试配置启用 `compute_query_id`。`pulse.max_entries` 是启动时设置，默认 1000；`pulse.track` 接受 none、top 或 all，默认 top。`pulse.queries()` 返回当前统计，`pulse.reset()` 清空统计。两个函数都撤销了 PUBLIC 权限并授权给 `pg_monitor`，因此该角色成员也能重置统计，应审查成员范围。共享内存计数不构成持久历史。独立的健康工作进程组件仍是骨架，不在本条目中宣称其能力。
