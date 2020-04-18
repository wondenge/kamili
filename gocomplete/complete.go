// Package main is complete tool for the go command line
package main

import "github.com/wondenge/kamili"

var (
	ellipsis   = kamili.PredictSet("./...")
	anyPackage = kamili.PredictFunc(predictPackages)
	goFiles    = kamili.PredictFiles("*.go")
	anyFile    = kamili.PredictFiles("*")
	anyGo      = kamili.PredictOr(goFiles, anyPackage, ellipsis)
)

func main() {
	build := kamili.Command{
		Flags: kamili.Flags{
			"-o": anyFile,
			"-i": kamili.PredictNothing,

			"-a":             kamili.PredictNothing,
			"-n":             kamili.PredictNothing,
			"-p":             kamili.PredictAnything,
			"-race":          kamili.PredictNothing,
			"-msan":          kamili.PredictNothing,
			"-v":             kamili.PredictNothing,
			"-work":          kamili.PredictNothing,
			"-x":             kamili.PredictNothing,
			"-asmflags":      kamili.PredictAnything,
			"-buildmode":     kamili.PredictAnything,
			"-compiler":      kamili.PredictAnything,
			"-gccgoflags":    kamili.PredictSet("gccgo", "gc"),
			"-gcflags":       kamili.PredictAnything,
			"-installsuffix": kamili.PredictAnything,
			"-ldflags":       kamili.PredictAnything,
			"-linkshared":    kamili.PredictNothing,
			"-pkgdir":        anyPackage,
			"-tags":          kamili.PredictAnything,
			"-toolexec":      kamili.PredictAnything,
		},
		Args: anyGo,
	}

	run := kamili.Command{
		Flags: kamili.Flags{
			"-exec": kamili.PredictAnything,
		},
		Args: goFiles,
	}

	test := kamili.Command{
		Flags: kamili.Flags{
			"-args": kamili.PredictAnything,
			"-c":    kamili.PredictNothing,
			"-exec": kamili.PredictAnything,

			"-bench":     predictBenchmark,
			"-benchtime": kamili.PredictAnything,
			"-count":     kamili.PredictAnything,
			"-cover":     kamili.PredictNothing,
			"-covermode": kamili.PredictSet("set", "count", "atomic"),
			"-coverpkg":  kamili.PredictDirs("*"),
			"-cpu":       kamili.PredictAnything,
			"-run":       predictTest,
			"-short":     kamili.PredictNothing,
			"-timeout":   kamili.PredictAnything,

			"-benchmem":             kamili.PredictNothing,
			"-blockprofile":         kamili.PredictFiles("*.out"),
			"-blockprofilerate":     kamili.PredictAnything,
			"-coverprofile":         kamili.PredictFiles("*.out"),
			"-cpuprofile":           kamili.PredictFiles("*.out"),
			"-memprofile":           kamili.PredictFiles("*.out"),
			"-memprofilerate":       kamili.PredictAnything,
			"-mutexprofile":         kamili.PredictFiles("*.out"),
			"-mutexprofilefraction": kamili.PredictAnything,
			"-outputdir":            kamili.PredictDirs("*"),
			"-trace":                kamili.PredictFiles("*.out"),
		},
		Args: anyGo,
	}

	fmt := kamili.Command{
		Flags: kamili.Flags{
			"-n": kamili.PredictNothing,
			"-x": kamili.PredictNothing,
		},
		Args: anyGo,
	}

	get := kamili.Command{
		Flags: kamili.Flags{
			"-d":        kamili.PredictNothing,
			"-f":        kamili.PredictNothing,
			"-fix":      kamili.PredictNothing,
			"-insecure": kamili.PredictNothing,
			"-t":        kamili.PredictNothing,
			"-u":        kamili.PredictNothing,
		},
		Args: anyGo,
	}

	generate := kamili.Command{
		Flags: kamili.Flags{
			"-n":   kamili.PredictNothing,
			"-x":   kamili.PredictNothing,
			"-v":   kamili.PredictNothing,
			"-run": kamili.PredictAnything,
		},
		Args: anyGo,
	}

	vet := kamili.Command{
		Flags: kamili.Flags{
			"-n": kamili.PredictNothing,
			"-x": kamili.PredictNothing,
		},
		Args: anyGo,
	}

	list := kamili.Command{
		Flags: kamili.Flags{
			"-e":    kamili.PredictNothing,
			"-f":    kamili.PredictAnything,
			"-json": kamili.PredictNothing,
		},
		Args: kamili.PredictOr(anyPackage, ellipsis),
	}

	doc := kamili.Command{
		Flags: kamili.Flags{
			"-c":   kamili.PredictNothing,
			"-cmd": kamili.PredictNothing,
			"-u":   kamili.PredictNothing,
		},
		Args: anyPackage,
	}

	tool := kamili.Command{
		Flags: kamili.Flags{
			"-n": kamili.PredictNothing,
		},
		Sub: kamili.Commands{
			"addr2line": {
				Args: anyFile,
			},
			"asm": {
				Flags: kamili.Flags{
					"-D":        kamili.PredictAnything,
					"-I":        kamili.PredictDirs("*"),
					"-S":        kamili.PredictNothing,
					"-V":        kamili.PredictNothing,
					"-debug":    kamili.PredictNothing,
					"-dynlink":  kamili.PredictNothing,
					"-e":        kamili.PredictNothing,
					"-o":        anyFile,
					"-shared":   kamili.PredictNothing,
					"-trimpath": kamili.PredictNothing,
				},
				Args: kamili.PredictFiles("*.s"),
			},
			"cgo": {
				Flags: kamili.Flags{
					"-debug-define":      kamili.PredictNothing,
					"debug-gcc":          kamili.PredictNothing,
					"dynimport":          anyFile,
					"dynlinker":          kamili.PredictNothing,
					"dynout":             anyFile,
					"dynpackage":         anyPackage,
					"exportheader":       kamili.PredictDirs("*"),
					"gccgo":              kamili.PredictNothing,
					"gccgopkgpath":       kamili.PredictDirs("*"),
					"gccgoprefix":        kamili.PredictAnything,
					"godefs":             kamili.PredictNothing,
					"import_runtime_cgo": kamili.PredictNothing,
					"import_syscall":     kamili.PredictNothing,
					"importpath":         kamili.PredictDirs("*"),
					"objdir":             kamili.PredictDirs("*"),
					"srcdir":             kamili.PredictDirs("*"),
				},
				Args: goFiles,
			},
			"compile": {
				Flags: kamili.Flags{
					"-%":              kamili.PredictNothing,
					"-+":              kamili.PredictNothing,
					"-B":              kamili.PredictNothing,
					"-D":              kamili.PredictDirs("*"),
					"-E":              kamili.PredictNothing,
					"-I":              kamili.PredictDirs("*"),
					"-K":              kamili.PredictNothing,
					"-N":              kamili.PredictNothing,
					"-S":              kamili.PredictNothing,
					"-V":              kamili.PredictNothing,
					"-W":              kamili.PredictNothing,
					"-asmhdr":         anyFile,
					"-bench":          anyFile,
					"-buildid":        kamili.PredictNothing,
					"-complete":       kamili.PredictNothing,
					"-cpuprofile":     anyFile,
					"-d":              kamili.PredictNothing,
					"-dynlink":        kamili.PredictNothing,
					"-e":              kamili.PredictNothing,
					"-f":              kamili.PredictNothing,
					"-h":              kamili.PredictNothing,
					"-i":              kamili.PredictNothing,
					"-importmap":      kamili.PredictAnything,
					"-installsuffix":  kamili.PredictAnything,
					"-j":              kamili.PredictNothing,
					"-l":              kamili.PredictNothing,
					"-largemodel":     kamili.PredictNothing,
					"-linkobj":        anyFile,
					"-live":           kamili.PredictNothing,
					"-m":              kamili.PredictNothing,
					"-memprofile":     kamili.PredictNothing,
					"-memprofilerate": kamili.PredictAnything,
					"-msan":           kamili.PredictNothing,
					"-nolocalimports": kamili.PredictNothing,
					"-o":              anyFile,
					"-p":              kamili.PredictDirs("*"),
					"-pack":           kamili.PredictNothing,
					"-r":              kamili.PredictNothing,
					"-race":           kamili.PredictNothing,
					"-s":              kamili.PredictNothing,
					"-shared":         kamili.PredictNothing,
					"-traceprofile":   anyFile,
					"-trimpath":       kamili.PredictAnything,
					"-u":              kamili.PredictNothing,
					"-v":              kamili.PredictNothing,
					"-w":              kamili.PredictNothing,
					"-wb":             kamili.PredictNothing,
				},
				Args: goFiles,
			},
			"cover": {
				Flags: kamili.Flags{
					"-func": kamili.PredictAnything,
					"-html": kamili.PredictAnything,
					"-mode": kamili.PredictSet("set", "count", "atomic"),
					"-o":    anyFile,
					"-var":  kamili.PredictAnything,
				},
				Args: anyFile,
			},
			"dist": {
				Sub: kamili.Commands{
					"banner":    {Flags: kamili.Flags{"-v": kamili.PredictNothing}},
					"bootstrap": {Flags: kamili.Flags{"-v": kamili.PredictNothing}},
					"clean":     {Flags: kamili.Flags{"-v": kamili.PredictNothing}},
					"env":       {Flags: kamili.Flags{"-v": kamili.PredictNothing, "-p": kamili.PredictNothing}},
					"install":   {Flags: kamili.Flags{"-v": kamili.PredictNothing}, Args: kamili.PredictDirs("*")},
					"list":      {Flags: kamili.Flags{"-v": kamili.PredictNothing, "-json": kamili.PredictNothing}},
					"test":      {Flags: kamili.Flags{"-v": kamili.PredictNothing, "-h": kamili.PredictNothing}},
					"version":   {Flags: kamili.Flags{"-v": kamili.PredictNothing}},
				},
			},
			"doc": doc,
			"fix": {
				Flags: kamili.Flags{
					"-diff":  kamili.PredictNothing,
					"-force": kamili.PredictAnything,
					"-r":     kamili.PredictSet("context", "gotypes", "netipv6zone", "printerconfig"),
				},
				Args: anyGo,
			},
			"link": {
				Flags: kamili.Flags{
					"-B":              kamili.PredictAnything,  // note
					"-D":              kamili.PredictAnything,  // address (default -1)
					"-E":              kamili.PredictAnything,  // entry symbol name
					"-H":              kamili.PredictAnything,  // header type
					"-I":              kamili.PredictAnything,  // linker binary
					"-L":              kamili.PredictDirs("*"), // directory
					"-R":              kamili.PredictAnything,  // quantum (default -1)
					"-T":              kamili.PredictAnything,  // address (default -1)
					"-V":              kamili.PredictNothing,
					"-X":              kamili.PredictAnything,
					"-a":              kamili.PredictAnything,
					"-buildid":        kamili.PredictAnything, // build id
					"-buildmode":      kamili.PredictAnything,
					"-c":              kamili.PredictNothing,
					"-cpuprofile":     anyFile,
					"-d":              kamili.PredictNothing,
					"-debugtramp":     kamili.PredictAnything, // int
					"-dumpdep":        kamili.PredictNothing,
					"-extar":          kamili.PredictAnything,
					"-extld":          kamili.PredictAnything,
					"-extldflags":     kamili.PredictAnything, // flags
					"-f":              kamili.PredictNothing,
					"-g":              kamili.PredictNothing,
					"-importcfg":      anyFile,
					"-installsuffix":  kamili.PredictAnything, // dir suffix
					"-k":              kamili.PredictAnything, // symbol
					"-libgcc":         kamili.PredictAnything, // maybe "none"
					"-linkmode":       kamili.PredictAnything, // mode
					"-linkshared":     kamili.PredictNothing,
					"-memprofile":     anyFile,
					"-memprofilerate": kamili.PredictAnything, // rate
					"-msan":           kamili.PredictNothing,
					"-n":              kamili.PredictNothing,
					"-o":              kamili.PredictAnything,
					"-pluginpath":     kamili.PredictAnything,
					"-r":              kamili.PredictAnything, // "dir1:dir2:..."
					"-race":           kamili.PredictNothing,
					"-s":              kamili.PredictNothing,
					"-tmpdir":         kamili.PredictDirs("*"),
					"-u":              kamili.PredictNothing,
					"-v":              kamili.PredictNothing,
					"-w":              kamili.PredictNothing,
					// "-h":           kamili.PredictAnything, // halt on error
				},
				Args: kamili.PredictOr(
					kamili.PredictFiles("*.a"),
					kamili.PredictFiles("*.o"),
				),
			},
			"nm": {
				Flags: kamili.Flags{
					"-n":    kamili.PredictNothing,
					"-size": kamili.PredictNothing,
					"-sort": kamili.PredictAnything,
					"-type": kamili.PredictNothing,
				},
				Args: anyGo,
			},
			"objdump": {
				Flags: kamili.Flags{
					"-s": kamili.PredictAnything,
					"-S": kamili.PredictNothing,
				},
				Args: anyFile,
			},
			"pack": {
				/* this lacks the positional aspect of all these params */
				Flags: kamili.Flags{
					"c":  kamili.PredictNothing,
					"p":  kamili.PredictNothing,
					"r":  kamili.PredictNothing,
					"t":  kamili.PredictNothing,
					"x":  kamili.PredictNothing,
					"cv": kamili.PredictNothing,
					"pv": kamili.PredictNothing,
					"rv": kamili.PredictNothing,
					"tv": kamili.PredictNothing,
					"xv": kamili.PredictNothing,
				},
				Args: kamili.PredictOr(
					kamili.PredictFiles("*.a"),
					kamili.PredictFiles("*.o"),
				),
			},
			"pprof": {
				Flags: kamili.Flags{
					"-callgrind":     kamili.PredictNothing,
					"-disasm":        kamili.PredictAnything,
					"-dot":           kamili.PredictNothing,
					"-eog":           kamili.PredictNothing,
					"-evince":        kamili.PredictNothing,
					"-gif":           kamili.PredictNothing,
					"-gv":            kamili.PredictNothing,
					"-list":          kamili.PredictAnything,
					"-pdf":           kamili.PredictNothing,
					"-peek":          kamili.PredictAnything,
					"-png":           kamili.PredictNothing,
					"-proto":         kamili.PredictNothing,
					"-ps":            kamili.PredictNothing,
					"-raw":           kamili.PredictNothing,
					"-svg":           kamili.PredictNothing,
					"-tags":          kamili.PredictNothing,
					"-text":          kamili.PredictNothing,
					"-top":           kamili.PredictNothing,
					"-tree":          kamili.PredictNothing,
					"-web":           kamili.PredictNothing,
					"-weblist":       kamili.PredictAnything,
					"-output":        anyFile,
					"-functions":     kamili.PredictNothing,
					"-files":         kamili.PredictNothing,
					"-lines":         kamili.PredictNothing,
					"-addresses":     kamili.PredictNothing,
					"-base":          kamili.PredictAnything,
					"-drop_negative": kamili.PredictNothing,
					"-cum":           kamili.PredictNothing,
					"-seconds":       kamili.PredictAnything,
					"-nodecount":     kamili.PredictAnything,
					"-nodefraction":  kamili.PredictAnything,
					"-edgefraction":  kamili.PredictAnything,
					"-sample_index":  kamili.PredictNothing,
					"-mean":          kamili.PredictNothing,
					"-inuse_space":   kamili.PredictNothing,
					"-inuse_objects": kamili.PredictNothing,
					"-alloc_space":   kamili.PredictNothing,
					"-alloc_objects": kamili.PredictNothing,
					"-total_delay":   kamili.PredictNothing,
					"-contentions":   kamili.PredictNothing,
					"-mean_delay":    kamili.PredictNothing,
					"-runtime":       kamili.PredictNothing,
					"-focus":         kamili.PredictAnything,
					"-ignore":        kamili.PredictAnything,
					"-tagfocus":      kamili.PredictAnything,
					"-tagignore":     kamili.PredictAnything,
					"-call_tree":     kamili.PredictNothing,
					"-unit":          kamili.PredictAnything,
					"-divide_by":     kamili.PredictAnything,
					"-buildid":       kamili.PredictAnything,
					"-tools":         kamili.PredictDirs("*"),
					"-help":          kamili.PredictNothing,
				},
				Args: anyFile,
			},
			"tour": {
				Flags: kamili.Flags{
					"-http":        kamili.PredictAnything,
					"-openbrowser": kamili.PredictNothing,
				},
			},
			"trace": {
				Flags: kamili.Flags{
					"-http":  kamili.PredictAnything,
					"-pprof": kamili.PredictSet("net", "sync", "syscall", "sched"),
				},
				Args: anyFile,
			},
			"vet": {
				Flags: kamili.Flags{
					"-all":                 kamili.PredictNothing,
					"-asmdecl":             kamili.PredictNothing,
					"-assign":              kamili.PredictNothing,
					"-atomic":              kamili.PredictNothing,
					"-bool":                kamili.PredictNothing,
					"-buildtags":           kamili.PredictNothing,
					"-cgocall":             kamili.PredictNothing,
					"-composites":          kamili.PredictNothing,
					"-compositewhitelist":  kamili.PredictNothing,
					"-copylocks":           kamili.PredictNothing,
					"-httpresponse":        kamili.PredictNothing,
					"-lostcancel":          kamili.PredictNothing,
					"-methods":             kamili.PredictNothing,
					"-nilfunc":             kamili.PredictNothing,
					"-printf":              kamili.PredictNothing,
					"-printfuncs":          kamili.PredictAnything,
					"-rangeloops":          kamili.PredictNothing,
					"-shadow":              kamili.PredictNothing,
					"-shadowstrict":        kamili.PredictNothing,
					"-shift":               kamili.PredictNothing,
					"-structtags":          kamili.PredictNothing,
					"-tags":                kamili.PredictAnything,
					"-tests":               kamili.PredictNothing,
					"-unreachable":         kamili.PredictNothing,
					"-unsafeptr":           kamili.PredictNothing,
					"-unusedfuncs":         kamili.PredictAnything,
					"-unusedresult":        kamili.PredictNothing,
					"-unusedstringmethods": kamili.PredictAnything,
					"-v":                   kamili.PredictNothing,
				},
				Args: anyGo,
			},
		},
	}

	clean := kamili.Command{
		Flags: kamili.Flags{
			"-i":         kamili.PredictNothing,
			"-r":         kamili.PredictNothing,
			"-n":         kamili.PredictNothing,
			"-x":         kamili.PredictNothing,
			"-cache":     kamili.PredictNothing,
			"-testcache": kamili.PredictNothing,
			"-modcache":  kamili.PredictNothing,
		},
		Args: kamili.PredictOr(anyPackage, ellipsis),
	}

	env := kamili.Command{
		Args: kamili.PredictAnything,
	}

	bug := kamili.Command{}
	version := kamili.Command{}

	fix := kamili.Command{
		Args: anyGo,
	}

	modDownload := kamili.Command{
		Flags: kamili.Flags{
			"-json": kamili.PredictNothing,
		},
		Args: anyPackage,
	}

	modEdit := kamili.Command{
		Flags: kamili.Flags{
			"-fmt":    kamili.PredictNothing,
			"-module": kamili.PredictNothing,
			"-print":  kamili.PredictNothing,

			"-exclude":     anyPackage,
			"-dropexclude": anyPackage,
			"-replace":     anyPackage,
			"-dropreplace": anyPackage,
			"-require":     anyPackage,
			"-droprequire": anyPackage,
		},
		Args: kamili.PredictFiles("go.mod"),
	}

	modGraph := kamili.Command{}

	modInit := kamili.Command{
		Args: kamili.PredictAnything,
	}

	modTidy := kamili.Command{
		Flags: kamili.Flags{
			"-v": kamili.PredictNothing,
		},
	}

	modVendor := kamili.Command{
		Flags: kamili.Flags{
			"-v": kamili.PredictNothing,
		},
	}

	modVerify := kamili.Command{}

	modWhy := kamili.Command{
		Flags: kamili.Flags{
			"-m":      kamili.PredictNothing,
			"-vendor": kamili.PredictNothing,
		},
		Args: anyPackage,
	}

	modHelp := kamili.Command{
		Sub: kamili.Commands{
			"download": kamili.Command{},
			"edit":     kamili.Command{},
			"graph":    kamili.Command{},
			"init":     kamili.Command{},
			"tidy":     kamili.Command{},
			"vendor":   kamili.Command{},
			"verify":   kamili.Command{},
			"why":      kamili.Command{},
		},
	}

	mod := kamili.Command{
		Sub: kamili.Commands{
			"download": modDownload,
			"edit":     modEdit,
			"graph":    modGraph,
			"init":     modInit,
			"tidy":     modTidy,
			"vendor":   modVendor,
			"verify":   modVerify,
			"why":      modWhy,
			"help":     modHelp,
		},
	}

	help := kamili.Command{
		Sub: kamili.Commands{
			"bug":         kamili.Command{},
			"build":       kamili.Command{},
			"clean":       kamili.Command{},
			"doc":         kamili.Command{},
			"env":         kamili.Command{},
			"fix":         kamili.Command{},
			"fmt":         kamili.Command{},
			"generate":    kamili.Command{},
			"get":         kamili.Command{},
			"install":     kamili.Command{},
			"list":        kamili.Command{},
			"mod":         modHelp,
			"run":         kamili.Command{},
			"test":        kamili.Command{},
			"tool":        kamili.Command{},
			"version":     kamili.Command{},
			"vet":         kamili.Command{},
			"buildmode":   kamili.Command{},
			"c":           kamili.Command{},
			"cache":       kamili.Command{},
			"environment": kamili.Command{},
			"filetype":    kamili.Command{},
			"go.mod":      kamili.Command{},
			"gopath":      kamili.Command{},
			"gopath-get":  kamili.Command{},
			"goproxy":     kamili.Command{},
			"importpath":  kamili.Command{},
			"modules":     kamili.Command{},
			"module-get":  kamili.Command{},
			"packages":    kamili.Command{},
			"testflag":    kamili.Command{},
			"testfunc":    kamili.Command{},
		},
	}

	// commands that also accepts the build flags
	for name, options := range build.Flags {
		test.Flags[name] = options
		run.Flags[name] = options
		list.Flags[name] = options
		vet.Flags[name] = options
		get.Flags[name] = options
	}

	gogo := kamili.Command{
		Sub: kamili.Commands{
			"build":    build,
			"install":  build, // install and build have the same flags
			"run":      run,
			"test":     test,
			"fmt":      fmt,
			"get":      get,
			"generate": generate,
			"vet":      vet,
			"list":     list,
			"doc":      doc,
			"tool":     tool,
			"clean":    clean,
			"env":      env,
			"bug":      bug,
			"fix":      fix,
			"version":  version,
			"mod":      mod,
			"help":     help,
		},
		GlobalFlags: kamili.Flags{
			"-h": kamili.PredictNothing,
		},
	}

	kamili.New("go", gogo).Run()
}
