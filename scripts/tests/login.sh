source ./scripts/tests/env.sh
export LF_JWT=$(http -b $LF_HOST:$LF_PORT/api/auth/login email=test@mail.com password=testpass | jq .jwt | awk '{gsub("\"","");print$1}')
echo $LF_JWT
