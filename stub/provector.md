## Usage

Sources:

- [README.md](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/README.md)
- [provector--1.0.sql](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/provector--1.0.sql)
- [test/test.sql](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/test/test.sql)
- [provector.cpp](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/provector.cpp)
- [provector.control](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/provector.control)

`provector` 1.0 is an early C++ extension example. Its reviewed SQL surface contains one text-conversion function; the project name and pgvector dependency do not establish a vector-search implementation.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION provector;
SELECT provector_demo('a title');
```

### Operational Boundaries

`provector_demo` accepts text, converts it to title case, and returns text. It is declared strict and volatile. The control file requires `vector` even though this example function accepts no vector argument.

Create the extension as a superuser after its library and dependencies are installed. There is no preload or restart requirement in the reviewed source. The build links locale/gettext facilities; behavior for non-ASCII input should be checked in the intended locale. No supported PostgreSQL-major matrix is published. Treat this as an educational prototype rather than a production search component.
