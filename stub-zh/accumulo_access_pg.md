## 用法

来源：

- [官方 README](https://github.com/larsw/accumulo-access-pg/blob/v0.3.0/README.md)
- [官方扩展控制文件](https://github.com/larsw/accumulo-access-pg/blob/v0.3.0/accumulo_access_pg.control)
- [官方 Rust 实现](https://github.com/larsw/accumulo-access-pg/blob/v0.3.0/src/lib.rs)
- [官方包元数据](https://github.com/larsw/accumulo-access-pg/blob/v0.3.0/Cargo.toml)
- [Apache-2.0 许可证](https://github.com/larsw/accumulo-access-pg/blob/v0.3.0/LICENSE_APACHE)
- [MIT 许可证](https://github.com/larsw/accumulo-access-pg/blob/v0.3.0/LICENSE_MIT)

`accumulo_access_pg` 0.3.0 在 PostgreSQL 内解析并计算 Apache Accumulo 风格的授权表达式，主要用于在行级安全策略中应用 `label1 & (label2 | label3)` 之类的表达式。扩展只计算调用方传入的标签；它不会自行认证用户，也不会取得可信授权声明。

### 核心流程

为每个受保护行保存授权表达式，并把可信的逗号分隔标签集传给 `sec_authz_check`。表达式或标签列表为空或 null 时，策略会拒绝该行。

```sql
CREATE EXTENSION accumulo_access_pg;

CREATE TABLE protected_notes (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    body text NOT NULL,
    authz_expr text NOT NULL
);

ALTER TABLE protected_notes ENABLE ROW LEVEL SECURITY;

CREATE POLICY protected_notes_read ON protected_notes
USING (
    sec_authz_check(
        authz_expr,
        current_setting('app.authorizations', true)
    )
);

SET app.authorizations = 'finance,eu';
SELECT * FROM protected_notes;
```

不要允许不可信数据库用户自由设置策略所读取的值。上游示例用会话 GUC 进行说明，但能自行选择 `app.authorizations` 的用户也能给自己授予标签。应通过受控的登录/会话初始化路径或经过安全审查的辅助函数导入声明。

### 函数

- `sec_authz_check(expression, tokens)` 使用逗号分隔标签计算一个表达式；表达式格式错误会抛出异常。
- `sec_expr_as_json(expression)` 和 `sec_expr_as_json_string(expression)` 暴露解析后的表达式供检查。
- `sec_authz_cache_stats()` 返回解析/计算缓存的 `hits`、`misses` 和 `size` 字段。
- `sec_authz_clear_cache()` 清空该缓存，并以布尔值返回是否成功。

表达式语法与转义遵循上游 `accumulo-access` Rust 库。应测试带引号标签、Unicode、空白、错误语法及应用需要的精确语义。

### 安装与安全边界

0.3.0 发布为 PostgreSQL 15–18 构建并测试，同时提供 PostgreSQL 19 beta 包；PostgreSQL 13–14 仅由 pgrx 支持但未构建或测试。控制文件仍要求超级用户安装且不可重定位。此版本拒绝非法引号词元（包括空引号词元），括号嵌套最多 256 层，带引号词元可包含控制字符。

行级安全取决于所有相关角色、策略、绕过权限和声明来源，而不只是 `sec_authz_check`。应验证表所有者、超级用户、具有 `BYPASSRLS` 的角色、连接池、会话状态重置/复用及畸形声明的行为。缓存统计只是运维辅助信息，不是授权审计日志。
