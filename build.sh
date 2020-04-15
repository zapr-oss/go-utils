#!/bin/bash
# Pull git go-utils/<version> code and build
if [ $# -lt 1 ]
then
	echo "Please provide the deployment severity. [major, minor, patch]\nSyntax: sh build.sh <severity>"
elif [ "$1" != "major" ] && [ "$1" != "minor" ] && [ "$1" != "patch" ]
then
	echo "Wrong argument provided for severity. Provide [major, minor, patch]"
else
	output_dir=bin
	(git reset --hard && git fetch -a && git checkout master && git pull --no-edit origin master && git pull --no-edit origin develop && git push origin master)
	if [ $? -eq 0 ]
	then
		current_tag=`git describe --abbrev=0 --tags`
		major=`echo $current_tag | cut -d . -f1 | cut -d v -f2`
		minor=`echo $current_tag | cut -d . -f2`
		patch=`echo $current_tag | cut -d . -f3`

		if [ "$1" = "major" ]
		then
			major=$((major+1))
			minor=0
			patch=0
		elif [ "$1" = "minor" ]
		then
			minor=$((minor+1))
			patch=0
		else
			patch=$((patch+1))
		fi

		new_tag="v$major.$minor.$patch"
		echo "Deploying new tag: $new_tag"

		(git tag $new_tag && git push origin $new_tag)
		if [ $? -eq 0 ]
		then
			echo "Done. Update the version number in dependent projects."
		else
			echo "Failure. Build script execution completed with failures."
		fi
	fi
fi
