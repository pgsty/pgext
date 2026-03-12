


## Usage

> [pljava: PL/Java procedural language](https://github.com/tada/pljava)

`pljava` enables writing PostgreSQL functions, triggers, and types in Java using the standard JDBC API.

```sql
CREATE EXTENSION pljava;
```

### Deploy Java Code

Package your Java classes in a JAR file with an SQLJ deployment descriptor, then install it:

```sql
SELECT sqlj.install_jar('file:///path/to/my_functions.jar', 'myjar', true);
SELECT sqlj.set_classpath('public', 'myjar');
```

### Create Functions

Write a Java class with static methods:

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

Declare the SQL function mapping:

```sql
CREATE FUNCTION add(int, int) RETURNS int
  AS 'com.example.MyFunctions.add'
  LANGUAGE java;

CREATE FUNCTION hello(varchar) RETURNS varchar
  AS 'com.example.MyFunctions.hello'
  LANGUAGE java;
```

### Set-Returning Functions

Implement `ResultSetProvider` for set-returning functions:

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

### Trigger Functions

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

### Database Access via JDBC

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

### JAR Management

```sql
SELECT sqlj.install_jar('file:///path/to/jar', 'jarname', true);
SELECT sqlj.replace_jar('file:///path/to/new.jar', 'jarname', true);
SELECT sqlj.remove_jar('jarname', true);
SELECT sqlj.set_classpath('schemaname', 'jar1:jar2');
SELECT sqlj.get_classpath('schemaname');
```
