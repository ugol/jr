//go:build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	/*
		fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))


		if err != nil {
			panic(err)
		}
		fmt.Printf("  cwd = %s\n", cwd)
		fmt.Printf("  os.Args = %#v\n", os.Args)

		for _, ev := range []string{"GOARCH", "GOOS", "GOFILE", "GOLINE", "GOPACKAGE", "DOLLAR"} {
			fmt.Println("  ", ev, "=", os.Getenv(ev))
		}
	*/

	tpl := `
package {{.package}}

func getInstance() {
  var o {{.type}}
  return &o
}

{{range $item, $element := array $iterations }}{{ $currentIteration := add $item 1 }}
    {{randoms "{\"product_id\":66,\"category\":\"calzone\",\"quantity\":4,\"unit_price\":18.78,\"net_price\":75.12}|{\"product_id\":71,\"category\":\"salad\",\"quantity\":4,\"unit_price\":4.9,\"net_price\":19.6}|{\"product_id\":79,\"category\":\"pizza\",\"quantity\":3,\"unit_price\":19.61,\"net_price\":58.83}|{\"product_id\":38,\"category\":\"calzone\",\"quantity\":2,\"unit_price\":13.05,\"net_price\":26.1}|{\"product_id\":26,\"category\":\"calzone\",\"quantity\":1,\"unit_price\":17.37,\"net_price\":17.37}|{\"product_id\":66,\"category\":\"calzone\",\"quantity\":3,\"unit_price\":24.63,\"net_price\":73.89}|{\"product_id\":62,\"category\":\"dessert\",\"quantity\":3,\"unit_price\":21.14,\"net_price\":63.42}|{\"product_id\":68,\"category\":\"salad\",\"quantity\":3,\"unit_price\":21.13,\"net_price\":63.39}|{\"product_id\":72,\"category\":\"salad\",\"quantity\":2,\"unit_price\":8.46,\"net_price\":16.92}|{\"product_id\":55,\"category\":\"calzone\",\"quantity\":2,\"unit_price\":3.6,\"net_price\":7.2}|{\"product_id\":77,\"category\":\"pizza\",\"quantity\":4,\"unit_price\":19.01,\"net_price\":76.04}|{\"product_id\":60,\"category\":\"pizza\",\"quantity\":3,\"unit_price\":1.49,\"net_price\":4.47}|{\"product_id\":8,\"category\":\"dessert\",\"quantity\":3,\"unit_price\":23.33,\"net_price\":69.99}|{\"product_id\":91,\"category\":\"calzone\",\"quantity\":4,\"unit_price\":11.24,\"net_price\":44.96}|{\"product_id\":78,\"category\":\"pizza\",\"quantity\":2,\"unit_price\":11.45,\"net_price\":22.9}|{\"product_id\":27,\"category\":\"dessert\",\"quantity\":3,\"unit_price\":22.02,\"net_price\":66.06}|{\"product_id\":58,\"category\":\"salad\",\"quantity\":1,\"unit_price\":9.88,\"net_price\":9.88}|{\"product_id\":9,\"category\":\"dessert\",\"quantity\":4,\"unit_price\":24.69,\"net_price\":98.76}|{\"product_id\":77,\"category\":\"salad\",\"quantity\":1,\"unit_price\":12.62,\"net_price\":12.62}|{\"product_id\":27,\"category\":\"wings\",\"quantity\":3,\"unit_price\":2.21,\"net_price\":6.63}|{\"product_id\":4,\"category\":\"calzone\",\"quantity\":1,\"unit_price\":19.04,\"net_price\":19.04}|{\"product_id\":48,\"category\":\"calzone\",\"quantity\":1,\"unit_price\":16.38,\"net_price\":16.38}|{\"product_id\":26,\"category\":\"dessert\",\"quantity\":5,\"unit_price\":9.24,\"net_price\":46.2}|{\"product_id\":75,\"category\":\"pizza\",\"quantity\":5,\"unit_price\":16.09,\"net_price\":80.45}|{\"product_id\":33,\"category\":\"salad\",\"quantity\":3,\"unit_price\":10.88,\"net_price\":32.64}|{\"product_id\":45,\"category\":\"calzone\",\"quantity\":5,\"unit_price\":22.64,\"net_price\":113.2}|{\"product_id\":65,\"category\":\"dessert\",\"quantity\":5,\"unit_price\":2.63,\"net_price\":13.15}|{\"product_id\":52,\"category\":\"wings\",\"quantity\":1,\"unit_price\":5.68,\"net_price\":5.68}|{\"product_id\":23,\"category\":\"pizza\",\"quantity\":3,\"unit_price\":5.49,\"net_price\":16.47}|{\"product_id\":67,\"category\":\"dessert\",\"quantity\":1,\"unit_price\":3.02,\"net_price\":3.02}|{\"product_id\":71,\"category\":\"dessert\",\"quantity\":4,\"unit_price\":16.69,\"net_price\":66.76}|{\"product_id\":69,\"category\":\"salad\",\"quantity\":3,\"unit_price\":17.07,\"net_price\":51.21}|{\"product_id\":41,\"category\":\"pizza\",\"quantity\":4,\"unit_price\":11.55,\"net_price\":46.2}|{\"product_id\":44,\"category\":\"wings\",\"quantity\":2,\"unit_price\":6.59,\"net_price\":13.18}|{\"product_id\":66,\"category\":\"wings\",\"quantity\":2,\"unit_price\":16.54,\"net_price\":33.08}|{\"product_id\":8,\"category\":\"calzone\",\"quantity\":5,\"unit_price\":12.07,\"net_price\":60.35}|{\"product_id\":43,\"category\":\"salad\",\"quantity\":5,\"unit_price\":13.9,\"net_price\":69.5}|{\"product_id\":13,\"category\":\"calzone\",\"quantity\":5,\"unit_price\":20.3,\"net_price\":101.5}|{\"product_id\":65,\"category\":\"salad\",\"quantity\":5,\"unit_price\":16.61,\"net_price\":83.05}|{\"product_id\":48,\"category\":\"dessert\",\"quantity\":1,\"unit_price\":4.18,\"net_price\":4.18}|{\"product_id\":59,\"category\":\"calzone\",\"quantity\":3,\"unit_price\":5.74,\"net_price\":17.22}|{\"product_id\":42,\"category\":\"calzone\",\"quantity\":1,\"unit_price\":15.65,\"net_price\":15.65}|{\"product_id\":22,\"category\":\"calzone\",\"quantity\":5,\"unit_price\":2.79,\"net_price\":13.95}|{\"product_id\":47,\"category\":\"pizza\",\"quantity\":2,\"unit_price\":15.35,\"net_price\":30.7}|{\"product_id\":72,\"category\":\"salad\",\"quantity\":5,\"unit_price\":24.97,\"net_price\":124.85}|{\"product_id\":60,\"category\":\"salad\",\"quantity\":1,\"unit_price\":2.85,\"net_price\":2.85}|{\"product_id\":58,\"category\":\"dessert\",\"quantity\":1,\"unit_price\":1.24,\"net_price\":1.24}|{\"product_id\":88,\"category\":\"calzone\",\"quantity\":4,\"unit_price\":24.0,\"net_price\":96.0}|{\"product_id\":98,\"category\":\"wings\",\"quantity\":1,\"unit_price\":6.66,\"net_price\":6.66}|{\"product_id\":37,\"category\":\"pizza\",\"quantity\":3,\"unit_price\":14.62,\"net_price\":43.86}|{\"product_id\":37,\"category\":\"calzone\",\"quantity\":3,\"unit_price\":14.21,\"net_price\":42.63}|{\"product_id\":54,\"category\":\"wings\",\"quantity\":1,\"unit_price\":18.88,\"net_price\":18.88}|{\"product_id\":52,\"category\":\"wings\",\"quantity\":3,\"unit_price\":9.47,\"net_price\":28.41}|{\"product_id\":80,\"category\":\"wings\",\"quantity\":1,\"unit_price\":13.39,\"net_price\":13.39}|{\"product_id\":6,\"category\":\"pizza\",\"quantity\":3,\"unit_price\":13.64,\"net_price\":40.92}|{\"product_id\":95,\"category\":\"wings\",\"quantity\":1,\"unit_price\":24.9,\"net_price\":24.9}|{\"product_id\":60,\"category\":\"pizza\",\"quantity\":5,\"unit_price\":7.63,\"net_price\":38.15}|{\"product_id\":81,\"category\":\"wings\",\"quantity\":4,\"unit_price\":20.31,\"net_price\":81.24}|{\"product_id\":42,\"category\":\"dessert\",\"quantity\":1,\"unit_price\":6.61,\"net_price\":6.61}|{\"product_id\":6,\"category\":\"salad\",\"quantity\":4,\"unit_price\":7.66,\"net_price\":30.64}|{\"product_id\":36,\"category\":\"calzone\",\"quantity\":4,\"unit_price\":3.23,\"net_price\":12.92}|{\"product_id\":27,\"category\":\"calzone\",\"quantity\":1,\"unit_price\":8.64,\"net_price\":8.64}|{\"product_id\":97,\"category\":\"salad\",\"quantity\":5,\"unit_price\":6.33,\"net_price\":31.65}|{\"product_id\":95,\"category\":\"salad\",\"quantity\":4,\"unit_price\":14.89,\"net_price\":59.56}|{\"product_id\":89,\"category\":\"pizza\",\"quantity\":1,\"unit_price\":8.81,\"net_price\":8.81}|{\"product_id\":97,\"category\":\"dessert\",\"quantity\":3,\"unit_price\":16.03,\"net_price\":48.09}|{\"product_id\":25,\"category\":\"salad\",\"quantity\":4,\"unit_price\":3.17,\"net_price\":12.68}|{\"product_id\":42,\"category\":\"wings\",\"quantity\":2,\"unit_price\":5.47,\"net_price\":10.94}|{\"product_id\":7,\"category\":\"salad\",\"quantity\":2,\"unit_price\":18.25,\"net_price\":36.5}|{\"product_id\":4,\"category\":\"wings\",\"quantity\":4,\"unit_price\":16.46,\"net_price\":65.84}|{\"product_id\":78,\"category\":\"calzone\",\"quantity\":1,\"unit_price\":2.15,\"net_price\":2.15}|{\"product_id\":71,\"category\":\"calzone\",\"quantity\":1,\"unit_price\":22.3,\"net_price\":22.3}|{\"product_id\":80,\"category\":\"dessert\",\"quantity\":5,\"unit_price\":12.69,\"net_price\":63.45}|{\"product_id\":83,\"category\":\"salad\",\"quantity\":1,\"unit_price\":1.66,\"net_price\":1.66}|{\"product_id\":44,\"category\":\"dessert\",\"quantity\":4,\"unit_price\":18.18,\"net_price\":72.72}|{\"product_id\":43,\"category\":\"calzone\",\"quantity\":4,\"unit_price\":2.01,\"net_price\":8.04}|{\"product_id\":91,\"category\":\"calzone\",\"quantity\":3,\"unit_price\":13.06,\"net_price\":39.18}|{\"product_id\":55,\"category\":\"wings\",\"quantity\":2,\"unit_price\":22.1,\"net_price\":44.2}|{\"product_id\":88,\"category\":\"wings\",\"quantity\":1,\"unit_price\":5.33,\"net_price\":5.33}|{\"product_id\":10,\"category\":\"dessert\",\"quantity\":2,\"unit_price\":8.68,\"net_price\":17.36}|{\"product_id\":76,\"category\":\"salad\",\"quantity\":2,\"unit_price\":1.39,\"net_price\":2.78}|{\"product_id\":13,\"category\":\"salad\",\"quantity\":5,\"unit_price\":17.69,\"net_price\":88.45}|{\"product_id\":27,\"category\":\"salad\",\"quantity\":4,\"unit_price\":9.22,\"net_price\":36.88}|{\"product_id\":94,\"category\":\"wings\",\"quantity\":3,\"unit_price\":7.16,\"net_price\":21.48}|{\"product_id\":89,\"category\":\"dessert\",\"quantity\":1,\"unit_price\":20.2,\"net_price\":20.2}|{\"product_id\":4,\"category\":\"wings\",\"quantity\":4,\"unit_price\":3.52,\"net_price\":14.08}|{\"product_id\":97,\"category\":\"dessert\",\"quantity\":4,\"unit_price\":6.62,\"net_price\":26.48}|{\"product_id\":57,\"category\":\"dessert\",\"quantity\":2,\"unit_price\":20.44,\"net_price\":40.88}|{\"product_id\":6,\"category\":\"dessert\",\"quantity\":4,\"unit_price\":10.59,\"net_price\":42.36}|{\"product_id\":86,\"category\":\"salad\",\"quantity\":3,\"unit_price\":24.6,\"net_price\":73.8}|{\"product_id\":85,\"category\":\"salad\",\"quantity\":3,\"unit_price\":11.83,\"net_price\":35.49}|{\"product_id\":95,\"category\":\"wings\",\"quantity\":3,\"unit_price\":15.27,\"net_price\":45.81}|{\"product_id\":18,\"category\":\"calzone\",\"quantity\":3,\"unit_price\":11.25,\"net_price\":33.75}|{\"product_id\":85,\"category\":\"wings\",\"quantity\":3,\"unit_price\":13.53,\"net_price\":40.59}|{\"product_id\":13,\"category\":\"wings\",\"quantity\":4,\"unit_price\":16.58,\"net_price\":66.32}|{\"product_id\":12,\"category\":\"salad\",\"quantity\":5,\"unit_price\":9.19,\"net_price\":45.95}|{\"product_id\":21,\"category\":\"pizza\",\"quantity\":1,\"unit_price\":13.65,\"net_price\":13.65}|{\"product_id\":14,\"category\":\"dessert\",\"quantity\":2,\"unit_price\":18.32,\"net_price\":36.64}|{\"product_id\":7,\"category\":\"calzone\",\"quantity\":5,\"unit_price\":23.88,\"net_price\":119.4}|{\"product_id\":71,\"category\":\"pizza\",\"quantity\":1,\"unit_price\":22.52,\"net_price\":22.52}"}}{{if lt $currentIteration $iterations }},{{end}}
{{end}}
`

	fmt.Println(tpl)

	var typesList = []string{}
	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && isAGoFile(path) && !isAGenerateFile(path) {
			name := strings.Split(filepath.Base(path), ".")
			typesList = append(typesList, name[0])
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(typesList)

}

func isAGoFile(path string) bool {
	return strings.HasSuffix(path, "go")
}

func isAGenerateFile(path string) bool {
	return strings.HasSuffix(path, "instance.go") || strings.HasSuffix(path, "generate.go")
}
