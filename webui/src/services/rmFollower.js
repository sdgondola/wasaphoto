import api from './axios'

import {
    BadFollowOperation,
    BadAuthException,
    InternalServerError,
    BlockedException,
    UserNotFoundException
} from './apiErrors'
import getLoginCookie from './getLoginCookie'

export default async function rmFollower(toRm) {
    const uid = getLoginCookie();
    if (uid == null) throw BadAuthException;
    let resp = await api.delete(`/users/${uid}/followers/${toRm}/remove`,
        { "headers": { "Authorization": `bearer ${uid}` } }
    );
    switch (resp.status) {
        case 204:
            return;
        case 400:
            throw BadFollowOperation;
        case 401:
            throw BadAuthException;
        case 403:
            throw BlockedException;
        case 404:
            throw UserNotFoundException;
        default:
            throw InternalServerError;
    }
}
