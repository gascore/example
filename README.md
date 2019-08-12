# Example - full example of gas ecosystem usage

### Dependencies:

0. [gasocre/gas](https://github.com/gascore/gas)
1. [gasocre/std](https://github.com/gascore/std)
2. [sass](https://sass-lang.com)
3. [golang-commonmark/markdown](https://gitlab.com/golang-commonmark/markdown) - markdown parser

### Gettings started:

```bash
// install nodejs, I recommend github.com/nvm-sh/nvm
npi i -g sass

GO111MODULE=off go get github.com/gascore/example
cd $GOPATH/src/github.com/gascore/example
export GO111MODULE=on

cd app && go get && cd ..

go run main.go -watch 
```
