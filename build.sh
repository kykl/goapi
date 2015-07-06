RUNMODE=$1
TAG=${RUNMODE:-latest}

if [ ! -z "${RUNMODE}" -a -d "${RUNMODE}-conf" ]
then 
	cp -rf "${RUNMODE}-conf"/* conf
fi

docker build -t goapi:$TAG .

