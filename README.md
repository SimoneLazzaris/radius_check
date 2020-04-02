# radius_check
A nagios plugin to check radius services, written in GO

To build this simple plugin, you'll need the layeh/radius go library. 
To download it and build the plugin:
```
go get -u layeh.com/radius/cmd/radius-dict-gen
go build radius_check

```
# Usage
You have to specify hostname and port, radius secret, username and password in the command line:
```
./radius_check -hostname my.radius.host:1812 -password secretpassword -username mydumbuser -secret supersecret
```
