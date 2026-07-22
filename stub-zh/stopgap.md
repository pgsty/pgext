## 用法

来源：

- [官方 README](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/README.md)
- [官方 Rust 软件包清单](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/stopgap/Cargo.toml)
- [官方扩展控制文件](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/stopgap/stopgap.control)

`stopgap` 为 TypeScript/JavaScript 函数包提供版本化部署和调用层，这些代码由配套的 `plts` 过程语言执行。它把环境、部署历史、激活、差异和回滚保存在 PostgreSQL 中，同时通过 `api.coolApi.myFn` 这样的路径定位应用代码。

### 安装与调用

两个扩展都不可信，只能由超级用户安装。先安装 `plts`，再安装 `stopgap`：

```sql
CREATE EXTENSION plts;
CREATE EXTENSION stopgap;

SELECT stopgap.call_fn(
  'api.coolApi.myFn',
  '{"id": 1}'::jsonb
);
```

当前上游方向是在项目本地 `stopgap/` 目录下维护具名 TypeScript 导出，并推荐使用 CLI 获得部署验证和操作便利性：

```bash
stopgap deploy --db "$STOPGAP_DB" --env prod --from-schema app --label v1.0
stopgap status --db "$STOPGAP_DB" --env prod
stopgap diff --db "$STOPGAP_DB" --env prod --from-schema app
stopgap rollback --db "$STOPGAP_DB" --env prod --steps 1
```

### 运行时与管理接口

- `stopgap.call_fn(path, args)` 解析已部署的具名导出，并把非空 JavaScript 值作为 `jsonb` 返回；JavaScript `undefined` 与 `null` 会变成 SQL `NULL`。
- `stopgap.status`、`stopgap.deployments`、`stopgap.diff` 和 `stopgap.rollback` 用于检查和管理环境状态。
- `stopgap.query` 处理函数使用只读数据库上下文并拒绝 `db.exec`；`stopgap.mutation` 和普通 `plts` 处理函数使用可读写上下文。
- 处理函数上下文发起的数据库操作与调用它的 SQL 语句运行在同一个 PostgreSQL 事务中。

### 安全与变更边界

- TypeScript/JavaScript 在数据库信任边界内运行。执行应用代码前，应审查包来源、依赖锁定、SQL 权限、模式访问、资源限制、错误处理和可观测性。
- 部署、激活和回滚权限应与普通调用权限分离。应把环境变更视为数据库发布，进行审查、备份、冒烟测试，并准备经过验证的回滚目标。
- 函数可以在调用事务期间持锁并消耗后端 CPU 或内存。应设置超时，不要在热点事务中执行网络或长耗时任务。
- 当前审阅的 README 说明产品在 2026 年 3 月进行方向调整，并把若干 SQL 签名标为迁移期间的目标形态。应固定版本 `0.1.3`、验证已安装签名，不能只依据路线图文字进行自动化。
- 编译器与运行时语义还取决于已安装的 `plts` 版本。应把 TypeScript 类型检查、模块解析、序列化、事务回滚和升级作为一个兼容单元测试。
