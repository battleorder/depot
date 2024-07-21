{ buildGoModule }:
buildGoModule {
  name = "battleorder-units";
  version = "v0.0.1";
  src = ./.;
  ldflags = [ "-s" "-w" ];
  vendorHash = "sha256-UCYYgHVw30sB2Hx3hR7Et+vVsHgJ4USFZAuovUuF5mU=";
  CGO_ENABLED = false;
}
