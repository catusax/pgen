# pgen

a project generator


## Why

when you are writing microservices or serveless function, it's really annoying to copy-paste project structure and code, especially when you want to change one file, then you have to do that agent.

with pgen, you cant wirte your own project skaffold with your favourite template engine, then generate / regenerate a project with one command.

## features

- [x] creating && regenerating project
- [x] customized project structure
- [x] customized template engine
- [x] customized function for template engine with wasm
- [ ] code formatting
- [ ] build-in project templates

## Usage

1. create a dir named `.template`
2. write your templates to `.template`
3. classify your files by onceFiles softFiles hardFiles and put the paths to `.pgen_config.yaml`
4. create your new project with `pgen new <project-name>`
5. regenerate project code with `PGEN_SOMEKEY=NEW-VALUE pgen generate`, your project is regenerated with NEW-VALUE

see example for more information [example project](example)

use `--help` to view more command line usage

## Config

config file is placed at `.template/.pgen_config.yaml`

```yaml

# OnceFiles only generate once when creating new project
once_files:
  - dir/filename.go
  - go.mod
# SoftFiles will be generate every time when using generate command,
# so you can use easily change file content by using ENV vairables.
# you dont need to commit these files to git
soft_files:
  - main.go
# HardFiles will be generate every time when using generate command with --hard flag
hard_files:
  - handler/handler.go
# DefaultENVs is default environment variables when creating new project.
# on existing project, Makefile variables will replace these variables.
# by default the program will inject a env NAME this value of your project name.
default_envs:
  PGEN_SOMEKEY: VALUE

# you can define your own functions with webassembly
wasm_funcs:
  # path of the webassembly library, relative to this config file's directory or absolute path
  - path: wasm/target/wasm32-unknown-unknown/debug/wasm.wasm
  # function to be exported, must have one cstring pointer input and one cstring pointer output
  # see wasm example for more information
    funcs:
      - greet

```
