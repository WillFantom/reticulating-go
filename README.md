# Reticulating Loading Messages...

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/willfantom/reticulating-go/Publish%20Rolling%20Release) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/willfantom/reticulating-go) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/willfantom/reticulating-go?sort=semver)

Use the loading messages from The Sims in your program!

![example](./example.gif)

---

Want some garbage messages you can add to a loading screen? Congrats 🎉 you've found the right repo!

## I Just Want Messages!

  1. Use the CLI tool (that can also run the API) on MacOS
     ```bash
      brew tap willfantom/reticulate
      brew install reticulate
      reticulate --help
     ```

  2. It can be as simple as a `cURL`!
      ```bash
      # From any game
      curl https://sims.willfantom.com/api/messages/random
      # From just the sims 2
      curl https://sims.willfantom.com/api/messages/sims2/random
      # From just the sims 2 university life
      curl https://sims.willfantom.com/api/messages/sims2/university/random
      ```

See [`examples`](./examples) for more uses!

---

## Add to Go Program

You can just add this to your golang program:

```bash
go get github.com/willfantom/reticulating-go
```

Example:
```go
package main

import (
  "fmt"

  "github.com/willfantom/reticulating-go"
)

func main() {
  fmt.Println(reticulating.GetLoadingMessage())
}
```

---

## I Wanna Host The API Myself...

Just use the built Docker image 🐳

```bash
docker run --rm -p 8080:8080 ghcr.io/willfantom/reticulating:latest
```

Or the CLI tool

```bash
reticulate api --port 8080
```

---

### TODO

1. Think of a better name for this project
2. MORE MESSAGES
3. Procrastinate by playing the Sims again
4. Stop working on this because its useless...

### Thanks

Thanks to [this post](https://modthesims.info/showthread.php?t=331011) for collating the messages!
