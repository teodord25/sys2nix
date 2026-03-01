{
  description = "sys2nix";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs =
    { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      packages.${system}.default = pkgs.buildGoModule {
        pname = "sys2nix";
        version = "0.1.0";
        src = ./.;

        # change this when deps change (checksum of project deps)
        vendorHash = "sha256-scg1vorrJ4a6pblnhEWJeLJjh60uv+PItVU7lCpLGxM=";
      };

      devShells.${system}.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          go
          gopls
          gotools
          go-tools
        ];
      };
    };
}
