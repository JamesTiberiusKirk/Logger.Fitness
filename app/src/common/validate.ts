
type ExpressionsType = {
    [key: string]: RegExp
}

const expressions: ExpressionsType = {
    email: /^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/,
}


function email(email: string): boolean {
    return expressions.email.test(email)
}

function isEmpty(data: string): boolean {
    return (data.length == 0 && data === '');
}

const Validate = {
    email,
    isEmpty,
}

export default Validate;