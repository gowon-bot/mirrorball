echo "Pulling latest code from Github"
git pull

echo "Building image"
docker build . -t mirrorball

echo "Restarting service"
cd $1
docker-compose up --force-recreate --no-deps -d mirrorball