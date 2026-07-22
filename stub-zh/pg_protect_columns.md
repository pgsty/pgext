## 用法

来源：

- [官方软件包页面](https://database.dev/danahartweg/pg_protect_columns)
- [database.dev 官方安装指南](https://database.dev/docs/install-a-package)

`pg_protect_columns` 通过行触发器阻止 UPDATE 改变指定列，适合实现应用级不可变规则，但不能替代权限、行级安全策略或约束。公开软件包页面目前只提供 0.0.1 产物，而目录元数据显示为 0.0.2。

### 核心流程

在具备所需可信语言扩展支持的环境中通过 database.dev/dbdev 安装软件包，然后为每张受保护表创建触发器。

```sql
SELECT dbdev.install('pg_protect_columns');
CREATE EXTENSION IF NOT EXISTS pg_protect_columns;

CREATE TABLE account (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tenant_id bigint NOT NULL,
    display_name text NOT NULL
);

CREATE TRIGGER _protect_columns_before_update
BEFORE UPDATE ON account
FOR EACH ROW
EXECUTE PROCEDURE protect_columns('tenant_id');
```

修改受保护值会抛出 restrict_violation。下划线前缀有助于让该触发器先于按字母排序靠后的触发器运行，但仍应显式审核触发器顺序。

### 受控绕过

- `disable_protection_on_column(text)` 为一个指定列设置事务局部绕过。
- `re_enable_column_protection()` 清除该绕过；受控更新完成后应立即在严格授权的例程内调用它。
- 这个绕过设置不是通用授权系统。不能把这些辅助函数授予必须禁止修改受保护值的角色。

### 注意事项

- 触发器只在已创建且启用它的位置生效。所有者、超级用户、禁用触发器的操作或复制路径可能绕过它。
- 如果软件包安装在当前 search path 之外，应使用 schema 限定函数名。
- database.dev 安装依赖 pg_tle/dbdev 支持。应遵循其可信语言扩展备份指引，并在把逻辑恢复用于灾难恢复前先行测试。
