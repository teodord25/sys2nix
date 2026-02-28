# sys2nix

## development

requires [nix (the package manager)](https://nixos.org/download/)

- `nix develop` — dev shell with all dependencies
- `nix build` — build the package (creates `result` symlink)
- `nix run` — build and run

or just manage deps yourself and use `go build/run/fmt`
