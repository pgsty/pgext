## 用法

来源：

- [项目 README](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/README.md)
- [扩展 control 文件](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/src/plhaskell.control)
- [5.0 版 SQL 语言定义](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/src/plhaskell--5.0.sql)
- [用法与查询 API](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/docs/Usage.md)
- [信任模型](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/docs/Trust.md)
- [已知错误](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/docs/Known_Bugs.md)

`plhaskell` 5.0 把 Haskell 嵌入为 PostgreSQL 过程语言。安装会创建用于带类型 `PGm` action 的 trusted `plhaskell`，以及用于 `IO` 的 untrusted `plhaskellu`；只有超级用户能用 untrusted 变体创建函数。数据库编码必须是 UTF-8。

### 定义 trusted 函数

```sql
CREATE EXTENSION plhaskell;

CREATE FUNCTION add_haskell(int, int)
RETURNS int
IMMUTABLE
AS $$
  import Data.Int (Int32)
  import PGutils (PGm)

  add_haskell :: Maybe Int32 -> Maybe Int32 -> PGm (Maybe Int32)
  add_haskell Nothing Nothing = return Nothing
  add_haskell (Just a) (Just b) = return $ Just $ a + b
  add_haskell a Nothing = return a
  add_haskell Nothing b = return b
$$
LANGUAGE plhaskell;

SELECT add_haskell(20, 22);
```

扩展还支持复合类型、数组、时间戳、advisory lock、参数化 SPI 查询、集合和匿名块。函数 volatility 决定内部查询是否只读；每条 SQL 查询和导入模块仍需审查。

### 运行时与信任边界

trusted 语言声明依赖扩展的 Haskell 类型和模块限制。同时，它会在 PostgreSQL 后端内调用 GHC API 和动态加载的 Haskell runtime，因此允许普通角色创建函数会显著扩大数据库执行表面。应独立审查 trusted 模块集合、package database、文件系统与 SELinux 策略、语句超时、角色权限和拒绝服务行为。`plhaskellu` 只应保留给严格控制的超级用户代码。

Haskell runtime 默认内存上限为 128 MB；`plhaskell.max_memory = 0` 会禁用上限，不适合多租户使用。上游记录了某些 `lc_messages` 值导致失败、多次调用集合返回函数时严重变慢，以及平台 runpath 问题。构建面向 Linux，且只在 x86-64 上测试。应固定匹配的 PostgreSQL、GHC、hint 与 runtime 库，并在生产部署前测试后端崩溃、错误、升级、复制和转储恢复。
