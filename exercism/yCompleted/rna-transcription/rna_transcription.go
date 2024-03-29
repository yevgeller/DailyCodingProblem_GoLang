package strand

func ToRNA(dna string) string {
	ret := ""
	for _, nuc := range dna {
		switch nuc {
		case 'A':
			ret += "U"
			break
		case 'C':
			ret += "G"
			break
		case 'G':
			ret += "C"
			break
		case 'T':
			ret += "A"
			break
		default:
			break
		}

	}
	return ret
}
