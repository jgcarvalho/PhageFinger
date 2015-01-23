package models

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/jgcarvalho/PhageAnalysis/pda"
	"github.com/jgcarvalho/PhageAnalysis/pep"
	"github.com/jgcarvalho/PhageAnalysis/prot"
)

type MyJob struct {
	ID         string
	User       string
	Status     bool
	SeqFile    *os.File
	PepLib     string
	ForwPrimer string
	RevPrimer  string
	Proteome   string
	RandLib    int
}

func (j *MyJob) Run() {

	fmt.Println("func running")
	time.Sleep(15 * time.Second)
	fmt.Println("func running yet")
	j.Status = true
	fmt.Println("ok")

}

type Job struct {
	ID   string
	User string
	// PepStatus    bool
	// RandStatus   bool
	ProtProgress int
	Status       bool
	SeqFile      *os.File
	PepLib       string
	ForwPrimer   string
	RevPrimer    string
	Proteome     string
	RandLib      int
	Peptides     []pep.Peptide
	Proteins     []prot.Protein
}

func (j *Job) Run() {
	fmt.Println("Running")
	var PepLen int
	switch j.PepLib {
	case "pep7nnk":
		PepLen = 7 * 3
	case "pep12nnk":
		PepLen = 12 * 3
	}
	template := pda.CreateTemplate(j.ForwPrimer, j.RevPrimer, PepLen)
	dna, err := pda.ReadMDna(j.SeqFile)
	if err != nil {
		fmt.Println("ERRO", err)
	}
	// fmt.Println(template)
	// fmt.Println(dna)
	// peptides, unreliable := pda.GetPeptides(dna, template)
	j.Peptides, _ = pda.GetPeptides(dna, template)
	// fmt.Println(j.Peptides)
	randomPeps := pep.RandomLibrary(j.Peptides)
	// proteomeDir deve apontar para a pasta onde estão os proteomas
	proteomeDir := "/home/zeh/gocode/src/github.com/jgcarvalho/PhageFinger/public/proteomes/"
	j.Proteins, err = pda.ReadMProteins(proteomeDir + j.Proteome)
	for i := 0; i < len(j.Proteins); i++ {
		j.Proteins[i].Analysis(j.Peptides, randomPeps)
		j.ProtProgress = (i + 2) * 100 / len(j.Proteins)
	}
	sort.Sort(sort.Reverse(prot.Proteins(j.Proteins)))
	fmt.Println("DONE")

	j.Status = true
}
