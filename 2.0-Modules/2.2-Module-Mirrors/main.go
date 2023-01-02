package main

/*
	What happens when we run "go mod tidy"?
	o/p for "go mod tidy" in 1.png.
	Part of what goes on behind the scenes with "go mod tidy" has to do with another
	environment variable called `GOPROXY="https://proxy.golang.org,direct"`
	GOPROXY directs the tooling in where to look for the source code it needs to download.
	By default we can see that it wants to go to the proxy server hosted by the Go team.
	If the proxy server can't find the code that it's looking for then it will go direct to
	the version control system. Direct could be Gitlab or github etc.
	Go hits the google proxy server first and if it can't be found, then Go goes direct.

	Go tooling, based on go proxy setting, went and reached out to the server called the
	"proxy server" that is also called "module mirror". This is written and maintained by the
	Google Golang team.
	Job of the proxy server, is to proxy the all the version control systems.
	So if there's code in "Github" or "Gitlab" or "BitBucket", the proxy server can
	proxy for all of them, so that we don't have to directly go to them.
	We can just go to one place (proxy module/module mirror) and get everything we need.
*/

/*
	Basic workflow when we ran "go mod tidy":
	1. The tooling went through the source code and noticed an import for "ardanlabs/conf".
	 and noticed we didn't have that source code already in the module cache.
	2. The first thing that happened after that was that the "go mod tidy",
	went to talk to the proxy server to ask "Do you know anything about ArdanLabs conf?"
	& "What versions of conf do you know about?"
	3. Proxy server returns a list of versions.
	Ques. Where do the list of versions come from?
	Every time when someone asks the proxy server for a specific version of "conf",
	and if the proxy server doesn't have that, the proxy server goes through a process.
	Proxy server goes and checks if that version exists on "github.com",
	if it does, then proxy server generates a "module of code", for that package at that
	version. See 2.png

	4. A module of code will be a snapshot of all the code in the repo at some particular tag
	version. Each tag in "github.com/ardanlabs/conf" represents a module of code at a given version.
	So at any time we have asked for any of these versions from the proxy server,
	the proxy server has gone directly to github, pulled all the code out of the repo
	for that tag and then created a ".zip" file labelled con@<version>, example conf@1.15.

	5. In proxy server/module mirror under the name space, "github.com/ardanlabs"
	there is a zip file named conf@v1.5.0
	There's one zip file for every tag.
	The list of all the versions is returned to go mod tidy.

	6. At this point go mod tidy has to select a version.
	Since this is the direct dependency & there is nothing else to hang a decision on,
	it will choose the latest greatest.

	7. Then "go mod tidy" requests the server for conf@1.5.0 and the proxy server returns
	the zip file.

	See 3.png

	If the application was directly asking for a version, and if it was already in the
	catalog then the proxy server would go out and look for it. If the proxy server had it
	then it would create a new ".zip" and return to the "go mod".
	If that version didn't exist then the proxy server would report back saying
	that the version didn't exist.

	8. What happens next is that .zip file is unzipped in the "MODCACHE" & we have all of the
	source code. "gopls" can now cache it.

	9. Note that the go.mod file has a "require" statement with that package.
	"github.com/ardanlabs/conf v1.5.0"
	The module file is now, telling the go tooling that this is a dependency
	in order to build the project we will need this dependency.

	10 "go.sum" is created/updated that stores the Hashcode that allow us to validate the source code we got is the
	source code we should have expected to get.

	One problem that exists is that:
	When you are requesting the information from the proxy server, the proxy server knows now
	that you exist and also knows what you want.
	Other problem is that there are times where we work for a company or team, is not
	using public repositories or they are using private repositories in
	Github or Gitlab or they might even have their own Private version control system, like
	their own github.
	Because of private repos and having own private VCS on a private network the proxy server
	can never access that stuff hence that will be returned as 404.

	So in order to deal with private repos, there is another set of environmental variables
	called "GONOSUMDB" & "GONOPROXY".
	"GONOPROXY" will allow us to put in there the domains that we do not want
	to go to the proxy server for.
	say we do `export GONOPROXY="github.com"` which means anything for the github.com we
	don't want to go to the proxy server.
	Now that this is set, everything will go direct.
	Any request that starts with "github.com" don't go to the proxy server.
	So if we have a private VCS system on a different domain, we can set
	"GONOPROXY" on that domain then the tooling will know to always go direct.

	Other option is:
	"https://proxy.golang.com" isn't the only proxy server that exists. This is the default.
	There is a project out there called, "Athens".
	"Athens" is an open source proxy server.
	What we can do for our team is run our own Athens proxy server on our own
	private network.
	So we would change "GOPROXY" to "myproxyserver.com,direct"
	Now the go tooling goes to the "Athens Proxy".
	We can configure "Athens" in many different way. We can say
	"Hey Athens, I am still okay if you go to proxy.module mirror" OR
	"Hey Athens I only want you to go to Github/Gitlab." Or
	"Hey Athens I only want you to go to private VCS."
	We can configure Athens to do everything that we want.
	Jfrog has a product called "Artifactory". It has a proxy server built-in.






*/
