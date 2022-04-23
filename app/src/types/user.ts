export type UserStore = {
    id: string,
    jwt: string,
    username: string,
    email: string,
    profile_picture: string,
    provider: string,
    provider_id: string,
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