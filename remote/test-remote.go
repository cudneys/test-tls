package remote

import (
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
	"time"
)

func Test(host string, port int, protocol string) error {
	if protocol == "" {
		protocol = "tcp"
	}

	llog := log.WithFields(log.Fields{"host": host, "port": port, "proto": protocol})

	addr := fmt.Sprintf("%s:%d", host, port)

	llog.WithField("addr", addr).Info("Testing Remote Host")

	conf := &tls.Config{
		InsecureSkipVerify: false,
	}

	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}, protocol, addr, conf)

	if err != nil {
		llog.WithFields(log.Fields{"error": err}).Error("Connection Unsuccessful!")

		if strings.HasSuffix(err.Error(), "i/o timeout") {
			return err
		}

		certs := GetRemoteCert(addr, protocol)
		for n, cert := range certs {
			clog := llog.WithField("cert_num", n)
			clog.WithFields(
				log.Fields{
					"field": "Issuer",
					"value": cert.IssuerName,
				},
			).Error("Faulty Cert Info")
			clog.WithFields(
				log.Fields{
					"field": "Expiry",
					"value": cert.Expiry,
				},
			).Error("Faulty Cert Info")
			clog.WithFields(
				log.Fields{
					"field": "CommonName",
					"value": cert.CommonName,
				},
			).Error("Faulty Cert Info")
			for sn, dns := range cert.DNSNames {
				clog.WithFields(
					log.Fields{
						"name_number": sn,
						"field":       "DNSName",
						"value":       dns,
					},
				).Error("Faulty Cert Info")
			}

			for sn, ip := range cert.IPs {
				clog.WithFields(
					log.Fields{
						"ip_number": sn,
						"field":     "IPAddr",
						"value":     ip,
					},
				).Error("Faulty Cert Info")
			}

			for sn, issuingCert := range cert.IssuingCerts {
				clog.WithFields(
					log.Fields{
						"issuer_number": sn,
						"field":         "Issuer",
						"value":         issuingCert,
					},
				).Error("Faulty Cert Info")
			}

			for sn, email := range cert.EmailAddrs {
				clog.WithFields(
					log.Fields{
						"email_number": sn,
						"field":        "Email",
						"value":        email,
					},
				).Error("Faulty Cert Info")
			}
		}

		return err
	}
	defer conn.Close()
	llog.WithFields(
		log.Fields{
			"remote_addr":      conn.RemoteAddr(),
			"cipher_suite":     conn.ConnectionState().CipherSuite,
			"negotiated_proto": conn.ConnectionState().NegotiatedProtocol,
			"version":          map_tls_version(conn.ConnectionState().Version),
		},
	).Info("Connection Established Successfully")
	return err
}
