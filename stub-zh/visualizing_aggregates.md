## 用法

来源：

- [上游 README](https://github.com/shortdiv/pg-visualize-aggregates/blob/659e6de3f404d8a43cdc92e41d7ddb9b74654521/README.md)
- [Rust 扩展入口](https://github.com/shortdiv/pg-visualize-aggregates/blob/659e6de3f404d8a43cdc92e41d7ddb9b74654521/src/lib.rs)
- [图形渲染器](https://github.com/shortdiv/pg-visualize-aggregates/blob/659e6de3f404d8a43cdc92e41d7ddb9b74654521/src/graph.rs)
- [版本与 PostgreSQL 特性矩阵](https://github.com/shortdiv/pg-visualize-aggregates/blob/659e6de3f404d8a43cdc92e41d7ddb9b74654521/Cargo.toml)

visualizing_aggregates 是一个尚未完成的 pgrx 原型，意在渲染聚合数据并创建 CodePen 可视化。0.0.0 版公开 hello_visualizing_aggregates()，并计划公开 draw_graph(text)。如不修复源码，它在已审查的上游提交中无法使用。

### 预期冒烟测试

修复并成功构建扩展后，可用以下语句检查其唯一不带副作用的 SQL 入口：

```sql
CREATE EXTENSION visualizing_aggregates;
SELECT hello_visualizing_aggregates();
```

预期的 draw_graph 函数会忽略 dataset 参数，从硬编码 climbs 表查询 route_name，记录选中值，渲染固定标记，再向 CodePen 发起阻塞式出站 HTTP POST。它目前并不会可视化调用者选择的聚合。

### 注意事项

- 已审查的提交无法编译：lib.rs 引入不存在的 bar_graph_template.html 文件，仓库树中却是 svg_template.html；它还使用三个参数调用 draw_svg，而 graph.rs 只定义了单参数方法。
- 该函数在后端代码中对数据库、模板、序列化和网络结果直接解包。表不存在、结果为空、渲染失败或远程服务失败都可能引发错误或 panic。
- 调用 draw_graph 会从 PostgreSQL 后端发起出站网络流量。它可能阻塞后端、依赖第三方可用性，并跨越安全边界；应限制执行权与数据库主机出站访问。
- 源码会记录硬编码查询与取回的 route_name。应将服务器日志和外部请求视为潜在数据泄露路径。
- 控制文件要求超级用户，并将扩展标记为不可信。上游没有提供运维文档、正式发布、许可证、除问候函数外的有意义测试，或生产兼容性声明。
