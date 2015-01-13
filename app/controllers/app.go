package controllers

import (
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/jgcarvalho/PhageAnalysis/pep"
	"github.com/jgcarvalho/PhageAnalysis/prot"
	"github.com/jgcarvalho/PhageFinger/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Run() revel.Result {
	return c.Render()
}

func (c App) Running(seqFile *os.File, pepLib, forwPrimer, revPrimer, proteome, randLib string) revel.Result {
	// user := "zeh"
	// job := "rick"
	// id := "AAAAA"
	fmt.Println(pepLib, forwPrimer, revPrimer, proteome, randLib)
	rl, _ := strconv.Atoi(randLib)
	j := &models.Job{ID: "AAAAA", User: "zeh", SeqFile: seqFile, PepLib: pepLib, ForwPrimer: forwPrimer, RevPrimer: revPrimer, Proteome: proteome, RandLib: rl}
	ListJobs["AAAAA"] = j
	go j.Run()
	// for i := 0; i < 5; i++ {
	//	time.Sleep(6 * time.Second)
	// 	t := time.Now().String()
	// 	c.Render(t)
	// }
	return c.Redirect("/results/%s/%s/", j.User, j.ID)
}

var htmlResume string = `	<p>User: %s</p>
													<p>File name: %s</p>
													<p>Pep Lib: %s</p>
													<p>Forward Primer: %s</p>
													<p>Reverse Primer: %s</p>
													<p>Proteome: %s</p>
													<p>Random Libs: %d</p>
													<p><a href="./peptides/1">Peptides</a></p>
													<p><a href="./proteins/1">Proteins</a></p>
													<p><a href="./peptidesfasta">Peptides Fasta</a></p>
													<p><a href="./proteinsrank">Proteins txt</a></p>`

var htmlProgress string = `<div class="bs-component">
                            <div class="progress">
                              <div class="progress-bar" style="width: %d%%;"></div>
                            </div>
                        	 </div>`
var htmlRefresh string = `<meta http-equiv="refresh" content="5">`

func (c App) Results(id string) revel.Result {
	if ListJobs[id].Status == false {
		status := "Processing proteome"
		resume := template.HTML(fmt.Sprintf(htmlProgress, ListJobs[id].ProtProgress))
		refresh := template.HTML(htmlRefresh)
		fmt.Println(ListJobs[id].Status)
		return c.Render(status, resume, refresh)
	} else {
		status := "Results"
		resume := template.HTML(fmt.Sprintf(htmlResume, ListJobs[id].User, ListJobs[id].SeqFile, ListJobs[id].PepLib, ListJobs[id].ForwPrimer, ListJobs[id].RevPrimer, ListJobs[id].Proteome, ListJobs[id].RandLib))
		fmt.Println(ListJobs[id].Status)
		return c.Render(status, resume)
	}

}

func (c App) Peptides(user, id string, page int) revel.Result {
	nperpage := 25
	lenPages := (len(ListJobs[id].Peptides) - 1) / nperpage
	pages := make([]int, lenPages+1)
	for i := 0; i < lenPages+1; i++ {
		pages[i] = i + 1
	}
	var peptides []pep.Peptide
	if (page)*nperpage > len(ListJobs[id].Peptides) {
		peptides = ListJobs[id].Peptides[(page-1)*nperpage:]
	} else {
		peptides = ListJobs[id].Peptides[(page-1)*nperpage : (page)*nperpage]
	}

	next, previous := page+1, page-1
	return c.Render(peptides, page, next, previous, pages)
}

func (c App) Proteins(user, id string, page int) revel.Result {
	nperpage := 50
	lenPages := (len(ListJobs[id].Proteins) - 1) / nperpage
	pages := make([]int, lenPages+1)
	for i := 0; i < lenPages+1; i++ {
		pages[i] = i + 1
	}
	var proteins []prot.Protein
	if (page)*nperpage > len(ListJobs[id].Proteins) {
		proteins = ListJobs[id].Proteins[(page-1)*nperpage:]
	} else {
		proteins = ListJobs[id].Proteins[(page-1)*nperpage : (page)*nperpage]
	}

	next, previous := page+1, page-1
	return c.Render(proteins, page, next, previous, pages)
}

func (c App) PeptidesFasta(user, id string) revel.Result {
	peptides := ListJobs[id].Peptides
	return c.Render(peptides)
}

func (c App) ProteinsRank(user, id string) revel.Result {
	proteins := ListJobs[id].Proteins
	return c.Render(proteins)
}

var ListJobs map[string]*models.Job = make(map[string]*models.Job, 100)
