## File Server
A small and simple file hosting server

## Usage
- Clone it
```sh
git clone https://github.com/kiwimoe/fileserver
```
- Compile
```sh
# Visit https://go.dev/doc/install and download according your platform, if you haven't
cd fileserver && go build

## If you want to change port or server URL, check out main.go file and change URL variable
## And run
./fileserver
```
- Now server should be available on localhost:1337 (if URL not modified else shall be active on your configuration)
