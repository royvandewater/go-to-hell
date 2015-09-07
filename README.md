# Go to Hell

Go Programming Challenges

> The best thing about Go is that it comes with its own linter. It literally won't let you write bad code!
>
> -- Random Guy sitting behind me at Defcon

This repository contains a number of challenges. Prove this guy wrong by completing the programming challenge whilst writing the worst code that will still pass `gofmt`, `golint`, and `go vet`. Or, show off how awesome you are by writing completing the challenge and writing the best code you can.

## How It Works

1. Fork the repo
2. List the challenges you're working on in submissions.txt:
  ```shell
  echo "challenge1" >> submissions.txt
  echo "challenge2" >> submissions.txt
  ```
3. Make the test pass `./test`
4. Submit a pull request, indicating whether or not this is an example of good or bad code by prefixing `[BAD]` or `[GOOD]` to the pull request title
5. Vote on other people's pull requests by leaving a comment or something

## Rules

* The tests must pass
* The following files are off-limits:
```
.travis.yml
test.sh
challenge1_test/challenge1_test.go
```

Eveything else is fair game
