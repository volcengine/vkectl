
$ErrorActionPreference = 'Stop';

$packageArgs = @{
  packageName   = $env:ChocolateyPackageName
  unzipLocation = "$(split-path -parent $MyInvocation.MyCommand.Definition)"
  url           = 'https://github.com/volcengine/vkectl/releases/download/v0.1.0/vkectl-windows-v0.1.0-386.zip'
  url64bit      = 'https://github.com/volcengine/vkectl/releases/download/v0.1.0/vkectl-windows-v0.1.0-amd64.zip'

  softwareName  = 'vkectl*'

  checksum      = '7ADE3153D8772C8115DADE4AAD45B55983CB0A6D5054BCA0927AA14EC8CA8FCC'
  checksumType  = 'sha256'
  checksum64    = '376BCB486F22E6CA460465AE4686BC83DE66C89CC9C7200437F651CB13C7E3E6'
  checksumType64= 'sha256'

  silentArgs    = "/qn /norestart /l*v `"$($env:TEMP)\$($packageName).$($env:chocolateyPackageVersion).MsiInstall.log`""
  validExitCodes= @(0, 3010, 1641)
}

Install-ChocolateyZipPackage @packageArgs



















