## 用法

来源：

- [官方 README](https://github.com/foxflow/pg_zen_engine/blob/e4924af4f5a0ad7da3d0031ee5e54ef6698db0b8/README.md)
- [Rust 实现](https://github.com/foxflow/pg_zen_engine/blob/e4924af4f5a0ad7da3d0031ee5e54ef6698db0b8/src/lib.rs)
- [扩展控制文件](https://github.com/foxflow/pg_zen_engine/blob/e4924af4f5a0ad7da3d0031ee5e54ef6698db0b8/pg_zen_engine.control)

`pg_zen_engine` 在 PostgreSQL 内针对 `jsonb` 输入求值 GoRules JSON Decision Model 图。它提供一个 Rust 实现的函数，适合把同一个已存储决策图应用到关系数据附近。

### 核心流程

将有效的 JDM 图保存为 `jsonb`，把每行输入文档传给 `evaluate_jdm()`，再从返回值中提取引擎结果。

```sql
CREATE EXTENSION pg_zen_engine;

SELECT name,
       evaluate_jdm(graph, jsonb_build_object('input', facts)) AS decision
FROM decision_models
JOIN decision_inputs USING (name);
```

`evaluate_jdm(graph jsonb, data jsonb) RETURNS jsonb` 被声明为 strict、immutable 和 parallel-safe。每次调用都会反序列化图、创建新的决策引擎、求值输入并序列化完整结果；扩展不提供模型缓存、管理表或验证函数。

### 运维边界

控制文件规定安装扩展需要超级用户权限。无效图 JSON、不受支持的模型内容或求值失败会到达使用 `unwrap()` 的 Rust 路径，因此应在关键查询之外验证模型，并在目标构建上演练失败情况。应把模型表达式视为可执行业务逻辑，并限制可修改已存储图的角色。

所审阅的 `0.0.1` crate 固定使用 Zen Engine `0.22.0`，并声明 PostgreSQL 12 到 16 的构建特性。即使 SQL 签名不变，决策模型语义也可能随引擎依赖变化；应固定完整构建、保留带预期结果的模型样例，并在对大规模扫描应用前衡量重复求值成本。
