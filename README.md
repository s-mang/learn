## Learn Go by Testing!

### NOTE:
**Read all the comments!** I wrote them just for you. 
I think you'll learn some stuff from them. The code too!

Install with:
```
go get github.com/adams-sarah/learn
```

Before we start, please `cd` into this directory.
```shell
cd $GOPATH/src/github.com/adams-sarah/learn
```

### Step 0: Go Test tool
During these exercises, you will need to use a few of the `go test` tool's 
great features.
So before we get started, I want you to look through the `go test` documentation
really quick. Just a skim is fine - as interested as you are.

To look at the documentation, run on your command line:
```
go test -help
```

Scroll up a bit, it's a few pages long.

##### Go Test Cover
In the first exercie, you will need to use the go test coverage tool.
Look in the documentation for the notes on the "-coverprofile" flag.

Then run the test command with the coverprofile flag set:
```shell
go test -coverprofile=cover.out
```

Your test should run, and you should see a statement:
`coverage: 65.4% of statements`.
This statement will exist when the -cover flag is present for the `go test` command.
As you read in the documentation, when the -coverprofile flag is set,
the -cover flag is automatically set to true for you.

Now please look in your working directory:
```
ls
```
you should see the new coverprofile file, 'cover.out'.

To read this file, we need to use the `go cover` tool.
Let's look at the documentation.

```shell
go tool cover -help
```

The `go cover` program will take a coverage profile like we just generated from our tests (cover.out)
and display it in a few different formats.
We will be using the HTML format, in a web browser.

Let's pull it up.

```shell
go tool cover -html=cover.out
```

You should see a black/red/green web page with some code on it.

In the drop down menu, you will see a few different files. These are all the files in 
this package.

To set us up for the first exercise, please select the `piglatin.go` file from the drop-down menu.

**Alright! You're ready for the exercises.**

### Step 1: Pig Latin
The first exercise is to write all the test cases
for the Pig Latin translation program I've written.

You can assume all the code I wrote is correct and complete.

The goal of this exercise is to help you get comfortable thinking of 
all the possible things that could go wrong.
You can use the coverage tool to help you, but remember that once you are
at 100% covereage for the file, that doesn't mean you've thought of all
the edge cases. 

**I found 8 test cases. How many can you find?**

Add more test cases to the `pigLatinTestCases` variable.


Be sure to also play around with the tests to find out how things work!