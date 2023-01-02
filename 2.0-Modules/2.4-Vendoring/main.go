package main

/*
	"go mod tidy" is great what it is doing is putting all the code in our module cache(local).
	So that the compiler/go tooling can find the code and we can build our binaries.
	There is one more thing we can do that adds more durability to our software,
	it's called "Vendoring".
	Idea of Vendoring is not to use the source code sitting in the module cache, but to bring
	all the 3rd party dependency code into the project.
	So the project own every line of code it needs to build.
	As per Bill we should vendor unless it's not reasonable or practical to do so.

	When is it not practical?
	When the amount of code that we manage in the vendor folder becomes huge.
	If we are building Kubernetes or docker then, there may be just too much
	dependency code for us to be able to match it in the project.

	We have our "go.mod" see 1.png.

	After we do "go mod tidy", we have to run the command,
	"go mod vendor". There is no output of that command, but we can see a
	"vendor" folder.
	Inside the vendor folder we now have all the source code for conf.
	This means that the build tooling will not look in the module cache anymore it will
	look inside the vendor folder.
	We will also be making sure that we are pushing the vendor code up with the repo.

	Benefits of vendoring:
	1. Since the code is in the project, if we want to read the code we have
	access to it.
	2. If we want to debug through this code, we can. Debugger can walk through it.
	3. We can add log statements to the 3rd party code.
	4. We can modify/hack it.
	5. Once we do the "git pull" we have all the source code.




*/
