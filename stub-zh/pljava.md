

## 用法

> [pljava: PL/Java 过程语言](https://github.com/tada/pljava)

`pljava` 允许使用标准 JDBC API 在 PostgreSQL 中编写 Java 函数、触发器和类型。

```sql
CREATE EXTENSION pljava;
```

### 部署 Java 代码

将 Java 类打包到包含 SQLJ 部署描述符的 JAR 文件中，然后安装：

```sql
SELECT sqlj.install_jar('file:///path/to/my_functions.jar', 'myjar', true);
SELECT sqlj.set_classpath('public', 'myjar');
```

### 创建函数

编写包含静态方法的 Java 类：

```java
package com.example;

import org.postgresql.pljava.annotation.Function;

public class MyFunctions {
    @Function
    public static int add(int a, int b) {
        return a + b;
    }

    @Function
    public static String hello(String name) {
        return "Hello, " + name + "!";
    }
}
```

声明 SQL 函数映射：

```sql
CREATE FUNCTION add(int, int) RETURNS int
  AS 'com.example.MyFunctions.add'
  LANGUAGE java;

CREATE FUNCTION hello(varchar) RETURNS varchar
  AS 'com.example.MyFunctions.hello'
  LANGUAGE java;
```

### 集合返回函数

实现 `ResultSetProvider` 接口以创建集合返回函数：

```java
import org.postgresql.pljava.ResultSetProvider;
import java.sql.ResultSet;
import java.sql.SQLException;

public class MySetFunction implements ResultSetProvider {
    public boolean assignRowValues(ResultSet receiver, int currentRow)
            throws SQLException {
        if (currentRow < 10) {
            receiver.updateInt(1, currentRow);
            receiver.updateString(2, "row " + currentRow);
            return true;
        }
        return false;
    }

    public void close() {}

    public static ResultSetProvider generate()
            throws SQLException {
        return new MySetFunction();
    }
}
```

### 触发器函数

```java
import org.postgresql.pljava.TriggerData;
import org.postgresql.pljava.annotation.Trigger;

public class MyTrigger {
    @Trigger(called = Trigger.Called.BEFORE, table = "my_table",
             events = {Trigger.Event.INSERT, Trigger.Event.UPDATE})
    public static void auditTrigger(TriggerData td) throws SQLException {
        ResultSet newRow = td.getNew();
        newRow.updateTimestamp("modified_at",
            new java.sql.Timestamp(System.currentTimeMillis()));
    }
}
```

### 通过 JDBC 访问数据库

```java
import java.sql.*;

public static int countUsers() throws SQLException {
    Connection conn = DriverManager.getConnection("jdbc:default:connection");
    PreparedStatement stmt = conn.prepareStatement("SELECT count(*) FROM users");
    ResultSet rs = stmt.executeQuery();
    rs.next();
    return rs.getInt(1);
}
```

### JAR 管理

```sql
SELECT sqlj.install_jar('file:///path/to/jar', 'jarname', true);
SELECT sqlj.replace_jar('file:///path/to/new.jar', 'jarname', true);
SELECT sqlj.remove_jar('jarname', true);
SELECT sqlj.set_classpath('schemaname', 'jar1:jar2');
SELECT sqlj.get_classpath('schemaname');
```
