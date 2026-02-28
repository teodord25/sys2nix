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
        vendorHash = null; # no deps yet
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
