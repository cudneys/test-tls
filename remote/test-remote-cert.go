package remote

import (
	"crypto/tls"
	"crypto/x509/pkix"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

type Cert struct {
	IssuerName   pkix.Name `json:"issuer-name"`
	Expiry       string    `json:"expiry"`
	CommonName   string    `json:"common-name"`
	DNSNames     []string  `json:"dns-names"`
	EmailAddrs   []string  `json:"email-addrs"`
	IPs          []net.IP  `json:"ips"`
	IssuingCerts []string  `json:"issuing-certs"`
}

func GetRemoteCert(addr string, protocol string) []Cert {
	var ret []Cert
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	//conn, err := tls.Dial(protocol, addr, conf)
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 2 * time.Second}, protocol, addr, conf)

	if err != nil {
		log.WithFields(log.Fields{"address": addr, "error": err}).Error("Failed To Connect")
		return ret
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		ret = append(
			ret,
			Cert{
				IssuerName:   cert.Issuer,
				Expiry:       cert.NotAfter.Format("2006-January-02"),
				CommonName:   cert.Issuer.CommonName,
				DNSNames:     cert.DNSNames,
				EmailAddrs:   cert.EmailAddresses,
				IPs:          cert.IPAddresses,
				IssuingCerts: cert.IssuingCertificateURL,
			},
		)
	}
	return ret
}
