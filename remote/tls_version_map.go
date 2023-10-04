package remote

func map_tls_version(verNum uint16) float32 {
	switch verNum {
	case 769:
		return 1.0
	case 770:
		return 1.1
	case 771:
		return 1.2
	case 772:
		return 1.3
	}
	return 0.0
}
