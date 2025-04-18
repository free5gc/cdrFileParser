# tttns

`tttns` is a command-line interface (CLI) tool designed to inspect 3GPP TS 32.297 CDR (Charging Data Record) files. The name "tttns" is derived from the first letters of "32297".

> Thanks to [Hao Li](https://github.com/haoli000) for the inspiration and initial implementation of this tool. The project has been developed to provide a user-friendly interface for working with CDR files, making it easier to extract and analyze data.

## Installation

```
go install main.go
```

## Usage

```bash
main [file|-] [flags]
main [command]
```

## Available Commands

- `cdr`: Print all CDR header info or use its sub commands
- `completion`: Generate the autocompletion script for the specified shell
- `file`: Print CDR file header info or use its sub commands
- `help`: Help about any command
- `version`: Print the version number of tttns

## Flags

- `-h, --help`: Display help for tttns
- `-j, --json`: Output in JSON format (except for cdr dump)

## Detailed Command Information

### cdr Command

The `cdr` command is used to print and manipulate CDR (Charging Data Record) information.

Usage:

```bash
main cdr [file|-] [flags]
main cdr [command]
```

Flags:

- `-h, --help`: Display help for the cdr command
- `-j, --json`: Output in JSON format

#### Available Subcommands

1. **count**: Get the number of CDRs in a file

   ```bash
   main cdr count [file|-]
   ```

2. **dump**: Dump the raw content of CDR to stdout

   ```bash
   main cdr dump [file|-] [index|1]
   ```

3. **header**: Print CDR header info

   ```bash
   main cdr header [file|-] [index|1]
   ```

For more information about a specific subcommand, use:

```bash
main cdr [subcommand] --help
```

## Examples

1. Get the number of CDRs in a file:

   ```bash
   main cdr count example/imsi-123456789012345.cdr
   ```

2. Print CDR header info of the 1st CDR:

   ```bash
   main cdr header example/imsi-123456789012345.cdr 1
   cat example/imsi-123456789012345.cdr | main cdr header 1
   ```

3. Dump the raw content of the 2nd CDR to stdout:

   ```bash
   main cdr dump example/imsi-123456789012345.cdr 1
   cat example/imsi-123456789012345.cdr | main cdr dump 1
   ```

4. Print CDR header info in JSON format:
  
   ```bash
   main cdr header example/imsi-123456789012345.cdr 1 --json
   ```

## License

Apache-2.0.

## Notes

The project has been scaffolded with the help of [kleiner](https://github.com/can3p/kleiner).
