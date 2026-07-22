## 用法

来源：

- [上游 README](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/README.md)
- [扩展控制文件](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/src/wasm.control)
- [版本 0.1.0 安装 SQL](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/src/wasm--0.1.0.sql)

`wasm` 版本 `0.1.0` 内嵌 Wasmer 运行时，使 PostgreSQL 能实例化服务器侧 WebAssembly 文件，并把导出函数暴露为动态生成的 SQL 函数。上游称该版本为实验性版本，并且只记录了 PostgreSQL 10 支持。

### 核心流程

项目分别安装 PL/pgSQL 扩展与动态库。使用 PostgreSQL 服务器可见的绝对路径完成初始化，然后在一个 SQL 函数命名空间下实例化模块。

```sql
CREATE EXTENSION wasm;
SELECT wasm_init('/absolute/path/to/libwasm.so');
SELECT wasm_new_instance('/absolute/path/to/simple.wasm', 'ns');
SELECT ns_sum(1, 2);
```

`wasm_new_instance` 返回实例 ID，并为每个受支持的导出创建一个以命名空间为前缀的函数。该原型最多支持十个参数。整数输入输出映射自 WebAssembly `i32` 与 `i64`；虽然自省会把浮点签名映射为 numeric 类型，但浮点调用仍未完整实现。

### 自省与安全

`wasm_init` 创建内部 FDW，以及用于查看当前进程实例与导出的 `wasm.instances`、`wasm.exported_functions` 外部表。模块路径由后端进程读取，已加载实例保存在后端内存中，而不是持久数据库存储。

初始化会通过 `CASCADE` 删除并重建内部包装器、服务器与 `wasm` 模式；重复运行可能销毁该模式中的对象。它还会根据传入的动态库路径动态创建 C 语言函数。应只允许可信管理员执行，并审核模块来源。没有文档要求预加载，但 PostgreSQL 11 及更高版本明确不在本次核对的兼容性声明中；只能在隔离环境中测试崩溃、资源耗尽、ABI 兼容性与后端重启行为。
