# Reticulating Loading Messages...

Use the loading messages from The Sims in your program!

---

Want some garbage messages you can add to a loading screen? Congrats 🎉 you've found the right repo!

## I Just Want Messages!

It can be as simple as a `cURL`!

```bash
curl https://sims.willfantom.com/api/messages/random
```

See [`examples`](./examples) for more!

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

---

### TODO

1. Think of a better name for this project
2. MORE MESSAGES
3. Procrastinate by playing the Sims again
4. Stop working on this because its useless...

### Thanks

Thanks to [this post](https://modthesims.info/showthread.php?t=331011) for collating the messages!
