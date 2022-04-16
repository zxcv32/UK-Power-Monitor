ref="$GITHUB_REF"
version=$(head -n 1 project.version)
source=$(echo $ref | cut -d / -f2-2)
tagConversion=$(echo $ref | cut -d / -f3- | tr  '/' '-')

if [[ "$source" == "heads" ]]
then
  if [[ "$tagConversion" == "main" ]]
  then
    tag="$version"
  else
    tag="$version-$tagConversion"
  fi
elif [[ "$source" == "tags" ]]
then
  tag="$version-$tagConversion"
elif [[ "$source" == "pull" ]]
then
  pr=$(echo $ref |  cut -d / -f3-3)
  tag="$version-rc$pr"
else
   tag="$version-unknown"
fi

echo "TAG=$tag"

