## 用法

来源：

- [上游来源 1](https://github.com/vital-ai/vital-graph/blob/b79dc29b4e4652adc49365404d9272fb7ee4f9b3/pg_extensions/vg_planner/test/sql/inert.sql)
- [上游来源 2](https://github.com/vital-ai/vital-graph/blob/b79dc29b4e4652adc49365404d9272fb7ee4f9b3/pg_extensions/vg_planner/vg_planner.control)
- [上游来源 3](https://github.com/vital-ai/vital-graph/blob/b79dc29b4e4652adc49365404d9272fb7ee4f9b3/pg_extensions/vg_planner/test/sql/advice.sql)

`vg_planner` 0.2 将 VitalGraph 生成的建议应用于 PostgreSQL 连接基数、方法与顺序。

### 核心流程

使用规划建议前，应加入预加载库并重启。生成的 SQL 通过结构化计划注释传递建议，没有建议的普通语句不受影响。`vg_planner_last_report` 说明各项建议是否生效，`vg_planner.enabled` 可作为会话级关闭开关。

```ini
shared_preload_libraries = 'vg_planner'
```

```sql
CREATE EXTENSION vg_planner;
SELECT vg_planner_info();
SHOW vg_planner.enabled;
SELECT vg_planner_last_report();
```

### 运行边界

已核对源码集成了 PostgreSQL 18 专用连接规划代码，安装需要超级用户。格式错误的建议会被拒绝，`vg_planner.on_mismatch` 控制文档规定的不匹配处理方式。请求的方法可能无法使用，因此需检查报告和执行计划。这是面向 VitalGraph 生成器的开发集成，并不承诺兼容任意提示。
