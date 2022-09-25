# How to

Generate a certificate for self-signing:

    openssl req -x509 -nodes -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 365

Then, run the server:

    go run cmd/tlsinfo/clienttlsinfo.go

Then make https requests to it, it will print the TLS fingerprints :

    curl -k https://localhost:1234/ffk

Try with different http clients to see the differences and consistency.

Note: the implemented algorithm is freely adapted from JA3.

Credits and inspiration:

- https://github.com/dreadl0ck/ja3
- https://gist.github.com/husobee/6e9f998653d66f7481da
