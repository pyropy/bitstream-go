# Example usage

### Encrypting the file
```
go run example/main.go -action=encrypt -in=bitcoin.pdf -out=bitcoin-enc.pdf -hash=00000000000000000000000000000000000000 -preimage=0000000000000000000000000000000000000000 
```

#### Decrypting the file
```
go run example/main.go -action=decrypt -in=bitcoin-enc.pdf -out=bitcoin-dec.pdf -hash=00000000000000000000000000000000000000 -preimage=0000000000000000000000000000000000000000 
```