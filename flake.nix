{

 description = "flake for nix development"
 
 inputs = {
   nixpkgs = "github:nixos/nixpkgs/nixos-unstable";
 };

 outputs = { self, nixpkgs }:

 let 
    pkgs = nixpkgs.legacyPackages."x86_64-linux";
 in

 {
   devShells."x86_64-linux".default = pkgs.mkShell {
    
    packages = [ pkgs.go pkgs.python3 ];
   
 };

 };

} 
