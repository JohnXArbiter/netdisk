import api from '../../../utils/apis/request.ts';
import type {Resp} from '../../../utils/apis/base.ts';
import type {File} from '../file.ts'

export function getFileDetailById(id: number) {
    return api.get<any, Resp<File>>(`/file/${id}`)
}