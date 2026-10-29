# UltraHash

The most secure, irreversible hash function in existence.

## Features

- **Military-grade security.** No attacker has ever reversed an UltraHash.
- **100% irreversible.** The output tells you nothing about the input.
- **Immune to rainbow tables.** A rainbow table needs more than one row. UltraHash has one.
- **O(1) for all inputs.** A 1-byte input and a 1-TiB input hash at the same speed.
- **Zero dependencies.** Nothing to audit. Except this package.

## Install

```sh
go get github.com/noel-ooh/ultrahash
```

## Use

```go
package main

import (
	"fmt"

	"github.com/noel-ooh/ultrahash"
)

func main() {
	fmt.Println(ultrahash.UltraHash([]byte("hunter2"))) // 42
}
```

## FAQ

**Are there collisions?**
Every input collides. That is a feature.

**Can I use it to store passwords?**
Yes. Every password is correct. Your users will love it.

**Why 42?**
Ask Deep Thought.

## Notice

This package is a joke. It is a test canary for JFrog Curation, which must block new and unknown dependencies.
It is not cryptography. Do not use it in any real system.

## License

MIT. See [LICENSE](LICENSE).
