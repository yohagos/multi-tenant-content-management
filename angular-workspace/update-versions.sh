VERSION=21.2.18

sed -Ei "s#\"@angular/([^\"]+)\": *\"[^\"]+\"#\"@angular/\1\": \"^${VERSION}\"#g" package.json

npm install
