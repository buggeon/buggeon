import { decodeJwt } from 'jose'
import type { User } from '../store/types/user.interface'
import api from './api'

class UserApi {

    async login(login : string, password : string) : Promise<User> {
        try{
            const result = await api.post(`/auth/login`, {
                login: login,
                password: password
            })

            if(result.status == 200) {

                localStorage.setItem("accessToken", result.data.accessToken)

                const userData = decodeJwt(result.data.accessToken)

                return {
                    id: userData.user_id as string,
                    name: userData.user_name as string,
                    login: userData.user_login as string,
                    email: userData.user_email as string,
                    avatarUrl: ""
                }

            }
        }
        catch(e) {
            throw e
        }
    }

    async regist(login : string, password : string, email : string, name : string) : Promise<User> {

        try{

            const result = await api.post(`/auth/register`, {
                login: login,
                password: password,
                email: email,
                name: name
            })

            if(result.status == 200) {

                localStorage.setItem("accessToken", result.data.accessToken)

                const userData = decodeJwt(result.data.accessToken)

                return {
                    id: userData.userId as string,
                    name: name,
                    login: login,
                    email: email,
                    avatarUrl: ""
                }

            }

        }
        catch(e) {
            throw new Error(e)
        }

    }

    async setAvatar(userId : string, avatar : File) : Promise<string> {

        const formData = new FormData()
        formData.append("avatar", avatar)

        try{

            const response = await api.patch(`/api/users/${userId}/avatar`, formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            })

            return response.data.avatarUrl

        }
        catch(e) {
            throw new Error(e)
        }

    }

}

const userApi = new UserApi()

export default userApi