export type UserStore = {
    id: string,
    jwt: string,
    username: string,
    email: string,
    avatar: string,
}

export type UserLoginDTO = {
    email: string,
    password: string,
}

export type UserRegisterDTO = UserLoginDTO & {
    username: string,
}

export type UserAuthHeader = {
    'x-access-token': string
}