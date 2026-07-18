## Usage

Sources:

- [Project README](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/README.md)
- [Extension control file](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/pljvm.control)
- [Version 1.0.0 installation SQL](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/pljvm--1.0.0.sql)
- [PostgreSQL call handler](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/src/pljvm.c)
- [Java RPC service](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/plj-java/src/main/java/org/postgresql/plj/PLJVM.java)

`pljvm` 1.0.0 defines a PostgreSQL procedural language that sends function arguments over RPC to code in a separate JVM process, then deserializes the returned value. A function body names a Java class and method; the Java-side service loads the class, creates an instance, and invokes the public method.

### Start the service and define a function

```shell
cd plj-java
mvn exec:java
```

```sql
CREATE EXTENSION pljvm;

CREATE FUNCTION pljvm_add(i int2, j int2)
RETURNS int2
AS 'org.postgresql.plj.test.Int.add'
LANGUAGE pljvm;

SELECT pljvm_add(2::int2, 3::int2);
```

The Java class must be on the service classpath. In the reviewed source both sides use a fixed TCP endpoint at `localhost:8000`; the extension does not start or supervise the JVM service.

### Security and implementation limits

Although the README describes a trusted execution engine, the installation SQL uses `CREATE LANGUAGE pljvm` without the `TRUSTED` keyword. PostgreSQL therefore treats it as an untrusted language, and function creation requires the corresponding elevated privilege. The reviewed RPC path has no configurable endpoint, authentication, or transport encryption; isolate the service and do not expose port 8000.

This snapshot is an early implementation: its C handler says triggers and multi-column results are unsupported, and the Java service uses older Netty/Maven dependencies. Treat JVM process lifecycle, classpath deployment, failure handling, resource limits, and exact PostgreSQL/JDK compatibility as operator responsibilities, and validate them before production use.
