package models

func Proteome(htmloption string) (filename string) {

	switch htmloption {
	case "prot-sample":
		filename = "uniprot-sample.fasta"
	case "rick-rick-br":
		filename = "uniprot-proteome%3AUP000007877.fasta"
	}
	return
}
