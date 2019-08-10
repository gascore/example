cd ~/go/src/github.com/gascore/burn.css
yarn run build
cd ~/go/src/github.com/gascore/example
cp ~/go/src/github.com/gascore/burn.css/dist/chota.min.css ~/go/src/github.com/gascore/example/app/static/chota.min.css
go run main.go -watch