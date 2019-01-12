// Copyright (c) 2019 Autoblocks <deslittle@gmail.com>
// All rights reserved. Use of this source code is governed by a LGPL v3
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"os"
	"encoding/json"
	"fmt"
	"ioutil"
	"github.com/Autoblocks/go-dsl"
	"github.com/Autoblocks/autoblocks-parser/ab"
	"github.com/Autoblocks/autoblocks-parser/il"
	"github.com/Autoblocks/autoblocks-parser/st"
	"github.com/Autoblocks/autoblocks-parser/sfc"
	"github.com/Autoblocks/autoblocks-parser/ld"
	"github.com/Autoblocks/autoblocks-parser/fbd"
)

func main() {

	flag.Parse()

	switch flag.NArg() {
	case 0:
		scanFn := ab.Scan
		parseFn := ab.Parse
		ts := ab.NewTokenSet()
		ns := ab.NewNodeSet()
		break
	case 1:
		arg, _ = ioutil.ReadFile(flag.Arg(0))
		if arg == "AB"{
			scanFn := ab.Scan
			parseFn := ab.Parse
			ts := ab.NewTokenSet()
			ns := ab.NewNodeSet()
		}else if arg == "IL"{
			scanFn := il.Scan
			parseFn := il.Parse
			ts := il.NewTokenSet()
			ns := il.NewNodeSet()
		}else if arg == "ST"{
			scanFn := st.Scan
			parseFn := st.Parse
			ts := st.NewTokenSet()
			ns := st.NewNodeSet()
		}else if arg == "SFC"{
			scanFn := sfc.Scan
			parseFn := sfc.Parse
			ts := sfc.NewTokenSet()
			ns := sfc.NewNodeSet()
		}else if arg == "LD"{
			scanFn := ld.Scan
			parseFn := ld.Parse
			ts := ld.NewTokenSet()
			ns := ld.NewNodeSet()
		}else if arg == "FBD"{
			scanFn := fbd.Scan
			parseFn := fbd.Parse
			ts := fbd.NewTokenSet()
			ns := fbd.NewNodeSet()
		}else{
			fmt.Printf("Language type muast be one of: AB, IL, ST, SFC, LD, FBD\n")
			os.Exit(1)
		}
		break
	default:
		fmt.Printf("Input must be from stdin or file\n")
		os.Exit(1)
	}

	
	reader := bufio.NewReader(os.Stdin)
	bufreader := bufio.NewReader(reader)

    logfilename := "log.txt"
    logfile, err := os.Create(logfilename)
    if err != nil {
		fmt.Println("Error: Could not create log file " + logfilename + ": " + err.Error())
        os.Exit(1)
	}
	ast, errs := dsl.ParseAndLog(parseFn, scanFn, ts, ns, bufreader, logfile)
    logfile.Close()
	outJson, _ := json.Marshal(ast.RootNode)
	outString := string(outJson)
	if errs != nil {
		for _, err := range errs {
			fmt.Print(err.String())
		}
	}else{
		fmt.Println(outString)
	}
    

}



