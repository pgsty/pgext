## 用法

来源：

- [官方上游 README](https://github.com/ricarim/secure-crdv/blob/3c224559b6c16bd850978845cf172e9fbcae6f00/README.md)
- [官方扩展控制文件 (crdv_pg_coordinator.control)](https://github.com/ricarim/secure-crdv/blob/3c224559b6c16bd850978845cf172e9fbcae6f00/mpc/crates/crdv_pg_coordinator/crdv_pg_coordinator.control)
- [官方实现源代码](https://github.com/ricarim/secure-crdv/blob/3c224559b6c16bd850978845cf172e9fbcae6f00/mpc/crates/crdv_pg_coordinator/src/lib.rs)

`crdv_pg_coordinator` 提供由多方安全计算（MPC）保护私密值的无冲突复制数据视图。它适用于相应的安全、审计或访问控制工作流。请以上方链接的固定上游修订为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION crdv_pg_coordinator;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `mpc_benchmark_set` 是一个扩展函数。
- `mpc_counter_dec` 是一个扩展函数。
- `mpc_counter_inc` 是一个扩展函数。
- `mpc_declassify` 是一个扩展函数。
- `mpc_input_secret_i64` 是一个扩展函数。
- `mpc_open` 是一个扩展函数。
- `mpc_reconfigure` 是一个扩展函数。
- `mpc_reset_party_state` 是一个扩展函数。
- `mpc_reshare` 是一个扩展函数。
- `mpc_set` 是一个扩展函数。
- `mpc_set_add` 是一个扩展函数。
- `mpc_set_contains` 是一个扩展函数。
- `mpc_set_coordinator_quorum` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与所链接的源代码中的固定版本进行比对。
