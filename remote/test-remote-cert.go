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
	llog := log.WithFields(log.Fields{"address": addr, "protocol": protocol})
	var ret []Cert
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	llog.WithFields(log.Fields{"InsecureSkipVerify": conf.InsecureSkipVerify}).Debug("Connecting To Host")
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 2 * time.Second}, protocol, addr, conf)
	if err != nil {
		log.WithFields(log.Fields{"address": addr, "error": err}).Error("Failed To Connect")
		return ret
	} else {
		log.WithFields(log.Fields{"server_name": conn.ConnectionState().ServerName}).Debug("Connection Established")
	}
	defer conn.Close()

	llog.Debug("Getting Peer Certificateds")
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		llog.WithFields(log.Fields{"common_name": cert.Issuer.CommonName}).Debug("Found Peer Cert")
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
	llog.Debug("Finished Collecting Peer Certs")
	return ret
}
