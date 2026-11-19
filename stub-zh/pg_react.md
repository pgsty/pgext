## 用法

来源：

- [官方资料](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/README.md)
- [官方资料](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/pg_react.control)
- [官方资料](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/LICENSE)
- [官方资料](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/sql/pg_react--0.46.2.sql)
- [官方资料](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/docs/reference/release-identity.md)
- [官方资料](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/docs/getting-started.md)
- [官方资料](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/docs/installation.md)
- [官方资料](https://github.com/trickle-labs/pg-react/blob/fc527241c7ed854066731f13c70f76a8652ccca3/docs/operations.md)

pg_react 将关系条件转换为持久化规则任务。源码 SQL 默认版本为 0.46.2，而官方验证的教学组合仍是 pg_react 0.46.0、pg_trickle 0.108.0、PostgreSQL 18.3 与 Linux amd64。Cargo 版本 0.48.0 及可选的模拟 MDM 适配器不构成更新部署组合已获验证的证据。

### 准备与检查

依次预加载 pg_trickle 与 pg_react 并重启。采用安装指南的触发器 CDC 与关闭 pg_trickle 调度器设置，为每个数据库配置托管工作进程及合适的执行角色。仅创建 SQL 对象不足以形成可运行部署。

```ini
shared_preload_libraries = 'pg_trickle,pg_react'
```

```sql
SELECT extname, extversion FROM pg_extension
WHERE extname IN ('pg_react','pg_trickle') ORDER BY extname;
SELECT pgreact.doctor();
SELECT pgreact_api.managed_status();
```

### 规则流程

创建表示条件的应用视图，并选择稳定的语义键。构造 pgreact.rule() 后执行验证、预览，检查问题，再用已审核的令牌部署。CONSTRAINT 记录匹配；COMMAND 需要带类型的执行函数。以下最小声明假定视图及其 order_id 键已存在。

```sql
SELECT pgreact.validate(pgreact.rule(
  name => 'review-required',
  condition => 'rule_def.review_required'::regclass,
  semantic_key => 'order_id'::name
));
SELECT pgreact.run();
SELECT * FROM pgreact.matches WHERE name = 'review-required';
SELECT * FROM pgreact.work WHERE name = 'review-required';
```

### 执行边界

验证不会部署规则。审核令牌绑定计划与状态，但不授予权限。默认初始化将已有匹配作为基线，不作为新任务处理。源变化需提交后工作进程才能看到；pgreact.run() 协调任务，而不等待全部执行结束。库内效果与执行记录处于同一事务；外部 outbox 投递为至少一次，需要去重。重试或移除前应检查 pgreact.attempts 与操作指南。
