# Dinkur

Task and time tracking utility.

## Install

```console
$ go install --tags fts5 github.com/dinkur/dinkur@latest
```

> The `--tags fts5` flags adds [Sqlite FTS5](https://www.sqlite.org/fts5.html)
> support, which is used for better and more performant full-text search.

For you CLI-power users, we recommend aliasing it to `ur`.

```sh
alias ur=dinkur
```

### CLI Autocompletion

Automatic generation of completions are powered by [Cobra](https://github.com/spf13/cobra),
which supports Bash, Zsh, Fish, and PowerShell.

Completions are provided by the `completion` subcommand. To get a more detailed
guide on how to install them, run the following (where `bash` can be exchanged
with `zsh`, `fish`, and `powershell`):

```sh
dinkur completion bash --help
```

## Development

Requires Go v1.17 (or higher)

```sh
go run --tags fts5 .
```
