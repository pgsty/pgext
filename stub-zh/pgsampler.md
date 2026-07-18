## 用法

来源：

- [官方扩展控制文件](https://github.com/no0p/pgsampler/blob/9feffce68a902dc9b35fd1ead5ef2c8095e25124/pgsampler.control)
- [官方上游文档](https://github.com/no0p/pgsampler/wiki)
- [官方项目或服务商页面](https://no0p.github.io/pgsampler/)

`pgsampler` — 实验性的 PostgreSQL 后台工作进程，用于采集时序状态与性能指标。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
