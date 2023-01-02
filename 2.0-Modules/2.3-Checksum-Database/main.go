package main

/*
	Let's try to understand the Hashcode in go.sum file.
	See 1.png.
	When the zip file of the right version came to us, we wrote to go mod.
	Then the go tooling generates 2 has codes. It creates a hash code for
	all the source code that unzipped & then it created a Hashcode specifically for
	"go.mod".
	Go tooling then took these hashcodes and a couple of things happened.
	Go tooling first asks "AM I writing these hashcodes, to the go.sum file for the
	first time?"
	If the answer is yes, we are writing the hashcodes to go.sum for the first time &
	"GONOSUMDB"(environmental variable) is either empty or not the same as the
	domain set for "GOPROXY",
	then there is a another round trip that happens to another service called
	"CheckSum database".
	There is only one "CheckSum database" service in existence which is run by the go team.
	Every time the "proxy server" creates a new module(new zip file) at some version,
	the "proxy server" generates the two HashCodes and "proxy server" writes the 2 HashCodes
	to the "CheckSum database".
	So for every module that is in the catalog at "proxy server" there are hashcodes
	related to that code in "Checksum DB".
	So what happens is, we downloaded the zip file in our app, we unzipped the source code,
	we made an entry in "go.mod", we generate the hashcode, and then the go tooling
	asks the "CheckSum DB" for the hashcodes that it has in it's system.
	Then we get those "HashCodes" from the checksum database and we compare them.
	If they are the same then "we are good".
	We know that we have the exact code that proxy server knows about/saw for the first time.

	If the "HashCodes" are already in the "go.sum" file, we don't take the extra trip.
	We can just take the "HashCode"(generated Hashcodes) and compare them to what's in "go.sum".

	But if they are not in "go.sum" we go to the "checksum DB" if we are allowed to based on
	"GONOSUMDB" flag.

	This gives us the durability, no matter where we pull the code from,
	direct, proxy, our own proxy, we always have to make sure that it's the same exact code.
	Nobody can go ahead and remove the tag change the code and put the tag back and we get
	the malicious code, and we had no idea that this is happening.





*/
