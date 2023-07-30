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
 *	go run script/generate.go
 *
 * # or use just
 * just gen
 * ```
 */

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
