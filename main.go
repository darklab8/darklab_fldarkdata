package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/configs_fixtures"
	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/freelancer_mapped/exe_mapped"
	"github.com/darklab8/darklab_goutils/goutils/utils"
)

/*
for profiling
*/

func main() {
	f, err := os.Create("prof.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	utils.TimeMeasure(func() {
		game_location := configs_fixtures.FixtureGameLocation()
		config := exe_mapped.FixtureFLINIConfig()
		ids := exe_mapped.GetAllInfocards(game_location, config.Resources.Dll)

		for id, text := range ids {
			fmt.Println(id)
			fmt.Println(text)
			break
		}
	}, "measures time")
}
