## 用法

来源：

- [PostgreSQL 扩展 README](https://github.com/benseverndev-oss/goldenmatch/blob/ec6ea0e4834fdd478961ff5808be2facc43740e2/packages/rust/extensions/postgres/README.md)
- [扩展 control 文件](https://github.com/benseverndev-oss/goldenmatch/blob/ec6ea0e4834fdd478961ff5808be2facc43740e2/packages/rust/extensions/postgres/goldenmatch_pg.control)
- [Cargo 软件包元数据](https://github.com/benseverndev-oss/goldenmatch/blob/ec6ea0e4834fdd478961ff5808be2facc43740e2/packages/rust/extensions/postgres/Cargo.toml)
- [PostgreSQL 桥接源码](https://github.com/benseverndev-oss/goldenmatch/tree/ec6ea0e4834fdd478961ff5808be2facc43740e2/packages/rust/extensions/postgres/src)

`goldenmatch_pg` `0.13.0` 版在固定的 `goldenmatch` 模式中提供 GoldenMatch 实体解析、去重、记录匹配、分析、验证、纠正和流水线辅助函数。部分标量与图算法内核使用原生 Rust；范围更广的操作会嵌入 CPython，并调用服务器安装的 `goldenmatch` 软件包。

### 从安全标量开始

```sql
CREATE EXTENSION goldenmatch_pg;
SELECT goldenmatch.goldenmatch_score(
  'John Smith', 'Jon Smyth', 'jaro_winkler'
);
```

文档化运行环境要求 PostgreSQL 15–17、Python 3.11+，以及安装在 PostgreSQL 所链接解释器中的兼容 `goldenmatch` Python 软件包。表/作业函数会读取命名关系并创建或删除结果状态；身份/纠正 API 可以接受服务器端数据库或内存文件路径。匹配分数是启发式结果，不是身份凭证。应限制运维型及文件型函数，审计动态标识符和生成对象，固定 Rust 与 Python 依赖，并用有标签数据和人工复核验证阈值。
