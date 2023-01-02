package main

/*
	Whenever we run "go mod tidy", we should run
	"go mod vendor" as well.
	When we run go mod tidy sometimes we see the word "incompatible".
	See 1.png.
	Even in the go.mod file. See 2.png.
	We didn't get it with the conf package, but we got with the "httptreemux".
	When Bill shot the vide the latest greatest was 5.3.0 but what was pulled down was
	5.0.1 (even in my case 5.0.1 was pulled and it said incompatible).

	For some reason go tooling chose 5.0.1 as the version of the repo/module.
	And in 5.0.1 there is no "go.mod" file.
	At 5.0.2 there is a go.mod file.
	The tooling for some reason chose the latest tag that didn't have a module file.
	If we look closer, the conf package is v1.5.0 - v1
	but the "httptreemux" is v5.

	If we take a look at the latest greatest tag version which is v5.3.0 and if we take a look
	at the go.mod file. See 3.png.
	Notice that the name of the module has "/v5" extension.
	Note in 3.png we have "github.com/dimfeld/httptreemux/v5"

	Go supports the idea that you can build against different major versions of a module.
	We could have situations where we are bringing in 2 different direct dependencies. Each
	requiring an API set for different major versions of this dependency.
	The way Go supports this is by name-spacing.
	If in 3.png it would have been module "github.com/dimfeld/httptreemux"
	(like we see in conf) then we assume that we are talking about v0 or v1 because we don't
	have to add that for backwards compatibility.
	In other words in this main file when we added in the import
	"github.com/dimfeld/httptreemux" we were telling the tooling that we want the
	latest and greatest version of major version1(v1) and not version 5.

	Here now that we see that 5.0.1 doesn't have module support, the go tooling said
	"I am going to treat 5.0.1 as the latest and greatest version of v1" because there is
	no module support and we need to choose one.

	Hence in go.mod we see,
	"github.com/dimfeld/httptreemux v5.0.1+incompatible"
	this is incompatible because we are asking for latest v1 and we have got some sought of v5.

	Question. So how do we get the tooling to give us the latest & greatest which is v5?
	We have to match the namespace in the module.

	Hence in the main.go, we ask to get us the latest greatest v5
	"github.com/dimfeld/httptreemux/v5"
	Next we do run "go mod tidy", "go mod vendor"
	And now in "go.mod" we see the v5.3.0 which is the latest and greatest.
	See 4.png and 5.png

	Question. What do we need to do when we look at a module/3rd party package that we
	want to bring up?
	We go to the repo in github where the package is and we look at the tags.
	If we see that the tags are going beyond, "v1" marker we go into the latest and greatest
	tag in this case "v5.3.0" and validate that the developer (who created the dependency)
	See 6.png, is using the right naming convention/naming space for that "v"
	As here "github.com/dimfeld/httptreemux/v5".

	MVS Algorithm - Minimum Version selection.
	How we can direct the MVS algorithm, to maybe upgrade these dependencies overtime?

	Say, we are building our App and our App has a direct dependency on "A".
	Say A's version v1.2
	Say in the module file for A it says that A depends on the module "D" version v1.3.2
	And say when we look at the repo D, it says that D is at version
	v1.9.1 (latest greatest of D).
	See 7.png

	So we can't build "A" unless we have "D" &  we know "D" v1.3.2 is compatible to do it.
	Latest greatest is v1.9.1.
	Imagine you are the build tool, and we want to build the app against A1.2, but what version
	of "D" are you going to pick/choose?
	You technically have 2 choices,
	We can use 1.9.1 or maybe use 1.3.2.
	Tis is where controversy comes in,
	if we are using traditional version selection tools based on algorithms that
	are called SAT solvers.
	SAT solvers want to pick the latest greatest version of a dependency.
	So if we were doing this build with SAT, most likely the compiler will choose
	"D" v1.9.1 & we might think this is what we want since it is the most stable version of the
	package and it has to be the most secure version.
	Problem is this is now what Go's tooling is going to do.
	Go uses the algorithm called "MVS", minimum version selection.
	It takes a different approach, it has a different opinion.
	Go's tooling says 1.9.1 isn't necessarily the most stable code for building this app
	because "A" is reporting that "A" knows that it is compatible with D v1.3.2
	- assuming all the test ran on A, against D's v1.3.2 - this is almost a guarantee that
	you will not have any problems.

	We have to remember that version semantic - v1.3/v1.9.
	We are assuming that the developer held the social contract and that the API
	didn't change. We are assuming that A can build against 1.9.1 and the build wouldn't break
	but that's an assumption.
	We would like to think that to A's v1.2 will build with D's v1.3.2 as a guarantee and that
	D's v1.9.1 is not a guarantee that our app will build.
	So the MVS algorithm says that "I want the guarantee".
	So what we will do is not choose D's v1.9.1 but choose D's 1.3.2 and that is what will be
	built for the app.
	See 8.png

	Now say we now we are building the app a bit more nad we bring in B's version v1.4
	and B also uses D, but it uses D at v1.5.2, see 9.png
	What version of D are you going to select?
	If you are SAT then you would select v1.9.1
	but if you are MVS then this is about greatest but not necessarily latest.
	Which version of D are going to choose, we have 3 versions of D.
	v1.9.1, v1.3.2 abd v1.5.2

	We will pick the latest version based on the potential versions, listed in A's and B's
	mod file.
	So in this case when we build the app we will be building it on D's v1.5.2
	See 10.png.
	If we are no longer using B's dependency on D v1.5.2 then as well
	MVS will be building against 1.5.2. See 11.png
	There is no historical document here, we will just maintain version 1.5.2
	MVS - Latest but not necessarily greatest.
	For direct dependency like A and B we will be using latest and greatest,
	for indirect dependency we wouldn't necessarily be doing that.

	If someone wants to use latest and greatest all the time the command we can use is,
	"go get -u -d -v ./..."
	What "go get" does is with with the above options, walks through the project tree,
	telling go get to essentially update every single dependency to it's latest and greatest.
	Which means is we would now use D's version v1.9.1

	* Go also solves the diamond dependency issue.
	See 12.png
	Say A is using D and B is using D.
	If everything is still in the same major version, hopefully we'll get lucky
	that APIs can still talk and the application doesn't break.

	what is A uses D's version1 and B uses D's version2.
	If B is using upgraded version2 major version change of D as opposed to A,
	we can't just pick one, because changing the major version we are saying that these
	APIs (D's v1 and D's v2) are different, there is something major changed.
	We can't just assume that B will build against D's v1 and A will build with D's v2.

	We need to have a way to be able to at major version level maintain the clear verticals.
	See 13.png.
	This is what we say with this notation,
	"github.com/dimfeld/httptreemux/v5 v5.3.0" what happens is on disk we are able to maintain
	the clear paths of major version.
	See 14.png
	We can see that we have different versions of the modules that we brought in.
	When we look inside "httptreemux" folder see 15.png,
	we see the "v5@v5.3.0", so the URL structure,
	"github.com/dimfeld/httptreemux/v5 v5.3.0" is following back on the disk,
	which means we can have v4 and v3 as well.
	Those codes sit in their own folders.
	Which means we can do 13.png in the build where we can do this
	in go.mod
	"github.com/dimfeld/httptreemux/v5 v4.1.0"
	"github.com/dimfeld/httptreemux/v4 v5.3.0"
	we have some code using the v4 version and some code using the v5 version.
	As long as that code imports that right version the  compiler can build it for us.

	Modules solve for us such problems as well.
*/

import (
	//"github.com/dimfeld/httptreemux"
	"github.com/dimfeld/httptreemux/v5"
)

func main() {
	httptreemux.New()
}
