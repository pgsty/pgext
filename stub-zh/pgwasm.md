## 用法

来源：

- [指定修订版的 pgwasm README](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/README.md)
- [pgwasm 架构与 SQL 生命周期](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/docs/architecture.md)
- [pgwasm GUC 参考](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/docs/guc.md)
- [pgwasm WIT 类型映射](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/docs/wit-mapping.md)
- [pgwasm 控制文件](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/pgwasm/pgwasm.control)

`pgwasm` 把 WebAssembly component 加载到 PostgreSQL，并把 WIT export 注册为带类型的 PostgreSQL 函数。编译产物保存在集群数据目录下，并由后端本地实例池复用。本文基于固定修订 `535b53363f8208af139e757e508e66c46309ee29`；源码声明版本 0.1.0，但没有提供带标签的 0.1.0 发行版。

### 核心流程

由超级用户创建扩展。使用以下文件加载流程前，管理员必须显式启用并限制其目录：

```sql
CREATE EXTENSION pgwasm;

ALTER SYSTEM SET pgwasm.allow_load_from_file = on;
ALTER SYSTEM SET pgwasm.module_path = '/srv/pgwasm';
ALTER SYSTEM SET pgwasm.allowed_path_prefixes = '/srv/pgwasm';
SELECT pg_reload_conf();

GRANT pgwasm_loader TO app_runtime;

SELECT pgwasm.pgwasm_load(
    'arith',
    '{"path":"arith.component.wasm"}'::json,
    '{}'::json
);

SELECT * FROM pgwasm.pgwasm_functions();
SELECT * FROM pgwasm.pgwasm_modules();

SELECT pgwasm.pgwasm_unload('arith');
```

`pgwasm_load(module_name text, bytes_or_path json, options json)` 只接受一个 `bytes` 或 `path` 来源。文件加载默认关闭。模块名会成为持久 catalog 键，也是经过清理后生成的 SQL 函数名前缀。

### 生命周期与类型映射

- `pgwasm_load` 执行校验、解析策略、创建所需 PostgreSQL 类型和函数、编译 AOT 产物并记录模块。
- `pgwasm_reload` 替换模块字节；签名兼容时会保留稳定标识。
- `pgwasm_reconfigure` 收窄或修改策略与资源限制。
- `pgwasm_unload` 删除生成函数、类型、catalog 行与产物；存在依赖时会阻止删除，除非显式选择级联。
- WIT record 映射为 composite type，enum 映射为 PostgreSQL enum，list 映射为数组或 `bytea`，受支持的 variant、flag、option、result 与 resource 则映射为文档规定的 PostgreSQL 表示。
- `pgwasm_modules()`、`pgwasm_functions()`、`pgwasm_wit_types()`、`pgwasm_policy_effective()` 与 `pgwasm_stats()` 用于检查。

授予执行权限或调用新加载 component 前，应先检查生成函数签名。涉及不兼容 WIT 变化的 reload 需要显式策略决策与依赖复核。

### 沙箱与权限

扩展创建 `pgwasm_loader` 用于生命周期变更，并创建 `pgwasm_reader` 用于可观测性。加载、重载、重配置和卸载要求超级用户或 loader 角色成员身份。

WASI 文件系统、环境变量、socket、HTTP 与 SPI host-query 访问默认全部关闭。管理员通过 `pgwasm.*` GUC 设置集群能力上限；每个模块的选项只能收窄，不能扩大该上限。应明确并尽量缩小 `pgwasm.allowed_hosts`、路径前缀与文件系统预开放范围。

### 资源与运维边界

- 默认模块大小限制是 32 MiB，调用内存为 1,024 个 WebAssembly page，墙钟时间限制为 5 秒。可以启用 fuel 计量，但默认关闭。
- `$PGDATA/pgwasm/<module_id>/` 下的产物由模块字节与 Wasmtime 构建生成。遇到不兼容的 Wasmtime 或 PostgreSQL 升级时应重新编译，不能把这些产物当作权威数据直接复制。
- 共享计数器依赖 postmaster 启动时分配共享内存。需要共享指标时应预加载 `pgwasm`；否则可观测性会退化到非共享计数器，并报告该状态。
- 源码提供 PostgreSQL 13 到 18 的构建 feature，默认使用 PostgreSQL 17，但该固定修订没有公开的支持矩阵。部署前应验证确切 PostgreSQL 大版本构建与全部所需 WIT 映射。
- 即使有沙箱，也应把 guest 代码视为靠近数据库的特权代码：限制模块加载者、约束每项能力与资源，并测试 trap、取消、reload、重启和回滚行为。
