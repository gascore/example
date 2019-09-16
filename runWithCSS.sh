cd ~/go/src/github.com/gascore/burn.css
yarn run build
cd ~/go/src/github.com/gascore/example
cp ~/go/src/github.com/gascore/burn.css/dist/burn.min.css ~/go/src/github.com/gascore/example/app/static/burn.min.css
go run main.go -w