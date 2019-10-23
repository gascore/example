cd ~/go/src/github.com/gascore/burn.css
yarn run build
cd ~/go/src/github.com/gascore/example
cp ~/go/src/github.com/gascore/burn.css/dist/burn.css ~/go/src/github.com/gascore/example/node_modules/burn.css/dist/burn.css
go run main.go -w