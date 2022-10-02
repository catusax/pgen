# pgen

a project generator


## Why

when you are writing microservices or serveless function, it's really annoying to copy-paste project structure and code, especially when you want to change one file, then you have to do that agent.

with pgen, you cant wirte your own project skaffold with your favourite template engine, then generate / regenerate a project with one command.

## Usage

1. create a dir named `.template`
2. write your templates to `.template`
3. classify your files by onceFiles softFiles hardFiles and put the paths to `.pgen_config.yaml`
4. create your new project with `pgen new <project-name>`
5. regenerate project code with `PGEN_SOMEKEY=NEW-VALUE pgen generate`, your project is regenerated with NEW-VALUE

an example is comming soon...

use `--help` to view more command line usage

## Config

config file is placed at `.template/.pgen_config.yaml`

```yaml

# OnceFiles only generate once when creating new project
onceFiles:
  - dir/filename.go
  - go.mod
# SoftFiles will be generate every time when using generate command,
# so you can use easily change file content by using ENV vairables.
# you dont need to commit these files to git
softFiles:
  - main.go
# HardFiles will be generate every time when using generate command with --hard flag
hardFiles:
  - handler/handler.go
# DefaultENVs is default environment variables when creating new project.
# on existing project, Makefile variables will replace these variables.
# by default the program will inject a env NAME this value of your project name.
defaultEnvs:
  PGEN_SOMEKEY: VALUE

```
