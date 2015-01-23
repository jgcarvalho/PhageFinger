// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Run(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Run", args).Url
}

func (_ tApp) Running(
		seqFile interface{},
		pepLib string,
		forwPrimer string,
		revPrimer string,
		proteome string,
		randLib string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "seqFile", seqFile)
	revel.Unbind(args, "pepLib", pepLib)
	revel.Unbind(args, "forwPrimer", forwPrimer)
	revel.Unbind(args, "revPrimer", revPrimer)
	revel.Unbind(args, "proteome", proteome)
	revel.Unbind(args, "randLib", randLib)
	return revel.MainRouter.Reverse("App.Running", args).Url
}

func (_ tApp) Results(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.Results", args).Url
}

func (_ tApp) Peptides(
		user string,
		id string,
		page int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("App.Peptides", args).Url
}

func (_ tApp) Proteins(
		user string,
		id string,
		page int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("App.Proteins", args).Url
}

func (_ tApp) ProtDetails(
		user string,
		id string,
		protid int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "protid", protid)
	return revel.MainRouter.Reverse("App.ProtDetails", args).Url
}

func (_ tApp) PeptidesFasta(
		user string,
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.PeptidesFasta", args).Url
}

func (_ tApp) ProteinsRank(
		user string,
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.ProteinsRank", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


