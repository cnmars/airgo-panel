
version='dev'

npm cache clean -f
npm install npm -g
npm install n -g
n v20.9.0

cd ../../web
sed -i 's/old-version/${version}/g' ./src/layout/footer/index.vue
npm i
# 前端打包
npm run build
# 将打包文件移动到后端，嵌入到go编译的二进制文件中
rm -rf ../server/web/web
mv web ../server/web/


cd ../server
sed -i 's/old-version/${version}/g' ./cmd/version.go
# 本机编译
CGO_ENABLED=1 go build -o AirGo -ldflags='-s -w' main.go