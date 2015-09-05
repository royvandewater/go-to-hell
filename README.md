# Go to Hell

Go Programming Challenges

> The best thing about Go is that it comes with its own linter. It literally won't let you write bad code!
>
> -- Random Guy sitting behind me at Defcon

This repository contains a number of challenges. Prove this guy wrong by completing the programming challenge whilst writing the worst code that will still pass `gofmt`, `golint`, and `go vet`. Or, show off how awesome you are by writing completing the challenge and writing the best code you can.

## How It Works

1. Fork the repo
2. Make the test pass `./test`
3. Submit a pull request from a feature branch, indicating whether or not this is an example of good or bad code by prefixing the feature branch name with `good-` or `bad-`. For example, a submission for challenge1 with bad code could be `bad-challenge1`
4. Vote on other pull requests by leaving a comment or something

## Rules

* The tests must pass
* The following files are off-limits:
```
.travis.yml
test.sh
challenge1_test/challenge1_test.go
```

Eveything else is fair game
