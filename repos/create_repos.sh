REPOS=$1
TOKEN=$2

echo "{\"name\":\"$REPOS\",\"auto_init\":false}" > tmp.json

curl -H "Authorization: bearer $TOKEN" -X POST \
    -d @tmp.json \
    https://api.github.com/user/repos

rm tmp.json
