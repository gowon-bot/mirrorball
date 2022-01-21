git pull

docker build . -t mirrorball

cd $1
docker-compose up --force-recreate --no-deps -d mirrorball