## 用法

来源：

- [官方扩展控制文件](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/brainfuck/brainfuck.control)
- [官方上游文档](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/brainfuck/README.md)
- [官方项目或服务商页面](https://database.dev/mansueli/brainfuck)

`brainfuck` — 使用 PLV8 可信语言扩展在 PostgreSQL 中运行 Brainfuck 程序。

已复核目录快照记录的版本为 `4.2.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plv8`。
整理后的兼容版本集合为 `15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "brainfuck";
SELECT extversion
FROM pg_extension
WHERE extname = 'brainfuck';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
