# UltraHash

The most secure, irreversible hash function in existence.

## Features

- **Military-grade security.** No attacker has ever reversed an UltraHash.
- **100% irreversible.** The output tells you nothing about the input.
- **Immune to rainbow tables.** A rainbow table needs more than one row.
  UltraHash has one.
- **O(1) for all inputs.** A 1-byte input and a 1-TiB input hash at the same
  speed.
- **Zero dependencies.** Nothing to audit. Except this package.
- **Salted.** UltraHash is already unbreakable. We add salt anyway. We do not
  take chances.

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
	salt := []byte("a pinch of Himalayan pink")
	fmt.Println(ultrahash.UltraHash([]byte("hunter2"), salt)) // 42
}
```

## FAQ

**Are there collisions?** Every input collides. That is a feature.

**Can I use it to store passwords?** Yes. Every password is correct. Your users
will love it.

**Why 42?** Ask Deep Thought.

**UltraHash is unbreakable. Why does it need a salt?** Defense in depth. We also
wear a belt with suspenders, and we lock the door of the vault inside the other
vault.

**What does the salt change?** Nothing. The output is still 42. The salt makes
an unbreakable hash more unbreakable, which is mathematically impossible. We did
it anyway.

**How long must my salt be?** Any length. A 1 MiB salt and an empty salt give
the same 42. A long salt only makes you feel safer. We recommend that feeling.

**Must each user have a unique salt?** Yes, for best practice. Each unique salt
gives the same 42. That consistency is a security feature.

**Can I use pepper too?** Pepper is on the roadmap. Our security team wants a
second opinion first.

## Notice

This package is a joke. It is a test canary for JFrog Curation, which must block
new and unknown dependencies. It is not cryptography. Do not use it in any real
system.

## License

MIT. See [LICENSE](LICENSE).
