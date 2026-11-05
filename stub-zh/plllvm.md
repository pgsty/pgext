## 用法

来源：

- [上游来源 1](https://github.com/matejformanek/plllvm/blob/9e439273ae695e74a13aba391636cf4df1ac79d3/README.md)
- [上游来源 2](https://github.com/matejformanek/plllvm/blob/9e439273ae695e74a13aba391636cf4df1ac79d3/plllvm.control)

`plllvm` 0.3 将非受信任 LLVM IR 作为 PostgreSQL 函数执行，并提供实验性 PL/pgSQL JIT 语言。

### 核心流程

函数在每个后端中单独编译。入口点返回 PostgreSQL Datum，非严格函数需显式处理空值。`plpgsql_jit` 接受受支持的 PL/pgSQL 函数体，`plllvm.opt_level` 与 `plllvm.auto_checks` 控制编译及插入的检查。LLVM IR 语言没有 DO 处理器。

```sql
CREATE EXTENSION plllvm;
CREATE FUNCTION llvm_add1(x int4) RETURNS int4
LANGUAGE plllvm STRICT AS $$
declare i32 @plllvm_arg_i32(ptr, i32)
define i64 @f(ptr %fcinfo) {
  %x = call i32 @plllvm_arg_i32(ptr %fcinfo, i32 0)
  %r = add i32 %x, 1
  %d = sext i32 %r to i64
  ret i64 %d
}
$$;
SELECT llvm_add1(41);
```

### 运行边界

上游仅测试了文档指定的 PostgreSQL 20devel、LLVM 22.1.8、macOS arm64 基线，必须匹配服务端使用的 LLVM 动态库。安装和语言函数创建均需超级用户权限；原生 IR 没有内存沙箱，可能导致后端崩溃或内存损坏。可选表达式钩子需显式加载。上游未提供扩展原地升级脚本，执行替换流程前应保全依赖函数。
