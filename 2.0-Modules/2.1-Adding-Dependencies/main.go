package main

import (
	"github.com/ardanlabs/conf"
)

/*
	"It's easier to add complexity than to remove it."
	Workflow for Go modules.
	What is a project?
	A Project is a github repo of code.
	If we are talking about application project like service project, we shouldn't think
	that the application project will only manage one binary because applications have
	too many binaries/services/tooling in play.
	More projects we have more opportunity is there to the
	different design philosophies/policies/guidelines in the company.
	One team should have one project in which they are working together regardless of number of services/tools that they need this maintains consistency not just related to development but also deployment.
	So every team could have their own project and every team could have their own
	philosophy, guidelines & policies.
*/
/*
	Say we want to use the "conf" package in our repo.
	Bill's workflow on using a new package.
	He adds the import of the package and then simply calls that package with a "New" function.
	It is possible that the package "conf" doesn't have a new function. But that is not important at this time. We syntactically pretend that here's a function called "New()".
	But the tooling doesn't clear up the import when we press "ctrl+S", save.
*/
/*
	go env will show us "GOMODCACHE"."GOMODCACHE" points to the module cache.
	Module cache is where all of our source code for any third party dependencies that we are downloading, so the go compile can build the source code.
*/
/*
	When we run vscode, something runs in the background called "gopls" (go please)
	"gopls" is the language server that the editor talks to, for all of the magic, that we see like "intellisense", showing us the errors (example, in imports).
	Editor itself is a blank canvas, it's the language server that makes the editor look smart.
	When the language server starts up, it actually stores in the memory a cache of the module cache. Language server has it's own sought of cache.
	"go mod tidy" is the next step. "go mod tidy" walks through our project looking to validate, that all the source code our project needs is on disk in module cache.
	"go mod tidy" goes to the link in the import, finds it, downloads it at some version 1.5.
	We will see that import, downloaded package in "GOMODCACHE/mod/github.com/ardanlabs/"
	Source code can be found in "GOMODCACHE/mod/github.com/ardanlabs/conf", this is the source code that "gopls" is using to give us intellisense.
*/

func main() {
	conf.New()
}
