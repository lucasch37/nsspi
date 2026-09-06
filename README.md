# nsspi

Pascal interpreter written in Go. Extended upon ideas from [Ruslan Spivak's series](https://ruslanspivak.com/lsbasi-part1/).

## Usage

```bash
# Run an example from examples/:
go run . -scope=false -stack=false examples/collatz.pas
```

## Todo

| Feature                | Done |
| ---------------------- | :-------: |
| AST                    |     [✓]     |
| Arithmetic Expressions |     [✓]     |
| SymTable               |     [✓]     |
| CallStack              |     [✓]     |
| Debugger               |     [✓]     |
| Nested Declarations    |     [✓]     |
| Recursion              |     [✓]     |
| `writeln` / `write`    |     [✓]     |
| Boolean Expressions    |     [✓]     |
| Conditionals           |     [✓]     |
| Functions              |     [✓]     |
| For / While Loops      |     [✓]     |
