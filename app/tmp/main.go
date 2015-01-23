// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/jgcarvalho/PhageFinger/app"
	controllers "github.com/jgcarvalho/PhageFinger/app/controllers"
	tests "github.com/jgcarvalho/PhageFinger/tests"
	controllers0 "github.com/revel/revel/modules/static/app/controllers"
	_ "github.com/revel/revel/modules/testrunner/app"
	controllers1 "github.com/revel/revel/modules/testrunner/app/controllers"
	os "os"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					20: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					24: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Running",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "seqFile", Type: reflect.TypeOf((**os.File)(nil)) },
					&revel.MethodArg{Name: "pepLib", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "forwPrimer", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "revPrimer", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "proteome", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "randLib", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Results",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					71: []string{ 
						"status",
						"resume",
						"refresh",
					},
					76: []string{ 
						"status",
						"resume",
					},
				},
			},
			&revel.MethodType{
				Name: "Peptides",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "page", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					96: []string{ 
						"peptides",
						"page",
						"next",
						"previous",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Proteins",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "page", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					114: []string{ 
						"proteins",
						"page",
						"next",
						"previous",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "ProtDetails",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "protid", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					119: []string{ 
						"protein",
					},
				},
			},
			&revel.MethodType{
				Name: "PeptidesFasta",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					124: []string{ 
						"peptides",
					},
				},
			},
			&revel.MethodType{
				Name: "ProteinsRank",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					129: []string{ 
						"proteins",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					48: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					78: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	revel.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
