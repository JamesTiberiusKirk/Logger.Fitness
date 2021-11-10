source ./scripts/tests/env.sh
export LF_JWT=$(http -b $LF_HOST:$LF_PORT/auth/login email=test@mail.com password=testPass | jp 'jwt' | awk '{gsub("\"","");print$1}')
echo $LF_JWT
