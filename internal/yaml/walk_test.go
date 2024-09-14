package yaml_test

import (
	"fmt"
	"strings"

	"github.com/rwxrob/cobracli/internal/yaml"
	yamlv3 "gopkg.in/yaml.v3"
)

func ExampleWalk() {

	docs, err := yaml.ParseString(`
%YAML 1.1
---
title: YAML Ain't Markup Language
foo: bar
some other map:
   title: Another Title
`)
	if err != nil {
		fmt.Println(err)
		return
	}

	// append trademark to end of any map value with key matching `title`
	aWalkFunc := func(n, f *yamlv3.Node, i int) error {
		if yaml.IsMapVal(n, f, i) && yaml.MapKeyFor(n, f, i) == `title` {
			n.Value = n.Value + `™`
		}
		return nil
	}

	if err := yaml.Walk(docs[0], aWalkFunc); err != nil {
		fmt.Println(err)
		return
	}

	yaml.Print(docs[0])

	// Output:
	// title: YAML Ain't Markup Language™
	// foo: bar
	// some other map:
	//     title: Another Title™
	//

}

func ExampleWalk_print_Aliases() {

	docs, err := yaml.ParseString(`
services:
  production-db: &database-definition
    image: mysql:5.7
    volumes:
      - db_data:/var/lib/mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: somewordpress
      MYSQL_DATABASE: wordpress
      MYSQL_USER: wordpress
      MYSQL_PASSWORD: wordpress
  test-db: *database-definition
`)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Walk(docs[0], func(n, f *yamlv3.Node, i int) error {
		if yaml.IsAlias(n) {
			fmt.Println(n.Value)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Output:
	// database-definition

}

func ExampleWalk_collect_Tags() {

	docs, err := yaml.ParseString(`
people:
  - name: Rob
    tags:
      - mentor
      - employed
  - name: Doris
    tags: employed artist
    pets:
      - name: Sam
        tag: dog
`)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(docs) == 0 {
		fmt.Println(`empty document slice returned`)
		return
	}

	// collection of tags with the number of times tag occurs
	tags := map[string]int{}

	// add every unique tag used anywhere in the document
	collectTags := func(n, f *yamlv3.Node, i int) error {
		if !yaml.IsMapVal(n, f, i) {
			return nil
		}
		switch v := yaml.MapKeyFor(n, f, i); v {
		case `tag`:
			tags[n.Value]++
		case `tags`:
			switch {
			case yaml.IsString(n):
				for _, tag := range strings.Fields(n.Value) {
					tags[tag]++
				}
			case yaml.IsSequence(n):
				for _, assumedScalar := range n.Content {
					tags[assumedScalar.Value]++
				}
			}
		}
		return nil
	}

	if err := yaml.Walk(docs[0], collectTags); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(yaml.ToString(tags))

	// Unordered output:
	// artist: 1
	// dog: 1
	// employed: 2
	// mentor: 1

}
