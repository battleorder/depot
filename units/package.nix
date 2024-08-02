{ buildGoModule }:
buildGoModule {
  name = "battleorder-units";
  version = "v0.0.1";
  src = ../.;
  ldflags = [ "-s" "-w" ];
  vendorHash = "sha256-zgZgFbFLRuYLYUAaXUbD8iZj39oPoC2Erxl5nEYnsVw=";
  CGO_ENABLED = false;
  env.GOWORK = "off";
  modRoot = "./units";
}
