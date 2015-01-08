package controllers

import (
	"fmt"
	"time"

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

func (c App) Results(id string) revel.Result {
	if ListJobs[id].status == false {
		done := "processing"
		fmt.Println(ListJobs[id].status)
		return c.Render(done)
	} else {
		done := "Done"
		fmt.Println(ListJobs[id].status)
		return c.Render(done)
	}

}

func (c App) Running() revel.Result {
	// user := "zeh"
	// job := "rick"
	// id := "AAAAA"
	j := &MyJob{id: "AAAAA"}
	ListJobs["AAAAA"] = j
	go j.Run()
	// for i := 0; i < 5; i++ {
	//	time.Sleep(6 * time.Second)
	// 	t := time.Now().String()
	// 	c.Render(t)
	// }
	return c.Redirect("/results?id=%s", j.id)
}

var ListJobs map[string]*MyJob = make(map[string]*MyJob, 100)

type MyJob struct {
	id string
	// user string
	// job string
	status bool
}

func (j *MyJob) Run() {

	fmt.Println("func running")
	time.Sleep(15 * time.Second)
	fmt.Println("func running yet")
	j.status = true
	fmt.Println("ok")

}
