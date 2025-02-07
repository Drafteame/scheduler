{
  description = "Draft CLI tool";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        draftVersion = "0.0.0";
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "scheduler";
          version = draftVersion;

          subPackages = [ "." ];

          src = ./.;

          ldflags = [
            "-s"
            "-w"
            "-X 'github.com/Drafteame/scheduler/cmd/commands.Version=nix-v${draftVersion}'"
          ];

          env.CGO_ENABLED = false;
          env.GOWORK = "off";

          vendorHash = "";

          meta = with pkgs.lib; {
            description = "Cron like application to schedule process that runs in background";
            mainProgram = "scheduler";
            license = licenses.mit;
          };
        };

        apps.default = flake-utils.lib.mkApp {
          drv = self.packages.${system}.default;
        };

        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            sd
            fd

            go
            gotools
            goimports-reviser
            golangci-lint

            husky
            go-task
          ];
        };
      }
    );
}