{
  lib,
  buildGoModule,
}:
buildGoModule {
  pname = "licencenow";
  version = "0.0.1";

  src = ./.;

  # lib.fakeSha256 should be used to when deps update, but its not working for me so im leaving this here
  #  sha256-0000000000000000000000000000000000000000000=
  vendorHash = "sha256-aVk0ccByDS4+gs2im4eU6S5daK3OVoRYoBxn3SSgDGw=";

  ldflags = ["-s" "-w"];

  preBuild = ''
    substituteInPlace main.go \
      --replace "./templates" "$out/share/templates" \
  '';

  postInstall = ''
    mkdir -p $out/share
    cp -r templates $out/share/templates
  '';

  meta = {
    description = "licencenow gets you a license for your project";
    homepage = "https://github.com/isabelroses/licencenow";
    license = with lib.licenses; [mit];
    maintainers = with lib.maintainers; [isabelroses];
    platforms = lib.platforms.all;
  };
}
