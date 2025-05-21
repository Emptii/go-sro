{
  description = "Go development environment (latest from nixpkgs)";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable"; # or stable if you prefer
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};
    in {
      devShells.default = pkgs.mkShell {
        buildInputs = [
          pkgs.go
        ];
        shellHook = ''
          echo "Go development environment ready with Go ${pkgs.go.version}"
        '';
      };
    });
}
