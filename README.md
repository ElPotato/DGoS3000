# DGoS3000

Dead simple command line tool to send number of HTTP requests concurently.

### Flags/Params
```
Mode: debug, silent, error
Workers: 1-∞
List: <path to the list with urls>
```

### Example
```
go run *.go -workers 128 -list list.txt -mode debug
```