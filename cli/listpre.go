/*
 * Warp (C) 2019-2020 MinIO, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package cli

import (
	"github.com/minio/cli"
	"github.com/minio/minio/pkg/console"
	"github.com/minio/warp/pkg/bench"
)

var (
	listpreFlags = []cli.Flag{
		cli.IntFlag{
			Name:  "objects",
			Value: 10000,
			Usage: "Number of objects to upload. Rounded to have equal concurrent objects.",
		},
		cli.StringFlag{
			Name:  "obj.size",
			Value: "1KB",
			Usage: "Size of each generated object. Can be a number or 10KB/MB/GB. All sizes are base 2 binary.",
		},
	}
)

var listpreCmd = cli.Command{
	Name:   "listpre",
	Usage:  "prepare list objects",
	Action: mainListpre,
	Before: setGlobalsFromContext,
	Flags:  combineFlags(globalFlags, ioFlags, listFlags, genFlags, benchFlags, analyzeFlags),
	CustomHelpTemplate: `NAME:
   {{.HelpName}} - {{.Usage}}
 
 USAGE:
   {{.HelpName}} [FLAGS]
   -> see https://github.com/minio/warp#list
 
 FLAGS:
   {{range .VisibleFlags}}{{.}}
   {{end}}`,
}

// mainDelete is the entry point for get command.
func mainListpre(ctx *cli.Context) error {
	checkListSyntax(ctx)
	src := newGenSource(ctx)

	b := bench.Listpre{
		Common: bench.Common{
			Client:      newClient(ctx),
			Concurrency: ctx.Int("concurrent"),
			Source:      src,
			Bucket:      ctx.String("bucket"),
			Location:    "",
			PutOpts:     putOpts(ctx),
		},
		CreateObjects: ctx.Int("objects"),
		NoPrefix:      ctx.Bool("noprefix"),
	} //b CreateObjects:10000
	return runBench(ctx, &b)
}

func checkListpreSyntax(ctx *cli.Context) {
	if ctx.NArg() > 0 {
		console.Fatal("Command takes no arguments")
	}

	checkAnalyze(ctx)
	checkBenchmark(ctx)
}
