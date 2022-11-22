# KryoSprite

A sprite library for Ebitengine. Kryosprite helps ease handling your sprite's texture, drawing, updating, and animations in one package. 

## Check Out Examples

I highly recommend checking out the [example](./examples/) code for a full demonstration.

```bash
go run -tags=example github.com/KalebHawkins/kryosprite/examples/basic@latest
```

```bash
go run -tags=example github.com/KalebHawkins/kryosprite/examples/fox@latest
```

## Using the Library

First you need to add the library to your project.

```bash
go get github.com/KalebHawkins/kryosprite
```

Then import it.

```bash
import (
    ks "github.com/KalebHawkins/kryosprite"
)
```

See the wiki for a walkthrough on how to build a simple game from scratch using Ebitengine and Kryosprite.

## TODO

- [ ] Make Wiki...
- [ ] Add loading and parsing images from JSON-exported sprite sheets.
  - [ ] LibreSprite/Aesprite
  - [ ] Others? (Open an issue for suggestions)
- [ ] Feedback? (Open an issue for feedback)
- [ ] Contributions? (Open a pull request)
- [ ] Create developer documentation and guides for issues and pull requests.