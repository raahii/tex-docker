# tex-docker

This is a CLI tool which enables compile TeX document without tex-live installation by using docker.

Currently supporting executing `latexmk` on your tex document directory. In addition, automatic compilation by your file changes is possible.

### Installation

```shell
go get -u github.com/raahii/tex-docker
```

### Example

- compile sources at once

  ```shell
  tex-docker -p <tex source directory>
  ```

- automatically compile every time file changes.

  ```shell
  tex-docker -p <tex source directory> -w <regex that matches files to watch e.g. main.tex>
  ```

### Usage

  ```
  Usage:
    tex-docker [OPTIONS]

  Application Options:
    -p, --path=           Latex source path to compile (execute latexmk)
    -n, --container-name= Docker container name
    -w, --watch=          Process any events whose filename matches the specified POSIX extended regular expression
    -r, --recursive       Watch all subdirectories of the --path directory

  Help Options:
    -h, --help            Show this help message
  ```


### TODO

- [x] auto compilation triggered by file changes
- [ ] support `.latexmk` customization

