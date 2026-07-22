## Usage

Sources:

- [Official README](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/brainfuck/README.md)
- [Official control file](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/brainfuck/brainfuck.control)
- [Official version 4.2.0 SQL](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/brainfuck/brainfuck--4.2.0.sql)
- [Official package page](https://database.dev/mansueli/brainfuck)

`brainfuck` runs Brainfuck programs inside PostgreSQL through PLV8. It is useful for demonstrations and experiments, not for application logic: arbitrary programs execute JavaScript in the database backend and can consume substantial CPU.

### Core Workflow

Install `plv8` first, then create `brainfuck`:

```sql
CREATE EXTENSION plv8;
CREATE EXTENSION brainfuck;

SELECT brainfuck('++++++++[>++++++++<-]>+.', '');
```

The second argument supplies input consumed by the `,` instruction. Pass an empty string when the program needs no input.

### API

`brainfuck(code text, input text) RETURNS text` is the extension's only SQL function. It is declared `IMMUTABLE STRICT`, uses a 256-cell wrapping tape, implements the eight Brainfuck instructions, and ignores other characters in `code`.

### Caveats

The extension requires `plv8`; the package registry workflow additionally uses `pg_tle`/`dbdev`, while a conventional extension installation does not. The actual SQL signature has two arguments even though an upstream README example shows one.

There is no execution limit or sandbox specific to the interpreted program. Malformed brackets fail during generated JavaScript compilation, and long-running or non-terminating programs occupy a PostgreSQL backend. Keep it out of untrusted SQL paths and use it only in controlled experiments.
