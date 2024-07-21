{
  inputs.nixpkgs.url = "github:nixos/nixpkgs?ref=nixpkgs-unstable";

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = inputs.nixpkgs.lib.systems.flakeExposed;

      perSystem = { config, pkgs, ... }: {
        devShells.default = pkgs.callPackage ./shell.nix { inherit config; };

        packages = {
          units = pkgs.callPackage ./units/package.nix { };
        };
      };
    };
}
