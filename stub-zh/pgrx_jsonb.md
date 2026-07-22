## 用法

来源：

- [Official README](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/README.md)
- [Extension control file](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/pgrx_jsonb.control)
- [Rust implementation](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/src/lib.rs)

pgrx_jsonb 是一个通过 Rust/pgrx 遍历 PostgreSQL 内部 JSONB 表示的概念验证。上游明确要求不要将其用于任何重要场景。它的三个演示函数只适合研究扩展内部实现。

### 核心流程

只能在一次性开发数据库中安装，并比较几条演示路径：

```sql
CREATE EXTENSION pgrx_jsonb;

SELECT jsonb_to_text('{"answer": 42}'::jsonb);
SELECT jsonb_test('{"answer": 42}'::jsonb);
SELECT jsonb_test2('{"answer": 42}'::jsonb);
```

第一个函数遍历底层表示并将其序列化，另外两个函数通过不同代码路径报告顶层值是否为对象。

### 重要对象

- `jsonb_to_text` 通过原型迭代器把 JSONB 值转换为文本。
- `jsonb_test` 通过自定义底层迭代器检查值是否为对象。
- `jsonb_test2` 使用 pgrx 的高层 JSONB 包装器执行比较。

### 安全与兼容性说明

迭代器依赖 PostgreSQL 内部 JSONB API，并包含断言、unwrap、panic 以及对某些 token 类型尚未实现的分支。不受支持的数据可能报错或终止后端进程。目录版本是 0.0.0，项目没有稳定兼容性承诺。实际扩展应使用 PostgreSQL 内置 JSONB 函数或仍受维护的 pgrx API。
