## 用法

来源：

- [官方上游 README](https://github.com/mohammadzainabbas/student-in-postgres/blob/14eb780cb9625aa1f965057df854a5c0b4bdfb1a/README.md)
- [官方扩展控制文件 (student.control)](https://github.com/mohammadzainabbas/student-in-postgres/blob/14eb780cb9625aa1f965057df854a5c0b4bdfb1a/student.control)
- [官方扩展 SQL (student--1.0.sql)](https://github.com/mohammadzainabbas/student-in-postgres/blob/14eb780cb9625aa1f965057df854a5c0b4bdfb1a/student--1.0.sql)

`student` 演示了如何在 PostgreSQL 中开发基础类型。应用数据需要此类型、域或相关运算符时可使用它。请以上方链接的固定上游修订为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION student;
```

在目标数据库中安装扩展；如果上游提供了最小示例，请运行该示例，并在集成到应用 SQL 前验证安装版本和返回值。

### 重要对象

- `student(text)` 是扩展函数，返回 `student`。
- `student_in(cstring)` 是扩展函数，返回 `student`。
- `student_out(student)` 是扩展函数，返回 `cstring`。
- `text(student)` 是扩展函数，返回 `text`。
- `student` 是扩展定义的类型。

### 要求与注意事项

- 审阅的控制文件声明默认版本为 `1.0`。
- 控制文件将该扩展标记为可重定位。
- 生产使用前，请根据固定版本源码确认权限、支持的 PostgreSQL 版本、升级行为和失败情形。
