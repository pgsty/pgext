


## 用法

> [pagevis: PostgreSQL 数据库页面的 ASCII 可视化](https://github.com/hollobon/pagevis)

pagevis 提供 PostgreSQL 数据库页面内容的 ASCII 图形化表示。它依赖 `pageinspect` 扩展。

### 函数

```sql
-- show_page(关系名, 块号, 行宽)
SELECT show_page('my_table', 0, 64);
```

### 输出字符

| 字符 | 含义 |
|-----------|---------|
| `P` | 页头 |
| `U` | 未使用的行指针 |
| `N` | 正常行指针 |
| `R` | 重定向行指针 |
| `D` | 已死行指针 |
| (空格) | 空闲空间（"空洞"） |
| `[n]` | 元组编号（叠加在元组头上） |
| `H` | 元组头 |
| `#` | 元组数据 |
| `-` | 对当前事务不可见的元组（已死） |
| `.` | 元组间填充 |

### 示例

```sql
SELECT show_page('pvtest', 0, 64);
                             show_page
------------------------------------------------------------------
 PPPPPPPPPPPPPPPPPPPPPPPP001N002N003N004N005N006N007N008N009N010N
 011N012N013N014N015N016N017N018N019N020N...

 [226]HHHHHHHHHHHHHHHHHHH####....[225]HHHHHHHHHHHHHHHHHHH####....
 [224]HHHHHHHHHHHHHHHHHHH####....[223]HHHHHHHHHHHHHHHHHHH####....
```

页头（`P`）和行指针（`N`）出现在开头，空闲空间在中间，元组从页面末尾向前增长。当行指针和元组之间没有剩余空间时，页面就满了。

需要超级用户权限（使用 `pageinspect` 的 `get_raw_page`）。
