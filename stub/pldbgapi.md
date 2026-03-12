


## Usage

> [pldbgapi: server-side support for debugging PL/pgSQL functions](https://github.com/EnterpriseDB/pldebugger)

`pldbgapi` provides a server-side API for interactive debugging of PL/pgSQL functions. It is typically used through a GUI client such as **pgAdmin**.

```sql
CREATE EXTENSION pldbgapi;
```

### Debugging with pgAdmin

The primary way to use the debugger is through pgAdmin's graphical interface:

- **Direct Debugging**: Right-click a function and select "Debug" to execute and step through it immediately
- **Global Breakpoints**: Select "Set Global Breakpoint" on a function, then wait for another session (e.g., a web application) to call that function -- the debugger will intercept the call and allow in-context debugging

### Debugging Capabilities

When connected through a debug client, you can:

- **Set breakpoints** on specific lines in PL/pgSQL functions
- **Step through** code line by line (step into, step over, step out)
- **Inspect variables** and their current values at each step
- **View the call stack** for nested function calls
- **Continue execution** to the next breakpoint

### Architecture

The debugging system has three components:

1. **Client GUI** (pgAdmin) -- displays source code, variables, and stack
2. **Target Backend** -- the session executing the PL/pgSQL code being debugged
3. **Debugging Proxy** -- coordinates between the client and target via a dedicated connection

### Supported Languages

The debugger works with PL/pgSQL functions and procedures. It requires the `pldbgapi` extension to be created in each database where debugging is needed.
