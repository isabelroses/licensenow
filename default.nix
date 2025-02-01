{
  lib,
  buildGoModule,
}:
buildGoModule {
  pname = "licencenow";
  version = "0.2.0";

  src = ./.;

  vendorHash = "sha256-Aev19JLghMqSs/GKIVSzrol8aLSjufHHLpjqPTaFQ88=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "licencenow gets you a license for your project";
    homepage = "https://github.com/isabelroses/licencenow";
    license = with lib.licenses; [ mit ];
    maintainers = with lib.maintainers; [ isabelroses ];
    mainProgram = "licencenow";
  };
}
