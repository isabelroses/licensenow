{
  lib,
  buildGoModule,
}:
buildGoModule {
  pname = "licencenow";
  version = "0.1.0";

  src = ./.;

  vendorHash = "sha256-aVk0ccByDS4+gs2im4eU6S5daK3OVoRYoBxn3SSgDGw=";

  ldflags = ["-s" "-w"];

  meta = {
    description = "licencenow gets you a license for your project";
    homepage = "https://github.com/isabelroses/licencenow";
    license = with lib.licenses; [mit];
    maintainers = with lib.maintainers; [isabelroses];
    platforms = lib.platforms.all;
  };
}
