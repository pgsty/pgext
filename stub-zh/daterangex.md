## 用法

来源：

- [官方文档](https://github.com/waaeer/pg_daterangex/blob/397429615dd2d89fe96f99c47ff5092a0f35af86/README.md)
- [扩展控制文件](https://github.com/waaeer/pg_daterangex/blob/397429615dd2d89fe96f99c47ff5092a0f35af86/daterangex.control)
- [官方仓库](https://github.com/waaeer/pg_daterangex)

`daterangex` 上界按闭区间显示的日期范围类型，并提供转换与操作符。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `daterangex`：

```sql
CREATE EXTENSION daterangex;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE EXTENSION daterangex;
SELECT daterange('[2024-01-01,2024-06-01]');
        daterange
-------------------------
 [2024-01-01,2024-06-02)

SELECT daterangex('[2024-01-01,2024-06-01]');
       daterangex
-------------------------
 [2024-01-01,2024-06-01]
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `date_minus` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `daterangex_canonical` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
