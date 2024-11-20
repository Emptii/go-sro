{ pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  )
, mkGoEnv ? pkgs.mkGoEnv
, gomod2nix ? pkgs.gomod2nix
, goPkgs ? import (fetchTarball {
      url ="https://github.com/NixOS/nixpkgs/archive/334ec8b503c3981e37a04b817a70e8d026ea9e84.tar.gz"; 
      sha256 = "158lqyrb0w1fsfy79zdk16zpli72q7yq3fk51a1wfnfy7d4xbmwm";
      }) { system = "x86_64-linux";}
, neovim ? pkgs.neovim
}:

let
  goEnv = mkGoEnv { go=goPkgs.go_1_18; pwd = ./.; };
in
pkgs.mkShell {
  packages = [
    goEnv
    gomod2nix
    goPkgs.go_1_18
    neovim
  ];
}
