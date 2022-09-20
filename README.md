# TODO

Installation:

```bash
go build .

mv todo /usr/bin/todo
```

## Usage:

- Adding items

```
todo "finish writing tests for todo"
```

- Viewing todo items

```
todo
```

Example output:

```
1 - finish writing tests for todo - 2022-09-19T23:31:15.51983599-07:00
```

- Removing items:

```
todo pop 1
```

- Clearing all items:

```
todo clear
```
