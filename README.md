# EFA
EFA ( Extreme Fabric Automation ) CLI Mock Tool  

cd EFA/cli

env GOOS=linux GOARCH=amd64 go build -o efa

-----
docker run -d -it -p 222:22 --name efa --hostname "(efa:extreme)extreme@tpvm1" efa:0.0.4