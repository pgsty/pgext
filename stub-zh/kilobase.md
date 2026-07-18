## 用法

来源：

- [官方扩展控制文件](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/kilobase.control)
- [官方上游文档](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/README.md)
- [官方项目或服务商页面](https://kbve.com/)

`kilobase` — 用于调度并刷新 PostgreSQL 物化视图的 pgrx 预加载后台工作进程。

已复核目录快照记录的版本为 `17.4.1`、类型为 `preload`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "kilobase";
SELECT extversion
FROM pg_extension
WHERE extname = 'kilobase';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
