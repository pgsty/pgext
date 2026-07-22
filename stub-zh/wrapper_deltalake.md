## 用法

来源：

- [Official README](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/README.md)
- [Extension control file](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/wrapper_deltalake.control)
- [Rust implementation](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/src/lib.rs)

wrapper_deltalake 是一个未完成的概念验证，原计划连接 Supabase Wrappers 与 delta-rs。所核验版本没有实现外部数据包装器、Delta Lake 读取器、服务器选项或表映射；它只暴露一个问候函数。

### 核心流程

完整的用户可见行为只有这条演示调用：

```sql
CREATE EXTENSION wrapper_deltalake;

SELECT hello_wrapper_deltalake();
```

它会返回固定问候语。项目没有定义外部服务器或查询 Delta 表的受支持方法。

### 重要对象

- `hello_wrapper_deltalake` 返回固定文本 `Hello, wrapper_deltalake`。
- `wrapper_deltalake` 是扩展名，但没有注册 FDW handler 或 validator。

### 运维说明

目录和 control 文件使用 0.0.0 版本，README 描述的是未来意图而非已实现行为。不要从 Supabase Wrappers 或 delta-rs 文档中臆造配置：那些项目的能力并不存在于此扩展。实际数据访问应选择仍受维护且已经实现的 Delta Lake 集成。
