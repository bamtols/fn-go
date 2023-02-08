TAG=$1

git tag -d "$TAG"
git push --delete origin "$TAG"

git tag "$TAG"
git push origin "$TAG"