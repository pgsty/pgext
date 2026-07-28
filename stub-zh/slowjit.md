## 用法

来源：

- [官方上游 README](https://github.com/higuoxing/pg_slowjit/blob/bb7dae7b7283e3822e546dafef50ed3e053e3dac/README.md)
- [官方扩展控制文件 (slowjit.control)](https://github.com/higuoxing/pg_slowjit/blob/bb7dae7b7283e3822e546dafef50ed3e053e3dac/slowjit.control)
- [官方实现源码](https://github.com/higuoxing/pg_slowjit/blob/bb7dae7b7283e3822e546dafef50ed3e053e3dac/slowjit.c)

`slowjit` 是一个教育性的 PostgreSQL JIT 提供程序实现。它在运行时生成 C 代码，调用 C 编译器，加载生成的共享库，并且目前仅处理少量表达式运算符。

### 核心工作流

构建并安装提供程序库，在 PostgreSQL 的 JIT 提供程序配置中选择它，并在一个隔离的服务器中强制一个简单的表达式通过 JIT 编译：

```ini
jit_provider = 'slowjit'
jit_above_cost = 0
```

```sql
EXPLAIN (SETTINGS ON)
SELECT 1;
```

README 仅使用此查询来演示一个函数被 JIT 编译。该模块并未安装面向用户的 SQL API，也未建立独立的 `CREATE EXTENSION` 工作流。

### 重要设置

- `slowjit.cc_path` 选择 C 编译器可执行文件；实现默认为 `cc`。

### 要求与注意事项

- 审查过的控制文件、注册表或目录证据标识版本 `1.0.0`。
- 控制文件将该扩展标记为可重定位。
- 上游明确称该提供程序非常低效，并仅文档化了对少数几个运算符的支持。
- 运行时编译执行服务器端编译器并创建可加载的代码。限制文件系统权限，并且永远不要将此原型视为经过加固的提供程序。
- JIT 提供程序接口是版本敏感的；在与 PostgreSQL 服务器源代码完全相同的版本中构建和测试它。
