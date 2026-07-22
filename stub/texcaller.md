## Usage

Sources:

- [Official PostgreSQL binding documentation](https://vog.github.io/texcaller/group__postgresql.html)
- [Extension SQL](https://github.com/vog/texcaller/blob/fc5fbc9862a6df3e907e2ee409e5d995aa175ef6/postgresql/texcaller.sql)
- [Conversion implementation](https://github.com/vog/texcaller/blob/fc5fbc9862a6df3e907e2ee409e5d995aa175ef6/c/texcaller.c)

`texcaller` compiles TeX-family source inside a PostgreSQL backend and returns the generated DVI or PDF document as `bytea`. It is suitable only for tightly controlled document-generation workloads: a call creates temporary files and starts an external TeX executable with the database server's operating-system identity.

### Core Workflow

Install the extension, escape any plain user-supplied fragment, assemble a complete document, and compile it with a bounded number of runs:

```sql
CREATE EXTENSION texcaller;

SELECT texcaller_convert(
  '\documentclass{article}' ||
  '\begin{document}' ||
  texcaller_escape_latex('A value with $ and % characters') ||
  '\end{document}',
  'LaTeX',
  'PDF',
  5
);
```

`texcaller_convert` is strict and returns binary output. `max_runs` must be at least 2; compilation repeats until the auxiliary file stabilizes or the limit is reached.

### Formats and Functions

- `texcaller_convert(source, source_format, result_format, max_runs)` compiles the document. Source formats are `TeX`, `LaTeX`, `XeTeX`, `XeLaTeX`, `LuaTeX`, and `LuaLaTeX` for PDF output; DVI output is available only from `TeX` or `LaTeX`.
- `texcaller_escape_latex` quotes special characters for inserting plain text into LaTeX source. It does not validate an entire template or create a security sandbox.

The required executable must be on the PostgreSQL server process's path: depending on the requested pair, the implementation invokes `tex`, `latex`, `pdftex`, `pdflatex`, `xetex`, `xelatex`, `luatex`, or `lualatex`.

### Security and Operations

The child process is started with batch, halt-on-error, file-line diagnostics, and `-no-shell-escape` options. Disabling shell escape is important but does not make arbitrary TeX safe: TeX programs can be computationally expensive and may access files allowed by the installed engine and operating-system policy. Revoke public execution and grant it only to a dedicated trusted role:

```sql
REVOKE EXECUTE ON FUNCTION
  texcaller_convert(text, text, text, integer)
FROM PUBLIC;

GRANT EXECUTE ON FUNCTION
  texcaller_convert(text, text, text, integer)
TO document_renderer;
```

Each call forks a child process and uses a private directory under `TMPDIR` or `/tmp`. Apply statement timeouts, input-size limits, concurrency controls, filesystem confinement, and monitoring for CPU, memory, disk, and process consumption. Do not concatenate untrusted TeX commands around escaped fragments.

The extension control file reports version 0, is relocatable, and has no preload setting. The upstream code is old; validate compiler binaries, fonts, TeX packages, PostgreSQL compatibility, cleanup behavior, and byte-for-byte output on the target host. Document output can vary with the TeX installation even when the SQL input is unchanged.
