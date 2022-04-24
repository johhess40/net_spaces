#!/bin/bash
##Taking first input of org/user and repo such as johhess40/net_spaces
LOOKUP=$1

##Confirming the latest tag version for download
confirm_download(){
  LOC=$(curl --silent "https://api.github.com/repos/$1/releases/latest" | jq .tag_name)
  echo "$LOC" | tr -d '"'
}

make_download(){
  echo "Downloading net_spaces executable..."
  DOWNLOAD_VERSION=$(confirm_download "$LOOKUP")
  echo "$DOWNLOAD_VERSION"
  curl -L "https://api.github.com/repos/$1/zipball/$DOWNLOAD_VERSION" -o $DOWNLOAD_VERSION.zip
}

remove_zipball(){
  echo "Creating zip ball and then removing unnecessary files..."
  DOWNLOAD_VERSION=$(confirm_download "$LOOKUP")
  mkdir app
  unzip -o $DOWNLOAD_VERSION.zip  -d app
  mv -f ./app/*/.[!.]* ./app
  mv -f ./app/*/* ./app
  AUTH=$(curl --silent "https://api.github.com/repos/$1/releases/latest" | jq .author.login | tr -d '"')
  echo "$AUTH"
  cd ./app && ls | grep $AUTH | xargs rm -rf

}

make_download "$LOOKUP"
remove_zipball "$LOOKUP"
