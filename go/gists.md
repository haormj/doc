### 代码片段

#### 验证证书是否合法

```go
package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	b, err := ioutil.ReadFile("./echo-api-cert-6-ats.pem")
	if err != nil {
		log.Fatalln(err)
	}
	p, _ := pem.Decode(b)
	crt, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		log.Fatalln(err)
	}
	certPool := x509.NewCertPool()
	g4Bytes, err := ioutil.ReadFile("./ca.crt")
	if err != nil {
		log.Fatalln(err)
	}

	if ok := certPool.AppendCertsFromPEM(g4Bytes); !ok {
		log.Fatalln(err)
	}
	_, err = crt.Verify(x509.VerifyOptions{
		DNSName:     "echo-api.amazon.com",
		CurrentTime: time.Now().UTC(),
		Roots:       certPool,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
```

