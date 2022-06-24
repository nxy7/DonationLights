{ pkgs ? import <nixpkgs> {}
}:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
  ];
  shellHook = ''
    echo 'Shell started...'
  '';
}
