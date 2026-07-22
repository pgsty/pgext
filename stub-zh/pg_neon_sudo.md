## 用法

来源：

- [扩展控制文件](https://github.com/neondatabase/pg_neon_sudo/blob/2fea160a4d68788a0ad57fc3c87d8af24ef2367a/pg_neon_sudo.control)
- [SQL 函数定义](https://github.com/neondatabase/pg_neon_sudo/blob/2fea160a4d68788a0ad57fc3c87d8af24ef2367a/pg_neon_sudo--0.0.1.sql)
- [C 实现](https://github.com/neondatabase/pg_neon_sudo/blob/2fea160a4d68788a0ad57fc3c87d8af24ef2367a/pg_neon_sudo.c)

`pg_neon_sudo` 0.0.1 是 Neon 专用的权限桥接扩展，用于启动和停止 PostgreSQL Anonymizer 的动态脱敏。它的两个函数会临时以 PostgreSQL 初始超级用户身份执行对应的 `anon` 例程。它是专用基础设施代码，不是通用 `sudo` 工具。

### 核心流程

扩展必须由超级用户安装，并会创建固定的 `sudo` 模式：

```sql
CREATE EXTENSION pg_neon_sudo;

SELECT sudo.anon_start_dynamic_masking();
SELECT sudo.anon_stop_dynamic_masking();
```

在受支持的 PostgreSQL 构建上，两次调用会分别执行 `anon.start_dynamic_masking()` 或 `anon.stop_dynamic_masking()`，然后返回 `operation successful`。

### 函数索引

- `sudo.anon_start_dynamic_masking()` 通过 PostgreSQL Anonymizer 启动动态脱敏。
- `sudo.anon_stop_dynamic_masking()` 停止动态脱敏。

### 依赖与安全边界

实现支持 PostgreSQL 16 及以上版本；较早版本会报“不支持该特性”。目标 `anon` 函数必须已经存在，但控制文件没有声明对 PostgreSQL Anonymizer 的依赖。该扩展无需预加载。

这些函数会主动跨越超级用户权限边界。应只向确实负责脱敏管理的角色授予 `EXECUTE`，并审计每一项授权。SQL 定义把函数标为 `IMMUTABLE`，但它们实际具有操作副作用，因此不要在索引、生成列或可能被规划器折叠的表达式中调用。该扩展面向 Neon 的受控环境；在其他环境使用前必须重新评估可移植性与安全假设。
