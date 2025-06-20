#!/bin/bash

helpFunction()
{
   echo ""
   echo "Usage: $0 -a paramARCH -o paramOS"
   echo -e "\t-a Enter the CPU architecture (amd64, x86_64, aarch64)"
   echo -e "\t-o Enter the Operating System targeted"
   echo ""
   exit 1 # Exit script after printing help
}

while [ "$1" != "" ]; do
    case $1 in
        -a )
            shift
            paramARCH=$1
        ;;
        -o )
            shift
            paramOS=$1
        ;;             
        -h )    helpFunction
            exit 1
        ;;
        * )              helpFunction
            exit 1
    esac
    shift
done

# Print helpFunction in case parameters are empty
if [ -z $paramARCH ] || [ -z $paramOS ]
then
   echo "Some or all of the parameters are empty";
   echo "Using environment variable from Docker";
   echo ""
   paramOS=$TARGETOS
   paramARCH=$TARGETARCH
   #helpFunction
fi

echo ""

# Begin script in case all parameters are correct
echo "OPERATING SYSTEM  = $paramOS"
echo "ARCHITECTURE      = $paramARCH"

version=$(curl -Ls -o /dev/null -w %{url_effective} https://github.com/apple/foundationdb/releases/latest)
version=${version##*/}

client_filename="foundationdb-clients_${version##*/}-1_$paramARCH"
server_filename="foundationdb-server_${version##*/}-1_$paramARCH"

echo "Latest version for FoundationDB [$version]"

echo ""

file_extension=""

if [ "$paramOS" == "linux" ]; then
  file_extension=".deb"
elif [ "$paramOS" == "darwin" ]; then
  file_extension=".pkg"
elif [ "$paramOS" == "win" ]; then
  file_extension=".msi"
else
  echo "Invalid OS exiting script"
  exit 1
fi

client_file="$client_filename$file_extension"
server_file="$server_filename$file_extension"

echo "FoundationDB Client filename : [$client_file]"
client_file_url="https://github.com/apple/foundationdb/releases/download/$version/$client_file"
echo "Download file url [$client_file_url]"

echo ""

echo "FoundationDB Server filename : [$server_file]"
server_file_url="https://github.com/apple/foundationdb/releases/download/$version/$server_file"
echo "Download file url [$server_file_url]"

echo ""

wget --directory-prefix=/tmp $client_file_url
apt install "/tmp/$client_file"

wget --directory-prefix=/tmp $server_file_url
apt install "/tmp/$server_file"

rm /tmp/*.deb