## 用法

来源：

- [官方上游 README](https://github.com/sam-harri/pgtensor/blob/969352abd2d425943c46aaac3508c41d0c731a49/README.md)
- [官方扩展控制文件 (pgtensor.control)](https://github.com/sam-harri/pgtensor/blob/969352abd2d425943c46aaac3508c41d0c731a49/pgtensor.control)
- [官方实现源代码](https://github.com/sam-harri/pgtensor/blob/969352abd2d425943c46aaac3508c41d0c731a49/src/lib.rs)

`pgtensor` — 开源的 Postgres 扩展，增加了张量类型，并使用动态工作进程实现了 ONNX 推理引擎。该扩展使用 pgrx 为 Rust 构建。将其用于相应的向量、模型或检索工作流。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pgtensor;

CREATE TABLE t (x tensor(2,3));
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 该扩展的目录记录版本 `0.0.0`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将该扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
