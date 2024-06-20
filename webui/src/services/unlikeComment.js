import {
    AccessDeniedException,
    BadAuthException,
    InternalServerError,
    LikeImpersonationException
} from './apiErrors';
import { authStatus } from './login';
import api from './axios'

export default async function unlikeComment(cid) {
    if (authStatus.status == null) throw BadAuthException;
    let resp = await api.delete(`/comments/${cid}/unlike/${authStatus.status}`,
        { "headers": { "Authorization": `bearer ${authStatus.status}` } });
    switch (resp.status) {
        case 201:
            return;
        case 401:
            throw BadAuthException;
        case 403:
            throw AccessDeniedException;
        case 404:
            throw LikeImpersonationException;
        default:
            throw InternalServerError;
    }
}
