## 用法

来源：

- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/README.md)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate_parser/percolate_parser.control)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate_parser/Cargo.toml)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate_parser/sql/surface.sql)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate_parser/sql/percolate_parser--0.2.29--0.2.30.sql)

`percolate_parser` 0.2.30 包含 Percolate 使用的编译型查询解析、YAML 工作流与图处理函数。这个 Rust 扩展当前声明支持 PostgreSQL 19，使用 pgrx 0.19.2，将原生函数与非超级用户持有的 SQL 应用分离。

### 安装与调用

```sql
CREATE EXTENSION percolate_parser;
SELECT extname, extversion FROM pg_extension WHERE extname = 'percolate_parser';
```

`aiq_parse(query)` 将应用查询方言解析为 JSONB；`p8_compile_workflow(yaml)` 将工作流编译为 JSONB。调用者必须提供合法的 Percolate 语法，不能假定任意 SQL 或 YAML 均可接受。匹配的 SQL 应用负责外围执行、身份与授权流程。

### 图接口与限制

SQL 接口还包含 `p8_graph_components`、`p8_graph_communities`、`p8_graph_distance` 和带预算的遍历函数。这些图操作依赖应用的图表和上下文，单独安装解析器并不能形成完整图服务。预算参数用于限制工作量，不构成普遍的延迟保证。

先由超级用户安装解析器，再以应用所有者安装 `percolate`。解析器不可迁移模式，没有声明预加载要求。解析器与应用版本应保持一致，生成的安装 SQL、版本化升级 SQL 必须与原生库匹配；仅复制共享库并不是完整升级。
