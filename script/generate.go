/*
 * This is a script that auto generates a json file containing all
 * versions of the dart sdk from the api. Creating a pre-generated
 * json file reduces the need to make a req to the api everytime.
 *
 * This process should be automated in some sort of way, for example
 * using github action's cron jobs
 *
 * Run the script from the root
 * ```
 *   $ go run script/generate.go
 *   # or use just
 *   $ just gen
 * ```
 */

package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/Yakiyo/dinx/utils"
	json "github.com/json-iterator/go"
)

func main() {
	var wg sync.WaitGroup
	m := sync.Map{}

	for _, chn := range utils.Channels {
		wg.Add(1)
		chn := chn
		go func() {
			fmt.Printf("Running goroutine for %v...\n", chn)
			defer wg.Done()
			do(chn, &m)
		}()
	}

	wg.Wait()

	plain := map[string][]string{}

	fmt.Println("Making map...")
	for _, chn := range utils.Channels {
		v, _ := m.Load(chn)
		plain[chn] = v.([]string)
	}

	out, err := json.MarshalIndent(plain, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println("Got json. Writing to file...")
	err = os.WriteFile("versions.json", out, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully ran script")
}

// This is the main func that runs thrice, once for every channel. We use
// goroutines to run them all, then wait on them to get the data and finally
// write em to the file
func do(channel string, m *sync.Map) {
	content := fetch(channel)
	stripped := []string{}

	for _, v := range content {
		strip := stripVers(v)
		stripped = append(stripped, strip)
	}
	m.Store(channel, stripped)
}
