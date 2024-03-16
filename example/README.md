# Example usage

### Encrypting the file
```
go run example/main.go -action=encrypt -in=bitcoin.pdf -out=bitcoin-enc.pdf -hash=84645f0ccdd6f4695d655b3b6e8dfe6f4f35ad55fac62981c98216d74f87bae5 -preimage=28a794564e8c20a64b6376c4ebdd84ae3525f45342b957f9334f6fdb86c7615d 
```

#### Decrypting the file
```
go run example/main.go -action=decrypt -in=bitcoin-enc.pdf -out=bitcoin-dec.pdf -hash=84645f0ccdd6f4695d655b3b6e8dfe6f4f35ad55fac62981c98216d74f87bae5 -preimage=28a794564e8c20a64b6376c4ebdd84ae3525f45342b957f9334f6fdb86c7615d 
```