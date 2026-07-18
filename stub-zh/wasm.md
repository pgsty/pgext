## 用法

来源：

- [上游 README](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/README.md)
- [扩展 control 文件](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/src/wasm.control)
- [0.1.0 版本安装 SQL](https://github.com/wasmerio/wasmer-postgres/blob/c36ceeb3b4638016896ea0a419a45d7a3f215384/src/wasm--0.1.0.sql)

`wasm` 嵌入 Wasmer runtime，使 PostgreSQL 能够实例化 WebAssembly module，并把导出函数暴露为动态生成的 SQL 函数。上游 `0.1.0` 版本属于实验性实现，仅支持 PostgreSQL 10；README 明确说明不支持 PostgreSQL 11。

编译动态库与 PL/pgSQL 扩展后，使用服务器端动态库的绝对路径进行初始化，再装载 WebAssembly module：

```sql
CREATE EXTENSION wasm;
SELECT wasm_init('/absolute/path/to/libwasm.so');
SELECT wasm_new_instance('/absolute/path/to/simple.wasm', 'ns');
SELECT ns_sum(1, 2);
```

namespace 参数会作为生成函数的前缀；以上示例会创建 `ns_sum`。该原型把 WebAssembly `i32`、`i64` 和 `v128` 值映射到 PostgreSQL 类型，浮点支持尚不完整。

`wasm_init()` 会动态创建 C 语言绑定和内部 FDW。其安装 SQL 使用 `CASCADE` 删除并重建内部 FDW、server 和 `wasm` schema；如果数据库已在该 schema 下保存对象，未经影响评估不要重复执行。module 与动态库路径必须是 PostgreSQL 服务器进程可见的绝对路径。上游没有要求预加载。
